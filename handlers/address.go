package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tldmonitor-ui-go/model"
)

// Address finds all domains with the address matching the search
func (wh *WebHandler) Address(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var all []model.TLD
	tag := p.ByName("id")
	// err := wh.mgc.Find(bson.M{"result.args.asn": bson.RegEx{Pattern: ".*" + tag + ".*", Options: ""}}).Select(bson.M{"name": 1, "level": 1}).All(&all)
	err := wh.mgc.Find(bson.M{"result.args.address": tag}).Select(bson.M{"name": 1, "level": 1}).All(&all)
	if err != nil {
		panic(err)
	}
	t := model.DomainList{Title: "Domains matching address " + tag, List: all}
	wh.t.ExecuteTemplate(w, "index.html", t)
}
