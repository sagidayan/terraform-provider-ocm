package e2e

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	CI "github.com/terraform-redhat/terraform-provider-rhcs/tests/ci"
	CON "github.com/terraform-redhat/terraform-provider-rhcs/tests/utils/constants"
)

var ctx context.Context
var token string

func TestRHCSProvider(t *testing.T) {
	token = CI.GetEnvWithDefault(CON.TokenENVName, "")
	ctx = context.Background()
	RegisterFailHandler(Fail)
	RunSpecs(t, "RHCS Provider Test")

}
