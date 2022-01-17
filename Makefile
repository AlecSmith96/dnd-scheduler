.PHONY: setup_swagger
setup_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: swagger_gen
swagger_gen:
	swagger generate spec -o ./docs/swagger.yaml --scan-models

.PHONY: swagger_serve
swagger_serve:
	swagger serve --no-open -p 5000 -F=swagger ./docs/swagger.yaml
