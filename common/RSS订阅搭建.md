## 搭建
 - 创建文件夹
   
 ```
mkdir rss
 ```
- 进入文件夹
  
```
cd rss/
```

- 创建文件：config.json（“listHeight”控制列表高度，源多可以试试300）
  
```
{
    "values": [
       	"https://www.v2ex.com/feed/ideas.xml",
        "https://www.zhihu.com/rss",
        "https://tech.meituan.com/feed/",
        "http://www.ruanyifeng.com/blog/atom.xml",
        "https://feeds.appinn.com/appinns/",
        "https://v2ex.com/feed/tab/tech.xml",
        "https://www.cmooc.com/feed",
        "http://www.sciencenet.cn/xml/blog.aspx?di=30",
        "https://www.douban.com/feed/review/book",
        "https://www.geekpark.net/rss",
        "https://rss.nodeseek.com/",
        "https://linux.do/latest.rss",
        "https://www.v2ex.com/feed/vps.xml",
        "https://hostloc.com/forum.php?mod=rss&fid=45&auth=389ec3vtQanmEuRoghE%2FpZPWnYCPmvwWgSa7RsfjbQ%2BJpA%2F6y6eHAx%2FKqtmPOg"
    ],
    "refresh": 6,
    "autoUpdatePush": 7,
    "listHeight": "600",
    "webTitle": "Hello suke",
    "webDes":"suke station"
}
```

- 创建文件：docker-compose.yml

```
version: "3"

services:
  server:
    image: topang/rss-reader-mix:latest
    container_name: rss-reader-mix
    restart: always
    ports:
      - "8880:8080"
    volumes:
      - "$PWD/config.json:/app/config.json"
```

- 启动
  
```
docker-compose up -d
访问8880端口（ip:8880）
```
