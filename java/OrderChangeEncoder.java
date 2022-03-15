/* Generated SBE (Simple Binary Encoding) message codec. */
package finslice_md;

import org.agrona.MutableDirectBuffer;


/**
 * Order Modify down
 */
@SuppressWarnings("all")
public final class OrderChangeEncoder
{
    public static final int BLOCK_LENGTH = 84;
    public static final int TEMPLATE_ID = 50;
    public static final int SCHEMA_ID = 1;
    public static final int SCHEMA_VERSION = 1;
    public static final java.nio.ByteOrder BYTE_ORDER = java.nio.ByteOrder.LITTLE_ENDIAN;

    private final OrderChangeEncoder parentMessage = this;
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

    public OrderChangeEncoder wrap(final MutableDirectBuffer buffer, final int offset)
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

    public OrderChangeEncoder wrapAndApplyHeader(
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

    public OrderChangeEncoder transactTime(final long value)
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

    public OrderChangeEncoder securityID(final long value)
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

    public OrderChangeEncoder orderAction(final OrderActionType value)
    {
        buffer.putByte(offset + 12, (byte)value.value());
        return this;
    }

    public static int sideId()
    {
        return 54;
    }

    public static int sideSinceVersion()
    {
        return 0;
    }

    public static int sideEncodingOffset()
    {
        return 13;
    }

    public static int sideEncodingLength()
    {
        return 1;
    }

    public static String sideMetaAttribute(final MetaAttribute metaAttribute)
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

    public OrderChangeEncoder side(final SideType value)
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

    public OrderChangeEncoder price(final double value)
    {
        buffer.putDouble(offset + 16, value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int newSizeId()
    {
        return 271;
    }

    public static int newSizeSinceVersion()
    {
        return 0;
    }

    public static int newSizeEncodingOffset()
    {
        return 24;
    }

    public static int newSizeEncodingLength()
    {
        return 8;
    }

    public static String newSizeMetaAttribute(final MetaAttribute metaAttribute)
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

    public static double newSizeNullValue()
    {
        return Double.NaN;
    }

    public static double newSizeMinValue()
    {
        return 4.9E-324d;
    }

    public static double newSizeMaxValue()
    {
        return 1.7976931348623157E308d;
    }

    public OrderChangeEncoder newSize(final double value)
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

    public OrderChangeEncoder securitySequence(final long value)
    {
        buffer.putLong(offset + 32, value, java.nio.ByteOrder.LITTLE_ENDIAN);
        return this;
    }


    public static int orderIDId()
    {
        return 37;
    }

    public static int orderIDSinceVersion()
    {
        return 0;
    }

    public static int orderIDEncodingOffset()
    {
        return 40;
    }

    public static int orderIDEncodingLength()
    {
        return 36;
    }

    public static String orderIDMetaAttribute(final MetaAttribute metaAttribute)
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

    public static byte orderIDNullValue()
    {
        return (byte)0;
    }

    public static byte orderIDMinValue()
    {
        return (byte)32;
    }

    public static byte orderIDMaxValue()
    {
        return (byte)126;
    }

    public static int orderIDLength()
    {
        return 36;
    }


    public OrderChangeEncoder orderID(final int index, final byte value)
    {
        if (index < 0 || index >= 36)
        {
            throw new IndexOutOfBoundsException("index out of range: index=" + index);
        }

        final int pos = offset + 40 + (index * 1);
        buffer.putByte(pos, value);

        return this;
    }

    public static String orderIDCharacterEncoding()
    {
        return java.nio.charset.StandardCharsets.US_ASCII.name();
    }

    public OrderChangeEncoder putOrderID(final byte[] src, final int srcOffset)
    {
        final int length = 36;
        if (srcOffset < 0 || srcOffset > (src.length - length))
        {
            throw new IndexOutOfBoundsException("Copy will go out of range: offset=" + srcOffset);
        }

        buffer.putBytes(offset + 40, src, srcOffset, length);

        return this;
    }

    public OrderChangeEncoder orderID(final String src)
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

    public OrderChangeEncoder orderID(final CharSequence src)
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

    public static int oldSizeId()
    {
        return 273;
    }

    public static int oldSizeSinceVersion()
    {
        return 0;
    }

    public static int oldSizeEncodingOffset()
    {
        return 76;
    }

    public static int oldSizeEncodingLength()
    {
        return 8;
    }

    public static String oldSizeMetaAttribute(final MetaAttribute metaAttribute)
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

    public static double oldSizeNullValue()
    {
        return Double.NaN;
    }

    public static double oldSizeMinValue()
    {
        return 4.9E-324d;
    }

    public static double oldSizeMaxValue()
    {
        return 1.7976931348623157E308d;
    }

    public OrderChangeEncoder oldSize(final double value)
    {
        buffer.putDouble(offset + 76, value, java.nio.ByteOrder.LITTLE_ENDIAN);
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

        final OrderChangeDecoder decoder = new OrderChangeDecoder();
        decoder.wrap(buffer, initialOffset, BLOCK_LENGTH, SCHEMA_VERSION);

        return decoder.appendTo(builder);
    }
}
