# argv

Extremely minimalistic argument parser for Go CLI apps. There is only
a single function `Parse` that returns a struct of all key/value
pairs for arguments provided in addition to any arguments provided
without corresponding keys.

## Rules

1. Any arguments that start with one or two hyphens (-arg or --arg) are considered keys
2. The value immediately following a key is considered that keys value.
  - NOTE: if the value of a key starts with a hyphen, you need to wrap the value in quotes,
otherwise it will be considered another key.
3. If an argument is provided without a key, it will be included in a []string field
of the returned struct called `nokeys`.

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
  // theArgs is a:
  // struct{
  //  nokeys []string
	//  keys   map[string]string
  //}
  theArgs := argv.Parse(os.Args[1:])
  fmt.Println("Arguments:", theArgs)
}
```

```sh
$ go build
$ ./test -arg1 123 --arg2 456 nokeyarg
# Arguments: {[nokeyarg] map[arg1:123 arg2:456]}

$ ./test -arg1 123 --arg2 --arg3 456 nokeyarg --arg4
# Arguments: {[nokeyarg] map[arg1:123 arg2: arg3:456 arg4:]}
```
