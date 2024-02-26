REPO_URL := https://github.com/OyePuru/grpc-gateway
BRANCH := main
PROTO_FOLDER := proto
GEN_FOLDER := gen
SERVICE_YAML := service.yaml

clean:
	rm -rf proto gen 

pull-views:
	git clone --branch $(BRANCH) --single-branch --depth 1 $(REPO_URL) tmp-repo
	cp -r tmp-repo/$(PROTO_FOLDER) .
	cp -r tmp-repo/$(GEN_FOLDER) .
	rm -rf tmp-repo


generate-stubs:
	rm -rf gen
	mkdir gen
	mkdir gen/go
	protoc -I . \
  --go_out ./gen/go/ \
  --go_opt paths=source_relative \
  --go-grpc_out ./gen/go/ \
  --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./gen/go \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt grpc_api_configuration=service.yaml \
  ./proto/**/*.proto