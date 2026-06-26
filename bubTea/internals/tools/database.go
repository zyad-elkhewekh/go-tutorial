package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	username  string
}
type ChoiceDetails struct {
	Choice   string
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserChoice(username string) *ChoiceDetails
	SetUpDatabase() error
	SetUpUser(username string, choice string)
}

var globalDB DatabaseInterface

func NewDatabase() (*DatabaseInterface, error) {
	// var database DatabaseInterface = &mockDB{}
	// var err error = database.SetUpDatabase()

	if globalDB != nil {
		return &globalDB, nil
	}
	globalDB = &memDB{}
	err := globalDB.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &globalDB, nil
}
