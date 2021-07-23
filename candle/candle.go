package candle

import "time"

type BaseCandle interface {
	Open() float64
	Close() float64
	High() float64
	Low() float64
	Time() time.Time
	Prev() BaseCandle
}

type baseCandle struct {
	open  float64
	close float64
	high  float64
	low   float64
	time  time.Time
	prev  BaseCandle
}

func (bc *baseCandle) Open() float64 {
	return bc.open
}

func (bc *baseCandle) Close() float64 {
	return bc.close
}

func (bc *baseCandle) High() float64 {
	return bc.high
}

func (bc *baseCandle) Low() float64 {
	return bc.low
}

func (bc *baseCandle) Time() time.Time {
	return bc.time
}

func (bc *baseCandle) Prev() BaseCandle {
	return bc.prev
}

func New(cTime time.Time, prev BaseCandle, cOpen, cClose, cHigh, cLow float64) BaseCandle {
	return &baseCandle{
		time:  cTime,
		open:  cOpen,
		close: cClose,
		high:  cHigh,
		low:   cLow,
		prev:  prev,
	}
}
