package entity

var UserType string

const CtxAuthenticatedUserKey string = "CtxAuthenticatedUserKey"

const TimeZone string = "Asia/Ho_Chi_Minh"

const (
	UserTypeAdmin   = "admin"
	UserTypeStudent = "student"
	UserTypeTeacher = "teacher"
	UserTypeDriver  = "driver"
	UserTypeChef    = "chef"
)

const (
	AccessKey              = "KMS_jwt_access"
	RefreshKey             = "KMS_jwt_refresh"
	CookiePath             = "/api/v1"
	CookiePathRefreshToken = "/api/v1/auth/refresh"
	CookieHTTPOnly         = false
	CookieMaxAge           = 14400
)
