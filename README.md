# Desafio de Cotação do Dólar

## Pré-requisitos

Certifique-se de ter o ambiente Go configurado corretamente antes de iniciar.

## Estilo 1 - Clean Architecture:

Na pasta "style-1", você encontrará uma implementação cuidadosamente estruturada seguindo os princípios da Clean Architecture. Essa abordagem enfatiza a separação clara de responsabilidades, promovendo a manutenibilidade, testabilidade e escalabilidade do código. O projeto está organizado em camadas distintas, como Entidades, Casos de Uso, Adaptadores e Frameworks Externos. As dependências são gerenciadas de forma a garantir uma arquitetura independente de frameworks, facilitando a substituição de componentes sem afetar o restante do sistema.

## Estilo 2 - Resolução do Desafio:

Na pasta "style-2", você encontrará a resolução direta do desafio de programação conforme as especificações fornecidas. Aqui, o foco está na entrega eficiente e direta da solução solicitada. O código pode ser estruturado de maneira mais tradicional, com a ênfase na lógica de negócios e na implementação dos requisitos do desafio. Este estilo pode ser mais pragmático e direto, proporcionando uma solução funcional e eficaz. Gostaria de ouvir sua opinião sobre a eficácia da resolução em atender aos requisitos do desafio, se os resultados estão corretos e se há algum aspecto que poderia ser melhorado em termos de clareza, eficiência ou boas práticas de programação.

## Passos para Execução - Estilo 2 - Resolução do Desafio:

1. **Clone o repositório:**

    ```bash
    git clone https://github.com/booscaaa/desafio-client-server-api-go-expert-pos.git
    ```

2. **Acesse o diretório do desafio:**

    ```bash
    cd desafio-client-server-api-go-expert-pos
    ```

3. **Execute o server.go:**
    - **Instale as dependências necessárias:**
        ```bash
        cd style-2/server
        go mod tidy

        go run server.go

        # O servidor estará disponível em http://localhost:8080.
        # O banco de dados estará disponível em ./desafio-client-server.db
        ```

3. **Execute o client.go:**
    - **Instale as dependências necessárias:**
        ```bash
        cd style-2/client
        go mod tidy

        go run client.go

        # O cliente realizará uma requisição ao servidor e salvará a cotação atual em um arquivo chamado cotacao.txt no formato: Dólar: {valor}.
        ```

## Passos para Execução - Estilo 1 - Clean Architecture:

1. **Clone o repositório:**

    ```bash
    git clone https://github.com/booscaaa/desafio-client-server-api-go-expert-pos.git
    ```

2. **Acesse o diretório do desafio:**

    ```bash
    cd desafio-client-server-api-go-expert-pos
    ```

3. **Execute o server.go:**
    - **Certifique-se que o docker está corretamente instalado**
    - **Rode os serviços:**
        ```bash
        cd style-1
        docker compose up --build -d

        # O servidor estará disponível em http://localhost:8080.
        # O banco de dados estará disponível em .data/desafio-client-server.db
        ```
    - **Acompanhe os logs:**
        ```bash
        docker logs server -f
        ```

3. **Execute o client:**
    - **Com o container rodando:**
        ```bash
        cd style-1
        docker exec -it server ./tmp/main client

        # O cliente realizará uma requisição ao servidor e salvará a cotação atual em um arquivo chamado cotacao.txt no formato: Dólar: {valor} na pasta ./data.
        ```
