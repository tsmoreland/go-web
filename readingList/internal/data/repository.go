package data

type Repository interface {
	Migrate() error
	Close() error
}
