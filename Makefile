setup: 
	docker-compose up -d

setup-down:
	docker-compose down

build:
	go build -o mongoparser main.go


parse: build
	./mongoparser parse
