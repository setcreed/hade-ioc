package injector

import "reflect"

var BeanFactory *FactoryImpl

func init() {
	BeanFactory = NewFactory()
}

type FactoryImpl struct {
	beanMapper BeanMapper
}

func NewFactory() *FactoryImpl {
	return &FactoryImpl{beanMapper: make(BeanMapper)}
}

func (f *FactoryImpl) Set(beanList ...interface{}) {
	if beanList == nil || len(beanList) == 0 {
		return
	}

	for _, bean := range beanList {
		f.beanMapper.add(bean)
	}
}

func (f *FactoryImpl) Get(bean interface{}) interface{} {
	if bean == nil {
		return nil
	}
	value := f.beanMapper.get(bean)
	if value.IsValid() {
		return value.Interface()
	}
	return nil
}

// 处理依赖注入
func (f *FactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return
	}
	beanValue := reflect.ValueOf(bean)
	if beanValue.Kind() == reflect.Ptr {
		beanValue = beanValue.Elem()
	}
	if beanValue.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < beanValue.NumField(); i++ {
		field := beanValue.Type().Field(i)
		// 注入对象首字母需要大写
		if beanValue.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			if k := f.Get(field.Type); k != nil {
				beanValue.Field(i).Set(reflect.ValueOf(k))
			}
		}
	}
}
