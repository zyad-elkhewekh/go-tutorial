package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
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
		Choice:   "cpp",
		Username: "zyad",
	},
	"salah": {
		Choice:   "java",
		Username: "salah",
	},
}

// func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
// 	time.Sleep(time.Second * 1)

// 	var cleintData = LoginDetails{}
// 	cleintData, ok := mockLoginDetails[username]
// 	if !ok {
// 		return nil
// 	}
// 	return &cleintData
// }

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

// func (d *mockDB) SetUpUser(username string, choice string) {
// 	mockChoiceDetails[username] = ChoiceDetails{
// 		Choice:   choice,
// 		Username: username,
// 	}
// }
