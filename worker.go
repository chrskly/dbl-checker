
package main

import (
    "fmt"
    "net"
)

type DblCheckResult struct {
    Domain string
    DblStatusIP net.IPV4
    DblStatusHuman string
}

func DblCheckOk(domain string) (bool, error) {
    // Check if domain is in spamhaus RBL
    lookup := fmt.Sprintf("%s.dbl.spamhaus.org", domain)
    result, err := net.LookupHost(lookup)
    if err != nil {
        return true, err
    }
    if len(result) < 1 {
        return true, nil
    }
    return false, nil
}

func DblWorker(worker_id int, domains <-chan string, results chan<- bool) {
    // Worker function. Reads in domain names from 'domains' channel. Passes
    // back bool to 'results' channel for each domain test.
    for domain := range domains {
        result, err := DblCheckOk(domain)
        _ = err
        if ! result {
            fmt.Println(domain)
        }
        results <- result
    }
}
