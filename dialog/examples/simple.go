package main

import (
	"encoding/json"
	"github.com/bbrodriges/mielofon"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input, output, err := mielofon.GetDialogPair(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		output.Response.Text = "Непонятно!"
		if input.HasKeyword("привет") {
			output.Response.Text = "Привет!"
			output.Response.EndSession = true
		}
		output.Response.Tts = output.Response.Text

		b, err := json.Marshal(output)
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

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
