package mogbot

type ReactionRole struct {
	ID        int64  `db:"id"`
	RoleID    string `db:"role_id"`
	MessageID string `db:"message_id"`
	Emoji     string `db:"emoji"`
}

type ReactionRoleService interface {
	AddReactionRole(ReactionRole) error
	GetReactionRole(string, string) (*ReactionRole, error)
	GetReactionRoles() ([]ReactionRole, error)
	UpdateReactionRole(int) (int, error)
	RemoveReactionRole(int) (int, error)
}
