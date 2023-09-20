perpare:
	go install github.com/cosmtrek/air@latest

run-web:
	go run .


run-air:
	air -c .air.toml

run-dkc:
	docker-compose up -d	
