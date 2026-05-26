package testutil

// LogWriter implementation of `io.Writer` redirecting the logs to `testing.T.Log()`
type LogWriter struct {
	T interface {
		Log(args ...any)
	}
}

func (w LogWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
