# 3.MicroserviceEngineering
AM-RentalManagement provides the functionality for the capability Management of the Rentals via API endpoints dedicated to individual use cases.

## Design AM-RentalManagement

[API Diagram](https://gitlab.kit.edu/kit/cm/teaching/carrentalapp/carrentalappv1/-/blob/main/pages/api_diagram_am-rental_management_v1.1.md)

[API Specification](./src/api/specification/api_specification_am_rental_management.proto)



## Run AM-RentalManagement

1. Create a copy of the `src/.env.example` file and rename it to `src/.env`. Afterward, update the variable settings within this new file.
  
2. In your terminal, change to the `src/` directory and execute the following command: 
    ```
    go run .
    ```


## Run Tests

1. Execute `go test -v ./...`


# Rental and Customer Service API

## 📌 Overview

This document provides detailed information on the gRPC APIs for the Rental and Customer services within the `rentalmanagement` package.

---

## Installing grpcurl

To install `grpcurl` using Go, make sure you have Go installed on your machine. Then, run the following command:

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7
```

## 🚗 Rental Service

### `RentCar`

#### Request Format

```protobuf
RentCarRequest {
    startDate: google.protobuf.Timestamp;
    endDate: google.protobuf.Timestamp;
    vin: string;
    customerId: string;
}
```

#### 📟 Command Line Request

```bash
grpcurl -d="{\"startDate\":\"2022-02-19T00:00:00Z\",\"endDate\":\"2022-02-20T00:00:00Z\",\"vin\":\"JH4DB1561NS000569\",\"customerId\":\"CUST12345\"}" -plaintext localhost:50051 rentalmanagement.RentalService/RentCar
```

#### 📤 Response Body

```json
{
    "rental": {
        "rentalId": "887fdad9-ec1d-4f9a-8592-bb3308aba8b9",
        "customerId": "CUST12345",
        "vin": "JH4DB1561NS000569",
        "startDate": {"seconds": 1645237120},
        "endDate": {"seconds": 1645937120}
    },
    "error": null
}
```
##### 📤 Error Response Body

```json
{
    "rental": null,
    "error": {
        "message": "Internal",
        "details": "car with VIN JH4DB1561NS000569 is not available for the specified time range"
    }
}
```


---

### `CancelRental`

#### Request Format

```protobuf
CancelRentalRequest {
    rentalId: string;
}
```

#### 📟 Command Line Request

```bash
grpcurl -d {\"rentalId\":\"887fdad9-ec1d-4f9a-8592-bb3308aba8b9\"} -plaintext localhost:50051 rentalmanagement.RentalService/CancelRental
```

#### 📤 Response Body

```json
{
    "error": null
}
```


##### 📤 Error Response Body

```json
{
    "error": {
        "message": "Internal",
        "details": "Failed to check existence of rental with ID 887fdad9-ec1d-4f9a-8592-bb3308aba8b9: Database failed to find rental with ID 887fdad9-ec1d-4f9a-8592-bb3308aba8b9"
    }
}
```

---

### `ListAvailableCars`

#### Request Format

```protobuf
ListAvailableCarsRequest {
    startDate: google.protobuf.Timestamp;
    endDate: google.protobuf.Timestamp;
}
```

#### 📟 Command Line Request

```bash
grpcurl -d="{\"startDate\":\"2022-02-19T00:00:00Z\",\"endDate\":\"2022-02-20T00:00:00Z\"}" -plaintext localhost:50051 rentalmanagement.RentalService/ListAvailableCars

```

#### 📤 Response Body

```json
{
    "cars": [
        {"vin": "JH4DB1561NS000569", "brand": "Tesla", "model": "Model S"},
        {"vin": "JH4DB1561NS000569", "brand": "BMW", "model": "M3"}
    ],
    "error": null
}
```

##### 📤 Error Response Body

```json
{
    "cars": [],
    "error": {
        "message": "Internal",
        "details": "StartDate must be before EndDate"
    }
}
```

---
### `ListCustomerRentals`

#### Request Format

```protobuf
ListCustomerRentalsRequest {
    customerId: string;
}
```

#### 📟 Command Line Request

```bash
grpcurl -d {\"customerId\":\"CUST12345\"} -plaintext localhost:50051 rentalmanagement.RentalService/ListCustomerRentals
```

#### 📤 Response Body

```json
{
    "rentals": [
        {
            "rentalId": "5cc65441-39ec-48bc-8c63-bb2f6ced9ba4",
            "customerId": "CUST12345",
            "vin": "JH4DB1561NS000566",
            "startDate": {"seconds": 1645237120},
            "endDate": {"seconds": 1645937120}
        },
        {
            "rentalId": "3ef41de2-2a19-46df-aceb-fb181d8cd552",
            "customerId": "CUST12345",
            "vin": "JH4DB1561NS000568",
            "startDate": {"seconds": 1646237120},
            "endDate": {"seconds": 1646937120}
        }
    ],
    "error": null
}
```
#### 📤 Error Response Body

```json
{
    "rentals": [],
    "error": {
        "message": "InvalidArgument",
        "details": "CustomerId is not valid"
    }
}
```

---

## 👥 Customer Service

### `RegisterCustomer`

#### Request Format

```protobuf
RegisterCustomerRequest {
    name: string;
}
```

#### 📟 Command Line Request

```bash
grpcurl -d {\"name\":\"John Doe\"} -plaintext localhost:50051 rentalmanagement.CustomerService/RegisterCustomer
```

#### 📤 Response Body

```json
{
    "customer": {
        "customerId": "CUST12345",
        "name": "John Doe"
    }
}
```

---

### `DeregisterCustomer`

#### Request Format

```protobuf
DeregisterCustomerRequest {
    customerId: string;
}
```

#### 📟 Command Line Request

```bash
grpcurl -d {\"customerId\":\"CUST12345\"} -plaintext localhost:50051 rentalmanagement.CustomerService/DeregisterCustomer
```

#### 📤 Response Body

```json
{
    "error": null
}
```
