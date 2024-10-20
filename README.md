# **`Golang + Gin` JWT authentication template**
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Grafana](https://img.shields.io/badge/grafana-%23F46800.svg?style=for-the-badge&logo=grafana&logoColor=white)
![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white)


### Small **REST API** template written in `Golang` using `Gin`, `PostgreSQL`, `JWT` and `Redis`. Project: user profile *CRUD* with authorization, authentication and token validation
---
### Endpoints:
1. `POST:/api/v1/users` - registrates new user
2. `POST:/api/v1/auth` - authorizes user, generates JWT token and writing it in cache
3. `POST:/api/v1/auth/changePassword` - changes password with deleting valid token in `Redis` valid tokens storage *(requires token in Bearer header)*
4. `GET:/api/v1/users/:username` - get user profile by username *(requires token in Bearer header)*
5. `GET:/api/v1/users/me` - returns current user profile *(requires token in Bearer header)*
6. `PATCH:/api/v1/users/me` - updates current user profile *(requires token in Bearer header)*

+ Visit http://localhost:8000/swagger/index.html to see `Swagger` specification
---
### Launch guide:
+ Make sure that your `.env` looks like this:
```
DB_HOST=localhost
DB_PORT=5432
DB_PASSWORD=password
DB_USER=postgres
DB_NAME=demo
JWT_SECRET=secret
REDIS_CONNECTION=localhost:6379
```
+ Make sure that you running Redis and PostgreSQL on your machine
+ Run `go run main.go` in terminal
