package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type httpbinResponse struct {
	Origin string `json:"origin"`
}

func main() {
	// 1. Fetch current IP from httpbin.org
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	hResp := httpbinResponse{}
	err = json.Unmarshal(body, &hResp)
	if err != nil {
		log.Fatal(err)
	}

	currentIp := hResp.Origin
	fmt.Println(currentIp)
}