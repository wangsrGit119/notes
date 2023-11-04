

## web连接页面

https://github.com/tabixio/tabix


## 服务端

```shell
docker pull clickhouse/clickhouse-server:23.8

docker run --rm -d --name=temp-clickhouse-server clickhouse/clickhouse-server:23.8
mkdir -p /storage/clickhouse/conf /storage/clickhouse/data /storage/clickhouse/log

docker cp temp-clickhouse-server:/etc/clickhouse-server/users.xml /storage/clickhouse/conf/users.xml
docker cp temp-clickhouse-server:/etc/clickhouse-server/config.xml /storage/clickhouse/conf/config.xm




```

### 修改配置

```shell
vi config.xml  修改 远程连接

vi users.xml 修改默认密码
```


> 8123 clockhouse  tcp端口   9004 mysql连接端口  9005 pgsql连接端口

```shell
docker run -d --name=clickhouse-server \
-p 8123:8123 \
-p 9004:9004 \
-p 9005:9005 \
--volume=/storage/clickhouse/conf/config.xml:/etc/clickhouse-server/config.xml \
--volume=/storage/clickhouse/conf/users.xml:/etc/clickhouse-server/users.xml \
--volume=/storage/clickhouse/data:/var/lib/clickhouse/ \
clickhouse/clickhouse-server:23.8
```