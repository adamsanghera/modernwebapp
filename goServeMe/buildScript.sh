CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o ./build/main .
# CGO_ENABLED=0 disables dynamic lib links
# GOOS=linux sets target OS to linux
# -a means rebuild all the packages used in the program

docker build -t website_server ./
# Build a container for hosting our compiled code.