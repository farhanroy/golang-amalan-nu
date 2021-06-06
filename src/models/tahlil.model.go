package models

import (
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Tahlil []TahlilElement

type TahlilElement struct {
	ID          int64  `json:"id"`         
	Title       string `json:"title"`      
	Arabic      string `json:"arabic"`     
	Translation string `json:"translation"`
}


func FetchTahlil() (Response, error) {
	var res Response

	jsonFile, err := os.Open("./src/data/tahlil.json")

	if err != nil {
		return res, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var tahlil Tahlil

	json.Unmarshal(byteValue, &tahlil)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = tahlil

	return res, nil
}