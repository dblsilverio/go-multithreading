package webclient

import (
	"errors"
	"io"
	"net/http"
)

func GET(url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		if err == nil {
			return "", errors.New("invalid response from gateway")
		}
		return "", err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(responseBody), nil
}
