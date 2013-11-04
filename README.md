logger
======

``logger`` is a simple level-logger.

## Installation

```sh
$ go get github.com/gonuts/logger
```

## Example

```go
package main

import (
	"github.com/gonuts/logger"
)

func main() {
	msg := logger.New("foo")
	msg.SetLevel(logger.DEBUG)
	msg.Infof("an info message [%s]\n", "hello")
	msg.Debugf("a debug message [%s]\n", "hallo")
	msg.Errorf("an error message [%s]\n", "hello")
}
```
