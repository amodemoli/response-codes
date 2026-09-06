## ⭓ Response Codes - Ready-Made Status Codes
Ready-made code for responding to clients, specifically for the server side in microservice and simple websites. Especially in microservices, you can equip all your services with this system using a simple script. Now you can broadcast your StatusCode and ResponseCode across all your services and update them whenever you want, instead of changing them manually and inconsistently, you just update them publicly for all services with the script you wrote. This way, services can communicate with each other in a common language and send and receive request statuses.
> ❐ It's suggested that you fork this project and make any changes you need yourself because this project is built based on my own architecture and personal use, and it might be a bit different from your projects.

### ❝ A sample Linux script that you can put on your service so it can take advantage of this feature:
```bash
#!/bin/bash

BASE_URL="https://raw.githubusercontent.com/amodemoli/response-codes/main"
RESPONSE_FILE="response.codes.go"
STATUS_FILE="status.codes.go"

RESPONSE_DIR="../internal/helpers/codes/response"
STATUS_DIR="../internal/helpers/codes/status"

# create directory's if not exists
mkdir -p "${RESPONSE_DIR}"
mkdir -p "${STATUS_DIR}"

echo "Downloading ${RESPONSE_FILE}..."
curl -s -o "$RESPONSE_DIR/$RESPONSE_FILE" "$BASE_URL/$RESPONSE_FILE"
if [ $? -eq 0 ]; then
    echo "> $RESPONSE_FILE saved to $RESPONSE_DIR/"
else
    echo "> Failed to download $RESPONSE_FILE"
fi

echo "Downloading ${STATUS_FILE}..."
curl -s -o "$STATUS_DIR/$STATUS_FILE" "$BASE_URL/$STATUS_FILE"
if [ $? -eq 0 ]; then
    echo "> $STATUS_FILE saved to $STATUS_DIR/"
else
    echo "> Failed to download $STATUS_FILE"
fi

echo "Done."
```
