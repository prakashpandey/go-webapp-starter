
.PHONY: build
build:
	go build -o go-sample-webapp-starter

.PHONY: run
run:
	./go-sample-webapp-starter

.PHONY:	clean
clean:
	rm go-sample-webapp-starter
