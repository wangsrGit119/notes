
## 使用 pkg将 node 服务打包成 可执行程序 比如 exe

```

npm i -g @vercel/ncc

## 生成必要的依赖文件
ncc build app.js -o dist
# 安装 pkg
npm install -g pkg
# 进入构建好的 dist中进行基础文件打包
cd dist && pkg index.js

```

## 用法

https://www.npmjs.com/package/pkg


## 举例
> node 服务如下 ，需要 node app.js 启动
> 构建执行程序步骤
> 1. ncc build app.js -o dist
> 2. cd dist && pkg index.js
> 3. app中引入的ffmpeg路径拷贝到dist中即可

```
const NodeMediaServer = require('node-media-server');
const config = {
  rtmp: {
    port: 1935,
    chunk_size: 60000,
    gop_cache: true,
    ping: 30,
    ping_timeout: 60
  },
  auth: {
      api : true,
      api_user: 'admin',
      api_pass: 'admin88888888',
    },
  http: {
      port: 8000,
      mediaroot: './media',
      allow_origin: '*'
    },
	trans: {
	    ffmpeg: './libs/win-amd-64/ffmpeg-amd64.exe',
	    tasks: [
	      {
	        app: 'live',
	        hls: true,
	        hlsFlags: '[hls_time=2:hls_list_size=3:hls_flags=delete_segments]',
	        hlsKeep: true, // to prevent hls file delete after end the stream
	        dash: true,
	        dashFlags: '[f=dash:window_size=3:extra_window_size=5]',
	        dashKeep: true ,// to prevent dash file delete after end the stream,
			mp4: true,
			mp4Flags: '[movflags=frag_keyframe+empty_moov]',
	      }
	    ]
	  }
};

var nms = new NodeMediaServer(config)
nms.run();
```
