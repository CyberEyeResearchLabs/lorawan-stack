---
title: "Create application"
description: ""
weight: 2
---

## Add an Application to {{%tts%}}

Devices need to be registered with an application to communicate with {{%tts%}}. Let’s add one.

Login to `https://thethings.example.com/`. You will be taken to {{%tts%}} console.

{{< figure src="001_adding_application.png" alt="{{%tts%}} console home page" >}}

Click on **Go to Applications**, and then click on **+ Add Application** to reach the application registration page.

{{< figure src="002_adding_application.png" alt="Applications list page" >}}

Fill and set the required fields:

- For `Application ID`, choose a unique ID by following the guidelines below:
  - Lower case, alphanumeric characters and nonconsecutive `-`
  - Maximum 100 characters are allowed
- For `Application Name`, give any suitable name to the application
- For `Description`, add the desired description of the application
- For `Linking`, Leave the checkbox checked to link the application automatically to the network server,

Click on **Create Application** to finish.

{{< figure src="003_adding_application.png" alt="Application registration page" >}}

Once the application is created, you can find it in the Console under the **Applications** tab.

{{< figure src="004_adding_application.png" alt="Application overview page" >}}