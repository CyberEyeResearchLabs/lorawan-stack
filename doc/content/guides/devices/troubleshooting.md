---
title: "Troubleshooting"
description: ""
weight: 10
---

If you cannot see the packets in The Things Stack server:

- Check whether you have a gateway running with a frequency configured to the EU band (EU868) and the forwarding server configured to The Things Stack.
- Check whether the Keys in the Arduino code match the Keys in The Things Stack console.
- Check whether the gateway is kept close to The Things Uno . A distance of more than 3m should be maintained between the gateway and The Things Uno.

If the join request is not getting accepted by The Things Stack in case of OTAA activation method:

- Ensure that The Things Uno is not immediately below the gateway.