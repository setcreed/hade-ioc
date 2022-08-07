package injector

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
