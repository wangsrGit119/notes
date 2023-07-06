
## 官网
https://doris.apache.org/zh-CN/docs/get-starting/

## GitHub
https://github.com/apache/doris/releases/tag/1.2.2-rc01

## 注意事项

1. 安装启动按照官网来即可
2. BE要按照 CPU支不支持 avx2 指令集，按照具体情况下载
3. 默认副本数为3。如果 BE 节点数量小于3，则需指定副本数小于等于 BE 节点数量 [官网说明](https://doris.apache.org/zh-CN/docs/sql-manual/sql-reference/Data-Definition-Statements/Create/CREATE-TABLE)；具体看官网starter示例。
   ```sql
    CREATE TABLE `aj_db` (
    `id` bigint(100) NOT NULL COMMENT 'id',
    `name` varchar(100) DEFAULT NULL COMMENT '姓名',
    `type` varchar(100) DEFAULT NULL COMMENT '类别',
    `money` double NOT NULL DEFAULT '0' COMMENT '金额',
    ) 
    COMMENT '测试表'
    DISTRIBUTED BY HASH(`id`) BUCKETS 1
    PROPERTIES (
        "replication_allocation" = "tag.location.default: 1"
    );

   ```
  4. 清空表 truncate table tablename;

## 节点伸缩请看官网
[节点伸缩](https://doris.apache.org/zh-CN/docs/admin-manual/cluster-management/elastic-expansion/#:~:text=%E5%88%A0%E9%99%A4%20BE%20%E8%8A%82%E7%82%B9%201%20%E8%AF%A5%E5%91%BD%E4%BB%A4%E7%94%A8%E4%BA%8E%E5%AE%89%E5%85%A8%E5%88%A0%E9%99%A4%20BE%20%E8%8A%82%E7%82%B9%E3%80%82%20%E5%91%BD%E4%BB%A4%E4%B8%8B%E5%8F%91%E5%90%8E%EF%BC%8CDoris,%27%2Fbackends%27%3B%20%E4%B8%AD%E7%9A%84%20TabletNum%20%E6%9F%A5%E7%9C%8B%EF%BC%8C%E5%A6%82%E6%9E%9C%E6%AD%A3%E5%9C%A8%E8%BF%9B%E8%A1%8C%EF%BC%8CTabletNum%20%E5%B0%86%E4%B8%8D%E6%96%AD%E5%87%8F%E5%B0%91%E3%80%82%205%20%E8%AF%A5%E6%93%8D%E4%BD%9C%E5%8F%AF%E4%BB%A5%E9%80%9A%E8%BF%87%3A%20)

##  启动常见异常指南

- BE启动异常：

```
 fail to open StorageEngine, res=file descriptors limit is too small
```
> 解决方案命令行输入下面：`ulimit -n 204800` 此为临时方案，永久请修改对应的配置文件
```
echo "* soft nofile 204800"  >> /etc/security/limits.conf
echo "* hard nofile 204800"  >> /etc/security/limits.conf
echo "* soft nproc 204800"  >> /etc/security/limits.conf
echo "* hard nproc 204800 "  >> /etc/security/limits.conf
```






