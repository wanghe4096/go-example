package service

import "gitlab.luojilab.com/zeroteam/artemis/engine"

type IService interface {
	Get(pk interface{}) interface{}
	MultiGet(pks []interface{}) map[string]interface{}
	List(pks []interface{}) []interface{}
	Add(data interface{}) (n int64, err error)
	Set(pk interface{}, data interface{}) (b bool, err error)
	MultiSet(map[string]interface{}) (n int64, err error)
}

type IController interface {
	IServicecdw
}

type IRouter interface {
	engine.IRouter
}
