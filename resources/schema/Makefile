BASE_GET_CMD := go run get.go
SWAGGER_PKG := github.com/go-swagger/go-swagger/cmd/swagger
SWAGGER_CMD := swagger
OPENAPI_GENERATOR := github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
OPENAPI_CMD := oapi-codegen
PACKAGE := ""
SCHEMA_FILE := ""
OUTPUT_FILE := ""
OUTDIR := $(dir $(OUTPUT_FILE))

SNIPE_VERSION := 6.1.2

.PHONY: swagger snipe

.create-output:
	mkdir -p $(OUTDIR)
	touch $(OUTPUT_FILE)

.get-swagger:
	go get $(SWAGGER_PKG)

install_swagger: .get-swagger
	go install $(SWAGGER_PKG)

.get-openapi:
	go get $(OPENAPI_GENERATOR)

install_openapi: .get-openapi
	go install $(OPENAPI_GENERATOR)

.generate-openapi:
	$(OPENAPI_CMD) -generate client,types -package $(PACKAGE) $(SCHEMA_FILE) > $(OUTPUT_FILE)

.jamf-tidy:
	mv jamf_swagger.yaml jamf/swagger.yaml

.jamf-run:
	$(BASE_GET_CMD) -jamf

.jamf-gen: 
	$(SWAGGER_CMD) generate client -f jamf/swagger.yaml -A jamf -t jamf

jamf: .jamf-run .jamf-tidy .jamf-gen

.snipe-tidy:
	mv snipe-*.json snipe/swagger.json

.snipe-get:
	$(BASE_GET_CMD) -schema $(SNIPE_VERSION)
	
.snipe-interactive-get:
	$(BASE_GET_CMD)

.snipe-generate: .generate-openapi

snipe: .snipe-get .snipe-tidy .create-output .snipe-generate

snipe-interactive: .snipe-interactive-get .snipe-tidy .create-output .snipe-generate
