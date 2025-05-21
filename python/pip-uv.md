
## uv 安装使用


### 安装
```
pip install uv  或者  conda install uv -c conda-forge
```

## 直接 安装 无py环境

方法 1：使用安装脚本（推荐）
Linux/macOS：
```
curl -LsSf https://astral.sh/uv/install.sh | sh
```
Windows（PowerShell）：
```
irm https://astral.sh/uv/install.ps1 | iex
```


## 创建虚拟环境 

```
#当前目录下创建 .venv 虚拟环境
uv venv --python 3.10
# 激活环境
source .venv/bin/activate

# uv init 会直接在当前目录初始化一个py项目 其中 pyproject.toml 为项目管控工具 如果要修改国内pip源的话 可以在这里

在 pyproject.toml 中添加以下内容：
---------------
[tool.uv.pip]
index-url = "https://mirrors.ustc.edu.cn/help/pypi.html"

## 当然也可以直接安装携带

uv pip install flask --index-url https://pypi.tuna.tsinghua.edu.cn/simple



```

### 默认镜像配置

```
mkdir -p ~/.config/uv/
touch  ~/.config/uv/uv.toml
vi ~/.config/uv/uv.toml

## 使用下面网站内的配置--------------------
https://mirrors.ustc.edu.cn/help/pypi.html
```
