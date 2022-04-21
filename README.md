# tmp

# About

https://docs.google.com/document/d/1NbxiClee0xhqxPmfyHGkxe8JkZNQ2zg2EgEySyJK4jY/edit


# Api doc

```
http request:

POST /sort_flights
{
    "flights": [
        ["D", "E"],
        ["A", "B"],
        ["B", "C"],
        ["E", "F"],
        ["C", "D"]
    ]
}

http response:

["A","F"]

```


# Start http server

```
sh start.sh
```

# Curl sample

```
http request:

curl --location --request POST 'http://localhost:9000/sort_flights' \
--header 'Content-Type: application/json' \
--data-raw '{
    "flights": [
        ["D", "E"],
        ["A", "B"],
        ["B", "C"],
        ["E", "F"],
        ["C", "D"]
    ]
}'

http response:

["A","F"]

```

# You can see algorithm details in flights directory 