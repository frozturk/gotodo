## GOLANG WEB APP WITH BASICS

Hi, all!
GO is a new language to me, in this repo I apply some basic web concepts.

## TODO
- [x] CONFIGS, [GoDotEnv](https://github.com/joho/godotenv)
- [x] REST API, [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/)
- [x] AUTHENTICATION, [jwt-go](https://github.com/dgrijalva/jwt-go), [go-redis](https://github.com/go-redis/redis)
- [x] Docker
- [ ] API GATEWAY
- [ ] KUBERNETES
- [ ] AUTHORIZATION
- [ ] ANGULAR UI

## HOW TO RUN?

This project uses GoDotEnv to read configs in .env files.
Place a ".env" file in root directory with your configurations;
```
DBHOST=localhost
DBPORT=5432
DBNAME=postgres
DBUSER=postgres
DBPASS=123456789
JWTSECRET=xxxxxxxx
REDISDSN=localhost:6379
```
Then,
```
//With docker
docker-compose up -d

//or
go run .
```
