---
title: "Arduino IDE Setup"
description: ""
weight: 1
---

Learn how to setup Arduino IDE

<!--more-->

## The Things Uno

[The Things Uno](https://www.thethingsnetwork.org/docs/devices/uno/) is the perfect board to start prototyping your IoT ideas or make your existing project wireless with up to 10km range by simply swapping boards.

Programming and adding The Things Uno to The Things Enterprise Stack V3 ({{%tts%}}) is almost the same as with the previous version and involves only minor modifications.

{{< figure src="001_things_uno.png" alt="The Things Uno" >}}

## Prerequisites

- Knowledge of Arduino IDE and related concepts
- Understanding of basic electronic concepts

### Setting up the Arduino IDE

The [Arduino Integrated Development Environment (IDE)](https://www.arduino.cc/en/main/software) is a cross-platform application (for Windows, macOS and Linux) used to write and upload programs to Arduino compatible boards.

Let us start by setting up the software development environment to program The Things Uno.

1. [Download](https://www.arduino.cc/en/Main/Software) and install the latest Arduino Software (IDE).
2. Navigate to **Sketch -> Include Library -> Manage Libraries** to open the library manager.
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