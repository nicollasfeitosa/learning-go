package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executado
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Go Discover"
	app.Usage = "Busca IPs e hostnames de servidores na internet"
	
	flags := []cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "nicollas.dev",
		},
	}

	app.Commands = []cli.Command { // Slice de structures do tipo cli.Command
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "server",
			Usage: "Busca o nome de servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context)  {
	host := c.String("host")

	fmt.Printf("IPs do host: %s\n", host);

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip);
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) // name-servers
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host);
	}
}