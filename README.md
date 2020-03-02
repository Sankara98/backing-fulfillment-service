# Fulfillment Service

This is the fulfillment service that is used with the backing-catalog service

For an interactive documentation and code samples visit the [Backing-fulfillment Services Docs](https://backingfulfillment.docs.apiary.io/#)

- `go build` - Creates the application binary, honoring the locally-vendored dependencies from the go-mod description file
- Run the application. By default, it will start on port \*_8000_, the port on which the default local copy of the [Catalog Service](https://github.com/Sankara98/backing-catalog) runs.

# Service API

This is a fake service, so it has no real data. It just returns manufactured payloads based on what you supply.

| Resource             | Method | Description                                                                                                |
| -------------------- | ------ | ---------------------------------------------------------------------------------------------------------- |
| products/{productId} | GET    | Fabricates a stock quantity and shipping time for an arbitrary productId. This productId can be any string |
