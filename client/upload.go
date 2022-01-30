package main

import (
	"cloud-watcher/proto/uploadpb"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gookit/color"
)

type Uploader struct {
	c            uploadpb.FileUploadServiceClient
	file         *os.File
	name         string
	searchTags   string
	pathToObject string

	args []string
}

func MakeUploader(uc uploadpb.FileUploadServiceClient, args []string) Uploader {
	u := Uploader{
		pathToObject: readInput("Enter path to file (example: ./folder/another Folder/file.txt): "),
		// searchTags:   readInput("Enter words assosiated with file, for easier search (example: code,go,app): "),
		searchTags: "ssadasd",
		c:          uc,
		args:       args,
	}
	u.name = filepath.Base(u.pathToObject)

	return u
}

func (u Uploader) Upload() {
	file, err := os.OpenFile(u.pathToObject, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}

	u.file = file
	defer u.file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stream, err := u.c.UploadFile(ctx)
	if err != nil {
		log.Fatalf("UploadFile error: %v", err)
	}

	//send first request thtat contains file info
	if err := stream.Send(infoRequest(u.name, u.searchTags)); err != nil {
		log.Fatalf("Send error: %v", err)
	}

	bufferSize := 2 * 1024

	var bytesStack [][]byte

	var offset int64 = 0

	for {
		bytesRead, err := readChunk(u.file, bufferSize, offset, 0)

		if err == io.EOF {
			if len(bytesStack) < 1 {
				log.Fatal("File is invalid.")
			}

			trimmedBytes := trimBytes(bytesStack[0])

			stream.Send(byteRequest(trimmedBytes))

			break
		}

		if err != nil {
			log.Fatal(err)
		}

		bytesStack = append(bytesStack, bytesRead)

		if len(bytesStack) > 1 {
			if err := stream.Send(byteRequest(bytesStack[0])); err != nil {
				log.Fatalf("Send error: %v", err)
			}

			bytesStack = bytesStack[1:]
		}

		offset += int64(bufferSize)
	}

	response, err := stream.CloseAndRecv()

	if err != nil {
		color.Error.Tips("Error uploading: " + u.name)
		log.Fatal(err)
	}

	if response.Ok {
		color.Success.Tips("Successfully uploaded: " + u.name)
	}
}

