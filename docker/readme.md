### Start Kafka
```
docker compose up -d
```

### Get Container id
```
docker ps
```

### Create Topic
```
docker exec -it <containr_id> /opt/bitnami/kafka/bin/kafka-topics.sh --bootstrap-server=localhost:9092 --create --replication-factor=1 --partitions=1 --topic=topic1
```
