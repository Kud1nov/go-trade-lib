package order

type Type int
type State int
type Direction string

const (
	LimitOrder     Type      = 0
	StopOrder      Type      = 1
	StateFilled    State     = 3
	StateCancelled State     = 4
	StateFailed    State     = 5
	Buy            Direction = "buy"
	Sell           Direction = "sell"
)

type BaseOrder interface {
	Ticker() string
	Price() float64
	Quantity() int
	State() State
	Direction() Direction
	Type() Type

	OnChanged(fn func())
}
