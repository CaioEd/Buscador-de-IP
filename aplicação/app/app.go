package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Gerar devolve a aplicação de linha de comando pronta para ser executada.
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor"

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
			Action: buscarIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) error {
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

func buscarServidores(c *cli.Context) error {
	host := c.String("host")
	if host == "" {
		return fmt.Errorf("nenhum host informado")
	}

	servidores, err := net.LookupNS(host)
	if err != nil {
		return fmt.Errorf("buscar servidores do host %q: %w", host, err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

	return nil
}
