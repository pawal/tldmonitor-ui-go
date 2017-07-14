package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tld-go/model"
)

// Tag finds all domain with the tag and renders the index web page (domain list)
func (wh *WebHandler) Tag(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var all []model.TLD
	tag := p.ByName("id")
	err := wh.mgc.Find(bson.M{"result.tag": tag}).Select(bson.M{"name": 1, "level": 1}).All(&all)
	if err != nil {
		panic(err)
	}
	t := model.DomainList{Title: "Domains matching tag " + tag, List: all}
	wh.t.ExecuteTemplate(w, "index.html", t)
}
