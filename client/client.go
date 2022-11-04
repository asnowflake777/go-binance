package client

import (
	"context"
	"fmt"
	"log"
	"time"

	externalClient "github.com/adshao/go-binance/v2"
	"github.com/asnowflake777/go-binance"
)

type Client struct {
	ctx           context.Context
	logger        log.Logger
	binanceClient *externalClient.Client
}

func New(ctx context.Context, apiKey, secretKey string) binance.Client {
	return &Client{ctx: ctx, binanceClient: externalClient.NewClient(apiKey, secretKey)}
}

func (c *Client) Ping(ctx context.Context) error {
	return c.binanceClient.NewPingService().Do(ctx)
}

func (c *Client) Time(ctx context.Context) (time.Time, error) {
	t, err := c.binanceClient.NewServerTimeService().Do(ctx)
	return time.UnixMilli(t), err
}

func (c *Client) OrderBook(ctx context.Context, obr binance.OrderBookRequest) (*binance.OrderBook, error) {
	depthResponse, err := c.binanceClient.NewDepthService().Symbol(obr.Symbol).Limit(obr.Limit).Do(ctx)
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

func (c *Client) AggTrades(ctx context.Context, atr binance.AggTradesRequest) ([]*binance.AggTrade, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Klines(ctx context.Context, kr binance.KlinesRequest) ([]*binance.Kline, error) {
	klineService := c.binanceClient.NewKlinesService().
		Symbol(kr.Symbol).
		Interval(string(kr.Interval))
	if kr.Limit > 0 {
		klineService = klineService.Limit(kr.Limit)
	}
	if !kr.StartTime.IsZero() {
		klineService = klineService.StartTime(kr.StartTime.UnixMilli())
	}
	if !kr.EndTime.IsZero() {
		klineService = klineService.StartTime(kr.EndTime.UnixMilli())
	}
	klines, err := klineService.Do(ctx)
	if err != nil {
		return nil, err
	}
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

func (c *Client) Ticker24(ctx context.Context, tr binance.TickerRequest) (*binance.Ticker24, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TickerAllPrices(ctx context.Context) ([]*binance.PriceTicker, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TickerAllBooks(ctx context.Context) ([]*binance.BookTicker, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) NewOrder(ctx context.Context, nor binance.NewOrderRequest) (*binance.ProcessedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) NewOrderTest(ctx context.Context, nor binance.NewOrderRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) QueryOrder(ctx context.Context, qor binance.QueryOrderRequest) (*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CancelOrder(ctx context.Context, cor binance.CancelOrderRequest) (*binance.CanceledOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) OpenOrders(ctx context.Context, oor binance.OpenOrdersRequest) ([]*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) AllOrders(ctx context.Context, aor binance.AllOrdersRequest) ([]*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Account(ctx context.Context, ar binance.AccountRequest) (*binance.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) MyTrades(ctx context.Context, mtr binance.MyTradesRequest) ([]*binance.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Withdraw(ctx context.Context, wr binance.WithdrawRequest) (*binance.WithdrawResult, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DepositHistory(ctx context.Context, hr binance.HistoryRequest) ([]*binance.Deposit, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) WithdrawHistory(ctx context.Context, hr binance.HistoryRequest) ([]*binance.Withdrawal, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) StartUserDataStream(ctx context.Context) (*binance.Stream, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) KeepAliveUserDataStream(ctx context.Context, s *binance.Stream) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CloseUserDataStream(ctx context.Context, s *binance.Stream) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DepthWebsocket(ctx context.Context, dwr binance.DepthWebsocketRequest) (chan *binance.DepthEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) KlineWebsocket(ctx context.Context, kwr binance.KlineWebsocketRequest) (chan *binance.KlineEvent, chan struct{}, error) {
	events := make(chan *binance.KlineEvent)
	doneC, stopC, err := externalClient.WsKlineServe(kwr.Symbol, string(kwr.Interval),
		func(event *externalClient.WsKlineEvent) {
			convertedEvent, err := ConvertWSKlineEvent(event)
			if err != nil {
				c.logger.Println(err)
			} else {
				events <- convertedEvent
			}
		},
		func(err error) {
			c.logger.Println(err)
		},
	)
	go func() {
		<-doneC
		close(events)
	}()
	go func() {
		<-ctx.Done()
		stopC <- struct{}{}
	}()
	if err != nil {
		return nil, nil, err
	}
	return events, doneC, nil
}

func (c *Client) TradeWebsocket(ctx context.Context, twr binance.TradeWebsocketRequest) (chan *binance.AggTradeEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserDataWebsocket(ctx context.Context, udwr binance.UserDataWebsocketRequest) (chan *binance.AccountEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}
