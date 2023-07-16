# Monitor de Sites

<p align="center">
  <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" alt="Gopher" width="200" height="200">
</p>

Este é um projeto Go que implementa um monitor de sites. A aplicação permite monitorar vários sites para verificar se estão online, exibir os logs com o resultado da verificação e sair da aplicação.

## Funcionalidades

A aplicação possui as seguintes funcionalidades:

1. **Monitorar Sites**: Essa opção permite monitorar vários sites para verificar se estão online. A aplicação faz requisições HTTP para cada site e verifica se o código de resposta é 200 (OK). Caso o site esteja offline, essa informação será exibida no terminal.

2. **Exibir Último Log**: Essa opção exibe os logs da verificação dos sites, mostrando a situação atual de cada site e o horário em que foi realizada a verificação.

3. **Sair da Aplicação**: Essa opção permite encerrar a aplicação.

## Pré-requisitos

Certifique-se de ter o Go instalado em seu ambiente de desenvolvimento. Você pode verificar se o Go está instalado executando o seguinte comando:

```bash
go version
```

Se o Go estiver instalado corretamente, você verá a versão instalada sendo exibida.

## Dependências

Este projeto utiliza a biblioteca [github.com/fatih/color](https://github.com/fatih/color) para colorir a saída do terminal. Certifique-se de instalar a biblioteca antes de executar o projeto.

```bash
go get github.com/fatih/color
```

## Como Executar o Projeto

Siga as etapas abaixo para executar o projeto:

1. Clone este repositório para o seu ambiente local:

```bash
git clone https://github.com/Lukasdias/curso-introdutorio-golang-alura
```

2. Navegue para o diretório do projeto:

```bash
cd curso-introdutorio-golang-alura
```

3. Execute o comando `go run` para compilar e executar o projeto:

```bash
go run main.go
```

4. No menu exibido no terminal, escolha uma das opções disponíveis para realizar a ação desejada.

## Contribuição

Contribuições são bem-vindas! Se você quiser melhorar este projeto, fique à vontade para enviar pull requests.
