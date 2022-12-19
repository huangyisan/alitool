package dns

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestDnsClient_getAllDnsDomains(t *testing.T) {
	var a *alidns.Client
	d := DnsClient{
		AccountName: "test",
		RegionId:    "cn-shanghai",
		I:           iDnsClient(a),
	}
	mockDescribeDomainsRequestWithOutNextPage := alidns.DescribeDomainsResponse{
		TotalCount: 19,
		Domains: alidns.DomainsInDescribeDomains{
			Domain: []alidns.DomainInDescribeDomains{
				{
					DomainName: "test.com",
				},
			},
		},
	}

	mockDescribeDomainsRequestWithNextPage := alidns.DescribeDomainsResponse{
		TotalCount: 40,
		Domains: alidns.DomainsInDescribeDomains{
			Domain: []alidns.DomainInDescribeDomains{
				{
					DomainName: "test.com",
				},
			},
		},
	}

	mockDescribeDomainsRequestWithErrorResponse := alidns.DescribeDomainsResponse{}

	convey.Convey("Mock normal response without next page", t, func() {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(a), "DescribeDomains", func(*alidns.Client, *alidns.DescribeDomainsRequest) (response *alidns.DescribeDomainsResponse, err error) {
			return &mockDescribeDomainsRequestWithOutNextPage, nil
		})
		defer patches.Reset()
		hasRecordDomains := d.getAllDnsDomains()
		convey.So(hasRecordDomains, convey.ShouldResemble, recordDnsDomains{"test.com": {}})
	})

	convey.Convey("Mock normal response with next page", t, func() {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(a), "DescribeDomains", func(*alidns.Client, *alidns.DescribeDomainsRequest) (response *alidns.DescribeDomainsResponse, err error) {
			return &mockDescribeDomainsRequestWithNextPage, nil
		})
		defer patches.Reset()
		hasRecordDomains := d.getAllDnsDomains()
		convey.So(hasRecordDomains, convey.ShouldResemble, recordDnsDomains{"test.com": {}})
	})

	convey.Convey("Mock error response", t, func() {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(a), "DescribeDomains", func(*alidns.Client, *alidns.DescribeDomainsRequest) (response *alidns.DescribeDomainsResponse, err error) {
			return &mockDescribeDomainsRequestWithErrorResponse, fmt.Errorf("Mock error response\n")
		})
		defer patches.Reset()
		hasRecordDomains := d.getAllDnsDomains()
		convey.So(hasRecordDomains, convey.ShouldResemble, recordDnsDomains{})
	})
}
