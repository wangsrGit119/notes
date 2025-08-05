

## web连接页面

https://github.com/tabixio/tabix


## 服务端

```shell
docker pull clickhouse/clickhouse-server:23.8

docker run  -d --name=temp-clickhouse-server clickhouse/clickhouse-server:23.8
mkdir -p /storage/clickhouse/conf /storage/clickhouse/data /storage/clickhouse/log

docker cp temp-clickhouse-server:/etc/clickhouse-server/users.xml /storage/clickhouse/conf/users.xml
docker cp temp-clickhouse-server:/etc/clickhouse-server/config.xml /storage/clickhouse/conf/config.xml

docker rm -f temp-clickhouse-server



```

### 修改配置

```shell
vi config.xml  修改 远程连接
找到  <listen_host>::</listen_host>  去掉注释

vi users.xml 修改默认密码
找到  <password>123456789</password> 修改密码
默认账户密码即为 default/12345678


```
### web连接工具

http://dash.tabix.io/dashboard

可私有化部署

```
docker pull spoonest/clickhouse-tabix-web-client
docker run -d -p 18080:80 --name tabix spoonest/clickhouse-tabix-web-client


```



> 8123 clockhouse  http端口   9004 mysql连接端口  9005 pgsql连接端口 9000端口是ClickHouse的默认端口，用于ClickHouse的原生协议

```shell
docker run -d --restart=always --name=clickhouse-server \
-p 19000:9000 \
-p 8123:8123 \
-p 9004:9004 \
-p 9005:9005 \
-v ${PWD}/clickhouse/conf/config.xml:/etc/clickhouse-server/config.xml \
-v ${PWD}/clickhouse/conf/users.xml:/etc/clickhouse-server/users.xml \
-v ${PWD}/clickhouse/data:/var/lib/clickhouse/ \
clickhouse/clickhouse-server:23.8
```

```
services:
  clickhouse-server:
    image: clickhouse/clickhouse-server:23.8
    container_name: clickhouse-server
    restart: always
    ports:
      - "19000:9000"
      - "8123:8123"
      - "9004:9004"
      - "9005:9005"
    volumes:
      - ./clickhouse/conf/config.xml:/etc/clickhouse-server/config.xml
      - ./clickhouse/conf/users.xml:/etc/clickhouse-server/users.xml
      - ./clickhouse/data:/var/lib/clickhouse/
```
