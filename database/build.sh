# remove lottery-db container (to prevent errors if it was previous running)
docker container rm -f lottery-db

# # remove lottery-db image 
# docker rmi -f lottery-db

# # remove mysql image
# docker rmi -f mysql:8.0

# build image with tag name lottery-db
docker build . -t lottery-db:latest

# runs container as detached in reference to image lottery-db 
# maps container port 3306 to our computer port 3306
docker run -d --name lottery-db -p 127.0.0.1:3306:3306 lottery-db

# docker stop lottery-db

#docker logs -f lottery-db

# NOTE: Trouble getting container to work property after the first build,
# after building and running twice, everything is normal 