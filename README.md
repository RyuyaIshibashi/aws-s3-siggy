# aws-s3-siggy
Generate presigned URLs for AWS S3 operations (upload, download, delete, multipart upload).


```
      _                   
  ___(_) __ _  __ _ _   _ 
 / __| |/ _' |/ _' | | | |
 |__ | | (_| | (_| | |_| |
 |___|_|___, |___, |___, |
        |___/ |___/ |___/ 
  Version: V*.*.*-*******
```

aws-s3-siggy is a command-line tool written in Go that generates temporary pre-signed URLs for AWS S3 objects.
It helps you securely grant time-limited access without exposing your AWS credentials.

## Prerequisites

- AWS credentials configured via one of the following:
  - Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
  - Shared credentials file (`~/.aws/credentials`)
  - AWS SSO (via AWS CLI v2)
  - IAM role (instance profile or web identity)

## Installation

```sh
brew install RyuyaIshibashi/aws-s3-siggy/siggy
```

## Usage

```sh
# Generate upload URL (PutObject)
siggy put -b <bucket_name> -k <object_key>

# Generate download URL (GetObject)
siggy get -b <bucket_name> -k <object_key>

# Generate delete URL (DeleteObject)
siggy delete -b <bucket_name> -k <object_key>

# Generate upload part URL for multipart upload (UploadPart)
siggy upload_part -b <bucket_name> -k <object_key> -u <upload_id> -p <part_number>
```

## Parameters

- `-b <bucket_name>`: Required. Name of the S3 bucket.
- `-k <object_key>`: Required. S3 object key.
- `-u <upload_id>`: Required for upload_part. The ID of the multipart upload.
- `-p <part_number>`: Required for upload_part. The part number of the multipart upload.

## Example

```sh
$ siggy get -b sample-bucket -k path/to/file.txt

https://sample-bucket.s3.amazonaws.com/path/to/file.txt?X-Amz-...
```
