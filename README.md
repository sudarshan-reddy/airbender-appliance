## Airbender appliance

Appliance code that would run of a GrovePi+ integrated Raspberry Pi

### Prequisite:

*Hardware*
RaspberryPi
GrovePi+
AirQuality Sensor v1.3

*Software*
Raspbian
`github.com/DexterInd/GrovePi`

### Installation: 
 ```bash 
 git clone github.com/sudarshan-reddy/airbender-appliance
 ```

### Run: 

#### Compile for Raspberry Pi:
    ```bash
    GOOS=linux GOARCH=arm go build
    ```
    scp binary and `variables.env` to pi

#### Running the binary:
    ```bash
    source variables.env
    run ./airbender-appliance
    ```
