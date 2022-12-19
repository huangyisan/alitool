package domain

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_getRegisteredDomainResponse(t *testing.T) {
	mockDomain := "test.com"
	mockAccountName := "testAccount"
	mockRegionId := "cn-shanghai"
	mockI := domain.Client{}
	mockResponse := &domain.QueryDomainByDomainNameResponse{
		DomainName: mockDomain,
	}
	d := DomainClient{
		mockAccountName,
		mockRegionId,
		&mockI,
	}
	var dc *domain.Client
	convey.Convey("QueryDomainByDomainName return error ", t, func() {
		patches := gomonkey.ApplyMethod(dc, "QueryDomainByDomainName", func(_ *domain.Client, request *domain.QueryDomainByDomainNameRequest) (response *domain.QueryDomainByDomainNameResponse, err error) {
			return nil, fmt.Errorf("mock error")
		})
		defer patches.Reset()
		res := d.getRegisteredDomainResponse(mockDomain)
		convey.So(res, convey.ShouldEqual, nil)
	})
	convey.Convey("QueryDomainByDomainName return mock response", t, func() {
		patches := gomonkey.ApplyMethod(dc, "QueryDomainByDomainName", func(_ *domain.Client, request *domain.QueryDomainByDomainNameRequest) (response *domain.QueryDomainByDomainNameResponse, err error) {
			return mockResponse, nil
		})
		defer patches.Reset()
		res := d.getRegisteredDomainResponse(mockDomain)
		convey.So(res, convey.ShouldEqual, mockResponse)
	})
}
