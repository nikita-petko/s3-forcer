package cache

type cachedData struct {
	Positions map[string]int `json:"positions"`
	Channels  []string       `json:"channels"`
}
