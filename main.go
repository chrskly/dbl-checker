
package main

import (
    "os"
    "fmt"
    "bufio"
)

func main () {

    // We'll send a series of domains into the domains channel for lookup
    domains := make(chan string, 300000)
    // Results channel to indicate when worker is done
    results := make(chan bool, 300000)

    // Start up some workers
    for w := 1; w <= 50; w++ {
        go RblWorker(w, domains, results)
    }

    // Take domains from stdin and push them into the domains channel
    scanner := bufio.NewScanner(os.Stdin)
    items := 0
    for scanner.Scan() {
        domains <- scanner.Text()
        items = items + 1
    }
    close(domains)

    // Read back the results
    bad_domains := 0
    for i := 0; i < items; i++ {
        if ! <-results {
            bad_domains = bad_domains + 1
        }
    }
    fmt.Printf("%d%% (%d/%d)\n", (100*bad_domains/items), bad_domains, items)

}
