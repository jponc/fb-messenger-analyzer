build:
	env go build -ldflags="-s -w" -o bin/fb-messenger-analyzer cmd/fb-messenger-analyzer/*.go

generate: build
	./bin/fb-messenger-analyzer -in-dir ./files -out-csv out.csv


