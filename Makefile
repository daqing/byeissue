all: tailwind
	go run main.go

tailwind:
	npx tailwindcss -i ./app/assets/tailwind.css -o ./app/static/css/tailwind.css --watch
