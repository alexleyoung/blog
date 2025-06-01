all: run

run: gen
	go run main.go

gen: $(wildcard components/*.templ)
	go tool templ generate
