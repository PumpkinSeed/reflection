# Reflection

The package try to solve the type issues when reflect.Set happens.

#### Usage 

```go
package main

import (
    "reflect"

    "github.com/PumpkinSeed/reflection"
    "github.com/volatiletech/null"
)

type Example struct {
    Integer     int
    String      string
    Nullable    null.Uint
    StringPtr   *string
}

func main() {
    var example = &Example{}
    
    err := reflection.Set(reflect.ValueOf(example.Integer), "12")
    // err check
    err = reflection.Set(reflect.ValueOf(example.String), 12)
    // err check
    err = reflection.Set(reflect.ValueOf(example.Nullable), 12)
    // err check
    err = reflection.Set(reflect.ValueOf(example.StringPtr), "12")
    // err check
}
```