package client

import (
	"context"
	externalClient "github.com/adshao/go-binance/v2"
	binance "go-binance"
	"time"
)

type Client struct {
	ctx           context.Context
	binanceClient *externalClient.Client

	pingService  *externalClient.PingService
	timeService  *externalClient.ServerTimeService
	depthService *externalClient.DepthService
}

func New(ctx context.Context, apiKey, secretKey string) binance.Client {
	c := externalClient.NewClient(apiKey, secretKey)
	return &Client{ctx: ctx, binanceClient: c}
}

func (c *Client) Ping() error {
	return c.pingService.Do(c.ctx)
}

func (c *Client) Time() (time.Time, error) {
	t, err := c.timeService.Do(c.ctx)
	return time.UnixMicro(t), err
}

func (c *Client) OrderBook(obr binance.OrderBookRequest) (*binance.OrderBook, error) {
	depthResponse, err := c.depthService.Symbol(obr.Symbol).Limit(obr.Limit).Do(c.ctx)
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

func (c *Client) AggTrades(atr binance.AggTradesRequest) ([]*binance.AggTrade, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Klines(kr binance.KlinesRequest) ([]*binance.Kline, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Ticker24(tr binance.TickerRequest) (*binance.Ticker24, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TickerAllPrices() ([]*binance.PriceTicker, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TickerAllBooks() ([]*binance.BookTicker, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) NewOrder(nor binance.NewOrderRequest) (*binance.ProcessedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) NewOrderTest(nor binance.NewOrderRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) QueryOrder(qor binance.QueryOrderRequest) (*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CancelOrder(cor binance.CancelOrderRequest) (*binance.CanceledOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) OpenOrders(oor binance.OpenOrdersRequest) ([]*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) AllOrders(aor binance.AllOrdersRequest) ([]*binance.ExecutedOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Account(ar binance.AccountRequest) (*binance.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) MyTrades(mtr binance.MyTradesRequest) ([]*binance.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Withdraw(wr binance.WithdrawRequest) (*binance.WithdrawResult, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DepositHistory(hr binance.HistoryRequest) ([]*binance.Deposit, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) WithdrawHistory(hr binance.HistoryRequest) ([]*binance.Withdrawal, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) StartUserDataStream() (*binance.Stream, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) KeepAliveUserDataStream(s *binance.Stream) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) CloseUserDataStream(s *binance.Stream) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) DepthWebsocket(dwr binance.DepthWebsocketRequest) (chan *binance.DepthEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) KlineWebsocket(kwr binance.KlineWebsocketRequest) (chan *binance.KlineEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TradeWebsocket(twr binance.TradeWebsocketRequest) (chan *binance.AggTradeEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) UserDataWebsocket(udwr binance.UserDataWebsocketRequest) (chan *binance.AccountEvent, chan struct{}, error) {
	//TODO implement me
	panic("implement me")
}
