package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	schema_entity_v1 "github.com/xmlking/entity-resolution/gen/go/er/schema/entity/v1"
	"github.com/xmlking/entity-resolution/gen/go/er/service/entity/v1/entityv1connect"
	// "github.com/xmlking/entity-resolution/internal/redis"
)

// TODO: https://cobra.dev/

func main() {
	// redis.Init()
	client := entityv1connect.NewEntityServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	names := []*schema_entity_v1.Name{
		{
			First: "sumo",
			Last:  "demo",
		},
		{
			First: "Dan",
			Last:  "Jose",
		},
	}

	req := connect.NewRequest(&schema_entity_v1.Member{
		ExternalId: "123e4567-e89b-12d3-a456-426614174000",
		Names:      names,
	})
	req.Header().Set("Some-Header", "hello from connect")
	res, err := client.Ingest(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(*res.Msg.Key)
	log.Println(res.Header().Get("Some-Other-Header"))
}
