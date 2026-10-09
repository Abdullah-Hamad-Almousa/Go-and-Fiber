# day6

This repo is not useful. It is a scratch pad of Go snippets with no app, no library, and no problem to solve. `master` is whatever `main.go` happens to do right now. The other branches are placeholders: names and notes only. They are not created yet.

## day6

## placeholder branches

## defer_example.go
it creates a file and write a text on it
after writing the text defer activate and delete the file

<br>

## chan_example.go
pipeline := make(chan int, 3) creating a channel of capacity 3
sends 10,20 and 30; into the buffer. These sends do not block bc the buffer has room for all three values
close(pipeline) closes the channel to signal that no more values.

<br>

## arrow_example.go
ch := make(chan string, 1) create a buffer string channel with capacity 1
data := <- in it's a parameter <-chan string can only receive.

<br>

## go_example.go

RunGoExample starts the background worker as a goroutine, then continues in the main.
then sleeps for 100ms
then it sleeps for 150ms giving the worker time to finish before RunGoExample returns.

<br>

## sync_example.go

Creates a sync.WaitGroup.
Loops three times; Each iteration increments, adds one to the wait group, then starts a worker goroutine.
The worker IDs: 1, 2, 3
Each worker prints its message, then defer wg.Done() reduces the wait-group count when that worker returns.
wg.Wait() blocks until all three workers have called Done().

<br>

## select_example.go

Creates FastCh, an unbuffered string channel.
Starts a goroutine that sleeps for 50ms, then tries to send "Fast server response" on fastCh.
Since the channel is unbuffered, that send waits until something receives the value.
Enters select, waiting for either
    - a value from fastCh or
    - the 200ms timeout from time.After
The goroutine's send is ready first, select receives the response and runs the first case.

<br>

| branch | what I did |
| --- | --- |
| `master` | Go Setup & Syntax *T and &T. |
| `day2` | for loop and range. make(), append(), delete(). |
| `day3` | Structs & Methods type and pointer |
| `day4` | type Speaker interface, s.(type) for Interfaces & Polymorphism |
| `day5` | Error Handling & Tests |
| `day6` | Concurrency Essentials |