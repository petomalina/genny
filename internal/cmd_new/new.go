package cmd_new

import "fmt"

func DefaultMakefileContents(projectName string) string {
	return fmt.Sprintf(`##             %s microservice Makefile
##
##  Simple Makefile containing implementation of targets for generating protobuf file
##
##    $ make gen
SHELL=/bin/bash

PROTO_INCLUDES=-I=proto -I=. $(shell cat protomodules | tr '\n' ' ')
PROTO_APIS=$(shell cat protoapis | tr '\n' ' ')
PROTO_GRPC_ARGS=paths=source_relative
##
##  \e[1mTargets\e[0m
##   \e[34mhelp\e[0m
##       Shows this help
help:
	@echo -e "$$(sed -n 's/^##//p' Makefile)"

##   \e[34mgen\e[0m
##       Shortcut for generate
gen:
	make generate SERVICE_NAME=user

##   \e[34mgenerate\e[0m
##       Generates Go
generate: generate/go

##   \e[34mgenerate/go\e[0m
##       Generates go grpc files and messages from proto file
generate/go:
	protoc $(PROTO_INCLUDES) \
		  proto/v1/${SERVICE_NAME}.proto \
		   --go_out=$(PROTO_GRPC_ARGS):go-sdk/${SERVICE_NAME} \
		   --go-grpc_out=$(PROTO_GRPC_ARGS):go-sdk/${SERVICE_NAME} \
		   --validate_out="lang=go,$(PROTO_GRPC_ARGS):go-sdk/${SERVICE_NAME}"
`, projectName)

}

func DefaultGitignoreContents() string {
	return `.terraform
.idea
vendor
*.tfplan`
}
