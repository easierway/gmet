package gmet

type MetWriter interface {
	// write the formatted metrics
	Write(msg string)
	// flush out the data from cache
	Flush()
	Close() error
}

type MetFormatter interface {
	Format(registry Registry) (string, error)
}
