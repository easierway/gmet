package gmet

// only for test
type DummyWriter struct {
}

func (writer *DummyWriter) Write(msg string) {
}

func (writer *DummyWriter) Flush() {
}

func (writer *DummyWriter) Close() error {
	return nil
}
