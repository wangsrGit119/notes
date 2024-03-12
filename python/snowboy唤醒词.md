## 安装训练识别模型
> web可视化录制界面
https://github.com/rhasspy/snowboy-seasalt
```
docker run -it -p 18005:8000 rhasspy/snowboy-seasalt
```


## 开始使用

> 不同环境编译使用
> https://github.com/Kitt-AI/snowboy/blob/master/README_ZH_CN.md


```

git clone https://github.com/Kitt-AI/snowboy.git

```

## swig环境 3.0.10

> 下载：https://sourceforge.net/projects/swig/files/swig/swig-3.0.10/

```

yum -y install gcc-c++
yum -y install gcc

tar -zxvf swig......压缩包

cd 进入后执行
./configure --prefix=/usr/local/swig-3.0.10 --without-pcre【即不需要安装pcre依赖】

make && make install
vi /etc/profile 追加下面内容
export PATH=$ANT_HOME/bin:$PATH
export PATH=/usr/local/swig-3.0.10/bin:$PATH
source /etc/profile
cd snowboy/swig/python3
make
```
