package linelogin

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	param := New()
	if reflect.TypeOf(param.ResponseType).Name() == "string" {
		t.Log("Params.ResponseType equal string")
	} else {
		t.Fatal("Params.ResponseType is not string")
	}
	if reflect.TypeOf(param.ClientID).Name() == "string" {
		t.Log("Params.ClientID equal string")
	} else {
		t.Fatal("Params.ClientID is not string")
	}
	if reflect.TypeOf(param.RedirectURL).Name() == "string" {
		t.Log("Params.RedirectURL equal string")
	} else {
		t.Fatal("Params.RedirectURL is not string")
	}
	if reflect.TypeOf(param.State).Name() == "string" {
		t.Log("Params.State equal string")
	} else {
		t.Fatal("Params.State is not string")
	}
	if reflect.TypeOf(param.Scope).Name() == "string" {
		t.Log("Params.Scope equal string")
	} else {
		t.Fatal("Params.Scope is not string")
	}
}

func TestParameters(t *testing.T) {
	param := New()
	channel_id := "test channel id"
	channel_secret := "test channel secret"
	redirect := "test redirect url"
	param.Parameters(channel_id, channel_secret, redirect)
	if param.ClientID == channel_id {
		t.Log("channel_id is clean")
	} else {
		t.Fatal("channel_id is invalid")
	}
	if param.RedirectURL == redirect {
		t.Log("redirect url is clean")
	} else {
		t.Fatal("redirect url is invalid")
	}
	if param.State == fmt.Sprint(time.Now().Unix()) {
		t.Log("state is clean")
	} else {
		t.Fatal("state is invalid")
	}
	if param.Scope == "profile openid" {
		t.Log("scope is clean")
	} else {
		t.Fatal("scope is invalid")
	}
}

func TestOutputURL(t *testing.T) {
	t.Log("nothing test")
}
