package mogbot

type Guild struct {
	ID             string `gorm:"primaryKey"`
	GreetMsg       string
	LeaveMsg       string
	MsgXP          int
	LevelMsg       string
	RankMsg        string
	ModLogID       string
	ServerLogID    string
	MemberLogID    string
	VoiceLogID     string
	MessageLogID   string
	JoinLeaveLogID string
	ModMailLogID   string
	Prefixes       []string
}

type GuildService interface {
	AddGuild(*Guild) error
	GetGuild(string) (*Guild, error)
	GetAllGuilds() ([]Guild, error)
	UpdateGuild(string, map[string]interface{}) error
	RemoveGuild(string) error
}

func (g *Guild) FillStruct(f map[string]interface{}) error {
	for k, v := range f {
		err := SetField(g, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
