package pac

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/getgauge-contrib/gauge-go/testsuit"
	"github.com/openshift-pipelines/release-tests/pkg/clients"
)

func VerifyPipelinesAsCodeEnable(cs *clients.Clients, section, inputField, enable string) (string, error) {
	// Construct the JSON payload based on the 'enable' parameter
	payload := fmt.Sprintf(`{"spec":{"platforms":{"openshift":{"%s":{"%s": %s}}}}}`, inputField, section, enable)

	cmd := exec.Command("oc", "patch", "tektonconfigs.operator.tekton.dev", "config", "--type", "merge", "-p", payload)

	// Run the 'oc' command
	if err := cmd.Run(); err != nil {
		// Step failed - Use testsuit.T.Fail to fail the step and provide an error message
		testsuit.T.Errorf("Failed to set PipelinesAsCode enable status: %v", err)
		return "", err
	}

	// Return a message indicating the status change
	return fmt.Sprintf("PipelinesAsCode enable status has been set to %s", enable), nil
}

