## webRTC获取当前客户端公共IP

```js
getPublicIP() {
		return new Promise((resolve, reject) => {
		  var configuration = {
			iceServers: [
			  {
				urls: "stun:stun.l.google.com:19302",
			  },
			],
		  };
		  var pc = new RTCPeerConnection(configuration);
		  pc.onicecandidate = function (event) {
			  console.log(event.candidate)
			if (event.candidate === null) {
				console.log(pc.localDescription.sdp)
			  var publicIP = pc.localDescription.sdp.match(/(c=IN IP4\s)([^\s]*)/)[2];
			  console.log("Public IP Address:", publicIP);
			  resolve(publicIP);
			}
		  };
		  pc.createDataChannel("");
		  pc.createOffer().then((desc) => {
			pc.setLocalDescription(desc);
		  });
		});
  }
```

```js



getPublicIP('stun:stun.l.google.com:19302')

function getPublicIP(stunserver) {
		return new Promise((resolve, reject) => {
		  var configuration = {
			iceServers: [
			  {
				urls: stunserver,
			  },
			],
		  };
		  var pc = new RTCPeerConnection(configuration);
		  pc.onicecandidate = function (event) {
			  console.log(event.candidate)
			if (event.candidate === null) {
				console.log(pc.localDescription.sdp)
			  var publicIP = pc.localDescription.sdp.match(/(c=IN IP4\s)([^\s]*)/)[2];
			  console.log("Public IP Address:", publicIP);
			  resolve(publicIP);
			}
		  };
		  pc.createDataChannel("");
		  pc.createOffer().then((desc) => {
			pc.setLocalDescription(desc);
		  });
		});
  }


```
