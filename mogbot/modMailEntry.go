package mogbot

type ModMailEntry struct {
	ID        int64  `gorm:"primaryKey"`
	GuildID   string `gorm:"foreignKey"`
	ChannelID string
	UserID    string `gorm:"foreignKey"`
}

type ModMailEntryService interface {
	AddModMailEntry(ModMailEntry) (int64, error)
	GetModMailEntry(int64) (*ModMailEntry, error)
	GetGuildModMailEntries(string) ([]ModMailEntry, error)
	GetAllModMailEntries() ([]ModMailEntry, error)
	UpdateModMailEntry(ModMailEntry) error
	RemoveModMailEntry(int64) error
}
