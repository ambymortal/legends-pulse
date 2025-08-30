package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Player struct {
	Guild  string
	Name   string
	Level  int
	Exp    string
	Gender string
	Job    string
	Quests int
	Cards  int
	Donor  bool
	Fame   int
}

// interactions with the legends api

// add a timeout to handle connections that are unresponsive
var client = &http.Client{Timeout: 10 * time.Second}

func ParseCharacterJSON(username string) (Player, error) {
	var data Player
	jsonUrl := fmt.Sprintf("https://maplelegends.com/api/character?name=%s", username)
	charInfo, err := client.Get(jsonUrl)
	if err != nil {
		return data, fmt.Errorf("this character does not exist")
	}

	body, _ := io.ReadAll(charInfo.Body)
	json.Unmarshal(body, &data)

	return data, nil
}

// return bytebuf of character based on url
func ParseCharacterImage(url string) (*bytes.Buffer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	imgBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image data: %w", err)
	}

	imgBuf := bytes.NewBuffer(imgBytes)

	return imgBuf, nil
}
