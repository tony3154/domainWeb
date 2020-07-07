package models

type Domain struct {
	ID                int    `orm:"column(id)"`
	DomainName        string `orm:"column(domainname)"`
	Project           string
	Service           string
	CDN               string `orm:"column(cdn)"`
	HTTPS             string `orm:"column(https)"`
	Backend           string
	WhiteList         string `orm:"column(whitelist)"`
	WhiteListLocation string `orm:"column(whitelistlocation)"`
	Notes             string
}
