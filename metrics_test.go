package tetsuo

import (
	"testing"
	"time"

<<<<<<< HEAD
	"github.com/marctc/tetsuo/backends"
=======
	"github.com/APSL/tetsuo/backends"
>>>>>>> 861b438... Initial commit
	"github.com/stretchr/testify/assert"
)

type DummyBackend struct {
	data []backends.ElasticPayload
}

func (db *DummyBackend) Report(metric, name string, value interface{}) (err error) {
	item := backends.ElasticPayload{
		Name:      name,
		TimeStamp: time.Now(),
		Value:     value,
	}
	db.data = append(db.data, item)
	return nil
}

func TestGauge(t *testing.T) {
	m := Metrics{&DummyBackend{}}
	m.Gauge("test_gauge", 5)
}

func TestIncrement(t *testing.T) {
	m := Metrics{&DummyBackend{}}
	m.Increment("test_increment")
}

func TestDecrement(t *testing.T) {
	m := Metrics{&DummyBackend{}}
	m.Decrement("test_decrement")
}

func TestTiming(t *testing.T) {
	m := Metrics{&DummyBackend{}}
	start := time.Now()
	elapsed := time.Since(start)
	m.Timing("test_timing", elapsed)
}

func TestElasticMetric(t *testing.T) {
	eb := backends.NewElasticBackend("test")
	m := Metrics{eb}
	assert.Nil(t, m.Gauge("test_elastic", 5))
}

func TestMongoMetric(t *testing.T) {
	mb := backends.NewMongoBackend("test", "test_mongo")
	m := Metrics{mb}
	assert.Nil(t, m.Gauge("test_elastic", 5))
}
