package kuwo

import (
	"strings"
	"testing"
)

func Test_GenSecretHeader(t *testing.T) {
	type args struct {
		t string
		e string
	}
	tests := []struct {
		name    string
		args    args
		wantSub string
	}{
		// {
		// 	name: "Test h",
		// 	args: args{
		// 		t: "Hm_Iuvt_cdb524f42f0cer9b268e4v7y735ewrq2324=rExGYaCYBn7w7DhFRnQmW7ZAiwjGSQj5; path=/; expires=Thu, 15 Feb 2024 06:55:29 GMT",
		// 		e: "Hm_Iuvt_cdb524f42f0cer9b268e4v7y735ewrq2324",
		// 	},
		// 	wantSub: "f1918b031ccd94b4a4e26bece59b4e6f8ccacdbf102a90f3a97fbf4af718095a",
		// },
		{
			name: "Test h 2",
			args: args{
				t: "Hm_Iuvt_cdb524f42f23cer9b268564v7y735ewrq2324=esCaXZwdwcHNktEEaTtakFrAbbpn68wc; path=/; expires=Thu, 10 Apr 2025 07:38:15 GMT",
				e: "Hm_Iuvt_cdb524f42f23cer9b268564v7y735ewrq2324",
			},
			wantSub: "4623c8c3497d1f4fa946e8df20cef468699b85b65cbcee466be20bf7b0be12ef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &crypto{
				cookie: tt.args.t,
				key:    tt.args.e,
			}
			got := c.GenSecretHeader()
			if !strings.HasPrefix(got, tt.wantSub) {
				t.Errorf("h() = %v, no prefix with %v", got, tt.wantSub)
			}
		})
	}
}
