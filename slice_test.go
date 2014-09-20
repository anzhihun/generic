package generic

import (
	"strconv"
	"testing"
)

type student struct {
	name string
	age  int
}

func (s student) Compare(other student) int {
	if s.age < other.age {
		return -1
	} else if s.age == other.age {
		return 0
	} else {
		return 1
	}
}

func TestSliceRemoveAt(t *testing.T) {
	values := []byte{1, 2, 3}
	err := Slice(&values).RemoveAt(0)
	if err != nil || len(values) != 2 || values[0] != 2 || values[1] != 3 {
		t.Fatal("Failed to RemoveAt first item!")
	}

	values2 := []byte{1, 2, 3}
	err = Slice(&values2).RemoveAt(2)
	if err != nil || len(values2) != 2 || values2[0] != 1 || values2[1] != 2 {
		t.Fatal("Failed to RemoveAt last item!")
	}

	values3 := []byte{1, 2, 3}
	err = Slice(&values3).RemoveAt(1)
	if err != nil || len(values3) != 2 || values3[0] != 1 || values3[1] != 3 {
		t.Fatal("Failed to RemoveAtAt middle item!")
	}

	values4 := []byte{1, 2, 3}
	err = Slice(&values4).RemoveAt(3)
	if err == nil {
		t.Fatal("It should be error when removing out of range item!")
	}

	values5 := []byte{1, 2, 3}
	err = Slice(values5).RemoveAt(1)
	if err == nil {
		t.Fatal("It should be error when the parameter is slice!")
	}

	values6 := [3]byte{1, 2, 3}
	err = Slice(&values6).RemoveAt(1)
	if err == nil {
		t.Fatal("It should be error when the parameter is array pointer!")
	}

	values7 := string("test")
	err = Slice(&values7).RemoveAt(1)
	if err == nil {
		t.Fatal("It should be error when the parameter is string!")
	}
}

func TestSliceRemove(t *testing.T) {
	values := []byte{1, 2, 3}
	err := Slice(&values).Remove(byte(1))
	if err != nil || len(values) != 2 || values[0] != 2 || values[1] != 3 {
		t.Fatal("Failed to Remove first item!")
	}

	err = Slice(&values).Remove(int(2))
	if err != nil && len(values) == 1 {
		t.Fatal("should not remove byte value by int value!")
	}

	students := []student{}
	students = append(students, student{name: "1", age: 100})
	students = append(students, student{name: "2", age: 100})
	students = append(students, student{name: "3", age: 100})
	err = Slice(&students).Remove(student{name: "3", age: 100})
	if err != nil || len(values) != 2 || students[0].name != "1" || students[1].name != "2" {
		t.Fatal("failed to remove struct from slice!")
	}
}

func TestSliceRemoveBy(t *testing.T) {
	values := []byte{1, 2, 3}
	err := Slice(&values).RemoveBy(func(value interface{}) bool {
		elem := value.(byte)
		return elem == byte(1)
	})

	if err != nil || len(values) != 2 || values[0] != 2 || values[1] != 3 {
		t.Fatal("Failed to remove first item through RemoveBy!")
	}

	students := []student{}
	students = append(students, student{name: "1", age: 100})
	students = append(students, student{name: "2", age: 100})
	students = append(students, student{name: "3", age: 100})
	err = Slice(&students).RemoveBy(func(value interface{}) bool {
		elem := value.(student)
		return elem.name == "3"
	})
	if err != nil || len(values) != 2 || students[0].name != "1" || students[1].name != "2" {
		t.Fatal("failed to remove struct from slice through RemoveBy!")
	}
}

func TestSliceForEach(t *testing.T) {
	values := []byte{1, 2, 3}
	sum := 0
	err := Slice(&values).ForEach(func(value interface{}, index int) {
		elem := value.(byte)
		sum = sum + int(elem)
	})

	if err != nil || sum != 6 {
		t.Fatal("Failed to iterate element of slice!")
	}
}

func TestSliceEach(t *testing.T) {
	values := []byte{1, 2, 3}
	sum := 0
	err := Slice(&values).Each(func(value interface{}, index int) {
		elem := value.(byte)
		sum = sum + int(elem)
	})

	if err != nil || sum != 6 {
		t.Fatal("Failed to iterate element of slice!")
	}
}

