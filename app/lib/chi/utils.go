package chi

import (
	"log"
	"net/http"

	"github.com/go-chi/render"

	xhttp "github.com/sknv/upsale/app/lib/net/http"
)

func DecodeRequest(w http.ResponseWriter, r *http.Request, v interface{}) {
	if err := render.DecodeJSON(r.Body, v); err != nil {
		log.Print("[ERROR] decode request: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		xhttp.AbortHandler()
	}
}
