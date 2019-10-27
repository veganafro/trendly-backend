package nyt

import (
	"log"
	"net"
	"time"
	"net/http"
	"encoding/json"
	"google.golang.org/grpc"
)

func Listen() {
	conn, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("server: Error creating tcp listener", err)
	}
	opts := make([]grpc.ServerOption, 0)
	grpcServer := grpc.NewServer(opts...)
	RegisterNytServiceServer(grpcServer, &NytServer{ Client: *NewClient() })
	grpcServer.Serve(conn)
}

func NewClient() *NytClient {
	client := &http.Client {
		CheckRedirect: func(request * http.Request, via[] * http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	endpoint := "https://api.nytimes.com/"
	nytClient := &NytClient{ Client: *client, Endpoint: endpoint }
	return nytClient
}

func GetJson(response *http.Response, target interface{}) (error) {
	return json.NewDecoder(response.Body).Decode(target)
}

