package main

import (
	"fmt"
)

type Item struct {
	metric string
	name   string
	value  interface{}
}

type Backend interface {
	report(metric string, name string, value interface{})
}

type ElasticBackend struct {
	data []Item
}

func (backend *ElasticBackend) report(metric string, name string, value interface{}) {
	item := Item{metric, name, value}
	backend.data = append(backend.data, item)
	fmt.Println(backend.data)
}

type MongoBackend struct {
	data []Item
}

func (backend *MongoBackend) report(metric string, name string, value interface{}) {
	item := Item{metric, name, value}
	backend.data = append(backend.data, item)
	fmt.Println(backend.data)
}

type Metrics struct {
	backend Backend
}

func (metrics *Metrics) Gauge(name string, value interface{}) {
	metrics.backend.report("gauge", name, value)
}

func main() {
	m := &Metrics{&ElasticBackend{}}
	m.Gauge("test.int", 5)
	m.Gauge("test.float", []int{1, 2})
	fmt.Println(m.backend)
}
