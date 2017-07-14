package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tld-go/model"
)

// TopListTags renders the model TopListTags as HTML
func (wh *WebHandler) TopListTags(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var stats model.TopListTags
	err := wh.mgstats.Find(bson.M{"tag": "tags"}).One(&stats)
	if err != nil {
		panic(err)
	}
	wh.t.ExecuteTemplate(w, "toplist_tags.html", stats)
}
