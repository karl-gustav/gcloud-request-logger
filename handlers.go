// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/karl-gustav/gae-go-hello-worl/loraFieldTesterParser"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var (
	payloadHexRegex = regexp.MustCompile(`"payload_hex":\s?"([^"]+)"`)
)

func Run() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

type RequestDTO struct {
	TS   time.Time
	Body string
}

func (r *RequestDTO) MarshalJSON() ([]byte, error) {
	var parsed loraFieldTesterParser.LoraFieldTester
	payload := payloadHexRegex.FindAllStringSubmatch(r.Body, 1)
	if len(payload) > 0 {
		hexBytes, err := hex.DecodeString(payload[0][1])
		log.Printf("hex: % x", hexBytes)
		if err == nil {
			parsed = loraFieldTesterParser.ParseFieldTester(hexBytes)
		}
	}
	type Alias RequestDTO
	return json.Marshal(&struct {
		Parsed loraFieldTesterParser.LoraFieldTester `json:"parsedPayload"`
		*Alias
	}{
		Parsed: parsed,
		Alias:  (*Alias)(r),
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.Method == http.MethodPut || r.Method == http.MethodPost {
		key := datastore.NewIncompleteKey(ctx, "Request", nil)
		inRequest := new(RequestDTO)
		inRequest.TS = time.Now()
		b, _ := ioutil.ReadAll(r.Body)
		inRequest.Body = string(b)
		if _, err := datastore.Put(ctx, key, inRequest); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		j, _ := json.Marshal(inRequest)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "%s", j)
	} else if r.Method == http.MethodGet {
		var requests []RequestDTO
		q := datastore.NewQuery("Request").Order("-TS")
		_, err := q.GetAll(ctx, &requests)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		j, _ := json.Marshal(requests)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "%s", j)
	}

}