func TestSliceQuickSort_Struct(t *testing.T) {
	students := []student{}
	err := Slice(&students).QuickSort()
	if err != nil {
		t.Fatal("Failed to quick sort empty slice! error: ", err)
	}

	students = append(students, student{name: "1", age: 100})
	if err = Slice(&students).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort slice contains only 1 element! error: ", err)
	}
	if len(students) != 1 || students[0].age != 100 {
		t.Fatal("After quick sort slice contains only 1 element, the element should be right!")
	}

	students = append(students, student{name: "2", age: 12})
	if err = Slice(&students).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort slice contains only 2 element! error: ", err)
	}
	if len(students) != 2 || students[0].age != 12 || students[1].age != 100 {
		t.Fatal("After quick sort slice contains only 2 element, the elements should be right and ordered!")
	}

	students = append(students, student{name: "3", age: 15})
	students = append(students, student{name: "4", age: 14})
	students = append(students, student{name: "5", age: 11})
	students = append(students, student{name: "6", age: 22})
	students = append(students, student{name: "7", age: 4})

	if err = Slice(&students).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort slice contains many elements! error: ", err)
	}
	for i := 0; i < len(students)-2; i++ {
		if students[i].age > students[i+1].age {
			t.Fatal("After quick sort slice contains many elements, the elements should be right and ordered! actual: ", students)
		}
	}

	// ordered slice
	students2 := []student{}
	for i := 0; i < 1000; i++ {
		students2 = append(students2, student{name: strconv.Itoa(i), age: i})
	}
	if err = Slice(&students2).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort ordered slice! error: ", err)
	}

	for i := 0; i < 1000; i++ {
		if students2[i].age != i {
			t.Fatal("After quick sort ordered slice, the elements should be right and ordered! actual: ", students2)
		}
	}

	students3 := []student{}
	for i := 999; i >= 0; i-- {
		students3 = append(students3, student{name: strconv.Itoa(i), age: i})
	}
	if err = Slice(&students3).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort reverse order slice! error: ", err)
	}

	count := len(students)
	for i := 0; i < count; i++ {
		if students3[i].age != i {
			t.Fatal("After quick sort reverse order slice, the elements should be right and ordered! actual: ", students3)
		}
	}

	students4 := []student{}
	students4 = append(students4, student{name: strconv.Itoa(10), age: 10})
	students4 = append(students4, student{name: strconv.Itoa(11), age: 9})
	students4 = append(students4, student{name: strconv.Itoa(12), age: 9})
	students4 = append(students4, student{name: strconv.Itoa(13), age: 8})
	students4 = append(students4, student{name: strconv.Itoa(14), age: 10})
	if err = Slice(&students4).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort slice contains repeat elements! error:", err)
	}

	count = len(students4)
	for i := 0; i < count-1; i++ {
		if students4[i].age > students4[i+1].age {
			t.Fatal("After quick sort slice contains repeat elements, the elements should be right and ordered! actual: ", students4)
		}
	}
}

func TestSliceQuickSort_Int8(t *testing.T) {
	testInt8s := []int8{}
	err := Slice(&testInt8s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support int8 slice! error: ", err)
	}

	testInt8s = append(testInt8s, 10)
	if err = Slice(&testInt8s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int8 slice contains with 1 element! error: ", err)
	}
	if len(testInt8s) != 1 || testInt8s[0] != 10 {
		t.Fatal("After quick sort int8 slice contains 1 element, the element should be right and ordered!")
	}

	testInt8s = append(testInt8s, 1)
	if err = Slice(&testInt8s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int8 slice contains 2 element! error: ", err)
	}
	if len(testInt8s) != 2 || testInt8s[0] != 1 || testInt8s[1] != 10 {
		t.Fatal("After quick sort int8 slice contains 2 element, the elements should be right and ordered!")
	}

	moreInt8s := []int8{}
	var index int8 = 0
	for index = 126; index >= -127; index-- {
		moreInt8s = append(moreInt8s, index)
	}
	if err = Slice(&moreInt8s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int8 slice contains many reverse order elements! error: ", err)
	}
	count := len(moreInt8s)
	if count != 254 {
		t.Fatal("After quick sort int8 slice contains many reverse order elements, the element count should be right!")
	}

	var i int
	for index = -127; index < 127; index++ {
		i = int(index) + 127
		if moreInt8s[i] != index {
			t.Fatal("After quick sort int8 slice contains many reverse order elements, the elements should be right and ordered!")
		}
	}
}

