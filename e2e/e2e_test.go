package e2e

import (
	"testing"

	_ "github.com/hashicorp/nomad/e2e/affinities"
	_ "github.com/hashicorp/nomad/e2e/example"
	_ "github.com/hashicorp/nomad/e2e/spread"
	_ "github.com/hashicorp/nomad/e2e/taskevents"
)

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}
