package postgres

import (
	"log"

	"github.com/captainmog/mogbot/internal/mogbot"
	"github.com/jmoiron/sqlx"
)

type RoleService struct {
	*sqlx.DB
}

func (s *RoleService) AddRole(r mogbot.Role) error {
	if _, err := s.Exec(
		"insert into guild_role (id, sticky_role, auto_role, mod_role) values ($1, $2, $3, $4)",
		r.ID,
		r.Sticky,
		r.Auto,
		r.Mod,
	); err != nil {
		log.Printf("Error inserting role '%s': %s", r.ID, err)
		return err
	}
	return nil
}

//func GetRole(roleID string) (*Role, error) {
//
//}
//
//func GetAllRoles() ([]*Role, error) {}
//
//func UpdateRole(nr *Role) error {}
//
//func RemoveRole(roleID string) error {}
