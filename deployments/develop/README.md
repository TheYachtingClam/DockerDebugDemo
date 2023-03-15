docker build . -f ./build/develop/Dockerfile --tag hello-image-dev
docker run --publish 1080:1080 --publish 4000:4000 --name debug-server hello-image-dev
curl http://localhost/hello
