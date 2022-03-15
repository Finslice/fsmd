/* Generated SBE (Simple Binary Encoding) message codec. */
package finslice_md;

import org.agrona.MutableDirectBuffer;


/**
 * Order Matched
 */
@SuppressWarnings("all")
public final class MatchEventEncoder
{
    public static final int BLOCK_LENGTH = 120;
    public static final int TEMPLATE_ID = 42;
    public static final int SCHEMA_ID = 1;
    public static final int SCHEMA_VERSION = 1;
    public static final java.nio.ByteOrder BYTE_ORDER = java.nio.ByteOrder.LITTLE_ENDIAN;

    private final MatchEventEncoder parentMessage = this;
    private MutableDirectBuffer buffer;
    private int initialOffset;
    private int offset;
    private int limit;

    public int sbeBlockLength()
    {
        return BLOCK_LENGTH;
    }

    public int sbeTemplateId()
    {
        return TEMPLATE_ID;
    }

    public int sbeSchemaId()
    {
        return SCHEMA_ID;
    }

    public int sbeSchemaVersion()
    {
        return SCHEMA_VERSION;
    }

    public String sbeSemanticType()
    {
        return "X";
    }

    public MutableDirectBuffer buffer()
    {
        return buffer;
    }

    public int initialOffset()
    {
        return initialOffset;
    }

    public int offset()
    {
        return offset;
    }

    public MatchEventEncoder wrap(final MutableDirectBuffer buffer, final int offset)
    {
        if (buffer != this.buffer)
        {
            this.buffer = buffer;
        }
        this.initialOffset = offset;
        this.offset = offset;
        limit(offset + BLOCK_LENGTH);

        return this;
    }

    public MatchEventEncoder wrapAndApplyHeader(
        final MutableDirectBuffer buffer, final int offset, final MessageHeaderEncoder headerEncoder)
    {
        headerEncoder
            .wrap(buffer, offset)
            .blockLength(BLOCK_LENGTH)
            .templateId(TEMPLATE_ID)
            .schemaId(SCHEMA_ID)
            .version(SCHEMA_VERSION);

        return wrap(buffer, offset + MessageHeaderEncoder.ENCODED_LENGTH);
    }

    public int encodedLength()
    {
        return limit - offset;
    }

    public int limit()
    {
        return limit;
    }

    public void limit(final int limit)
    {
        this.limit = limit;
    }

    public static int transactTimeId()
    {
        return 60;
    }

    public static int transactTimeSinceVersion()
    {
        return 0;
    }

    public static int transactTimeEncodingOffset()
    {
        return 0;
    }

    public static int transactTimeEncodingLength()
    {
        return 8;
    }

