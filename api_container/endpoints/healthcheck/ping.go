package healthcheck

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping() httprouter.Handle {
	return func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		log.Println("Started: check health of server")
		writer.WriteHeader(http.StatusOK)
	}
}
