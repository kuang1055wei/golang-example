package main

import "sync"

//单列模式，运用sync.once保证线程安全且只执行一次
type singleton struct {

}

var instance *singleton
var once sync.Once

func GetInstance() *singleton  {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
