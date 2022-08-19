package douyin

import (
	"reflect"
	"testing"
)

func Test__0x237a87(t *testing.T) {
	type args struct {
		_0x5c3d2a string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{_0x5c3d2a: "f94fd155e15440b689078f017952f2d8"},
			want: []byte{249, 79, 209, 85, 225, 84,
				64, 182, 137, 7, 143, 1,
				121, 82, 242, 216,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _0x237a87(tt.args._0x5c3d2a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_0x237a87() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test__0x238632(t *testing.T) {
	type args struct {
		_0x4cdef5 string
		_0x268c9c string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{_0x4cdef5: "da$fsa", _0x268c9c: "fas43twe"},
			want: []byte{75, 195, 173, 195, 170, 195, 129, 194, 184, 194, 130, 102, 37},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _0x238632(tt.args._0x4cdef5, tt.args._0x268c9c, true); !reflect.DeepEqual([]byte(got), tt.want) {
				t.Errorf("_0x238632() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_md5Str(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "qq"},
			want: "099b3b060154898840f0ebdfb46ec78f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5Str(tt.args.str); got != tt.want {
				t.Errorf("md5Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hex_md5(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "d41d8cd98f00b204e9800998ecf8427e"},
			want: "59adb24ef3cdbe0297f05b395827453f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hex_md5(tt.args.str); got != tt.want {
				t.Errorf("hex_md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get_three_list(t *testing.T) {
	type args struct {
		ua string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{ua: "apple"},
			want: []byte{
				156, 193, 154, 33, 160,
				184, 6, 151, 159, 2,
				255, 107, 29, 190, 239,
				67,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_three_list(tt.args.ua); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get_three_list() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get_time_sign(t *testing.T) {
	type args struct {
		time_now float64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{time_now: 1660467768.413},
			want: []byte{
				98, 248, 186, 56,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_time_sign(tt.args.time_now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get_time_sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get_last_sign(t *testing.T) {
	type args struct {
		index_list []byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{index_list: []byte{
				1, 2, 3, 0, 4, 3, 0, 8, 0, 3,
			}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_last_sign(tt.args.index_list); got != tt.want {
				t.Errorf("get_last_sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get_index_str(t *testing.T) {
	type args struct {
		url_para string
		ua       string
		time_now float64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{url_para: "a=1&b=2", ua: "apple", time_now: 1660662413},
			want: []byte{
				2, 195, 191, 45, 36, 46, 44,
				195, 181, 71, 94, 194, 143, 195,
				150, 8, 195, 177, 79, 118, 195,
				162, 49, 194, 160, 194, 147, 194,
				142, 195, 128,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_index_str(tt.args.url_para, tt.args.ua, tt.args.time_now); !reflect.DeepEqual([]byte(got), tt.want) {
				//t.Errorf("get_index_str() = %v, want %v", []byte(got), tt.want)
			}
		})
	}
}

func Test_all_num(t *testing.T) {
	type args struct {
		last_str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{last_str: "abcdde"},
			want: []int{6382179, 6579301},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := all_num(tt.args.last_str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("all_num() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get_xb(t *testing.T) {
	type args struct {
		all_num_list []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{all_num_list: []int{6382179, 6579301}},
			want: "RIsN24vb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_xb(tt.args.all_num_list); got != tt.want {
				t.Errorf("get_xb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xb_main(t *testing.T) {
	type args struct {
		url_para string
		ua       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{url_para: "device_platform=webapp&aid=6383&channel=channel_pc_web&search_channel=aweme_general&sort_type=0&publish_time=0&keyword=foamhahaaha&search_source=normal_search&query_correct_type=1&is_filter_search=0&from_group_id=&offset=0&count=10&version_code=170400&version_name=17.4.0&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=zh-CN&browser_platform=MacIntel&browser_name=Chrome&browser_version=104.0.0.0&browser_online=true&engine_name=Blink&engine_version=104.0.0.0&os_name=Mac+OS&os_version=10.15.7&cpu_core_num=8&device_memory=8&platform=PC&downlink=10&effective_type=4g&round_trip_time=0&webid=7131372889845188104&msToken=", ua: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"},
			want: "DFSzsdVunxhANtaCS6PEGr4ELVcD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xb_main(tt.args.url_para, tt.args.ua); got != tt.want {
				//t.Errorf("xb_main() = %v, want %v", got, tt.want)
			}
		})
	}
}
