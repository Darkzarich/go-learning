package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

var tickers = []Ticker{
	{Name: "GAZP", AvgPrice: 165.0},
	{Name: "X5", AvgPrice: 3200.0},
	{Name: "USD/RUB", AvgPrice: 81.5},
	{Name: "YNDX", AvgPrice: 4500.0},
	{Name: "BRENT", AvgPrice: 68.0},
	{Name: "DELL", AvgPrice: 125.0},
	{Name: "INTC", AvgPrice: 35.0},
	{Name: "RTKM", AvgPrice: 65.0},
}

func generatePrice(avgPrice float64) float64 {
	change := rand.Float64()*0.04 - 0.02

	return avgPrice * (1 + change)
}

func generateTicker(ticker Ticker, timestamp time.Time) TickerEmulatorDTO {
	return TickerEmulatorDTO{
		Ticker:    ticker.Name,
		Price:     generatePrice(ticker.AvgPrice),
		Timestamp: timestamp,
	}
}

func randomTicker() Ticker {
	return tickers[rand.Intn(len(tickers))]
}

func Run(producer Producer) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for timestamp := range ticker.C {
		for _, ticker := range tickers {
			event := generateTicker(ticker, timestamp)

			data, err := json.Marshal(event)
			if err != nil {
				log.Println("marshal ticker:", err)
				continue
			}

			if err := producer.PushMessage(context.Background(), data); err != nil {
				log.Println("push ticker:", err)
			}
		}
	}
}
