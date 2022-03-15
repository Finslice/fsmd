/* Generated SBE (Simple Binary Encoding) message codec */
#ifndef _FINSLICE_MD_ORDERACTIONTYPE_H_
#define _FINSLICE_MD_ORDERACTIONTYPE_H_

#if !defined(__STDC_LIMIT_MACROS)
#  define __STDC_LIMIT_MACROS 1
#endif

#include <cstdint>
#include <iomanip>
#include <limits>
#include <ostream>
#include <stdexcept>
#include <sstream>
#include <string>

#define SBE_NULLVALUE_INT8 (std::numeric_limits<std::int8_t>::min)()
#define SBE_NULLVALUE_INT16 (std::numeric_limits<std::int16_t>::min)()
#define SBE_NULLVALUE_INT32 (std::numeric_limits<std::int32_t>::min)()
#define SBE_NULLVALUE_INT64 (std::numeric_limits<std::int64_t>::min)()
#define SBE_NULLVALUE_UINT8 (std::numeric_limits<std::uint8_t>::max)()
#define SBE_NULLVALUE_UINT16 (std::numeric_limits<std::uint16_t>::max)()
#define SBE_NULLVALUE_UINT32 (std::numeric_limits<std::uint32_t>::max)()
#define SBE_NULLVALUE_UINT64 (std::numeric_limits<std::uint64_t>::max)()

namespace finslice_md {

class OrderActionType
{
public:
    enum Value
    {
        NewMarketOrder = static_cast<std::uint8_t>(1),
        NewLimitOrder = static_cast<std::uint8_t>(2),
        OrderOpen = static_cast<std::uint8_t>(3),
        OrderFilled = static_cast<std::uint8_t>(8),
        OrderCancel = static_cast<std::uint8_t>(9),
        MatchEvent = static_cast<std::uint8_t>(12),
        OrderChange = static_cast<std::uint8_t>(13),
        NULL_VALUE = static_cast<std::uint8_t>(255)
    };

    static OrderActionType::Value get(const std::uint8_t value)
    {
        switch (value)
        {
            case static_cast<std::uint8_t>(1): return NewMarketOrder;
            case static_cast<std::uint8_t>(2): return NewLimitOrder;
            case static_cast<std::uint8_t>(3): return OrderOpen;
            case static_cast<std::uint8_t>(8): return OrderFilled;
            case static_cast<std::uint8_t>(9): return OrderCancel;
            case static_cast<std::uint8_t>(12): return MatchEvent;
            case static_cast<std::uint8_t>(13): return OrderChange;
            case static_cast<std::uint8_t>(255): return NULL_VALUE;
        }

        throw std::runtime_error("unknown value for enum OrderActionType [E103]");
    }

    static const char *c_str(const OrderActionType::Value value)
    {
        switch (value)
        {
            case NewMarketOrder: return "NewMarketOrder";
            case NewLimitOrder: return "NewLimitOrder";
            case OrderOpen: return "OrderOpen";
            case OrderFilled: return "OrderFilled";
            case OrderCancel: return "OrderCancel";
            case MatchEvent: return "MatchEvent";
            case OrderChange: return "OrderChange";
            case NULL_VALUE: return "NULL_VALUE";
        }

        throw std::runtime_error("unknown value for enum OrderActionType [E103]:");
    }

    template<typename CharT, typename Traits>
    friend std::basic_ostream<CharT, Traits> & operator << (
        std::basic_ostream<CharT, Traits> &os, OrderActionType::Value m)
    {
        return os << OrderActionType::c_str(m);
    }
};

}

#endif
