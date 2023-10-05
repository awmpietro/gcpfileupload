# gcpfileupload

gcpfileupload is a Go package that simplifies the process of uploading a file to a bucket in Google Cloud Storage.

## Installation

To use this package, install it using go get:

```bash
go get github.com/awmpietro/gcpfileupload
```

## Usage:

First, import the package:

```go
import "github.com/awmpietro/gcpfileupload"
```

Create an instance of FileUploader:

```go
/**
/* The params for NewFileUploader are:
/* r *http.Request - pointer to the current http.Request
/* formFile string - name of field in form that holds the file
/* fileName string - if you want to change the name of original file. default is original name.
/* bucketName string - the name of the bucket where file is going to be store in Google Cloud Storage
*/
fileUpload := gcpfileupload.NewFileUploader(r, "file", "my-new-name", "my-bucket")
```

Call the Upload function:

```go
hasFile, err := fileUpload.Upload()
```

The Upload function returns:

A boolean hasFile: true if there is a file to upload.
false if not.

An error:
Returns nil if there is no error.
Returns the error encountered during the upload process otherwise.

##Configuration
Make sure to set up the GOOGLE_STORAGE_BUCKET_NAME environment variable in your environment. This should point to the path of the Google user account key you've previously created:

```yaml
export GOOGLE_STORAGE_BUCKET_NAME=path_to_your_google_account_key.json
```

Replace path_to_your_google_account_key.json with the actual path to your Google user account key file.

##Contributing
Feel free to open issues or PRs if you find any problems or have suggestions!
