GO_VERSION=1.22

make:
	go build -v -o bin/api.exe
	# docker compose up -d
run:
	./bin/api.exe
	# docker compose up -d
clean:
	go clean
	rm -rf bin
	docker compose down