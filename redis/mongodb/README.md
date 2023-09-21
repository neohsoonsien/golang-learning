# Setup the MongoDB Database Using Docker-Compose
1.  Initialize the mongodb database using docker-compose.
    ```bash
        docker-compose up --build -d
    ```
2.  Login to the MongoDB server with
    ```bash
        mongosh -u root --host localhost --port 3001 -p --authenticationDatabase admin
    ```
    The `password` would be ***example***.
3.  Create an e-commerce database
    ```bash
        test> use e-commerce;
    ```
    Output:
    ```bash
        switched to db e-commerce
    ```
4.  Insert five documents into a new `products` collection
    ```bash
        e-commerce> db.products.insertMany([
            {"product_id" : 1,
             "product_name" : "INSTANT WATER HEATER",
             "retail_price" : 45.55
            },
            {"product_id" : 2,
             "product_name" : "DOUBLE SOCKET WITH PATTRESS",
             "retail_price" : 6.65
            },
            {"product_id" : 3,
             "product_name" : "80MM USB PRINTER",
             "retail_price" : 125.95
            },
            {"product_id" : 4,
             "product_name" : "FITNESS SMARTWATCH",
             "retail_price" : 39.85
            },
            {"product_id" : 5,
             "product_name" : "3.1A FAST CHARGER",
             "retail_price" : 23.90
            }
        ]);
    ```
    Output:
    ```bash
        {
            acknowledged: true,
            insertedIds: {
                '0': ObjectId("6451ec655b20925eda45a93a"),
                '1': ObjectId("6451ec655b20925eda45a93b"),
                '2': ObjectId("6451ec655b20925eda45a93c"),
                '3': ObjectId("6451ec655b20925eda45a93d"),
                '4': ObjectId("6451ec655b20925eda45a93e")
            }
        }
    ```
5.  Query the `products` collection to verify the documents:
    ```bash
        e-commerce> db.products.find();
    ```
    Output
    ```bash
        [
            {
                _id: ObjectId("6451ec655b20925eda45a93a"),
                product_id: 1,
                product_name: 'INSTANT WATER HEATER',
                retail_price: 45.55
            },
            {
                _id: ObjectId("6451ec655b20925eda45a93b"),
                product_id: 2,
                product_name: 'DOUBLE SOCKET WITH PATTRESS',
                retail_price: 6.65
            },
            {
                _id: ObjectId("6451ec655b20925eda45a93c"),
                product_id: 3,
                product_name: '80MM USB PRINTER',
                retail_price: 125.95
            },
            {
                _id: ObjectId("6451ec655b20925eda45a93d"),
                product_id: 4,
                product_name: 'FITNESS SMARTWATCH',
                retail_price: 39.85
            },
            {
                _id: ObjectId("6451ec655b20925eda45a93e"),
                product_id: 5,
                product_name: '3.1A FAST CHARGER',
                retail_price: 23.9
            }
        ]
    ```

# Test the Application from localhost
1.  Launch the application with
    ```bash
        go run ./
    ```
2.  One can obtain the response from the application via
    ```bash
        curl localhost:8080/products
    ```
    If one runs the same command within the `redis` expiration time, the `_source` would be _"Redis Cache"_ instead of _"MongoDB database"_.

# Browse the Redis Cache via Redis-Insight
1.  One can lauch the _Redis-Insight_ application at
    ```bash
        localhost:8001
    ```
2.  One should provide the redis `host:port` and other connection details as
    -   Host:     redis
    -   Port:     6379
    -   Name:     e-commerce
    -   Password: PASSWORD
    when prompted.