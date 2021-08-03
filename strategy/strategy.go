package strategy

type TransferType int

const (
	Status TransferType = iota
	OrderBook
	LastPrice
	TradeStat
	candles
)

type TransferMessage struct {
	Event TransferType
	Data  interface{}
}

type Strategy interface {
	DataReceiver(TransferMessage) error
}