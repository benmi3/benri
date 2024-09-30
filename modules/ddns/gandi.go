package ddns

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (ds *DdnsSettings) GandiSingelByType(itemNum int, rrset_type, method string) string {
	//Creates a new record. Will raise a 409 conflict if the record already exists, and return a 200 OK if the record already exists with the same values
	// TODO: item size and itemNum check

	//url := "https://api.gandi.net/v5/livedns/domains/example.com/records/www/CNAME"
	url := fmt.Sprintf("https://api.gandi.net/v5/livedns/domains/%s/records/%s/%s", ds.Record[itemNum].Domain, ds.Record[itemNum].Name, rrset_type)

	var rrset_values string

	if ds.Record[itemNum].AAAA {
		rrset_values = ds.CurrentIPS.ipv6
	} else if ds.Record[itemNum].A {
		rrset_values = ds.CurrentIPS.ipv4
	} else {
		// Config is not correct
		// This should not happen here
		return "500 Internal ERROR"
	}

	authToken := fmt.Sprintf("Bearer %s", ds.AuthKey)

	payloadString := fmt.Sprintf("{\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}", rrset_values, ds.Record[itemNum].Ttl)

	payload := strings.NewReader(payloadString)

	req, _ := http.NewRequest(method, url, payload)

	// Remember that the new Api tokens in gandi are scoped
	// This encreases the failure by 403 if mis-configured
	// So currectly show the user why the request failed
	req.Header.Add("authorization", authToken)
	req.Header.Add("content-type", "application/json")

	res, err := ds.client.Do(req)
	if err != nil {
		// Handle error that actually could happen
		return "500 Internal ERROR"
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if ds.Record[itemNum].AAAA {
		ds.Record[itemNum].CurrentIPS.ipv6 = ds.CurrentIPS.ipv6
	} else if ds.Record[itemNum].A {
		ds.Record[itemNum].CurrentIPS.ipv4 = ds.CurrentIPS.ipv4
	} else {
		// Config is not correct
		// This should not happen here
		return "500 Internal ERROR"
	}

	return string(body)
}

func (ds *DdnsSettings) GandiCreateSingelByType(itemNum int, rrset_type string) string {

	return ds.GandiSingelByType(itemNum, rrset_type, "POST")
}

func (ds *DdnsSettings) GandiUpdateSingelByType(itemNum int, rrset_type string) string {

	return ds.GandiSingelByType(itemNum, rrset_type, "PUT")
}

func (ds *DdnsSettings) GandiUpdateMultipleByDomain(domain, payloadString, authToken string) string {
	//url := "https://api.gandi.net/v5/livedns/domains/example.com/records"
	url := fmt.Sprintf("https://api.gandi.net/v5/livedns/domains/%s/records", domain)

	payload := strings.NewReader(payloadString)

	// Creating domains is limited to one, so this will only be PUT anyway
	req, _ := http.NewRequest("PUT", url, payload)

	// Remember that the new Api tokens in gandi are scoped
	// This encreases the failure by 403 if mis-configured
	// So currectly show the user why the request failed
	req.Header.Add("authorization", authToken)
	req.Header.Add("content-type", "application/json")

	res, err := ds.client.Do(req)
	if err != nil {
		// Handle error that actually could happen
		return "500 Internal ERROR"
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body)
}

func (ds *DdnsSettings) GandiUpdateAll() string {

	// 	{
	//   "items": [
	//     {
	//       "rrset_name": "www",
	//       "rrset_type": "A",
	//       "rrset_values": [
	//         "192.0.2.1"
	//       ]
	//     },
	//     {
	//       "rrset_name": "www",
	//       "rrset_type": "AAAA",
	//       "rrset_values": [
	//         "2001:db8::1",
	//         "2001:db8::2"
	//       ]
	//     },
	//     {
	//       "rrset_name": "@",
	//       "rrset_type": "TXT",
	//       "rrset_values": [
	//         "\"v=spf1 include:_mailcust.gandi.net ?all\""
	//       ]
	//     }
	//   ]
	// }
	var needUpdate bool

	for i := 0; i < int(ds.RecordCount); i++ {
		// TODO: make this logic perfect
		needUpdate = false
		payloadFormat := "{\"items\":[\"%s\"]}"
		if ds.Record[i].A {
			needUpdate = true
			aRecord := fmt.Sprintf("{\"rrset_name\":\"A\",\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}", ds.CurrentIPS.ipv4, ds.Record[i].Ttl)

		}
		if ds.Record[i].AAAA {
			aaaaRecord := fmt.Sprintf("{\"rrset_name\":\"AAAA\",\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}", ds.CurrentIPS.ipv6, ds.Record[i].Ttl)
			if !needUpdate {
				needUpdate = true
			} else {
			}
		}
		payloadString := fmt.Sprintf("{\"items\":[\"%s\"]}")
	}

	var rrset_values string

	if ds.Record[itemNum].AAAA {
		rrset_values = ds.CurrentIPS.ipv6
	} else if ds.Record[itemNum].A {
		rrset_values = ds.CurrentIPS.ipv4
	} else {
		// Config is not correct
		// This should not happen here
		return "500 Internal ERROR"
	}

	authToken := fmt.Sprintf("Bearer %s", ds.AuthKey)

	payloadString := fmt.Sprintf("{\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}", rrset_values, ds.Record[itemNum].Ttl)

	payload := strings.NewReader("{\"rrset_name\":\"www\",\"rrset_type\":\"A\",\"rrset_values\":[\"192.0.2.1\"]}")

	payload := strings.NewReader(payloadString)

	req, _ := http.NewRequest("PUT", url, payload)

	// Remember that the new Api tokens in gandi are scoped
	// This encreases the failure by 403 if mis-configured
	// So currectly show the user why the request failed
	req.Header.Add("authorization", authToken)
	req.Header.Add("content-type", "application/json")

	res, err := ds.client.Do(req)
	if err != nil {
		// Handle error that actually could happen
		return "500 Internal ERROR"
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if ds.Record[itemNum].AAAA {
		ds.Record[itemNum].CurrentIPS.ipv6 = ds.CurrentIPS.ipv6
	} else if ds.Record[itemNum].A {
		ds.Record[itemNum].CurrentIPS.ipv4 = ds.CurrentIPS.ipv4
	} else {
		// Config is not correct
		// This should not happen here
		return "500 Internal ERROR"
	}

	return string(body)
}

func (ds *DdnsSettings) GandiList() string {

	url := "https://api.gandi.net/v5/livedns/domains/example.com/records"
	// There should be no way that the request should error
	req, _ := http.NewRequest("GET", url, nil)

	authToken := fmt.Sprintf("Bearer %s", ds.AuthKey)

	// Remember that the new Api tokens in gandi are scoped
	// This encreases the failure by 403 if mis-configured
	// So currectly show the user why the request failed
	req.Header.Add("authorization", authToken)

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
