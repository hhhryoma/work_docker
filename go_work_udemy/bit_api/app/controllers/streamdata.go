package controllers

import (
	"../../app/models"
	"../../bitFlyer"
	"../../config"
)

func StreamIngestionData() {
	var tickerChanel = make(chan bitFlyer.Ticker)
	apiClient := bitFlyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChanel)
	go func() {
		for ticker := range tickerChanel {
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated == true && duration == config.Config.TradeDuration {
					// TODO
				}
			}
		}
	}()

}
