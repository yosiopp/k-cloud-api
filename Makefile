BIN := csv2json
.PHONY: all test clean build run

all: clean test build run

test:
	go test cmd/csv2json/main.go

build:
	go build -o $(BIN) cmd/csv2json/main.go
	chmod +x $(BIN)

clean:
	rm $(BIN)

run:
	scripts/run.sh