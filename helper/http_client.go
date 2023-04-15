package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func RequestGet(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, fmt.Errorf("error from rpi with code %s, response: %s", strconv.Itoa(resp.StatusCode), string(body))
	}

	return body, nil
}
