// Echo1 prints its command line arguments
package main

import "fmt"
import "os"

func main () {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " -_- "
    }
    fmt.Println(s)
}
