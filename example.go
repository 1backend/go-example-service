package example

import (
	"net/http"

	utils "github.com/1backend/go-utils"
	httpr "github.com/julienschmidt/httprouter"
)

func Hi(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	utils.Write(w, "Hi! This endpoint is using an imported service")
}
