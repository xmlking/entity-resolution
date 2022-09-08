package main

import (
	"log"
	"net/http"

	"github.com/xmlking/entity-resolution/gen/go/er/service/entity/v1/entityv1connect"
	"github.com/xmlking/entity-resolution/service/entity/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	entityHandler, _ := handler.NewEntityServiceHandler()
	mux.Handle(entityv1connect.NewEntityServiceHandler(entityHandler))
	err := http.ListenAndServe(
		"0.0.0.0:8080",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		//h2c.NewHandler(mux, &http2.Server{}),
		h2c.NewHandler(newCORS().Handler(mux), &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
