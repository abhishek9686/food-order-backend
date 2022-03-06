.PHONY: cmd all dockerImg clean

all: 
	go mod tidy
	${MAKE} build-bin dockerImg clean
build-bin:
	GOOS=linux CGO_ENABLED=0 go build -o foodOrder


dockerImg:
	docker build -f ./build/Dockerfile . -t abhi9686/food-order-bd:latest

watch:
	reflex -s -r '\.go$$' make run

run-postgres-dev:
	docker run -d --name postgres-dev -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=root \
	-e POSTGRES_DB=food-order-app postgres
clean:
	go clean
	@rm -f ./foodOrder	