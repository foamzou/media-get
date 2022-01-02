package netease

import (
	"testing"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/test_helper"
)

func TestCore_FetchMetaAndResourceInfo(t *testing.T) {
	type fields struct {
		Opts *args.Options
	}
	tests := []struct {
		name          string
		fields        fields
		wantMediaMeta *meta.MediaMeta
		wantErr       bool
	}{
		{
			name:   "Test video",
			fields: fields{Opts: &args.Options{Url: "https://music.163.com/#/video?id=723A994A39CBC7930CB52764296BD4E5"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "超超超甜翻唱《我多喜欢你你会知道》",
				Description: "一只雨仙女吖发表的视频《超超超甜翻唱《我多喜欢你你会知道》》，简介：超甜的《致我们单纯的小美好》主题曲翻唱~\\n我多喜欢你们你们会知道吗！~\\n\\n原唱：王俊琪\\n翻唱：慕诺Nico\\n修音：星離\\n混音：小白\\n视频摄影/剪辑：衔飞\\n\\n哇头一次尝试这种",
				Duration:    199,
				Artist:      "一只雨仙女吖",
				Album:       "网易Video",
				Audios:      []meta.Audio{{Url: ".mp4"}},
				Videos:      []meta.Video{{Url: ".mp4"}},
			},
		},
		{
			name:   "Test MV",
			fields: fields{Opts: &args.Options{Url: "https://music.163.com/#/mv?id=10928596"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "勇气",
				Description: "梁静茹《勇气》高清MV在线观看，发布时间2000-08-01。更多梁静茹相关歌曲高清MV尽在网易云音乐",
				Duration:    358,
				Artist:      "梁静茹",
				Album:       "网易Video",
				Audios:      []meta.Audio{{Url: ".mp4"}},
				Videos:      []meta.Video{{Url: ".mp4"}},
			},
		},
		{
			name:   "Test mlog",
			fields: fields{Opts: &args.Options{Url: "https://st.music.163.com/mlog/mlog.html?id=a1kgGPQY0uK3j5w&type=2&userid=52644372&songId=4875075&startTime=0"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "一首最近超火的《遇到》，网友调侃：是没谈恋爱不敢听的歌曲",
				Description: "二次元音乐秀发表的视频《一首最近超火的《遇到》，网友调侃：是没谈恋爱不敢听的歌曲》，简介：二次元音乐秀：听到这首歌，脑海中满满都是湘琴和直树两个人的爱情故事。更多二次元音乐秀相关的视频尽在网易云音乐",
				Duration:    146,
				Artist:      "二次元音乐秀",
				Album:       "网易Video",
				Audios:      []meta.Audio{{Url: ".mp4"}},
				Videos:      []meta.Video{{Url: ".mp4"}},
			},
		},
		{
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://music.163.com/#/song?id=156182"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "越爱越难过",
				Description: "越爱越难过",
				Duration:    272,
				Artist:      "吴克群",
				Album:       "MagiK Great Hits 新歌+精选",
				Audios:      []meta.Audio{{Url: ".mp3"}},
			},
		},
		{
			name:   "Test program",
			fields: fields{Opts: &args.Options{Url: "https://y.music.163.com/m/program?id=2496095914&userid=52644372&djId=1287164974"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "Black Bird",
				Description: "愿你喜欢",
				Duration:    28,
				Artist:      "原来是二智",
				Album:       "是二智呀",
				Audios:      []meta.Audio{{Url: ".mp3"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Core{
				Opts: tt.fields.Opts,
			}
			gotMediaMeta, err := c.FetchMetaAndResourceInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMetaAndResourceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			test_helper.TestMediaMeta(t, gotMediaMeta, tt.wantMediaMeta)
		})
	}
}
