# frete-rapido

Projeto em Golang de integração com API Frete Rápido

Foram utilizados os pacotes:
- Golang 1.20
- GoChi,
- PQ,
- Viper
- Docker

## configuração

Precisamos primeiro fazer a criação de dois arquivos diferentes de configurações antes de subir o docker.

O primeiro é o `.env` que é o responsável por toda a configuração do docker-compose. Para facilitar temos o `.env-example`. Execute o comando `mv env-example .env` no terminal e edite as variaveis de ambiente se precisar.

O segundo é o `config.toml` que é onde deixamos nossas variaveis de ambiente do código em si. Basta executar o comando `mv env-example .env` no terminal. Para esse caso precisamos sim editar as variaveis.

Um ponto importante é que os valores setadas em ambos os arquivos pro database devem ser os mesmos, se não, teremos problemas com a conexão.

As configurações de API no `.config.toml` é a api que estamos usando de terceiros. Então pra isso você precisa dos dados certinhos!

## hora de rodar :)

Pra rodar, bora executar o docker. Tanto o banco de dados quanto a nossa aplicação estão lá.

Execute os comando a seguir no terminal:

`docker-compose build`

`docker-compose up -d`

## 

Temos dois endpoints.

O primeiro é onde geramos a simulação de cotação de frete usando o método POST `/quote`.

Essa API recebe apenas o corpo a seguir:

```{
   "recipient":{
      "address":{
         "zipcode":"01311000"
      }
   },
   "volumes":[
      {
         "category":7,
         "amount":1,
         "unitary_weight":5,
         "price":349,
         "sku":"abc-teste-123",
         "height":0.2,
         "width":0.2,
         "length":0.2
      },
      {
         "category":7,
         "amount":2,
         "unitary_weight":4,
         "price":556,
         "sku":"abc-teste-527",
         "height":0.4,
         "width":0.6,
         "length":0.15
      }
   ]
}
```

Ela nos retorna uma lista com as transportadoras disponiveis. Segue exemplo de retorno.

```
{
   "recipient":{
      "address":{
         "zipcode":"01311000"
      }
   },
   "volumes":[
      {
         "category":7,
         "amount":1,
         "unitary_weight":5,
         "price":349,
         "sku":"abc-teste-123",
         "height":0.2,
         "width":0.2,
         "length":0.2
      },
      {
         "category":7,
         "amount":2,
         "unitary_weight":4,
         "price":556,
         "sku":"abc-teste-527",
         "height":0.4,
         "width":0.6,
         "length":0.15
      }
   ]
}

```

