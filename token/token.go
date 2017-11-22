package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var LINE_TOKEN_URL = "https://api.line.me/oauth2/v2.1/token"
var CONTENT_TYPE = "application/x-www-form-urlencoded"

type (
	Request struct {
		GrantType    string
		Code         string
		RedirectURL  string
		ClientID     string
		ClientSecret string
	}

	Response struct {
		AccessToken  string
		ExpiresIn    int64
		IDToken      string
		RefreshToken string
		Scope        string
		TokenType    string
	}
)

func New() Request {
	return Request{}
}

func (req *Request) Parameters(code string, url string, channel_id string, channel_secret string) error {
	req.GrantType = "authorization_code"
	req.Code = code
	req.RedirectURL = url
	req.ClientID = channel_id
	req.ClientSecret = channel_secret
	return nil
}

func (req *Request) BuildParams() url.Values {
	value := url.Values{}
	value.Set("grant_type", req.GrantType)
	value.Add("code", req.Code)
	value.Add("redirect_url", req.RedirectURL)
	value.Add("client_id", req.ClientID)
	value.Add("client_secret", req.ClientSecret)
	return value
}

func GetToken(req Request) (res Response, err error) {
	newreq, err := http.NewRequest(
		"POST",
		LINE_TOKEN_URL,
		strings.NewReader(req.BuildParams().Encode()),
	)
	if err != nil {
		return res, err
	}

	newreq.Header.Set("Content-Type", CONTENT_TYPE)

	client := &http.Client{}
	resp, err := client.Do(newreq)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return res, errors.New(fmt.Sprintf("response is invalid, status code is %d", resp.StatusCode))
	}
	body_byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(body_byte, &res)
	if err != nil {
		return Response{}, err
	}
	return res, nil
}
