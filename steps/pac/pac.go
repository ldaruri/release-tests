package pac

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/openshift-pipelines/release-tests/pkg/pac"
	"github.com/openshift-pipelines/release-tests/pkg/store"
)

var _ = gauge.Step("Set <auto-configure-new-github-repo> section under <pipelinesAsCode> to <auto-configure-new-github-repo: true|false>", func(section, inputField, isEnable string) {
	pac.VerifyPipelinesAsCodeEnable(store.Clients(), section, inputField, isEnable)
})

