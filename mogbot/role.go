package mogbot

type Role struct {
	ID      string `gorm:"primaryKey"`
	GuildID string `gorm:"primaryKey"`
	Sticky  bool
	Auto    bool
	Mod     bool
}

type RoleService interface {
	AddRole(Role) error
	GetRole(string) (*Role, error)
	GetAllRoles() ([]*Role, error)
	UpdateRole(*Role) error
	RemoveRole(string) error
}
