
## 手动申请 free.ssl

> 文件验证请注意 txt文件中内容
> 如果出现 `DCV 预检测通过，但 CA 未通过验证，请稍后再试` 这是正常的情况，可能是当前验证排队的证书过多，证书验证需要一定时间，一般是10min-30min。如果30min后证书仍未颁发，再考虑是否配置有误。




[教程](https://u.sb/acme-sh-ssl/)

## 下载安装
```
curl https://get.acme.sh | sh -s email=2090183831@qq.com
```
## 开启自动更新
```
acme.sh --upgrade --auto-upgrade
```

## 切换
目前 acme.sh 支持四个正式环境 CA，分别是 Let's Encrypt、Buypass、ZeroSSL 和 SSL.com，默认使用 ZeroSSL，如果需要更换可以使用如下命令：
切换 Let's Encrypt：
```
acme.sh --set-default-ca --server letsencrypt
```

## 使用 HTTP 验证签发证书
> 通配符不支持 HTTP签发
首先我们要做一下准备工作，假设你域名是 example.com，解析到你的服务器让其生效后，我们建立一个目录：

```
mkdir -p /var/www/letsencrypt
```

我们的目的是绑定 http://example.com/.well-known/acme-challenge 到这个目录。

如果您用的 Nginx，那么新建一个配置文件：

```
server {
	listen 80;
	listen [::]:80;
	server_name example.com;

	location /.well-known/acme-challenge {
		root /var/www/letsencrypt;
	}

	location / {
		rewrite	^/(.*)$ https://$host/$1 permanent;
	}
}
```
## 申请多域名

```
acme.sh --issue -d wangsrbus.cn -d www.wangsrbus.cn -w /home/nginxWebUI/letsebcrypt
```
## 泛域名仅支持 DNS 验证（具体根据不同厂商来 dns验证）
申请密钥(DNSPod Token)
腾讯云 https://console.dnspod.cn/account/token/apikey

```
cd ~/.acme.sh
vi account.conf
## 添加如下 dnspod 密钥
export DP_Id="xxxx"
export DP_Key="12222222222"
```

### 开始
```
acme.sh --issue --log --dns dns_dp -d wangsrbus.cn -d *.wangsrbus.cn --key-file /home/nginxWebUI/letsebcrypt/private.key --fullchain-file /home/nginxWebUI/letsebcrypt/cert.crt
```