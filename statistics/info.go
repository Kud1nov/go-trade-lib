package statistics

type Info interface {
	Strategy() Strategy
	ProfitLoss() ProfitLoss
	Trades() Trades
}

type statistic struct {
	strategy   Strategy
	profitLoss ProfitLoss
	trades     Trades
}

func (s *statistic) Strategy() Strategy {
	return s.strategy
}
func (s *statistic) ProfitLoss() ProfitLoss {
	return s.profitLoss
}
func (s *statistic) Trades() Trades {
	return s.trades
}

func New() Info {
	return &statistic{
		strategy:   NewStrategy(),
		profitLoss: NewProfitLoss(),
		trades:     NewTrades(),
	}
}
