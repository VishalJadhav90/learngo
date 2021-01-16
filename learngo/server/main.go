package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var g int

type prediction_struct struct {
	ClassName   string
	Probability float32
}

type result_struct struct {
	Origin      string
	Src         string
	Predictions []prediction_struct
}

type nsfwData struct {
	Porn   float32
	Sexy   float32
	Hentai float32
}

type configData struct {
	NSFW_CATEGORIES        nsfwData
	PULL_CONFIG_AFTER_MINS int
	MODEL_PATH             string
	IMAGE_SIZE             int
	RANDOM_IMGS            int
	MAX_IMGS               int
	STOP_SCRAPING          bool
	STOP_SCANING           bool
	LOG_LEVEL              string
}

func main() {
	g = 0

	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var t result_struct
		decoder.Decode(&t)
		log.Println("------------")
		log.Println("Origin:", t.Origin)
		log.Println("Src:", t.Src)
		log.Println("Predictions:", t.Predictions)
	})

	http.HandleFunc("/config", func(res http.ResponseWriter, req *http.Request) {
		log.Println("------------")
		// STOP_SCRAPING & STOP_SCANING will be populated from basegen of policy assigned to user
		g = g + 1
		var STOP_SCRAPING bool
		var STOP_SCANNING bool
		if false {
			STOP_SCRAPING = true
			STOP_SCANNING = true
		} else {
			STOP_SCRAPING = false
			STOP_SCANNING = false
		}
		confData := configData{nsfwData{0.9, 0.9, 0.9}, 5, "../model", 299, 5, 15, STOP_SCRAPING, STOP_SCANNING, "DEBUG"}
		confJson, err := json.Marshal(confData)
		if err != nil {
			log.Println("Error: %s", err)
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(confJson)
	})

	// run server on port "9000"
	log.Fatal(http.ListenAndServe(":9000", nil))

}
