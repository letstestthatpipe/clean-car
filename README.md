# Clean car API project

### API documentation under /swagger-ui

### To build clone this repo and run 
```editorconfig
go build
```

### To run the tests
```editorconfig
go test
```

### To deploy on sloppy.io
1. create an account under sloppy.io
2. in travis set ```SLOPPY_APITOKEN``` env variable 
3. in travis register ```DOCKER_USERNAME``` and ```DOCKER_PASSWORD```
4. in travis set ```DOCKER_IMAGE``` which is the docker image name for the app
5. in travis set ```DOMAIN``` which is the sloppy domain name or an external one if you have any


### mercedes account 
1. register for an account here https://me.secure.mercedes-benz.com
2. login with the account to https://developer.mercedes-benz.com/apis/connected_vehicle_experimental_api and subscribe to the API
3. you will need to create a new app there and set a callback URL such as http://localhost:3333/callback for local dev
4. Client Id and Client secret need to be exported as ```CLIENT_ID``` and ```CLIENT_SECRET``` in docker container, this can be achieved via either ```docker run --env CLIENT_SECRET=xxxxx```
