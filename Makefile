.PHONY: swagger_gen
swagger_gen:
	swagger generate spec -o ./swagger.yaml --scan-models

.PHONY: swagger_view
swagger_view:
	swagger serve -F=swagger swagger.yaml
