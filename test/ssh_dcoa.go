package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
)

// defaultKeyCandidates is the ordered list of private keys tried when
// DCOA_KEY_PATH is not set.
var defaultKeyCandidates = []string{"dcoa_ed25519", "id_ed25519", "id_rsa"}

// loadSigner resolves and parses the SSH private key used for authentication.
// It honors DCOA_KEY_PATH first, then falls back to common key names under
// ~/.ssh. The key never leaves the local machine.
func loadSigner() (ssh.Signer, error) {
	var candidates []string
	if p := os.Getenv("DCOA_KEY_PATH"); p != "" {
		candidates = []string{p}
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("cannot resolve home directory: %w", err)
		}
		for _, name := range defaultKeyCandidates {
			candidates = append(candidates, filepath.Join(home, ".ssh", name))
		}
	}

	for _, path := range candidates {
		key, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("failed to parse SSH private key %s (encrypted keys are not supported — use an unencrypted key or an ssh-agent-backed key): %w", path, err)
		}
		return signer, nil
	}

	return nil, fmt.Errorf("no usable SSH private key found (set DCOA_KEY_PATH or place a key at ~/.ssh/{%s})", defaultKeyCandidates[0])
}

func main() {
	host := os.Getenv("DCOA_HOST")
	user := os.Getenv("DCOA_USER")

	if host == "" || user == "" {
		log.Fatal("DCOA_HOST and DCOA_USER environment variables must be set")
	}

	signer, err := loadSigner()
	if err != nil {
		log.Fatalf("SSH key auth setup failed: %v", err)
	}

	addr := host
	if _, _, err := net.SplitHostPort(host); err != nil {
		addr = host + ":22"
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         15 * time.Second,
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalf("failed to dial %s: %v", addr, err)
	}
	defer client.Close()

	cmd := "hostname && uname -a"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("stdout pipe: %v", err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatalf("stderr pipe: %v", err)
	}

	if err := session.Start(cmd); err != nil {
		log.Fatalf("failed to start command: %v", err)
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	if err := session.Wait(); err != nil {
		log.Fatalf("command failed: %v", err)
	}
	fmt.Println("--- ssh session completed ---")
}
