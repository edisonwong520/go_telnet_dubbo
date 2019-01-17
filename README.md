# go-telnet-dubbo


### Status: **Done**


### Problem Definition

Give a ip , port ,method name and parameter, test whether the Dubbo service on the ip:port works or not

### Main code
run Dubbo command in telnet
~~~
func (g *goTelnet) run() {
	telnetClient := g.createTelnetClient()
	cmd:=`invoke org.apache.dubbo.demo.DemoService.sayhello("hxx")`+"\n"
	telnetClient.ProcessData(cmd, os.Stdout)
}
~~~

<br>
modify telnet ip and port  

~~~
func (g *goTelnet) createTelnetClient() *client.TelnetClient {
	host:="localhost"
	port:=20880
	telnetClient := client.NewTelnetClient(host,port)
	return telnetClient
}
~~~

### Usage ###
run command
~~~
go install && go-telnet
~~~
or
~~~
go run main.go
~~~

### Output ###
~~~
"hello,hxx!"
elapsed: 0 ms.
dubbo>
Process finished with exit code 0
~~~
