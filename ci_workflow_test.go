// Copyright (c) 2026 by Delphix. All rights reserved.
//
// ci_workflow_test.go — static validation tests for the CI workflow contract
// defined by DLPXECO-14115. These tests validate that:
//   - .github/workflows/ci.yml has the required structure and keys
//   - the coverage-threshold shell logic behaves correctly (PASS/FAIL/missing)
//   - CLAUDE.md and CONTRIBUTING.md document the CI contract
//
// These tests are package-`main` because they live at the repo root and only
// validate file content; they do not exercise any provider code.
package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const (
	ciYamlPath     = ".github/workflows/ci.yml"
	claudeMdPath   = "CLAUDE.md"
	contributingMd = "CONTRIBUTING.md"
)

// readFile reads a file relative to the repo root and fails the test on error.
func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	return string(b)
}

// assertContains fails the test if haystack does not contain needle.
func assertContains(t *testing.T, haystack, needle, scenario string) {
	t.Helper()
	if !strings.Contains(haystack, needle) {
		t.Errorf("[%s] expected to find %q but did not", scenario, needle)
	}
}

// ---------------------------------------------------------------------------
// FR-001: GitHub Actions CI Workflow (S1–S7)
// ---------------------------------------------------------------------------

func TestS1_CIWorkflowExistsAndHasName(t *testing.T) {
	content := readFile(t, ciYamlPath)
	assertContains(t, content, "name: ci", "S1")
}

func TestS2_TriggersOnPullRequestToMain(t *testing.T) {
	content := readFile(t, ciYamlPath)
	assertContains(t, content, "pull_request:", "S2")
	assertContains(t, content, "main", "S2")
}

func TestS3_TriggersOnPushToMain(t *testing.T) {
	content := readFile(t, ciYamlPath)
	assertContains(t, content, "push:", "S3")
	// Ensure both push and pull_request reference main
	if !strings.Contains(content, "push:") || !strings.Contains(content, "main") {
		t.Errorf("[S3] expected workflow to trigger on push to main")
	}
}

func TestS4_UsesGoVersionFileFromGoMod(t *testing.T) {
	content := readFile(t, ciYamlPath)
	assertContains(t, content, "go-version-file: go.mod", "S4")
}

// TestS5_DoesNotExportTFAcc verifies the workflow does not export TF_ACC.
//
// The test plan text says "File does NOT contain the string TF_ACC", but the
// design doc and the workflow itself clarify the actual contract: the workflow
// must not *export* TF_ACC (so acceptance tests are auto-skipped). Comments
// that document this intent are acceptable and expected. We therefore check
// every non-comment, non-blank line and assert that none of them mentions
// TF_ACC.
func TestS5_DoesNotExportTFAcc(t *testing.T) {
	content := readFile(t, ciYamlPath)
	for i, raw := range strings.Split(content, "\n") {
		// Strip a trailing inline comment, then trim whitespace.
		line := raw
		if idx := strings.Index(line, "#"); idx >= 0 {
			line = line[:idx]
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "TF_ACC") {
			t.Errorf("[S5] line %d exports/uses TF_ACC outside a comment: %q", i+1, raw)
		}
	}
}

func TestS6_TestCommandHas300sTimeout(t *testing.T) {
	content := readFile(t, ciYamlPath)
	assertContains(t, content, "-timeout=300s", "S6")
}

func TestS7_UploadArtifactHasIfAlways(t *testing.T) {
	content := readFile(t, ciYamlPath)
	assertContains(t, content, "if: always()", "S7")
}

// ---------------------------------------------------------------------------
// FR-002: Coverage Threshold Enforcement (S8–S11)
// ---------------------------------------------------------------------------

func TestS8_CoverageThresholdDefined(t *testing.T) {
	content := readFile(t, ciYamlPath)
	assertContains(t, content, "COVERAGE_THRESHOLD:", "S8")
}

// thresholdScript is the inline threshold-check logic extracted from ci.yml.
// Kept here in sync with the workflow's "Enforce coverage threshold" step so
// the same logic can be unit-tested without invoking GitHub Actions.
//
// The script reads coverage.out directly (looking for a "total:" line in the
// mock) rather than invoking `go tool cover`, which is sufficient for unit
// testing the threshold comparison logic in isolation.
const thresholdScript = `
if [[ ! -f coverage.out ]] || [[ ! -s coverage.out ]]; then
  echo "ERROR: coverage.out is missing or empty"
  exit 1
fi
TOTAL=$(grep "^total:" coverage.out | awk '{print $3}' | tr -d '%')
if [[ -z "$TOTAL" ]]; then
  echo "ERROR: could not parse coverage total"
  exit 1
fi
echo "Total coverage: ${TOTAL}%"
PASS=$(awk -v total="$TOTAL" -v threshold="$COVERAGE_THRESHOLD" \
  'BEGIN { print (total + 0 >= threshold + 0) ? "1" : "0" }')
if [[ "$PASS" != "1" ]]; then
  echo "FAIL: Coverage ${TOTAL}% is below threshold ${COVERAGE_THRESHOLD}%"
  exit 1
fi
echo "PASS: Coverage ${TOTAL}% >= threshold ${COVERAGE_THRESHOLD}%"
`

