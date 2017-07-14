package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AboutZonemaster renders the about page about Zonemaster
func (wh *WebHandler) AboutZonemaster(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	wh.t.ExecuteTemplate(w, "zonemaster.html", nil)
}

// AboutTLDMonitor renders the about page about TLDMonitor
func (wh *WebHandler) AboutTLDMonitor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	wh.t.ExecuteTemplate(w, "tldmonitor.html", nil)
}
