// Package alllcoin provides Coins struct with methods to operate on a state of
// itself find coins for a symbols without explicit delimiter

// Holds Constructors and Storing methods for Coins struct
package allcoin

import (
	"encoding/json"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"os"
)

// Url to pull all of the most recet coins
const AllCoinsUrl = "https://min-api.cryptocompare.com/data/all/coinlist"

// Building Coins object from API response
func NewFromAPI() (Coins, error) {
	type ApiCoins struct {
		Data Coins `json:Data`
	}

	var allCoins ApiCoins

	resp, err := resty.R().Get(AllCoinsUrl)

	if err != nil {
		return allCoins.Data, err
	}

	err = json.Unmarshal(resp.Body(), &allCoins)

	if err != nil {
		return allCoins.Data, err
	}

	return allCoins.Data, nil
}

// Building Coins object from stored JSON file
func NewFromJSON(fileName string) (Coins, error) {
	var cs Coins

	dat, err := ioutil.ReadFile(fileName)

	if err != nil {
		return cs, err
	}

	err = json.Unmarshal(dat, &cs)

	return cs, nil
}

// Storing currently operated coins in a JSON format
func WriteToFile(cs Coins, fileName string) error {
	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	jsonEncodedCoins, err := cs.EncodeJSON()

	if err != nil {
		return err
	}

	_, err = f.WriteString(jsonEncodedCoins)

	if err != nil {
		return err
	}

	return nil
}
