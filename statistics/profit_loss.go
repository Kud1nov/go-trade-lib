package statistics

type ProfitLoss interface {
	PnL() float64
	TotalCommission() float64
	MaximumDrawDown() float64
	AverageProfitTrade() float64
	AverageLossTrade() float64
	MaximumPnL() float64
	MinimumPnL() float64

	UpdatePnL(float64)
	UpdateTotalCommission(float64)
	UpdateMaximumDrawDown(float64)
	UpdateAverageProfitTrade(float64)
	UpdateAverageLossTrade(float64)
	UpdateMaximumPnL(float64)
	UpdateMinimumPnL(float64)
}

type profitLoss struct {
	currentPnL         float64
	maximumPnL         float64
	minimumPnL         float64
	totalCommission    float64
	maximumDrawDown    float64
	averageProfitTrade float64
	averageLossTrade   float64
}

func (p *profitLoss) PnL() float64 {
	return p.currentPnL
}

func (p *profitLoss) MaximumPnL() float64 {
	return p.maximumPnL
}

func (p *profitLoss) MinimumPnL() float64 {
	return p.minimumPnL
}

func (p *profitLoss) TotalCommission() float64 {
	return p.totalCommission
}

func (p *profitLoss) MaximumDrawDown() float64 {
	return p.maximumDrawDown
}

func (p *profitLoss) AverageProfitTrade() float64 {
	return p.averageProfitTrade
}

func (p *profitLoss) AverageLossTrade() float64 {
	return p.averageLossTrade
}

func (p *profitLoss) UpdatePnL(v float64) {
	p.currentPnL = v
}

func (p *profitLoss) UpdateTotalCommission(v float64) {
	p.totalCommission += v
}

func (p *profitLoss) UpdateMaximumDrawDown(v float64) {
	p.maximumDrawDown = v
}

func (p *profitLoss) UpdateAverageProfitTrade(v float64) {
	p.averageProfitTrade = v
}

func (p *profitLoss) UpdateAverageLossTrade(v float64) {
	p.averageLossTrade = v
}

func (p *profitLoss) UpdateMaximumPnL(v float64) {
	if p.maximumPnL >= v {
		return
	}
	p.maximumPnL = v
}

func (p *profitLoss) UpdateMinimumPnL(v float64) {
	if p.minimumPnL <= v {
		return
	}
	p.minimumPnL = v
}

func NewProfitLoss() ProfitLoss {
	return &profitLoss{}
}
