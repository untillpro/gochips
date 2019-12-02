# Errors
Multiple errors container. Implemented as wrapper for `[]error`.
# Usage
```go
var errs errs.Errors
errs.Add("error")
fmt.Println(errs) 
// output: error

errs.Add("next error")
fmt.Println(errs) 
// output: Multiple errors: 
//		error
//		next error  

errs.AddE(errors.New("3rd error"))
fmt.Println(errs) 
// output: Multiple errors: 
//		error
//		next error  
//		3rd error

errs = append(errs, errors.New("4th error"))
fmt.Println(errs) 
// output: Multiple errors: 
//		error
//		next error  
//		3rd error
//		4th error

fmt.Printn(errs[3].Error())
// output: 4th error
```