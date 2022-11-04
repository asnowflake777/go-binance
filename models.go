package binance

import "time"

// Order represents single order information.
type Order struct {
	Price    float64
	Quantity float64
}

// OrderBook represents Bids and Asks.
type OrderBook struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []*Order
	Asks         []*Order
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

// ProcessedOrder represents data from processed order.
type ProcessedOrder struct {
	Symbol        string
	OrderID       int64
	ClientOrderID string
	TransactTime  time.Time
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

// CanceledOrder represents data about canceled order.
type CanceledOrder struct {
	Symbol            string
	OrigClientOrderID string
	OrderID           int64
	ClientOrderID     string
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

// Balance groups balance-related information.
type Balance struct {
	Asset  string
	Free   float64
	Locked float64
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
