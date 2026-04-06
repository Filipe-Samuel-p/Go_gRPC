# Estudo de gRPC com Golang

Este repositório contém um estudo sobre gRPC (Google Remote Procedure Call) usando Golang, explorando seus diferentes padrões de interação de API.

## O que é gRPC?

gRPC é um framework de RPC (Remote Procedure Call) moderno, de alta performance, de código aberto e universal que pode ser executado em qualquer ambiente. Ele permite que os serviços se comuniquem entre si de forma eficiente, sendo frequentemente usado para conectar serviços em arquiteturas de microsserviços, conectar dispositivos móveis e navegadores a serviços de backend.

**Principais características e vantagens:**
*   **Performance:** Construído sobre HTTP/2 para transporte, Protocol Buffers (protobuf) como linguagem de descrição de interface e serialização binária. Isso resulta em comunicação altamente eficiente.

*   **Independente de Linguagem:** Suporta várias linguagens através da geração de código, permitindo que clientes e servidores escritos em diferentes linguagens se comuniquem perfeitamente.
*   **Tipagem Forte:** Protocol Buffers impõem um contrato rigoroso entre cliente e servidor, reduzindo erros em tempo de execução e melhorando a manutenibilidade.
*   **Interoperabilidade:** Projetado para comunicação entre plataformas e entre linguagens.

**Quando usar gRPC:**
*   **Microsserviços:** Ideal para comunicação interserviço de alta performance dentro de uma arquitetura de microsserviços.
*   **Serviços de tempo real ponto a ponto:** Onde baixa latência e alta taxa de transferência são críticas.
*   **Ambientes poliglotas:** Quando você tem serviços escritos em diferentes linguagens de programação.
*   **Móvel e IoT:** Comunicação eficiente para dispositivos com recursos limitados.

## Tipos de API gRPC

gRPC suporta quatro tipos principais de métodos de serviço, fornecendo diferentes padrões de comunicação entre o cliente e o servidor:

### 1. RPC Unário

Este é o tipo mais simples de RPC, onde o cliente envia uma única requisição ao servidor e recebe uma única resposta de volta. É semelhante a um modelo tradicional de requisição/resposta HTTP.

### 2. RPC de Streaming do Servidor

Em um RPC de streaming do servidor, o cliente envia uma única requisição ao servidor, e o servidor envia de volta uma sequência de respostas. Após enviar todas as respostas, o servidor fecha o stream. O cliente lê do stream até que não haja mais mensagens.

### 3. RPC de Streaming do Cliente

Em um RPC de streaming do cliente, o cliente envia uma sequência de mensagens ao servidor. Uma vez que o cliente terminou de escrever suas mensagens, ele espera que o servidor envie de volta uma única resposta.

### 4. RPC de Streaming Bidirecional

Em um RPC de streaming bidirecional, tanto o cliente quanto o servidor enviam uma sequência de mensagens um para o outro usando um stream de leitura e escrita. Os dois streams operam independentemente, então clientes e servidores podem ler e escrever em qualquer ordem.

## Visão Geral da Estrutura do Projeto

Este repositório está organizado em diretórios, cada um demonstrando um tipo específico de API gRPC implementado em Golang.

### `API_Unary/`

Este diretório contém um exemplo de **RPC Unário**. Ele geralmente demonstra um padrão básico de requisição-resposta, onde um cliente envia uma única mensagem (por exemplo, solicitando detalhes de um produto) e o servidor responde com um único resultado.

### `API_Server_Streaming/`

Este diretório apresenta um **RPC de Streaming do Servidor**. Aqui, um cliente envia uma única requisição, e o servidor transmite múltiplas respostas de volta (por exemplo, transmitindo uma lista de departamentos).

### `API_Client_Streaming/`

Este diretório fornece um exemplo de **RPC de Streaming do Cliente**. O cliente envia um fluxo de mensagens para o servidor (por exemplo, enviando múltiplas requisições de cálculo), e o servidor as processa para retornar uma única resposta agregada.

### `API_Bidirectional_Streaming/`

Este diretório demonstra um **RPC de Streaming Bidirecional**. Tanto o cliente quanto o servidor trocam fluxos de mensagens de forma independente e concorrente (por exemplo, uma atualização em tempo real de um carrinho de compras onde cliente e servidor podem enviar atualizações a qualquer momento).

### `Protocol_Buffer/`

Este diretório contém um exemplo básico de uso de Protocol Buffers, separado dos tipos específicos de API gRPC. Ele mostra como os arquivos `.proto` são definidos e como são compilados para código Go (`.pb.go`) para serialização e desserialização de dados, o que é fundamental para o gRPC.
