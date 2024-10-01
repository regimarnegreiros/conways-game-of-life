# Jogo da Vida - John Conway

## Descrição

Este é um programa simples escrito em Go que simula o jogo da vida de Conway, é o exemplo mais bem conhecido de autômato celular. O programa cria uma grade bidimensional de células que podem estar vivas ou mortas, e essas células evoluem em cada iteração de acordo com um conjunto de regras.

## Evolução das Células

Em cada iteração, o estado de cada célula é atualizado com base no número de vizinhos vivos, seguindo as regras do jogo da vida:
- Toda célula morta com exatamente três vizinhos vivos torna-se viva (nascimento).
- Toda célula viva com menos de dois vizinhos vivos morre por isolamento.
- Toda célula viva com mais de três vizinhos vivos morre por superpopulação.
- Toda célula viva com dois ou três vizinhos vivos permanece viva.

## Como Executar

1. Clone o repositório ou copie o código para o seu ambiente local.
2. Compile o programa usando o seguinte comando no terminal:

   ```bash
   go build main.go
   ```

3. Execute o programa:

   ```bash
   ./main
   ```
4. Ou execute diretamente com:
    ```bash
    go run main.go
    ```

A simulação será exibida no terminal e será atualizada a cada `FRAMETIME` milissegundos.

## Parâmetros Personalizáveis

- `HEIGHT`: Altura da grade.
- `WIDTH`: Largura da grade.
- `FRAMETIME`: Intervalo de tempo entre as atualizações (em milissegundos).

## Referências

Este programa é baseado no Jogo da Vida, criado pelo matemático John Conway. Para mais detalhes sobre o jogo, consulte [Jogo da Vida na Wikipedia](https://pt.wikipedia.org/wiki/Jogo_da_vida).