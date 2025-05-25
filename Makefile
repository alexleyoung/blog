all: run

run: gen
	go run main.go

gen: $(wildcard components/*.templ)
	templ generate
