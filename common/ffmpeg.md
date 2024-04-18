

## 安装

### linux 最新 可运行文件下载
https://johnvansickle.com/ffmpeg/

### win最新
https://www.gyan.dev/ffmpeg/builds/
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

```
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
```
## RTP推流

- srs
> ffmpeg -re -i target.mp4 -vcodec h264 -acodec aac -f rtp_mpegts rtp://124.x.x.x:8000

## 推流转码优化

> 如果指定视频编码为 copy则后面再设置FPS/maxrate/minrate则不会生效，同时设置分辨率也是不行的，会报错。

```shell
// -c:v 为 copy则不能设置 -r fps
ffmpeg -i {0} -c:v libx264 -preset veryfast -r 10 -minrate 100k -maxrate 200k -bufsize 500k -c:a aac -f flv {1}
```

## ffmpeg录制视频

> 录制视频的时候如果录制为mp4,且突然中断进程，那么大概率文件是损坏的，无法播放。解决方案

1.录制参数更改

> 使用  `"-movflags", "faststart" `或者 `"-movflags", "frag_keyframe+empty_moov"`,但是这样有个缺点，就是录制的视频前端播放加载的时候要等整个视频加载完毕才能播放，如果视频比较大则加载很耗时

```golang
ffmpegArgs := []string{
		"-i", url,
		"-c:v", "copy",
		"-c:a", "aac",
		"-f", "mp4",
		// "-movflags", "faststart", //moov box移动到文件的开头，以便在网络传输时能够更快地开始播放。这样可以减少用户等待时间，并减少因为网络传输未完成而导致的中断损坏的风险
		"-movflags", "frag_keyframe+empty_moov",//告诉FFmpeg在关键帧之间进行分段。这样可以确保每个分段都从关键帧开始，也可以防止突然终止造成的文件损坏
		fileResPath,
	}

```

2.录制为 HLS或者DASH，然后转mp4提供下载

> fileList.txt内容

```txt
file 'segment-0.ts'
file 'segment-1.ts'
file 'segment-2.ts'
file 'segment-3.ts'
```

```golang
// mp4FilePath为输出mp4文件
cmd := exec.Command(ffmpegPath, "-y", "-f", "concat", "-safe", "0", "-i", fileList.txt, "-c", "copy", mp4FilePath)

ffmpeg -y -f concat -safe 0 -i filelist_tmp.txt -c copy ef849a79-4b1b-4ee1-9397-53243c1bf967.mp4


ffmpeg -y -f concat  -i filelist_tmp.txt -c copy ef849a79-4b1b-4ee1-9397-53243c1bf967.mp4

ffmpeg -allowed_extensions ALL -i ef849a79-4b1b-4ee1-9397-53243c1bf967.m3u8 -c copy ef849a79-4b1b-4ee1-9397-53243c1bf967.mp4


```


