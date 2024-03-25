## 下载
https://ollama.com/download/windows

## docker 下载安装
https://ollama.com/blog/ollama-is-now-available-as-an-official-docker-image

## 使用

> 第一次对话会下载对应的模型

- 单条输入
```
ollama run llama2-chinese "天空为什么是蓝色的？"
```
- 对话模式
```
ollama run llama2-chinese

```

- server 模式
```
#win环境变量
OLLAMA_ORIGINS = *
OLLAMA_HOST=0.0.0.0 # win

export OLLAMA_ORIGINS=* #linux
export OLLAMA_HOST=0.0.0.0 #linux
ollama server #  先要终止之前的启动服务
```

## gui
https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web
```
docker run -d -p 13001:8080 -e OLLAMA_BASE_URL=http://192.168.0.14:11434 -v open-webui:/app/backend/data --name open-webui --restart always ghcr.io/open-webui/open-webui:main
```
