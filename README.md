# PoW client
An implementation of a client requesting the challenge, solving it and getting the quote from the server.

## Build
docker build -t powclient .

## Run
docker run -e "SERVER=192.168.1.13:8080" powclient

SERVER points to the server that implements challenge mechanism.