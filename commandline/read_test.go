package commandline

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadRequiredArguments(t *testing.T) {
	// given
	host := "aHost"
	port := uint64(92832932)
	SetCommandLineArgs(host, fmt.Sprintf("%d", port))

	// when
	commandLine := Read()

	// then
	assert := assert.New(t)
	assert.Equal(host, commandLine.Host(), "Host read from command line is different than set earlier.")
	assert.Equal(port, commandLine.Port(), "Port read from command line is different than set earlier.")
}

func TestReadOptionalFlags(t *testing.T) {
	// given
	timeout := time.Duration(5 * time.Second)
	SetCommandLineArgs("-t", fmt.Sprintf("%v", timeout), "host", "123")

	// when
	commandLine := Read()

	// then
	assert := assert.New(t)
	assert.Equal(timeout, commandLine.Timeout(), "Timeout read from command line is different than set earlier.")
}
