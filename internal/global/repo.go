package global

type Repo interface {
	Setup(*Global)
	Get() interface{}
	Close() error
}
