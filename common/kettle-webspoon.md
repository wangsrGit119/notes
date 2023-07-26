
# kettle-spoon是 kettle客户端的bs版本

## 仓库分支：
https://github.com/pentaho/pentaho-kettle/tree/webspoon-9.2.0.3

## 构建
看readme 

## 基础教程
https://devpress.csdn.net/gitcode/6412b20f986c660f3cf92998.html

## docker 容器

> 很久没更新了

https://hub.docker.com/r/hiromuhota/webspoon

```
docker run -d -p 18086:8080 \
-v /home/kettle-webspoon/kettle:/home/tomcat/.kettle -v /home/kettle-webspoon/pentaho:/home/tomcat/.pentaho \
hiromuhota/webspoon
```