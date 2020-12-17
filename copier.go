// Package copier 相同结构进行深拷贝
package copier

import (
	"errors"
	"fmt"
	"reflect"
)

// Copy 做入参基本校验以及数据类型萃取
func Copy(a interface{}, b interface{}) (err error) {
	// reflect 相关操作极易引发 panic
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("copy err :%v", e)
			println(err.Error())
		}
	}()

	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if av.IsNil() || bv.IsNil() {
		err = errors.New("params is nil")
		return
	}

	err = copyFields(av, bv.Elem(), reflect.StructField{})
	return
}

// copyFields 递归进行同命名属性的深度填充  bv -> av
func copyFields(av reflect.Value, bv reflect.Value, field reflect.StructField) error {
	if !bv.IsValid() {
		return nil
	}

	ak := av.Kind()
	at := av.Type()
	bt := bv.Type()
	if ak == reflect.Ptr {
		// 指针类型需要判空，并进行 New
		vPtr := av
		isNew := false
		if av.IsNil() {
			isNew = true
			vPtr = reflect.New(at.Elem())
		}

		err := copyFields(vPtr.Elem(), bv, field)
		if err != nil {
			return err
		}

		if isNew {
			av.Set(vPtr)
		}
	} else if ak == reflect.Array || ak == reflect.Slice {
		// 切片、数组类型需要深拷贝
		// 切片需要单独 make
		if ak == reflect.Slice {
			if bv.IsNil() {
				return nil
			}

			slice := reflect.MakeSlice(field.Type, bv.Len(), bv.Cap())
			av.Set(slice)
		}

		for i := 0; i < av.Len(); i++ {
			// 防治越界
			if i >= bv.Len() {
				continue
			}

			bvi := bv.Index(i)
			if bvi.Kind() == reflect.Ptr {
				bvi = bvi.Elem()
			}

			copyFields(av.Index(i), bvi, field)
		}
	} else if ak == reflect.Map {
		if bv.IsNil() {
			return nil
		}

		// map 类型需要深拷贝
		mmap := reflect.MakeMap(field.Type)
		av.Set(mmap)

		for _, key := range bv.MapKeys() {
			// map 的 key 不接受对象，value 不接受 map 类型(map[string]map[string]string)
			if key.Kind() >= reflect.Array && key.Kind() != reflect.String {
				continue
			}

			bvi := bv.MapIndex(key)
			var avi reflect.Value
			if bvi.Kind() == reflect.Ptr {
				// 指针类型需要进行判空，防止 new 出新对象
				if bvi.IsNil() {
					avi = reflect.NewAt(field.Type.Elem().Elem(), nil)
					av.SetMapIndex(key, avi)
					continue
				}

				bvi = bvi.Elem()
				avi = reflect.New(field.Type.Elem().Elem())
				copyFields(avi, bvi, field)
				av.SetMapIndex(key, avi)
			} else if bvi.Kind() == reflect.Struct {
				avi = reflect.New(field.Type.Elem())
				copyFields(avi, bvi, field)
				av.SetMapIndex(key, avi)
			} else {
				avi = reflect.New(field.Type.Elem()).Elem()
				copyFields(avi, bvi, field)
				av.SetMapIndex(key, avi)
			}
		}
	} else if ak == reflect.Struct {
		// 类型相等且非匿名，可以直接使用
		if at == bt && !field.Anonymous {
			av.Set(bv)
			return nil
		}

		// struct 类型需要深拷贝
		for i := 0; i < av.NumField(); i++ {
			af := at.Field(i)
			b := bv.FieldByName(af.Name)
			if b.Kind() == reflect.Ptr {
				b = b.Elem()
			}

			err := copyFields(av.Field(i), b, af)
			if err != nil {
				return err
			}
		}
	} else {
		if ak == bv.Kind() {
			if at == bv.Type() {
				av.Set(bv)
			} else {
				// 此处处理自定义类型
				switch bk := bv.Kind(); bk {
				case reflect.String:
					av.SetString(bv.String())
				case reflect.Bool:
					av.SetBool(bv.Bool())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					av.SetInt(bv.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
					av.SetUint(bv.Uint())
				case reflect.Float32, reflect.Float64:
					av.SetFloat(bv.Float())
				case reflect.Complex64, reflect.Complex128:
					av.SetComplex(bv.Complex())
				default:
					return errors.New("unsupport kind")
				}
			}
		} else {
			return errors.New("different kind")
		}
	}

	return nil
}
