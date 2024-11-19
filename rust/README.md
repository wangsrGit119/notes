## 常见开发异常

> rust cargo build一直出现 `Blocking waiting for file lock on package cache`
如果确定没有多个程序占用，可以删除`rm -rf ~/.cargo/.package-cache`，然后再执行
