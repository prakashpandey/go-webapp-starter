
.PHONY: build
build:
	go build -o go-sample-webapp-starter

.PHONY: run
run:
	./go-sample-webapp-starter

.PHONY:	clean
clean:
	go clean
	rm -f go-sample-webapp-starter
