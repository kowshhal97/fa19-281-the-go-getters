## APIs for Review Microservice

**1. Ping the API Endpoint**

```Request```
```
  GET /reviews/ping
  Content-Type: application/json
```
```Response```
```
  {
    "The Go Getters Review API version 1.0 ALIVE!"
  }
```
---
**2. GET request to get all Reviews of a particular Menu Item**

```Request```
```
  GET /getReviews/{itemName}
  Content-Type: application/json
 ```
 ```Response```
```
INPUT: {itemName}="Cereals"
OUTPUT: [
    {
        "id": "5c0ac7571ddecdd2d1906677",
        "ItemName": "Cereals",
        "Reviews": [
            {
                "ReviewerName": "sbw"
                "Comment": "qqq",
                "Rating": 3
            },            
        ]
        "ReviewDate": "2019.11.15 20.44.54" 
    }
]
```

---
