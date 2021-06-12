package models

import (
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Shalawat []ShalawatElement

type ShalawatElement struct {
	ID          int64  `json:"id"`         
	Title       string `json:"title"`      
	Arabic      string `json:"arabic"`     
	Translation string `json:"translation"`
}


func FetchShalawat() (Response, error) {
	var res Response

	jsonFile, err := os.Open("./src/data/Shalawat.json")

	if err != nil {
		return res, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var shalawat Shalawat

	json.Unmarshal(byteValue, &shalawat)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = shalawat

	return res, nil
}