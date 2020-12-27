protoc pkg/upload/uploadpb/upload.proto --go_out=plugins=grpc:.

protoc pkg/download/downloadpb/download.proto --go_out=plugins=grpc:.

protoc pkg/auth/authpb/auth.proto --go_out=plugins=grpc:.

protoc pkg/scanner/scannerpb/scanner.proto --go_out=plugins=grpc:.