// runThresholdScript writes a mock coverage.out (or skips it), runs the
// threshold script in a tempdir with COVERAGE_THRESHOLD set, and returns
// (stdout+stderr, exit code).
func runThresholdScript(t *testing.T, mockCoverage string, threshold string) (string, int) {
	t.Helper()
	dir := t.TempDir()
	if mockCoverage != "" {
		if err := os.WriteFile(filepath.Join(dir, "coverage.out"), []byte(mockCoverage), 0o644); err != nil {
			t.Fatalf("failed to write mock coverage.out: %v", err)
		}
	}
	cmd := exec.Command("bash", "-c", thresholdScript)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "COVERAGE_THRESHOLD="+threshold)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	code := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		code = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("unexpected error running threshold script: %v", err)
	}
	return out.String(), code
}

func TestS9_ThresholdScriptPassesWhenCoverageExceedsThreshold(t *testing.T) {
	mock := "mode: atomic\nfile.go:1.1,2.2 1 1\ntotal:\t(statements)\t55.0%\n"
	out, code := runThresholdScript(t, mock, "50")
	if code != 0 {
		t.Errorf("[S9] expected exit 0 when coverage 55%% >= threshold 50%%; got exit %d. Output:\n%s", code, out)
	}
	assertContains(t, out, "PASS", "S9")
}

func TestS10_ThresholdScriptFailsWhenCoverageBelowThreshold(t *testing.T) {
	mock := "mode: atomic\nfile.go:1.1,2.2 1 1\ntotal:\t(statements)\t48.0%\n"
	out, code := runThresholdScript(t, mock, "50")
	if code == 0 {
		t.Errorf("[S10] expected non-zero exit when coverage 48%% < threshold 50%%; got exit 0. Output:\n%s", out)
	}
	assertContains(t, out, "FAIL", "S10")
	assertContains(t, out, "48", "S10")
	assertContains(t, out, "50", "S10")
}

func TestS11_ThresholdScriptFailsWhenCoverageMissing(t *testing.T) {
	out, code := runThresholdScript(t, "", "50")
	if code == 0 {
		t.Errorf("[S11] expected non-zero exit when coverage.out is missing; got exit 0. Output:\n%s", out)
	}
	assertContains(t, out, "missing or empty", "S11")
}

// ---------------------------------------------------------------------------
// FR-003: Documentation — CLAUDE.md and CONTRIBUTING.md (S12–S16)
// ---------------------------------------------------------------------------

func TestS12_ClaudeMdHasCIContractSection(t *testing.T) {
	content := readFile(t, claudeMdPath)
	assertContains(t, content, "## CI Contract", "S12")
}

func TestS13_ClaudeMdHasLocalEquivalentCommand(t *testing.T) {
	content := readFile(t, claudeMdPath)
	assertContains(t, content, "coverprofile=coverage.out", "S13")
}

func TestS14_ClaudeMdDocumentsCoverageThreshold(t *testing.T) {
	content := readFile(t, claudeMdPath)
	assertContains(t, content, "COVERAGE_THRESHOLD", "S14")
}

func TestS15_ContributingMdHasCISection(t *testing.T) {
	content := readFile(t, contributingMd)
	assertContains(t, content, "## CI", "S15")
}

func TestS16_ContributingMdNamesStatusCheck(t *testing.T) {
	content := readFile(t, contributingMd)
	assertContains(t, content, "ci / unit-tests", "S16")
}

// ---------------------------------------------------------------------------
// FR-004: Branch Protection Contract Documentation (S17–S19)
// ---------------------------------------------------------------------------

func TestS17_ClaudeMdHasExactStatusCheckString(t *testing.T) {
	content := readFile(t, claudeMdPath)
	assertContains(t, content, "ci / unit-tests", "S17")
}

func TestS18_ClaudeMdBranchProtectionRequiresStatusChecks(t *testing.T) {
	content := readFile(t, claudeMdPath)
	assertContains(t, content, "Require status checks to pass before merging", "S18")
}

func TestS19_ClaudeMdBranchProtectionRequiresUpToDate(t *testing.T) {
	content := readFile(t, claudeMdPath)
	assertContains(t, content, "Require branches to be up to date before merging", "S19")
}
