## Arquitetura geral do sistema
![](./Pastedimage20231020211732.png)
- Teremos o Next com BFF dos bancos, o correntista faz manipulações que resultam em chamada http
- Irá para o Nest.js que salvará em um Banco de dados PostgreSQL
- Quando criarmos e consultamos chaves PIX, faremos via gRPC ao microsserviço Golang(Banco Central)
- E quando um correntista quiser fazer uma transação, usaremos o Kafka

## Casos de uso

## 1 - Criar chaves pix

### Caso de uso
![](./Pastedimage20231020212101.png)

### Fluxo das operações

![](./Pasted%20image%2020231020212349.png)

## 2 - Consultar minhas chaves pix

### Caso de uso
![](./Pasted%20image%2020231020212432.png)

#### Fluxo das operações
![](./Pasted%20image%2020231020212507.png)

## 3 - Criar transferência

### Caso de uso
![](./Pasted%20image%20231020212807.png)

### Fluxo das operações - Iniciar transferência
![](./Pasted%20image%2020231020213018.png)

### Fluxo das operações - Notificar transferência para outro banco (alvo)
![](./Pasted%20image%2020231020213202.png)

### Fluxo das operações - Notificar transferência para banco origem
![](./Pasted%20image%2020231020213441.png)

## 4 - Listar minhas transações

### Caso de uso
![](./Pasted%20image%2020231020213548.png)

### Fluxo de operações
![](./Pasted%20image%2020231020213626.png)
