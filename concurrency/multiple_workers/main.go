package multipleworkers

// Create 3 worker goroutines that each receive numbers from a shared input channel, square them, and send results to
// a shared output channel. The main goroutine should send numbers 1-9 and collect all results.
