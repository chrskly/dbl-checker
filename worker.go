
package main

import (
    "fmt"
    "net"
)

// Type to hold possible set of human-friendly DBL lookup results

type DblStatusHuman string

const (
    NOT_LISTED DblStatusHuman = "Not listed"
    SPAM_DOMAIN DblStatusHuman = "Spam domain"
    PHISH_DOMAIN DblStatusHuman = "Phish domain"
    MALWARE_DOMAIN DblStatusHuman = "Malware domain"
    BOTNET_DOMAIN DblStatusHuman = "Botbet C&C domain"
    ABUSED_LEGIT_SPAM DblStatusHuman = "Abused legit spam"
    ABUSED_SPAMMED_REDIRECTOR DblStatusHuman = "Abused spammed redirector domain"
    ABUSED_LEGIT_PHISH DblStatusHuman = "Abused legit phish"
    ABUSED_LEGIT_MALWARE DblStatusHuman = "Abused legit malware"
    ABUSED_LEGIT_BOTNET DblStatusHuman = "Abused legit botnet C&C"
    IP_QUERIES_PROHIBITED DblStatusHuman = "IP queries prohibited!"
)

// Type to hold a single DBL lookup result

type DblCheckResult struct {
    Domain string
    OK bool
    StatusIP net.IP
    StatusHuman DblStatusHuman
}

func (r *DblCheckResult) SetDomain(domain string) {
    r.Domain = domain
}

func (r *DblCheckResult) SetOK(ok bool) {
    r.OK = ok
}

func (r *DblCheckResult) SetStatusIP(IP net.IP) {
    r.StatusIP = IP
}

func (r *DblCheckResult) SetStatusHuman(status DblStatusHuman) {
    r.StatusHuman = status
}


// Check if a domain is in spamhaus DBL

func DblCheck(domain string) (DblCheckResult, error) {
    lookup := fmt.Sprintf("%s.dbl.spamhaus.org", domain)
    DblResult := DblCheckResult{}
    DblResult.SetDomain(domain)
    // Do DNS lookup
    result, err := net.LookupHost(lookup)
    if err != nil || len(result) < 1 {
        // FIXME verify that we're getting NXDOMAIN
        DblResult.SetOK(true)
        DblResult.SetStatusHuman(NOT_LISTED)
        return DblResult, err
    }
    // We have a spammy domain
    DblResult.SetOK(false)
    ip := net.ParseIP(result[0])
    DblResult.SetStatusIP(ip)
    switch result[0] {
    case "127.0.1.2":
        DblResult.SetStatusHuman(SPAM_DOMAIN)
    case "127.0.1.4":
        DblResult.SetStatusHuman(PHISH_DOMAIN)
    case "127.0.1.5":
        DblResult.SetStatusHuman(MALWARE_DOMAIN)
    case "127.0.1.6":
        DblResult.SetStatusHuman(BOTNET_DOMAIN)
    case "127.0.1.102":
        DblResult.SetStatusHuman(ABUSED_LEGIT_SPAM)
    case "127.0.1.103":
        DblResult.SetStatusHuman(ABUSED_SPAMMED_REDIRECTOR)
    case "127.0.1.104":
        DblResult.SetStatusHuman(ABUSED_LEGIT_PHISH)
    case "127.0.1.105":
        DblResult.SetStatusHuman(ABUSED_LEGIT_MALWARE)
    case "127.0.1.106":
        DblResult.SetStatusHuman(ABUSED_LEGIT_BOTNET)
    case "127.0.1.255":
        DblResult.SetStatusHuman(IP_QUERIES_PROHIBITED)
    }
    return DblResult, nil
}

func DblWorker(worker_id int, domains <-chan string, results chan<- DblCheckResult) {
    // Worker function. Reads in domain names from 'domains' channel. Passes
    // back bool to 'results' channel for each domain test.
    for domain := range domains {
        result, err := DblCheck(domain)
        _ = err
        //if result.StatusHuman != NOT_LISTED {
        //    fmt.Println(domain)
        //}
        results <- result
    }
}
