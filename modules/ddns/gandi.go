package ddns

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (ds *DdnsSettings) GandiSingelByType(itemNum int, rrset_type, rrset_value, method string) (int, error) {
	//Creates a new record. Will raise a 409 conflict if the record already exists, and return a 200 OK if the record already exists with the same values
	// TODO: item size and itemNum check
	if itemNum > int(ds.RecordCount) {
		return 0, fmt.Errorf("Item number does not exist")
	}

	//url := "https://api.gandi.net/v5/livedns/domains/example.com/records/www/CNAME"
	url := fmt.Sprintf("https://api.gandi.net/v5/livedns/domains/%s/records/%s/%s", ds.Record[itemNum].Domain, ds.Record[itemNum].Name, rrset_type)

	authToken := fmt.Sprintf("Bearer %s", ds.AuthKey)

	payloadString := fmt.Sprintf("{\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}", rrset_value, ds.Record[itemNum].Ttl)

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
		return 0, err
	}

	return res.StatusCode, nil
}

func (ds *DdnsSettings) GandiCreateSingelByType(itemNum int, rrset_type string) (int, error) {
	var rrset_value string

	switch rrset_type {
	case "A":
		rrset_value = ds.CurrentIPS.ipv4
		break
	case "AAAA":
		rrset_value = ds.CurrentIPS.ipv6
		break
	case "CNAME":
		rrset_value = ds.Record[itemNum].CNAME
		break
	default:
		return 0, fmt.Errorf("Unsupported record type was set")
	}
	reCode, err := ds.GandiSingelByType(itemNum, rrset_type, rrset_value, "POST")

	if err != nil || reCode != 200 {
		return reCode, err
	}

	switch rrset_type {
	case "A":
		ds.Record[itemNum].CurrentIPS.ipv4 = ds.CurrentIPS.ipv4
		break
	case "AAAA":
		ds.Record[itemNum].CurrentIPS.ipv6 = ds.CurrentIPS.ipv6
		break
	case "CNAME":
		break
	default:
		break

	}

	return 0, nil

}

func (ds *DdnsSettings) GandiUpdateSingelByType(itemNum int, rrset_type string) (int, error) {

	var rrset_value string

	switch rrset_type {
	case "A":
		rrset_value = ds.CurrentIPS.ipv4
		break
	case "AAAA":
		rrset_value = ds.CurrentIPS.ipv6
		break
	case "CNAME":
		rrset_value = ds.Record[itemNum].CNAME
		break
	default:
		return 0, fmt.Errorf("Unsupported record type was set")
	}
	reCode, err := ds.GandiSingelByType(itemNum, rrset_type, rrset_value, "PUT")

	if err != nil || reCode != 200 {
		return reCode, err
	}

	switch rrset_type {
	case "A":
		ds.Record[itemNum].CurrentIPS.ipv4 = ds.CurrentIPS.ipv4
		break
	case "AAAA":
		ds.Record[itemNum].CurrentIPS.ipv6 = ds.CurrentIPS.ipv6
		break
	case "CNAME":
		break
	default:
		break

	}

	return 0, nil
}

func (ds *DdnsSettings) GandiUpdateMultipleByDomain(domain, payloadString, authToken string) (int, error) {
	//url := "https://api.gandi.net/v5/livedns/domains/example.com/records"
	url := fmt.Sprintf("https://api.gandi.net/v5/livedns/domains/%s/records", domain)

	payload := strings.NewReader(payloadString)

	// Creating domains is limited to one, so this will only be PUT anyway
	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return 0, err
	}

	// Remember that the new Api tokens in gandi are scoped
	// This encreases the failure by 403 if mis-configured
	// So currectly show the user why the request failed
	req.Header.Add("authorization", authToken)
	req.Header.Add("content-type", "application/json")

	res, err := ds.client.Do(req)
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil

	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)

	// return string(body)
}

func (ds *DdnsSettings) GandiUpdateAll() (int, error) {

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

	authToken := fmt.Sprintf("Bearer %s", ds.AuthKey)

	for i := 0; i < int(ds.RecordCount); i++ {
		// TODO: make this logic perfect
		needUpdate = false
		payloadFormat := "{\"items\":[\"%s\"]}"
		payloadStr := ""
		payloadString := ""
		if ds.Record[i].A && ds.CurrentIPS.ipv4 != ds.Record[i].CurrentIPS.ipv4 {
			needUpdate = true
			payloadStr = fmt.Sprintf("{\"rrset_name\":\"A\",\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}", ds.CurrentIPS.ipv4, ds.Record[i].Ttl)

		}
		if ds.Record[i].AAAA && ds.CurrentIPS.ipv6 != ds.Record[i].CurrentIPS.ipv6 {
			aaaaRecord := fmt.Sprintf("{\"rrset_name\":\"AAAA\",\"rrset_values\":[\"%s\"],\"rrset_ttl\":%d}", ds.CurrentIPS.ipv6, ds.Record[i].Ttl)
			if !needUpdate {
				needUpdate = true
				payloadStr = aaaaRecord
			} else {
				payloadStr = payloadStr + "," + aaaaRecord
			}
		}
		payloadString = fmt.Sprintf(payloadFormat, payloadStr)
		if needUpdate {
			// Execute update logic
			// Else skip and refresh
			retStatusCode, err := ds.GandiUpdateMultipleByDomain(ds.Record[i].Domain, payloadString, authToken)
			if err != nil {
				// This error would be errors creating, sending or reciving requests
				// And have nothing to do with the http statusCode
				return i, err
			}
			if retStatusCode == 200 {
				// For now, the main logic will be
				// If the ip in the record is the same as main setting
				// it has been updated
				ds.Record[i].CurrentIPS = ds.CurrentIPS
			} else {
				// As the domain ips should be checked before we send the update request
				// There should only be 200 returned
				// All else needs to be checked and handled
				// Thats why we return here
				return i, fmt.Errorf("Unexpeced code returned on index: %d, Code: %d", i, retStatusCode)
			}

		}

	}

	return 0, nil
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
