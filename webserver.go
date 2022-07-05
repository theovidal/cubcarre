package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"log"
	"net/http"
	"strconv"

	"github.com/theovidal/cubcarre/lib"
)

func validMAC(bot *lib.Bot, dataString, hash string) bool {
	token := hmac.New(sha256.New, []byte("WebAppData"))
	token.Write([]byte(bot.Token))
	key := token.Sum(nil)

	received := hmac.New(sha256.New, key)
	received.Write([]byte(dataString))
	expectedData := received.Sum(nil)

	return hmac.Equal([]byte(hash), expectedData)
}

func webServer(bot lib.Bot) {
	http.HandleFunc("/save", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		/*
			dataString := req.FormValue("data_string")
			hash := req.FormValue("hash")
			if !validMAC(&bot, dataString, hash) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}*/

		queryID := req.FormValue("queryID")
		userID, _ := strconv.ParseInt(req.FormValue("userID"), 10, 0)
		time, _ := strconv.ParseUint(req.FormValue("time"), 10, 0)
		cube := req.FormValue("cube")
		if queryID == "" || userID == 0 || time == 0 || cube == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := bot.WebCallback(&bot, queryID, userID, time, cube); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
