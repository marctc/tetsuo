package backends

type BaseBackend interface {
	Report(metric, name string, value interface{}) (err error)
}
