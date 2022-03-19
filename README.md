中文 | [English](./README_en.md)

# media-get
在给定的 url 里解析出媒体文件，例如音频视频

## 示例
```shell
# 下载 b 站的视频
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0"

# 只下载音频，加上该参数 `-t audio`
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0" -t audio
```

## 如何下载或更新该工具
以下方式任选一种
- 在 release 页面下载： [Click me](https://github.com/foamzou/media-get/releases)
- Mac 用户可以使用 brew install (等 star 多了才能在 brew 申请通过，可以支持一下...)

## 依赖
需要预先安装 `ffmpeg`
- 可以从官网下载: https://www.ffmpeg.org/download.html
- Mac用户可以: `brew install ffmpeg`

## 参数说明
```shell
Usage:
  media-get [OPTIONS]

Application Options:
  -u, --url=         以 http[s]:// 开头的 url
  -o, --out=         下载文件存储目录, 文件名可选. (默认: 当前目录)
  -t, --type=        希望下载的媒体类型 [auto/audio/video/all] (默认: auto)
      --addMediaTag  将歌曲名、专辑名、歌手等 tag 信息添加到歌曲文件里
  -m, --metaOnly     只获取媒体信息，不下载文件
      --metaFormat=  媒体信息以何种形式展示。支持 plain/json. 默认为 plain
  -l, --logLevel=    日志输出级别。支持 silence/error/warn/info/debug. 默认为 info

Help Options:
  -h, --help         Show this help message
```

## 编译
```shell
make build
```

## 支持的网站
Site | audio | video | Remark
:------------ | :------------- | :------------- | :-------------
[b站](https://www.bilibili.com/) | ✅ | ✅ | 
[youtube](https://www.youtube.com/) | ✅ | ✅ | 
[网易云音乐](https://music.163.com/) | ✅ | ✅ | 包含 歌曲、MV、Mlog、电台
[qq音乐](https://y.qq.com/) | ✅ | ⌛ | 
[抖音](https://www.douyin.com/) | ✅ | ✅ |
[咪咕音乐](https://music.migu.cn/) | ✅ | ⌛ |
[酷狗](https://www.kugou.com/) | ✅ | ⌛ |
[酷我](https://www.kuwo.cn/) | ✅ | ⌛ |

## 致谢
- Netease Encrypt Method: https://github.com/872409/music-get