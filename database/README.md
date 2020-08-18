Directory that contains dockerfile and database file

## How to run MySQL in docker with the following commands- 

remove lottery-db container (to prevent errors if it was previous running):

###'docker container rm -f lottery-db'

#build image with tag name lottery-db
### 'docker build . -t lottery-db:latest'

#runs container as detached in reference to image lottery-db maps container port 3306 to our computer port 3306
### 'docker run -d --name lottery-db -p 127.0.0.1:3306:3306 lottery-db'

# To stop container from running: 
docker stop lottery-db 

# IF NEEDED: 
#execute bash in lottery-db container

### 'docker exec -it lottery-db /bin/bash'
