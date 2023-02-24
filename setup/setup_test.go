package setup

import (
	"github.com/captainsuchard/kong-admin-api-wrapper/vars"
	"testing"
)

func TestNewKongClient(t *testing.T) {
	kongHost := "localhost"
	kongPort := "8001"
	useTLS := true
	err := NewKongClient(kongHost, kongPort, useTLS)
	if err != nil {
		t.Errorf("error setting up kong client: %v", err)
	}
	if vars.KongClient.BaseRootURL() != "https://localhost:8001" {
		t.Errorf("expected kong client base root url to be 'https://localhost:8001', got '%v'", vars.KongClient.BaseRootURL())
	}
}

func BenchmarkNewKongClient(b *testing.B) {
	kongHost := "localhost"
	kongPort := "8001"
	useTLS := true
	for i := 0; i < b.N; i++ {
		err := NewKongClient(kongHost, kongPort, useTLS)
		if err != nil {
			b.Errorf("error setting up kong client: %v", err)
		}
	}
}

func TestNewKongClientNoTLS(t *testing.T) {
	kongHost := "localhost"
	kongPort := "8001"
	useTLS := false
	err := NewKongClient(kongHost, kongPort, useTLS)
	if err != nil {
		t.Errorf("error setting up kong client: %v", err)
	}
	if vars.KongClient.BaseRootURL() != "http://localhost:8001" {
		t.Errorf("expected kong client base root url to be 'http://localhost:8001', got '%v'", vars.KongClient.BaseRootURL())
	}
}

func TestNewKongClientEmptyHost(t *testing.T) {
	kongHost := ""
	kongPort := "8001"
	useTLS := true
	err := NewKongClient(kongHost, kongPort, useTLS)
	if err == nil {
		t.Error("expected error setting up kong client, got nil")
	}
}

func TestNewKongClientEmptyPort(t *testing.T) {
	kongHost := "localhost"
	kongPort := ""
	useTLS := true
	err := NewKongClient(kongHost, kongPort, useTLS)
	if err == nil {
		t.Error("expected error setting up kong client, got nil")
	}
}
