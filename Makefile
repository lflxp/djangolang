run:
	cd examples && go run example.go

swag:
	swag init --parseDependency --parseDepth=1

sample:
	cd examples/sample && go run main.go