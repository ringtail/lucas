package handlers

import (
	"net/http"
	"github.com/ringtail/lucas/backend/services"
	"encoding/json"
	"github.com/ringtail/lucas/backend/types"
	log "github.com/Sirupsen/logrus"
)

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	var options *types.Opts
	var store *services.Store
	var err error
	if opts := r.Context().Value("opts"); opts != nil {
		options = opts.(*types.Opts)
	} else {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	if options.Key == "" || options.Cert == "" || options.Ca == "" {
		store, err = services.NewWithOutTLS(options.Endpoints)
	} else {
		store, err = services.New(options.Endpoints, options.Ca, options.Key, options.Cert)
	}
	if err != nil {
		log.Errorf("Failed to init etcd client because of %s", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	if r.Method == http.MethodGet {
		key := r.URL.Query().Get("key")
		keys := store.List(key)
		bytes, err := json.Marshal(keys)
		if err != nil {
			log.Errorf("Failed to list etcd keys because of %s", err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Write(bytes)
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		key := r.FormValue("key")
		value := r.FormValue("value")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := store.Put(key, value)
		if err != nil {
			log.Errorf("Failed to update value because of %s", err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	if r.Method == http.MethodDelete {
		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := store.Delete(key)
		if err != nil {
			log.Errorf("Failed to delete value because of %s", err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
