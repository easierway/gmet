package gmet

type GMet interface {
	Metric(key string)
	// TODO: add stats
	Flush()
}

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
