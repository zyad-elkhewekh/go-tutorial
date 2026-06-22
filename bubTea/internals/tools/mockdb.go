package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]loginDetails{
	"zyad": {
		AuthToken: "123abc",
		username:  "zyad",
	},
	"salah": {
		AuthToken: "456def",
		username:  "salah",
	},
}
var mockChoiceDetails = map[string]ChoiceDetails{
	"zyad": {
		choice:   "cpp",
		username: "zyad",
	},
	"salah": {
		choice:   "java",
		username: "salah",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *loginDetails {
	time.Sleep(time.Second * 1)

	var cleintData = loginDetails{}
	cleintData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &cleintData
}

func (d *mockDB) GetUserChoice(username string) *ChoiceDetails {
	time.Sleep(time.Second * 1)

	var cleintData = ChoiceDetails{}
	cleintData, ok := mockChoiceDetails[username]
	if !ok {
		return nil
	}
	return &cleintData
}

func (d *mockDB) SetUpDatabase() error {
	return nil
}
