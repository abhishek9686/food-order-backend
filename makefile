.PHONY: cmd all dockerImg clean

all: 
	go mod tidy
	${MAKE} build dockerImg clean
build:
	GOOS=linux CGO_ENABLED=0 go build -o foodOrder


dockerImg:
	docker build -f ./deployments/Dockerfile . -t abhi9686/foodorder:v1

watch:
	reflex -s -r '\.go$$' make run

clean:
	go clean
	@rm -f ./foodOrder	