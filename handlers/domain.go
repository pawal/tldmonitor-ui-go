package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tldmonitor-ui-go/model"
	"gopkg.in/mgo.v2/bson"
)

// Domain renders the page with the log from a domain
func (wh *WebHandler) Domain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	var tld model.TLD
	domain := p.ByName("id")
	w.Header().Set("Cache-Control", "max-age=3600, must-revalidate")

	// Get entry from MongoDB
	err := wh.mgc.Find(bson.M{"name": domain}).One(&tld)
	if err != nil {
		fmt.Println(err)
	}
	if tld.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// if render JSON
	if queryValues.Get("json") != "" {
		wh.renderJSON(w, tld)
	} else {
		// prepare log entries and format args
		var newRes []model.Entry
		for _, res := range tld.Result {
			res.ArgsString = template.HTML(formatArgs(res.Args))
			newRes = append(newRes, res)
		}
		// replace Result with formatted log
		tld.Result = newRes
		// render HTML
		wh.t.ExecuteTemplate(w, "domain.html", tld)
		return
	}
}

// format the args as a string for rendering
func formatArgs(args map[string]string) (res string) {
	for key, value := range args {
		switch key {
		case "nsnlist":
			res += formatNsnlist(value)
		case "nsset":
			res += formatNsset(value)
		case "glue":
			res += formatGlue(value)
		case "ns":
			res += formatNs(value)
		case "address":
			res += formatAddress(value)
		case "prefix":
			res += formatPrefix(value)
		case "asn":
			res += formatASN(key, value)
		default:
			res += fmt.Sprintf("%s: %s<br/>", key, value)
		}
	}
	if res != "" {
		// fmt.Println(res)
	}
	return
}

// removeTrailingDot removes any trailing dot from a hostname
func removeTrailingDot(ns string) (res string) {
	return strings.TrimSuffix(ns, ".")
}

// nsnlist is a list of nameserver names separated by ","
func formatNsnlist(value string) (res string) {
	var s string
	for _, ns := range strings.Split(value, ",") {
		ns = template.URLQueryEscaper(removeTrailingDot(ns))
		s += fmt.Sprintf("<a href=\"/ns/%s\">%s</a><br/>", ns, ns)
	}
	return fmt.Sprintf("nsnlist: %s<br/>", s)
}

// nsset is a list of nameserver names separated by ","
func formatNsset(value string) (res string) {
	var s string
	for _, ns := range strings.Split(value, ",") {
		ns = template.URLQueryEscaper(removeTrailingDot(ns))
		s += fmt.Sprintf("<a href=\"/ns/%s\">%s</a><br/>", ns, ns)
	}
	return fmt.Sprintf("nsset: %s<br/>", s)
}

// glue is a list of glue nameserver records separated by ";"
func formatGlue(value string) (res string) {
	var s string
	for _, str := range strings.Split(value, ";") {
		str = template.URLQueryEscaper(str)
		s += fmt.Sprintf("<a href=\"/ns/%s\">%s</a><br/>", str, str)
	}
	return fmt.Sprintf("glue: %s<br/>", s)
}

// ns is either a nameserver name, or a ns and address separated by "/"
// can also be a list of ns separated by ";"...
func formatNs(value string) (res string) {
	i := strings.Index(value, "/")
	// If the ns args contains both ns and address separated by "/"
	// TODO: URL escape strings
	if i != -1 {
		ns := fmt.Sprintf("<a href=\"/ns/%s\">%s</a>",
			value[0:i], value[0:i])
		a := fmt.Sprintf("<a href=\"/address/%s\">%s</a>",
			value[i+1:len(value)], value[i+1:len(value)])
		return fmt.Sprintf("ns: %s %s<br/>", ns, a)
	}

	// If the ns is a list of ns separated by ";"
	if strings.Contains(value, ";") {
		var s string
		for _, str := range strings.Split(value, ";") {
			str = template.URLQueryEscaper(str)
			s += fmt.Sprintf("<a href=\"/ns/%s\">%s</a><br/>", str, str)
		}
		return fmt.Sprintf("ns: %s<br/>", s)
	}

	// only the ns name
	value = template.URLQueryEscaper(value)
	res = fmt.Sprintf("ns: <a href=\"/ns/%s\">%s</a><br/>", value, value)
	return res
}

// address is either an ipv4 or ipv6 address
func formatAddress(value string) (res string) {
	v := template.URLQueryEscaper(value)
	return fmt.Sprintf("address: <a href=\"/address/%s\">%s</a><br/>", v, value)
}

// address is either an ipv4 or ipv6 address
func formatPrefix(value string) (res string) {
	v := template.URLQueryEscaper(value)
	return fmt.Sprintf("prefix: <a href=\"/prefix/%s\">%s</a><br/>", v, value)
}

// asn is either a single value, or comma separated list of values
func formatASN(key, value string) (res string) {
	// if comma separated list
	if strings.Contains(value, ",") {
		for _, str := range strings.Split(value, ",") {
			str = template.URLQueryEscaper(str)
			res += fmt.Sprintf("<a href=\"/asn/%s\">%s</a> ", str, str)
		}
		return "asn: " + res + "<br/>"
	}
	// single value
	value = template.URLQueryEscaper(value)
	return fmt.Sprintf("asn: <a href=\"/asn/%s\">%s</a><br/>", value, value)
}
