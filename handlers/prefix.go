package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tldmonitor-ui-go/model"
)

// Prefix finds all domains with the prefix matching the search
func (wh *WebHandler) Prefix(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var all []model.TLD
	prefix := p.ByName("id") + "/" + p.ByName("len")
	err := wh.mgc.Find(bson.M{"result.args.prefix": prefix}).Select(bson.M{"name": 1, "level": 1}).All(&all)
	if err != nil {
		panic(err)
	}
	t := model.DomainList{Title: "Domains matching prefix " + prefix, List: all}
	wh.t.ExecuteTemplate(w, "index.html", t)
}
