# Errors
Multiple errors container. Implemented as wrapper for `[]error`.
# Usage
```go
var errs Error
errs = errs.Add("error")
fmt.Println(errs) // output: error
errs = errs.Add("next error")
fmt.Println(errs) // output: Multiple errors: 
				  //		 error
				  //		 next error  
errs = errs.AddE(errors.New("3rd error"))
fmt.Println(errs) // output: Multiple errors: 
				  //		 error
				  //		 next error  
				  //         3rd error
```