func countUp(max int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 1; i <= max; i++ {
            out <- i
        }
    }()
    return out
}
