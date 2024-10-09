package ddns

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type DnsRecord struct {
	CurrentIPS CurIp
	Domain     string
	Name       string
	CNAME      string
	Ttl        int
	A          bool
	AAAA       bool
}

type CurIp struct {
	ipv4 string
	ipv6 string
}

type DdnsSettings struct {
	client            *http.Client
	CurrentIPS        CurIp
	Service           string
	AuthKey           string
	AutoCnameDefault  string
	Record            []DnsRecord
	RecordCount       int8
	AutoCnameCreation bool
}

func getBodyOfThis(client *http.Client, url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Handle error
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		// Handle error
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// Handle error
		return "", err
	}
	return string(body), nil
}

func getIP(client *http.Client, ipv4 bool, ipv6 bool) (string, string) {
	var ipv4_address string
	var ipv6_address string
	if ipv4 {
		ipv4_response, err := getBodyOfThis(client, "https://api.ipify.org")
		if err != nil {
			// Handle error
		}
		ipv4_address = ipv4_response
	}
	if ipv6 {
		ipv6_response, err := getBodyOfThis(client, "https://api64.ipify.org")
		if err != nil {
			// Handle error
		}
		ipv6_address = ipv6_response
	}
	return ipv4_address, ipv6_address
}

func new(httpClient *http.Client) DdnsSettings {
	// Want to not keep needing to recreate http clients
	// so will add pointer to struct
	// This way, its easier to set up if I want to record
	// cookies or not later
	this := DdnsSettings{
		client: httpClient,
	}
	return this
}

func (ds *DdnsSettings) Ddns() error {
	ipv4, ipv6 := getIP(ds.client, true, true)
	if ipv4 == ds.CurrentIPS.ipv4 && ipv6 == ds.CurrentIPS.ipv6 {
		return nil
	}

	// TODO: Create a good logic that if the ipadress has not changed, dont try to update

	ds.GandiUpdateAll()

	// CloudflareUpdate()
	return nil
}

func Test() {
	// Not sure
	qjar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: qjar,
	}

	// TODO: Create a good logic that if the ipadress has not changed, dont try to update

	ds := new(client)
	ds.GandiUpdateAll()

	// CloudflareUpdate()
}
