package gocurrency

import (
	"github.com/beevik/etree"
	"strconv"
	"net/http"
	"log"
	"time"
)

const (
	RatesSource = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	CacheExpire = time.Hour * 4
)

var (
	ExchangeRates = make(map[string]float64)
	CacheTime time.Time
)

func init() {
	ExchangeRates["EUR"] = 1 //EUR doesn't exist by default, because values are relative to EUR
}

func RefreshRates() {
	resp, err := http.Get(RatesSource)

	defer func() {
		if r := recover(); r != nil {
			updateError(r.(error))
		}
	}()

	if err != nil {
		panic(err)
	}

	root := etree.NewDocument()

	if _, err := root.ReadFrom(resp.Body); err != nil {
		panic(err)
	}

	for _, v := range selectRecursive(root.Element, "gesmes:Envelope", "Cube", "Cube", "Cube"){
		rate, _ := strconv.ParseFloat(v.SelectAttr("rate").Value, 32)
		ExchangeRates[v.SelectAttr("currency").Value] = rate
	}

	CacheTime = time.Now()
}

func RefreshIfRequired() {
	if CacheTime.Add(CacheExpire).Before(time.Now()) {
		RefreshRates()
	}
}

func updateError(err error) {
	if CacheTime.IsZero() {
		log.Panicf("Unable to update exchange rates. Error: \n%v", err)
	}

	log.Printf("<IMPORTANT> Unable to update exchange rates. Current rates are from %v", CacheTime)
}

func selectRecursive(elem etree.Element, path ...string) ([]*etree.Element) {
	relem := &elem
	for _, v := range path[:len(path) - 1] {
		relem = relem.SelectElement(v)
	}
	return relem.SelectElements(path[len(path) - 1])
}