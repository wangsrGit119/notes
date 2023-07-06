
> socket.io 组件前端访问的时候默认携带了路径 `socket.io`,因此，如果nginx代理后，如下：
```nginx
location /ssso/ {
			proxy_pass http://127.0.0.1:18083/socket.io/;
			proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection 'upgrade';
            proxy_set_header Host $host;
            proxy_cache_bypass $http_upgrade;
		
}
```
> 则前端需要重新添加代理配置路径

```javascript
//this.$serverSocketUrl 只需要配置 wss/ws+ip即可，无需添加代理路径 即时添加也不生效 （要生效则同时要在后端socketio服务器配置 namespace）
this.linkSocket = io(this.$serverSocketUrl, {
    path:'/ssso',
    reconnectionDelayMax: 10000,
    transports: ["websocket"],
    query: this.formInline
});
```