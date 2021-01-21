### Runner 

**Runner** provides simple wrapper for operation which may fails, thus can be retried.
It handle common OS signals which indicate program being interupted, closed etc. 

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