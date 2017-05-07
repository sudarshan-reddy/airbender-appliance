
## Airbender appliance

Appliance code that would run of a GrovePi+ integrated Raspberry Pi

### Prequisite:

*Hardware*
RaspberryPi
GrovePi+
Grove AirQuality Sensor v1.3
Grove SHT31 Sensor

*Software*
Raspbian
git clone `https://github.com/DexterInd/GrovePi`
Update the Firmware using 
```
sudo ./GrovePi/Firmware/firmware_update.sh
```

### Installation:
#### From the source
In your raspberry pi
 ```
 go get github.com/sudarshan-reddy/airbender-appliance
 glide up
 go build
 ```
#### Raspberry pi executable Binary
```
wget https://github.com/sudarshan-reddy/airbender-appliance/releases/download/0.1b/airbender-appliance
chmod+x airbender-appliance
```

#### Running the binary:
```
    source variables.env
    ./airbender-appliance
```

### Variables

| Variable |Sample Value| Description |
|----------|-------------|-------------|
|AA_MQTT_URL| | mqtt url
|AA_MQTT_TOPIC|event.test.airquality | mqtt publish topic
|AA_MQTT_CLIENT|airbender | mqtt clientid
|AA_APPLIANCE_NAME| home_bedroom_1| name for the appliance
|AA_LOG_FILE_NAME|logs.txt| logfile to write logs to
|AA_TIME_ZONE|Asia/Kolkata| timezone to use
|MONITOR_INTERVAL|10s| Polling interval

### MQTT message format

```json
{
  "name": "home_bedroom_1",
  "timeStamp": "20170507111302",
  "readingAQ": 34,
  "celsius": 32,
  "farenheit": 90,
  "humidity": 70,
  "timeZone": "Asia/Kolkata",
  "error": ""
}
```
