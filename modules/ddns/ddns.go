package ddns

import (
	"log"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

func main() {
	// Not sure if im keeping this
	qjar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: qjar,
	}

	resp, err := client.Get("http://example.com")
	// ...

	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	// ...

	GandiUpdate()

	CloudflareUpdate()

}
