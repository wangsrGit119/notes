
## ssl证书配置

```
server {
  listen 8834 ssl http2;
  ssl_certificate "D:\phpstudy_pro\Extensions\Nginx1.15.11\conf\ssl/localhost_ssl.crt";
  ssl_certificate_key "D:\phpstudy_pro\Extensions\Nginx1.15.11\conf\ssl/localhost_key.key";
  ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3;
  location / {
      proxy_pass http://192.168.1.13:18005/;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Host $http_host;
    proxy_set_header X-Forwarded-Port $server_port;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    
  }

}

```
