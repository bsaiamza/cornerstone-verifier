# docker commands
build_docker:
	docker build -t cornerstone-verifier:0.0.1 .
	docker tag cornerstone-verifier:0.0.1 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-verifier:0.0.1

push_docker:
	docker push 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-verifier:0.0.1

# golang commands
fmt:
	go fmt ./...

lint: 
	golint ./...

test:
	go test -v -cover ./...

.PHONY: build_docker push_docker fmt lint test