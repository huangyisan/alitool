package ssl

import (
	"alitool/internal/pkg/common"
	. "alitool/internal/pkg/mylog"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"
)

type expireCertsInfo map[string]string

func (s *SSLClient) getListCertResponse() *cas.ListUserCertificateOrderResponse {
	request := cas.CreateListUserCertificateOrderRequest()
	request.Scheme = "https"
	response, err := s.I.ListUserCertificateOrder(request)
	if err != nil {
		LoggerNoT.Println(err.Error())
	}
	return response

}

func (s *SSLClient) getListUploadCertResponse() *cas.ListUserCertificateOrderResponse {
	request := cas.CreateListUserCertificateOrderRequest()
	request.Scheme = "https"
	request.OrderType = "UPLOAD"
	response, err := s.I.ListUserCertificateOrder(request)
	if err != nil {
		LoggerNoT.Println(err.Error())
	}
	return response
}

func (s *SSLClient) getExpireCertByAccount(expireDay int) (certs expireCertsInfo) {
	certs = make(expireCertsInfo)
	res := s.getListCertResponse()
	//LoggerNoT.Printf("%#v", res.CertificateOrderList)
	if len(res.CertificateOrderList) > 0 {
		for _, v := range res.CertificateOrderList {
			if common.DaysBetweenNowAndTimestamp(v.CertEndTime) >= 0 && common.DaysBetweenNowAndTimestamp(v.CertEndTime) < expireDay {
				certs[v.Domain] = common.FormatTimestampToDateString(v.CertEndTime)
			}
		}
		return certs
	}
	return nil
}

func (s *SSLClient) getExpireUploadCertByAccount(expireDay int) (certs expireCertsInfo) {
	certs = make(expireCertsInfo)
	res := s.getListUploadCertResponse()
	if len(res.CertificateOrderList) > 0 {
		for _, v := range res.CertificateOrderList {
			if common.DaysBetweenNowAndDate(v.EndDate) >= 0 && common.DaysBetweenNowAndDate(v.EndDate) < expireDay {
				certs[v.Sans] = v.EndDate
			}
		}
		return certs
	}
	return nil
}
