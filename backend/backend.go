package backend

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ringtail/lucas/backend/handlers"
	"golang.org/x/net/context"
	"github.com/ringtail/lucas/backend/types"
	log "github.com/Sirupsen/logrus"
)

type LucasServer struct {
	Handler *mux.Router
}

func (ls *LucasServer) Start(opts *types.Opts) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/store", handlers.StoreHandler)
	contextMux := ls.Middleware(opts,mux)
	log.Fatal(http.ListenAndServe(":8080", contextMux))
}


func (ls *LucasServer) Middleware(opts *types.Opts, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "opts", opts)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
