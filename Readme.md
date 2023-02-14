protoc -I ./proto \ #A
--go_out ./golang \ #B
--go_opt paths=source_relative \ #C
--go-grpc_out ./golang \ #D
--go-grpc_opt paths=source_relative \ #E
./proto/order.proto #F

-I is used to specify import path where imported packages in proto files are searched.
--go_out to specify where to put generated go code for messages --go_opt is used for configuring options for go souce code generation like paths=source_relative to keep same folder structure after source code generation.
--go-grpc_out to define destination folder of gRPC specific go source code like calling a service function.
--go-grpc_opt is used to configure options for gRPC related operations like paths=source_relative to have same folder structure after source code generation.


protoc -I ./org --go_out ./golang/org --go_opt paths=source_relative --go-grpc_out ./golang/org --go-grpc_opt paths=source_relative ./org/org.proto

protoc -I ./role --go_out ./golang/role --go_opt paths=source_relative --go-grpc_out ./golang/role --go-grpc_opt paths=source_relative ./role/role.proto


git tag -a golang/org/v1.2.3 -m “golang/org/v1.2.3”
git tag -a golang/role/v1.2.8 -m “golang/role/v1.2.8”
git push --tags