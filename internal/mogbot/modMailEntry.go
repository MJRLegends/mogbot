package mogbot

type ModMailEntry struct {
	ID        int64  `db:"id"`
	GuildID   string `db:"guild_id"`
	ChannelID string `db:"channel_id"`
	UserID    string `db:"user_id"`
}

type ModMailEntryService interface {
	AddModMailEntry(ModMailEntry) (int64, error)
	GetModMailEntry(int64) (*ModMailEntry, error)
	GetGuildModMailEntries(string) ([]ModMailEntry, error)
	GetAllModMailEntries() ([]ModMailEntry, error)
	UpdateModMailEntry(ModMailEntry) error
	RemoveModMailEntry(int64) error
}
