import paho.mqtt.client as mqtt #import the client1
import time
broker_address="localhost" 
#broker_address="iot.eclipse.org" #use external broker
client = mqtt.Client("P1") #create new instance
client.connect(broker_address, 1883) #connect to broker
while True:
    client.publish("house/main-light","1")#publish    
    time.sleep(1)