func readChunk(file *os.File, chunkSize int, off int64, whence int) ([]byte, error) {
	bytes := make([]byte, chunkSize)

	file.Seek(off, whence)
	_, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func infoRequest(name, searchTags string) *uploadpb.FileUploadRequest {
	return &uploadpb.FileUploadRequest{
		Data: &uploadpb.FileUploadRequest_Info{
			Info: &uploadpb.FileUploadInfo{
				Name:       name,
				SearchTags: searchTags,
			},
		},
	}
}

func byteRequest(bytes []byte) *uploadpb.FileUploadRequest {
	return &uploadpb.FileUploadRequest{
		Data: &uploadpb.FileUploadRequest_Body{
			Body: &uploadpb.FileUploadBody{Bytes: bytes},
		},
	}
}

// func trimBytes(bytes []byte) []byte {
// 	var firstZeroByteIndex int
// 	firstZeroWasHitted := false

// 	for i, singleByte := range bytes {
// 		if singleByte == 0 && !firstZeroWasHitted {
// 			firstZeroByteIndex = i
// 			firstZeroWasHitted = true
// 		}
// 		if singleByte != 0 {
// 			firstZeroByteIndex = -1
// 			firstZeroWasHitted = false
// 		}
// 	}

// 	if firstZeroByteIndex == -1 {
// 		firstZeroByteIndex = len(bytes)
// 	}

// 	bytes = bytes[:firstZeroByteIndex]

// 	return bytes
// }

// -----------------------------------------------------------------------------------------------------------------------

// func uploadFromTerminal(uc uploadpb.FileUploadServiceClient, args []string) {
// 	path := readInput("Enter path to file (example: ./folder/another Folder/file.txt): ")

// 	for _, arg := range args {
// 		if arg == "-f" {
// 			uploadFolderFromTerminal(uc, path)
// 			return
// 		}
// 	}

// 	uploadFileFromTerminal(uc, path)
// }

// func uploadFolderFromTerminal(uc uploadpb.FileUploadServiceClient, path string) {
// 	var filePaths []string

// 	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
// 		if !info.IsDir() {
// 			filePaths = append(filePaths, path)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		log.Println("upload reading files from folder: ", err)
// 		return
// 	}

// 	uploadMany(uc, filePaths)
// }

// func uploadFileFromTerminal(uc uploadpb.FileUploadServiceClient, path string) {
// 	upload := newUpload(uc)

// 	ctx := context.Background()

// 	if err := upload.openAndPrepareFile(ctx, path); err != nil {
// 		log.Fatalf("unable to open file %v", err)
// 	}

// 	if err := upload.handle(ctx); err != nil {
// 		log.Fatal("unable to handle file streaming")
// 	}
// }

// func newUpload(client uploadpb.FileUploadServiceClient) *upload {
// 	return &upload{c: client}
// }

// func uploadMany(uc uploadpb.FileUploadServiceClient, filePaths []string) {
// 	var mu sync.Mutex

// 	for _, filePathF := range filePaths {
// 		mu.Lock()
// 		filePath := filePathF

// 		go func() {
// 			ctx := context.Background()

// 			upload := newUpload(uc)

// 			if err := upload.openAndPrepareFile(ctx, filePath); err != nil {
// 				log.Fatalf("unable to open file %v", err)
// 			}

// 			if err := upload.handle(ctx); err != nil {
// 				log.Fatal("unable to handle file streaming")
// 			}

// 			mu.Unlock()
// 		}()
// 	}
// }

// type upload struct {
// 	c        uploadpb.FileUploadServiceClient
// 	file     *os.File
// 	fileName string
// }

// func (u *upload) openAndPrepareFile(ctx context.Context, fullPath string) error {
// 	if err := u.openFile(ctx, fullPath); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // func (u *upload) getFileData(ctx context.Context, fullPath string) error {
// // 	pathSliceNotFiltered := strings.Split(fullPath, "/")
// // 	if len(pathSliceNotFiltered) < 1 {
// // 		return errInvalidPath
// // 	}

// // 	var pathSlice []string

// // 	for _, name := range pathSliceNotFiltered {
// // 		if name == "." {
// // 			continue
// // 		}

// // 		pathSlice = append(pathSlice, name)
// // 	}

// // 	nameSlice := strings.Split(pathSlice[(len(pathSlice)-1)], ".")
// // 	if len(nameSlice) < 2 {
// // 		return errInvalidPath
// // 	}

// // 	u.fileName = nameSlice[0]
// // 	u.fileExtension = nameSlice[1]
// // 	u.path = strings.Join(pathSlice[:(len(pathSlice)-1)], "/")

// // 	return nil
// // }

// func (u *upload) openFile(ctx context.Context, fullPath string) error {
// 	file, err := os.OpenFile(mainDirPath+"/"+fullPath, os.O_RDONLY, 0777)
// 	if err != nil {
// 		return err
// 	}

// 	u.file = file

// 	return nil
// }

// func (u *upload) handle(ctx context.Context) error {
// 	defer u.file.Close()

// 	stream, err := u.c.UploadFile(ctx)
// 	if err != nil {
// 		fmt.Printf("UploadFile error1: %v", err)
// 		return err
// 	}

// 	stream.Send(&uploadpb.FileUploadRequest{
// 		Data: &uploadpb.FileUploadRequest_Info{
// 			Info: &uploadpb.FileUploadInfo{
// 				Name:       u.fileName,
// 				SearchTags: u.searchTags,
// 			},
// 		},
// 	})

// 	var bytesStack [][]byte

// 	for {
// 		bytes, err := u.readFile(ctx)
// 		if err == io.EOF {
// 			if len(bytesStack) < 1 {
// 				log.Println("file is invalid: ", u.fileName+"."+u.fileExtension)
// 				break
// 			}

// 			trimmedBytes := trimBytes(bytesStack[0])

// 			req := byteRequest(trimmedBytes)

// 			stream.Send(req)

// 			break
// 		}

// 		if err != nil {
// 			fmt.Printf("UploadFile error2: %v", err)
// 			return err
// 		}

// 		bytesStack = append(bytesStack, bytes)

// 		if len(bytesStack) > 1 {
// 			req := byteRequest(ctx, bytesStack[0])

// 			stream.Send(req)

// 			bytesStack = bytesStack[1:]
// 		}
// 	}
// }

// 	if err != nil {
// 		color.Error.Tips("Error uploading: " + u.fileName + "." + u.fileExtension)
// 		return err
// 	}

// 	if response.Ok {
// 		color.Success.Tips("Successfully uploaded: " + u.fileName + "." + u.fileExtension)
// 	}

// 	return nil
// }

// func (u *upload) infoRequest(ctx context.Context) *uploadpb.FileUploadRequest {
// 	return &uploadpb.FileUploadRequest{
// 		Data: &uploadpb.FileUploadRequest_Info{
// 			Info: &uploadpb.FileUploadInfo{
// 				Name:             u.fileName,
// 				Extension:        u.fileExtension,
// 				Path:             u.path,
// 				ToPersonalFolder: true,
// 				ExternalFolderID: "",
// 			},
// 		},
// 	}
// }

// func byteRequest(ctx context.Context, bytes []byte) *uploadpb.FileUploadRequest {
// 	return &uploadpb.FileUploadRequest{
// 		Data: &uploadpb.FileUploadRequest_Body{
// 			Body: &uploadpb.FileUploadBody{Bytes: bytes},
// 		},
// 	}
// }

// func (u *upload) readFile(ctx context.Context) ([]byte, error) {
// 	bytes := make([]byte, 4*1024)
// 	_, err := u.file.Read(bytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return bytes, nil
// }
