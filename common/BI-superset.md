
## 代码

git clone https://github.com/apache/superset.git

git checkout 2.1.0


## docker 


docker pull apache/superset:2.1.1rc2

https://hub.docker.com/r/apache/superset

## 基础配置
/app/docker/pythonpath_dev/superset_config.py


```
LANGUAGES = {
    'en': {'flag': 'us', 'name': 'English'},
    'zh': {'flag': 'cn', 'name': 'Chinese'},
}
PUBLIC_ROLE_LIKE_GAMMA = True


SECRET_KEY = 'suc'

```

## 启用图表选择联动

- 需要允许图表互相筛选在config.py  修改为以下

```
"DASHBOARD_NATIVE_FILTERS": True,
"DASHBOARD_CROSS_FILTERS": True,
"DASHBOARD_NATIVE_FILTERS_SET": True,
```

- docker 则 修改/superset-master/docker/superset_config.py

```
FEATURE_FLAGS = {
        "DASHBOARD_NATIVE_FILTERS": True,
        "DASHBOARD_CROSS_FILTERS": True,
        "DASHBOARD_NATIVE_FILTERS_SET": True
}
```

## 更多配置
https://superset.apache.org/docs/installation/configuring-superset/


## docker-compose 启动
https://superset.apache.org/docs/installation/installing-superset-using-docker-compose/#installing-superset-locally-using-docker-compose
docker-compose -f docker-compose-non-dev.yml pull
docker-compose -f docker-compose-non-dev.yml up -d


## 官方镜像启动

docker pull apache/superset:2.1.1rc2
https://hub.docker.com/r/apache/superset

```
docker run -d -p 18083:8088 -e "SUPERSET_SECRET_KEY=my_secret_key" --name superset apache/superset:2.1.1rc2

```

## 初始化容器内数据

Setup your local admin account
```
docker exec -it superset superset fab create-admin \
              --username admin \
              --firstname Superset \
              --lastname Admin \
              --email admin@superset.com \
              --password admin
```
Migrate local DB to latest

```
docker exec -it superset superset db upgrade
```

Load Examples

```
 docker exec -it superset superset load_examples

```
Setup roles
```
docker exec -it superset superset init
```
