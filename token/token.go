package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
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
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		IDToken      string `json:"id_token"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
	}

	Token struct {
		Iss     string
		Sub     string
		Aud     string
		Exp     int64  // token enable limit unix time
		Iat     int64  // generate id_token unix time
		Nonece  string // require set "nonece" value
		Name    string // require add "profile" to scope
		Picture string // require add "profile" to scope
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
	value.Add("redirect_uri", req.RedirectURL)
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

func ParseIDToken(res Response, secret string) (token Token, err error) {
	// Parse jwt use secret
	t, err := jwt.Parse(res.IDToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return token, fmt.Errorf("unexpexted signing method: %v", t.Header["token"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return token, err
	}
	claim, ok := t.Claims.(jwt.MapClaims)
	if ok && t.Valid {
		if !res.Nonce {
			return token, errors.New("unset nonece data")
		}
		if res.Nonce != claim["nonce"] {
			return token, errors.New("invalid nonce from id_token")
		}
		token.Iss = claim["iss"]
		token.Sub = claim["iss"]
		token.Aud = claim["aud"]
		token.Exp = claim["exp"]
		token.Iat = claim["iat"]
		token.Nonece = claim["nonce"]
		token.Name = claim["name"]
		token.Picture = claim["picture"]
	}
	return token, errors.New("parse missing.")
}
