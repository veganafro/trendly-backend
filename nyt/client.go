package nyt

import (
	"log"
	"path"
	"strconv"
	"net/http"
)

type NytClient struct {
	Client http.Client
	Endpoint string
}

type NytMostShared struct {
	Response string
}

func (nytClient * NytClient) GetMostShared(token string, period int32) (*http.Response, error) {
	request, error := http.NewRequest("GET", nytClient.Endpoint, nil)

	if error != nil {
		log.Fatal("client: Error creating new request", error)
	}

	request.URL.Path = path.Join(
		request.URL.Path,
		"svc/mostpopular/v2/shared",
		strconv.Itoa(int(period)),
	)

	query := request.URL.Query()
	query.Add("api-key", token)
	request.URL.RawQuery = query.Encode()

	request.Header.Add("Accept", "application/json")

	response, error := nytClient.Client.Do(request)
	defer response.Body.Close()

	if error != nil {
		log.Fatal("client: Error executing request for most shared", error)
	}
	log.Println("client: Request status:", response.Status)
	return response, error
}

