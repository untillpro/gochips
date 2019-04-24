# gochips

Small useful pieces of golang code

```go
import (
  gc "github.com/untillpro/gochips"
) 
```

# Chips for Tests

```go
func Test_Get(t *testing.T) {
  err := MyFunc()
  gc.FatalIfError(t, err, "Error in MyFunc()")
}
```

# Chips for Logging


```go

  gc.Error(args ...interface{})
  gc.Info(args ...interface{})
  
  // Adds "..." suffux to  argument
  gc.Doing("Getting things done")

  // Does nothing until you set gc.IsVerbose = true
  gc.Verbose(args ...interface{})

}
```
Prefixes:

- Verbose() output is prefixed with `gc.VerbosePrefix = "--- "`
- Error() output is prefixed with `gc.ErrorPrefix = "*** "`

Change default behavior:
- `var Output func(funcName, s string)`
  - This function is called by default implementation of `Error`, `Info` etc
- It is possible to redefine all implementations, ref:
```go
func init() {
	Doing = implDoing
	Info = implInfo
	Verbose = implVerbose
	Error = implError
	Output = implOutput
	VerboseWriters = implVerboseWriters
}
```



# Chips for Exeс

Piped execution a-la shell

`echo README.md | grep README.md | sed s/READ/forgive/`:
```go
	err := new(PipedExec).
		Command("echo", "README.md").WorkingDir("/").
		Command("grep", "README.md").
		Command("echo", "good").
		Run(os.Stdout, os.Stdout)
```

Avoid stdout until `gc.IsVerbose` is set to true:

```go
	err := new(PipedExec).
		Command("echo", "README.md").WorkingDir("/").
		Command("grep", "README.md").
		Command("echo", "good").
		Run(gc.VerboseWriters())
```


Ref. also [PipedExec_test.go](PipedExec_test.go)





