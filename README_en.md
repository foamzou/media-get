English | [中文](./README.md)

# media-get
Get the media through the url

## Quick Start
### Download media file from the url 
```shell
# download the bilibili video with the url
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0"

# only download the audio, use `-t audio`
media-get -u "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0" -t audio
```

### Search media with keyword

The program would list the result sorted by keyword related, then you can download with the url
```shell
# With the keyword, such as "songName artistName"
media-get -k "搁浅 周杰伦" 

# The result will be sorted more accurate if specific these info 
media-get --searchSongName="搁浅" --searchArtist="周杰伦" 
media-get --searchSongName="搁浅" --searchArtist="周杰伦" --searchAlbum="七里香"

# Just want to search in some source
media-get -k "搁浅 周杰伦" --sources="migu,bilibili"

# Search in all source but some source
media-get -k "搁浅 周杰伦" --excludeSource="youtube"
```
### Proxy setting
The proxy is set in the form of environment variables or json files, with priority environment variables > configuration files.

#### Set by env
You only need to set the ENV_MEDIA_GET environment variable to take effect, the environment variable is also in json format
```markdown
{
    "proxy":{
        "douyin":"http://a.com",
        "bilibili":"http://b.com"
    }
}
```
#### Set by config file
You need to create a file named .media-get.json in the HOME directory of the current login, with the same contents as the environment variables set above, also a json format configuration

### Call the program on your application
The program support output info with JSON. You can search with keyword then filter the result you wanna.
After that, call the program with the url to download media file

BTW, my another open source project [Personal Music Assistant](https://github.com/foamzou/personal-music-assistant) , will use the program to do awesome thing.
```shell
media-get --searchSongName="그리움에 가까운" --searchArtist="Hello Gayoung" --infoFormat=json 
```

## Download / Update the tool
One of the following way can be downloaded
- from the latest release page. [Click me](https://github.com/foamzou/media-get/releases)
- brew install (coming soon...)

## Requirement
Require `ffmpeg`, please install it before use the tool.
- From official website: https://www.ffmpeg.org/download.html
- Mac: `brew install ffmpeg`

## Usage
```shell
Usage:
  main [OPTIONS]

Application Options:
  -u, --url=            Url from website. Start with http[s]://
  -o, --out=            The output file will be place to here(dir/filename). Default: current dir
  -t, --type=           The resource type you want to get [auto/audio/video/all] (default: auto)
  -m, --metaOnly        Output meta info only. Would not download audio
      --infoFormat=     Support plain/json. Default: plain
      --addMediaTag     Add media tag into file from meta information
  -l, --logLevel=       Support: silence/error/warn/info/debug. Default: info

Search Options:
  -k, --keyword=        If search song, keyword or searchSongName is required. Would use keyword to search if both of them are specify. Such as "songName singer"
      --searchSongName= song name
      --searchArtist=   artist name
      --searchAlbum=    album name
      --searchType=     Support: song, Default: song
      --sources=        Search in the specific source, separate with comma. Support: bilibili,douyin,kugou,kuwo,migu,netease,qq,youtube,qmkg. Default search in all
      --excludeSource=  Search not in the specific source, separate with comma

Help Options:
  -h, --help            Show this help message

```

## Build
```shell
make build
```

## Support source
Site | Source Name | Audio | Video | Search
:------------ | :------------- | :------------- | :------------- | :------------- 
[bilibili](https://www.bilibili.com/) | bilibili | ✅ | ✅ | ✅
[youtube](https://www.youtube.com/) |youtube | ✅ | ✅ | ✅
[netease](https://music.163.com/) |netease | ✅ | ✅ | ✅
[qq music](https://y.qq.com/) | qq |✅ | ⌛ | ✅
[douyin](https://www.douyin.com/) | douyin |✅ | ✅  | ✅
[migu music](https://music.migu.cn/) |migu | ✅ | ⌛ | ✅
[kugou](https://www.kugou.com/) |kugou | ✅ | ⌛ | ✅
[kuwo](https://www.kuwo.cn/) | kuwo |✅ | ⌛ | ✅
[qmkg](https://kg.qq.com/) | qmkg |✅ | ⌛ | ⌛


## Thanks
- DouYin XB: https://www.jianshu.com/p/4db4f4d6a536
- Netease Encrypt Method: https://github.com/872409/music-get