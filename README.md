# Prove of concept of MQTT protocol using KAFKA as a broker

## 1. Setup
   ### KAFKA
   ### 1. For Docker
   #### Step 1: Download Docker compose file
          $ curl -sSL https://raw.githubusercontent.com/bitnami/bitnami-docker-kafka/master/docker-compose.yml > docker-compose.yml
   #### Step 2: Edit the file for external host
    environment:
      - KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181
      - ALLOW_PLAINTEXT_LISTENER=yes
      - KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CLIENT:PLAINTEXT,EXTERNAL:PLAINTEXT
      - KAFKA_CFG_LISTENERS=CLIENT://:9092,EXTERNAL://:9093
      - KAFKA_CFG_ADVERTISED_LISTENERS=CLIENT://kafka:9092,EXTERNAL://localhost:9093
      - KAFKA_CFG_INTER_BROKER_LISTENER_NAME=CLIENT
    ports:
      - '9092:9092'
      - '9093:9093'
   #### Step 3: Run the Docker compose
      $ docker-compose up -d
   #### Step 4: Inside Kafka container CLI navigate to following path:  
      cd opt/bitnami/kafka/bin
   #### Step 5: Create Topic according to the requirement
       kafka-topics.sh --create --topic <topic name> --replication-factor <value> --partitions <value> --bootstrap-server localhost:9092
   #### Note: The kafka will run in localhost:9093
   
   ### 2. Local Machine
   #### Step 1: Download kafka from: https://www.apache.org/dyn/closer.cgi?path=/kafka/3.1.1/kafka_2.13-3.1.1.tgz and unzip
   #### Step 2: Navigate to the bin of the extracted folder and open 3 terminals
   #### Step 3: Run these commands in seperate terminals 1.Start Zookeeper 2.Start Kafka_server
  
         $ bin/zookeeper-server-start.sh config/zookeeper.properties
         $ bin/kafka-server-start.sh config/server.properties
  #### Step 4: Create topic as per requirement
         kafka-topics.sh --create --topic <topic name> --replication-factor <value> --partitions <value> --bootstrap-server localhost:9092
  ### MQTT Broker
      
## 2. Code
   1. [Kafka producer](cmd/kafka_producer) <br>
     This Producer can send message to the kafka broker using pre-defined balancers and balancers can also be implemented using custom balancer using different logics.<br>
   2. [Kafka consumer](cmd/kafka_consumer) <br>
   The consumer can recieve message from the kafka broker and the implemented concept of consumer group is used for no duplicate processing of same data.<br>
   3. [Kafka balancer](inernal/balancer)   <br>
   This balancer method is present in the internal folder from where it is used in kafka producer.This method is created for making own load balancer using different logics.<br>
   4. [MQTT publisher](cmd/publish)        <br>
   This MQTT publisher simply sends a message to the MQTT broker to a particular topic.<br>
   5. [MQTT subscriber](cmd/subscribed)<br>
   This MQTT subscribers simply read the messages from a particular topic.
   
