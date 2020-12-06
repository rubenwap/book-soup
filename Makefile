run:
	go run ./cmd/main.go

createdb: 
	psql -U ruben -d booksoup -h localhost  -p 5432 -f db/structure.sql

