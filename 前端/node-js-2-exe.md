
## 使用 pkg将 node 服务打包成 可执行程序 比如 exe

```

npm i -g @vercel/ncc

ncc build app.ts -o dist

npm install -g pkg

pkg app.js -t node16-win-x64  --debug -o app.exe 

```

## 用法

https://www.npmjs.com/package/pkg
