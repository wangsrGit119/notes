



## 安装 

```shell
# on Ubuntu or Debian
sudo apt update && sudo apt install ffmpeg

# on Arch Linux
sudo pacman -S ffmpeg

# on MacOS using Homebrew (https://brew.sh/)
brew install ffmpeg

# on Windows using Chocolatey (https://chocolatey.org/)
choco install ffmpeg

# on Windows using Scoop (https://scoop.sh/)
scoop install ffmpeg
```





https://www.cnblogs.com/famhuai/p/ffmpeg.html

## 视频控制
//渐入
i in.mp4 -vf fade=in:0:90 out.mp4                 
//黑白                    
i in.mp4 -vf lutyuv="u=128:v=128" out.mp4   
//锐化
i in.mp4 -vf unsharp=luma_msize_x=7:luma_msize_y=7:luma_amount=2.5 out.mp4   
//反锐化
i in.mp4 -vf unsharp=7:7:-2:7:7:-2 out.mp4
//渐晕
i in.mp4 -vf vignette=PI/4 out.mp4
//闪烁渐晕
i in.mp4 -vf vignette='PI/4+random(1)*PI/50':eval=frame out.mp4
//视频颤抖
i in.mp4 -vf crop="in_w/2:in_h/2:(in_w-out_w)/2+((in_w-out_w)/2)*sin(n/10):(in_h-out_h)/2+((in_h-out_h)/2)*sin(n/7)" out.mp4  
//色彩变幻
i in.mp4 -vf hue="H=2*PI*t:s=sin(2*PI*t)+1" out.mp4
//模糊处理
i in.mp4 -vf boxblur=5:1:cr=0:ar=0 out.mp4
//镜像翻转
i in.mp4 -vf crop=iw/2:ih:0:0,split[left][tmp];[tmp]hflip[right];[left]pad=iw*2[a];[a][right]overlay=w out.mp4
//水平翻转
i in.mp4 -vf geq=p(W-X\\,Y) out.mp4
//垂直翻转
i in.mp4 -vf vflip out.mp4
//浮雕效果
i in.mp4 -vf format=gray,geq=lum_expr='(p(X,Y)+(256-p(X-4,Y-4)))/2' out.mp4
//均匀噪声
i in.mp4 -vf noise=alls=20:allf=t+u out.mp4

## RTP推流

- srs
> ffmpeg -re -i target.mp4 -vcodec h264 -acodec aac -f rtp_mpegts rtp://124.x.x.x:8000

## 推流转码优化

> 如果指定视频编码为 copy则后面再设置FPS/maxrate/minrate则不会生效，同时设置分辨率也是不行的，会报错。

```shell
// -c:v 为 copy则不能设置 -r fps
ffmpeg -i {0} -c:v libx264 -preset veryfast -r 10 -minrate 100k -maxrate 200k -bufsize 500k -c:a aac -f flv {1}
```