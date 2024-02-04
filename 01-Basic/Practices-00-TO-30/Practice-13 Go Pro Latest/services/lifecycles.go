package services

type lifecycle int

const (
	// چرخه زندگی گذرا
	Transient lifecycle = iota
	Singleton
	Scoped
)
