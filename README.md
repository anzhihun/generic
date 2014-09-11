generic
=======

It's a Golang generic library

Motivations
-----------
*   Golang doesn't support generics now.
*   We need generics in projects. 
*   There is not a thirdpart generic library now.
*   Challenge myself.

Features
----------
*   Remove element in slice of any type. (available now)[see](#api-slice-remove)
*   Sort elements in slice of any type. (available now, support int8, int16, int32, int, int64, uint8, uint16, uint32, uint, uint64 and struct which contains a compare function)
*   Find element in slice of any type. 
*   


APIs
-----------
*   <a name="api-slice-remove" id="api-slice-remove">Remove</a>
    >`func (s *slice) Remove(index int) error`
 
    > Remove element of slice at index. The slice can be any type slice, include struct slice.
    
    > Example
    
    >```
    >intSlice := []int{1, 2, 3}
    >Slice(&intSlice).Remove(1)
    >fmt.Println(intSlice) // the result should be [1, 3]
    >```
    
*   QuickSort
    >`func (s *slice) QuickSort() error `
 
    > Sort the elements of slice in ascending order. The slice can be any int and uint slice, and struct slice.  The struct must contains the compare function `func (s structName) Compare(other structName) int`, which should return a int value to indicate which one is more greater. If the return value is equal to 0. The element is equal to other. If the return value is less than 0. The other element is more greater. If the return value is greater than 0. The other element is more less.
    
    > Example
    
    >```
    >type student struct {
    >   age int
    >}
    >
    >func (s student) Compare(other student) int {
    >   return s.age - other.age
    >}
    >
    >students := []student{}
    >students = append(students, student{3})
    >students = append(students, student{1})
    >students = append(students, student{5})
    >Slice(&students).QuickSort()
    >fmt.Println(students) // the result should be [{1} {3} {5}]
    >```
 