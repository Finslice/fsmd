/* Generated SBE (Simple Binary Encoding) message codec */

#ifndef _FINSLICE_MD_ORDERCHANGE_H_
#define _FINSLICE_MD_ORDERCHANGE_H_

#include <errno.h>
#if !defined(__STDC_LIMIT_MACROS)
#define __STDC_LIMIT_MACROS 1
#endif
#include <limits.h>
#define SBE_FLOAT_NAN NAN
#define SBE_DOUBLE_NAN NAN
#include <math.h>
#include <stdbool.h>
#include <stdint.h>
#include <string.h>

#include "messageHeader.h"
#include "orderActionType.h"
#include "sideType.h"

#ifdef __cplusplus
#define SBE_ONE_DEF inline
#else
#define SBE_ONE_DEF static inline
#endif

/*
 * Define some byte ordering macros
 */
#if defined(WIN32) || defined(_WIN32)
    #define SBE_BIG_ENDIAN_ENCODE_16(v) _byteswap_ushort(v)
    #define SBE_BIG_ENDIAN_ENCODE_32(v) _byteswap_ulong(v)
    #define SBE_BIG_ENDIAN_ENCODE_64(v) _byteswap_uint64(v)
    #define SBE_LITTLE_ENDIAN_ENCODE_16(v) (v)
    #define SBE_LITTLE_ENDIAN_ENCODE_32(v) (v)
    #define SBE_LITTLE_ENDIAN_ENCODE_64(v) (v)
#elif __BYTE_ORDER__ == __ORDER_LITTLE_ENDIAN__
    #define SBE_BIG_ENDIAN_ENCODE_16(v) __builtin_bswap16(v)
    #define SBE_BIG_ENDIAN_ENCODE_32(v) __builtin_bswap32(v)
    #define SBE_BIG_ENDIAN_ENCODE_64(v) __builtin_bswap64(v)
    #define SBE_LITTLE_ENDIAN_ENCODE_16(v) (v)
    #define SBE_BIG_ENDIAN_ENCODE_32(v) __builtin_bswap32(v)
    #define SBE_BIG_ENDIAN_ENCODE_64(v) __builtin_bswap64(v)
    #define SBE_LITTLE_ENDIAN_ENCODE_16(v) (v)
    #define SBE_LITTLE_ENDIAN_ENCODE_32(v) (v)
    #define SBE_LITTLE_ENDIAN_ENCODE_64(v) (v)
#elif __BYTE_ORDER__ == __ORDER_BIG_ENDIAN__
    #define SBE_LITTLE_ENDIAN_ENCODE_16(v) __builtin_bswap16(v)
    #define SBE_LITTLE_ENDIAN_ENCODE_32(v) __builtin_bswap32(v)
    #define SBE_LITTLE_ENDIAN_ENCODE_64(v) __builtin_bswap64(v)
    #define SBE_BIG_ENDIAN_ENCODE_16(v) (v)
    #define SBE_BIG_ENDIAN_ENCODE_32(v) (v)
    #define SBE_BIG_ENDIAN_ENCODE_64(v) (v)
#else
    #error "Byte Ordering of platform not determined. Set __BYTE_ORDER__ manually before including this file."
#endif

#if defined(SBE_NO_BOUNDS_CHECK)
    #define SBE_BOUNDS_CHECK_EXPECT(exp,c) (false)
#elif defined(_MSC_VER)
    #define SBE_BOUNDS_CHECK_EXPECT(exp,c) (exp)
#else
    #define SBE_BOUNDS_CHECK_EXPECT(exp,c) (__builtin_expect(exp,c))
#endif

#define SBE_NULLVALUE_INT8 INT8_MIN
#define SBE_NULLVALUE_INT16 INT16_MIN
#define SBE_NULLVALUE_INT32 INT32_MIN
#define SBE_NULLVALUE_INT64 INT64_MIN
#define SBE_NULLVALUE_UINT8 UINT8_MAX
#define SBE_NULLVALUE_UINT16 UINT16_MAX
#define SBE_NULLVALUE_UINT32 UINT32_MAX
#define SBE_NULLVALUE_UINT64 UINT64_MAX

