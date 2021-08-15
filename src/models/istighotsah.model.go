package models

import (
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Istighotsah []IstighotsahElement

type IstighotsahElement struct {
	ID          int64  `json:"id"`         
	Title       string `json:"title"`      
	Arabic      string `json:"arabic"`     
	Translation string `json:"translation"`
}


func FetchIstighotsah() (Response, error) {
	var res Response

	jsonFile, err := os.Open("./src/data/istighotsah.json")

	if err != nil {
		return res, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var istighotsah Istighotsah

	json.Unmarshal(byteValue, &istighotsah)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = istighotsah

	return res, nil
}