func TestSliceQuickSort_Int16(t *testing.T) {
	testint16s := []int16{}
	err := Slice(&testint16s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support int16 slice! error: ", err)
	}

	testint16s = append(testint16s, 30000)
	if err = Slice(&testint16s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int16 slice contains 1 element! error: ", err)
	}
	if len(testint16s) != 1 || testint16s[0] != 30000 {
		t.Fatal("After quick sort int16 slice contains 1 element, the elements should be right and ordered!")
	}

	testint16s = append(testint16s, -30000)
	if err = Slice(&testint16s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int16 slice contains 2 element! error: ", err)
	}
	if len(testint16s) != 2 || testint16s[0] != -30000 || testint16s[1] != 30000 {
		t.Fatal("After quick sort int16 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Int32(t *testing.T) {
	testint32s := []int32{}
	err := Slice(&testint32s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support int32 slice! error: ", err)
	}

	testint32s = append(testint32s, 2000000000)
	if err = Slice(&testint32s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int32 slice contains 1 element! error: ", err)
	}
	if len(testint32s) != 1 || testint32s[0] != 2000000000 {
		t.Fatal("After quick sort int32 slice contains 1 element, the elements should be right and ordered!")
	}

	testint32s = append(testint32s, -2000000000)
	if err = Slice(&testint32s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int32 slice contains 2 element! error: ", err)
	}
	if len(testint32s) != 2 || testint32s[0] != -2000000000 || testint32s[1] != 2000000000 {
		t.Fatal("After quick sort int32 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Int64(t *testing.T) {
	testint64s := []int64{}
	err := Slice(&testint64s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support int64 slice! error: ", err)
	}

	testint64s = append(testint64s, 6000000000)
	if err = Slice(&testint64s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int64 slice contains 1 element! error: ", err)
	}
	if len(testint64s) != 1 || testint64s[0] != 6000000000 {
		t.Fatal("After quick sort int64 slice contains 1 element, the elements should be right and ordered!")
	}

	testint64s = append(testint64s, -6000000000)
	if err = Slice(&testint64s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int64 slice contains 2 element! error: ", err)
	}
	if len(testint64s) != 2 || testint64s[0] != -6000000000 || testint64s[1] != 6000000000 {
		t.Fatal("After quick sort int64 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Int(t *testing.T) {
	testints := []int{}
	err := Slice(&testints).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support int slice! error: ", err)
	}

	testints = append(testints, 75536)
	if err = Slice(&testints).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int slice contains 1 element! error: ", err)
	}
	if len(testints) != 1 || testints[0] != 75536 {
		t.Fatal("After quick sort int slice contains 1 element, the elements should be right and ordered!")
	}

	testints = append(testints, 65536)
	if err = Slice(&testints).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort int slice contains 2 element! error: ", err)
	}
	if len(testints) != 2 || testints[0] != 65536 || testints[1] != 75536 {
		t.Fatal("After quick sort int slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Uint(t *testing.T) {
	testuints := []uint{}
	err := Slice(&testuints).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support uint slice! error: ", err)
	}

	testuints = append(testuints, 10)
	if err = Slice(&testuints).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint slice contains 1 element! error: ", err)
	}
	if len(testuints) != 1 || testuints[0] != 10 {
		t.Fatal("After quick sort uint slice contains 1 element, the elements should be right and ordered!")
	}

	testuints = append(testuints, 1)
	if err = Slice(&testuints).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint slice contains 2 element! error: ", err)
	}
	if len(testuints) != 2 || testuints[0] != 1 || testuints[1] != 10 {
		t.Fatal("After quick sort uint slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Uint8(t *testing.T) {
	testuint8s := []uint8{}
	err := Slice(&testuint8s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support uint8 slice! error: ", err)
	}

	testuint8s = append(testuint8s, 255)
	if err = Slice(&testuint8s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint8 slice contains 1 element! error: ", err)
	}
	if len(testuint8s) != 1 || testuint8s[0] != 255 {
		t.Fatal("After quick sort uint8 slice contains 1 element, the elements should be right and ordered!")
	}

	testuint8s = append(testuint8s, 128)
	if err = Slice(&testuint8s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint8 slice contains 2 element! error: ", err)
	}

	if len(testuint8s) != 2 || testuint8s[0] != 128 || testuint8s[1] != 255 {
		t.Fatal("After quick sort uint8 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Uint16(t *testing.T) {
	testuint16s := []uint16{}
	err := Slice(&testuint16s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support uint16 slice! error: ", err)
	}

	testuint16s = append(testuint16s, 65535)
	if err = Slice(&testuint16s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint16 slice contains 1 element! error: ", err)
	}
	if len(testuint16s) != 1 || testuint16s[0] != 65535 {
		t.Fatal("After quick sort uint16 slice contains 1 element, the elements should be right and ordered!")
	}

	testuint16s = append(testuint16s, 15535)
	if err = Slice(&testuint16s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint16 slice contains 2 element! error: ", err)
	}

	if len(testuint16s) != 2 || testuint16s[0] != 15535 || testuint16s[1] != 65535 {
		t.Fatal("After quick sort uint16 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Uint32(t *testing.T) {
	testuint32s := []uint32{}
	err := Slice(&testuint32s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support uint32 slice! error: ", err)
	}

	testuint32s = append(testuint32s, 4000000000)
	if err = Slice(&testuint32s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint32 slice contains 1 element! error: ", err)
	}
	if len(testuint32s) != 1 || testuint32s[0] != 4000000000 {
		t.Fatal("After quick sort uint32 slice contains 1 element, the elements should be right and ordered!")
	}

	testuint32s = append(testuint32s, 1000000000)
	if err = Slice(&testuint32s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint32 slice contains 2 element! error: ", err)
	}

	if len(testuint32s) != 2 || testuint32s[0] != 1000000000 || testuint32s[1] != 4000000000 {
		t.Fatal("After quick sort uint32 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Uint64(t *testing.T) {
	testuint64s := []uint64{}
	err := Slice(&testuint64s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support uint64 slice! error: ", err)
	}

	testuint64s = append(testuint64s, 10000000001000000000)
	if err = Slice(&testuint64s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint64 slice contains 1 element! error: ", err)
	}
	if len(testuint64s) != 1 || testuint64s[0] != 10000000001000000000 {
		t.Fatal("After quick sort uint64 slice contains 1 element, the elements should be right and ordered!")
	}

	testuint64s = append(testuint64s, 10000000000000000000)
	if err = Slice(&testuint64s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort uint64 slice contains 2 element! error: ", err)
	}
	if len(testuint64s) != 2 || testuint64s[0] != 10000000000000000000 || testuint64s[1] != 10000000001000000000 {
		t.Fatal("After quick sort uint64 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Float32(t *testing.T) {
	testfloat32s := []float32{}
	err := Slice(&testfloat32s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support float32 slice! error: ", err)
	}

	testfloat32s = append(testfloat32s, 12.1)
	if err = Slice(&testfloat32s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort float32 slice contains 1 element! error: ", err)
	}
	if len(testfloat32s) != 1 || testfloat32s[0] != 12.1 {
		t.Fatal("After quick sort float32 slice contains 1 element, the elements should be right and ordered!")
	}

	testfloat32s = append(testfloat32s, 11.9)
	if err = Slice(&testfloat32s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort float32 slice contains 2 element! error: ", err)
	}
	if len(testfloat32s) != 2 || testfloat32s[0] != 11.9 || testfloat32s[1] != 12.1 {
		t.Fatal("After quick sort float32 slice contains 2 element, the elements should be right and ordered!")
	}
}

func TestSliceQuickSort_Float64(t *testing.T) {
	testfloat64s := []float64{}
	err := Slice(&testfloat64s).QuickSort()
	if err != nil {
		t.Fatal("Quick sort should support float64 slice! error: ", err)
	}

	testfloat64s = append(testfloat64s, 12.1)
	if err = Slice(&testfloat64s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort float64 slice contains 1 element! error: ", err)
	}
	if len(testfloat64s) != 1 || testfloat64s[0] != 12.1 {
		t.Fatal("After quick sort float64 slice contains 1 element, the elements should be right and ordered!")
	}

	testfloat64s = append(testfloat64s, 11.9)
	if err = Slice(&testfloat64s).QuickSort(); err != nil {
		t.Fatal("Failed to quick sort float64 slice contains 2 element! error: ", err)
	}
	if len(testfloat64s) != 2 || testfloat64s[0] != 11.9 || testfloat64s[1] != 12.1 {
		t.Fatal("After quick sort float64 slice contains 2 element, the elements should be right and ordered!")
	}
}
