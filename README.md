# Buscador de IP
 Aplicação de linha de comando feita em Go. 
 Recebe um endereço web, procura o IP desse endereço e depois devolve o nome do servidor ou IP (IPv4 e IPv6) desse endereço.

 ## Como rodar
 - Certifique-se de ter Go instalado em sua máquina
 - Clone o repositório
 - Navegue até a pasta do projeto
 - Execute o comando `go mod tidy` para instalar todas as dependências do projeto 
 - Execute o comando `go build -o main.exe`
 - Execute o comando `./main.exe --help` para verificar os comandos e funcionalidades disponíveis
 - `./main.exe ip --host=google.com` para buscar IPs
 - `./main.exe servidores --host=google.com` para buscar servidores
