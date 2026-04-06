package ssh

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"platoon-go/internal/config"
	"strconv"

	"golang.org/x/crypto/ssh"
)

func RunShell(target *config.TargetConfig, command string) ([]byte, error) {
	return exec.Command("ssh", target.Username+"@"+target.Host, "-p"+strconv.Itoa(target.Port), command).CombinedOutput()
}

func Run(target *config.TargetConfig, keyPath string, command string) {
	config := &ssh.ClientConfig{
		User: target.Username,
		Auth: []ssh.AuthMethod{
			publicKeyAuth(keyPath),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // not in prod!
	}

	client, err := ssh.Dial("tcp", target.Host, config)
	if err != nil {
		log.Fatalf("failed to connect to the server: %v", err)
	}
	defer client.Close()

	output, err := runCommand(client, command)
	if err != nil {
		log.Fatalf("failed to run command: %v", err)
	}
	fmt.Println(string(output))
}

func runCommand(client *ssh.Client, command string) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	return session.CombinedOutput(command)
}

func publicKeyAuth(keyPath string) ssh.AuthMethod {
	key, err := os.ReadFile(keyPath)

	if err != nil {
		log.Fatalf("failed to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)

	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}

	return ssh.PublicKeys(signer)
}
