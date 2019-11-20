## Mongo DB Sharded cluster

#### The Cluster consists of one replica set containing two config server, two shards containing 2 nodes in each shard and 1 mongos. 

#### Steps 

#### 1. Create 7(4 instances for sharding, 2 config servers and 1 query router) Ubuntu instances and have MongoDB installed in them. 
Refer - https://github.com/paulnguyen/cmpe281/blob/master/labs/lab5/aws-mongodb-replica-set.md

#### 2. Create a config server replica Set 

Edit /etc/mongod.conf file in each of the two config servers. 
```
net:
  port: 27019
  bindIp: 0.0.0.0
replication:
  replSetName: <replication_set_name>
sharding:
  clusterRole: configsvr
  ```
 
 #### 3. Start mongod
 ```
 sudo mongod --config /etc/mongod.conf
 ```
 
 #### 4. Connect to one config server from jumpbox and fire rs.initiate on it
 ```
  mongo --host <host-ip> --port 27019
  
  mongo> > rs.initiate(
...   {
...     _id: <replication_set_name>,
...     members: [
...       { _id : 0, host : "config_server_1> },
...       { _id : 1, host : "config_server_2" }
...     ]
...   }
... )

Response - 

{
        "ok" : 1,
        "operationTime" : Timestamp(1574228545, 1),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1574228545, 1),
                "signature" : {
                        "hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
                        "keyId" : NumberLong(0)
                }
        }
}
```
---

 #### 1. Create Shard server 
 
 Repeat the below steps for each of the two shard servers 
 
 #### 1. Edit /etc/mongod.conf file in each member of the shards
 ```
 net:
  port: 27018
  bindIp: 0.0.0.0
 replication:
  replSetName: <replication_set_name_for_shard_server>
sharding:
  clusterRole: shardsvr
 ```
 
 #### 2. Start mongod
 ```
 sudo mongod --config /etc/mongod.conf
 ```
 
 #### 3. Connect to one shard member from jumpbox and fire rs.initiate on it
 ```
  mongo --host <host-ip> --port 27018
  
  mongo> > rs.initiate(
...   {
...     _id: <replication_set_name>,
...     members: [
...       { _id : 0, host : "shard_member_1> },
...       { _id : 1, host : "shard_member_2" }
...     ]
...   }
... )

Response - 

{
        "ok" : 1,
        "operationTime" : Timestamp(1574228545, 1),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1574228545, 1),
                "signature" : {
                        "hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
                        "keyId" : NumberLong(0)
                }
        }
}
```

---

#### Mongos (Query router) Configuration

#### 1. Edit /etc/mongod.conf

```
net:
  port: 27017
sharding:
  configDB: config_server_replication_set/config_server_1:27019,config_server_2:27019
```
Comment out the storage part as mongos doesnt need that.

#### 2. Start mongos 
```
sudo mongos --config /etc/mongod.conf --fork

Response -

2019-11-20T06:12:25.104+0000 W SHARDING [main] Running a sharded cluster with fe                                                                                        wer than 3 config servers should only be done for testing purposes and is not re                                                                                        commended for production.
about to fork child process, waiting until server is ready for connections.
forked process: 10408

child process started successfully, parent exiting
```

#### 3. Connect to mongos from the jumpbox 

```
 mongo -port 27017
 ```
 
 #### 4. Add each shard to mongos
 ```
 mongos> sh.addShard("<replication_set_of_shard1>/<shard_member_1_ip:27018,shard_member_2_ip:27018")
 
 Expected Response - 
 
{
        "shardAdded" : "repl2",
        "ok" : 1,
        "operationTime" : Timestamp(1574230763, 5),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1574230763, 5),
                "signature" : {
                        "hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
                        "keyId" : NumberLong(0)
                }
        }
}
```
Add the second shard similarly
```
mongos> sh.addShard("<replication_set_of_shard2>/<shard_member_1_ip:27018,shard_member_2_ip:27018")
{
        "shardAdded" : "repl3",
        "ok" : 1,
        "operationTime" : Timestamp(1574230795, 6),
        "$clusterTime" : {
                "clusterTime" : Timestamp(1574230795, 6),
                "signature" : {
                        "hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
                        "keyId" : NumberLong(0)
                }
        }
}
```
#### 5. Steps to enable sharding 

```
mongos> use testdb
mongos> db.runCommand({enablesharding: "testdb"})
```
This enables sharding on database 'testdb'

#### 6. Shard a collection by specifying the shard key. 
Note - Once the shard key is specified, it is permanent. Choosing a shard key wisely helps as it effects the performance.
```
mongos> db.runCommand( {shardcollection: "testdb.orders", key: { orderId :1 }});
```

I have done sharding on the basis of OrderId here. 

---


References - 
https://docs.mongodb.com/manual/tutorial/deploy-shard-cluster/

https://docs.mongodb.com/manual/reference/replica-configuration/

https://docs.mongodb.com/manual/core/hashed-sharding/#hashed-sharding-shard-key


