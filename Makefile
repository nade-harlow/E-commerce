run: |
	gofmt -w adapter core ports
	go run main.go