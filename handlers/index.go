package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tldmonitor-ui-go/model"
)

// Index renders the index web page
func (wh *WebHandler) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var all []model.TLD
	err := wh.mgc.Find(nil).Select(bson.M{"name": 1, "level": 1}).All(&all)
	if err != nil {
		panic(err)
	}
	t := model.DomainList{Title: "All TLDs", List: all}
	wh.t.ExecuteTemplate(w, "index.html", t)
}
