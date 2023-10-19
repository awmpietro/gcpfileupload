package gcpfileupload

import (
	"context"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
)

type FileUploader struct {
	Request    *http.Request
	FormFile   string
	FileName   string
	BucketName string
}

func NewFileUploader(r *http.Request, formFile string, fileName string, bucketName string) *FileUploader {
	return &FileUploader{
		Request:    r,
		FormFile:   formFile,
		FileName:   fileName,
		BucketName: bucketName,
	}
}

// Upload a file to a bucket in Google Cloud Storage
// Make sure you have defined the env var GOOGLE_STORAGE_BUCKET_NAME in your env file
// The env should point to the path of your google user account keys
func (fu *FileUploader) Upload(errCh chan error) {
	go func(errCh chan<- error) {
		file, header, err := fu.Request.FormFile(fu.FormFile)
		if err != nil {
			errCh <- err
			return
		}
		defer file.Close()
		if file != nil {

			if fu.FileName == "" {
				fu.FileName = header.Filename
			} else {
				ext := filepath.Ext(header.Filename)
				fu.FileName = fu.FileName + ext
			}

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, time.Second*50)
			defer cancel()

			client, err := storage.NewClient(ctx)
			if err != nil {

				errCh <- err
				return
			}
			bucket := client.Bucket(fu.BucketName)
			obj := bucket.Object(fu.FileName)
			wc := obj.NewWriter(ctx)
			if _, err := io.Copy(wc, file); err != nil {

				errCh <- err
				return
			}
			if err := wc.Close(); err != nil {
				errCh <- err
				return
			}
			close(errCh)
		}
	}(errCh)

}
