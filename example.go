package example

import (
	"fmt"
	"net/http"

	utils "github.com/1backend/go-utils"
	httpr "github.com/julienschmidt/httprouter"
)

func Hi(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	utils.Write(w, "Hi! This endpoint is using an imported service")
}

// This Rectangle type is not the same as what is used in the "types" section of this service on 1backend.com
// The other one is used to generate clients, this one is for the server.
// They are intentionally kepts separate for now.
type Rectangle struct {
	SideA int
	SideB int
}

func CalculateRectangleArea(w http.ResponseWriter, r *http.Request, p httpr.Params) {
	input := struct {
		Unit string
		Rect Rectangle
	}{}
	err := utils.ReadJsonBody(r, &input)
	if err != nil {
		utils.Write400(w, err)
		return
	}
	area := input.Rect.SideA * input.Rect.SideA
	utils.Write(w, fmt.Sprintf("%v%v", area, input.Unit))
	return
}
