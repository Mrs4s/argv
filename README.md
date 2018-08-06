# argv

Extremely minimalistic argument parser for Go CLI apps. There is only
a single function #Parse that returns a map[string]string of all key/value
pairs for arguments provided.

## Rules

- Any arguments that start with one or two hyphens (-arg or --arg) are considered keys
- The value immediately following a key is considered that keys value.
NOTE: if the value of a key starts with a hyphen, you need to wrap the value in quotes,
otherwise it will be considered another key.

## Install

```sh
$ go get github.com/whatl3y/argv
```

## Usage

If you setup the following script, you should see the following
output to the console given the provided arguments.

```go
// ./argv.go
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
$ ./argv -arg1 123 --arg2 456
# Arguments: map[arg1:123 arg2:456]

$ ./argv -arg1 123 --arg2 --arg3 456 --arg4
# Arguments: map[arg1:123 arg2: arg3:456 arg4:]
```
