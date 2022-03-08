package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"hasrateNoCrawlerScrapper/config"
	"hasrateNoCrawlerScrapper/data"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func DualType(e *colly.HTMLElement) {
	var dual = regexp.MustCompile(`([a-zA-Z]{0,}?\+[a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/s)([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/s)([0-9]{0,})w([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/w)([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/w)\$([0-9]{0,}?\.[0-9]{0,})\$([0-9]{0,}?\.[0-9]{0,}) ([0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}) ([a-zA-Z]{0,})`)
	if dual.MatchString(e.Text) {
		DualCrawled := data.DualCoin{}
		coin := dual.ReplaceAllString(e.Text, `$1`)
		coinsName := strings.Split(coin, "+")
		DualCrawled.CoinPrimary = coinsName[0]
		DualCrawled.CoinAlt = coinsName[1]
		DualCrawled.CoinName = coin
		DualCrawled.Conso, _ = strconv.Atoi(dual.ReplaceAllString(e.Text, `$6`))
		DualCrawled.Income, _ = strconv.ParseFloat(dual.ReplaceAllString(e.Text, `$13`), 64)
		DualCrawled.HashrateAlt, _ = strconv.ParseFloat(dual.ReplaceAllString(e.Text, `$4`), 64)
		DualCrawled.HashratePrimary, _ = strconv.ParseFloat(dual.ReplaceAllString(e.Text, `$2`), 64)
		DualCrawled.UnitPrimary = dual.ReplaceAllString(e.Text, `$3`)
		DualCrawled.UnitAlt = dual.ReplaceAllString(e.Text, `$5`)
		for index, elem := range data.CardsResult.CardHarvested {
			if elem.Name == e.Request.URL.Path[1:] {
				data.CardsResult.CardHarvested[index].DualCoin = append(data.CardsResult.CardHarvested[index].DualCoin, DualCrawled)
			}
		}
	}
}

func Singletype(e *colly.HTMLElement) {
	var single = regexp.MustCompile(`([a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/s)([0-9]{0,})w([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/w)\$([0-9]{0,}?\.[0-9]{0,})\$([0-9]{0,}?\.[0-9]{0,}) ([0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}) ([a-zA-Z]{0,})`)
	if single.MatchString(e.Text) {
		SingleCrawled := data.SingleCoin{}
		coin := single.ReplaceAllString(e.Text, `$1`)
		SingleCrawled.CoinPrimary = coin
		SingleCrawled.CoinName = coin
		SingleCrawled.HashratePrimary, _ = strconv.ParseFloat(single.ReplaceAllString(e.Text, `$2`), 64)
		SingleCrawled.UnitPrimary = single.ReplaceAllString(e.Text, `$3`)
		SingleCrawled.Conso, _ = strconv.Atoi(single.ReplaceAllString(e.Text, `$4`))
		SingleCrawled.Income, _ = strconv.ParseFloat(single.ReplaceAllString(e.Text, `$9`), 64)
		for index, elem := range data.CardsResult.CardHarvested {
			if elem.Name == e.Request.URL.Path[1:] {
				data.CardsResult.CardHarvested[index].SingleCoin = append(data.CardsResult.CardHarvested[index].SingleCoin, SingleCrawled)
			}
		}
	}
}

func DispatchType(e *colly.HTMLElement) {
	DualType(e)
	Singletype(e)
}

func runCrawler() {
	data.CardsResult = &data.CardsHarvested{}
	c := colly.NewCollector(
		colly.AllowedDomains("www.hashrate.no", "hashrate.no"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème : ", err)
	})
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		DispatchType(e)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("harvesting : ", r.URL.String())
	})
	c.OnScraped(func(s *colly.Response) {
		log.Println("harvested : ", s.Request.URL.String())
	})
	for _, elem := range config.Cards.CardsList {
		url := fmt.Sprintf("https://www.hashrate.no/%s", elem)
		TmpCardHarvested := data.CardHarvested{}
		TmpCardHarvested.Name = elem
		data.CardsResult.CardHarvested = append(data.CardsResult.CardHarvested, TmpCardHarvested)
		c.Visit(url)
	}
}

func main() {
	config.LoadYamlConfig()
	runCrawler()
	JsonCardResult, _ := json.Marshal(data.CardsResult)
	log.Println(string(JsonCardResult))
}
