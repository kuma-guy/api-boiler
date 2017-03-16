package store

type Store interface {
	// Migrate migrates the database schema to the latest available version.
	Migrate() error
	// Ping checks the contnectivity with the store.
	Ping() error
}
