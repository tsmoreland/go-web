package data

type Repository interface {
	Migrate() error
	Ping() error
	InsertOne(title string, published int, pages int, rating float64, genres []string) (*Book, error)
	Close() error
}
