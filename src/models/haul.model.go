package models

import (
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Haul []HaulElement

type HaulElement struct {
	ID          int64  `json:"id"`         
	Title       string `json:"title"`      
}


func FetchHaul() (Response, error) {
	var res Response

	jsonFile, err := os.Open("./src/data/haul.json")

	if err != nil {
		return res, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var haul Haul

	json.Unmarshal(byteValue, &haul)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = haul

	return res, nil
}