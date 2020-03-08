package swarm

import (
	"bytes"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

// GetNodeCount gets the docker swarm nodes count
func GetNodeCount() int {
	//create command
	ncCmd := exec.Command("docker", "node", "ls")
	wcCmd := exec.Command("wc", "-l")

	//make a pipe
	reader, writer := io.Pipe()
	var buf bytes.Buffer

	//set the output of "node ls" command to pipe writer
	ncCmd.Stdout = writer
	//set the input of the "wc" command pipe reader

	wcCmd.Stdin = reader

	//cache the output of "wc" to memory
	wcCmd.Stdout = &buf

	//start to execute "node ls" command
	ncCmd.Start()

	//start to execute "wc" command
	wcCmd.Start()

	ncCmd.Wait()
	writer.Close()

	wcCmd.Wait()
	reader.Close()

	strStdout := string(buf.Bytes())
	firstLine := strings.Split(strStdout, "\n")[0]
	res, err := strconv.Atoi(firstLine)

	if err != nil {
		panic(err)
	}

	// return number of available workers
	// minus 1 bc ignoring header
	if res > 0 {
		res--
	}

	return res
}