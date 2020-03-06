---
title: "The Things Uno"
description: ""
weight: 3
---

[The Things Uno](https://www.thethingsnetwork.org/docs/devices/uno/) is the perfect board to start prototyping your IoT ideas or make your existing project wireless with up to 10km range by simply swapping boards.

Programming and adding The Things Uno to The Things Enterprise Stack V3 ({{%tts%}}) is almost the same as with the previous version and involves only minor modifications.

<!--more-->
{{< figure src="001_things_uno.png" alt="The Things Uno" >}}

This document provides you with a stepwise process for adding a device (in the current case, The Things Uno) to {{%tts%}} server and programming it to send and receive data via {{%tts%}} using LoRaWAN<sup>®</sup>.

## Prerequisites

- Knowledge of Arduino IDE and related concepts
- Understanding of basic electronic concepts

### Setting up the Arduino IDE

The [Arduino Integrated Development Environment (IDE)](https://www.arduino.cc/en/main/software) is a cross-platform application (for Windows, macOS and Linux) that is used to write and upload programs to Arduino compatible boards.

Let us start by setting up the software development environment to program The Things Uno.

1. [Download](https://www.arduino.cc/en/Main/Software) and install the latest Arduino Software (IDE).
2. Navigate to **Sketch -> Include Library -> Manage Libraries** to open library manager.
3. Search for **TheThingsNetwork** and click on the result to select it.
4. Click on the **Install** button to install it.

{{< figure src="001_arduino_library.png" alt="The Things Network library" >}}

> Note: The Arduino IDE will notify you of updates for the IDE and library automagically.

### Connecting The Things Uno to a Computer

To connect your device, proceed with the following steps.

- Use a Micro-USB cable to connect The Things Uno to a USB port of your computer.
- In Arduino IDE, select **Tools -> Board -> Arduino Leonardo**.

{{< figure src="001_arduino_ide.png" alt="Arduino Leonardo Board" >}}

- Navigate to **Tools -> Port** and select the port that identifies as Arduino Leonardo.

{{< figure src="002_arduino_ide_port.png" alt="COM Port for Arduino Leonardo Board" >}}

>If you don't see a port that identifies as Arduino Leonardo, make sure that The Things Uno's power LED is on, and check the cable and USB port you have used. On Windows, you might need to [install drivers](https://www.arduino.cc/en/Guide/ArduinoLeonardoMicro#toc2). See [Arduino Troubleshooting](https://www.arduino.cc/en/Guide/Troubleshooting#toc16) for more suggestions.

## Getting the Device Information

To communicate via {{%tts%}}, you need to register your device. For this, you would require the Device's unique EUI (a unique address hardcoded into the LoRa module). To get the device EUI, do the following:

Select **File -> Examples -> TheThingsNetwork -> [DeviceInfo](https://github.com/TheThingsNetwork/arduino-device-lib/blob/master/examples/DeviceInfo/DeviceInfo.ino)** to open the `DeviceInfo` example code.

{{< figure src="001_getting_device_info.png" alt="DeviceInfo example code" >}}

Replace `REPLACE_ME` in the code with `TTN_FP_EU868` (if you are using the EU868 Frequency Plan).

The final code should look like the following:

```bash
#include <TheThingsNetwork.h>

#define loraSerial Serial1
#define debugSerial Serial

// Replace REPLACE_ME with TTN_FP_EU868 or TTN_FP_US915
#define freqPlan TTN_FP_EU868

TheThingsNetwork ttn(loraSerial, debugSerial, freqPlan);

void setup()
{
  loraSerial.begin(57600);
  debugSerial.begin(9600);
}

void loop()
{
  debugSerial.println("Device Information");
  debugSerial.println();
  ttn.showStatus();
  debugSerial.println();
  debugSerial.println("Use the EUI to register the device for OTAA");
  debugSerial.println("-------------------------------------------");
  debugSerial.println();

  delay(10000);
}
```

The above code snippet does the following:

- Uses [#define](https://www.arduino.cc/en/Reference/Define) to create more meaningful aliases for the [Serial](https://www.arduino.cc/en/Reference/Serial) ports for the LoRa modem, the  USB connection as well as for the frequency plan.
- Creates an instance `ttn` of `TheThingsNetwork` class, passing the serial ports and the frequency plan as arguments to the instance.
- Calls [begin()](https://www.arduino.cc/en/Serial/Begin) to set the data rate for both serial ports.
- Uses [println()](https://www.arduino.cc/en/Serial/Println) to print on the Serial Monitor and calls `ttn.showStatus()` to get the device information to be used in the further steps.

In the Arduino IDE, select **Sketch -> Upload** to upload the sketch.

>Note: Uploads might fail if the Serial Monitor is open or if the IDE loses track of the port you selected in the 'Connect your Device' section. When that happens, close the Serial Monitor, check the port selected and try again. If it still fails, check [Arduino Troubleshooting](https://www.arduino.cc/en/Guide/Troubleshooting).

Select **Tools -> Serial Monitor** to open the [Serial Monitor](https://www.arduino.cc/en/Guide/Environment#toc12).

>Note: The serial monitor is the 'tether' between your computer and The Things Uno. It lets you send and receive text messages, and comes handy in debugging and controlling The Things Uno from a keyboard.

A print log, similar to the one shown below, gets displayed in the Serial Monitor:

```bash
Device Information

EUI: 0004A30B001BB05B
Battery: 3253
AppEUI: 7878787F8787F8F8
DevEUI: 0004A30B001BB05B
Data Rate: 5
RX Delay 1: 1000
RX Delay 2: 2000
```

We will use the EUI (which is meant to be the Device EUI) value to register our device.

We have now successfully connected the device, uploaded the sketch into the device and retrieved Device EUI from the logs in the Serial Monitor.

## Add an Application to {{%tts%}}

Devices need to be registered with an application to communicate with {{%tts%}}. Let’s add one.

Login to `https://thethings.example.com/`. You will be taken to {{%tts%}} console.

{{< figure src="001_adding_application.png" alt="{{%tts%}} console home page" >}}

Click on **Go to Applications**, and then click on **+ Add Application** to reach the application registration page.

{{< figure src="002_adding_application.png" alt="Applications list page" >}}

Fill and set the required fields:

- For `Application ID`, choose a unique ID by following the guidelines below:
  - Lower case, alphanumeric characters and nonconsecutive ‘ - ’ 
  - Maximum 100 characters are allowed
- For `Application Name`, give any suitable name to the application
- For `Description`, add the desired description of the application
- For `Linking`, Leave the checkbox checked to link the application automatically to the network server,

Click on **Create Application** to finish.

{{< figure src="003_adding_application.png" alt="Application registration page" >}}

Once the application is created, you can find it in the Console under the **Applications** tab.

{{< figure src="004_adding_application.png" alt="Application overview page" >}}

## Activation

All the communication in LoRaWAN is done with a dynamic 32bit device address (DevAddr) of which 7 bits are fixed for {{%tts%}}, leaving 25 bits that can be assigned to individual devices, through a procedure called **Activation**.

For the Activation of a device, we have the following two methods in LoRaWAN:

1. Over The Air Activation (OTAA)
2. Activation by personalization (ABP)

Follow the below sub-sections to understand these activation methods.

### Activation of the device using OTAA

Over-the-Air Activation (OTAA) is the preferred and most secure way to connect with {{%tts%}}. Devices perform a join-procedure with the network, during which a dynamic `DevAddr` is assigned and security keys are negotiated with the device.

To get your device activated using the OTAA method, do the following.

#### Device Registration

You are now ready to register your device to the application for activating it using the OTAA method.

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

You will be redirected to the newly registered device, where you can find the generated keys under the &quot;Session Information&quot; section, which we will need in the further steps.

{{< figure src="003_adding_device.png" alt="Device overview page" >}}

#### Device Activation

Now that you have registered the device, you can activate the connection from your device itself.

Go to **File -> Examples -> TheThingsNetwork -> SendOTAA** and click on it to open the SendOTAA sketch.

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

You can now directly proceed with the &quot;Sending Message on Interrupt&quot; section to learn how to add a push button and send messages on pressing the button.

Alternatively, you can follow the next section for activating the device using the ABP method.

### Activation of the device using ABP

In some cases you might need to hardcode the DevAddr as well as the security keys in the device. This means activating a device by personalization (ABP). This strategy might seem simpler, because you skip the join procedure, but it has some downsides related to security.

To get your device activated using the ABP method, follow the steps below.

#### Device Registration

The following steps will walk you through the process of adding the device to be activated using the ABP activation method:

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

You will be redirected to the newly registered device, where you can find the generated Keys under the &quot;Session Information&quot; section which we will need in further steps.

{{< figure src="003_adding_device.png" alt="Device overview page" >}}

#### Device Activation

Now that you have registered the device, you can activate the connection from your device itself.

Go to **File -> Examples -> TheThingsNetwork -> SendABP** and click on it to open the example code.

{{< figure src="001_abp_activation.png" alt="SendABP example code location" >}}

Replace `devAddr`, `nwkSKey`, and `appSKey` with the keys obtained from the device page in {{%tts%}} console.

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

## Sending Message on Interrupt

To make it interactive, let us create a setup so that we can toggle the onboard LED and send the status to {{%tts%}} on pressing a button.

The following is a simple breadboard based circuit to add a button to The Things Uno board. Make the connections to The Things Uno as per the circuit shown below:

{{< figure src="001_adding_button.png" alt="Hardware connections of the button to The Things Uno" >}}

Connect the wires as follows:

- Connect the red wire to 5V pin
- Connect the black wire to GND pin
- Connect the yellow wire to 4th Digital pin

Once the hardware connections are done, it’s time to write the code in Arduino.

Replace the **loop()** section of the code snippet used for device activation with the following:

```bash
void loop(){
  static int count = 0;
  byte payload[1];

  if (digitalRead(4) == HIGH) {
    count ++;
    Serial.println(payload[0]);
    ttn.sendBytes(payload, sizeof(payload));
    delay(5000);
  }

  if ( count % 2 == 0) {
    payload[0] = 1;
    digitalWrite(LED_BUILTIN, HIGH);
  }
  else {
    payload[0] = 0;
    digitalWrite(LED_BUILTIN, LOW);
  }
}
```

The above code snippet does following:

- `digitalRead(4)` reads the status of the button connected on digital pin 4.
- After each button press, `count++` will increase the count by one unit.
- `count%2 == 0` will check whether the count value is even or odd.
- The `payload[0]` value will be toggled according to the button press.

Select **Sketch -> Upload** to upload the sketch and then **Tools -> Serial Monitor** to open the Serial Monitor.

The output on Serial Monitor should appear as below after pressing the button twice:

```bash
1
Sending: mac tx uncnf 1 01
Successful transmission
0
Sending: mac tx uncnf 1 00
Successful transmission
```

Go to the **Data** section of your device page in {{%tts%}} Console to see the data being received in the server.

{{< figure src="001_device_activation.png" alt="Received uplinks in The Things Stack console" >}}

We receive the status of the LED in Base64 encoded form in {{%tts%}}.

## Decoding the Message Payload

What you see in the **Data** section of the device in {{%tts%}} Console as below, is the raw payload in base64 format.

{{< figure src="001_uplink_message.png" alt="Received message payload in JSON format" >}}

Let us decode that into a meaningful message.

Go to **Payload Formatters** tab in {{%tts%}} Console.

In the `Formatter Type` section, select `Javascript`.

{{< figure src="001_decoding_data.png" alt="Adding payload formatter code" >}}

In the `Formatter Parameter` box, paste the following code snippet:

```bash
function Decoder(bytes, port) {
 // Decode an uplink message from a buffer
 // (array) of bytes to an object of fields.
 var decoded = {};

 if (port === 1) decoded.led = bytes[0];

 return decoded;
}
```

The above code snippet does the following:

- `decoded.led = bytes[0]` will decode the base64 encoding into base8 integer.
- `return decoded` will return the decoded message.

Go to the **Data** tab and click on any one of the recently received packets. You should be able to see the actual LED status in it as shown below.

{{< figure src="002_decoding_data.png" alt="Decoded message" >}}

Once you get the LED status in its final decoded form in {{%tts%}} Console, you can follow the next section to add an MQTT Integration to get the status of the LED in a custom application.

## Node-Red Integration

> TODO: [MQTT Integration to a third party application]

## Troubleshooting

If you cannot see the packets in {{%tts%}} server:

- Check whether you have a gateway running with a frequency configured to the EU band (EU868) (or the relevant Frequency Band) and the forwarding server configured to {{%tts%}}.
- Check whether the Keys in the Arduino code match the Keys in {{%tts%}} console.
- Ensure that the gateway is not close to The Things Uno. A distance of more than 3m should be maintained between the gateway and The Things Uno.

If the join request is not getting accepted by {{%tts%}} in case of OTAA activation method:

- Ensure that The Things Uno is not immediately below the gateway.
