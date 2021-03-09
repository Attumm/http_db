# LambdaDB
In memory database that uses filters to get the data you need.
Lambda DB has a tiny codebase which does a lot
Lambda is not ment as a persistance storage or a replacement for a traditional
Database but as fast analytics engine cache representation engine.

powers: https://dego.vng.nl

## Properties:

- Insanely fast API. 1ms respsonses
- Fast to setup.
- Easy to deploy.
- Easy to customize.
- Easy export data

- Implement custom authorized filters.

## Indexes

- S2 geoindex for fast point lookup
- Bitarrays
- Mapping

- Your own special needs indexes!

## Flow:

Generate a model and load your data.
The API is generated from your model.
Deploy.

Condition: Your dataset must fit in memory.

Can be used for your needs by changing the `models.go` file to your needs.
Creating and registering of the functionality that is needed.


### Steps
You can start the database with only a csv.
Go over steps below, And see the result in your browser.

1. place csv file, in dir extras.
2. `python3 create_model_v2.py`  answer the questions..
3. go fmt model.go
4. mv model.go ../
5. go build
6. ./lambda --help
7. ./lambda  --csv assets/items.csv or `python3 ingestion.py -b 1000`
9. curl 127.0.0.1:8000/help/
10. browser 127.0.0.1:8000/


11. instructions curl 127.0.0.1:8000/help/ | python -m json.tool


### TODO

- load data directly from a database (periodic)
- use a remote source for CSV
- use some compression faster to load than gzip
- generate swagger API
