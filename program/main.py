from machine import Pin,ADC
import urequests as requests
import network
import time
import json

WIFI_SSID = 'Wokwi-GUEST'
WIFI_PASSWORD = ''

SERVER_URL = 'http://77.221.159.40:1337/process'
TIME_URL = 'http://worldtimeapi.org/api/timezone/etc/utc'

wlan = network.WLAN(network.STA_IF)

class AnalogSensor:
    def __init__(self, pin, sensor_type):
        self.sensor_id = pin
        self.pin = pin
        self.sensor_type = sensor_type

        self.adc = ADC(Pin(pin), atten = ADC.ATTN_11DB)
    
    def get_data(self) -> int:
        value = self.adc.read_u16()
        return value

pot4 = AnalogSensor(4, "potentiometr")
pot5 = AnalogSensor(5, "potentiometr")
pot6 = AnalogSensor(6, "potentiometr")

def do_connect():
    wlan.active(True)
    if not wlan.isconnected():
        print('connecting to network...')
        wlan.connect(WIFI_SSID, WIFI_PASSWORD)
        while not wlan.isconnected():
            pass
    print('network config:', wlan.ifconfig())

def send_data(sensors, current_time):
    data = []
    for sensor in sensors :
        data.append({
            "id": sensor.sensor_id,
            "sensor_type": sensor.sensor_type,
            "value": sensor.get_data()
        })

    jd = {
        "timestamp": current_time,
        "data": data
    }
    post_data = json.dumps(jd)

    headers = {'content-type': 'application/json'}

    res = requests.post(SERVER_URL, headers = headers, data = post_data)

def get_current_time() -> str:
    res = requests.get(TIME_URL)
    res_str = json.loads(res.text)
    return res_str["datetime"]

def loop():
    while True:
        current_time = get_current_time()
        print('data sended. Timestamp: {0}'.format(current_time))
        send_data([pot4, pot5, pot6], current_time)
        time.sleep_ms(10000)

do_connect()
loop()
