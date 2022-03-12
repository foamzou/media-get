# media-get
Get the media through the url

## Quick Start
```shell
# download the bilibili video with the url
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0"

# only download the audio, use `-t audio`
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0" -t audio
```

## Download / Update the tool
One of the following way can be download
- from the latest release page. [Click me](https://github.com/foamzou/media-get/releases)
- brew install (coming soon...)

## Requirement
Require `ffmpeg`, please install it before use the tool

## Usage
```shell
Usage:
  media-get [OPTIONS]

Application Options:
  -u, --url=         Url from website. Start with http[s]://
  -o, --out=         The output file will be place to here(dir/filename). Default: current dir
  -t, --type=        The resource type you want to get [auto/audio/video/all] (default: auto)
      --addMediaTag  Add media tag into file from meta information. Default: false
  -m, --metaOnly     Output meta info only. Would not download audio
      --metaFormat=  Support plain/json. Default: plain
  -l, --logLevel=    Support: silence/error/warn/info/debug. Default: info

Help Options:
  -h, --help         Show this help message
```

## Build
```shell
make build
```

## Support media
Site | audio | video | Remark
:------------ | :------------- | :------------- | :-------------
[bilibili](https://www.bilibili.com/) | ✅ | ✅ | 
[youtube](https://www.youtube.com/) | ✅ | ✅ | 
[netease](https://music.163.com/) | ✅ | ✅ | include `song`, `mv`, `mlog`, `radio`
[douyin](https://www.douyin.com/) | ✅ | ✅ |
[migu music](https://music.migu.cn/) | ✅ | - |

## Thanks
- Netease Encrypt Method: https://github.com/872409/music-get