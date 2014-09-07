generic
=======

It's a Golang generic library

APIs
-----------
*   Remove
    >`func (s *slice) Remove(index int) error`
 
    > remove element of slice at index. The slice can be any type slice, include struct slice. 
    
    > Example
    
    > `intSlice := []{1, 2, 3}
    
    Slice(&intSlice).Remove(1)
    
    fmt.Println(intSlice) // the result should be [1, 3]
    `
 