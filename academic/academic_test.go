package academic

import (
	"fmt"
	"testing"
)

var key string = "-" // this code is sample , forget this :) my token has been changed !

func TestCalcHistogram(t *testing.T) {
	hist, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := hist.CalcHistogram("And(Composite(AA.AuN=='jaime teevan'),Y>2012)",
		CalcHistogramRequest{
			Model: "latest",
		},
	)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}
