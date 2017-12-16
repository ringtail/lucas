package handlers

import (
	"net/http"
	"github.com/ringtail/lucas/frontend"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(frontend.HOME_PAGE))
}
