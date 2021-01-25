package main

import (
	"errors"
	"net/http"
	"strconv"
)

func getIntParam(r *http.Request) (int, error) {
	if strN := r.URL.Query().Get("n"); len(strN) == 0 {
		return 0, errors.New("parameter n missing")
	} else if n, err := strconv.Atoi(strN); err != nil {
		return 0, errors.New("invalid parameter n")
	} else {
		return n, nil
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}
