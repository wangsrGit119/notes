## 安装依赖
> pip install opencv-python

## Demo代码

```python
import cv2

video_capture = cv2.VideoCapture(0)

def openCamera():
	while True:
		ret, frame = video_capture.read()
		cv2.imshow('Video', frame)
		# Hit 'q' on the keyboard to quit!
		if cv2.waitKey(1) & 0xFF == ord('q'):
			break
	video_capture.release()
	cv2.destroyAllWindows()
if __name__ == '__main__':
	openCamera()

```