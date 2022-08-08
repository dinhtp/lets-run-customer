# Let's GO Customer API Gateway

## I. What to do
In this service, you will develop a backend server side functions that handle the **customer** API resources.
This service will contain 2 sub-services which are the Rest and gRPC service. Rest service will handle the http requests 
while gRPC will forward the request to the designated ecommerce platform. In order to get the correct gRPC address for 
the ecommerce service, `platform_id` is provided to identify which platform to fetch from the **platform** service 
database.

## II. API Output
### Customer API Output
#### Get a customer by ID.
    - URL: [GET] {customer_url}/run/{platform_id}/customer/{id}
    - Response: 
        {
            "id": "string",
            "first_name": "string",
            "middle_name": "string",
            "last_name": "string",
            "email": "string",
            "phone": "string",
            "note": "string",
            "status": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Create a customer.
    - URL: [POST] {customer_url}/run/{platform_id}/customer
    - Payload:
        {
            "first_name": "string",
            "middle_name": "string",
            "last_name": "string",
            "email": "string",
            "phone": "string",
            "note": "string",
            "status": "string", // allowed values: "active", "inactive"
        }
    - Response:
        {
            "id": "string",
            "first_name": "string",
            "middle_name": "string",
            "last_name": "string",
            "email": "string",
            "phone": "string",
            "note": "string",
            "status": "string", // allowed values: "active", "inactive"
            "created_at": "string",
            "updated_at": "string"
        }
#### Update a customer by ID.
    - URL: [PUT] {customer_url}/run/{platform_id}/customer/{id}
    - Payload:
        {
            "id": "string",
            "first_name": "string",
            "middle_name": "string",
            "last_name": "string",
            "email": "string",
            "phone": "string",
            "note": "string",
            "status": "string", // allowed values: "active", "inactive"
        }
    - Response:
        {
            "id": "string",
            "first_name": "string",
            "middle_name": "string",
            "last_name": "string",
            "email": "string",
            "phone": "string",
            "note": "string",
            "status": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Delete a customer by ID.
    - URL: [DELETE] {customer_url}/run/{platform_id}/customer/{id}
    - Status: 200
#### List customers by page and limit
    - URL: [GET] {customer_url}/run/{platform_id}/customers
    - Query: ?page=1&limit=10
    - Response:
        {
            "items": [
                {
                    "id": "string",
                    "first_name": "string",
                    "middle_name": "string",
                    "last_name": "string",
                    "email": "string",
                    "phone": "string",
                    "note": "string",
                    "status": "string",
                    "created_at": "string",
                    "updated_at": "string"
                },
                ...
            ]
            "max_page": 0,
            "total_count": 0,
            "page": 1,
            "limit": 10,
        }