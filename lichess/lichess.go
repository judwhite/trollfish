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
	White         int    `json:"white"`
	Draws         int    `json:"draws"`
	Black         int    `json:"black"`
}

type PositionResults struct {
	White int    `json:"white"`
	Draws int    `json:"draws"`
	Black int    `json:"black"`
	Moves []Move `json:"moves"`
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

	return result, nil
}
