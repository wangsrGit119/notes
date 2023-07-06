
## face_recognition库中 函数基础知识 
> https://face-recognition.readthedocs.io/en/latest/face_recognition.html
> https://blog.csdn.net/u011428210/article/details/106475285
> https://github.com/ageitgey/face_recognition


## 常用函数

### load_image_file
load_image_file(file, mode='RGB') #加载一个图像文件到一个numpy array类型的对象上。



### face_encodings
face_encodings(face_image, known_face_locations=None, num_jitters=1,model='small')  # 给定一个图像，返回图像中每个人脸的128脸部编码（特征向量）。
>face_image： 			输入的人脸图像
>known_face_locations：  可选参数，如果你知道每个人脸所在的边界框
>num_jitters=1： 		在计算编码时要重新采样的次数。越高越准确，但速度越慢（100就会慢100倍）
>model 可选——要使用的模型。“large”或“small”（默认）仅返回 5 点但速度更快

### face_distance
face_distance(face_encodings, face_to_compare,tolerance) # 给定一组面部编码，将它们与已知的面部编码进行比较，得到欧氏距离。对于每一个比较的脸，
> face_encodings 已存在人脸数据集合  
> face_to_compare 要比较的人脸数据
> tolerance：两张脸之间有多少距离才算匹配。该值越小对比越严格，0.6是典型的最佳值默认也是0.6


### batch_face_locations
batch_face_locations(images, number_of_times_to_upsample=1, batch_size=128) #使用cnn人脸检测器返回一个包含人脸特征的二维数组，如果使用了GPU，这个函数能够更快速的返回结果；如果不使用GPU的话，该函数就没必要使用。
> images : 一个包含图像数据的list，每个成员都是一个 numpy array类型
> number_of_times_to_upsample： 从images的样本中查找多少次人脸，该参数的值越高的话越能发现更小的人脸
> batch_size： 每个GPU一次批处理多少个image



### face_locations
face_locations(image, number_of_times_to_upsample=1, model='cnn')  # 单个人脸信息检索和前面批量函数对比
> model="cnn" GPU加速 使用哪种人脸检测模型。“hog” 准确率不高，但是在CPUs上运行更快，“cnn” 更准确更深度（且GPU/CUDA加速，如果有GPU支持的话），默认是“hog”

### compare_faces
compare_faces(known_face_encodings, face_encoding_to_check, tolerance=0.6) #比较脸部编码列表和候选编码，看看它们是否匹配

### face_landmarks
face_landmarks(face_image,face_locations=None,model="large") # 人脸特征提取
>face_image：			输入的人脸图片
>face_locations=None：	可选参数，默认值为None，代表默认解码图片中的每一个人脸。若输入face_locations()[i]可指定人脸进行解码
>model="large"：			输出的特征模型，默认为“large”，可选“small”。当选择为"small"时，只提取左眼、右眼、鼻尖这三种脸部特征。


## 安装请看成型项目章节



## 人脸识别脸切割

> pil_image.save 保存文件为识别后的人脸
> 

```python
import face_recognition
from PIL import Image

def shouPosFace(image_path):
	image = face_recognition.load_image_file(image_path)
	# model="cnn" GPU加速 使用哪种人脸检测模型。“hog” 准确率不高，但是在CPUs上运行更快，“cnn” 更准确更深度（且GPU/CUDA加速，如果有GPU支持的话），默认是“hog”
	# face_locations = face_recognition.face_locations(image, number_of_times_to_upsample=0, model="cnn")
	face_locations = face_recognition.face_locations(image)
	for face_location in face_locations:
		print("face_location",face_location)
		top, right, bottom, left = face_location
		face_image = image[top:bottom, left:right]
		pil_image = Image.fromarray(face_image)
		pil_image.save('./rc.jpg')

if __name__ == '__main__':
	image = './images/wckvFVSMS86VAAAAABJRU5ErkJggg==.png'
	shouPosFace(image)

```
## 关键点位检测

> image_path 本地图片位置

```python

def showMaskPoint(image_path):
	image = face_recognition.load_image_file(image_path)
	face_landmarks_list = face_recognition.face_landmarks(image)
	pil_image = Image.fromarray(image)
	d = ImageDraw.Draw(pil_image)
	for face_landmarks in face_landmarks_list:
		 for facial_feature in face_landmarks.keys():
		 	d.line(face_landmarks[facial_feature], width=5)

```

