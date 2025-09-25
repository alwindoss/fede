.PHONY: clean build run

build: clean
	go build -o bin/fede

clean:
	rm -rf bin

run: build
	./bin/fede run server
