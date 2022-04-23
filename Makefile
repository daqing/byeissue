all: tailwind

tailwind:
	npx tailwindcss -i ./app/assets/tailwind.css -o ./app/static/css/tailwind.css --watch

run:
	go run ./main.go
