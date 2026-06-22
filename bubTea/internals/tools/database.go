package tools

import (
	log "github.com/sirupsen/logrus"
)

type loginDetails struct {
	AuthToken string
	username  string
}
type ChoiceDetails struct {
	choice   string
	username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *loginDetails
	GetUserChoice(username string) *ChoiceDetails
	SetUpDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}
	var err error = database.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
