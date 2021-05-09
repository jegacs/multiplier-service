# multiplier-service

## Goal

Design a service to multiply two numbers in Go. Use both gRPC
and REST to serve the requests.


## Project structure
- The service logic is in /services directory.
- Both gRPC and HTTP server are located in /server directory.
- gRPC service and protocol buffer definitions are located in /protos directory.
- A couple of custom errors are located in /errors directory.
- The REST payload and responses are located in /dto directory.
- An example to test the gRPC mode of the server is located in /grpc-client directory (further explanation below).

## Usage

Change to the server/ directory and compile the service.

```
go build .
```
Then run the server. To run in http mode (rest), do as 
follows:
```
./server -http
```
You can test http mode using curl: 
```
curl -d '{"first":"100.01", "second":"-100"}' -H "Content-Type: application/json" -X POST http://localhost:8000/multiplier -v
```
Otherwise, run it in gRPC mode:

```
./server -gRPC
```

To test gRPC mode, go to the root of the service and change to grpc_client. Pass the two numbers to multiply as an 
argument:
```
go run main.go 0 100
```
### Run tests
In order to run the tests and print the coverage of the project, go to the root of project and run the following command:
```
go test ./... -coverprofile cover.out
```
## Limitations
- The two numbers inserted are processed as strings, being parsed with math/big and processed package.
- The precission setting used for the multiplication and parsing the numbers is 256 
- An arbitrary upper and limit for each number has been set [-1000, 1000]. These limits were chosen as arbitrary as possible in order to validate in unit tests the captabilities
of the API.
- The result of the product is truncated to two decimals and 
so it is returned.

# Further work and considerations.
I would improve this project by:
- Create an interface for the service. Maybe call it `Operation` with a method `Calculate`. This way we could 
extend this for other operations, like `Sum`, `Substract`, etc. Also, this is one of the reasons for treating the 
input and result numbers as strings.
- By doing the mentioned above, decouple the service logic a 
little bit more.
- Being more specific about the level of precision for the
operations being done.
- Think of a way to scale the service by calculating several
operations concurrently.


