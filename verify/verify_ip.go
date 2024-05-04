package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var DATA_FILE string

func init() {
	flag.StringVar(&DATA_FILE, "f", "Country.mmdb", "specify MaxMind database file")
	flag.Parse()
}

func main() {
	fmt.Println("Open MaxMind database file:", DATA_FILE)
	db, err := geoip2.Open(DATA_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil

	var list = []string{"103.200.30.143", "103.228.130.61", "216.58.200.238",
		"103.200.30.245", "118.184.26.113", "103.200.31.172", "69.171.235.101", "123.126.55.41", "117.23.61.238", "27.152.182.60",
		"2402:f000:1:404:166:111:4:100", "2001:4860:4860::8844"}
	// "103.200.30.143", "103.228.130.61", "216.58.200.238",
	//	"103.200.30.245", "118.184.26.113", "103.200.31.172", "69.171.235.101"
	// is custom ip from gfw and be clarified as HK in ipip but china in maxmind
	// "123.126.55.41", "117.23.61.238" is example ip of china
	// "2402:f000:1:404:166:111:4:100": www.tsinghua.edu.cn
	// "2001:4860:4860::8844": dns.google

	for _, ipTxt := range list {
		ip := net.ParseIP(ipTxt)
		record, err := db.Country(ip)
		if err != nil || record == nil {
			log.Fatal(err)
		}

		fmt.Printf("IP:%s-Locale:%s\n", ipTxt, record.Country.IsoCode)
		//fmt.Printf("%d, %s, %s|\n %s, %d, %s, %v|\n %s, %d, %v, %s|\n %s, %s, %v, %v|\n ", record.Continent.GeoNameID, record.Continent.Code, record.Continent.Names,
		//	record.Country.Names, record.Country.GeoNameID, record.Country.IsoCode, record.Country.IsInEuropeanUnion,
		//	record.RegisteredCountry.Names, record.RegisteredCountry.GeoNameID, record.RegisteredCountry.IsInEuropeanUnion, record.RegisteredCountry.IsoCode,
		//	record.RepresentedCountry.IsoCode, record.RepresentedCountry.Type, record.Traits.IsAnonymousProxy, record.Traits.IsSatelliteProvider)
	}
}
