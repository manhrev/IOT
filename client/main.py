import serial.tools.list_ports
import paho.mqtt.client as mqtt
import time
import sys
import json


FEED_IDS = ["led-1", "led-2", "pump-1", "pump-2"]
isMicrobitConnected = False

def connected(client, userdata, flags, rc):
    print("Ket noi thanh cong...")
    for feed in FEED_IDS:
        client.subscribe(feed)

def subscribe(client , userdata , mid , granted_qos):
    print("Subcribe thanh cong...")

def disconnected(client):
    print("Ngat ket noi...")
    sys.exit (1)

def message(client, userdata, msg):
    data = str(msg.payload)
    data = data.split(":")[1].replace("}", "").replace(" ", "").replace("'", "")
    print("Nhan du lieu: " + data)
    
    if isMicrobitConnected:
        ser.write((data + "#").encode())


def send_microbit(value):
    if isMicrobitConnected:
        ser.write((str(value) + "#").encode())


broker_address= "localhost"
client = mqtt.Client("P1") # create new instance
client.on_connect = connected
client.on_disconnect = disconnected
client.on_message = message
client.on_subscribe = subscribe
client.connect(broker_address, 1883) # connect to broker
client.loop_forever()



def getPort():
    ports = serial.tools.list_ports.comports()
    N = len(ports)
    commPort = "None"
    for i in range(0, N):
        port = ports[i]
        strPort = str(port)
        if "USB Serial Device" in strPort:
            splitPort = strPort.split(" ")
            commPort = (splitPort[0])
    return commPort


if getPort() != "None":
    ser = serial.Serial( port=getPort(), baudrate=115200)
    isMicrobitConnected = True


def processData(data):
    data = data.replace("!", "")
    data = data.replace("#", "")
    splitData = data.split(":")
    print(splitData)

    myData = {
        "value": splitData[2]
    }
    pubData = json.dumps(myData) # format before publish

    try:
        if splitData[1] == "GRHUMI1":
            client.publish("grhumi-1", pubData)
        elif splitData[1] == "LIGHT1":
            client.publish("light-1", pubData)
        elif splitData[1] == "GRHUMI2":
            client.publish("grhumi-2", pubData)
        elif splitData[1] == "LIGHT2":
            client.publish("light-2", pubData)

    except:
        pass

mess = ""
def readSerial():
    bytesToRead = ser.inWaiting()
    if (bytesToRead > 0):
        global mess
        mess = mess + ser.read(bytesToRead).decode("UTF-8")
        while ("#" in mess) and ("!" in mess):
            start = mess.find("!")
            end = mess.find("#")
            processData(mess[start:end + 1])
            if (end == len(mess)):
                mess = ""
            else:
                mess = mess[end+1:]

while True:
    if isMicrobitConnected:
        readSerial()

    time.sleep(1)