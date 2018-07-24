# menu-api

### About

*menu-api* is an example of a basic (GETs only atm) web server exposing a RESTful API mimicking an automated restaurant menu ordering system following a specific set of requirements.

The design of this repo is modular in nature, and consists of three main components:

1. cmd executables
   - this is where static binaries go. programs making use of repo packages live here
2. api package
   - the web server and router
3. api/storage
   - database driver subpackage
   - contains importable data types used in API server

**Note:** This is not a consumer/front facing API. It is meant to be used as a backend query engine of a front end user interface.

### 

### Usage

To use, first clone the github repo. 

**Requirements**

The only requirement is *docker* and *docker-compose*.

To install docker see here: https://docs.docker.com/install/

To install docker compose see here: https://docs.docker.com/compose/install/



**Running tests**

To run the unit and integration tests run the command: 

`docker-compose up test`

Everything (should) pass!

Then run:

`docker-compose down`

to tear down the containers.



**Running the web server**

To run the web server on your local machine run the command:

`docker-compose up web`

This launches an HTTP server listening on ***localhost:5000*** and inserts some mock data into the database.

When done run:

`docker-compose down`

to tear down the containers.



### Data Types

#### Item

| Field         | Type         | Description                                                |
| ------------- | ------------ | ---------------------------------------------------------- |
| id            | int          | the Item ID                                                |
| restaurant_id | int          | the restaurant ID                                          |
| name          | string       | the name of the Item                                       |
| Kind          | string       | the kind of Item this Item is                              |
| description   | string       | a description of the item                                  |
| modifiers     | list of Item | a list of other Items that can modify or add to this order |

#### Item kinds

| kind    | description                                   |
| ------- | --------------------------------------------- |
| entree  | a main course                                 |
| side    | a side dish                                   |
| topping | a topping to go on a main course or side dish |



### GET Endpoints

#### Items list

`/restaurants/{restaurantId}/item` 



**URL Parameters**

| param        | type | description                      |
| ------------ | ---- | -------------------------------- |
| restaurantId | int  | The ID of the desired restaurant |

**Returns**

A list of Items for the given restaurantId.

Currently the only content-encoding supported is JSON.

**Examples**

```[
   curl -X GET localhost:5000/restaurant/1/item
   
   {
      "id":1,
      "restaurant_id":1,
      "name":"cheese",
      "kind":"topping",
      "description":"cheddar cheese",
      "modifiers":"null"
   },
   {
      "id":5,
      "restaurant_id":1,
      "name":"french fries",
      "kind":"side",
      "description":"fried potatoes",
      "modifiers":[
         {
            "id":1,
            "restaurant_id":1,
            "name":"cheese",
            "kind":"topping",
            "description":"cheddar cheese",
            "modifiers":"null"
         }
      ]
   }
]
```



#### TODOs and future thoughts...

1. items endpoint as a list of nested item objects is a bit hard to work with, would prefer a true JSON dictionary hierarchy structure
2. data model could be improved upon, separating out items into their respective kinds as static types
3. similarly, DB backend storage structure could be improved, instead of a single items table that has an inherent recursive nature forced upon it
4. could move out types into separate "types" package for more modular/cross platform use
5. easy to add a global api route level LRU cache layer on top of storage related caches
6. add metrics tracking/instrumentation and profiling
7. add continuous integration for testing and possibly integrated deployment, like circleci 
8. add package vendoring with dep 
9. add additional HTTP content negotiation support, such as gzip compression, XML encoding format, etc
10. gather more measurements/estimates to assist in projections of expected future load/usage to better assess scalability & other engineering needs/requirements

