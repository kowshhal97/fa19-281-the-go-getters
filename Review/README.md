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
```
  INPUT: {itemName}=MexicanPizza
 ```
 ```Response```
```
OUTPUT: 
[
  {
    "id": "5dd38c07e1af6e9229916850",
    "ItemName": "MexicanPizza",
    "Reviews": [
      {
        "ReviewerName": "Ayushman",
        "Comment": "good",
        "Rating": 4
      }
    ],
    "ReviewDate": "2019-11-18 22:30:31"
  },
  {
    "id": "5dd38c15e1af6e922991685a",
    "ItemName": "MexicanPizza",
    "Reviews": [
      {
        "ReviewerName": "Naman",
        "Comment": "good",
        "Rating": 4
      }
    ],
    "ReviewDate": "2019-11-18 22:30:45"
  }
]
```

---

**3. POST request to post a new review**

```Request```
```
   POST /postReview
   Content-Type: application/json
```

```
INPUT:
{
	"ItemName" : "MexicanPizza",
	"Reviews" : [
					{
						"ReviewerName" : "Naman",
						"Comment" : "good",
						"Rating" : 4
					}
				]
}
```
```Response```
```

OUTPUT:
{
  "Response": "Review added"
}
```
---
**4. Update previous review**

```Request```
```
   PUT /updateReview
   Content-Type: application/json
```
```
INPUT:
{
	"ItemName" : "MexicanPizza",
	"Reviews" : [
					{
						"ReviewerName" : "Naman",
						"Comment" : "Best Pizza Ever",
						"Rating" : 5
					}
				]
}
```
```Response```
```
OUTPUT:
{
  "Response": "Review updated"
}
```
---

**5. Delete Last review**

```Request```
```
   DELETE /deleteReview
   Content-Type: application/json
```

```Response```
```
OUTPUT:
{
  "Response": "Review deleted"
}
```
---
