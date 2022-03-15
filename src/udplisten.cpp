#include "udplisten.h"
#include <iostream>
#include <net/if.h>
#include <netdb.h>
#include <sys/socket.h>

UDPListen::UDPListen(const char *iface, const char *addr, uint16_t port) {
  struct addrinfo hints = {};
  struct addrinfo *group;
  int error;
  std::string p = std::to_string(port);
  hints.ai_socktype = SOCK_DGRAM;
  hints.ai_flags = AI_NUMERICHOST;

  error = getaddrinfo(addr, p.c_str(), &hints, &group);
  if (error) {
    throw std::runtime_error("Can't resolve group-id info" + std::string(addr));
  }
  _socket = socket(group->ai_family, group->ai_socktype, group->ai_protocol);
  if (_socket < 0) {
    throw std::runtime_error("Socket Creation Error");
  }

  /* MISC socket attributes: Allow Reuse, 1s timeout, buffer to 64k */
  error = 0;
  int reuse = 1;
  error += setsockopt(_socket, SOL_SOCKET, SO_REUSEADDR, &reuse, sizeof(reuse));

  struct timeval read_timeout;
  read_timeout.tv_sec = 1;
  read_timeout.tv_usec = 0;
  error += setsockopt(_socket, SOL_SOCKET, SO_RCVTIMEO, &read_timeout, sizeof(read_timeout));

  size_t rcvbuf = 1 << 15;
  error += setsockopt(_socket, SOL_SOCKET, SO_RCVBUF, &rcvbuf, sizeof(rcvbuf));

  if (error) {
    throw std::runtime_error("Socket misc options setting error");
  }

  error = bind(_socket, group->ai_addr, group->ai_addrlen);
  if (error) {
    throw std::runtime_error("Socket Bind Error");
  }

  /* Mcast Join */
  struct group_req multicast_request;
  int ip_proto = IPPROTO_IP;
  uint32_t ifidx = if_nametoindex(iface);
  if (ifidx == 0) {
    throw std::runtime_error("Can't resolve interface name " + std::string(iface));
  }
  multicast_request.gr_interface = ifidx;
  memcpy(&multicast_request.gr_group, group->ai_addr, group->ai_addrlen);
  ip_proto = IPPROTO_IP;
  if (group->ai_family == AF_INET6) {
    ip_proto = IPPROTO_IPV6;
  }
  error = setsockopt(_socket, ip_proto, MCAST_JOIN_GROUP, (char *)&multicast_request,
                     sizeof(multicast_request));
  if (error) {
    throw std::runtime_error("Multicast Joining Error");
  }
}

UDPListen::~UDPListen() {
  ::close(_socket);
  _socket = -1;
}

int UDPListen::getMsg(char *buffer, int len, uint32_t &src_ip, uint16_t &src_port) {
  sockaddr_in clntAddr;
  socklen_t addrLen = sizeof(clntAddr);
  int rtn;
  rtn = recvfrom(_socket, (void *)buffer, len, 0, (sockaddr *)&clntAddr, (socklen_t *)&addrLen);
  if (rtn >= 0) {
    src_port = ntohs(clntAddr.sin_port);
    src_ip = ntohl(clntAddr.sin_addr.s_addr);
  }
  return rtn;
}