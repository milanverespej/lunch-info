package main

import (
	"github.com/milanverespej/lunch-info/basen"
	"github.com/milanverespej/lunch-info/eurest"
	"sync"
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
