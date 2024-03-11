
## 查看浏览器支持编码格式

```
const supportsSetCodecPreferences = window.RTCRtpTransceiver &&
	'setCodecPreferences' in window.RTCRtpTransceiver.prototype;
console.log(supportsSetCodecPreferences)
const {
	codecs
} = RTCRtpSender.getCapabilities('video');
console.log(codecs)
```
