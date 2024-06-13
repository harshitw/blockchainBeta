build:
	go build -o ./bin/blockchainBeta
run : build
	./bin/blockchainBeta
test : 
	go test -v ./...