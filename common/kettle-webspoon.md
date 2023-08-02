
# kettle-spoon是 kettle客户端的bs版本

## 仓库分支：
https://github.com/pentaho/pentaho-kettle/tree/webspoon-9.2.0.3

## 构建
看readme 

## 基础教程
https://devpress.csdn.net/gitcode/6412b20f986c660f3cf92998.html

## docker 容器

> 截至 20230727 最新 9.0.22版本

https://hub.docker.com/r/hiromuhota/webspoon

> 下面启动前请配置相关配置

```

docker run -d -p 18086:8080 --name webspoon-9.0.22 hiromuhota/webspoon

#拷贝环境变量配置
docker cp webspoon-9.0.22:/usr/local/tomcat/bin/setenv.sh .

# 添加汉化关键语句到 envsh中
cat setenv.sh 
CLASSPATH=/usr/local/tomcat/lib/webspoon-security-9.0.0.0-423-22.jar
CATALINA_OPTS="-Dorg.apache.tomcat.util.buf.UDecoder.ALLOW_ENCODED_SLASH=true"
JAVA_OPTS="-Duser.language=zh -Duser.region=CN -Dfile.encoding=UTF-8"

docker cp webspoon-9.0.22:/usr/local/tomcat/lib /home/kettle-webspoon/
docker cp webspoon-9.0.22:/usr/local/tomcat/samples /home/kettle-webspoon/

# 移除

docker rm -f webspoon-9.0.22


# 挂在ktr目录 和 配置目录 
docker run -d -p 18086:8080 -u 0 --name webspoon-9.0.22  -v /home/kettle-webspoon/.kettle/data:/root/.kettle/data  -v /home/kettle-webspoon/lib:/usr/local/tomcat/lib -v /home/kettle-webspoon/setenv.sh:/usr/local/tomcat/bin/setenv.sh  hiromuhota/webspoon


docker exec -it -u 0 webspoon-9.0.22 /bin/bash


```

## 汉化

https://huaweicloud.csdn.net/6356684dd3efff3090b5de9c.html

## 其他问题
https://gitee.com/uxue/dataCollection/blob/master/doc/QA.md#webspoon


https://github.com/HiromuHota/pentaho-kettle/wiki/User%3A-Query-parameters-for-File-open