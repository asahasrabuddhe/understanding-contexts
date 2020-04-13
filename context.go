package context

import "time"

// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
	// START DONE OMIT
	// Done returns a channel that is closed when this Context is canceled
	// or times out.
	Done() <-chan struct{} // HLDone
	// END DONE OMIT
	// START ERR OMIT
	// Err indicates why this context was canceled, after the Done channel
	// is closed.
	Err() error // HLErr
	// END ERR OMIT

	// START DEADLINE OMIT
	// Deadline returns the time when this Context will be canceled, if any.
	Deadline() (deadline time.Time, ok bool) // HLDeadline
	// END DEADLINE OMIT

	// START VALUE OMIT
	// Value returns the value associated with key or nil if none.
	Value(key interface{}) interface{} // HLValue
	// END VALUE OMIT
}
