package linelogin

import (
	"fmt"
	"net/url"
	"time"
)

type Params struct {
	ResponseType string
	ClientID     string
	RedirectURL  string
	State        string
	Scope        string
}

var LINE_LOGIN_URL = "https://access.line.me/oauth2/v2.1/authorize"

func New() Params {
	return Params{}
}

func (param *Params) Parameters(channel_id string, channel_secret string, redirect string) error {
	param.ResponseType = "code"
	param.ClientID = channel_id
	param.RedirectURL = redirect
	param.State = fmt.Sprint(time.Now().Unix())
	param.Scope = "profile"
	return nil
}

func (param *Params) OutputURL() string {
	value := url.Values{}
	value.Set("response_type", param.ResponseType)
	value.Add("client_id", param.ClientID)
	value.Add("redirect_uri", param.RedirectURL)
	value.Add("state", param.State)
	value.Add("scope", param.Scope)
	value.Add("nonce", fmt.Sprint(time.Now().Unix()))
	values := value.Encode()
	return fmt.Sprintf("%s?%s", LINE_LOGIN_URL, values)
}
