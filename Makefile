proto:
	protoc api-gateway/pkg/**/pb/*.proto --go_out=plugins=grpc:.
	protoc auth-service/internal/pb/*.proto --go_out=plugins=grpc:.
	protoc product-service/internal/pb/*.proto --go_out=plugins=grpc:.
	protoc user-service/internal/pb/*.proto --go_out=plugins=grpc:.
