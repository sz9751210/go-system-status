build:
	@go build -o bin/system-monitor

run: build
	@./bin/system-monitor
