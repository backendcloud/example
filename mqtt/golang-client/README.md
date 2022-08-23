```bash
[root@centos9 golang-client]# docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:latest
Emulate Docker CLI using podman. Create /etc/containers/nodocker to quiet msg.
✔ docker.io/emqx/emqx:latest
Trying to pull docker.io/emqx/emqx:latest...
Getting image source signatures
Copying blob 4f4fb700ef54 done  
Copying blob 1b7ca6aea1dd done  
Copying blob 547d45a2540b done  
Copying blob 754c16f4511a done  
Copying blob 8e098de503c8 done  
Copying blob eeb6bfc6f84e done  
Copying blob 3741a835589f done  
Copying blob 0803a8486f11 done  
Copying config eb0d480e05 done  
Writing manifest to image destination
Storing signatures
49f999af11bf5d01aae98c67f0b8baa502b45049ca1054fc6d34ccd99fd8651b
[root@centos9 golang-client]# go run main.go 
[client]   Connect()
[store]    memorystore initialized
[client]   about to write new connect msg
[client]   socket connected to broker
[client]   Using MQTT 3.1.1 protocol
[net]      connect started
[net]      received connack
[client]   startCommsWorkers called
[client]   client is connected/reconnected
[net]      incoming started
[net]      startIncomingComms started
[net]      outgoing started
[net]      startComms started
[client]   startCommsWorkers done
[client]   exit startClient
[client]   enter Subscribe
[client]   SUBSCRIBE: dup: false qos: 1 retain: false rLength: 0 MessageID: 1 topics: [go-mqtt/sample]
[client]   sending subscribe message, topic: go-mqtt/sample
[net]      outgoing waiting for an outbound message
[net]      logic waiting for msg on ibound
[net]      startIncomingComms: inboundFromStore complete
[net]      logic waiting for msg on ibound
[pinger]   keepalive starting
[net]      obound priority msg to write, type *packets.SubscribePacket
[net]      outgoing waiting for an outbound message
[client]   exit Subscribe
[net]      startIncoming Received Message
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received suback, id: 1
[net]      startIncomingComms: granted qoss [0]
[net]      logic waiting for msg on ibound
[client]   enter Publish
[client]   sending publish message, topic: go-mqtt/sample
[net]      obound msg to write 0
[net]      obound wrote msg, id: 0
[net]      outgoing waiting for an outbound message
[client]   enter Publish
[client]   sending publish message, topic: go-mqtt/sample
[net]      obound msg to write 0
[net]      obound wrote msg, id: 0
[net]      outgoing waiting for an outbound message
[client]   enter Publish
[client]   sending publish message, topic: go-mqtt/sample
[net]      obound msg to write 0
[net]      obound wrote msg, id: 0
[net]      outgoing waiting for an outbound message
[client]   enter Publish
[client]   sending publish message, topic: go-mqtt/sample
[net]      obound msg to write 0
[net]      obound wrote msg, id: 0
[net]      outgoing waiting for an outbound message
[client]   enter Publish
[client]   sending publish message, topic: go-mqtt/sample
[net]      obound msg to write 0
[net]      obound wrote msg, id: 0
[net]      outgoing waiting for an outbound message
[net]      startIncoming Received Message
[net]      startIncoming Received Message
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received publish, msgId: 0
[net]      logic waiting for msg on ibound
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received publish, msgId: 0
[net]      startIncoming Received Message
[net]      logic waiting for msg on ibound
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received publish, msgId: 0
[net]      logic waiting for msg on ibound
[net]      startIncoming Received Message
TOPIC: go-mqtt/sample
MSG: this is msg #0!
TOPIC: go-mqtt/sample
MSG: this is msg #1!
TOPIC: go-mqtt/sample
MSG: this is msg #2!
[net]      startIncoming Received Message
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received publish, msgId: 0
TOPIC: go-mqtt/sample
MSG: this is msg #3!
[net]      logic waiting for msg on ibound
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received publish, msgId: 0
[net]      logic waiting for msg on ibound
TOPIC: go-mqtt/sample
MSG: this is msg #4!
[pinger]   ping check 0.999037599
[pinger]   ping check 1.9988296490000002
[pinger]   ping check 2.999283086
[pinger]   keepalive sending ping
[net]      startIncoming Received Message
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received pingresp
[net]      logic waiting for msg on ibound
[pinger]   ping check 0.999980534
[pinger]   ping check 1.99975846
[pinger]   ping check 3.000078984
[pinger]   keepalive sending ping
[client]   enter Unsubscribe
[client]   sending unsubscribe message, topics: [go-mqtt/sample]
[client]   exit Unsubscribe
[net]      startIncoming Received Message
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received pingresp
[net]      logic waiting for msg on ibound
[net]      obound priority msg to write, type *packets.UnsubscribePacket
[net]      outgoing waiting for an outbound message
[net]      startIncoming Received Message
[net]      startIncomingComms: got msg on ibound
[net]      startIncomingComms: received unsuback, id: 2
[net]      logic waiting for msg on ibound
[client]   disconnecting
[client]   calling WaitTimeout
[net]      obound priority msg to write, type *packets.DisconnectPacket
[net]      outbound wrote disconnect, closing connection
[client]   WaitTimeout done
[client]   stopCommsWorkers called
[net]      incoming complete
[client]   stopCommsWorkers waiting for workers
[net]      startIncomingComms: ibound complete
[net]      startIncomingComms goroutine complete
[router]   matchAndDispatch exiting
[client]   startCommsWorkers output redirector finished
[net]      outgoing waiting for an outbound message
[net]      outgoing waiting for an outbound message
[net]      outgoing waiting for an outbound message
[net]      outgoing waiting for an outbound message
[net]      outgoing comms stopping
[net]      startComms closing outError
[client]   incoming comms goroutine done
[pinger]   keepalive stopped
[client]   stopCommsWorkers waiting for comms
[client]   stopCommsWorkers done
[client]   forcefully disconnecting
[msgids]   cleaned up
[client]   disconnected
[store]    memorystore closed
[root@centos9 golang-client]# 
```
