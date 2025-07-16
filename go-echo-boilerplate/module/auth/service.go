package auth

import (
	"context"
	"errors"
	"time"

	"boilerplate/config"
	"boilerplate/constant"
	"boilerplate/entity"
	initDB "boilerplate/init"
	"boilerplate/pkg/helper"
	"boilerplate/pkg/utils"
	"boilerplate/schema/request"
	"boilerplate/schema/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthService struct {
	config    *config.Config
	database  *mongo.Database
	jwtHelper *helper.JWTHelper
}

func NewAuthService(config *config.Config) *AuthService {
	return &AuthService{
		config:    config,
		database:  initDB.MongoDB,
		jwtHelper: helper.NewJWTHelper(config.JWTSecret, config.JWTExpireHours),
	}
}

// RegisterUser handles user registration logic
func (a *AuthService) RegisterUser(ctx context.Context, req *request.RegisterRequest) (*response.LoginResponse, error) {
	// Get MongoDB collection
	client := a.database.Client()
	coll := client.Database(a.config.DatabaseName).Collection(constant.UsersCollection)

	// Check if user already exists
	var existingUser entity.User
	err := coll.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}
	if err != mongo.ErrNoDocuments {
		return nil, errors.New("database error while checking existing user")
	}

	// Hash password
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Set default role if not provided
	role := req.Role
	if role == "" {
		role = constant.RoleUser
	}

	// Create new user
	user := entity.User{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert user into database
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate JWT token
	token, err := a.jwtHelper.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
	)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Prepare response
	response := &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return response, nil
}

// AuthenticateUser handles user login logic
func (s *AuthService) AuthenticateUser(ctx context.Context, req *request.LoginRequest) (*response.LoginResponse, error) {
	// Get MongoDB collection
	client := s.database.Client()
	coll := client.Database(s.config.DatabaseName).Collection(constant.UsersCollection)

	// Find user by email
	var user entity.User
	err := coll.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("invalid email or password")
		}
		return nil, errors.New("database error")
	}

	// Verify password
	if !helper.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := s.jwtHelper.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
	)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Prepare response
	response := &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return response, nil
}

// GetUserByID retrieves user information by ID
func (s *AuthService) GetUserByID(ctx context.Context, userID string) (*response.UserResponse, error) {
	// Convert string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	// Get MongoDB collection
	client := s.database.Client()
	coll := client.Database(s.config.DatabaseName).Collection(constant.UsersCollection)

	// Find user by ID
	var user entity.User
	err = coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("database error")
	}

	// Prepare response
	response := &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

// GetUsers returns a list of users
func (s *AuthService) GetUsers(ctx context.Context, page int, limit int) (*response.GetUsersResponse, error) {
	// Get MongoDB collection
	client := s.database.Client()
	coll := client.Database(s.config.DatabaseName).Collection(constant.UsersCollection)

	// Count total documents
	total, err := coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("database error")
	}

	// Find users with pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	// Find users
	var users []entity.User
	cursor, err := coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, errors.New("database error")
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.New("database error")
	}

	// Prepare response
	userResponses := make([]*response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = &response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	// Create pagination metadata
	meta := utils.CreateMeta(page, limit, total)

	response := &response.GetUsersResponse{
		Users: userResponses,
		Meta:  *meta,
	}

	return response, nil
}

func (s *AuthService) UpdateUser(ctx context.Context, email string, user *request.UpdateUserRequest) error {
	// Get MongoDB collection
	client := s.database.Client()
	coll := client.Database(s.config.DatabaseName).Collection(constant.UsersCollection)

	// Update user
	result, err := coll.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": user})
	if err != nil {
		return errors.New("database error")
	}

	// Handle not found data
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
