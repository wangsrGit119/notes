
## 官网下载安装

- arm平台
  
https://www.cnblogs.com/fang-d/p/17832995.html

Miniforge3

```
wget https://github.com/conda-forge/miniforge/releases/latest/download/Miniforge3-Linux-aarch64.sh
bash Miniforge3-Linux-aarch64.sh
source ~/.bashrc
```

- conda 
https://repo.anaconda.com/archive/
> 如下 注意提示 按照提示进行即可比如输入 yes、no、回车
```
wget https://repo.anaconda.com/archive/Anaconda3-2023.07-1-Linux-x86_64.sh
chmod +x Anaconda3-2023.07-1-Linux-x86_64.sh
./Anaconda3-2023.07-1-Linux-x86_64.sh
```


- miniconda 
https://docs.conda.io/en/latest/miniconda.html#windows-installers
linux安装：https://www.cnblogs.com/qq419139624/p/14862793.html

windows上 下载安装完成后 将安装目录授予普通用户读写权限 不然非管理员cmd打开创建环境会提示无权限
## 配置国内源

https://mirrors.tuna.tsinghua.edu.cn/help/anaconda/

## 查看本地配置
```
conda info
conda config --show
```
## pip源设置
> 推荐阿里云
```
pip config set global.index-url https://repo.huaweicloud.com/repository/pypi/simple/
pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/
pip config set global.index-url https://pypi.org/
pip config set global.index-url https://pypi.douban.com/simple/
pip config set global.index-url https://mirror.baidu.com/pypi/simple
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple/
pip config set global.index-url https://mirrors.cloud.tencent.com/pypi/simple/
```

## pip包默认安装到 C没有到 conda对应环境则按照下面
python -m site

>USER_BASE: 'C:\\Users\\12156\\AppData\\Roaming\\Python' (exists)
>USER_SITE: 'C:\\Users\\12156\\AppData\\Roaming\\Python\\Python311\\site-packages' (exists)
>ENABLE_USER_SITE: True

修改配置文件追加：
找到安装的Anaconda路径 依次找到路径D:\ProgramData\miniconda3\lib\，里面有一个site.py文件，推荐使用编辑器打开，然后将相应的部分修改成下面的格式
USER_SITE="D:\ProgramData\miniconda3\lib\site-packages"
USER_BASE="D:\ProgramData\miniconda3\Scripts"
**上面仅仅是修改conda默认环境的，如果新创建 env，则要到具体的python环境下找到lib然后修改**



## 环境
```
conda env list | 查看所有环境
conda create -n venv39[自定义名称] python=3.9
conda create -p D:\ProgramData\miniconda3\envs\paddle_ocr_38 python=3.8【指定-p后不能再指定-n，直接在文件夹指定环境名称即可】
conda create -p D:\ProgramData\miniconda3\envs\yolo_311 python=3.11
conda remove --name venv39 --all
```
## 激活环境
### on windows
activate venv39
### on linux
conda activate venv39

## 退出环境
### on windows
deactivate
### on linux
conda deactivate

## 虚拟环境 移动
如果你已经创建了一个虚拟环境，但想要将其移动到另一个文件夹中，可以按照以下步骤操作：

首先，确保你的虚拟环境处于非激活状态。如果当前虚拟环境处于激活状态，请输入以下命令来关闭它：


conda deactivate
复制整个虚拟环境文件夹到你想要存储它的新位置。

打开Anaconda Prompt或终端，并使用以下命令来更新虚拟环境的路径：

```
conda config --add envs_dirs /path/to/new/directory
```
其中，/path/to/new/directory是你想要将虚拟环境移动到的新文件夹路径。

现在，你可以重新激活虚拟环境并在新位置中使用它了：
```
conda activate /path/to/new/directory/your_env_name
```

## 清除缓存

```
conda clean --packages --tarballs
```

### 常见错误

> IMPORTANT: You may need to close and restart your shell after running 'conda init'.

一般执行  `source activate`即可
