#pragma once
#include "MatchEvent.h"
#include "MessageHeader.h"
#include "OrderActionType.h"
#include "OrderChange.h"
#include "OrderUpdate.h"
#include "SideType.h"
#include <cstdint>
#include <memory>
#include <string>

namespace finslice_md {
class MDMessage {
private:
  char *_buffer = nullptr;

public:
  MDMessage() = default;
  MDMessage(char *buffer) : _buffer(buffer) {}
  uint64_t getChannelTimestamp() {
    uint64_t val;
    std::memcpy(&val, _buffer, sizeof(std::uint64_t));
    return val;
  }
  uint32_t getChannelSequence() {
    uint32_t val;
    std::memcpy(&val, _buffer, sizeof(std::uint32_t));
    return val;
  }
  MessageHeader getMsgHeader() {
    MessageHeader hdr(_buffer + 12, 8);
    return hdr;
  }
  template <typename T> T getEvent() {
    T event(_buffer + 20, 192);
    // auto hdr = getMsgHeader();
    // event.wrapForDecode(_buffer + 12, hdr.encodedLength(), hdr.blockLength(), hdr.version(),
    // 192);
    return event;
  }

  template <typename CharT, typename Traits>
  friend std::basic_ostream<CharT, Traits> &operator<<(std::basic_ostream<CharT, Traits> &builder,
                                                       MDMessage writer) {
    builder << '{';
    builder << R"("ChTimestamp": )";
    builder << +writer.getChannelTimestamp();

    builder << ", ";
    builder << R"("ChSequence": )";
    builder << +writer.getChannelSequence();
    builder << '}';
    auto hdr = writer.getMsgHeader();
    builder << hdr;
    switch (hdr.templateId()) {
    case (MatchEvent::sbeTemplateId()): {
      builder << writer.getEvent<MatchEvent>();
      break;
    }
    case (OrderUpdate::sbeTemplateId()): {
      builder << writer.getEvent<OrderUpdate>();
      break;
    }
    case (OrderChange::sbeTemplateId()): {
      builder << writer.getEvent<OrderChange>();
      break;
    }
    }
    return builder;
  }
};
} // namespace finslice_md
