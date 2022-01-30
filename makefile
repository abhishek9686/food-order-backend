.PHONY: cmd all dockerImg clean

all: 
	go mod tidy
	${MAKE} cmd dockerImg clean

cmd:
	make -C cmd all
dockerImg:
	docker build -f ./deployments/Dockerfile . -t abhi9686/foodorder:v1
clean:
	go clean
	@rm -f ./cmd/foodOrder	