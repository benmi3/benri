package ddns

import (
	"io"
	"net/http"
	"strings"
)

// For more info check https://api.gandi.net/docs/livedns/
func (ds *DdnsSetting)GandiUpdateSingelByType(itemNum int, rrset_type string) string {
	// This function will only be used if the user allows only ipv4 or ipv6
	// Other time updating by domain will lead to less api calls,
	// So this function will not be used then.

	// TODO: item size and itemNum check

	//url := "https://api.gandi.net/v5/livedns/domains/example.com/records/www/CNAME"
	url := fmt.Sprintf("https://api.gandi.net/v5/livedns/domains/%s/records/%s/%s", ds.DnsRecord[ItemNum].Domain,ds.DnsRecord[ItemNum].Name, rrset_type)
	
	var rrset_values string
	
	if ds.DnsRecord[ItemNum].AAAA {
		rrset_values := ds.CurrentIPS.ipv6
	} else if ds.DnsRecord[ItemNum].A {
		rrset_values := ds.CurrentIPS.ipv4
	} else {
		// Config is not correct
		// This should not happen here
		return "500 Internal ERROR"
	}

	authToken := fmt.Sprintf("Bearer %s",ds.AuthKey)

	payloadString := fmt.Sprintf("{\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}",rrset_values,ds.DnsRecord[ItemNum].Ttl)

	payload := strings.NewReader(payloadString)

	req, _ := http.NewRequest("PUT", url, payload)
	
	// Remember that the new Api tokens in gandi are scoped
	// This encreases the failure by 403 if mis-configured
	// So currectly show the user why the request failed
	req.Header.Add("authorization", authToken)
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		// Handle error that actually could happen
		return "500 Internal ERROR"
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if ds.DnsRecord[ItemNum].AAAA {
		ds.DnsRecord[ItemNum].currentIPS.ipv6 = ds.CurrentIPS.ipv6
	} else if ds.DnsRecord[ItemNum].A {
		ds.DnsRecord[ItemNum].currentIPS.ipv4 = ds.CurrentIPS.ipv4
	} else {
		// Config is not correct
		// This should not happen here
		return "500 Internal ERROR"
	}

	return string(body)
}

func (ds *DdnsSetting)GandiList() string {

	url := "https://api.gandi.net/v5/livedns/domains/example.com/records/www"
	// There should be no way that the request should error
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer pat_abc-123")

	res, err := ds.client.Do(req)
	if err != nil {
		// Handle error that actually could happen
		return "500 Internal ERROR"
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

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
