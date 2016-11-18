package snapcard

import (
	"testing"
)

func TestRate(t *testing.T) {
	_, err := Rate("BTCUSD")
	if err != nil {
		t.Fail()
	}
}

func TestInvalidRatePair(t *testing.T) {
	_, err := Rate("foobar")
	if err == nil {
		t.Fail()
	}
}
