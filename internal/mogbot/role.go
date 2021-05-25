package mogbot

import "github.com/bwmarrin/discordgo"

type Role struct {
	ID     string
	Sticky bool
	Auto   bool
	Mod    bool
}

func NewRole(r *discordgo.Role) *Role {
	return &Role{r.ID, true, false, false}
}

type RoleService interface {
	AddRole(Role) error
	GetRole(string) (*Role, error)
	GetAllRoles() ([]*Role, error)
	UpdateRole(*Role) error
	RemoveRole(string) error
}
