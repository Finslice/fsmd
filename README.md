# Finslice Market Date

The low latency binary Market Data feed by Finslice adopts a high performance FIX standard called [Simple Binary Encoding](https://www.fixtrading.org/standards/sbe-online/). You can find the schema xmal and generated codes for parsing the data in this repository.

## Installation

The generated codes are header only, you can either copy them directly to your project or include from this repo. They are stored by language:
- C++: [cpp](cpp)
- C: [c](c)
- Java: [java](java)
- Golang: [golang](golang)
- Rust: [rust](rust)

A detailed user guide for each generated code can be found in: [SBE Wiki](https://github.com/real-logic/simple-binary-encoding/wiki)


## Demo C++ Client
A simple C++ client is included in this repository. To build, make sure you have a rescent version of g++ and cmake installed in your dev environment. Run:

```
mkdir build
cd build
cmake -DCMAKE_BUILD_TYPE=RELEASE ..
make
```
After build you can simple start the client by typing something like:
```
./finslice_demo eth0 224.0.0.1 19200
```
The program takes 3 arguements:
- network interface to subscribe/listen to UDP Multicast
- UDP multicast group IP/ID
- UDP multicast distination port

It will print parsed market data on screen

## Packet Frame Format
The UDP payload part of each packet is encoded in Little Endian, and looks like:
```
+------------------------------------------------------------------------------------------+
|                           +---------------------+----------------------------------+     |
| Channel Header[12 bytes]  | SBE Header[8 bytes] | Message [x bytes]                |     |
|                           +---------------------+----------------------------------+     |
+------------------------------------------------------------------------------------------+
```
## Channels

## Products mapping


## Types of Feed

