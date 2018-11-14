package basen

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
    fmt.Println("Basen:", element.Find(".nazev").Text(), element.Find(".price").Text())
}

func DailyMenu() {
    // Make HTTP request
    response, err := http.Get("http://www.jedna-basen.cz/bistro")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    // Create a goquery document from the HTTP response
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)

    }
    document.Find(".denni-menu-polozka").Each(processElement)
}
