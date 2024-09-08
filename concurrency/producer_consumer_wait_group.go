package concurency

import (
        "fmt"
        "sync"
        "time"
)

// Shared buffer channel
var dataChannel = make(chan int, 10) // Buffer size of 10

// WaitGroup to synchronize goroutines
var wg sync.WaitGroup

func producer(n int) {
        defer wg.Done()
        for i := 0; i < n; i++ {
                dataChannel <- i
                fmt.Println("Producer sent:", i)
                time.Sleep(time.Millisecond * 100)
        }
}

func consumer() {
        defer wg.Done()
        for data := range dataChannel {
                fmt.Println("Consumer received:", data)
                time.Sleep(time.Millisecond * 200)
        }
}

func main() {
        // Create producer and consumer goroutines
        wg.Add(2)
        go producer(10)
        go consumer()

        // Wait for all goroutines to finish
        wg.Wait()

        // Close the channel to signal the consumer to stop
        close(dataChannel)

        fmt.Println("All done!")
}