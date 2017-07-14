package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tld-go/model"
)

// TopListNS renders the model TopListNS as HTML
func (wh *WebHandler) TopListNS(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stats model.TopListNS
	err := wh.mgstats.Find(bson.M{"tag": "ns"}).One(&stats)
	if err != nil {
		panic(err)
	}
	wh.t.ExecuteTemplate(w, "toplist_ns.html", stats)
}
