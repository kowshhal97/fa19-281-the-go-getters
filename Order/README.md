## APIs for Order Microservice

**1. Ping the API Endpoint**

```Request```
```
  GET /order/ping
  Content-Type: application/json
```
```Response```
```
  {
    "Order API is alive"
  }
```
---
**2. GET request to get a particular order**

```Request```
```
  GET /order/orderid
  Content-Type: application/json
 ```
 ```Response```
 
| Parameter      | Type          | Description  |
| ------------- |:-------------:| -----:|
| OrderId      | String | OrderId of the order placed |
| UserId      | String      | Id of the user placing the order |
| Items | Struct      |    List of all items ordered |
| TotalAmount | float32      |    Total amount of placed order |

---

**3. POST request to place an order**

```Request```
```
   POST /order
   Content-Type: application/json
```
| Parameter      | Type          | Description  |
| ------------- |:-------------:| -----:|
| UserId      | String      | Id of the user placing the order |
| ItemName | String      |   Name of the item  |
| ItemPrice      | float32      | Price of the order |
| ItemQuantity | Integer      |   Quantity of the item  |

```Response```

|Parameter	|Type	|Description  |
|----|----|----|
|orderId |String | Order ID of the order placed |
|userId | String | Id of user who has placed the order |
|orderStatus | String  | Status: Order Processing |
|items |Struct | ItemName,Itemprice, ItemQuantity|
|totalAmount	| float32 | Total amount of the order placed by user|

---

**4. Update the order status after payment is done**

```Request```
```
   PUT /order/orderid
   Content-Type: application/json
```
```Response```

|Parameter	|Type	|Description  |
|----|----|----|
|orderId |String | Order ID of the order placed |
|userId | String | Id of user who has placed the order |
|orderStatus | String  | Status: Placed |
|items |Struct | ItemName,Itemprice, ItemQuantity|
|totalAmount	| float32 | Total amount of the order placed by user|

---

**5.Delete the order**

```Request```
```
   DELETE /order/orderid
   Content-Type: application/json
```

Status code: 200


|Parameter	|Type |	Description|
|-----|-----|------|
|messsage	|String| Order deleted |

Status code: 400

|Parameter	|Type |	Description|
|-----|-----|------|
|messsage	|String| Sorry, order not found |

---
