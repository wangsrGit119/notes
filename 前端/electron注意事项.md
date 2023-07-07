
## 构建打包

1.electron 打包会根据package.json中的build下面files文件过滤，因此打包可以在此处排除文件夹

2.打包会将当前目录的所有文件都打包进去包括比如vue项目下`src/*`所有源码，因此打包可以排除 当然核心的electron-main.js 相关的还是打包在内的["src/*.js","src/assets/**/*","dist/**/*"],需要明确的是 我所提到的整个目录指的是packjson同层级的所有因此不止src还有public和dist等

3.主窗体加载打包后的文件也是根据electron的main.js的路径来的，如果你的这个核心js在src下面，那么打包后获取静态文件如下：
> mainWindow.loadURL(path.join(__dirname,'../dist/index.html'+"?username=123"));