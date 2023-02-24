// Package setup provides a simple way to setup a new connection to the kong admin api
package setup

import (
	"fmt"
	"github.com/captainsuchard/kong-admin-api-wrapper/vars"
	"github.com/kong/go-kong/kong"
	"strings"
)

// NewKongClient sets up a new connection to the kong admin api and stores the connection in vars.KongClient
func NewKongClient(host, port string, useTLS bool) error {
	if strings.TrimSpace(host) == "" {
		return fmt.Errorf("host cannot be empty")
	}
	if strings.TrimSpace(port) == "" {
		return fmt.Errorf("port cannot be empty")
	}
	var kongAdminURL string
	if useTLS {
		kongAdminURL = "https://" + host + ":" + port
	} else {
		kongAdminURL = "http://" + host + ":" + port
	}
	var err error
	vars.KongClient, err = kong.NewClient(&kongAdminURL, nil)
	if err != nil {
		return fmt.Errorf("error creating kong client: %v", err)
	}
	return nil
}

// SetKongClient allows the user to supply their own kong client
func SetKongClient(client *kong.Client) error {
	if client == nil {
		return fmt.Errorf("client cannot be nil")
	}
	vars.KongClient = client
	return nil
}
