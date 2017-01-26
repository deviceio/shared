# Abstract

The Command and Control Streaming Protocol (CCSP) is an application-level protocol intended to provide real-time control of distributed devices through lightweight remote object access mechanisms and standardized HTTP APIs. The CCSP is intended to be used over the Internet, private Local Area Networks (LANs), or Corporate Networks.

CCSP is targeting implimentors in disciplines such as: Internet of Things (IOT), Configuration Management (CM), Computer Automation (CA), and other disciplines where direct control and coordination of devices and streaming data is required.

# Introduction

CCSP aims to solve the issue of remote device control, reliable device data streaming, and strong security controls by leveraging existing functionality from protocols such as TCP, TLS, and HTTP, By specifying a formalized application-level protocol to exhange commands with remote devices, send or recieve whole or stream data to/from a device, and enforce connectivity and access security.

Existing protocols in this space have failed to express a common application-level protocol that promotes cross-platform command and control and data exchange between an owner of a device, and the device itself or between two or more devices. Instead, existing protocols have soley focused on transport level connectivity and data transmission protocols, with limited development of message passing symantics or security controls leading the industry to fragmentation on how to impliment owner-to-device or device-to-device command and control in a efficient and forward compatable way.

# Connectivity Models

CCSP starts with two connectivity models: D2H (Device to Hub) or D2D (Device to Device). These connectivity models express distribution and connectivity of devices under operation of the CCSP:

D2H: Device to Hub

```
            +----------------------+
            |                      |
            |                      |
            |         HUB          |
            |                      |
       +---->                      <----+
       |    +-----------^----------+    |
       |                |               |
       |                |               |
       |                |               |
       |                |               |
+------------+   +------------+  +------------+
|            |   |            |  |            |
|   DEVICE   |   |   DEVICE   |  |   DEVICE   |
|            |   |            |  |            |
+------------+   +------------+  +------------+
```

The above diagram depics one or more devices that connect to a central Hub component. In this model, the Hub component is under explicit control of an owner or organization and permits access to Device Resources through an exposed HTTP API, which is only made available through the Hub component. The API exposed by the Hub component is described by this protocol to provide direct command and control and data streaming of remote Device Resources. 

D2D: Device to Device

```
+------------+                      +------------+
|            +---------------------->            |
|   DEVICE   |                      |   DEVICE   |
|            <----------------------+            |
+------------+                      +------------+
```

The above diagram depics two or more devices that can inter-connect with each other to form a direct Device to Device communication scheme. In this model, the connection and message passing symantics of a D2D scenario are the same as in the D2H model as described by the CCSP, however CCSP does not describe any required API for Device interaction for use by an owner or organization. Instead, the implimentor of the CCSP protocol for a Device may expose a vendor specific API for command and control. This mode permits Devices implimenting CCSP from different vendors to interact with each other over a common protocol with resources advertised and discovered between the two Devices.

As described further in the CCSP, an optional Device Behavior Plugin system is described for implimentors allowing user-customization of Device Behaviors in the D2D scenario.
As described further in the CCSP, an optional Device Resource Plugin system is described for implimentors allowing user-customization of Device Resources in D2D and D2H scenarios.
As described further in the CCSP, Devices can participate in either the D2H or D2D connectivity models, or both simultaniously.
As described further in the CCSP, all implimentors of CCSP MUST impliment the D2H connectivity model, but MAY impliment the D2D connectivity model.
As described further in the CCSP, Autodiscovery and Authorization mechanisms exist in both the D2H and D2D scenarios.

# Resources

In the context of CCSP, a Resource is the principle abstraction that permits command and control and data streaming to/from a remote distributed Device. Resources are modeled based on traditional Object Oriented Programming (OOP) concepts and permit many of the data access and state manipulation benifits present in OOP implimentations. In the CCSP, there exists 3 categories of Resources:

* Protocol Defined Resources
* Vendor Defined Resources
* User Defined Resources

Protocol Defined Resources are Resources that are described by the CCSP and MAY be implimented by a implimentor of the CCSP. These Resources are designed to be abstract and as cross-platform as possible, providing a rich base of command and control capabilities to owners and organizations looking to orchestrate devices without platform specific peculiarities or vender specific implimentation. As stated previously, Protocol Defined Resources are optional, and implimentors of the CCSP may choose which resources, if any, to impliment that appropriatly match the capabilities of the Device they are targeting. 

Vendor Defined Resources are Resources that are specifically designed and implimented by an implimentor of the CCSP. These Resources are designed to expose functionality not present in Protocol Defined Resources such as: Platform specific functionality, or Vendor specific functionality.

User Defined Resources are Resources that are designed and implimented by a user or organization for a CCSP specific implimentation that supports the Device Resource Plugin system. These Resources are designed to meet owner or organizational needs that cannot be met by the CCSP or vendors.

Resources are implimented on the Device, and as described later in the CCSP are advertised to a connecting peer in both the D2H and D2D models, making the Resource portion of the CCSP self-describing.

# Resource Structure

Resources are structured conceputally after OOP concepts:

* A Resource is a representation of a data structure, existing on the remote Device.
* A Resource MUST impliment one or more Fields or Methods
* Resources can be created (Instance Resources) or can be called directly (Static Resources)
* 

# Resource interaction

Interaction with Resources is conducted through either the Hub API component in the D2H model, or by a peer Device in the D2D model. 




# Concepts

## Devices

In the context of CCSP, a Device is any computer system or electronic device that is capabile of running modern Network, Crytography, and Messanging protocols suitable for Internet or LAN communications. 

# Components

# Transport

# Messages