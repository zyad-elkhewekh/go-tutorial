package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/zyad-elkhewekh/go-tutorial/bubTea/api"
	"github.com/zyad-elkhewekh/go-tutorial/bubTea/internals/tools"
)

var unautherr = errors.New("invalid token")

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorize")
		var err error

		if username == "" || token == "" {
			log.Error(unautherr)
			api.RequestErrorHandler(w, unautherr)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(unautherr)
			api.RequestErrorHandler(w, unautherr)
			return
		}

		next.ServeHTTP(w, r)
	})
}
