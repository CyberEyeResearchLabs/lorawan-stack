---
title: "Messaging on Interrupt"
description: ""
weight: 4
---

Learn how to send a LoRaWAN packet on interrupt

<!--more-->

## Step by step

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

  if ( count % 2 == 0) {
    payload[0] = 1;
    digitalWrite(LED_BUILTIN, HIGH);
  }
  else {
    payload[0] = 0;
    digitalWrite(LED_BUILTIN, LOW);
  }

  if (digitalRead(4) == HIGH) {
    count ++;
    Serial.println(payload[0]);
    ttn.sendBytes(payload, sizeof(payload));
    delay(5000);
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