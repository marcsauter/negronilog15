# negronilog15 [![godoc reference](https://godoc.org/gopkg.in/inconshreveable/log15.v2?status.png)](https://github.com/marcsauter/negronilog15)

[log15](https://github.com/inconshreveable/log15) middleware for [negroni](https://github.com/codegangsta/negroni)

## Example

```go
 n := negroni.New(negronilog15.NewMiddleware())
```
