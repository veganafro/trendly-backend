package nyt

import (
	"log"
	"time"
	"path"
	"strconv"
	"net/http"
)

type NytClient struct {
	Client http.Client
	Endpoint string
	Token string
}

func NewClient(token string) *NytClient {
	client := &http.Client {
		CheckRedirect: func(request * http.Request, via[] * http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	endpoint := "https://api.nytimes.com/"
	nytClient := &NytClient{ Client: *client, Endpoint: endpoint, Token: token }
	return nytClient
}

func (nytClient * NytClient) GetMostShared(period int) {
	request, error := http.NewRequest("GET", nytClient.Endpoint, nil)

	if error != nil {
		log.Fatal("Error creating new request", error)
	}

	request.URL.Path = path.Join(
		request.URL.Path,
		"svc/mostpopular/v2/shared",
		strconv.Itoa(period),
	)

	query := request.URL.Query()
	query.Add("api-key", nytClient.Token)
	request.URL.RawQuery = query.Encode()

	request.Header.Add("Accept", "application/json")

	response, error := nytClient.Client.Do(request)
	defer response.Body.Close()

	if error != nil {
		log.Fatal("Error executing request for most shared", error)
	}
	log.Println("Request status:", response.Status)
}

