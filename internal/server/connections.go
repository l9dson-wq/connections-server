package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func SimulateConnections(ids []string, url string) {
    //fmt.Println("SimulateConnections is being called")
    var wg sync.WaitGroup

    // setting a response time for the server
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    // send the requests using a for loop
    for _, id := range ids {
        wg.Add(1)

        // using go rutines
        go func(id string) {
            defer wg.Done()

            // The server URL 
            requestURL := fmt.Sprintf("%s?id=%s", url, id)

            // make the request to the server
            resp, err := client.Get(requestURL)
            if err != nil {
                fmt.Printf("Error sending the HTPP request for the ID %s: %v\n", id, err)
                return
            }
            defer resp.Body.Close()

            fmt.Printf("Response for the ID %s: %v\n", id, resp.StatusCode) 

            // Simulate disconnections of Each user after a period of time
            time.Sleep(40 * time.Second)
            fmt.Println("User disconnected after the 40SEG")
        }(id)
    }

    // wait until all wg group count is zero (0)
    wg.Wait()
}
