package ddns

import (
	"log"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

func main() {
	// Not sure
	qjar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: qjar,
	}

	//req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	//resp, err := client.Do(req)
	// ...

	GandiUpdate(client)

	//CloudflareUpdate()

}
