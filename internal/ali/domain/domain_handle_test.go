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

func Test_getAllRegisteredDomainsResponse(t *testing.T) {
	mockDomain := "test.com"
	mockAccountName := "testAccount"
	mockRegionId := "cn-shanghai"
	mockI := domain.Client{}
	var dc *domain.Client
	d := DomainClient{
		mockAccountName,
		mockRegionId,
		&mockI,
	}
	mockQueryDomainListResponse := &domain.QueryDomainListResponse{
		NextPage:  false,
		RequestId: "testID",
		Data: domain.DataInQueryDomainList{
			Domain: []domain.Domain{
				{
					DomainName: mockDomain,
				},
			},
		},
	}
	convey.Convey("mock QueryDomainList return err", t, func() {
		patches := gomonkey.ApplyMethod(dc, "QueryDomainList", func(_ *domain.Client, request *domain.QueryDomainListRequest) (response *domain.QueryDomainListResponse, err error) {
			return nil, fmt.Errorf("mock error")
		})
		defer patches.Reset()
		res := d.getAllRegisteredDomainsResponse()
		convey.So(res, convey.ShouldEqual, nil)
	})
	convey.Convey("mock QueryDomainList", t, func() {
		patches := gomonkey.ApplyMethod(dc, "QueryDomainList", func(_ *domain.Client, request *domain.QueryDomainListRequest) (response *domain.QueryDomainListResponse, err error) {
			return mockQueryDomainListResponse, nil
		})
		defer patches.Reset()
		res := d.getAllRegisteredDomainsResponse()
		convey.So(res, convey.ShouldResemble, []*domain.QueryDomainListResponse{mockQueryDomainListResponse})
	})
}
