package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func callUrl(method, url, body string, rc chan string) {
	if url == "" {
		// TODO: nice error handle
		return
	}

	var d bytes.Buffer
	if body != "" {
		json.Compact(&d, []byte(body))
	}

	req, _ := http.NewRequest(method, url, bytes.NewReader(d.Bytes()))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		rc <- fmt.Sprintf("Response: %v", err)
		return
	}

	t, err := io.ReadAll(res.Body)
	if err != nil {
		// TODO: nice error handle
		return
	}

	// NOTE: not sure accurate, but it's beta
	if strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
		rc <- jsonFormat(t)
	} else {
		rc <- string(t)
	}
}

func jsonFormat(b []byte) string {
	var d bytes.Buffer
	json.Indent(&d, b, "", "\t")
	return d.String()
}

func normalizeMethod(m string) string {
	var r string

	switch strings.ToUpper(m) {
	case http.MethodGet:
		r = " GET    | "
	case http.MethodPost:
		r = " POST   | "
	case http.MethodPut:
		r = " PUT    | "
	case http.MethodPatch:
		r = " PATCH  | "
	case http.MethodDelete:
		r = " DELETE | "
	}

	return r
}
