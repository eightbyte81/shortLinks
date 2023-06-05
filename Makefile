migrateup:
	migrate -path ./schema -database 'postgres://me:password@localhost:5432/short_link_db?sslmode=disable' up

migratedown:
	migrate -path ./schema -database 'postgres://me:password@localhost:5432/short_link_db?sslmode=disable' down