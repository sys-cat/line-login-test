package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var LINE_PROFILE_URL = "https://api.line.me/v2/profile"

type (
	Response struct {
		DisplayName   string `json:"displayName"`
		UserID        string `json:"userId"`
		PictureURL    string `json:"pictureUrl"`
		StatusMessage string `json:"statusMessage"`
	}
)

func CreateHeaderParam(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

func GetProfileData(token string) (Response, error) {
	req, err := http.NewRequest(
		"GET",
		LINE_PROFILE_URL,
		nil,
	)
	if err != nil {
		return Response{}, err
	}

	if token == "" {
		return Response{}, errors.New("nil token")
	}
	req.Header.Set("Authorization", CreateHeaderParam(token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Response{}, errors.New(fmt.Sprintf("response is invalid, status code is %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}
