---
title: "Getting Device Information"
description: ""
weight: 4
---

To communicate via The Things Stack, you need to register your device. For this, we require the Device's unique EUI (a unique address hardcoded into the LoRa module). 

To get the device EUI, follow the below steps.

Select **File -> Examples -> TheThingsNetwork -> [DeviceInfo](https://github.com/TheThingsNetwork/arduino-device-lib/blob/master/examples/DeviceInfo/DeviceInfo.ino)** to open the DeviceInfo example code.

{{< figure src="001_getting_device_info.png" alt="DeviceInfo example code" >}}

Replace **REPLACE_ME** with **TTN_FP_EU868**.

The final code will look like following:

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

The above-mentioned code snippet does the following:

- Uses [#define](https://www.arduino.cc/en/Reference/Define) to create more meaningful aliases for the [Serial](https://www.arduino.cc/en/Reference/Serial) ports for the LoRa modem, and the  USB connection, and as well as for the frequency plan.
- Creates an instance tts of TheThingsNetwork class, passing the serial ports and the frequency plan as arguments to the instance.
- Calls [begin()](https://www.arduino.cc/en/Serial/Begin) to set the data rate for both serial ports.
- Uses [println()](https://www.arduino.cc/en/Serial/Println) to print on the Serial Monitor and calls **tts.showStatus()** to get the device information to be used in the further steps.

In the Arduino IDE, select **Sketch -> Upload** - **Ctrl+U (In Windows)** or **⌘+U (In Mac OS)** - to upload the sketch.

>*Note:* Uploads might fail if the Serial Monitor is open or if the IDE loses track of the port you selected in the 'Connect your Device' section. Close the Serial Monitor, check the port selection and try again. If it still fails, check Arduino Troubleshooting.

Select **Tools -> Serial Monitor**, **Ctrl+Shift+M (in Windows)** / **⌘+Shift+M (in Mac OS)** to open the [Serial Monitor](https://www.arduino.cc/en/Guide/Environment#toc12).

>**Note:** The serial monitor is the 'tether' between the computer and your The Things Uno. It lets you send and receive text messages, and comes handy in debugging and controlling The Things Uno from a keyboard.

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

We will use the **EUI** (which is meant to be the **Device EUI**) value to register our device.

We have now successfully connected the device, uploaded the sketch into the device and retrieved Device EUI from the logs in the Serial Monitor. 

Follow the next section for adding an application in The Things Stack.