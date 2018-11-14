
package main

import (
    "fmt"
    "net"
)

func RblCheckOk(domain string) (bool, error) {
    // Check if domain is in RBL
    lookup := fmt.Sprintf("%s.dbl.spamhaus.org", domain)
    result, err := net.LookupHost(lookup)
    //fmt.Printf(">> %s %s\n", lookup, result)
    if err != nil {
        //fmt.Println("WARNING lookup for ", domain, " failed : ", err)
        return true, err
    }
    if len(result) < 1 {
        return true, nil
    }
    return false, nil
}

func RblWorker(worker_id int, domains <-chan string, results chan<- bool) {
    fmt.Printf("Worker #%d started\n", worker_id)
    for domain := range domains {
        //fmt.Printf("Worker #%d testing %s\n", worker_id, domain)
        result, err := RblCheckOk(domain)
        _ = err
        if ! result {
            fmt.Println(domain)
        }
        results <- result
        //fmt.Printf("Worker #%d finished\n", worker_id)
    }
}
