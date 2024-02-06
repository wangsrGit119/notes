
## 使用 pkg将 node 服务打包成 可执行程序 比如 exe

```

npm i -g @vercel/ncc

## 生成必要的依赖文件
ncc build app.ts -o dist
# 安装 pkg
npm install -g pkg
# 进入构建好的 dist中进行基础文件打包
cd dist && pkg index.js

```

## 用法

https://www.npmjs.com/package/pkg
