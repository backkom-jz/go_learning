package client

import (
	"go_learning/src/ch22/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	t.Log(series.Squire(2))
}
