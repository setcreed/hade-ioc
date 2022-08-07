package injector

import "reflect"

type BeanMapper map[reflect.Type]reflect.Value

func (bm BeanMapper) add(bean interface{}) {
	beanType := reflect.TypeOf(bean)
	if beanType.Kind() != reflect.Ptr {
		panic("require ptr object!")
	}
	bm[beanType] = reflect.ValueOf(bean)
}

func (bm BeanMapper) get(bean interface{}) reflect.Value {
	beanType := reflect.TypeOf(bean)
	if v, ok := bm[beanType]; ok {
		return v
	}
	return reflect.Value{}
}
