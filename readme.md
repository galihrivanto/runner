<div align="center">
  <h1>runner</h1>
  
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/galihrivanto/runner)
  
**Runner** provides simple wrapper for operation which may fails, thus can be retried.
It handle common OS signals which indicate program being interupted, closed etc.  
</div>


simple usage:
```go
ctx, cancel := context.WithCancel(context.Background())

runner.RunWithRetry(func(mctx context.Context) error {
    ...
})
.Handle(func(sig os.Signal){
    cancel()
})

```

# license
[MIT](https://choosealicense.com/licenses/mit/)