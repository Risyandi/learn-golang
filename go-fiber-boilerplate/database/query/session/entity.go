package session

import "time"

type (
	UserSessionEntity struct {
		SessionID string    `json:"sessionId" bson:"session_id"`
		UserID    string    `json:"userId" bson:"user_id"`
		IsActive  bool      `json:"isActive" bson:"is_active"`
		IPAddress string    `json:"ipAddress" bson:"ip_address"`
		URI       string    `json:"uri" bson:"uri"`
		UserAgent string    `json:"userAgent" bson:"user_agent"`
		DeviceID  string    `json:"deviceId" bson:"device_id"`
		CreatedAt time.Time `json:"-" bson:"created_at"`
		UpdatedAt time.Time `json:"-" bson:"updated_at"`
		ExpiredAt time.Time `json:"-" bson:"expired_at"`
	}

	UserLogEntity struct {
		SessionID string      `json:"sessionId" bson:"session_id"`
		UserID    string      `json:"userId" bson:"user_id"`
		Method    string      `json:"method" bson:"method"`
		URI       string      `json:"uri" bson:"uri"`
		Request   interface{} `json:"request" bson:"request"`
		Response  interface{} `json:"response" bson:"response"`
		IPAddress string      `json:"ipAddress" bson:"ip_address"`
		UserAgent string      `json:"userAgent" bson:"user_agent"`
		CreatedAt time.Time   `json:"-" bson:"created_at"`
	}

	AdminLogEntity struct {
		UserId    string      `json:"userId" bson:"user_id"`
		AdminId   string      `json:"adminId" bson:"admin_id"`
		IP        string      `json:"ip" bson:"ip"`
		Method    string      `json:"method" bson:"method"`
		Uri       string      `json:"uri" bson:"uri"`
		UserAgent string      `json:"userAgent" bson:"user_agent"`
		Request   interface{} `json:"request" bson:"request"`
		CreatedAt time.Time   `json:"-" bson:"created_at"`
	}
)
