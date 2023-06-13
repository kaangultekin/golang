package message

const (
	ErrEnvFailed              = "Error loading .env file"
	ErrConnectDBFailed        = "Failed to connect to database!"
	ErrConnectRedisFailed     = "Failed to connect to Redis!"
	ErrUsedEmail              = "Email has been used before."
	ErrUserNotFound           = "User not found."
	ErrEPNotFound             = "Endpoint not found."
	ErrValidationFailed       = "Validation failed."
	ErrFormNotFound           = "Form not found."
	ErrCompareHashAndPassword = "Password is incorrect."
	ErrUserStatusPassive      = "User's status is passive."
	ErrInvalidToken           = "Invalid token."
	ErrFailedLogout           = "Failed to logout."
)
