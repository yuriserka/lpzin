# lpzin

Pra rodar primeiro tem que clonar o repo e executar:

### `git submodule update --recursive --remote`
### `cd ./lpzin_frontend`
### `npm install`

isso fará com que instale as dependencia pro front funcionar

### `npm start`

Roda o client-side, irá rodar na porta 3000, mas tem o proxy que redireciona para o servidor.
essa config ta no package.json

# Para rodar o servidor precisa abrir outro terminal na pasta raíz e executar:

### `go build`
### ./lpzin

O servidor irá rodar na porta 8080.