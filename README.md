# Finslice Market Date

Low latency, binary encoded market data(MD) feed by Finslice.

We adopted a common, high performance FIX standard called [Simple Binary Encoding](https://www.fixtrading.org/standards/sbe-online/). You can find the schema xml and generated codes for parsing the data in this repository.

## Installation

The generated codes are header only, you can either copy them directly to your project or include from this repo. They are stored by language:
- C++: [cpp](cpp)
- C/C#: [c](c)
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
| Channel Header[12 bytes]  | SBE Header[8 bytes] | Message Body [x bytes]           |     |
|                           +---------------------+----------------------------------+     |
+------------------------------------------------------------------------------------------+
```
The Channel Header contains: {Timestamp(8), ChannelSequence(4)}:

 - Timestamp is Epoch time in nanoseconds
 - Channel Sequence is uint32_t, increase by 1 per packet, wraps at 0xffffffff

 The SBE header is tsandard 8 bytes that contains:
 - Block Length(16 bit): number of bytes including header and body
 - TemplateID (16 bit): Message type
 - SchemaId(16 bit): Finslice MD(1) 
 - Version(16 bit): current version

For different message body type checkout [finslice_md.xml](finslice_md.xml)
## Channels and products mapping
The market data is grouped by UDP multicast channels, each data source(exchange) can have multiple channels.

Multiple products(pairs) can be carried in a single channel, each product will have a globally unique, 32 bit product ID.

The full list of channels and products will be published in [config](config), an example channel and its products looks like this:
```
Channel:
  name: coinbase-1
  a-feed: 204.0.0.1:10000
  b-feed: 204.0.0.101:14000
  products:
    BTC-USD: 10000
    BTC-EUR: 10001
    ETH-USD: 20000
    ETH-EUR: 20001
    SOL-USD: 80000
```
## Types of Feeds
The design of Finslice is to support fast distribution of full market data, so we discarded all un-necessary lookup and convertions. For each exchange(source) we will support their lowest latency feed type, for example: order based feed(L3) from Coinbase and price level based feed(L2) from Kraken. The SBE encoding supports either message types, and clients can use each exchange's native recovery methods to initilaize or re-sync their full book.
