# yuheng

yuheng ☁️ is a convenient and friendly backend service to manage the users' and companies' information.


## Quick Start
1. Adjust the `application.yaml` file according to your needs (see application.yaml.example for more details).

2. Install the dependencies:
```shell
go mod tidy
```

3. Run the application:
```shell
go run main.go
```

## Docker

```dockerfile
docker pull cocoshe/yuheng:api
```

if you want to make your own docker image:
```dockerfile
docker build -t cocoshe/yuheng:tagname .
docker push cocoshe/yuheng:tagname
```

## API doc
http://localhost:port/swagger/index.html
http://psyqlk.space/api/v1/swagger/index.html