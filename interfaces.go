package binance

import (
	"context"
	"fmt"
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

// Error represents Client error structure with error code and message.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// Error returns formatted error message.
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// OrderBook represents Bids and Asks.
type OrderBook struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []*Order
	Asks         []*Order
}

type DepthEvent struct {
	WSEvent
	UpdateID int
	OrderBook
}

// Order represents single order information.
type Order struct {
	Price    float64
	Quantity float64
}

// OrderBookRequest represents OrderBook request data.
type OrderBookRequest struct {
	Symbol string
	Limit  int
}

// AggTrade represents aggregated trade.
type AggTrade struct {
	ID             int
	Price          float64
	Quantity       float64
	FirstTradeID   int
	LastTradeID    int
	Timestamp      time.Time
	BuyerMaker     bool
	BestPriceMatch bool
}

type AggTradeEvent struct {
	WSEvent
	AggTrade
}

// AggTradesRequest represents AggTrades request data.
type AggTradesRequest struct {
	Symbol    string
	FromID    int64
	StartTime int64
	EndTime   int64
	Limit     int
}

// KlinesRequest represents Klines request data.
type KlinesRequest struct {
	Symbol    string
	Interval  Interval
	Limit     int
	StartTime int64
	EndTime   int64
}

// Kline represents single Kline information.
type Kline struct {
	OpenTime                 time.Time
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                time.Time
	QuoteAssetVolume         float64
	NumberOfTrades           int
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
}

type KlineEvent struct {
	WSEvent
	Interval     Interval
	FirstTradeID int64
	LastTradeID  int64
	Final        bool
	Kline
}

// TickerRequest represents Ticker request data.
type TickerRequest struct {
	Symbol string
}

// Ticker24 represents data for 24hr ticker.
type Ticker24 struct {
	PriceChange        float64
	PriceChangePercent float64
	WeightedAvgPrice   float64
	PrevClosePrice     float64
	LastPrice          float64
	BidPrice           float64
	AskPrice           float64
	OpenPrice          float64
	HighPrice          float64
	LowPrice           float64
	Volume             float64
	OpenTime           time.Time
	CloseTime          time.Time
	FirstID            int
	LastID             int
	Count              int
}

// PriceTicker represents ticker data for price.
type PriceTicker struct {
	Symbol string
	Price  float64
}

// BookTicker represents book ticker data.
type BookTicker struct {
	Symbol   string
	BidPrice float64
	BidQty   float64
	AskPrice float64
	AskQty   float64
}

// NewOrderRequest represents NewOrder request data.
type NewOrderRequest struct {
	Symbol           string
	Side             OrderSide
	Type             OrderType
	TimeInForce      TimeInForce
	Quantity         float64
	Price            float64
	NewClientOrderID string
	StopPrice        float64
	IcebergQty       float64
	Timestamp        time.Time
}

// ProcessedOrder represents data from processed order.
type ProcessedOrder struct {
	Symbol        string
	OrderID       int64
	ClientOrderID string
	TransactTime  time.Time
}

// QueryOrderRequest represents QueryOrder request data.
type QueryOrderRequest struct {
	Symbol            string
	OrderID           int64
	OrigClientOrderID string
	RecvWindow        time.Duration
	Timestamp         time.Time
}

// ExecutedOrder represents data about executed order.
type ExecutedOrder struct {
	Symbol        string
	OrderID       int
	ClientOrderID string
	Price         float64
	OrigQty       float64
	ExecutedQty   float64
	Status        OrderStatus
	TimeInForce   TimeInForce
	Type          OrderType
	Side          OrderSide
	StopPrice     float64
	IcebergQty    float64
	Time          time.Time
}

// CancelOrderRequest represents CancelOrder request data.
type CancelOrderRequest struct {
	Symbol            string
	OrderID           int64
	OrigClientOrderID string
	NewClientOrderID  string
	RecvWindow        time.Duration
	Timestamp         time.Time
}

// CanceledOrder represents data about canceled order.
type CanceledOrder struct {
	Symbol            string
	OrigClientOrderID string
	OrderID           int64
	ClientOrderID     string
}

// OpenOrdersRequest represents OpenOrders request data.
type OpenOrdersRequest struct {
	Symbol     string
	RecvWindow time.Duration
	Timestamp  time.Time
}

// AllOrdersRequest represents AllOrders request data.
type AllOrdersRequest struct {
	Symbol     string
	OrderID    int64
	Limit      int
	RecvWindow time.Duration
	Timestamp  time.Time
}

// AccountRequest represents Account request data.
type AccountRequest struct {
	RecvWindow time.Duration
	Timestamp  time.Time
}

// Account represents user's account information.
type Account struct {
	MakerCommision  int64
	TakerCommision  int64
	BuyerCommision  int64
	SellerCommision int64
	CanTrade        bool
	CanWithdraw     bool
	CanDeposit      bool
	Balances        []*Balance
}

type AccountEvent struct {
	WSEvent
	Account
}

// Balance groups balance-related information.
type Balance struct {
	Asset  string
	Free   float64
	Locked float64
}

// MyTradesRequest represents MyTrades request data.
type MyTradesRequest struct {
	Symbol     string
	Limit      int
	FromID     int64
	RecvWindow time.Duration
	Timestamp  time.Time
}

// Trade represents data about trade.
type Trade struct {
	ID              int64
	Price           float64
	Qty             float64
	Commission      float64
	CommissionAsset string
	Time            time.Time
	IsBuyer         bool
	IsMaker         bool
	IsBestMatch     bool
}

// WithdrawRequest represents Withdraw request data.
type WithdrawRequest struct {
	Asset      string
	Address    string
	Amount     float64
	Name       string
	RecvWindow time.Duration
	Timestamp  time.Time
}

// WithdrawResult represents Withdraw result.
type WithdrawResult struct {
	Success bool
	Msg     string
}

// HistoryRequest represents history-related calls request data.
type HistoryRequest struct {
	Asset      string
	Status     *int
	StartTime  time.Time
	EndTime    time.Time
	RecvWindow time.Duration
	Timestamp  time.Time
}

// Deposit represents Deposit data.
type Deposit struct {
	InsertTime time.Time
	Amount     float64
	Asset      string
	Status     int
}

// Withdrawal represents withdrawal data.
type Withdrawal struct {
	Amount    float64
	Address   string
	TxID      string
	Asset     string
	ApplyTime time.Time
	Status    int
}

// Stream represents stream information.
//
// Read web docs to get more information about using streams.
type Stream struct {
	ListenKey string
}

type WSEvent struct {
	Type   string
	Time   time.Time
	Symbol string
}

type DepthWebsocketRequest struct {
	Symbol string
}

type KlineWebsocketRequest struct {
	Symbol   string
	Interval Interval
}

type TradeWebsocketRequest struct {
	Symbol string
}

type UserDataWebsocketRequest struct {
	ListenKey string
}
