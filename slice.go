package generic

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type slice struct {
	slicePtr interface{}
}

// New a slice with slice ptr
func Slice(slicePtr interface{}) *slice {
	return &slice{slicePtr}
}

// Remove element at index of slice
func (s *slice) RemoveAt(index int) error {
	err := s.checkSlice()
	if err != nil {
		return err
	}

	slicePtrValue := reflect.ValueOf(s.slicePtr)
	sliceValue := slicePtrValue.Elem()
	if index < 0 || index >= sliceValue.Len() {
		return errors.New("index out of range!")
	}
	sliceValue.Set(reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, sliceValue.Len())))
	return nil
}

func (s *slice) checkSlice() error {
	if s.slicePtr == nil {
		return errors.New("slice is nil!")
	}

	slicePtrValue := reflect.ValueOf(s.slicePtr)
	// should be pointer
	if slicePtrValue.Type().Kind() != reflect.Ptr {
		return errors.New("should be slice pointer!")
	}

	sliceValue := slicePtrValue.Elem()
	// should be slice
	if sliceValue.Type().Kind() != reflect.Slice {
		return errors.New("should be slice pointer!")
	}

	return nil
}

// sort slice by quick sort algorithm
// support slice of all int and uint types
// and support stuct which has the compare function, function name should be "Compare", and return int. such as
// type student stuct {
// 	age int
// }
//
// // return value:
// // if value == 0, equal with other student,
// // if value < 0, less than other,
// // if value > 0, greater than other student.
// func (s student) Compare(other student) int {
// 	return s.age - other.age
// }
//
func (s *slice) QuickSort() error {
	err := s.checkSlice()
	if err != nil {
		return err
	}

	slicePtrValue := reflect.ValueOf(s.slicePtr)
	sliceValue := slicePtrValue.Elem()
	if sliceValue.Len() <= 1 {
		return nil
	}
	compareFuncName := "Compare"
	err = checkTypeOfSort(sliceValue.Index(0), compareFuncName)
	if err != nil {
		return err
	}
	quickSort(sliceValue, 0, sliceValue.Len()-1, compareFuncName)
	return nil
}

// if the type is not supported, or the struct doesn't have compare function, then return error.
func checkTypeOfSort(elem reflect.Value, funcName string) error {
	switch elem.Type().Kind() {
	case reflect.Struct:
		funcValue := elem.MethodByName(funcName)
		if funcValue.IsNil() {
			return errors.New("no compare function!")
		}
		if !strings.HasSuffix(funcValue.Type().String(), "int") {
			return errors.New("compare function should return int!")
		}
		break
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		break
	default:
		fmt.Println(elem.Type().Kind())
		return errors.New("unsupport type: " + elem.Type().Kind().String())
	}

	return nil
}

// Basicly it is same as QuickSort function.
// It just give you choice to decide the compare function which is used by struct
func (s *slice) QuickSortBy(compareFuncName string) error {
	err := s.checkSlice()
	if err != nil {
		return err
	}

	slicePtrValue := reflect.ValueOf(s.slicePtr)
	sliceValue := slicePtrValue.Elem()
	if sliceValue.Len() <= 1 {
		return nil
	}
	if err = checkTypeOfSort(sliceValue.Index(0), compareFuncName); err != nil {
		return err
	}

	quickSort(sliceValue, 0, sliceValue.Len()-1, compareFuncName)
	return nil
}

// the internal function for implementing quick sort algorithm.
func quickSort(slice reflect.Value, lowIndex, highIndex int, compareFuncName string) {
	if lowIndex < 0 {
		lowIndex = 0
	}
	if highIndex > slice.Len()-1 {
		highIndex = slice.Len() - 1
	}
	if lowIndex >= highIndex {
		return
	}

	firstIndex := lowIndex
	lastIndex := highIndex
	key := clone(slice.Index(firstIndex))

	for firstIndex < lastIndex {

		for firstIndex < lastIndex && compare(slice.Index(lastIndex), key, compareFuncName) >= 0 {
			lastIndex = lastIndex - 1
		}
		swap(slice, firstIndex, lastIndex)

		for firstIndex < lastIndex && compare(slice.Index(firstIndex), key, compareFuncName) <= 0 {
			firstIndex = firstIndex + 1
		}
		swap(slice, firstIndex, lastIndex)
	}

	quickSort(slice, lowIndex, firstIndex-1, compareFuncName)
	quickSort(slice, lastIndex+1, highIndex, compareFuncName)
}

func clone(value reflect.Value) reflect.Value {
	newValue := reflect.New(value.Type()).Elem()
	newValue.Set(value)
	return newValue
}

// compare two elements of the slice
// return value:
// if value == 0, va1 is equal to val2,
// if value < 0, va1 is less than val2,
// if value > 0, va1 is greater than val2.
func compare(val1, val2 reflect.Value, compareFuncName string) int {

	switch val1.Type().Kind() {
	case reflect.Int8:
		return int(val1.Interface().(int8)) - int(val2.Interface().(int8))
		break
	case reflect.Int16:
		return int(val1.Interface().(int16)) - int(val2.Interface().(int16))
		break
	case reflect.Int32:
		return int(val1.Interface().(int32)) - int(val2.Interface().(int32))
		break
	case reflect.Int64:
		v1 := val1.Interface().(int64)
		v2 := val2.Interface().(int64)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
		break

	case reflect.Int:
		return val1.Interface().(int) - val2.Interface().(int)
		break
	case reflect.Uint:
		return int(val1.Interface().(uint)) - int(val2.Interface().(uint))
		break
	case reflect.Uint8:
		return int(val1.Interface().(uint8)) - int(val2.Interface().(uint8))
		break
	case reflect.Uint16:
		return int(val1.Interface().(uint16)) - int(val2.Interface().(uint16))
		break
	case reflect.Uint32:
		return int(val1.Interface().(uint32)) - int(val2.Interface().(uint32))
		break
	case reflect.Uint64:
		v1 := val1.Interface().(uint64)
		v2 := val2.Interface().(uint64)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
		break
	case reflect.Float32:
		v1 := val1.Interface().(float32)
		v2 := val2.Interface().(float32)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
		break
	case reflect.Float64:
		v1 := val1.Interface().(float64)
		v2 := val2.Interface().(float64)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
		break
	default:
		compareFuncValue := val1.MethodByName(compareFuncName)
		return compareFuncValue.Call([]reflect.Value{val2})[0].Interface().(int)
	}

	return 0
}

// swap two elements of the slice
func swap(sliceValue reflect.Value, index1, index2 int) {
	if index1 == index2 {
		return
	}

	cache := clone(sliceValue.Index(index1))
	sliceValue.Index(index1).Set(sliceValue.Index(index2))
	sliceValue.Index(index2).Set(cache)
}
