---
title: "Activation"
description: ""
weight: 3
---

Activating the device using different activation methods

<!--more-->

## Step by step

All the communication in LoRaWAN is done with a dynamic 32bit device address (DevAddr) of which 7 bits are fixed for {{%tts%}}, leaving 25 bits that can be assigned to individual devices, through a procedure called **Activation**.

For the Activation of a device, we have the following two methods in LoRaWAN:

1. Over The Air Activation (OTAA)
2. Activation by personalization (ABP)

Follow the below sub-sections to understand these activation methods.

### Activation of the device using OTAA

Over-the-Air Activation (OTAA) is the preferred and most secure way to connect with {{%tts%}}. Devices perform a join-procedure with the network, during which a dynamic `DevAddr` is assigned and security keys are negotiated with the device.

To get your device activated using the OTAA method, do the following:

#### Device Registration

On the application’s screen, select **Devices** from the bottom right menu.

In the **Devices** section, click **+ Add Device**.

- For Device ID, choose a unique ID of lower case, alphanumeric characters and nonconsecutive `-` and `_`
  e.g.: my_new_device, device_1.
- For `Device Name`, give any name you desire according to the device.
- For `Device Description`, add the description as desired.
- For `MAC Version`, select MAC V1.0.2
- For `PHY Version`, select PHY V1.0.2 REV B
- For `Frequency Plan`, select Europe 863-870 MHz, (or the Frequency Plan you are currently in)
- For `Network Server Address`, put the URL of your deployment. For example: `https://thethings.example.com/`
- For `Application Server Address`, put the same as the domain above.
- For `Activation Mode`, select Over The Air Activation (OTAA)
- For `Join EUI`, put any 16 character string which will work as App EUI in Arduino Sketch.
- For `Dev EUI`, use the Dev EUI from the device information retrieved in the `Get your Device Information` section.
- For `External Join Server`, uncheck the box.

