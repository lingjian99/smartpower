diagram_api:
	cd ediagram && goctl api format --dir=./api && \
	goctl api go -api ./api/diagram.api -home ../deploy/goctl/1.5.3 -dir .  -style goZero

device_rpc:

	goctl rpc protoc ./proto/device.proto --go_out=./ --go-grpc_out=./ --zrpc_out=./ -m
	