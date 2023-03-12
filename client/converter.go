package client

import (
	"fmt"
	"strconv"
	"time"

	externalClient "github.com/adshao/go-binance/v2"
	"github.com/asnowflake777/go-binance"
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
		OpenTime:                 time.UnixMicro(kline.OpenTime),
		Open:                     open,
		High:                     high,
		Low:                      low,
		Close:                    closed,
		Volume:                   volume,
		CloseTime:                time.UnixMicro(kline.CloseTime),
		QuoteAssetVolume:         quoteAssetValume,
		TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
		TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
	}
	return k, nil
}
