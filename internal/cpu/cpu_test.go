package cpu_test

import (
	"testing"

	"github.com/anno9474/sysmon/internal/cpu"
)

func TestGetCPUUsage(t *testing.T) {
	usage, err := cpu.GetCPUUsage()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if usage < 0 || usage > 100 {
		t.Fatalf("expected usage between 0 and 100, got %v", usage)
	}
}
