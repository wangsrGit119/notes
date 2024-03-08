## SRS有编译好的直接下载即可

https://github.com/ossrs/srs/tags


## zlm 使用vcpkg安装

- 安装 vs2017 一定要选择 c++桌面开发选择，不然 使用 vcpkg安装基础包的时候报错,如下：

```
\Microsoft Visual Studio\2017\Community\VC\Auxiliary/Build\vcvarsall.bat
```

- 安装 vcpkg
https://github.com/microsoft/vcpkg/blob/master/README_zh_CN.md#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B-windows

- 然后 准备环境（如果是linux环境直接执行下面的即可安装zlm）其他的安装完 

```
./vcpkg/vcpkg install zlmediakit\[core,mp4,openssl,webrtc,sctp\]

```

- 按照教程开始编译zlm

  https://docs.zlmediakit.com/zh/guide/install/compilation_instructions_for_windows_version.html#%E7%BC%96%E8%AF%91

- 基础依赖
```
#vcpkg安装openssl
vcpkg install --triplet=x64-windows-static openssl
#编译 libsrtp,并且打开OPENSSL, 需要 ENABLE_OPENSSL, 可编辑 c:\vcpkg\ports\libsrtp\portfile.cmake, 修改
vcpkg_configure_cmake 为如下:
vcpkg_configure_cmake(
  SOURCE_PATH ${SOURCE_PATH}
  PREFER_NINJA
  OPTIONS
    -DENABLE_OPENSSL:BOOL=ON
)
#编译libsrtp
vcpkg install --triplet=x64-windows-static libsrtp
vcpkg install --triplet=x64-windows-static libsrtp

```
