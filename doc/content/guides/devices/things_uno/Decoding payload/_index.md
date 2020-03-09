---
title: "Decoding Payload"
description: ""
weight: 5
---

Learn how to decode the payload in {{%tts%}}

<!--more-->


## Step by step

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