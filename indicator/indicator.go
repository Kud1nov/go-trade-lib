package indicator

type BaseIndicator interface {
	Add(float64) error
	IsFormed() bool
	Value() float64
	Reset() error
}
