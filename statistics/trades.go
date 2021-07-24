package statistics

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
	trades        int
	orders        int
	longTrades    int
	shortTrades   int
	winningTrades int
	losingTrades  int
}

func (t *trades) TradesCount() int {
	return t.trades
}

func (t *trades) OrdersCount() int {
	return t.orders
}

func (t *trades) LongTrades() int {
	return t.longTrades
}

func (t *trades) ShortTrades() int {
	return t.shortTrades
}

func (t *trades) WinningTrades() int {
	return t.winningTrades
}

func (t *trades) LosingTrades() int {
	return t.losingTrades
}

func (t *trades) UpdateTrades(v int) {
	t.trades = v
}

func (t *trades) UpdateOrders(v int) {
	t.orders = v
}

func (t *trades) UpdateLongTrades(v int) {
	t.longTrades = v
}

func (t *trades) UpdateShortTrades(v int) {
	t.shortTrades = v
}

func (t *trades) UpdateWinningTrades(v int) {
	t.winningTrades = v
}

func (t *trades) UpdateLosingTrades(v int) {
	t.losingTrades = v
}

func NewTrades() Trades {
	return &trades{}
}
