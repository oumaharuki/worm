package structs

import "time"

type Anime struct {
	Id           string    `xorm:"not null pk VARCHAR(40)"`
	Name         string    `xorm:"not null VARCHAR(100)"`
	Chapter      string    `xorm:"not null VARCHAR(100)"`
	Total        int       `xorm:"not null int"`
	Update       time.Time `xorm:"TIMESTAMP"`
	Index        int       `xorm:"null int"`
	Created      time.Time `xorm:"TIMESTAMP created"`
	Year         string    `xorm:"not null VARCHAR(100)`
	Area         string    `xorm:"not null VARCHAR(100)`
	Picture      string    `xorm:"not null VARCHAR(255)`
	Introduction string    `xorm:"text"`
	Form         string    `xorm:"not null VARCHAR(100)"`
	Flag         bool
}

type Chapter struct {
	Id      string    `xorm:"not null pk VARCHAR(40)"`
	Pid     string    `xorm:"not null VARCHAR(40)"`
	Name    string    `xorm:"not null VARCHAR(150)"`
	Path    string    `xorm:"text"`
	Source  string    `xorm:"not null VARCHAR(150)"`
	JX      string    `xorm:"text"`
	Created time.Time `xorm:"TIMESTAMP created"`
}

type Star struct {
	Id         string    `xorm:"not null pk VARCHAR(40)"`
	Name       string    `xorm:"not null VARCHAR(150)"`
	Pid        string    `xorm:"not null VARCHAR(40)"`
	CreateTime time.Time `orm:"TIMESTAMP created"`
}
type Director struct {
	Id         string    `xorm:"not null pk VARCHAR(40)"`
	Name       string    `xorm:"not null VARCHAR(150)"`
	Pid        string    `xorm:"not null VARCHAR(40)"`
	CreateTime time.Time `orm:"TIMESTAMP created"`
}

type UrlData struct {
	Type    string `json:"type"`
	File    string `json:"file"`
	Label   string `json:"label"`
	Default string `json:"default"`
}
type DramaPlay struct {
	Flay   string
	Encry  int
	Link   string
	Name   string
	From   string
	Trysee int
	Url    string
}
