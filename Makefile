
compile:
	go build -o api ./cmd/server
	go build -o worker ./cmd/kafka
