package gogitversion

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	expected := "9999999.9999999.1"

	version = expected
	defer func() { version = "" }()

	actual := Get()

	if expected != actual {
		t.Errorf("'%s' expected but got '%s'", expected, actual)
	}
}

func TestGet_GitUnknown(t *testing.T) {
	// Force `git describe` to fail by pointing GIT_DIR at a path that is not
	// a git repository. With GIT_DIR explicitly set, git uses it directly and
	// does not walk up the filesystem looking for a parent repo, so this
	// works regardless of where `go test` is invoked from. t.Setenv reverts
	// the env var on test completion, so this only affects this test.
	t.Setenv("GIT_DIR", t.TempDir())

	expected := "unknown"
	actual := Get()

	if expected != actual {
		t.Errorf("'%s' expected but got '%s'", expected, actual)
	}
}

// TestGitDescribeVersion_IgnoresNonVersionTags verifies that a non-`v*` tag
// (for example a `demo-*` tag placed on a demo commit) does not shadow an
// earlier `v*` tag. Regression test for the case where `git describe --tags`
// alone would pick up the most recent demo tag.
func TestGitDescribeVersion_IgnoresNonVersionTags(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not available on PATH")
	}

	repo := t.TempDir()
	gitEnv := append(os.Environ(),
		"GIT_AUTHOR_NAME=test", "GIT_AUTHOR_EMAIL=test@test",
		"GIT_COMMITTER_NAME=test", "GIT_COMMITTER_EMAIL=test@test",
	)
	run := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", args...)
		cmd.Dir = repo
		cmd.Env = gitEnv
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}

	run("init", "-q")
	run("commit", "--allow-empty", "-m", "release")
	run("tag", "v0.1.0")
	run("commit", "--allow-empty", "-m", "demo")
	run("tag", "demo-something")

	// gitDescribeVersion shells out to git in the process cwd, so chdir in
	// for the call and restore afterwards.
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(orig) }()
	if err := os.Chdir(repo); err != nil {
		t.Fatal(err)
	}

	got := gitDescribeVersion()
	if !strings.HasPrefix(got, "v0.1.0") {
		t.Errorf("expected version derived from v0.1.0 tag, got %q", got)
	}
	if strings.Contains(got, "demo") {
		t.Errorf("version should not include demo tag, got %q", got)
	}
}
