package domain

type Collection struct {
	id    string `json:"id"`
	name  string `json:"name"`
	user  User   `json:"user"`
	games []Game `json:"games"`
}
