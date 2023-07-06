## 容器搭建
-  说明
>   默认用户 postgres
- 环境变量：
>   POSTGRES_DB: dbname
>   POSTGRES_USERS_USER: postgres
>   POSTGRES_PASSWORD: postgres
- 启动
```shell
docker pull postgres:14.5
mkdir -p /home/pgsql
docker run --name postgres-test \
    --restart=always \
    -e POSTGRES_PASSWORD=61332433 \
    -p 15432:5432 \
    -v /home/pgsql/data:/var/lib/postgresql/data \
    -d postgres:14.5

```