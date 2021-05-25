package mogbot

type Member struct {
	UserID  string `gorm:"primaryKey"`
	GuildID string `gorm:"primaryKey"`
	Roles   []Role
}

type MemberService interface {
	AddMember(m *Member) error
	GetMember(uid, gid string) (*Member, error)
	GetGuildMembers(gid string) ([]*Member, error)
	UpdateMember(nm *Member) error
	RemoveMember(uid, gid string) error
}
