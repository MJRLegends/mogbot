package mogbot

const (
	LogTypeJoinLeave = iota
	LogTypeMemberUpdate
	LogTypeServerUpdate
	LogTypeMessageUpdate
	LogTypeChannelUpdate
	LogTypeModmail
)

type Log struct {
	ID    string
	Types []int
}

type LogService interface {
	AddLog(*Log) error
	GetLogByID(string) (*Log, error)
	GetLogsByType(int) ([]*Log, error)
	UpdateLog(string, map[string]interface{}) (*Log, error)
	RemoveLog(string) error
}
