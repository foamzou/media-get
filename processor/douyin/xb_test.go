package douyin

import (
	"testing"
)

func Test_genXB(t *testing.T) {
	type args struct {
		urlParams string
		ua        string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				urlParams: "device_platform=webapp&aid=6383&channel=channel_pc_web&search_channel=aweme_general&sort_type=0&publish_time=0&keyword=foamhahaaha&search_source=normal_search&query_correct_type=1&is_filter_search=0&from_group_id=&offset=0&count=10&version_code=170400&version_name=17.4.0&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=zh-CN&browser_platform=MacIntel&browser_name=Chrome&browser_version=104.0.0.0&browser_online=true&engine_name=Blink&engine_version=104.0.0.0&os_name=Mac+OS&os_version=10.15.7&cpu_core_num=8&device_memory=8&platform=PC&downlink=10&effective_type=4g&round_trip_time=0&webid=7131372889845188104&msToken=",
				ua:        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := genXB(tt.args.urlParams, tt.args.ua); err != nil {
				t.Errorf("genXB() faield, err: %v", err)
			}
		})
	}
}
