package user

// CreateUserRequest defines the expected fields for user sign-up.
type CreateUserRequest struct {
	Email          string   `json:"email" binding:"required,email"`
	Name           string   `json:"name" binding:"required,name"`
	Phone          string   `json:"phone" binding:"required,phone"`
	Role           []string `json:"role" binding:"required"`
	Password       string   `json:"password,omitempty"`
	ProfilePic     string   `json:"profilePic,omitempty"`
	TargetPlatform string   `json:"targetPlatform,omitempty"`
	DeviceToken    string   `json:"deviceToken,omitempty"`
	IsGoogleLogin  *bool    `json:"isGoogleLogin,omitempty"`
}
