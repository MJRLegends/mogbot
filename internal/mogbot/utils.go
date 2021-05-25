package mogbot

const (
	DiscordEpoch   = 1420070400000
)

func FindRole(roleID string, roles []*Role) (*Role, bool) {
	for _, r := range roles {
		if roleID == r.ID {
			return r, true
		}
	}
	return nil, false
}
