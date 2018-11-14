package eurest

import (
    "fmt"
    "log"
//    "time"
    "net/http"
    "encoding/json"
)

type Menu struct {
//    Date string
    Categories []struct {
        Dishes []struct {
            Title string
            Price int
        }
    }
}
func DailyMenu() {
    // Make HTTP request
    response, err := http.Get("http://restaurateka.filiptesar.cz/api/menus/16")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
    var menu []Menu
    json.NewDecoder(response.Body).Decode(&menu)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

    for _, category := range menu[0].Categories {
        for _, dish := range category.Dishes {
            fmt.Println("Eurest:", dish.Title, dish.Price, "Kč")
        }
    }
}
