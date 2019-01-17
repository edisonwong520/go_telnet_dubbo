package main


import (
	"os"

	"github.com/edison/go-telnet/client"

)

type goTelnet struct{}

func newGoTelnet() *goTelnet {
	return new(goTelnet)
}

func (g *goTelnet) run() {
	telnetClient := g.createTelnetClient()
	cmd:=`invoke org.apache.dubbo.demo.DemoService.sayhello("hxx")`+"\n"
	telnetClient.ProcessData(cmd, os.Stdout)
}

func (g *goTelnet) createTelnetClient() *client.TelnetClient {
	host:="localhost"
	port:=20880
	telnetClient := client.NewTelnetClient(host,port)
	return telnetClient
}
func main() {
	app := newGoTelnet()
	app.run()
}
