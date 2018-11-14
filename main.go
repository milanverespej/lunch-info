package main

import (
    "sync"
    "github.com/milanverespej/lunch_info/basen"
    "github.com/milanverespej/lunch_info/eurest"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(2)

    go func() {
        basen.DailyMenu()
        defer wg.Done()
    }()
     go func() {
        eurest.DailyMenu()
        defer wg.Done()
    }()

    wg.Wait()
}
