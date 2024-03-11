GO_VERSION=1.22

make:
	go mod tidy
	go build -v -o bin/api.exe
	# docker compose up -d
dev:
	air
run:
	./bin/api.exe
	# docker compose up -d
clean:
	go clean
	rm -rf bin
	docker compose down