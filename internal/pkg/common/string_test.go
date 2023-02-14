package common

import "testing"

func TestDomainSuffix(t *testing.T) {
	type args struct {
		domainName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base case01",
			args: args{
				domainName: "baidu.com",
			},
			want: "baidu.com",
		},
		{
			name: "base case02",
			args: args{
				domainName: "www.baidu.com",
			},
			want: "baidu.com",
		},
		{
			name: "base case03",
			args: args{
				domainName: "cc.www.baidu.com",
			},
			want: "baidu.com",
		},
		{
			name: "base case04",
			args: args{
				domainName: "baidu",
			},
			want: "baidu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainSuffix(tt.args.domainName); got != tt.want {
				t.Errorf("DomainSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLastMonth(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "base case 01",
			want: "2023-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLastMonth(); got != tt.want {
				t.Errorf("GetLastMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
