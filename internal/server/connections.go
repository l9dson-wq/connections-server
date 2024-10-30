package server

import (
	"fmt"
    _"log"
	"net/http"
	"sync"
)

func SimulateConnecionts(ids []string, url string) {
    var wg sync.WaitGroup

    // send the requests using a for loop (probably this can be improve)
    for _, id := range ids {
        wg.Add(1)

        // using go rutines
        go func(id string) {
            defer wg.Done()

            // My server URL which for now is a localhost
            requestURL := fmt.Sprintf("%s%s", url, id)

            // make the request to the server
            resp, err := http.Get(requestURL)
            if err != nil {
                fmt.Printf("Error sending the HTPP request for the ID %s: %v\n", id, err)
                return
            }
            defer resp.Body.Close()

            fmt.Printf("Response for the ID %s: %v\n", id, resp.StatusCode)
        }(id)
    }

    // wait until all wg group count is zero (0)
    wg.Wait()
}
