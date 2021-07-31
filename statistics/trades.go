package statistics

import "sync"

type Trades interface {
	TradesCount() int
	OrdersCount() int
	LongTrades() int
	ShortTrades() int
	WinningTrades() int
	LosingTrades() int

	UpdateTrades(int)
	UpdateOrders(int)
	UpdateLongTrades(int)
	UpdateShortTrades(int)
	UpdateWinningTrades(int)
	UpdateLosingTrades(int)
}

type trades struct {
	mx            sync.RWMutex
	trades        int
	orders        int
	longTrades    int
	shortTrades   int
	winningTrades int
	losingTrades  int
}

func (t *trades) TradesCount() int {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.trades
}

func (t *trades) OrdersCount() int {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.orders
}

func (t *trades) LongTrades() int {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.longTrades
}

func (t *trades) ShortTrades() int {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.shortTrades
}

func (t *trades) WinningTrades() int {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.winningTrades
}

func (t *trades) LosingTrades() int {
	t.mx.RLock()
	defer t.mx.RUnlock()
	return t.losingTrades
}

func (t *trades) UpdateTrades(v int) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	t.trades = v
}

func (t *trades) UpdateOrders(v int) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	t.orders = v
}

func (t *trades) UpdateLongTrades(v int) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	t.longTrades = v
}

func (t *trades) UpdateShortTrades(v int) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	t.shortTrades = v
}

func (t *trades) UpdateWinningTrades(v int) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	t.winningTrades = v
}

func (t *trades) UpdateLosingTrades(v int) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	t.losingTrades = v
}

func NewTrades() Trades {
	return &trades{}
}
