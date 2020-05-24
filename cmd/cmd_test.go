package cmd_test

import (
	"testing"

	cmd "github.com/narendranathreddythota/podtnl/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {

	t.Run("test the code generation", func(t *testing.T) {
		if !assert.NotEmpty(t, cmd.GetCode(false)) {
			t.Errorf("Code should be generated")
		}
	})

}
