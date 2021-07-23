package candles

import "time"

type Candle interface {
	Open() float64
	Close() float64
	High() float64
	Low() float64
	Time() time.Time
	Prev() Candle
}

type baseCandle struct {
	open  float64
	close float64
	high  float64
	low   float64
	time  time.Time
	prev  Candle
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

func (bc *baseCandle) Prev() Candle {
	return bc.prev
}

func New (
	time time.Time,
	open float64,
	close float64,
	high float64,
	low float64,
	prev Candle,
	) Candle {
	return &baseCandle{
		time: time,
		open: open,
		close: close,
		high: high,
		low: low,
		prev: prev,
	}
}