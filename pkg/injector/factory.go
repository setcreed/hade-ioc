package injector

import (
	"fmt"
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
			// 不同方式注入
			if field.Tag.Get("inject") == "-" {
				// 在容器里找，有就赋值
				if k := f.Get(field.Type); k != nil {
					beanValue.Field(i).Set(reflect.ValueOf(k))
				}
			} else {
				// 表达式方式处理注入
				fmt.Println("使用了表达式的方式注入")
				resultSet := expr.BeanExpr(field.Tag.Get("inject"), f.ExprMap)
				if resultSet != nil && !resultSet.IsEmpty() {
					beanValue.Field(i).Set(reflect.ValueOf(resultSet[0]))
				}
			}
		}
	}
}
