package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	openapi "github.com/delphix/dct-sdk-go"
)

func main() {

	var KEY string
	var HOST string

	fmt.Println("Enter Your KEY (excluding 'apk'): ")
	fmt.Scanln(&KEY)

	fmt.Println("Enter Your HOST ('hostname.domain'): ")
	fmt.Scanln(&HOST)

	apiKeyMap := make(map[string]openapi.APIKey)
	apiKeyMap["ApiKeyAuth"] = openapi.APIKey{
		Key:    KEY,
		Prefix: "apk",
	}

	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, apiKeyMap)

	cfg := openapi.NewConfiguration()
	cfg.Host = HOST
	cfg.Scheme = "https"
	cfg.HTTPClient = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	client := openapi.NewAPIClient(cfg)

	req := client.EnginesApi.GetEngines(ctx)

	res, _, err := client.EnginesApi.GetEnginesExecute(req)

	if err != nil {
		fmt.Print("\n__________ERR__________\n")
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("List of engines are: \n")

	for _, j := range res.Items {
		fmt.Printf("%s \n", *j.Hostname)
	}
}
