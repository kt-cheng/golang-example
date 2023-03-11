# golang_example

### Go module
- Run `go mod init <module_name>` to generate `go.mod` file.
- Run `go get -u <package_name>` to install packages, related information is recorded in `go.sum` file.

### Build image
docker build -t golang_example .

### Run Golang docker container
docker run -d -p 8080:8080 golang_example

### Reference
- https://docs.docker.com/language/golang/build-images/
- https://ithelp.ithome.com.tw/articles/10233671

### Check response
Run `curl http://localhost:8080/`