## gRPC
-  É um framework desenvolvido pela google que tem o objetivo de facilitar o processo de comunicação entre sistemas de uma forma extremamente rápida,leve,independente de linguagem.
- É utilizada como uma das fontes primárias de comunicação da Google
- Faz parte da CNCF(Cloud Native Computation Foundation)
### Casos de uso
- Ideal para microsserviços
- Mobile, Browsers e Backend
- Geração das bibliotecas de forma automática com stubs(códigos prontos)
- Streaming bidirecional utilizando HTTP/2
	- Enviar e receber dados de forma contínua sem precisar de request e response(Protobuffers)
-  Usar gRPC para fazer comunicação entre browser e servidor, é uma tecnologia imatura ainda.
- gRPC é utilizada para comunicação entre sistemas
### Linguagens(Suporte oficial)
- gRPC-GO
- gRPC-JAVA
- gRPC-C
	- C++
	- Python
	- Ruby
	- Objective C
	- PHP
	- C#
	- Node.js
	- Dart
	- Kotlin/JVM
### RPC - Remote Procedure Call
![Pasted image 20231017134037.png](./Pasted%20image%2020231017134037.png)
#### Protocol Buffers
- Procol buffers são uma linguagem e plataforma neutras criada pelo Google, extensível e possui uma forma de você conseguir criar uma estrutura de dados e serializar esse tipo de dado(Pense em um XML, mas menor,mais rápido e mais simple)
#### Protocol Buffers vs JSON
- Arquivos binários < JSON
- Processo de serialização é mais leve(CPU) do que JSON
- Gasta menos recursos de rede
- Processo é mais rápido, mais leve
#### Exemplo de arquivo procol buffer
```
syntax = "proto3";

message SearchRequest {
	string query = 1;
	int32 page_number = 2;
	int32 result_per_page = 3;
}
```
`
### HTTP/2
- Nome original criado pela Google era SPDY
- Lançado em 2015
- Dados trafegados são binários e não texto como no HTTP 1.1
- Utiliza a mesma conexão TCP para enviar e receber os dados do cliente e do servidor(Multiplex)
- Server Push
- Headers são comprimidos
- Gasta menos recurso de rede
- Processo é mais veloz
### gRPC - API "unary"
![Pasted image 20231017135328.png](./Pasted%20image%2020231017135328.png)
### gRPC - API "Server streaming"
![Pasted image 20231017135423.png](./Pasted%20image%2020231017135423.png)
- No Server streaming, a response é contínua, ele não espera acabar e ficar pronta.
### gRPC - "Client streaming"
![Pasted image 20231017135610.png](./Pasted%20image%2020231017135610.png)
- No client streaming, as requests são enviadas aos poucos, como streaming
### gRPC - API "Bi direcional streaming"
![Pasted image 20231017135701.png](./Pasted%20image%2020231017135701.png)
- Tanto as requests, quanto as responses são enviadas por streaming

### REST vs gRPC
##### REST
- Texto / JSON
- Unidirecional
- Alta latência
- Sem contrato (maior chance de erros)
- Não tem suporte a streaming(Request/Response)
- Design pré-definido
- Bibliotecas de terceiro
##### gRPC
- Protocol Buffers
- Bidirecional e Assíncrono
- Baixa latência
- Contrato definido(.proto)
- Suporte a streaming
- Design é livre
- Geração de código(stubs)