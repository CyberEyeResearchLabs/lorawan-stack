---
title: "MultiTech Conduit AEP"
description: ""
weight: 8
---

## Connecting MultiTech Conduit to {{%tts%}}

The [MultiConnect® Conduit™](http://www.multitech.net/developer/products/multiconnect-conduit-platform/) is a configurable, scalable cellular communications gateway for industrial IoT applications. The technical specifications of this conduit can be found in the **[official documentation](https://www.multitech.com/documents/publications/software-guides/s000727%20-%20mPower%20Edge%20Intelligence%20Conduit_AEP_software_guide.pdf).**

This guide will help you set up the gateway to communicate over **{{%tts%}}**.

{{< figure src="001_Multitech_Conduit_AEP.png" alt="multitech Conduit AEP" >}}

<!--more-->

<br>
### Prerequisites

-   MultiTech Conduit [AEP](http://www.multitech.net/developer/software/aep/) model. 
-   MultiTech [MTAC-LoRa LoRa accessory card](http://www.multitech.net/developer/products/accessory-cards/mtac-lora/), installed as [instructed](http://www.multitech.net/developer/products/accessory-cards/installing-an-accessory-card/).
Do not forget to mount the antenna to the mCard after fitting it in the conduit.
-   Computer with an Ethernet port.
-   LoRa Antenna - 1
-   LTE Antennas - 2
-   WiFi Antenna - 1
-   LAN Cable - 1
-   GPS Antenna - 1
-   Screwdriver
-   SIM Card - 1 (If you use cellular as backhaul)
<br>
{{< figure src="002_Hardware_Setup_Parts.png" alt="Hardware Setup Parts" >}}

<br>
## Hardware Setup:

- Attach the LoRa antenna to LoRa mCards’s RF connector in the AP2 slot located at the back of the conduit. Finger-tighten the antenna. (Ignore this step, if the antenna is already attached.)
- Similarly, attach GPS, WiFi, CELL and LoRa antennas.
- Attach the power supply to the Multitech Conduit.

<br>
{{< figure src="003_Multitech_Conduit_Ports.png" alt="Hardware Setup Parts" >}}
<br>
{{< figure src="004_Antenna_Assembled.png" alt="004_Antenna_Assembled.png" >}}

<br>

## Configuring the AEP Model

### Login to the Web Interface

To start configuring your conduit, connect a computer to its LAN interface and open the following URL in a browser: [http://192.168.2.1/](http://192.168.2.1/).  The URL redirects you to the commissioning window to set a Username and Password.

- The system requires you to set up an **Administrator** user.  Read the requirements carefully, set the desired username and password and click on OK. 

> This article addresses the minimum configuration required to connect your MultiTech Conduit AEP model to {{%tts%}}. You can find a detailed description of the components and their configurations, in this **[software guide](https://www.multitech.com/documents/publications/software-guides/s000727%20-%20mPower%20Edge%20Intelligence%20Conduit_AEP_software_guide.pdf).** 

<br>
{{< figure src="005_Setting_Login_Username.png" alt="Setting Login Username" >}}
<br>
{{< figure src="006_Setting_Login_Password.png" alt="Setting Login Password" >}}
<br>

-  Log into the conduit using your new username and password. 

{{< figure src="007_User_Login.png" alt="User Login" >}}
<br>


### 


### First-Time Setup Wizard


After you log in, you will be directed to First-Time Setup Wizard page.
This process takes a while. Be patient!

{{< figure src="008_First_Time_Setup_Wizard.png" alt="First Time Setup Wizard" >}}

Click on **Next** to begin.

<br>
### 


#### Call Home Remote Management

-  Soon after the setup process begins, the **Call Home Remote Management** window appears.


> The "Call Home" button enables the device to call home for configuration files, firmware updates, custom applications, and adds your DeviceHQ account key to the device. Also, clicking the Call Home button results in the device being reset to factory defaults. 

-   You can choose to either enable or disable the **Call Home Remote Management** feature as desired. 

{{< figure src="009_Call_Home_Remote_Management.png" alt="Call Home Remote Management" >}}

- Leaving the default settings as-is, click on **Next**.
- You can read more about the Call Home Remote Management feature **[here](http://www.multitech.net/developer/software/aep/call-home-remote-management/).**

### 
<br>
### 


#### Time Configuration

Select your time zone and update both date and time as required.

{{< figure src="010_Time_Configuration.png" alt="Time Configuration" >}}
<br>
Click on **Next** to proceed.

<br>

#### Network Interface Configuration

You can modify the network interfaces as desired. However, for the initial setup, you can leave the default settings as-is. 

{{< figure src="011_Network_Interface_Configuration.png" alt="Network Interface Configuration" >}}


Click on **Next** to proceed.

<br>

#### Cellular Setup:

- Set the Cellular Configuration to **Enabled** and add the corresponding APN if you are using Cellular as the backhaul. 

> APN stands for Access Point Name. The APN is assigned by your wireless service provider.

{{< figure src="012_Cellular_Configuration.png" alt="Cellular Configuration " >}}


- Click on **Next**, and the Cellular Authentication window appears. You can leave the default setting as-is. 


{{< figure src="013_Cellular Authentication.png" alt="Cellular Authentication" >}}

Click on **Next** again to proceed. 

<br>

#### Remote Management

If you prefer Remote Management, you can **Enable** the feature and provide the necessary Account Key.  To know more about how to configure Remote Management, you can refer to this **[software guide](https://www.multitech.com/documents/publications/software-guides/s000727%20-%20mPower%20Edge%20Intelligence%20Conduit_AEP_software_guide.pdf).**


{{< figure src="014_Remote_Management.png" alt="Remote Management" >}}

Click on **Next** to proceed.


#### 

<br>
#### HTTP/HTTPS Access

The device allows only secure access to its Web UI. This panel defines how the device handles HTTP traffic like, automatically redirecting HTTP requests to the device's secure HTTPS port, etc. You can choose **LAN** or **WAN** as desired.


{{< figure src="015_Access_Configuration.png" alt="Access Configuration" >}}

Click on **Next** to proceed.

#### 

<br>
#### Bootloader Protection

This feature enables password authentication to access the device bootloader. It is disabled by default.

{{< figure src="016_Bootloader_Protection.png" alt="Bootloader Protection" >}}

Click on **Finish**.
#### 

<br>
#### Adding {{%tts%}} Server Address in the Gateway
<br>

- Click on "**LoRaWAN**" in the menu on the left. It opens the Gateway's configuration page.

{{< figure src="017_001_Gateway_Menu_LoRaWAN.png" alt="Gateway Menu" >}}
<br>

- Select the mode as "**Packet Forwarder**" 
{{< figure src="018_Packet_Forwarder_Gateway.png" alt="Packet Forwarder Gateway" >}}

<br>

- Select your **frequency plan** (EU868 in this case.)

{{< figure src="019_Frequency_Plan.png" alt="Frequency Plan" >}}

| <ul><li>Scroll down and add:</li></ul> 	|  	|
|----------------------------------------------------------	|-------------------------	|
| {{%tts%}} URL provided to you as the **Server Address**. 	| `thethings.example.com` 	|
| Upstream Port 	| 1700 	|
| Downstream Port 	| 1700 	|

<br>

> If you followed the **[Getting Started guide](https://enterprise.thethingsstack.io/v3.4.0/guides/getting-started/)** , the domain of {{%tts%}} is the same as what you use instead of **https://thethings.example.com**. For example, \<domain\>.eu1.cloud.thethings.industries

<br>

{{< figure src="020_Server_Setup_Configuration.png" alt="Server Configuration" >}}


- Click on **Submit**.

<br>


#### 


#### Network Interfaces

Network interface shows the configuration for one or more network interfaces available.


- Select **Setup > Network Interfaces** from the left menu:

{{< figure src="021_Network_Interfaces.png" alt="Network Interfaces" >}}
<br>

- Click on the pencil icon(edit) in the "Options" column corresponding to **eth0**. You will be redirected to the Network Interface Configuration window. 
<br>
{{< figure src="022_Network_Interface_Configuration_Window.png" alt="Network Interface Configuration" >}}
<br>

- In the configuration window, change the **Bridge** from **br0** to **--**.
- In IPv4 Settings, change the **Mode** from **Static** to **DHCP Client**.
- Click on **Submit**.
- Now click on **Save and Restart** from the menu.

{{< figure src="017_002_Gateway_Menu_Save-And-Restart.png" alt="Gateway Menu" >}}

<br>

- You will be prompted to confirm the restart. Choose **OK** to proceed.

{{< figure src="023_System_Confirmation_Prompt.png" alt="System Confirmation" >}}


- Disconnect the temporary network cable and connect the Conduit to the target network while it restarts.

- Ensure that the target network has Internet access. The Gateway obtains its IP address from the target network after the restart.

To know more about other features of the MultiTech Conduit gateway, you can refer to this **[software guide](https://www.multitech.com/documents/publications/software-guides/s000727%20-%20mPower%20Edge%20Intelligence%20Conduit_AEP_software_guide.pdf).**


<br>
## Adding MultiTech Conduit to {{%tts%}} Console

- Login to **[https://thethings.example.com/](https://thethings.example.com/)**. 
- You will be taken to the {{%tts%}} console. 

> If you followed the **[Getting Started guide](https://enterprise.thethingsstack.io/v3.4.0/guides/getting-started/)** , the domain of {{%tts%}} is the same as what you use instead of **https://thethings.example.com**. For example, \<domain\>.eu1.cloud.thethings.industries


{{< figure src="025_The_Things_Stack_Console.png" alt="The Things Stack Console" >}}


- Go to **Gateways** in the top menu, and click on "**+ Add Gateway**" to reach the gateway registration page.

- Fill and set the required fields:
    1. Gateway ID
    2. Gateway EUI (Obtained during the gateway LoRaWAN settings configuration)
    3. Gateway Address (Use your default {{%tts%}} URL **thethings.example.com**)
    4. Enable/Disable Duty Cycle

- Click on **Create Gateway** to create the gateway. 

{{< figure src="026_Adding_a_Gateway.png" alt="Adding a Gateway" >}}

<br>

- Once created, you can find the newly added Gateway and the configuration under the “Gateways” section of the Console.

{{< figure src="027_Gateway_Console_Connection_Status.png" alt="Gateway Connection Status" >}}

<br>

## Troubleshooting

### Packet Forwarder Logs
You can find the packet forwarder logs using the gateway&apos;s web console. Do the following to download the logs onto your device:

- Login to your gateway&apos;s web console.
- Navigate to **Administration -> Debug Options**
- Click on the **Download Logs** button in the **Logging** section to download the logs.

{{< figure src="028_Packet_Forwarder.png" alt="Packet forwarder" >}}

<br>

- Now, in the downloaded logs, open the **lora-pkt-fwd-1.log** file to see the packet forwarder logs.

<br>

### Upgrading the Firmware


Refer to the top of your configuration software window to check your firmware version. You have to upgrade the device's firmware to the latest version.

You can download firmware upgrades from the downloads section of the MultiTech website ([http://www.multitech.net/developer/downloads/](http://www.multitech.net/developer/downloads/)).

Do the following to upgrade the firmware on your device:

-	Before you upgrade your firmware, **save your present configuration as a backup**.
-   Go to the MultiTech website, **locate the firmware upgrade file** you want for your device(MTCDT AEP x.x.x in the current case, where x.x.x is the latest version available on the website), and **download** this file to a known location.
-   Goto **Administration -> Firmware Upgrade**.
-   Click the **Choose Firmware Upgrade File** button:
    -  Click **Browse** to find where the firmware file resides that you want to apply.
    -  Select the file and click **Open**. The file name appears next to the **Choose Firmware Upgrade File** button. Make sure you select the correct BIN file; otherwise, your device can become inoperable.
-   Click on **Start Upgrade**.
-   A message about the time needed to upgrade appears. Click on **OK**.
-   A progress bar appears indicating the status of the upgrade. When the upgrade is completed, your device reboots.
-   After the firmware upgrade is complete, verify your configuration to make sure it is what you expected.

> **Note**: The new firmware is written into flash memory. It may take up to 10 minutes to upgrade the firmware. Do not interrupt the devices' power or press the reset button during this time.
<br>

{{< figure src="024_Firmware_Upgrade.png" alt="Firmware Upgrade" >}}

<br>
