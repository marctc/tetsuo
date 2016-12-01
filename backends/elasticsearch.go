package backends

import (
	"time"

	elastic "gopkg.in/olivere/elastic.v3"
)

type ElasticBackend struct {
	Client    *elastic.Client
	IndexName string
}

type ElasticPayload struct {
	Name      string      `json:"name"`
	TimeStamp time.Time   `json:"@timestamp"`
	Value     interface{} `json:"value"`
}

func NewElasticBackend(indexName string) (eb *ElasticBackend) {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	return &ElasticBackend{Client: client, IndexName: indexName}
}

func (eb *ElasticBackend) GetIndexName() string {
	return eb.IndexName + "-" + time.Now().Format("2006.01.02")
}

func (eb *ElasticBackend) Report(metric, name string, value interface{}) (err error) {
	ep := ElasticPayload{
		Name:      name,
		TimeStamp: time.Now(),
		Value:     value,
	}
	_, err = eb.Client.Index().Index(eb.GetIndexName()).Type(metric).BodyJson(ep).Refresh(true).Do()
	return err
}
