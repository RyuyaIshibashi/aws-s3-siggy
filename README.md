# aws-s3-siggy
Generate presigned URL for AWS S3

```
      _                   
  ___(_) __ _  __ _ _   _ 
 / __| |/ _' |/ _' | | | |
 |__ | | (_| | (_| | |_| |
 |___|_|___, |___, |___, |
        |___/ |___/ |___/ 
  Version: V*.*.*-*******
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
```

## Parameters

- `-b <bucket_name>`: Required. Name of the S3 bucket.
- `-k <object_key>`: Required. S3 object key.

## Example

```sh
$ siggy -m get -b sample-bucket -k path/to/file.txt

https://sample-bucket.s3.amazonaws.com/path/to/file.txt?X-Amz-...  
```
