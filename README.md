# Telezap

Aplicação web desenvolvida utilizando Go, React e Postgresql.

## Linguagens de Programação - 2019/1

Trabalho final da disciplina Linguagens de Programação do período 2019/1.

### Professor

*Marcelo Ladeira*.

### Alunos

*Yuri Serka*;

*Henrique Mendes*;

*Henrique Lopes*;

*Gabriel Alves*.

## Instruções de uso

Execute:

> `$cd ./front`
>
> `$npm install`
>
> `$yarn`
>
> `$yarn build`

Isso fará com que instale as dependências para o front funcionar:

> `$npm start`

E rodar o client-side, será executada na porta 3000, porém existe o proxy que redireciona para o servidor.

Essa configuração se encontra no _package.json_

Para rodar o servidor precisa abrir outro terminal na pasta raíz e executar:

> `$go build`
>
> ./lpzin

O servidor irá escutará na porta 8080.
