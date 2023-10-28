# 💳 Codepix - Um simulador de Transações Pix

Codepix é um programa desenvolvido na "Imersão Full Stack && Full Cycle" da [Full Cycle](https://www.youtube.com/fullcycle)

<p align="center">
  <img src="documentation/home-frontend.png" alt="Bank Interface" />
</p>

<p align="center">
  <img alt="Node" src="https://img.shields.io/badge/-Node.js-44475a?logo=node.js&logoColor=white" />
  <img alt="Go" src="https://img.shields.io/badge/-Go-44475a?logo=go&logoColor=white" />
  <img alt="Next.js" src="https://img.shields.io/badge/-Next.js-44475a?logo=next.js&logoColor=white" />
  <img alt="NestJS" src="https://img.shields.io/badge/-NestJS-44475a?logo=nestjs&logoColor=white" />
  <img alt="PostgreSQL" src="https://img.shields.io/badge/-PostgreSQL-44475a?logo=postgresql&logoColor=white" />
  <img alt="Apache Kafka" src="https://img.shields.io/badge/-Apache Kafka-44475a?logo=apache-kafka&logoColor=white" />
  <img alt="Docker" src="https://img.shields.io/badge/-Docker-44475a?logo=docker&logoColor=white" />
</p>


## Como rodar o projeto

### Como rodar os serviços gRPC e Kafka

```
docker compose up -d
```
```
docker exec -it codepix-app-1 bash 
```
```
cd codepix-go
```
```
go run main.go all
```



## 2 Pilares de um Full Cycle Developer
- Operate what you build 
	- O Desenvolvedor fica responsável desde a concepção até o deploy do projeto.
- Ferramentas para escalar
	- O Desenvolvedor precisa dominar ferramentas que auxiliaram no processo de escala e Deploy contínuo como CI/CD, Docker, Kubernetes

## Conteúdos essenciais para ser um Full Cycle

## Estudo de caso CodePix
- É uma solução para simularmos transferências de valores em bancos fictícios através de chaves(email, cpf)
- Simularemos diversos bancos e contas bancárias que possuem uma chave PIX atribuída
- Cada conta poderá cadastrar suas chaves pix
- Uma conta poderá realizar uma transferência para outro banco utilizando a chave Pix de destino
- Uma transação não pode ser perdida mesmo que o CodePix esteja fora do ar ou que o Banco de destino esteja fora do ar
### Sobre os Bancos
- O banco será um microsserviço com funções limitadas a cadastro de contas e chaves Pix, bem como transferência de valores.
- Utilizaremos a mesma aplicação para simularmos diversos bancos, mudando apenas cores, nome e código.
- **Nest.js no backend**
- **Next.js no frontend**
	- Iremos mudar cada comportamento de banco baseado em variáveis de ambiente(Css, logos,etc)
	
## Sobre o CodePix
Fluxo básico de trabalho
- O microsserviço CodePix será responsável por intermediar as transferências bancárias
- Receberá a transação de transferência
- Encaminhará a transação para o banco de destino(Status: "pending")
- Recebe a confirmação do banco de destino (Status: "confirmed")
- Envia confirmação para o banco de origem informando quando o banco de destino processou
- Recebe a confirmação do banco de origem de que ele processou (Status: "completed")
- Marca a transação como completa.(Status: "completed")
### Cadastro e consulta de chave Pix
![Cadastro e consulta de chave Pix](./documentation/aula1/Pasted%20image%2020231016103143.png)
- O CodePix precisa armazenar TODAS as chaves pix
- Cada banco precisa armazenar apenas suas chaves pix

### Dinâmica do processo
![Dinâmica do processo](./documentation/aula1/Pasted%20image%2020231016104330.png)

### Principais desafios
- Comunicação rápida e eficiente(REST não é a opção mais rápida)
- Criação e consulta instantânea  das chaves (Síncrona)
- Garantia de que nenhuma transação seja perdida, mesmo que qualquer dos 3 sistemas esteja fora(Assíncrona)
- Iremos utilizar 2 tecnologias para comunicação
	- GRPC(Framework que irá trabalha com HTTP2 e protocol buffers em binário,velocidade)
	- Apache Kafka(Faz processamento de dados, consome e produz, sistema de stream de dados e filas)

## CodePix
- Será capaz de atuar como um servidor gRPC
- Consumir e publicar mensagens no Apache Kafka, garantindo resiliência para que as mensagens não sejam perdidas
- Ambas operações devem ser realizadas de forma simultânea ao executar o serviço
- Trabalhar com um design focado em só solucionar o problema do domínio(DDD)
- Teremos uma camada de domínio focada em regras de negócio
- Deixar a complexidade técnica para "camada de aplicação", responsável pelo servidor gRPC e Kafka
- Deixar a aplicação desacoplada e flexível para a implementação de outros formatos de comunicação, como API Rest, CLI clients, etc. **SEM** alterar nenhum outro componente da aplicação ou modelo de domínio
### Arquitetura de software
- Arquitetura é uma forma de você conseguir usar padrões já existentes, para criar algo sem desperdício, e que esse algo se pague durante o tempo, sendo sustentável
- Crescimento sustentável... e isso não é para amador
- Manutenção...  é óbvio, mas é complexo
- A complexidade inicial se paga ao longo do projeto
- O seu software deve ser definido e desenhado por você e não pelo seu framework
- As peças têm que se encaixar, mas eventualmente podem ser substituídas ou adicionadas
### Arquitetura Hexagonal / Ports and Adapters
- Não tem nada haver com um hexágono
- Criada pelo Dr. Alistair Cockburn
- Muito conhecido por ser um dos assinantes do Manifesto Agile
- Ports and Adapters - 2005
- Temos a aplicação no meio e adaptadores em volta da aplicação que conseguem se comunicar com ela
 ![Arquitetura Hexagonal](./documentation/aula1/Pasted%20image%2020231016111022.png)
### Onion Architecture
- Criada por Jeffrey Palermo em 2008
- Ports and Adapters mais detalhada
  ![Onion Architecture](./documentation/aula1/Pasted%20image%2020231016111811.png)
## Clean Architecture
- Criada por Robert C. Martin aka Uncle Bob, também criador do SOLID
- Famoso pelo conceito de Clean Code
  ![Clean Architecture](./documentation/aula1/Pasted%20image%2020231016112154.png)
## Estrutura e camadas do CodePix
  ![Estrutura e camadas do CodePix](./documentation/aula1/Pasted%20image%2020231016114129.png)

## Recursos a serem utilizados
- Docker
- Golang
- Apache Kafka
- Postgres