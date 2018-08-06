# argv

Extremely minimalistic argument parser for Go CLI apps. There is only
a single function `Parse` that returns a map[string]string of all key/value
pairs for arguments provided.

## Rules

1. Any arguments that start with one or two hyphens (-arg or --arg) are considered keys
2. The value immediately following a key is considered that keys value.
  - NOTE: if the value of a key starts with a hyphen, you need to wrap the value in quotes,
otherwise it will be considered another key.

## Install

```sh
$ go get github.com/whatl3y/argv
```

## Usage

The following script is an example of how to use the library.

```go
// ./test.go
package main

import (
  "fmt"
  "os"
  "github.com/whatl3y/argv"
)

func main() {
  theArgs := argv.Parse(os.Args[1:])
  fmt.Println("Arguments:", theArgs)
}
```

```sh
$ go build
$ ./test -arg1 123 --arg2 456
# Arguments: map[arg1:123 arg2:456]

$ ./test -arg1 123 --arg2 --arg3 456 --arg4
# Arguments: map[arg1:123 arg2: arg3:456 arg4:]
```
