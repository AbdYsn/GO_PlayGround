package abdysn_playground

import (
   "fmt"
   "rsc.io/quote"
)

func Hello() string{
        fmt.Println("Hello, playground")
        return quote.Hello()
}
