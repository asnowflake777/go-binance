package client

import (
	"fmt"
	externalClient "github.com/adshao/go-binance/v2"
	"github.com/asnowflake777/go-binance"
	"strconv"
	"time"
)

func ConvertKline(kline *externalClient.Kline) (*binance.Kline, error) {
	open, err := strconv.ParseFloat(kline.Open, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	high, err := strconv.ParseFloat(kline.High, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	low, err := strconv.ParseFloat(kline.Low, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	closed, err := strconv.ParseFloat(kline.Close, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	volume, err := strconv.ParseFloat(kline.Volume, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	quoteAssetValume, err := strconv.ParseFloat(kline.QuoteAssetVolume, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	takerBuyBaseAssetVolume, err := strconv.ParseFloat(kline.TakerBuyBaseAssetVolume, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	takerBuyQuoteAssetVolume, err := strconv.ParseFloat(kline.TakerBuyQuoteAssetVolume, 64)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	k := &binance.Kline{
		OpenTime:                 time.UnixMilli(kline.OpenTime),
		Open:                     open,
		High:                     high,
		Low:                      low,
		Close:                    closed,
		Volume:                   volume,
		CloseTime:                time.UnixMilli(kline.CloseTime),
		QuoteAssetVolume:         quoteAssetValume,
		TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
		TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
	}
	return k, nil
}

func ConvertWSKlineEvent(event *externalClient.WsKlineEvent) (*binance.KlineEvent, error) {
	kline, err := ConvertWSKline(&event.Kline)
	if err != nil {
		return nil, err
	}
	klineEvent := &binance.KlineEvent{
		WSEvent: binance.WSEvent{
			Type:   event.Event,
			Time:   time.UnixMilli(event.Time),
			Symbol: event.Symbol,
		},
		Interval:     binance.Interval(event.Kline.Interval),
		FirstTradeID: event.Kline.FirstTradeID,
		LastTradeID:  event.Kline.LastTradeID,
		Final:        event.Kline.IsFinal,
		Kline:        *kline,
	}
	return klineEvent, nil
}

func ConvertWSKline(wsKline *externalClient.WsKline) (*binance.Kline, error) {
	kline := &externalClient.Kline{
		OpenTime:                 wsKline.StartTime,
		Open:                     wsKline.Open,
		High:                     wsKline.High,
		Low:                      wsKline.Low,
		Close:                    wsKline.Close,
		Volume:                   wsKline.Volume,
		CloseTime:                wsKline.EndTime,
		QuoteAssetVolume:         wsKline.QuoteVolume,
		TradeNum:                 wsKline.TradeNum,
		TakerBuyBaseAssetVolume:  wsKline.ActiveBuyVolume,
		TakerBuyQuoteAssetVolume: wsKline.ActiveBuyQuoteVolume,
	}
	convertedKline, err := ConvertKline(kline)
	return convertedKline, err
}
