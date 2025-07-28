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

### AWS Credentials
Configure AWS credentials using one of these methods:
- Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
- Shared credentials file (`~/.aws/credentials`)
- AWS SSO (via `aws sso login` or `aws-vault` etc.)
- IAM role (EC2 instance profile or web identity)

### AWS Region Configuration
You must configure an AWS region. Choose one of these methods:

1. **Set default region in AWS config file** (Recommended):
   ```sh
   # Add to ~/.aws/config
   [default]
   region = ap-northeast-1
   ```

2. **Use environment variable**:
   ```sh
   export AWS_REGION=ap-northeast-1
   # or
   export AWS_DEFAULT_REGION=ap-northeast-1
   ```

3. **Use AWS profile with region**:
   ```sh
   export AWS_PROFILE=your-profile-name
   ```

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

## Troubleshooting

### Error: "Invalid region: region was not a valid DNS name"
This error occurs when AWS region is not configured. See [AWS Region Configuration](#aws-region-configuration) section above.

### Error: "failed to refresh cached credentials"
This error occurs when AWS credentials are not properly configured or have expired. Try:
- For AWS SSO: Run `aws sso login --profile <profile-name>` or use `aws-vault exec <profile-name>`
- For IAM users: Check your `~/.aws/credentials` file
- For temporary credentials: Ensure your session hasn't expired

### Error: "Access Denied"
Ensure your AWS credentials have the necessary S3 permissions:
- `s3:GetObject` for download URLs
- `s3:PutObject` for upload URLs
- `s3:DeleteObject` for delete URLs
- `s3:PutObject` for multipart upload URLs
