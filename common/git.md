
## 提交或者拉不下来常用
> refusing to merge unrelated histories
```shell
git pull origin main --allow-unrelated-histories
```

## 远程仓库创建 本地已存在文件夹执行下面
```shell
git init 
git remote add origin http://xxxxx/xxxx-xxxx.git
git branch -M main
git push -uf origin main
```
## 查看关联仓库地址
```
git remote -v
```