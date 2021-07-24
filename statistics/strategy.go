package statistics

type Strategy interface {
	ExpectedValue() float64
	AbsExpectedValue() float64
	StandardDeviation() float64
	SharpeRatio() float64
	SortinoRatio() float64
	ProfitFactor() float64
	RecoveryFactor() float64
	ZScore() float64

	UpdateExpectedValue(float64)
	UpdateAbsExpectedValue(float64)
	UpdateStandardDeviation(float64)
	UpdateSharpeRatio(float64)
	UpdateSortinoRatio(float64)
	UpdateProfitFactor(float64)
	UpdateRecoveryFactor(float64)
	UpdateZScore(float64)
}
type strategy struct {
	expectedValue     float64
	absExpectedValue  float64
	standardDeviation float64
	sharpeRatio       float64
	sortinoRatio      float64
	profitFactor      float64
	recoveryFactor    float64
	zScore            float64
}

func (s *strategy) ExpectedValue() float64 {
	return s.expectedValue
}

func (s *strategy) AbsExpectedValue() float64 {
	return s.absExpectedValue
}

func (s *strategy) StandardDeviation() float64 {
	return s.standardDeviation
}

func (s *strategy) SharpeRatio() float64 {
	return s.sharpeRatio
}

func (s *strategy) SortinoRatio() float64 {
	return s.sortinoRatio
}

func (s *strategy) ProfitFactor() float64 {
	return s.profitFactor
}

func (s *strategy) RecoveryFactor() float64 {
	return s.recoveryFactor
}

func (s *strategy) ZScore() float64 {
	return s.zScore
}

func (s *strategy) UpdateExpectedValue(v float64) {
	s.expectedValue = v
}
func (s *strategy) UpdateAbsExpectedValue(v float64) {
	s.absExpectedValue = v
}
func (s *strategy) UpdateStandardDeviation(v float64) {
	s.standardDeviation = v
}
func (s *strategy) UpdateSharpeRatio(v float64) {
	s.sharpeRatio = v
}
func (s *strategy) UpdateSortinoRatio(v float64) {
	s.sortinoRatio = v
}
func (s *strategy) UpdateProfitFactor(v float64) {
	s.profitFactor = v
}
func (s *strategy) UpdateRecoveryFactor(v float64) {
	s.recoveryFactor = v
}
func (s *strategy) UpdateZScore(v float64) {
	s.zScore = v
}

func NewStrategy() Strategy {
	return &strategy{}
}
