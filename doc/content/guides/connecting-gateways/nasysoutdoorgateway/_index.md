---
title: "NASYS LoRaWAN® Outdoor Gateway"
description: ""
weight: 1
---

Technical information regarding the LoRaWAN® Outdoor Gateway can be found in [the official product page][1]. This page guides you to connect it to {{% tts %}}.

## Prerequisites

1. User account on {{% tts %}} with rights to create Gateways.
2. LoRaWAN® Outdoor Gateway connected to the Internet (or your local network) via Ethernet.

## Registration

Create a gateway by following the instructions for the [Console]({{< ref "/guides/getting-started/console#create-gateway" >}}) or the [CLI]({{< ref "/guides/getting-started/cli#create-gateway" >}}). Typically, the **EUI** field for your gateway should exist on the sticker at the bottom.

## Configuration using a Terminal

Find the IP address the gateway. This can be done in various ways. You can connect your machine to the same local network as that of the gateway Ethernet connection and scan for open SSH ports or assign a static IP to the gateway and use that. Once the gateway IP address is found, ssh into it.

```bash
$ ssh root@<GatewayIP>
```

The default username is **root**, and the default password can be also found in the sticker. It is typically suggested that you change the default password using `passwd` after logging in.

Your gateway should come with a slightly modified version of the [Lora-net UDP packet forwarder](https://github.com/Lora-net/packet_forwarder) pre-installed at `/opt/nas-lgw`. There are two configuration files `global_conf.json` and `local_conf.json`, both located in `/opt/nas_lgw`.

Typically, you should not need to touch the `global_conf.json` file. Instead, edit `local_conf.json`

```bash
$ cp /opt/nas-lgw/local_conf.json /opt/nas-lgw/local_conf.json.bak
$ vi /opt/nas-lgw/local_conf.json
```

Edit the server parameters.

1. **gateway_ID**: Set to the gateway EUI (in lowercase).
2. **server_address**: Address of the Gateway Server. If you followed the [Getting Started guide]({{< ref "/guides/getting-started" >}}) this is the same as what you use instead of `https://thethings.example.com`.
3. **serv_port_up**: UDP upstream port of the Gateway Server, typically 1700.
4. **serv_port_down**: UDP downstream port of the Gateway Server, typically 1700.

You can also update the rest of the parameters, as you see fit. For example, update the **ref_latitude**, **ref_longitude** and **ref_altitude** parameters to update the Fake GPS location reported by the gateway in its status messages.

Save the configuration and restart the packet forwarder:

```bash
$ systemctl restart nas_lgw
```

If your configuration was successful, your gateway will connect to {{% tts %}} after a couple of seconds. You can easily check if your gateway has connected using the Console.

The rest of the document discusses further topics that are not required for the gateway to work, but may be useful for your deployments.


## Troubleshooting

If the gateway does not connect to {{% tts %}} after a few minutes, issue a `reboot` command, or disconnect and reconnect the power supply to power-cycle the gateway.

You can access the gateway system logs using journalctl. Use the `-f` flag to follow the journal (i.e. do not exit and display new logs as they arrive), `-u nas_lgw` to only show the logs related to the packet forwarder, and `-n 1000` to show the last 1000 log lines.

```bash
$ journalctl [-f] [-u nas_lgw] [-n 1000]
```


## Extras: VPN

Your gateway will most likely be installed on the field and connect to the network server via the Internet (using cellular data or an Ethernet connection). Most likely this means that your gateway will be sitting behind a NAT, so connecting to it directly (for example, trying to SSH should something break) is impossible.

Thankfully, your gateway comes with a pre-installed OpenVPN client. In order to connect the gateway to your VPN, first generate a new client configuration file for your server and name it `CLIENT.conf`. Setting up the VPN server itself and generating the client configuration file is out-of-scope for this document.

Copy the file to the gateway with `scp`:

```bash
$ scp CLIENT.conf root@<GatewayIP>:CLIENT.conf
```

And then, enable the openvpn client for your configuration:

```bash
$ ssh root@<GatewayIP>
> systemctl start openvpn@CLIENT.service
> systemctl status openvpn@CLIENT.service    # <--- verify that it works
> systemctl enable openvpn@CLIENT.service    # <--- enable so that the service will restart on reboot
```

Your gateway should now be connected to your VPN server.

> NOTE: The installed OpenVPN client version may be old and not support the `tls-crypt` configuration command. You need to use `tls-auth` instead. You have to configure your VPN server accordingly.


## Extras: Logging

The gateway logs rotate when they reach about 15M in size, which means that you will generally not be able to access very old logs. At times of dense traffic (e.g. ~1000s of devices) this typically means that you only have logs available for the previous 2 hours. If you want to keep historical data (for whatever reason), then you can forward the logs to an other server using `netcat`:

```bash
$ journalctl -f | nc server-hostname port-number &
```


## Extras: Monitoring

The gateway comes with `nrpe` pre-installed. This means that you can connect it out of the box with your Nagios installation. The following steps are required for this:

```bash
$ vi /opt/nrpe/nrpe_local.cfg
```

Edit the **allowed_hosts** variable and add your Nagios host. Then, create a new configuration file for Nagios with the following contents:

<details><summary># gateway-checks.conf</summary>
<pre>
# gateway-checks.conf

define host {
     use                           generic-host
     statusmap_image               base/ubuntu.gd2
     icon_image_alt                Ubuntu
     icon_image                    base/ubuntu.png
     host_name                     YOUR_GATEWAY_NAME            # change this
     address                       YOUR_GATEWAY_ADDRESS         # change this
     vrml_image                    ubuntu.png
     check_interval                10
     max_check_attempts            1
}

define command {
    command_name            check-disk-command-nrpe
    command_line            /usr/lib/nagios/plugins/check_nrpe -H "$HOSTADDRESS$" -c "check_disk" -t 30 -p 5666
}

define service {
    use                     generic-service
    host_name               YOUR_GATEWAY_NAME
    check_command           check-disk-command-nrpe
    description             check-disk
    max_check_attempts      1
    check_interval          10
}

define command {
    command_name            check-nas-lgw-command-nrpe
    command_line            /usr/lib/nagios/plugins/check_nrpe -H "$HOSTADDRESS$" -c "check_service_nas-lgw" -t 30 -p 5666
}

define service {
    use                     generic-service
    host_name               YOUR_GATEWAY_NAME
    check_command           check-nas-lgw-command-nrpe
    description             check-nas-lgw
    max_check_attempts      1
    check_interval          10
}


define command {
    command_name            check-temp-command-nrpe
    command_line            /usr/lib/nagios/plugins/check_nrpe -H "$HOSTADDRESS$" -c "check_temp" -t 30 -p 5666
}

define service {
    use                     generic-service
    host_name               YOUR_GATEWAY_NAME
    check_command           check-temp-command-nrpe
    description             check-temp
    max_check_attempts      1
    check_interval          10
}
</pre>
</details>

Add this file under `/etc/nagios3/conf.d`, restart Nagios and you should have Nagios checks for your gateway's free disk space, temperature, and the packet forwarder service.


## References / Useful links

1. [Official product page][1]
2. [Quick start guide][2]


[1]: https://www.nasys.no/product/lorawan-gateway/ "Official Product Page"
[2]: https://www.nasys.no/wp-content/uploads/lorawan_gateway_quick_start_guide_1.pdf "Quick start guide"
