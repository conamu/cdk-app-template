for dir in internal/app/lambda/*/ ; do
  echo "Building from $dir for arch $arch"
  cd "$dir" && GOOS=linux GOARCH=$arch go build -o bootstrap . && zip bootstrap.zip bootstrap
  cd ../../../../
done