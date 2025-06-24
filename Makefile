run:
	go run cmd/main.go

dockerbuild:
	docker build . -t microblog  

dockerrun:
	docker run -dp 8080:8080 microblog