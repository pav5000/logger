# Logger
An easy way to add structured logs to your project. Based on uber-go/zap.

```go
import (
    "net/http"

    "github.com/pav5000/logger"
    "go.uber.org/zap"
)

const addr = ":8080"

func main() {
    logger.Info("listening http", zap.String("addr", addr))
    err := http.ListenAndServe(addr, nil)
    logger.Fatal("http.ListenAndServe exited", zap.Error(err))
}
```

## Configuration

You can enable dev mode and/or enable adding stacktraces to messages of level Error and above (both are turned off by default).
```go
import (
	"flag"

	"github.com/pav5000/logger"
	"go.uber.org/zap"
)

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "Debug mode")
	flag.Parse()

	logger.Init(logger.Settings{
		DevMode: debugMode,
        StackTracedErrors: true,
	})

    // ......
}
```


In dev mode messages look like this:
```
2023-02-19T17:39:50.896+0300    INFO    app/main.go:22  listening http  {"addr": ":8080"}
```

In production mode (default) messages look like this:
```
{"level":"info","ts":1676818605.186886,"msg":"listening http","addr":":8080"}
```
