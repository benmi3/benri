package ddns

import (
	"io"
	"net/http"
	"strings"
)

// For more info check https://api.gandi.net/docs/livedns/
func GandiUpdate(client *http.Client) string {

	url := "https://api.gandi.net/v5/livedns/domains/example.com/records/www/CNAME"
	// There should be no way that the request should error
	payload := strings.NewReader("{\"rrset_values\":[\"www.example.org\"],\"rrset_ttl\":320}")

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("authorization", "Bearer pat_abc-123")
	req.Header.Add("content-type", "application/json")

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	return string(body)

}

//	func example() {
//		resp, err := client.Get("http://example.co")
//		// ...
//
//		req, err := http.NewRequest("GET", "http://example.com", nil)
//		// ...
//		req.Header.Add("If-None-Match", `W/"wyzzy"`)
//		resp, err := client.Do(req)
//		// ...
//	}
func GandiList(client *http.Client) string {

	url := "https://api.gandi.net/v5/livedns/domains/example.com/records/www"
	// There should be no way that the request should error
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer pat_abc-123")

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	return string(body)

}

/*
[
  {
    "rrset_name": "www",
    "rrset_type": "A",
    "rrset_values": [
      "192.0.2.1"
    ],
    "rrset_ttl": 320,
    "rrset_href": "https://api.test/v5/livedns/domains/example.com/records/www/A"
  },
  {
    "rrset_name": "www",
    "rrset_type": "TXT",
    "rrset_values": [
      "some-text-value"
    ],
    "rrset_ttl": 320,
    "rrset_href": "https://api.test/v5/livedns/domains/example.com/records/www/TXT"
  }
]
*/
