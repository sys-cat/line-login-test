package login

import (
	"encoding/base64"
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

func (param Params) Parameters(channel_id string, channel_secret string, redirect string) error {
	param.ResponseType = "code"
	param.ClientID = channel_id
	param.RedirectURL = base64.StdEncoding.EncodeToString([]byte(redirect))
	param.State = time.Now().Unix()
	param.Scope = "Profile"
	return nil
}

func (param Params) OutputURL() string {
	value := url.Values{}
	value.Set("response_type", param.ResponseType)
	value.Add("client_id", param.ClientID)
	value.Add("redirect_url", param.RedirectURL)
	value.Add("state", param.State)
	value.Add("scope", param.Scope)
	values := value.Encode()
	return fmt.Sprintf("%s?%s", LINE_LOGIN_URL, values)
}
