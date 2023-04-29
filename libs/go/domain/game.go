package domain

type Game struct {
	id string `json:"id"`
	name string `json:"name"`
	minPlayers int `json:"min_players"`
	maxPlayers int `json:"max_players"`
	avgPlayTime int `json:"avg_play_time"`
}