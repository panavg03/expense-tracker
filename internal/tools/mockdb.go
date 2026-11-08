package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Panav": {
		AuthToken: "p1",
		Username:  "panav",
	},
	"Panav2": {
		AuthToken: "p2",
		Username:  "panav2",
	},
	"Panav3": {
		AuthToken: "p3",
		Username:  "panav3",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Panav": {
		Coins:    1000,
		Username: "panav",
	},
	"Panav2": {
		Coins:    2000,
		Username: "panav2",
	},
	"Panav3": {
		Coins:    3000,
		Username: "panav3",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
