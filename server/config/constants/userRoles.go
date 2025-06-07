package constants

// UserRoles defines the roles available in the system.
type UserRoles string

const (
	// UserRoleAdmin represents an admin user.
	UserRoleAdmin UserRoles = "admin"
	// UserRoleUser represents a regular user.
	UserRoleUser UserRoles = "user"
	// UserRoleGuest represents a guest user.
	UserRoleGuest UserRoles = "guest"
	// UserRoleSuperAdmin represents a super admin user.
	UserRoleSuperAdmin UserRoles = "superadmin"
	// UserRoleModerator represents a moderator user.
	UserRoleModerator UserRoles = "moderator"
	// UserRoleSupport represents a support user.
	UserRoleSupport UserRoles = "support"
)

// UserRolesList is a list of all user roles.
var UserRolesList = []UserRoles{
	UserRoleAdmin,
	UserRoleUser,
	UserRoleGuest,
	UserRoleSuperAdmin,
	UserRoleModerator,
	UserRoleSupport,
}

// IsValid checks if the given role is a valid user role.
func IsValid(role UserRoles) bool {
	for _, r := range UserRolesList {
		if r == role {
			return true
		}
	}
	return false
}
