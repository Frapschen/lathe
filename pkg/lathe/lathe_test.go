package lathe

import (
	"strings"
	"testing"
)

func TestRootHelpExposesAgentHint(t *testing.T) {
	root := NewApp(testManifest())
	out, err := execute(root, "--help")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"For agents:",
		"myctl commands --json",
		"myctl commands show",
		"myctl search",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("--help missing %q\nfull output:\n%s", want, out)
		}
	}
}
