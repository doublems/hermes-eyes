package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/doublems/goquery"
)

func main() {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s", "https://store.lining.com/shop/goodsCate-sale,desc,1,16s16_126_14,16_129_14,16_128_14,16_127_15,16_128_14_l-0-0-16_126_14,16_129_14,16_128_14,16_127_15,16_128_14_l-0s0-5-0-min,max-acs45gewiov4eslkfluneh44a55tra5wvkbg.html"),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	res, err := client.Do(req)

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".selMainPic a").Each(func(i int, s *goquery.Selection) {
		href, exist := s.Attr("href")
		if !exist {
			href = "none"
		}
		fmt.Printf("%s\n", href)
	})

	document.Find(".selMainPic a img").Each(func(i int, s *goquery.Selection) {
		href, exist := s.Attr("orginalsrc")
		if !exist {
			href = "none"
		}
		fmt.Printf("%s\n", href)
	})

	document.Find(".hgoodsName").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s\n", s.Text())
	})

	document.Find(".price").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s\n", s.Text())
	})
}