>Note: The `Frequency Plan` may vary based on the country you are in. Check for the appropriate Frequency Plan for your region [here](https://www.thethingsnetwork.org/docs/lorawan/frequency-plans.html).
>Also, if the `External Join Server` is not disabled, you will not get `AppKey` in the device overview page.

{{< figure src="001_adding_device.png" alt="Device registration page" >}}

{{< figure src="004_adding_device.png" alt="Device registration for Activation using OTAA" >}}

Click on **Create Device** to add the device.

You will be redirected to the newly registered device, where you can find the generated keys under the **Session Information** section, which we will need in the further steps.

{{< figure src="003_adding_device.png" alt="Device overview page" >}}

#### Device Activation

Now that you have registered the device, you can activate the connection from your device itself.

Go to **File -> Examples -> TheThingsNetwork -> SendOTAA** and click on it to open the `SendOTAA` sketch.

{{< figure src="001_otaa_activation.png" alt="SendOTAA example code location" >}}

Replace the `appEui` and `appKey` with the keys obtained from {{%tts%}} console.

>Note: `appEui` is the `JoinEUI` in {{%tts%}} console.

Replace `REPLACE_ME` with one of the lines below, depending on the frequency plan of your device and your country.

- `TTN_FP_EU868` (Europe, Middle East, Africa)
- `TTN_FP_US915` (Americas, except Brazil)
- `TTN_FP_AU915` (Oceania, Brazil)
- `TTN_FP_IN865_867` (India)
- `TTN_FP_KR920_923` (Korea)
- `TTN_FP_AS920_923` (Japan, Singapore, Malaysia)
- `TTN_FP_AS923_925` (Southeast Asia)

>Note: The code below uses `TTN_FP_EU868`

The final code should look like this:

```bash
#include <TheThingsNetwork.h>

// Set your AppEUI and AppKey
const char *appEui = "B07A78992F8789F8";
const char *appKey = "D708031B2BBAB3345645B8E11DE5E18";

#define loraSerial Serial1
#define debugSerial Serial

// Replace REPLACE_ME with TTN_FP_EU868 or TTN_FP_US915
#define freqPlan TTN_FP_EU868

TheThingsNetwork ttn(loraSerial, debugSerial, freqPlan);

void setup()
{
  loraSerial.begin(57600);
  debugSerial.begin(9600);

  // Wait a maximum of 10s for Serial Monitor
  while (!debugSerial && millis() < 10000)
    ;

  debugSerial.println("-- STATUS");
  ttn.showStatus();

  debugSerial.println("-- JOIN");
  ttn.join(appEui, appKey);
}

void loop()
{
  debugSerial.println("-- LOOP");

  // Prepare payload of 1 byte to indicate LED status
  byte payload[1];
  payload[0] = (digitalRead(LED_BUILTIN) == HIGH) ? 1 : 0;

  // Send it off
  ttn.sendBytes(payload, sizeof(payload));

  delay(10000);
}
```

Select **Sketch -> Upload** to upload the sketch and then **Tools -> Serial Monitor** to open the Serial Monitor.

The Serial Monitor output should look similar to this:

```bash
-- STATUS
EUI: 0004A307771BB05B
Battery: 3253
AppEUI: B07A78992F8789F8
DevEUI: 0004A307771BB05B
Data Rate: 5
RX Delay 1: 1000
RX Delay 2: 2000
-- JOIN
Model: RN2483
Version: 1.0.1
Sending: mac set deveui 0004A307771BB05B
Sending: mac set adr off
Sending: mac set deveui 0004A307771BB05B
Sending: mac set appeui B07A78992F8789F8
Sending: mac set appkey D708031B2BBAB3345645B8E11DE5E18
Sending: mac save
Sending: mac set rx2 3 869525000
Sending: mac set ch drrange 1 0 6
Sending: mac set ch dcycle 0 799
Sending: mac set ch dcycle 1 799
Sending: mac set ch dcycle 2 799
Sending: mac set ch dcycle 3 799
Sending: mac set ch freq 3 867100000
Sending: mac set ch drrange 3 0 5
Sending: mac set ch status 3 on
Sending: mac set ch dcycle 4 799
Sending: mac set ch freq 4 867300000
Sending: mac set ch drrange 4 0 5
Sending: mac set ch status 4 on
Sending: mac set ch dcycle 5 799
Sending: mac set ch freq 5 867500000
Sending: mac set ch drrange 5 0 5
Sending: mac set ch status 5 on
Sending: mac set ch dcycle 6 799
Sending: mac set ch freq 6 867700000
Sending: mac set ch drrange 6 0 5
Sending: mac set ch status 6 on
Sending: mac set ch dcycle 7 799
Sending: mac set ch freq 7 867900000
Sending: mac set ch drrange 7 0 5
Sending: mac set ch status 7 on
Sending: mac set pwridx 1
Sending: mac set retx 7
Sending: mac set dr 5
Sending: mac join otaa
Join accepted. Status: 00000401
DevAddr: 27000019
```

The Join requests and the messages on {{%tts%}} Console should look like the following:

{{< figure src="001_device_activation.png" alt="Uplink messages in {{%tts%}} console" >}}

Your device is now activated using OTAA Activation method and is connected to {{%tts%}}.

You can now directly proceed with the **Sending Message on Interrupt** section to learn how to add a push button and send messages on pressing the button.

Alternatively, you can follow the next section for activating the device using the ABP method.

### Activation of the device using ABP

In some cases you might need to hardcode the DevAddr as well as the security keys in the device. This means activating a device by personalization (ABP). This strategy might seem simpler, because you skip the join procedure, but it has some downsides related to security.

To get your device activated using the ABP method, follow the steps below:

#### Device Registration

In the application’s screen, select **Devices** from the bottom right menu.

In the **Devices** box, click **+ Add Device**.

- For `Device ID`, choose a unique ID of lower case, alphanumeric characters and nonconsecutive `-` and `_`.
- For `Device Name`, give any name you desire according to the device.
- For `Device Description`, add the description as desired.
- For `MAC Version`, select MAC V1.0.2
- For `PHY Version`, select PHY V1.0.2 REV B
- For `Frequency Plan`, select Europe 863-870 MHz, (or the Frequency Plan you are currently in)
- For `Network Server Address`, put the domain of {{%tts%}}. For example, `https://thethings.example.com/`
- For `Application Server Address`, put the same as the domain above.
- For `Activation Mode`, select Activation By Personalization (ABP).
- For `Device Address`, you can choose any combination of 8 characters having letters and numbers.

{{< figure src="001_adding_device.png" alt="Device registration page" >}}

{{< figure src="002_adding_device.png" alt="Registration of device using ABP activation" >}}

Click on **Create Device** to add the device.

You will be redirected to the newly registered device, where you can find the generated Keys under the **Session Information** section which we will need in further steps.

{{< figure src="003_adding_device.png" alt="Device overview page" >}}

#### Device Activation

Now that you have registered the device, you can activate the connection from your device itself.

Go to **File -> Examples -> TheThingsNetwork -> SendABP** and click on it to open the example code.

{{< figure src="001_abp_activation.png" alt="SendABP example code location" >}}

Replace `devAddr`, `nwkSKey`, and `appSKey` with the keys obtained from the device page in {{%tts%}} console.

> Note: `nwkSKey` can be replaced with the `NwkSEncKey` from {{%tts%}} console.

Replace `REPLACE_ME` with one of the lines below, depending on the frequency plan of your device and your country.

- `TTN_FP_EU868` (Europe, Middle East, Africa)
- `TTN_FP_US915` (Americas, except Brazil)
- `TTN_FP_AU915` (Oceania, Brazil)
- `TTN_FP_IN865_867` (India)
- `TTN_FP_KR920_923` (Korea)
- `TTN_FP_AS920_923` (Japan, Singapore, Malaysia)
- `TTN_FP_AS923_925` (Southeast Asia)

>Note: The code below uses `TTN_FP_EU868`

The final code should look like the following:

```bash
#include <TheThingsNetwork.h>

// Set your DevAddr, NwkSKey, AppSKey and the frequency plan
const char *devAddr = "2700006C";
const char *nwkSKey = "A41B7BD15237BE21ED931D1E6ED9EF2C";
const char *appSKey = "F7A7BE235FCB5BADEC505B79B4C729BA";

#define loraSerial Serial1
#define debugSerial Serial

// Replace REPLACE_ME with TTN_FP_EU868 or TTN_FP_US915
#define freqPlan TTN_FP_EU868

TheThingsNetwork ttn(loraSerial, debugSerial, freqPlan);

void setup()
{
  loraSerial.begin(57600);
  debugSerial.begin(9600);

  // Wait a maximum of 10s for Serial Monitor
  while (!debugSerial && millis() < 10000)
    ;

  debugSerial.println("-- PERSONALIZE");
  ttn.personalize(devAddr, nwkSKey, appSKey);

  debugSerial.println("-- STATUS");
  ttn.showStatus();
}

void loop()
{
  debugSerial.println("-- LOOP");

  // Prepare payload of 1 byte to indicate LED status
  byte payload[1];
  payload[0] = (digitalRead(LED_BUILTIN) == HIGH) ? 1 : 0;

  // Send it off
  ttn.sendBytes(payload, sizeof(payload));

  delay(10000);
}
```

Select **Sketch -> Upload** to upload the sketch and then **Tools -> Serial Monitor** to open the Serial Monitor.

The code will activate the device and send the status of the built-in LED to {{%tts%}} server.

The Serial Monitor output should look similar to this:

```bash
-- PERSONALIZE
Model: RN2483
Version: 1.0.4
Sending: mac set deveui 0004A30B001BC88B
Sending: mac set adr off
Sending: mac set devaddr 2700006C
Sending: mac set nwkskey A41B7BD15237BE21ED931D1E6ED9EF2C
Sending: mac set appskey F7A7BE235FCB5BADEC505B79B4C729BA
Sending: mac set rx2 3 869525000
Sending: mac set ch drrange 1 0 6
Sending: mac set ch dcycle 0 799
Sending: mac set ch dcycle 1 799
Sending: mac set ch dcycle 2 799
Sending: mac set ch dcycle 3 799
Sending: mac set ch freq 3 867100000
Sending: mac set ch drrange 3 0 5
Sending: mac set ch status 3 on
Sending: mac set ch dcycle 4 799
Sending: mac set ch freq 4 867300000
Sending: mac set ch drrange 4 0 5
Sending: mac set ch status 4 on
Sending: mac set ch dcycle 5 799
Sending: mac set ch freq 5 867500000
Sending: mac set ch drrange 5 0 5
Sending: mac set ch status 5 on
Sending: mac set ch dcycle 6 799
Sending: mac set ch freq 6 867700000
Sending: mac set ch drrange 6 0 5
Sending: mac set ch status 6 on
Sending: mac set ch dcycle 7 799
Sending: mac set ch freq 7 867900000
Sending: mac set ch drrange 7 0 5
Sending: mac set ch status 7 on
Sending: mac set pwridx 1
Sending: mac set retx 7
Sending: mac set dr 5
Sending: mac join abp
Personalize accepted. Status: 00000001
-- STATUS
EUI: 0004A30B001BC88B
Battery: 3325
AppEUI: 70B3D57ED001ACE0
DevEUI: 0004A30B001BC88B
Data Rate: 5
RX Delay 1: 1000
RX Delay 2: 2000
-- LOOP
Sending: mac tx uncnf 1 00
Successful transmission
```

Go to the **Data** tab of the **Devices** section in {{%tts%}} Console to see if the data is being sent to the server.

{{< figure src="001_device_activation.png" alt="Received uplinks in The Things Stack console" >}}

We have completed the activation of our device using ABP activation method.