package validate

import (
	e "github.com/cloudposse/atmos/internal/exec"
	u "github.com/cloudposse/atmos/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateComponent(t *testing.T) {
	_, err := e.ExecuteValidateComponent("infra/vpc", "tenant1-ue2-dev", "validate-infra-vpc-component.rego", "opa")
	u.PrintError(err)
	assert.Error(t, err)
}

func TestValidateComponent2(t *testing.T) {
	_, err := e.ExecuteValidateComponent("infra/vpc", "tenant1-ue2-prod", "", "")
	u.PrintError(err)
	assert.Error(t, err)
}

func TestValidateComponent3(t *testing.T) {
	_, err := e.ExecuteValidateComponent("infra/vpc", "tenant1-ue2-staging", "", "")
	u.PrintError(err)
	assert.Error(t, err)
}
