package user

type CreateUserRequest struct {
	Email          string   `json:"email" binding:"required,email"`
	Name           string   `json:"name" binding:"required"` // <-- FIXED
	Phone          string   `json:"phone" binding:"required"`
	Role           []string `json:"role" binding:"required"`
	Password       string   `json:"password,omitempty"`
	ProfilePic     string   `json:"profilePic,omitempty"`
	TargetPlatform string   `json:"targetPlatform,omitempty"`
	DeviceToken    string   `json:"deviceToken,omitempty"`
	IsGoogleLogin  *bool    `json:"isGoogleLogin,omitempty"`
}
