.PHONY: setup_swagger
setup_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: swagger_gen
swagger_gen:
	swagger generate spec -o ./swagger.yaml --scan-models

.PHONY: swagger_serve
swagger_serve:
	swagger serve -F=swagger swagger.yaml
