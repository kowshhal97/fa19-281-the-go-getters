## The Go Getters

## Team Members 

1. [Anjana Menon](https://github.com/AnjanaMenonCherubala)
2. [Naman Agrawal](https://github.com/agrawalnaman)
3. [Pavani Somarouthu](https://github.com/Pavanisoma)
4. [Shivangi Jain](https://github.com/shivangi-jain)
5. [Kowshhal Uppu](https://github.com/kowshhal97)

---

## Description 

#### Frontend

1. The User frontend Screen can be used by to Login / Register into the application. The request is transferred from the Frontend to the User microservice
in backend through the Amazon API gateway.

2. The Menu frontend Screen displays all the items on the screen . It displays Itemname, ItemId, ItemDescription, ItemPrice to the user so the user and the data
is retreived from backend 'Menu' microservice.

3. The Order frontend screen will show the status of the order, total amount ,order Id and the related information of the order placed.
User has the option to cancel the order from frontend and also to select the order history and corresponding data is retrieved from backend.

4. The Payment frontend screen will have the option for the user to input card details. On clicking submit button , the payment is posted and the backend GO microservice is hit.

6. The Reviews frontend screen has facility to input the comment of a partucular Item and give appropriate rating. Also, user has an option to 
see all the reviews for a particular menu and it is time based. 

#### Amazon API Gateway 

The Amazon API Gateway is  used to route the resquests from the Frontend to the LoadBalancer which in turn route them to the individual GO APIs deployed in docker hosts.

#### Features of each Microservice

1. Login Microservice 
```
Create a new user (signup)
Login into account
```

2. Menu Microservice
```
Get all the Menu Items 
Get menu items based on Item Id.
```

3. Orders Microservice
```
Post an Order 
Get order details
Get all order details based on UserId
Delete Order
```

4. Payments Microservice
```
Post a payment
Delete Payment
```
5. Reviews Microservice
```
Post a review
Get reviews based on menuId
```

#### Mongo DB sharded cluster

1. The sharded cluster consists of two config servers in one replica set, 2 shards containing two nodes in each of them and one mongos(query router).
 - Each shard contains a subset of the sharded data. Shards are deployed as replica sets.
 - Mongos is the query router and it providing an interface between client applications and the sharded cluster.
 - Config Servers store the metadata and the configuration settings.
 
2. Shard Key  
The shard key determines the distribution of documents in the database. The shard key is immutable.
MongoDB aims to distribute data evenly among the shards and the shard key has the direct relationship with the effectiveness
of data distribution.

Please refer [this](https://github.com/nguyensjsu/fa19-281-the-go-getters/blob/master/Sharding_Steps.md) for the detailed steps.


