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

// Remove element of slice
func (s *slice) Remove(elem interface{}) error {
	err := s.checkSlice()
	if err != nil {
		return err
	}

	slicePtrValue := reflect.ValueOf(s.slicePtr)
	sliceValue := slicePtrValue.Elem()
	if sliceValue.Len() <= 0 {
		return nil
	}

	for index := 0; index < sliceValue.Len(); index++ {
		if reflect.DeepEqual(sliceValue.Index(index).Interface(), elem) {
			s.RemoveAt(index)
			return nil
		}
	}

	return nil
}

// Remove element of slice when equal function return true
func (s *slice) RemoveBy(equal func(interface{}) bool) error {
	err := s.checkSlice()
	if err != nil {
		return err
	}

	slicePtrValue := reflect.ValueOf(s.slicePtr)
	sliceValue := slicePtrValue.Elem()
	if sliceValue.Len() <= 0 {
		return nil
	}

	for index := 0; index < sliceValue.Len(); index++ {
		if equal(sliceValue.Index(index).Interface()) {
			s.RemoveAt(index)
			return nil
		}
	}

	return nil
}

// Iterate to each element in slice. And then you can do anything in iterate function.
func (s *slice) ForEach(iterate func(interface{}, int)) error {
	err := s.checkSlice()
	if err != nil {
		return err
	}

	slicePtrValue := reflect.ValueOf(s.slicePtr)
	sliceValue := slicePtrValue.Elem()
	if sliceValue.Len() <= 0 {
		return nil
	}

	for index := 0; index < sliceValue.Len(); index++ {
		iterate(sliceValue.Index(index).Interface(), index)
	}

	return nil
}

// Iterate to each element in slice. It is same as ForEach
func (s *slice) Each(iterate func(interface{}, int)) error {
	return s.ForEach(iterate)
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

	go quickSort(slice, lowIndex, firstIndex-1, compareFuncName)
	go quickSort(slice, lastIndex+1, highIndex, compareFuncName)
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
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Int:
		return int(val1.Int() - val2.Int())
		break
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return int(val1.Uint() - val2.Uint())
		break
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		v1 := val1.Float()
		v2 := val2.Float()
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
		return int(compareFuncValue.Call([]reflect.Value{val2})[0].Int())
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
