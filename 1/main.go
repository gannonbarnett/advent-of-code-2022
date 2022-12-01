package main

import (  
    "fmt"
    "os"
    "log"
    "strconv"
    "bufio"
)

func main() {  
    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    max := 0
    curr := 0
    for scanner.Scan() {
        if scanner.Text() == "" {
            if curr > max {
                max = curr
            }
            curr = 0
            continue
        }

        val, err := strconv.Atoi(scanner.Text())        
        if err != nil {
            log.Fatal(err)
        }
        curr += val
    }

    fmt.Printf("Answer: %v\n", max)
}
