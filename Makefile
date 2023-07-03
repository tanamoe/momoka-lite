prepare:
	go mod download
	go mod verify

serve:
	go run ./cmd/serve.go serve

fmt:
	find -name "*.go" -exec go fmt {} \;
