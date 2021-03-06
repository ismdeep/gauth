install:
	go mod tidy

test:
	go test ./... -count=1

lint:
	golint -set_exit_status ./...
	gosec -exclude=G505 -fmt golint ./...