#define E100 -50100 // E_BUF_SHORT
#define E103 -50103 // VAL_UNKNOWN_ENUM
#define E104 -50104 // I_OUT_RANGE_NUM
#define E105 -50105 // I_OUT_RANGE_NUM
#define E106 -50106 // I_OUT_RANGE_NUM
#define E107 -50107 // BUF_SHORT_FLYWEIGHT
#define E108 -50108 // BUF_SHORT_NXT_GRP_IND
#define E109 -50109 // STR_TOO_LONG_FOR_LEN_TYP
#define E110 -50110 // CNT_OUT_RANGE

#ifndef SBE_STRERROR_DEFINED
#define SBE_STRERROR_DEFINED
SBE_ONE_DEF const char *sbe_strerror(const int errnum)
{
    switch (errnum)
    {
        case E100:
            return "buffer too short";
        case E103:
            return "unknown value for enum";
        case E104:
            return "index out of range";
        case E105:
            return "index out of range";
        case E106:
            return "length too large";
        case E107:
            return "buffer too short for flyweight";
        case E108:
            return "buffer too short to support next group index";
        case E109:
            return "std::string too long for length type";
        case E110:
            return "count outside of allowed range";
        default:
            return "unknown error";
    }
}
#endif

struct finslice_md_orderChange
{
    char *buffer;
    uint64_t buffer_length;
    uint64_t offset;
    uint64_t position;
    uint64_t acting_block_length;
    uint64_t acting_version;
};

enum finslice_md_orderChange_meta_attribute
{
    finslice_md_orderChange_meta_attribute_EPOCH,
    finslice_md_orderChange_meta_attribute_TIME_UNIT,
    finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE,
    finslice_md_orderChange_meta_attribute_PRESENCE
};

union finslice_md_orderChange_float_as_uint
{
    float fp_value;
    uint32_t uint_value;
};

union finslice_md_orderChange_double_as_uint
{
    double fp_value;
    uint64_t uint_value;
};

struct finslice_md_orderChange_string_view
{
    const char* data;
    size_t length;
};

SBE_ONE_DEF uint64_t finslice_md_orderChange_sbe_position(
    const struct finslice_md_orderChange *const codec)
{
    return codec->position;
}

SBE_ONE_DEF bool finslice_md_orderChange_set_sbe_position(
    struct finslice_md_orderChange *const codec,
    const uint64_t position)
{
    if (SBE_BOUNDS_CHECK_EXPECT((position > codec->buffer_length), false))
    {
        errno = E100;
        return false;
    }
    codec->position = position;

    return true;
}

