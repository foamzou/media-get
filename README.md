中文 | [English](./README_en.md)

# media-get
在给定的 url 里解析出媒体文件，例如音频视频。并提供代理功能，针对不同的源可以设置不同的代理地址

## 示例
### 在 url 所在的网页下载资源
```shell
# 下载 b 站的视频
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0"

# 只下载音频，加上该参数 `-t audio`
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0" -t audio
```
### 使用关键字搜索资源

本应用会以关键字相关性，对搜索结果进行排序并列出来。然后你可以用列出的 url 下载资源
```shell
# 使用 keyword，例如 "歌名 歌手名"
media-get -k "搁浅 周杰伦" 

# 如果你传入更多信息的话，排序会更准一些
media-get --searchSongName="搁浅" --searchArtist="周杰伦" 
media-get --searchSongName="搁浅" --searchArtist="周杰伦" --searchAlbum="七里香"

# 如果只希望从指定网站搜索
media-get -k "搁浅 周杰伦" --sources="migu,bilibili"

# 如果希望排除一些网站，例如 youtube 国内一般搜索不了。后续会考虑加入代理配置（粒度会在网站而不是全局，因为考虑到某些网站判断到非国内 IP 就不返回结果）
media-get -k "搁浅 周杰伦" --excludeSource="youtube"
```

### 设置代理
代理通过环境变量或者 json 文件的形式设置，优先级环境变量 > 配置文件

#### 环境变量的设置方式
只需要设置 ENV_MEDIA_GET 这个环境变量即可生效，环境变量也是 json 格式的
```markdown
{
    "proxy":{
        "douyin":"http://a.com",
        "bilibili":"http://b.com"
    }
}
```
#### 配置文件的设置方式
需要在当前登陆用的 HOME 目录下新建 .media-get.json 文件，内容如上面环境变量设置的一样，也是一个 json 格式的配置

### 你可以在你的程序里调用本应用
本应用支持以 JSON 格式输出信息。你可以使用搜索功能得到一些列表，然后以你的需求找到目标后再调用本应用下载资源

另外, 我的另一个开源项目 [Personal Music Assistant](https://github.com/foamzou/personal-music-assistant) , 也会使用本应用做一些很棒的事情
```shell
media-get --searchSongName="그리움에 가까운" --searchArtist="Hello Gayoung" --infoFormat=json 
```

## 如何下载或更新该工具
支持 Win、Mac 和 Linux

以下方式任选一种
- 在 release 页面下载： [Click me](https://github.com/foamzou/media-get/releases)
- Mac 用户可以使用 brew install (等 star 多了才能在 brew 申请通过，可以支持一下)

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
      --infoFormat=  媒体信息以何种形式展示。支持 plain/json. 默认为 plain
  -l, --logLevel=    日志输出级别。支持 silence/error/warn/info/debug. 默认为 info

Search Options:
  -k, --keyword=        若想使用搜索功能， keyword 或 searchSongName 是必传的. 如果都传入的话，会使用 keyword 进行搜索. Such as "歌名 歌手名"
      --searchSongName= 歌名
      --searchArtist=   歌手名
      --searchAlbum=    专辑名
      --searchType=     暂时只支持: song, 默认: song
      --sources=        在指定的网站中搜索，使用英文逗号隔开. 目前支持: bilibili,douyin,kugou,kuwo,migu,netease,qq,youtube,qmkg. 默认在全部网站中搜索
      --excludeSource=  排除指定的网站，使用英文逗号隔开

Help Options:
  -h, --help         Show this help message
```

## 编译
```shell
make build
```

## 支持的网站
Site | Source Name | Audio | Video | Search | Environment dependencies
:------------ | :------------- | :------------- | :------------- | :------------- | :-------------
[b站](https://www.bilibili.com/) | bilibili | ✅ | ✅ | ✅ | -
[youtube](https://www.youtube.com/) |youtube | ✅ | ✅ | ✅| -
[网易云音乐](https://music.163.com/) |netease | ✅ | ✅ | ✅ | -
[qq音乐](https://y.qq.com/) | qq |✅ | ⌛ | ✅ | -
[抖音](https://www.douyin.com/) | douyin |✅ | ✅  | ✅ | -
[咪咕音乐](https://music.migu.cn/) |migu | ✅ | ⌛ | ✅ | -
[酷狗](https://www.kugou.com/) |kugou | ✅ | ⌛ | ✅ | -
[酷我](https://www.kuwo.cn/) | kuwo |✅ | ⌛ | ✅ | 需要先安装 Node.js
[全民K歌](https://kg.qq.com/) | qmkg |✅ | ⌛ | ⌛ | -


## 致谢
- DouYin XB: https://www.jianshu.com/p/4db4f4d6a536
- Netease Encrypt Method: https://github.com/872409/music-get
