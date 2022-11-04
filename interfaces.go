package binance

import (
	"context"
	"time"
)

// Client is wrapper for Client API.
//
// Read web documentation for more endpoints descriptions and list of
// mandatory and optional params. Wrapper is not responsible for client-side
// validation and only sends requests further.
//
// For each API-defined enum there's a special type and list of defined
// enum values to be used.
type Client interface {
	// Ping tests connectivity.
	Ping(ctx context.Context) error
	// Time returns server time.
	Time(ctx context.Context) (time.Time, error)
	// OrderBook returns list of orders.
	OrderBook(ctx context.Context, obr OrderBookRequest) (*OrderBook, error)
	// AggTrades returns compressed/aggregate list of trades.
	AggTrades(ctx context.Context, atr AggTradesRequest) ([]*AggTrade, error)
	// Klines returns klines/candlestick data.
	Klines(ctx context.Context, kr KlinesRequest) ([]*Kline, error)
	// Ticker24 returns 24hr price change statistics.
	Ticker24(ctx context.Context, tr TickerRequest) (*Ticker24, error)
	// TickerAllPrices returns ticker data for symbols.
	TickerAllPrices(ctx context.Context) ([]*PriceTicker, error)
	// TickerAllBooks returns tickers for all books.
	TickerAllBooks(ctx context.Context) ([]*BookTicker, error)

	// NewOrder places new order and returns ProcessedOrder.
	NewOrder(ctx context.Context, nor NewOrderRequest) (*ProcessedOrder, error)
	// NewOrderTest places testing order.
	NewOrderTest(ctx context.Context, nor NewOrderRequest) error
	// QueryOrder returns data about existing order.
	QueryOrder(ctx context.Context, qor QueryOrderRequest) (*ExecutedOrder, error)
	// CancelOrder cancels order.
	CancelOrder(ctx context.Context, cor CancelOrderRequest) (*CanceledOrder, error)
	// OpenOrders returns list of open orders.
	OpenOrders(ctx context.Context, oor OpenOrdersRequest) ([]*ExecutedOrder, error)
	// AllOrders returns list of all previous orders.
	AllOrders(ctx context.Context, aor AllOrdersRequest) ([]*ExecutedOrder, error)

	// Account returns account data.
	Account(ctx context.Context, ar AccountRequest) (*Account, error)
	// MyTrades list user's trades.
	MyTrades(ctx context.Context, mtr MyTradesRequest) ([]*Trade, error)
	// Withdraw executes withdrawal.
	Withdraw(ctx context.Context, wr WithdrawRequest) (*WithdrawResult, error)
	// DepositHistory lists deposit data.
	DepositHistory(ctx context.Context, hr HistoryRequest) ([]*Deposit, error)
	// WithdrawHistory lists withdraw data.
	WithdrawHistory(ctx context.Context, hr HistoryRequest) ([]*Withdrawal, error)

	// StartUserDataStream starts stream and returns Stream with ListenKey.
	StartUserDataStream(ctx context.Context) (*Stream, error)
	// KeepAliveUserDataStream prolongs stream livespan.
	KeepAliveUserDataStream(ctx context.Context, s *Stream) error
	// CloseUserDataStream closes opened stream.
	CloseUserDataStream(ctx context.Context, s *Stream) error

	DepthWebsocket(ctx context.Context, dwr DepthWebsocketRequest) (chan *DepthEvent, chan struct{}, error)
	KlineWebsocket(ctx context.Context, kwr KlineWebsocketRequest) (chan *KlineEvent, chan struct{}, error)
	TradeWebsocket(ctx context.Context, twr TradeWebsocketRequest) (chan *AggTradeEvent, chan struct{}, error)
	UserDataWebsocket(ctx context.Context, udwr UserDataWebsocketRequest) (chan *AccountEvent, chan struct{}, error)
}
