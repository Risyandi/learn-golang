package constants

type contextKey string

const (
	JwtClaimsKey contextKey = "userInfo"

	ClaimsUserSubscriptionIsActive         = "userSubscriptionIsActive"
	ClaimsUserSubscriptionLevel            = "userSubscriptionLevel"
	ClaimsUserSubscriptionName             = "userSubscriptionName"
	ClaimsUserSubscriptionProductLimit     = "userSubscriptionProductLimit"
	ClaimsUserSubscriptionCoLimit          = "userSubscriptionCoLimit"
	ClaimsUserSubscriptionProductLimitInCo = "userSubscriptionProductLimitInCo"
	ClaimsUserSubscriptionExpiredAt        = "userSubscriptionExpiredAt"
	ClaimsUserSubscriptionMemberLimit      = "userSubscriptionMemberLimit"
	ClaimsID                               = "id"
	ClaimsToken                            = "token"
	ClaimsIsSus                            = "isSuspended"
	ClaimsName                             = "name"
	ClaimsEmail                            = "email"
	ClaimsRole                             = "role"
	ClaimsParentID                         = "parentId"
	ClaimsAccess                           = "userAccess"
	ClaimsUserType                         = "userType"
	ClaimsTokenID                          = "jti"
	ClaimsOAID                             = "oaId"
	ClaimsSub                              = "sub"
	ClaimsUserRole                         = "userRole"
	ClaimsUserSubscriptionIsCsRotator      = "userSubscriptionIsCsRotator"
)
