# frubanagoChallenge
We have a catalog in a Json file with all varieties of Fruits and Vegetables (products.json)

• Run a small REST service enabling an endpoint "/ products".
Where all items will be sorted by name or by weight (weight)

• Enable another endpoint to our REST service: "/ orders" where you can select by Route, by Order ("orderID"), by quantity of products to be delivered, by order total (amount x price of the product).

# Software instructions

You can execute the solution over Docker or on your own Go environment:

## For Docker

* Download image

```
$ docker login (user/password of DockerHub)
$ docker pull jortizh1/frubanagochallenge:latest

```

* Run server Container
```
$ docker run -ti -d -p8585:8585 --name frubanagochallenge c197da6aeea6 /bin/bash

* Start Aplication
```
$ docker exec -ti frubanagochallenge ./home/GO/frubanagoChallenge/frubanagoChallenge

## REST services documentation

http://127.0.0.1:8585/api/products

//Parameters
Field string `json:"value" xml:"field"`

http://127.0.0.1:8585/api/orders

//Parameters
Field string `json:"field" xml:"field"`
Value int    `json:"value" xml:"value"`
