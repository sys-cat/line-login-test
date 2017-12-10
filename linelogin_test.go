package linelogin

import (
	"reflect"
	"testing"
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
