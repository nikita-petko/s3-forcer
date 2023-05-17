package cache

type cachedData struct {
	Positions map[string]uint64 `json:"positions"`
	Channels  []string          `json:"channels"`
}
