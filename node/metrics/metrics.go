package metric

import (
	metrics "github.com/rcrowley/go-metrics"
)

var (
	SayHelloInMeter   = metrics.NewRegisteredMeter("airmsExample/SayHello/in", nil)
	SayHelloOutMeter  = metrics.NewRegisteredMeter("airmsExample/SayHello/out", nil)
	SayHelloTimeMeter = metrics.NewRegisteredMeter("airmsExample/SayHello/time", nil)
)
