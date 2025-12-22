package game

type TrackGamesRequest struct {
	Game string `json:"game"`
}

type UntrackGamesRequest struct {
	Game string `json:"game"`
}

type Game struct {
	Id               int     `json:"id" bson:"_id"`
	Name             string  `json:"name" bson:"name"`
	Cover            Cover   `json:"cover" bson:"cover"`
	AggregatedRating float64 `json:"aggregated_rating" bson:"aggregated_rating"`
}

type Cover struct {
	Id  int    `json:"id" bson:"_id"`
	Url string `json:"url" bson:"url"`
}
