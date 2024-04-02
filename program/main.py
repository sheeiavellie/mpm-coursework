from machine import Pin,ADC
import urequests as requests
import network
import time

adc4 = ADC(Pin(4), atten = ADC.ATTN_11DB)
adc5 = ADC(Pin(5), atten = ADC.ATTN_11DB)
adc6 = ADC(Pin(6), atten = ADC.ATTN_11DB)

def do_connect():
    wlan = network.WLAN(network.STA_IF)
    wlan.active(True)
    if not wlan.isconnected():
        print('connecting to network...')
        wlan.connect('Wokwi-GUEST', '')
        while not wlan.isconnected():
            pass
    print('network config:', wlan.ifconfig())

def read_data():
    read = adc4.read()
    read_u16 = adc4.read_u16()
    read_uv = adc4.read_uv()
    print("read={0},read_u16={1},read_uv={2}".format(read,read_u16,read_uv))
    time.sleep_ms(100)

def loop():
    while True:
        read_data()


loop()
