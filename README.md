## GOLANG WEB APP WITH BASICS

Hi, all!
GO is a new language to me, in this repo I apply some basic web concepts.

## TODO
- [x] Configs, [GoDotEnv](https://github.com/joho/godotenv)
- [x] Rest API, [Gin](https://github.com/gin-gonic/gin), [GORM](https://github.com/go-gorm/gorm)
- [x] Authentication, [jwtauth](https://github.com/frozturk/jwtauth)
- [x] Docker
- [ ] Test
- [ ] CI/CD
- [ ] Api gateway
- [ ] Kubernetes
- [ ] Authorization
- [ ] Angular ui
- [ ] Query Cancellation
- [ ] Logging
- [ ] Access limit
- [ ] Security
- [ ] Sessions
- [ ] Server side rendering
- [ ] Stat collection
- [ ] Swagger
- [ ] Caching

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
