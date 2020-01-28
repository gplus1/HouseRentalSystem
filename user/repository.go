package Tourist

import "github.com/gplus1/HouseRentalSystem/entity"

// UserRepository specifies application Tourist related database operations
type UserRepository interface {
	Users() ([]entity.Tourist, []error)
	Tourist(id uint) (*entity.Tourist, []error)
	UserByEmail(email string) (*entity.Tourist, []error)
	UpdateUser(Tourist *entity.Tourist) (*entity.Tourist, []error)
	DeleteUser(id uint) (*entity.Tourist, []error)
	StoreUser(Tourist *entity.Tourist) (*entity.Tourist, []error)
	PhoneExists(phone string) bool
	EmailExists(email string) bool
	UserRoles(*entity.Tourist) ([]entity.Role, []error)
}

// RoleRepository speifies application Tourist role related database operations
type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// SessionRepository specifies logged in Tourist session related database operations
type SessionRepository interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}
