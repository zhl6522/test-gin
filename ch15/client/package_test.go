package client

import (
	"github.com/my/repo/test-gin/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	t.Log(series.Square(5))
}
