
package main

import (
    "fmt"
    "net"
)

func RblWorker(worker_id int, domains <-chan string, results chan<- string) {
    fmt.Println("Worker #%d started", worker_id)
    for domain := range domains {
        fmt.Println("Worker #%d testing %s", worker_id, domain)
        net.LookupHost("example.com.dbl.spamhaus.org")
    }
}