    public static String transactTimeMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "UTCTimestamp";
        }

        return "";
    }

    public static long transactTimeNullValue()
    {
        return 0xffffffffffffffffL;
    }

    public static long transactTimeMinValue()
    {
        return 0x0L;
    }

    public static long transactTimeMaxValue()
    {
        return 0xfffffffffffffffeL;
    }

    public MatchEventEncoder transactTime(final long value)
    {
        buffer.putLong(offset + 0, value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int securityIDId()
    {
        return 48;
    }

    public static int securityIDSinceVersion()
    {
        return 0;
    }

    public static int securityIDEncodingOffset()
    {
        return 8;
    }

    public static int securityIDEncodingLength()
    {
        return 4;
    }

    public static String securityIDMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "int";
        }

        return "";
    }

    public static long securityIDNullValue()
    {
        return 4294967295L;
    }

    public static long securityIDMinValue()
    {
        return 0L;
    }

    public static long securityIDMaxValue()
    {
        return 4294967294L;
    }

    public MatchEventEncoder securityID(final long value)
    {
        buffer.putInt(offset + 8, (int)value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int orderActionId()
    {
        return 279;
    }

    public static int orderActionSinceVersion()
    {
        return 0;
    }

    public static int orderActionEncodingOffset()
    {
        return 12;
    }

    public static int orderActionEncodingLength()
    {
        return 1;
    }

    public static String orderActionMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "int";
        }

        return "";
    }

    public MatchEventEncoder orderAction(final OrderActionType value)
    {
        buffer.putByte(offset + 12, (byte)value.value());
        return this;
    }

    public static int makerSideId()
    {
        return 54;
    }

    public static int makerSideSinceVersion()
    {
        return 0;
    }

    public static int makerSideEncodingOffset()
    {
        return 13;
    }

    public static int makerSideEncodingLength()
    {
        return 1;
    }

    public static String makerSideMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "int";
        }

        return "";
    }

    public MatchEventEncoder makerSide(final SideType value)
    {
        buffer.putByte(offset + 13, (byte)value.value());
        return this;
    }

    public static int priceId()
    {
        return 270;
    }

    public static int priceSinceVersion()
    {
        return 0;
    }

    public static int priceEncodingOffset()
    {
        return 16;
    }

    public static int priceEncodingLength()
    {
        return 8;
    }

    public static String priceMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "double";
        }

        return "";
    }

    public static double priceNullValue()
    {
        return Double.NaN;
    }

    public static double priceMinValue()
    {
        return 4.9E-324d;
    }

    public static double priceMaxValue()
    {
        return 1.7976931348623157E308d;
    }

    public MatchEventEncoder price(final double value)
    {
        buffer.putDouble(offset + 16, value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int sizeId()
    {
        return 271;
    }

    public static int sizeSinceVersion()
    {
        return 0;
    }

    public static int sizeEncodingOffset()
    {
        return 24;
    }

    public static int sizeEncodingLength()
    {
        return 8;
    }

    public static String sizeMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "double";
        }

        return "";
    }

    public static double sizeNullValue()
    {
        return Double.NaN;
    }

    public static double sizeMinValue()
    {
        return 4.9E-324d;
    }

    public static double sizeMaxValue()
    {
        return 1.7976931348623157E308d;
    }

    public MatchEventEncoder size(final double value)
    {
        buffer.putDouble(offset + 24, value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int securitySequenceId()
    {
        return 83;
    }

    public static int securitySequenceSinceVersion()
    {
        return 0;
    }

    public static int securitySequenceEncodingOffset()
    {
        return 32;
    }

    public static int securitySequenceEncodingLength()
    {
        return 8;
    }

    public static String securitySequenceMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "int";
        }

        return "";
    }

    public static long securitySequenceNullValue()
    {
        return 0xffffffffffffffffL;
    }

    public static long securitySequenceMinValue()
    {
        return 0x0L;
    }

    public static long securitySequenceMaxValue()
    {
        return 0xfffffffffffffffeL;
    }

    public MatchEventEncoder securitySequence(final long value)
    {
        buffer.putLong(offset + 32, value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int makerOrderIDId()
    {
        return 37;
    }

    public static int makerOrderIDSinceVersion()
    {
        return 0;
    }

    public static int makerOrderIDEncodingOffset()
    {
        return 40;
    }

    public static int makerOrderIDEncodingLength()
    {
        return 36;
    }

    public static String makerOrderIDMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "String";
        }

        return "";
    }

    public static byte makerOrderIDNullValue()
    {
        return (byte)0;
    }

    public static byte makerOrderIDMinValue()
    {
        return (byte)32;
    }

    public static byte makerOrderIDMaxValue()
    {
        return (byte)126;
    }

    public static int makerOrderIDLength()
    {
        return 36;
    }


    public MatchEventEncoder makerOrderID(final int index, final byte value)
    {
        if (index < 0 || index >= 36)
        {
            throw new IndexOutOfBoundsException("index out of range: index=" + index);
        }

        final int pos = offset + 40 + (index * 1);
        buffer.putByte(pos, value);

        return this;
    }

    public static String makerOrderIDCharacterEncoding()
    {
        return java.nio.charset.StandardCharsets.US_ASCII.name();
    }

    public MatchEventEncoder putMakerOrderID(final byte[] src, final int srcOffset)
    {
        final int length = 36;
        if (srcOffset < 0 || srcOffset > (src.length - length))
        {
            throw new IndexOutOfBoundsException("Copy will go out of range: offset=" + srcOffset);
        }

        buffer.putBytes(offset + 40, src, srcOffset, length);

        return this;
    }

    public MatchEventEncoder makerOrderID(final String src)
    {
        final int length = 36;
        final int srcLength = null == src ? 0 : src.length();
        if (srcLength > length)
        {
            throw new IndexOutOfBoundsException("String too large for copy: byte length=" + srcLength);
        }

        buffer.putStringWithoutLengthAscii(offset + 40, src);

        for (int start = srcLength; start < length; ++start)
        {
            buffer.putByte(offset + 40 + start, (byte)0);
        }

        return this;
    }

    public MatchEventEncoder makerOrderID(final CharSequence src)
    {
        final int length = 36;
        final int srcLength = null == src ? 0 : src.length();
        if (srcLength > length)
        {
            throw new IndexOutOfBoundsException("CharSequence too large for copy: byte length=" + srcLength);
        }

        buffer.putStringWithoutLengthAscii(offset + 40, src);

        for (int start = srcLength; start < length; ++start)
        {
            buffer.putByte(offset + 40 + start, (byte)0);
        }

        return this;
    }

    public static int tradeIDId()
    {
        return 76;
    }

    public static int tradeIDSinceVersion()
    {
        return 0;
    }

    public static int tradeIDEncodingOffset()
    {
        return 76;
    }

    public static int tradeIDEncodingLength()
    {
        return 8;
    }

    public static String tradeIDMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "int";
        }

        return "";
    }

    public static long tradeIDNullValue()
    {
        return 0xffffffffffffffffL;
    }

    public static long tradeIDMinValue()
    {
        return 0x0L;
    }

    public static long tradeIDMaxValue()
    {
        return 0xfffffffffffffffeL;
    }

    public MatchEventEncoder tradeID(final long value)
    {
        buffer.putLong(offset + 76, value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int takerOrderIDId()
    {
        return 11;
    }

    public static int takerOrderIDSinceVersion()
    {
        return 0;
    }

    public static int takerOrderIDEncodingOffset()
    {
        return 84;
    }

    public static int takerOrderIDEncodingLength()
    {
        return 36;
    }

    public static String takerOrderIDMetaAttribute(final MetaAttribute metaAttribute)
    {
        if (MetaAttribute.PRESENCE == metaAttribute)
        {
            return "required";
        }
        if (MetaAttribute.SEMANTIC_TYPE == metaAttribute)
        {
            return "String";
        }

        return "";
    }

    public static byte takerOrderIDNullValue()
    {
        return (byte)0;
    }

    public static byte takerOrderIDMinValue()
    {
        return (byte)32;
    }

    public static byte takerOrderIDMaxValue()
    {
        return (byte)126;
    }

    public static int takerOrderIDLength()
    {
        return 36;
    }


    public MatchEventEncoder takerOrderID(final int index, final byte value)
    {
        if (index < 0 || index >= 36)
        {
            throw new IndexOutOfBoundsException("index out of range: index=" + index);
        }

        final int pos = offset + 84 + (index * 1);
        buffer.putByte(pos, value);

        return this;
    }

    public static String takerOrderIDCharacterEncoding()
    {
        return java.nio.charset.StandardCharsets.US_ASCII.name();
    }

    public MatchEventEncoder putTakerOrderID(final byte[] src, final int srcOffset)
    {
        final int length = 36;
        if (srcOffset < 0 || srcOffset > (src.length - length))
        {
            throw new IndexOutOfBoundsException("Copy will go out of range: offset=" + srcOffset);
        }

        buffer.putBytes(offset + 84, src, srcOffset, length);

        return this;
    }

    public MatchEventEncoder takerOrderID(final String src)
    {
        final int length = 36;
        final int srcLength = null == src ? 0 : src.length();
        if (srcLength > length)
        {
            throw new IndexOutOfBoundsException("String too large for copy: byte length=" + srcLength);
        }

        buffer.putStringWithoutLengthAscii(offset + 84, src);

        for (int start = srcLength; start < length; ++start)
        {
            buffer.putByte(offset + 84 + start, (byte)0);
        }

        return this;
    }

    public MatchEventEncoder takerOrderID(final CharSequence src)
    {
        final int length = 36;
        final int srcLength = null == src ? 0 : src.length();
        if (srcLength > length)
        {
            throw new IndexOutOfBoundsException("CharSequence too large for copy: byte length=" + srcLength);
        }

        buffer.putStringWithoutLengthAscii(offset + 84, src);

        for (int start = srcLength; start < length; ++start)
        {
            buffer.putByte(offset + 84 + start, (byte)0);
        }

        return this;
    }

    public String toString()
    {
        if (null == buffer)
        {
            return "";
        }

        return appendTo(new StringBuilder()).toString();
    }

    public StringBuilder appendTo(final StringBuilder builder)
    {
        if (null == buffer)
        {
            return builder;
        }

        final MatchEventDecoder decoder = new MatchEventDecoder();
        decoder.wrap(buffer, initialOffset, BLOCK_LENGTH, SCHEMA_VERSION);

        return decoder.appendTo(builder);
    }
}
