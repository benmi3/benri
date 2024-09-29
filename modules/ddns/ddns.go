package ddns

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type DnsRecord struct {
	Name string
	A    bool
	AAAA bool
}

type DdnsSettings struct {
	Service string
	AuthKey string
	Record  []DnsRecord
}
type IpMemory struct {
	ipv4 string
	ipv6 string
}

var currentIPs IpMemory

func getBodyOfThis(client *http.Client, url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle error
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		//Handle error
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		//Handle error
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
			//Handle error
		}
		ipv4_address = ipv4_response
	}
	if ipv6 {
		ipv6_response, err := getBodyOfThis(client, "https://api64.ipify.org")
		if err != nil {
			//Handle error
		}
		ipv6_address = ipv6_response
	}
	return ipv4_address, ipv6_address
}

func Ddns() error {
	// Not sure
	qjar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: qjar,
	}

	ipv4, ipv6 := getIP(client, true, true)
	if ipv4 == currentIPs.ipv4 && ipv6 == currentIPs.ipv6 {
		return nil
	}

	// TODO: Create a good logic that if the ipadress has not changed, dont try to update

	GandiUpdate(client)

	//CloudflareUpdate()
	return nil
}

func main() {
	// Not sure
	qjar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: qjar,
	}

	// TODO: Create a good logic that if the ipadress has not changed, dont try to update

	//req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	//resp, err := client.Do(req)
	// ...

	GandiUpdate(client)

	//CloudflareUpdate()

}