SBE_ONE_DEF uint64_t *finslice_md_orderChange_sbe_position_ptr(
    struct finslice_md_orderChange *const codec)
{
    return &codec->position;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_reset(
    struct finslice_md_orderChange *const codec,
    char *buffer,
    const uint64_t offset,
    const uint64_t buffer_length,
    const uint64_t acting_block_length,
    const uint64_t acting_version)
{
    codec->buffer = buffer;
    codec->offset = offset;
    codec->buffer_length = buffer_length;
    codec->acting_block_length = acting_block_length;
    codec->acting_version = acting_version;
    if (!finslice_md_orderChange_set_sbe_position(codec, offset + acting_block_length))
    {
        return NULL;
    }

    return codec;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_copy(
    struct finslice_md_orderChange *const codec,
    const struct finslice_md_orderChange *const other)
{
     codec->buffer = other->buffer;
     codec->offset = other->offset;
     codec->buffer_length = other->buffer_length;
     codec->acting_block_length = other->acting_block_length;
     codec->acting_version = other->acting_version;
     codec->position = other->position;

     return codec;
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_sbe_block_length(void)
{
    return (uint16_t)84;
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_sbe_template_id(void)
{
    return (uint16_t)50;
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_sbe_schema_id(void)
{
    return (uint16_t)1;
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_sbe_schema_version(void)
{
    return (uint16_t)1;
}

SBE_ONE_DEF const char *finslice_md_orderChange_sbe_semantic_type(void)
{
    return "X";
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_offset(
    const struct finslice_md_orderChange *const codec)
{
    return codec->offset;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_wrap_and_apply_header(
    struct finslice_md_orderChange *const codec,
    char *buffer,
    const uint64_t offset,
    const uint64_t buffer_length,
    struct finslice_md_messageHeader *const hdr)
{
    finslice_md_messageHeader_wrap(
        hdr, buffer + offset, 0, finslice_md_messageHeader_sbe_schema_version(), buffer_length);

    finslice_md_messageHeader_set_blockLength(hdr, finslice_md_orderChange_sbe_block_length());
    finslice_md_messageHeader_set_templateId(hdr, finslice_md_orderChange_sbe_template_id());
    finslice_md_messageHeader_set_schemaId(hdr, finslice_md_orderChange_sbe_schema_id());
    finslice_md_messageHeader_set_version(hdr, finslice_md_orderChange_sbe_schema_version());

    finslice_md_orderChange_reset(
        codec,
        buffer + offset + finslice_md_messageHeader_encoded_length(),
        0,
        buffer_length - finslice_md_messageHeader_encoded_length(),
        finslice_md_orderChange_sbe_block_length(),
        finslice_md_orderChange_sbe_schema_version());

    return codec;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_wrap_for_encode(
    struct finslice_md_orderChange *const codec,
    char *buffer,
    const uint64_t offset,
    const uint64_t buffer_length)
{
    return finslice_md_orderChange_reset(
        codec,
        buffer,
        offset,
        buffer_length,
        finslice_md_orderChange_sbe_block_length(),
        finslice_md_orderChange_sbe_schema_version());
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_wrap_for_decode(
    struct finslice_md_orderChange *const codec,
    char *buffer,
    const uint64_t offset,
    const uint64_t acting_block_length,
    const uint64_t acting_version,
    const uint64_t buffer_length)
{
    return finslice_md_orderChange_reset(
        codec,
        buffer,
        offset,
        buffer_length,
        acting_block_length,
        acting_version);
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_sbe_rewind(
    struct finslice_md_orderChange *const codec)
{
    return finslice_md_orderChange_wrap_for_decode(
        codec,
        codec->buffer,
        codec->offset,
        codec->acting_block_length,
        codec->acting_version,
        codec->buffer_length);
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_encoded_length(
    const struct finslice_md_orderChange *const codec)
{
    return finslice_md_orderChange_sbe_position(codec) - codec->offset;
}

SBE_ONE_DEF const char *finslice_md_orderChange_buffer(
    const struct finslice_md_orderChange *const codec)
{
    return codec->buffer;
}

SBE_ONE_DEF char *finslice_md_orderChange_mut_buffer(
    struct finslice_md_orderChange *const codec)
{
    return codec->buffer;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_buffer_length(
    const struct finslice_md_orderChange *const codec)
{
    return codec->buffer_length;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_acting_version(
    const struct finslice_md_orderChange *const codec)
{
    return codec->acting_version;
}

SBE_ONE_DEF const char *finslice_md_orderChange_TransactTime_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "UTCTimestamp";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_transactTime_id(void)
{
    return 60;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_transactTime_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_transactTime_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_transactTime_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_transactTime_encoding_offset(void)
{
    return 0;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_transactTime_null_value(void)
{
    return SBE_NULLVALUE_UINT64;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_transactTime_min_value(void)
{
    return UINT64_C(0x0);
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_transactTime_max_value(void)
{
    return UINT64_C(0xfffffffffffffffe);
}

SBE_ONE_DEF size_t finslice_md_orderChange_transactTime_encoding_length(void)
{
    return 8;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_transactTime(
    const struct finslice_md_orderChange *const codec)
{
    uint64_t val;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(&val, codec->buffer + codec->offset + 0, sizeof(uint64_t));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return SBE_LITTLE_ENDIAN_ENCODE_64(val);
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_transactTime(
    struct finslice_md_orderChange *const codec,
    const uint64_t value)
{
    uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(codec->buffer + codec->offset + 0, &val, sizeof(uint64_t));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_SecurityID_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "int";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_securityID_id(void)
{
    return 48;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_securityID_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_securityID_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_securityID_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_securityID_encoding_offset(void)
{
    return 8;
}

SBE_ONE_DEF uint32_t finslice_md_orderChange_securityID_null_value(void)
{
    return SBE_NULLVALUE_UINT32;
}

SBE_ONE_DEF uint32_t finslice_md_orderChange_securityID_min_value(void)
{
    return UINT32_C(0x0);
}

SBE_ONE_DEF uint32_t finslice_md_orderChange_securityID_max_value(void)
{
    return UINT32_C(0xfffffffe);
}

SBE_ONE_DEF size_t finslice_md_orderChange_securityID_encoding_length(void)
{
    return 4;
}

SBE_ONE_DEF uint32_t finslice_md_orderChange_securityID(
    const struct finslice_md_orderChange *const codec)
{
    uint32_t val;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(&val, codec->buffer + codec->offset + 8, sizeof(uint32_t));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return SBE_LITTLE_ENDIAN_ENCODE_32(val);
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_securityID(
    struct finslice_md_orderChange *const codec,
    const uint32_t value)
{
    uint32_t val = SBE_LITTLE_ENDIAN_ENCODE_32(value);
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(codec->buffer + codec->offset + 8, &val, sizeof(uint32_t));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_OrderAction_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "int";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_orderAction_id(void)
{
    return 279;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_orderAction_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_orderAction_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_orderAction_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_orderAction_encoding_offset(void)
{
    return 12;
}

SBE_ONE_DEF size_t finslice_md_orderChange_orderAction_encoding_length(void)
{
    return 1;
}

SBE_ONE_DEF bool finslice_md_orderChange_orderAction(
    const struct finslice_md_orderChange *const codec,
    enum finslice_md_orderActionType *const out)
{
    uint8_t val;
    memcpy(&val, codec->buffer + codec->offset + 12, sizeof(uint8_t));

    return finslice_md_orderActionType_get((val), out);
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_orderAction(
    struct finslice_md_orderChange *const codec,
    const enum finslice_md_orderActionType value)
{
    uint8_t val = (value);
    memcpy(codec->buffer + codec->offset + 12, &val, sizeof(uint8_t));

    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_Side_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "int";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_side_id(void)
{
    return 54;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_side_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_side_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_side_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_side_encoding_offset(void)
{
    return 13;
}

SBE_ONE_DEF size_t finslice_md_orderChange_side_encoding_length(void)
{
    return 1;
}

SBE_ONE_DEF bool finslice_md_orderChange_side(
    const struct finslice_md_orderChange *const codec,
    enum finslice_md_sideType *const out)
{
    uint8_t val;
    memcpy(&val, codec->buffer + codec->offset + 13, sizeof(uint8_t));

    return finslice_md_sideType_get((val), out);
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_side(
    struct finslice_md_orderChange *const codec,
    const enum finslice_md_sideType value)
{
    uint8_t val = (value);
    memcpy(codec->buffer + codec->offset + 13, &val, sizeof(uint8_t));

    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_Price_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "double";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_price_id(void)
{
    return 270;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_price_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_price_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_price_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_price_encoding_offset(void)
{
    return 16;
}

SBE_ONE_DEF double finslice_md_orderChange_price_null_value(void)
{
    return SBE_DOUBLE_NAN;
}

SBE_ONE_DEF double finslice_md_orderChange_price_min_value(void)
{
    return 4.9E-324;
}

SBE_ONE_DEF double finslice_md_orderChange_price_max_value(void)
{
    return 1.7976931348623157E308;
}

SBE_ONE_DEF size_t finslice_md_orderChange_price_encoding_length(void)
{
    return 8;
}

SBE_ONE_DEF double finslice_md_orderChange_price(
    const struct finslice_md_orderChange *const codec)
{
    union finslice_md_orderChange_double_as_uint val;
    memcpy(&val, codec->buffer + codec->offset + 16, sizeof(double));
    val.uint_value = SBE_LITTLE_ENDIAN_ENCODE_64(val.uint_value);
    return val.fp_value;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_price(
    struct finslice_md_orderChange *const codec,
    const double value)
{
    union finslice_md_orderChange_double_as_uint val;
    val.fp_value = value;
    val.uint_value = SBE_LITTLE_ENDIAN_ENCODE_64(val.uint_value);
    memcpy(codec->buffer + codec->offset + 16, &val, sizeof(double));
    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_NewSize_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "double";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_newSize_id(void)
{
    return 271;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_newSize_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_newSize_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_newSize_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_newSize_encoding_offset(void)
{
    return 24;
}

SBE_ONE_DEF double finslice_md_orderChange_newSize_null_value(void)
{
    return SBE_DOUBLE_NAN;
}

SBE_ONE_DEF double finslice_md_orderChange_newSize_min_value(void)
{
    return 4.9E-324;
}

SBE_ONE_DEF double finslice_md_orderChange_newSize_max_value(void)
{
    return 1.7976931348623157E308;
}

SBE_ONE_DEF size_t finslice_md_orderChange_newSize_encoding_length(void)
{
    return 8;
}

SBE_ONE_DEF double finslice_md_orderChange_newSize(
    const struct finslice_md_orderChange *const codec)
{
    union finslice_md_orderChange_double_as_uint val;
    memcpy(&val, codec->buffer + codec->offset + 24, sizeof(double));
    val.uint_value = SBE_LITTLE_ENDIAN_ENCODE_64(val.uint_value);
    return val.fp_value;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_newSize(
    struct finslice_md_orderChange *const codec,
    const double value)
{
    union finslice_md_orderChange_double_as_uint val;
    val.fp_value = value;
    val.uint_value = SBE_LITTLE_ENDIAN_ENCODE_64(val.uint_value);
    memcpy(codec->buffer + codec->offset + 24, &val, sizeof(double));
    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_SecuritySequence_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "int";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_securitySequence_id(void)
{
    return 83;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_securitySequence_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_securitySequence_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_securitySequence_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_securitySequence_encoding_offset(void)
{
    return 32;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_securitySequence_null_value(void)
{
    return SBE_NULLVALUE_UINT64;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_securitySequence_min_value(void)
{
    return UINT64_C(0x0);
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_securitySequence_max_value(void)
{
    return UINT64_C(0xfffffffffffffffe);
}

SBE_ONE_DEF size_t finslice_md_orderChange_securitySequence_encoding_length(void)
{
    return 8;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_securitySequence(
    const struct finslice_md_orderChange *const codec)
{
    uint64_t val;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(&val, codec->buffer + codec->offset + 32, sizeof(uint64_t));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return SBE_LITTLE_ENDIAN_ENCODE_64(val);
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_securitySequence(
    struct finslice_md_orderChange *const codec,
    const uint64_t value)
{
    uint64_t val = SBE_LITTLE_ENDIAN_ENCODE_64(value);
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(codec->buffer + codec->offset + 32, &val, sizeof(uint64_t));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_OrderID_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "String";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_orderID_id(void)
{
    return 37;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_orderID_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_orderID_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_orderID_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_orderID_encoding_offset(void)
{
    return 40;
}

SBE_ONE_DEF char finslice_md_orderChange_orderID_null_value(void)
{
    return (char)0;
}

SBE_ONE_DEF char finslice_md_orderChange_orderID_min_value(void)
{
    return (char)32;
}

SBE_ONE_DEF char finslice_md_orderChange_orderID_max_value(void)
{
    return (char)126;
}

SBE_ONE_DEF size_t finslice_md_orderChange_orderID_encoding_length(void)
{
    return 36;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_orderID_length(void)
{
    return 36;
}

SBE_ONE_DEF const char *finslice_md_orderChange_orderID_buffer(
    const struct finslice_md_orderChange *const codec)
{
    return codec->buffer + codec->offset + 40;
}

SBE_ONE_DEF char finslice_md_orderChange_orderID_unsafe(
    const struct finslice_md_orderChange *const codec,
    const uint64_t index)
{
    char val;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(&val, codec->buffer + codec->offset + 40 + (index * 1), sizeof(char));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return (val);
}

SBE_ONE_DEF bool finslice_md_orderChange_orderID(
    const struct finslice_md_orderChange *const codec,
    const uint64_t index,
    char *const out)
{
    if (index >= 36)
    {
        errno = E104;
        return false;
    }

    char val;
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(&val, codec->buffer + codec->offset + 40 + (index * 1), sizeof(char));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    *out = (val);
    return true;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_orderID_unsafe(
    struct finslice_md_orderChange *const codec,
    const uint64_t index,
    const char value)
{
    char val = (value);
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(codec->buffer + codec->offset + 40 + (index * 1), &val, sizeof(char));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return codec;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_orderID(
    struct finslice_md_orderChange *const codec,
    const uint64_t index,
    const char value)
{
    if (index >= 36)
    {
        errno = E105;
        return NULL;
    }

    char val = (value);
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmaybe-uninitialized"
#endif
    memcpy(codec->buffer + codec->offset + 40 + (index * 1), &val, sizeof(char));
#if defined(__GNUG__) && !defined(__clang__)
#pragma GCC diagnostic pop
#endif
    return codec;
}

SBE_ONE_DEF char *finslice_md_orderChange_get_orderID(
    const struct finslice_md_orderChange *const codec,
    char *dst,
    const uint64_t length)
{
    if (length > 36)
    {
        errno = E106;
        return NULL;
    }

    memcpy(dst, codec->buffer + codec->offset + 40, sizeof(char) * length);

    return dst;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_put_orderID(
    struct finslice_md_orderChange *const codec,
    const char *src)
{
    memcpy(codec->buffer + codec->offset + 40, src, sizeof(char) * 36);

    return codec;
}

SBE_ONE_DEF const char *finslice_md_orderChange_OldSize_meta_attribute(
    const enum finslice_md_orderChange_meta_attribute attribute)
{
    switch (attribute)
    {
        case finslice_md_orderChange_meta_attribute_EPOCH: return "";
        case finslice_md_orderChange_meta_attribute_TIME_UNIT: return "";
        case finslice_md_orderChange_meta_attribute_SEMANTIC_TYPE: return "double";
        case finslice_md_orderChange_meta_attribute_PRESENCE: return "required";
    }

    return "";
}

SBE_ONE_DEF uint16_t finslice_md_orderChange_oldSize_id(void)
{
    return 273;
}

SBE_ONE_DEF uint64_t finslice_md_orderChange_oldSize_since_version(void)
{
    return 0;
}

SBE_ONE_DEF bool finslice_md_orderChange_oldSize_in_acting_version(
    const struct finslice_md_orderChange *const codec)
{
#if defined(__clang__)
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wtautological-compare"
#endif
    return codec->acting_version >= finslice_md_orderChange_oldSize_since_version();
#if defined(__clang__)
#pragma clang diagnostic pop
#endif
}

SBE_ONE_DEF size_t finslice_md_orderChange_oldSize_encoding_offset(void)
{
    return 76;
}

SBE_ONE_DEF double finslice_md_orderChange_oldSize_null_value(void)
{
    return SBE_DOUBLE_NAN;
}

SBE_ONE_DEF double finslice_md_orderChange_oldSize_min_value(void)
{
    return 4.9E-324;
}

SBE_ONE_DEF double finslice_md_orderChange_oldSize_max_value(void)
{
    return 1.7976931348623157E308;
}

SBE_ONE_DEF size_t finslice_md_orderChange_oldSize_encoding_length(void)
{
    return 8;
}

SBE_ONE_DEF double finslice_md_orderChange_oldSize(
    const struct finslice_md_orderChange *const codec)
{
    union finslice_md_orderChange_double_as_uint val;
    memcpy(&val, codec->buffer + codec->offset + 76, sizeof(double));
    val.uint_value = SBE_LITTLE_ENDIAN_ENCODE_64(val.uint_value);
    return val.fp_value;
}

SBE_ONE_DEF struct finslice_md_orderChange *finslice_md_orderChange_set_oldSize(
    struct finslice_md_orderChange *const codec,
    const double value)
{
    union finslice_md_orderChange_double_as_uint val;
    val.fp_value = value;
    val.uint_value = SBE_LITTLE_ENDIAN_ENCODE_64(val.uint_value);
    memcpy(codec->buffer + codec->offset + 76, &val, sizeof(double));
    return codec;
}

#endif
