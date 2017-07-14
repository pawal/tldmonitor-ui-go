package model

import (
	"html/template"

	"gopkg.in/mgo.v2/bson"
)

// Entry is used for all log entries for a TLD
type Entry struct {
	Timestamp  float64           `json:"timestamp" bson:"timestamp"`
	Module     string            `json:"module" bson:"module"`
	Level      string            `json:"level" bson:"level"`
	Tag        string            `json:"tag" bson:"tag"`
	Args       map[string]string `json:"args" bson:"args"`
	ArgsString template.HTML     // Only used for rendering results
}

// TLD is the complete log for a domain
type TLD struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Level  string        `json:"level" bson:"level"`
	Result []Entry       `json:"result" bson:"result"`
}

// TopListASN is the toplists for the ASN numbers
type TopListASN struct {
	ID   bson.ObjectId        `json:"id" bson:"_id"`
	Tag  string               `json:"tag" bson:"tag"`
	Data map[string][]ASNList `json:"data" bson:"data"`
}

// ASNList used by TopListASN
type ASNList struct {
	ASN   string `json:"asn" bson:"asn"`
	Count int    `json:"count" bson:"count"`
}

// TopListNS is the toplists for the name servers
type TopListNS struct {
	ID   bson.ObjectId       `json:"id" bson:"_id"`
	Tag  string              `json:"tag" bson:"tag"`
	Data map[string][]NSList `json:"data" bson:"data"`
}

// NSList used by TopListNS
type NSList struct {
	NS    string `json:"ns" bson:"ns"`
	Count int    `json:"count" bson:"count"`
}

// TopListTags is the toplist for the tags sorted by error levels
type TopListTags struct {
	ID   bson.ObjectId        `json:"id" bson:"_id"`
	Tag  string               `json:"tag" bson:"tag"`
	Data map[string][]TagList `json:"data" bson:"data"`
}

// TagList used by TopListTags
type TagList struct {
	Tag   string `json:"tag" bson:"tag"`
	Count int    `json:"count" bson:"count"`
}

// DomainList is used for rendering a domain list page
type DomainList struct {
	Title    string
	List     []TLD
	Notice   int
	Warning  int
	Error    int
	Critical int
}
