# Myclub

Has 2 parts:

1. Server which will add player data and will produce into a topic in kafka.
2. Worker which will pull the message from kafka and add it in a map. 


## Steps to start KAFKA in local

Go to kafka location and start zoo keeper

```
bin/zookeeper-server-start.sh config/zookeeper.properties
```

Inside `config/server.properties` add broker address(`listeners=PLAINTEXT://:9093`) and start broker instance

```
bin/kafka-server-start.sh config/server.properties
```

Create topic `real-madrid`
```
bin/kafka-topics.sh --create --topic real-madrid --bootstrap-server localhost:9093 --partitions 1 --replication-factor 1
```


## Steps to run locally

```
make compile
```

`./api` will start the server
`./worker add` will run the worker to consume messages from producer
Send a POST request to POST `/v1/players`
Worker will consume the message once message is produced. 



[Detailed explaination about installing and running Kafka in local can be found here](https://www.sohamkamani.com/install-and-run-kafka-locally/)

