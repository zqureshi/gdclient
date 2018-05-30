package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/miekg/dns"
)

type httpbinResponse struct {
	Origin string `json:"origin"`
}

func main() {
	hostIP, err := getHostIP()
	if err != nil {
		log.Fatal(err)
	}

	domainIP, err := getDomainIP("home.zeeshanq.com.")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IP of current host is " + hostIP)
	fmt.Println("IP of domain is " + domainIP)

	if domainIP != hostIP {
		fmt.Println("Updating domain IP to " + hostIP)
		// TODO implement and finish off
	} else {
		fmt.Println("IPs match, nothing to update")
	}
}

func getHostIP() (string, error) {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	hResp := httpbinResponse{}
	err = json.Unmarshal(body, &hResp)
	if err != nil {
		return "", err
	}

	return hResp.Origin, nil
}

func getDomainIP(domain string) (string, error) {
	m := new(dns.Msg)
	m.SetQuestion(domain, dns.TypeA)

	in, err := dns.Exchange(m, "1.1.1.1:53")
	if err != nil {
		return "", err
	}

	a, ok := in.Answer[0].(*dns.A);
	if !ok {
		return "", fmt.Errorf("could not read DNS answer: %s", in)
	}

	return a.A.String(), nil
}