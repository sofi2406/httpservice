package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	data struct {
		Name  string `json:"id"`
		Sound string `json:"value"`
	}
)

var dataset []data

func init() {
	wd, err := os.Getwd()
	if err != nil {
		er(err)
	}

	file, err := os.Open(fmt.Sprintf("%v/database/data.json", wd))
	if err != nil {
		er(err)
	}

	byteFile, _ := ioutil.ReadAll(file)
	if err != nil {
		er(err)
	}

	err = json.Unmarshal(byteFile, &dataset)
	if err != nil {
		er(err)
	}
}

func GetDataSet(name string) (*data, error) {
	for _, v := range dataset {
		if name == v.Name {
			return &v, nil
		}
	}
	return nil, errors.New("No data found")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
