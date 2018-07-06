package jsonaggre

import (
	"encoding/json"
	"io/ioutil"
)

type TestData1 struct {
	Name     string
	Id       int
	Personas []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ProcessJson(fileName string) TestData1 {
	rawData, err := ioutil.ReadFile(fileName)
	check(err)
	var payload TestData1

	json.Unmarshal(rawData, &payload)
	return payload
}
