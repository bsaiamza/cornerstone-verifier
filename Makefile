# docker commands
build_docker:
	docker build -t cornerstone-verifier:latest .
	docker tag cornerstone-verifier:latest 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-verifier:latest

push_docker:
	docker push 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-verifier:latest

# golang commands
fmt:
	go fmt ./...

lint: 
	golint ./...

test:
	go test -v -cover ./...

.PHONY: build_docker push_docker fmt lint test