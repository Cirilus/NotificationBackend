swagger:
	swag init -g server/app.go --output docs/
create_migration:
	goose -dir ./migrations/ create $(name) sql
