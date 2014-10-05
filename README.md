generic
=======

It's a Golang generic library

Motivations
-----------
*   Golang doesn't support generics now.
*   We need generics in projects. 
*   There is not a thirdparty generic library now.
*   Challenge myself.

Features
----------
*   Remove element in slice of any type. (available now) API: [RemoveAt](#api-slice-removeAt) [Remove](#api-slice-remove) [RemoveBy](#api-slice-removeBy)
*   Iterate elements in slice. API: [Each](#api-slice-each) [ForEach](#api-slice-forEach)
*   Sort elements in slice of any type. (available now, support int8, int16, int32, int, int64, uint8, uint16, uint32, uint, uint64, float32, float64 and struct which contains a compare function) API: [QuickSort](#api-slice-quicksort)
*   Find element in slice of any type. API: [Find](#api-slice-find) [FindBy](#api-slice-findBy)


APIs
-----------
*   <a name="api-slice-removeAt" id="api-slice-removeAt">RemoveAt</a>
    >`func (s *slice) RemoveAt(index int) error`
 
    > Remove element of slice at index. The slice can be any type slice, include struct slice.
    
    > Example
    
    >```
    >intSlice := []int{1, 2, 3}
    >Slice(&intSlice).RemoveAt(1)
    >fmt.Println(intSlice) // the result should be [1, 3]
    >```

*   <a name="api-slice-remove" id="api-slice-remove">Remove</a>
    >`func (s *slice) Remove(elem interface{}) error`
 
    > Remove element of slice. The slice can be any type slice, include struct slice. The parameter `elem` is contained by slice.
    
    > Example
    
    >```
    >byteSlice := []byte{1, 2, 3}
    >Slice(&byteSlice).Remove(byte(1))
    >fmt.Println(byteSlice) // the result should be [2, 3]
    >```

*   <a name="api-slice-removeBy" id="api-slice-removeBy">RemoveBy</a>
    >`func (s *slice) RemoveBy(equal func(interface{}) bool) error`
 
    > Remove element of slice when `equal` function return true. The slice can be any type slice, include struct slice. 
    
    > Example
    
    >```
    >values := []byte{1, 2, 3}
    >err := Slice(&values).RemoveBy(func(value interface{}) bool {
    >    elem := value.(byte)
    >    return elem == byte(1)
    >})
    >fmt.Println(byteSlice) // the result should be [2, 3]
    >```

*   <a name="api-slice-find" id="api-slice-find">Find</a>
    >`func (s *slice) Find(elem interface{}) (int, error)`
 
    > Find element of slice. The parameter `elem` is contained by slice. Return -1 if not find, otherwise return the index of element in slice. It is available to any type slice, include struct slice. 
    
    > Example
    
    >```
    >byteSlice := []byte{1, 2, 3}
    >index, err := Slice(&byteSlice).Find(byte(1))
    >fmt.Println(index) // the result should be 0
    >```

*   <a name="api-slice-findBy" id="api-slice-findBy">FindBy</a>
    >`func (s *slice) FindBy(equal func(interface{}) bool) (int, error) `
 
    > Find element of slice when `equal` function return true. Return -1 if not find, otherwise return the index of element in slice. It is available to any type slice, include struct slice. 
    
    > Example
    
    >```
    >values := []byte{1, 2, 3}
    >index, err := Slice(&values).FindBy(func(value interface{}) bool {
    >    elem := value.(byte)
    >    return elem == byte(1)
    >})
    >fmt.Println(index) // the result should be 0
    >```

*   <a name="api-slice-each" id="api-slice-each">Each</a>
    >`func (s *slice) Each(iterate func(interface{}, int)) error`
 
    > Iterate to each element in slice. And then you can do anything in iterate function. The slice can be any type slice, include struct slice. 
    
    > Example
    
    >```
    >values := []byte{1, 2, 3}
    >sum := 0
    >err := Slice(&values).Each(func(value interface{}, index int) {
    >    elem := value.(byte)
    >    sum = sum + int(elem)
    >})
    >fmt.Println(sum) // the result should be 6
    >```

*   <a name="api-slice-forEach" id="api-slice-forEach">ForEach</a>
    >`func (s *slice) ForEach(iterate func(interface{}, int)) error`
 
    > Iterate to each element in slice. It is same as Each. The slice can be any type slice, include struct slice. 
    
    > Example
    
    >```
    >values := []byte{1, 2, 3}
    >sum := 0
    >err := Slice(&values).ForEach(func(value interface{}, index int) {
    >    elem := value.(byte)
    >    sum = sum + int(elem)
    >})
    >fmt.Println(sum) // the result should be 6
    >```

*   <a name="api-slice-quicksort" id="api-slice-quicksort">QuickSort</a>
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
 
Helping Generic
-----------

#### I found a bug!

If you found a bug, please [search existing issues](https://github.com/anzhihun/generic/issues) first  to
see if it's already there. If not, please create a new [issue](https://github.com/anzhihun/generic/issues), Include steps to consistently reproduce the problem, actual vs. expected results, screenshots, and your OS and
Generic version number. 


#### I have a new suggestion, but don't know how to program!

For feature requests please [search existing feature issues](https://github.com/anzhihun/generic/issues) to
see if it's already there; you can comment to upvote it if so. If not, feel free to create an new issue; we'll
change it to the feature issue for you.


#### I want to help with the code!

Awesome! Please feel free to push your request.

License
-----------
MIT