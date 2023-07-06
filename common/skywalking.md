[docker](https://www.cnblogs.com/MrSong97/p/16733791.html)

win:  https://juejin.cn/post/7159460571312553998

https://github.com/apache/skywalking/tags


9版本以上  jdk11

历史 版本 https://archive.apache.org/dist/skywalking/8.9.0/

微服务中请改下面参数 ：DSW_AGENT_NAME，修改为服务的名字

-javaagent:D:\dev_tool\apache-skywalking-apm-bin\agent\skywalking-agent.jar -DSW_AGENT_NAME=gateway -DSW_AGENT_COLLECTOR_BACKEND_SERVICES=192.168.99.60:11800

