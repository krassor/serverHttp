#go build -o bin/app-serverHttp app/main.go./
echo "stopping app-server-http-news container"
docker stop app-server-http-news
echo "removing app-server-http-news container"
docker rm app-server-http-news
echo "removing app-server-http-news image"
docker rm app-server-http-news
echo "building new app-server-http-news image"
docker build -t app-server-http-news .
docker run -it -d --publish 8001:8001 --name app-server-http-news-v1.1 app-server-http-news