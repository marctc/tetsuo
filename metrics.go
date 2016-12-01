package tetsuo

import (
	"time"
	"github.com/marctc/tetsuo/backends"
)


// Metrics reporting struct
type Metrics struct {
	Backend backends.BaseBackend // Backend stores the reported data.
}

// Gauge records the value of a metric
//    m.Gauge('users.notifications', 13)
func (m *Metrics) Gauge(name string, value interface{}) (err error) {
	return m.Backend.Report("gauge", name, value)
}

// Increment a counter
//    m.Increment('user.profile.views')
func (m *Metrics) Increment(name string) (err error) {
	return m.Backend.Report("counter", name, 1)
}

// Decrement a counter
//    m.Decrement('hotel.occupation')
func (m *Metrics) Decrement(name string) (err error) {
	return m.Backend.Report("counter", name, -1)
}

// Timing records a time value
//	  func foo() {
//		  start := time.Now()
//		  // do stuff...
//		  elapsed := time.Since(start)
//		  m.Timing("foo_operation", elapsed)
//	  }
func (m *Metrics) Timing(name string, value time.Duration) (err error) {
	return m.Backend.Report("timing", name, value)
}
