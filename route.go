package main

import (
	"github.com/Mallekoppie/goslow/platform"
	"github.com/Mallekoppie/gravitee-api-importer/service"
	"net/http"
)

var Routes = platform.Routes{
	platform.Route{
		Path:        "/",
		Method:      http.MethodGet,
		HandlerFunc: service.FirstHelloWorld,
		SlaMs:       0,
	},
}
