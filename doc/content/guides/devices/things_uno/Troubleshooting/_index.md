---
title: "Troubleshooting"
description: ""
weight: 7
---

Resolving common issues.

<!--more-->

## Troubleshooting

If you cannot see the packets in {{%tts%}} server:

- Check whether you have a gateway running with a frequency configured to the EU band (EU868) (or the relevant Frequency Band) and the forwarding server configured to {{%tts%}}.
- Check whether the Keys in the Arduino code match the Keys in {{%tts%}} console.
- Ensure that the gateway is not close to The Things Uno. A distance of more than 3m should be maintained between the gateway and The Things Uno.

If the join request is not getting accepted by {{%tts%}} in case of OTAA activation method:

- Ensure that The Things Uno is not immediately below the gateway.
