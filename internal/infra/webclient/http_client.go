package webclient

import (
	"io"
	"net/http"
)

func GET(url string) string {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
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

	return string(responseBody)
}
