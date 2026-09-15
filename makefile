.PHONY: dev prod

dev:
	APP_ENV=development go run main.go

prod:
	APP_ENV=production go run main.go
