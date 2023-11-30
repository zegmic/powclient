Build
docker build -t powclient .

Run
docker run -e "SERVER=192.168.1.13:8080" powclient