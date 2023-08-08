package client

import (
	"context"
	"fmt"
	"time"

	extBinanceClient "github.com/adshao/go-binance/v2"
	"github.com/asnowflake777/go-binance"
)

type Client struct {
	ctx    context.Context
	client *extBinanceClient.Client
}

func New(ctx context.Context, apiKey, secretKey string) binance.Client {
	c := extBinanceClient.NewClient(apiKey, secretKey)
	client := &Client{
		ctx:    ctx,
		client: c,
	}
	return client
}

func (c *Client) Ping(ctx context.Context) error {
	return c.client.NewPingService().Do(ctx)
}

func (c *Client) Time(ctx context.Context) (time.Time, error) {
	t, err := c.client.NewServerTimeService().Do(ctx)
	return time.UnixMilli(t), err
}

func (c *Client) OrderBook(ctx context.Context, obr binance.OrderBookRequest) (*binance.OrderBook, error) {
	depthResponse, err := c.client.NewDepthService().Symbol(obr.Symbol).Limit(obr.Limit).Do(ctx)
	if err != nil {
		return nil, err
	}
	ob := &binance.OrderBook{
		LastUpdateID: depthResponse.LastUpdateID,
	}
	for _, ask := range depthResponse.Asks {
		price, quantity, err := ask.Parse()
		if err != nil {
			return nil, err
		}
		ob.Asks = append(ob.Asks, &binance.Order{Price: price, Quantity: quantity})
	}
	for _, bid := range depthResponse.Bids {
		price, quantity, err := bid.Parse()
		if err != nil {
			return nil, err
		}
		ob.Bids = append(ob.Bids, &binance.Order{Price: price, Quantity: quantity})
	}
	return ob, err
}

func (c *Client) AggTrades(context.Context, binance.AggTradesRequest) ([]*binance.AggTrade, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Klines(ctx context.Context, kr binance.KlinesRequest) ([]*binance.Kline, error) {
	klineService := c.client.NewKlinesService().
		Symbol(kr.Symbol).
		Interval(string(kr.Interval))
	if kr.Limit > 0 {
		klineService = klineService.Limit(kr.Limit)
	}
	if kr.StartTime.IsZero() {
		klineService = klineService.StartTime(kr.StartTime.Unix())
	}
	if kr.EndTime > 0 {
		klineService = klineService.StartTime(kr.EndTime)
	}
	klines, err := klineService.Do(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(klines)
	var innerKlines []*binance.Kline
	for _, kline := range klines {
		innerKline, err := ConvertKline(kline)
		if err != nil {
			return nil, fmt.Errorf("failed to convert kline: %w", err)
		}
		innerKlines = append(innerKlines, innerKline)
	}
	return innerKlines, nil
}

func (c *Client) Ticker24(context.Context, binance.TickerRequest) (*binance.Ticker24, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TickerAllPrices(context.Context) ([]*binance.PriceTicker, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TickerAllBooks(context.Context) ([]*binance.BookTicker, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) NewOrder(context.Context, binance.NewOrderRequest) (*binance.ProcessedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) NewOrderTest(context.Context, binance.NewOrderRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) QueryOrder(_ context.Context, _ binance.QueryOrderRequest) (*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CancelOrder(_ context.Context, _ binance.CancelOrderRequest) (*binance.CanceledOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) OpenOrders(_ context.Context, _ binance.OpenOrdersRequest) ([]*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) AllOrders(_ context.Context, _ binance.AllOrdersRequest) ([]*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Account(_ context.Context, _ binance.AccountRequest) (*binance.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) MyTrades(_ context.Context, _ binance.MyTradesRequest) ([]*binance.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Withdraw(_ context.Context, _ binance.WithdrawRequest) (*binance.WithdrawResult, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DepositHistory(_ context.Context, _ binance.HistoryRequest) ([]*binance.Deposit, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) WithdrawHistory(_ context.Context, _ binance.HistoryRequest) ([]*binance.Withdrawal, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) StartUserDataStream(_ context.Context) (*binance.Stream, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) KeepAliveUserDataStream(_ context.Context, _ *binance.Stream) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CloseUserDataStream(_ context.Context, _ *binance.Stream) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DepthWebsocket(_ context.Context, _ binance.DepthWebsocketRequest) (chan *binance.DepthEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) KlineWebsocket(_ context.Context, _ binance.KlineWebsocketRequest) (chan *binance.KlineEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TradeWebsocket(_ context.Context, _ binance.TradeWebsocketRequest) (chan *binance.AggTradeEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserDataWebsocket(_ context.Context, _ binance.UserDataWebsocketRequest) (chan *binance.AccountEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}
