<!-- TITLE -->
<h1 align="center" color="black">Golang API</h1>

<!-- THUMB -->
<p align="center">
        <img src="./doc_thumb.png" width="250px" style="box-shadow: 1px 2px 4px gray;" alt="Logo do Projeto" object-fit="cover">
</p>

<!-- STATUS -->
<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/skGab/Go-API.svg)](https://github.com/skGab/Go-API/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/skGab/Go-API.svg)](https://github.com/skGab/Go-API/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<!-- DESCRIPTION -->
<p align="center"> 
        💡 
         API escrita em Golang, utilizando ferramentas como Gin e GORM, oferece operações CRUD e autenticação de usuários. Além de seu propósito de aprendizado na construção de APIs, suporta um projeto pessoal, Coopers.
  <br> 
</p>

<!-- INTRO -->

## Índice

- [Tecnologias](#tecnologies)
- [Objetivo](#goal)
- [Funcionalidades](#features)
- [Requerimentos de qualidade](#quality)
- [Instruções de Uso](#glossary)
- [Autor](#authors)

## Tecnologias <a name="tecnologies"></a>

- Golang
- GORM
- GIN
- AIR (Hot reload)
- PostGresSQL

## Objetivo <a name="goal"></a>

O projeto tem como principal objetivo o estudo sobre a lingaguem Golang. Utilizando tecnologias como Gin e GORM, a API busca garantir um desempenho otimizado e uma manutenção simplificada, aproveitando as melhores práticas e padrões de código. Coopers, sendo o projeto subjacente, depende diretamente dessa API para gerenciar suas tarefas e autenticar seus usuários.

## Funcionalidades <a name="features"></a>

- Requisição de repositorios - API

## Requerimentos de qualidade <a name="quality"></a>

- Disponibilidade
- Sustentável

## Instruções de Uso <a name="glossary"></a>

- Certifique-se de ter o Golang instalado em seu sistema. Em seguida, execute o seguinte comando para instalar as dependências do projeto:

`go install`

<br>

- Para iniciar o servidor de desenvolvimento local, utilize o seguinte comando no root do projeto:

`go run`

O sistema estará disponível em http://localhost:8080.

<br>

- Para criar a versão final do projeto otimizada para produção, execute o seguinte comando:

`go build`

Os arquivos finais serão gerados em um executavel, "go-api.exe".
Para executar o arquivo, abra o terminal e rode o seguinte comando para iniciar o servidor:

`go-api.exe`

<br>

- Para manutenção de dependencias utilize o comando:

`go mod tidy`

## Autor <a name="authors"></a>

- [@Gabriel Assunção](https://github.com/skGab) - Ideia e Construção.
