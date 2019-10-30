package handlers

import (
	"encoding/json"
	"example-rest-api/dummy_db"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func encode(w http.ResponseWriter, obj interface{}, errorCode int) {
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(obj)
}

func AddLink(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		encode(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = dummy_db.NewLink(body)

	if err != nil {
		encode(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encode(w, nil, http.StatusCreated)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	link, err := dummy_db.Links.FindById(id)

	if err != nil {
		encode(w, err, http.StatusNotFound)
		return
	}

	encode(w, link, http.StatusAccepted)

}

func RedirectToRealURL(w http.ResponseWriter, r *http.Request) {

	shortName := mux.Vars(r)["url"]

	realUrl, err := dummy_db.Links.FindByName(shortName)

	if err != nil {
		encode(w, err, http.StatusNotFound)
		return
	}

	http.Redirect(w, r, realUrl.RealURL, http.StatusTemporaryRedirect)

	log.Printf("Link with short name \"%s\" was redirected to \"%s\"\n",
		shortName, realUrl.RealURL)
}
