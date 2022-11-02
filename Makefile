proto:
	protoc adapter/pkg/**/pb/*.proto --go_out=plugins=grpc:.
