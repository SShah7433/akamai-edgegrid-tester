package main

import (
	"fmt"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"net/http"
)


func createRequest(host string, path string, client_token string, client_secret string, access_token string) {
	// format configuration for edgegrid
	config := edgegrid.Config{
		Host:         host,
		ClientToken:  client_token,
		ClientSecret: client_secret,
		AccessToken:  access_token,
		MaxBody:      1024,
		Debug:        true, // enable logging of auth flow to console
	}

	// create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s%s", config.Host, path), nil)
	edgegrid.AddRequestHeader(config, req)
}

func main() {

	// setup variables for inputs from user
	var host string
	var client_token string
	var client_secret string
	var access_token string
	var path string

	// receive inputs from user
	fmt.Print("Enter akamai host (xxxxxx.luna.akamaiapis.net): ")
	fmt.Scanln(&host)

	fmt.Print("Enter path with query string params (/path?param1=value1&param2=value2): ")
	fmt.Scanln(&path)

	fmt.Print("Enter client token (xxxx-xxxxxxxxxxx-xxxxxxxxxxx): ")
	fmt.Scanln(&client_token)

	fmt.Print("Enter client secret (xxxxxxxxxxxxxxxxxxxxxxxxxxxxx): ")
	fmt.Scanln(&client_secret)

	fmt.Print("Enter access token (xxxx-xxxxxxxxxxx-xxxxxxxxxxx): ")
	fmt.Scanln(&access_token)

	createRequest(host, path, client_token, client_secret, access_token)
}
