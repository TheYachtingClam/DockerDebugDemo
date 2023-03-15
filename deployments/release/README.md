docker build . -f ./build/release/Dockerfile --tag hello-image
docker run --publish 1080:1080 --name hello-image hello-image
curl http://localhost/hello
