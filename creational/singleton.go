package creational

import "sync"

type Singleton struct {
}

var (
	singleInstance *Singleton
	once           sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		singleInstance = &Singleton{}
	})
	return singleInstance
}
