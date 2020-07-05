package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	//resp, err := http.Get(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("referer", url)
	req.Header.Set("user-agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36 SocketLog(tabid=982&client_id=)")

	req.Header.Set("sec-fetch-dest","ocument")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		return nil, err
	}
		defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Error: status code", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}
