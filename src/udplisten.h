#pragma once
#include <arpa/inet.h>
#include <cstring>
#include <stdexcept>
#include <unistd.h>

class UDPListen {
public:
  UDPListen(const char *iface, const char *addr, uint16_t port);
  ~UDPListen();
  int getMsg(char *buffer, int len, uint32_t &src_ip, uint16_t &src_port);

private:
  int _socket;
};