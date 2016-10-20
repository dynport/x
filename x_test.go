package x

import (
	"os/exec"
	"sort"
	"strings"
	"testing"
)

func TestDependencies(t *testing.T) {
	b, err := exec.Command("go", "list", "-f", `{{ join .Deps "\n" }}`, "./...").CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	withDots := []string{}
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		if strings.Contains(line, ".") {
			withDots = append(withDots, line)
		}
	}
	if len(withDots) > 0 {
		sort.Strings(withDots)
		t.Errorf("expected to external dependencies bot git:\n%s", strings.Join(withDots, "\n"))
	}
}
