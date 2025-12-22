package game

type TrackGamesRequest struct {
	Game string `json:"game"`
}

type UntrackGamesRequest struct {
	Game string `json:"game"`
}

type Game struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Cover            Cover   `json:"cover"`
	AggregatedRating float64 `json:"aggregated_rating"`
}

type Cover struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}
