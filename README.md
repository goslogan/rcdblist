# rcutils

A simple library of code to parse the logs returned from Redis Cloud into
structs. Contains parsers for both the database and the system log exports.

## Examples

```
package main

import (
    "log"
    "os"

    "github.com/dimchansky/utfbom"
    "github.com/goslogan/rcutils"
)


func main() {

    if reader, err := os.Open("databases.csv"); err != nil {
            log.Fatalf("Unable to open databases.csv - %v", err)
    } else if databases, err := rcutils.Databases(utfbom.SkipOnly(reader)); err != nil {
            log.Fatalf("Unable to parse databases.csv - %v", err)
    } else {
        // PROCESS DATABASE OUTPUT
    }
}
```


```
package main

import (
    "log"
    "os"

    "github.com/dimchansky/utfbom"
    "github.com/goslogan/rcutils"
)


func main() {

    if reader, err := os.Open("system_log.csv"); err != nil {
            log.Fatalf("Unable to open system_log.csv - %v", err)
    } else if events, err := rcutils.SystemLog(utfbom.SkipOnly(reader), func(e *rcutils.LogEvent) bool { return e.Activity == "Configuration" }); err != nil {
            log.Fatalf("Unable to parse system_log.csv - %v", err)
    } else {
        // PROCESS CONFIG EVENTSs
    }
}
```

