package mogbot

type Guild struct {
	ID             string
	GreetMsg       string
	LeaveMsg       string
	MsgXP          int
	LevelMsg       string
	RankMsg        string
	//ModLogID       string
	//ServerLogID    string
	//MemberLogID    string
	//VoiceLogID     string
	//MessageLogID   string
	//JoinLeaveLogID string
	//ModMailLogID   string
	Prefixes       []string
}

type GuildService interface {
	AddGuild(*Guild) error
	GetGuild(string) (*Guild, error)
	GetAllGuilds() ([]Guild, error)
	UpdateGuild(string, map[string]interface{}) (*Guild, error)
	RemoveGuild(string) error
}
