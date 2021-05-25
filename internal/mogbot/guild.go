package mogbot

type Guild struct {
	ID             string
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
	ModMailCatID   string
	Members        []Member
	Roles          []Role
}

type GuildService interface {
	AddGuild(*Guild) error
	GetGuild(string) (*Guild, error)
	GetAllGuilds() ([]*Guild, error)
	UpdateGuild(*Guild) error
	RemoveGuild(string) error
}
