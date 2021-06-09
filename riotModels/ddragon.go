package riotModels

type Ddragon struct {
	Type    string `json:"type"`
	Format  string `json:"format"`
	Version string `json:"version"`
	Data    map[string]struct {
		Version string `json:"version"`
		ID      string `json:"id"`
		Key     int    `json:"key,string"`
		Name    string `json:"name"`
		Title   string `json:"title"`
		Blurb   string `json:"blurb"`
		Info    struct {
			Attack     int `json:"attack"`
			Defense    int `json:"defense"`
			Magic      int `json:"magic"`
			Difficulty int `json:"difficulty"`
		} `json:"info"`
		Image struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
		Tags    []string `json:"tags"`
		Partype string   `json:"partype"`
	} `json:"data"`
}
