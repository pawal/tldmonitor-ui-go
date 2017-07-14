package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/pawal/tld-go/handlers"
	"github.com/spf13/viper"
)

func main() {
	_ = godotenv.Load(".env")
	viper.AutomaticEnv()
	viper.SetDefault("listen", "localhost:5000")
	viper.SetDefault("mongourl", "mongodb://localhost")
	viper.SetDefault("mongodb", "results")
	viper.SetDefault("mongocollection", "tlds")
	viper.SetDefault("mongostats", "tlds_stats")
	viper.SetDefault("template", "./template")

	tmpl, err := template.ParseGlob(viper.GetString("template") + "/*.html")
	if err != nil {
		panic(err)
	}

	wh, err := handlers.InitWebHandler(tmpl,
		viper.GetString("mongodb"),
		viper.GetString("mongocollection"),
		viper.GetString("mongostats"))
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	router.GET("/", wh.Index)
	router.GET("/domain/:id", wh.Domain)
	router.GET("/tag/:id", wh.Tag)
	router.GET("/asn/:id", wh.ASN)
	router.GET("/address/:id", wh.Address)
	router.GET("/ns/:id", wh.NS)
	router.GET("/toplist/ns", wh.TopListNS)
	router.GET("/toplist/asn", wh.TopListASN)
	router.GET("/toplist/tags", wh.TopListTags)
	router.GET("/about/zonemaster", wh.AboutZonemaster)
	router.GET("/about/tldmonitor", wh.AboutTLDMonitor)
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(viper.GetString("listen"), router))

}
