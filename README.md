To run this project please replace actual values for postgres database

```db, err = gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=test password=postgres sslmode=disable",
	)
```

Please install all dependencies with ```go mod tidy```\
Run the project with ```go run .```\
Run all test cases ```go test -v```
