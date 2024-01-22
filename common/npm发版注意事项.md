
## 发布流程

- 更改版本号
- npm login  注意邮箱接受到的二次验证动态验证码
- npm publish --access public


## 发布注意
1.正式包发布24小时内同一个包只能发布一次

2.同一个账号一天发布不同的包也是限制到 5个


## 本地调试
3.尽量在本地测试再发布到，别的项目引用，使用 npm  link 

**【如果 npm link执行 失败建议使用 yarn】**

```
cd 组件目录
npm link 
```
## 解释
当你在一个包项目下执行npm link后,会发生以下事情:
1. 它会在本地的npm全局包缓存中创建一个该包的符号链接。这个全局包缓存的路径一般为:
- Mac/Linux: ~/.npm-global
- Windows: %AppData%\npm-global
2. 你可以打开这个路径,会看到一个与你的包名同名的文件夹,这就是npm link创建的全局符号链接文件夹。


然后在需要引入的项目执行
```
npm link   【前面组件的名字】
```

先在使用npm包的项目的文件目录下解除特定的链接
npm unlink packageName
2.32 再在npm包所在的文件目录下去除全局链接
npm unlink 

到这里其实就OK了，但是如果你还想：
2.33 强制解除创建的某个特定全局链接
sudo npm rm --global packageName
复制代码
2.34 查看所有创建的全局链接名称
npm ls --global --depth 0
