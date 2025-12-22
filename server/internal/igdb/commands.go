package igdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

type Games []Game

var client *http.Client = &http.Client{}

func GetGame(id int) Game {
	url := "https://api.igdb.com/v4/games"

	body := bytes.NewReader([]byte(fmt.Sprintf("fields id,name,aggregated_rating,cover.url; where id=%d;", id)))

	resp, err := newRequest(url, body)
	if err != nil {
		log.Fatalf("Something went wrong with the request: %s", err.Error())
	}

	defer resp.Body.Close()

	result := Games{}

	byte_resp, _ := io.ReadAll(resp.Body)

	json.Unmarshal(byte_resp, &result)

	return result[0]
}

func GetGamesByIds(ids []int) Games {
	url := "https://api.igdb.com/v4/games"

	idsString := ""
	for i, id := range ids {
		if i != 0 {
			idsString += ","
		}
		idsString += fmt.Sprintf("%d", id)
	}

	bodyString := fmt.Sprintf(`
		fields id,name,aggregated_rating,cover.url;
		where id = (%s);
		`,
		idsString,
	)

	body := bytes.NewReader([]byte(bodyString))
	resp, err := newRequest(url, body)
	if err != nil {
		log.Fatalf("Something went wrong with the request: %s", err.Error())
	}

	defer resp.Body.Close()

	result := Games{}

	byte_resp, _ := io.ReadAll(resp.Body)
	json.Unmarshal(byte_resp, &result)

	return result
}

func GetGamesByName(name string) Games {
	url := "https://api.igdb.com/v4/games"

	// game_type = 0 makes sure that only games are shown (0 is the reference ID for "main_games")
	bodyString := fmt.Sprintf(`
        fields id,name,aggregated_rating,cover.url; 
        where game_type = 0;
        search "%s";
        `,
		name,
	)

	body := bytes.NewReader([]byte(bodyString))

	resp, err := newRequest(url, body)
	if err != nil {
		log.Fatalf("Something went wrong with the request: %s", err.Error())
	}

	defer resp.Body.Close()

	result := Games{}

	byte_resp, _ := io.ReadAll(resp.Body)

	json.Unmarshal(byte_resp, &result)

	return result
}

func newRequest(url string, body *bytes.Reader) (*http.Response, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalf("Getting game did not work, %s", err.Error())
	}
	request.Header.Add("Client-ID", ApiKeys.Client_id)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", AccessToken.Access_token))

	resp, err := client.Do(request)
	return resp, err
}
