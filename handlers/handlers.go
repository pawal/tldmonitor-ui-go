package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

// WebHandler contains handler info (mongo db) for the handlers
type WebHandler struct {
	mgc     *mgo.Collection
	mgstats *mgo.Collection
	t       *template.Template
}

// InitWebHandler connects to the mongo db, and adds the two collections
func InitWebHandler(t *template.Template, mdb, mgc, stats string) (*WebHandler, error) {
	var wh WebHandler
	wh.t = t // template stuff for HTML rendering

	session, err := mgo.Dial(viper.GetString("mongourl"))
	if err != nil {
		return nil, err
	}
	wh.mgc = session.DB(mdb).C(mgc)
	wh.mgstats = session.DB(mdb).C(stats)
	return &wh, err
}

// internal JSON renderer that sets the correct content-type
func (wh *WebHandler) renderJSON(w http.ResponseWriter, v interface{}) error {
	j, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", j)
	return nil
}
