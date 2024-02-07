package eventsourcing

type (
	Versioner interface {
		GetVersion() int
		GetPendingVersion() int
		setVersion(version int)
	}
)
