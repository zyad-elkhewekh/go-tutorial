package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zyad-elkhewekh/go-tutorial/bubTea/api"
	"github.com/zyad-elkhewekh/go-tutorial/bubTea/internals/tools"
)

func submitChoice(w http.ResponseWriter, r *http.Request) {
	var body api.SubmitChoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	db, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	(*db).SetUpUser(body.Username, body.Choice)
	w.WriteHeader(http.StatusCreated)
}
