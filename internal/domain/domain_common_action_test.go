package domain

import (
	"testing"
)

func TestDomainClient_checkDomainStatus(t *testing.T) {
	dc := NewDomainClient("cn-shanghai", "", "")
	dc.GetExpireDomains(300)
}
