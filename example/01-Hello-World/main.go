package main

// http://play.golang.org/p/jZ5pa944O1 <- will not display the colors
import "fmt"

const (
        InfoColor    = "\033[0;31m%s\033[0m"
        NoticeColor  = "\033[0;32m%s\033[0m"
        WarningColor = "\033[0;33m%s\033[0m"
        ErrorColor   = "\033[0;34m%s\033[0m"
        DebugColor   = "\033[0;35m%s\033[0m"
)

func main() {
        fmt.Printf("\033[0;31m%s\033[0m\n", "Hello, world.")
        fmt.Printf("\033[0;32m%s\033[0m\n", "Hello, world.")
        fmt.Printf("\033[0;33m%s\033[0m\n", "Hello, world.")
        fmt.Printf("\033[0;34m%s\033[0m\n", "Hello, world.")
        fmt.Printf("\033[0;35m%s\033[0m\n", "Hello, world.")
        fmt.Printf("\033[0;36m%s\033[0m\n", "Hello, world.")
        fmt.Printf("\033[0;37m%s\033[0m\n", "Hello, world.")
        fmt.Printf("\033[0;38m%s\033[0m\n", "Hello, world.")
		fmt.Scanln()
}