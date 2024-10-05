package ddns

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// for more info check https://developers.cloudflare.com/api/operations/dns-records-for-a-zone-list-dns-records
func CloudflareUpdate() {

	url := "https://api.cloudflare.com/client/v4/zones/zone_id/dns_records/dns_record_id"

	payload := strings.NewReader("{\n  \"comment\": \"Domain verification record\",\n  \"name\": \"example.com\",\n  \"proxied\": true,\n  \"settings\": {},\n  \"tags\": [],\n  \"ttl\": 3600,\n  \"content\": \"198.51.100.4\",\n  \"type\": \"A\"\n}")

	req, _ := http.NewRequest("PATCH", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Email", "")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func CloudflareList() {

	url := "https://api.cloudflare.com/client/v4/zones/zone_id/dns_records"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Email", "")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle error
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

/*
{
  "errors": [],
  "messages": [],
  "success": true,
  "result_info": {
    "count": 1,
    "page": 1,
    "per_page": 20,
    "total_count": 2000
  },
  "result": [
    {
      "comment": "Domain verification record",
      "name": "example.com",
      "proxied": true,
      "settings": {},
      "tags": [],
      "ttl": 3600,
      "content": "198.51.100.4",
      "type": "A",
      "comment_modified_on": "2024-01-01T05:20:00.12345Z",
      "created_on": "2014-01-01T05:20:00.12345Z",
      "id": "023e105f4ecef8ad9ca31a8372d0c353",
      "meta": {},
      "modified_on": "2014-01-01T05:20:00.12345Z",
      "proxiable": true,
      "tags_modified_on": "2025-01-01T05:20:00.12345Z"
    }
  ]
}
*/
