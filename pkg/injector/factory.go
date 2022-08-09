package injector

import (
	"github.com/shenyisyn/goft-expr/src/expr"
	"reflect"
)

var BeanFactory *FactoryImpl

func init() {
	BeanFactory = NewFactory()
}

type FactoryImpl struct {
	beanMapper BeanMapper
	ExprMap    map[string]interface{}
}

func NewFactory() *FactoryImpl {
	return &FactoryImpl{beanMapper: make(BeanMapper), ExprMap: make(map[string]interface{})}
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
			// 如果是表达式注入 支持多例
			if field.Tag.Get("inject") != "-" {
				// 表达式方式处理注入
				resultSet := expr.BeanExpr(field.Tag.Get("inject"), f.ExprMap)
				if resultSet != nil && !resultSet.IsEmpty() {
					retValue := resultSet[0]
					if retValue != nil {
						f.Set(retValue)
						beanValue.Field(i).Set(reflect.ValueOf(retValue))
					}
				}
			} else {
				// 如果是 inject:"-"，直接从容器里找
				if k := f.Get(field.Type); k != nil {
					beanValue.Field(i).Set(reflect.ValueOf(k))
				}
			}
		}
	}
}

func (f *FactoryImpl) Config(cfgs ...interface{}) {
	for _, cfg := range cfgs {
		t := reflect.TypeOf(cfg)
		if t.Kind() != reflect.Ptr {
			panic("require ptr object")
		}
		f.Set(cfg)                // 把config本身也加入bean
		f.ExprMap[t.Name()] = cfg // 自动构建 exprMap
		value := reflect.ValueOf(cfg)
		for i := 0; i < t.NumMethod(); i++ {
			method := value.Method(i)
			callRet := method.Call(nil)
			if callRet != nil && len(callRet) == 1 {
				f.Set(callRet[0].Interface())
			}
		}
	}
}
