package mogbot

type Member struct {
	UserID  string `gorm:"primaryKey"`
	GuildID string `gorm:"primaryKey"`
	Roles   []Role
}

type MemberService interface {
	AddMember(*Member) error
	GetMember(string, string) (*Member, error)
	GetGuildMembers(string) ([]*Member, error)
	UpdateMember(string, string, map[string]interface{}) error
	RemoveMember(string, string) error
}

func (m *Member) FillStruct(f map[string]interface{}) error {
	for k, v := range f {
		err := SetField(m, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
