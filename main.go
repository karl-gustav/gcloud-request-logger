// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

type RequestDTO struct {
	TS      time.Time
	Body    string
	Headers []Header
}

type Header struct {
	Key   string
	Value string
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.Method == http.MethodPut || r.Method == http.MethodPost {
		key := datastore.NewIncompleteKey(ctx, "Request", nil)
		inRequest := new(RequestDTO)
		inRequest.TS = time.Now()
		b, _ := ioutil.ReadAll(r.Body)
		inRequest.Body = string(b)
		for key, value := range r.Header {
			inRequest.Headers = append(inRequest.Headers, Header{key, strings.Join(value, ", ")})
		}
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
