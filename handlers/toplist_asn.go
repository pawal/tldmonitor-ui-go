package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tld-go/model"
)

// TopListASN renders the model TopListASN as HTML
func (wh *WebHandler) TopListASN(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stats model.TopListASN
	err := wh.mgstats.Find(bson.M{"tag": "asn"}).One(&stats)
	if err != nil {
		panic(err)
	}
	wh.t.ExecuteTemplate(w, "toplist_asn.html", stats)
}
