package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bbrodriges/mielofon"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil || len(b) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var input mielofon.Input
		if err := json.Unmarshal(b, &input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var keywordFound bool
		for _, token := range input.Request.Nlu.Tokens {
			if token == "привет" {
				keywordFound = true
			}
		}

		var output mielofon.Output
		output.Version = input.Version
		output.Session = input.Session
		output.Response.Text = "Непонятно!"

		if keywordFound {
			output.Response.Text = "Привет!"
			output.Response.EndSession = true
		}

		output.Response.Tts = output.Response.Text

		b, err = json.Marshal(output)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	})

	http.ListenAndServe(":80", nil)
}
