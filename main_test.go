package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

const (
	rootCmdMsg  = "A client for developping best practice of how to load config"
	serveCmdMsg = "serve called\n"
)

// Capturing Stdout/err and return as string. thread unsafe
func captureOutput(f func()) (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("creating pipe: %s", err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	os.Stdout = w
	os.Stderr = w
	defer func() {
		// cleanup
		os.Stdout = stdout
		os.Stderr = stderr
		if err := r.Close(); err != nil {
			log.Errorf("%s\n", err)
		}
	}()
	// call function
	f()
	// copy stdout
	var buf bytes.Buffer
	// have to close write FD before reading (otherwise pipe continues)
	if err := w.Close(); err != nil {
		return "", fmt.Errorf("closing pipe: %s", err)
	}
	// read bytes written to pipe
	if _, err := io.Copy(&buf, r); err != nil {
		return "", fmt.Errorf("copying buffer: %s", err)
	}
	return buf.String(), nil
}

func TestMain(t *testing.T) {
	// setup & teardown
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
	}()
	testServeCommand(t)
	testServeCommandFlags(t)
}

func testServeCommand(t *testing.T) {
	t.Helper()
	os.Args = []string{"command", "serve"}
	actual, err := captureOutput(main)
	if err != nil {
		t.Fatalf("Error should not occurred: %s\n", err)
	}
	expect := serveCmdMsg
	if actual != expect {
		t.Fatalf("Expect: %s, Actual: %s\n", expect, actual)
	}
}

func testServeCommandFlags(t *testing.T) {
	t.Helper()
	os.Args = []string{"command", "serve",
		"--port", "9090",
		"--github-user", "test",
		"--github-secret", "testpass",
	}
	actual, err := captureOutput(main)
	if err != nil {
		t.Fatalf("Error should not occurred: %s\n", err)
	}
	expect := serveCmdMsg
	if actual != expect {
		t.Fatalf("Expect: %s, Actual: %s\n", expect, actual)
	}
	port := viper.GetInt("server.port")
	if port != 9090 {
		t.Fatalf("Expect: 9090, Actual: %d\n", port)
	}
	user := viper.GetString("github.user")
	if user != "test" {
		t.Fatalf("Expect: test, Actual: %s\n", user)
	}
	secret := viper.GetString("github.secret")
	if secret != "testpass" {
		t.Fatalf("Expect: testpass, Actual: %s\n", secret)
	}
}

// TODO: implemented
func testServeCommandEnvs(t *testing.T) {
	t.Helper()
	os.Args = []string{"command", "serve"}
	actual, err := captureOutput(main)
	if err != nil {
		t.Fatalf("Error should not occurred: %s\n", err)
	}
	expect := serveCmdMsg
	if actual != expect {
		t.Fatalf("Expect: %s, Actual: %s\n", expect, actual)
	}
}

// TODO: implemented
func testServeCommandConfigFile(t *testing.T) {
	t.Helper()
	os.Args = []string{"command", "serve", "--config", "./test.yaml"}
	actual, err := captureOutput(main)
	if err != nil {
		t.Fatalf("Error should not occurred: %s\n", err)
	}
	expect := serveCmdMsg
	if actual != expect {
		t.Fatalf("Expect: %s, Actual: %s\n", expect, actual)
	}
}
