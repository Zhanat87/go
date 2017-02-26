package model

import (
	"testing"
	"encoding/json"
)

func TestGetsTheCorrectCurrencies(t *testing.T) {
	// arrange

	// act
	result := GetCurrencies()

	// assert
	resultJson, _ := json.Marshal(result)
	expected, _ := json.Marshal([]*Currency{
		{ID: 1, Name: "USD"},
		{ID: 2, Name: "EUR"},
	})
	if string(resultJson) != string(expected) {
		t.Error("Did not get expected currencies")
	}
}
