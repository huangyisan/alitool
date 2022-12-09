package common

import (
	"testing"
)

func TestIsValidDomain(t *testing.T) {
	type args struct {
		domainName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "base case01",
			args: args{
				domainName: "www.baidu.com",
			},
			want: true,
		},
		{
			name: "true case",
			args: args{
				domainName: "www.baidu.com",
			},
			want: true,
		},
		{
			name: "false case",
			args: args{
				domainName: "baidu",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDomain(tt.args.domainName); got != tt.want {
				t.Errorf("IsValidDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
