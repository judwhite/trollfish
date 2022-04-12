package lichess

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const allRatings = "1600,1800,2000,2200,2500"
const allSpeeds = "bullet,blitz,rapid,classical"

type Move struct {
	UCI           string `json:"uci"`
	SAN           string `json:"san"`
	AverageRating int    `json:"averageRating"`

	White      int `json:"white"`
	Black      int `json:"black"`
	Draws      int `json:"draws"`
	TotalGames int `json:"total_games"`

	WhitePercent      float64 `json:"white_pct"`
	BlackPercent      float64 `json:"black_pct"`
	DrawsPercent      float64 `json:"draws_pct"`
	PopularityPercent float64 `json:"popularity_pct"`
}

type PositionResults struct {
	White      int    `json:"white"`
	Draws      int    `json:"draws"`
	Black      int    `json:"black"`
	Moves      []Move `json:"moves"`
	TotalGames int    `json:"total_games"`
}

func Lookup(fen, play string) (PositionResults, error) {
	var result PositionResults

	u, err := url.Parse("https://explorer.lichess.ovh/lichess")
	if err != nil {
		return result, err
	}
	q := u.Query()
	q.Add("fen", fen)
	if play != "" {
		q.Add("play", play)
	}
	q.Add("recentGames", "0")
	q.Add("topGames", "0")
	q.Add("speeds", allSpeeds)
	q.Add("ratings", allRatings)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("http status code %d", resp.StatusCode)
	}

	total := result.White + result.Black + result.Draws
	result.TotalGames = total

	for i := 0; i < len(result.Moves); i++ {
		move := result.Moves[i]

		moveTotal := move.White + move.Black + move.Draws

		popularity := float64(moveTotal) / float64(total) * 100
		white := float64(move.White) / float64(moveTotal) * 100
		black := float64(move.Black) / float64(moveTotal) * 100
		draw := float64(move.Draws) / float64(moveTotal) * 100

		result.Moves[i].WhitePercent = white
		result.Moves[i].BlackPercent = black
		result.Moves[i].DrawsPercent = draw
		result.Moves[i].PopularityPercent = popularity
		result.Moves[i].TotalGames = moveTotal
	}

	return result, nil
}
