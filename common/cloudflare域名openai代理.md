
## 基础步骤

1.国内域名先修改DNS解析托管cloudflare

## vercel获取外部gpt-demo反代地址
1.fork https://github.com/ddiu8081/chatgpt-demo，fork名字注意不要重复
https://github.com/Yidadaa/ChatGPT-Next-Web
2.注册登录vercel，创建project，导入github该fork项目
3.注意配置openai的key和访问密码
3.设置中添加自定义域名

## cloudflare 解析cname
域名dns托管到cloudflare后，添加canme解析到vercel生成的地址

## 出现重定向次数过多
在 Cloudflare 的 SSL/TLS 设置里设置成 FULL 模式
在不起作用 则 修改 dns  proxy status


https://dash.cloudflare.com/

https://vercel.com/


