openapi: 1.0.0
servers: []
info:
  description: 'Esta API esta diseñada para consultar cervezas'
  version: "1.0.0"
  title: 'API Beers'
  contact:
    email: 'yvillalonsilva@gmail.com '
tags:
  - name: cerveza
    description: rica cerveza..
paths:
  /beers:
    get:
      tags:
        - cerveza
      summary: Lista todas las cervezas
      operationId: searchBeers
      description: |
        Lista todas las cervezas que se encuentran en la base de datos
      responses:
        '200':
          description: 'Operacion exitosa'
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/BeerItem'
    post:
      tags:
        - cerveza
      summary: 'Ingresa una nueva cerveza'
      operationId: addBeers
      description: 'Ingresa una nueva cerveza'
      responses:
        '201':
          description: 'Cerveza creada'
        '400':
          description: 'Request invalida'
        '409':
          description: 'El ID de la cerveza ya existe'
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/BeerItem'
        description: 'Ingresa una nueva cerveza'
        
  /beers/{beerID}:
    get:
      tags:
        - cerveza
      summary: Lista el detalle de la marca de cervezas
      operationId: searchBeerById
      description: |
        Busca una cerveza por su Id
      parameters:
        - name: beerID
          in: path
          description: 'Busca una cerveza por su Id'
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: 'Operacion exitosa'
        '404':
          description: 'El Id de la cerveza no existe'
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/BeerItem'
            
  /beers/{beerID}/boxprice:
    get:
      tags:
        - cerveza
      summary: Lista el precio de una caja de cervezas de una marca
      operationId: boxBeerPriceById
      description: |
        Obtiene el precio de una caja de cerveza por su Id
      parameters:
        - name: beerID
          in: path
          description: 'Busca una cerveza por su Id'
          required: true
          schema:
            type: integer
        - in: query
          name: currency
          schema:
            type: string
          description: Tipo de moneda con la que pagará
        - in: query
          name: quantity
          schema:
            type: integer
            default: 6
          description: La cantidad de cervezas a comprar
      responses:
        '200':
          description: 'Operacion exitosa'
        '404':
          description: 'El Id de la cerveza no existe'
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/BeerBox'
components:
  schemas:
    BeerItem:
      type: object
      required:
        - Id
        - Name
        - Brewery
        - Country
        - Price
        - Currency
      properties:
        Id:
          type: integer
          example: 1
        Name:
          type: string
          example: 'Golden'
        Brewery:
          type: string
          example: 'Kross'
        Country:
          type: string
          example: 'Chile'
        Price:
          type: number
          example: 10.5
        Currency:
          type: string
          example: 'EUR'
    BeerBox:
      type: object
      properties:
        Price Total:
          type: number 
      
