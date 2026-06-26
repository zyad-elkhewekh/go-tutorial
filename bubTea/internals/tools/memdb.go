package tools

import (
	"sync"
	//"time"
)

type memDB struct {
	mu      sync.Mutex
	choices map[string]ChoiceDetails
	logins  map[string]LoginDetails
}

func (d *memDB) GetUserLoginDetails(username string) *LoginDetails {
	d.mu.Lock()
	defer d.mu.Unlock()
	details, ok := d.logins[username]
	if !ok {
		return nil
	}
	return &details
}

func (d *memDB) SetUpUser(username string, choice string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.choices[username] = ChoiceDetails{
		Choice:   choice,
		Username: username,
	}
}

func (d *memDB) SetUserChoice(username, choice string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.choices[username] = ChoiceDetails{Choice: choice, Username: username}
}

func (d *memDB) GetUserChoice(username string) *ChoiceDetails {
	d.mu.Lock()
	defer d.mu.Unlock()
	c, ok := d.choices[username]
	if !ok {
		return nil
	}
	return &c
}

// SetUpDatabase just initializes the maps
func (d *memDB) SetUpDatabase() error {
	d.choices = make(map[string]ChoiceDetails)
	d.logins = make(map[string]LoginDetails)
	return nil
}
