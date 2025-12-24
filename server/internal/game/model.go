package game

type TrackGamesRequest struct {
	Game string `json:"game"`
}

type UntrackGamesRequest struct {
	Game string `json:"game"`
}
