# 🛒 Shop2 – Google Cloud Spanner (Emulator Setup)

## Initialize Google spanner database emulator
```sh 
 docker compose up -d
```
## Create Instance and Database

```sh
gcloud spanner instances create test-instance \
  --config=emulator-config \
  --description="Local Test Instance" \
  --nodes=1

gcloud spanner databases create shop \
  --instance=test-instance
```

## Export Env 
```sh 
export SPANNER_PROJECT=emulator-project \
export SPANNER_INSTANCE=test-instance  \
export SPANNER_DATABASE=shop \
export APP_PORT=9001 
```
## Migrate Schema 
```sh 
gcloud spanner databases ddl update shop   --instance=test-instance   --ddl-file=migrations/001_initial_schema.sql
```

## Run The app 
```sh 
  go run cmd/server/main.go
```
## Manual Test Product Create Product Using GRPC Client
```
   go run tests/grpc/product.go
```

