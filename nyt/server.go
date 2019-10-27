package nyt

import (
	"log"
	context "context"
)

type NytServer struct {
	UnimplementedNytServiceServer
	Client NytClient
}

func (nytServer * NytServer) GetMostShared(ctx context.Context, request * NytRequest) (*NytResponse, error) {
	response, error := nytServer.Client.GetMostShared(request.Token, request.Period)
	if error != nil {
		log.Fatal("server: Error executing request for most shared", error)
	}
	json := NytMostShared{}
	GetJson(response, &json)
	return &NytResponse{ Response: json.Response }, error
}

