package metrics

import (
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
)

// CountGoroutines - count the numlber of goroutines

// here to create my custom metrics i used the  prometheus builtin function "prometheus.NewGaugeFunc()"
// So that we can pass an anonymous function
var Goroutine = prometheus.NewGaugeFunc(
	prometheus.GaugeOpts{
		Subsystem: "runtime",
		Name:      "number_goroutines",
		Help:      "the number of goroutines",
	},
	func() float64 { return float64(runtime.NumGoroutine()) },
)

var Countcpu = prometheus.NewGaugeFunc(
	prometheus.GaugeOpts{
		Subsystem: "runtime",
		Name:      "number_process_used",
		Help:      "the number of process used",
	},
	func() float64 { return float64(runtime.NumCPU()) },
)
