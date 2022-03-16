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

  static constexpr size_t channelHdrLength() noexcept { return 12; }

  uint64_t getChannelTimestamp() {
    uint64_t val;
    std::memcpy(&val, _buffer, sizeof(std::uint64_t));
    return val;
  }

  void setChannelTimestamp(const uint64_t value) { std::memcpy(_buffer, &value, sizeof(uint64_t)); }

  uint32_t getChannelSequence() {
    uint32_t val;
    std::memcpy(&val, _buffer + 8, sizeof(std::uint32_t));
    return val;
  }

  void setChannelSequence(const uint32_t value) {
    std::memcpy(_buffer + 8, &value, sizeof(uint32_t));
  }

  MessageHeader getMsgHeader() {
    MessageHeader hdr(_buffer + channelHdrLength(), 8);
    return hdr;
  }

  template <typename T> T getEvent() {
    T event(_buffer + channelHdrLength() + MessageHeader::encodedLength(), 172);
    // auto hdr = getMsgHeader();
    // event.wrapForDecode(_buffer + 12, hdr.encodedLength(), hdr.blockLength(), hdr.version(),
    // 192);
    return event;
  }

  char *getSBEBuf() { return _buffer + 12; }

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

struct MDBuffer {
  size_t len;
  char buffer[192];
  void clean() { std::memset(buffer, 0, len); }
  MDMessage getMD() {
    MDMessage md(buffer);
    return md;
  }
  template <typename T> T getEvent() {
    T event(buffer + MDMessage::channelHdrLength() + MessageHeader::encodedLength(), 172);
    return event;
  }
};

using MDBufferPtr = std::unique_ptr<MDBuffer>;

} // namespace finslice_md
