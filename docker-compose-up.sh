#!/bin/bash
if [ -z "$AWS_PROFILE" ]; then
  echo "AWS_PROFILE is not set"
  exit 1
fi

docker compose down

# Store aws-vault credentials in temporary environment variables
eval $(aws-vault exec "$AWS_PROFILE" -- env | grep AWS_ | sed 's/^/export /')

AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
AWS_SESSION_TOKEN=$AWS_SESSION_TOKEN \
AWS_REGION=${AWS_REGION:-ap-northeast-1} \
docker compose up
