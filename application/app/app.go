package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// GenerateCommandLineApp devolve a aplicação de linha de comando pronta para ser executada.
func GenerateCommandLineApp() *cli.App {
	app := cli.NewApp()   // instância para gerar nova CLI app
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca de IPs e nomes de servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
			Usage: "Host a ser consultado (por exemplo, devbook.com.br)",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIPs(c *cli.Context) error {
	host := c.String("host")
	if host == "" {
		return fmt.Errorf("nenhum host informado")
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return fmt.Errorf("buscar IPs do host %q: %w", host, err)
	}

	for _, ip := range ips {
		fmt.Println(ip.String())
	}

	return nil
}

func searchServers(c *cli.Context) error {
	host := c.String("host")
	if host == "" {
		return fmt.Errorf("nenhum host informado")
	}

	servers, err := net.LookupNS(host)
	if err != nil {
		return fmt.Errorf("buscar servidores do host %q: %w", host, err)
	}

	for _, servers := range servers {
		fmt.Println(servers.Host)
	}

	return nil
}
