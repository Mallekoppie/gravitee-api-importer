package service

import (
	"net/http"

	"github.com/Mallekoppie/goslow/platform"
)

func FirstHelloWorld(w http.ResponseWriter, r *http.Request) {
	platform.Logger.Info("We arrived at a new world!!!!")

	w.Write([]byte("Hello World"))
}
