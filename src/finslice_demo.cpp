#include "finslice_md/finslice_md.h"
#include "udplisten.h"
#include <chrono>
#include <csignal>
#include <iostream>

/* Demo C++ program to parse and decode Finslice multicast market data
   Using plain Linux UDP Multicast socket */

static int done = 0;
void signalHandler(int signum) {
  std::cout << "Interrupt signal (" << signum << ") received.\n";
  done = 1;
}

// Make sure receieve packet buffer > 192 bytes
const int MTU_SIZE = 1500;

uint64_t getTimeStampNS() {
  using namespace std::chrono_literals;
  using clock = std::chrono::system_clock;
  using tpns = std::chrono::time_point<clock, std::chrono::duration<uint64_t, std::nano>>;
  tpns tp = clock::now();
  return tp.time_since_epoch().count();
}

int main(int argc, char **argv) {
  if (argc != 4) {
    std::cerr << "Usage: " << argv[0] << " <Interface> <Multicast Address> <Port>"
              << "\n";
    exit(1);
  }
  signal(SIGINT, signalHandler);
  uint16_t port = atoi(argv[3]);
  UDPListen sock(argv[1], argv[2], port);
  char buf[MTU_SIZE];
  uint32_t src_ip;
  uint16_t src_port;
  uint32_t myseq = 0;
  while (!done) {
    int bytes = sock.getMsg(buf, MTU_SIZE, src_ip, src_port);
    if (bytes > 16) {
      uint64_t ts = getTimeStampNS();
      finslice_md::MDMessage msg(buf);
      if ((myseq != 0) && (myseq != msg.getChannelSequence())) {
        std::cerr << ts << " Gap Detected! Expected sequence: " << myseq
                  << " Received sequence: " << msg.getChannelSequence() << "\n";
      }
      std::cout << ts << " " << msg << "/n";
      myseq = msg.getChannelSequence() + 1;
      // Example of parse & convert to event types
      // auto hdr = msg.getMsgHeader();
      // if (hdr.templateId() == finslice_md::MatchEvent::sbeTemplateId())
      // {
      //   auto event = msg.getEvent<finslice_md::MatchEvent>();
      //   std::cout << event;
      // }
    }
  }
  std::cout << "Program exited normally\n";
  return done;
}