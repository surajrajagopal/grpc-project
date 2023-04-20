# grpc-project

Steps to build grpc-project :
1. Install Go in your system, if it linux follow below link :
   - https://golangdocs.com/install-go-linux

2. Install Protocol Buffer : follow below links 
   - https://grpc.io/docs/protoc-installation/

3. Go plugins for the protocol compiler : 
   - https://grpc.io/docs/languages/go/quickstart/
   - if your getting some error while installing follow below link :
     https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable
     https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable

4. Generating the stubs :
   - protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
     --proto_path -> this is where your proto exists.
     --go_out -> generate a struct in file (.pb.go)
     --go-grpc_out -> generate a server and client interface in file (_grpc.pb.go)

5. After generating of stubs :
   - Implement service by using API's that we declare in proto. Ex: helloservice folder in this project.

6. Implement Server side Connection. Ex: Server folder in this project.

7. Implement Client side Connection to call server side APIs. Ex: Client folder in this project.