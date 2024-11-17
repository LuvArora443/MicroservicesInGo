# .DEFAULT_GOAL := swagger

# install_swagger:
# 	go get -u github.com/go-swagger/go-swagger/cmd/swagger

# swagger:
# 	@echo Ensure you have the swagger CLI or this command will fail.
# 	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
# 	@echo ....

# 	swagger generate spec -o ./swagger.yaml --scan-models

.DEFAULT_GOAL := swagger

# Install Swagger CLI
install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

# Generate Swagger specification
swagger:
	@echo "Ensure you have the swagger CLI or this command will fail."
	@echo "You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger"
	@echo "Generating Swagger spec..."
	swagger generate spec -o ./swagger.yaml --scan-models

# Generate Swagger client from the specification
swagger-client:
	@echo "Generating Swagger client from the spec..."
	swagger generate client -f ./swagger.yaml -A product-api
