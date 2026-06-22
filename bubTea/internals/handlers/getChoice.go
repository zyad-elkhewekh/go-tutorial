package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/zyad-elkhewekh/go-tutorial/bubTea/api"
	"github.com/zyad-elkhewekh/go-tutorial/bubTea/internals/tools"
)

func getChoices(w http.ResponseWriter, r *http.Request) {
	var params = api.ChoicesParam{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.ChoiceDetails
	tokenDetails = (*database).GetUserChoice(params.username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.ChoicesResponse{
		Information: (*tokenDetails).Choice,
		Code:        http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
