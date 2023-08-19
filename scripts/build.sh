for dir in internal/app/lambda/*/ ; do
  echo "Building from $dir"
  cd "$dir" && GOOS=linux GOARCH=arm64 go build -o bootstrap . && zip bootstrap.zip bootstrap
done