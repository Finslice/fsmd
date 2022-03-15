/* Generated SBE (Simple Binary Encoding) message codec. */
package finslice_md;

import org.agrona.DirectBuffer;


/**
 * Order Matched
 */
@SuppressWarnings("all")
public final class MatchEventDecoder
{
    public static final int BLOCK_LENGTH = 120;
    public static final int TEMPLATE_ID = 42;
    public static final int SCHEMA_ID = 1;
    public static final int SCHEMA_VERSION = 1;
    public static final java.nio.ByteOrder BYTE_ORDER = java.nio.ByteOrder.LITTLE_ENDIAN;

    private final MatchEventDecoder parentMessage = this;
    private DirectBuffer buffer;
    private int initialOffset;
    private int offset;
    private int limit;
    int actingBlockLength;
    int actingVersion;

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

    public DirectBuffer buffer()
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

    public MatchEventDecoder wrap(
        final DirectBuffer buffer,
        final int offset,
        final int actingBlockLength,
        final int actingVersion)
    {
        if (buffer != this.buffer)
        {
            this.buffer = buffer;
        }
        this.initialOffset = offset;
        this.offset = offset;
        this.actingBlockLength = actingBlockLength;
        this.actingVersion = actingVersion;
        limit(offset + actingBlockLength);

        return this;
    }

    public MatchEventDecoder wrapAndApplyHeader(
        final DirectBuffer buffer,
        final int offset,
        final MessageHeaderDecoder headerDecoder)
    {
        headerDecoder.wrap(buffer, offset);

        final int templateId = headerDecoder.templateId();
        if (TEMPLATE_ID != templateId)
        {
            throw new IllegalStateException("Invalid TEMPLATE_ID: " + templateId);
        }

        return wrap(
            buffer,
            offset + MessageHeaderDecoder.ENCODED_LENGTH,
            headerDecoder.blockLength(),
            headerDecoder.version());
    }

    public MatchEventDecoder sbeRewind()
    {
        return wrap(buffer, initialOffset, actingBlockLength, actingVersion);
    }

    public int sbeDecodedLength()
    {
        final int currentLimit = limit();
        sbeSkip();
        final int decodedLength = encodedLength();
        limit(currentLimit);

        return decodedLength;
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

    public long transactTime()
    {
        return buffer.getLong(offset + 0, java.nio.ByteOrder.LITTLE_ENDIAN);
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

    public long securityID()
    {
        return (buffer.getInt(offset + 8, java.nio.ByteOrder.LITTLE_ENDIAN) & 0xFFFF_FFFFL);
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

    public short orderActionRaw()
    {
        return ((short)(buffer.getByte(offset + 12) & 0xFF));
    }

    public OrderActionType orderAction()
    {
        return OrderActionType.get(((short)(buffer.getByte(offset + 12) & 0xFF)));
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

    public short makerSideRaw()
    {
        return ((short)(buffer.getByte(offset + 13) & 0xFF));
    }

    public SideType makerSide()
    {
        return SideType.get(((short)(buffer.getByte(offset + 13) & 0xFF)));
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

    public double price()
    {
        return buffer.getDouble(offset + 16, java.nio.ByteOrder.LITTLE_ENDIAN);
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

    public double size()
    {
        return buffer.getDouble(offset + 24, java.nio.ByteOrder.LITTLE_ENDIAN);
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

    public long securitySequence()
    {
        return buffer.getLong(offset + 32, java.nio.ByteOrder.LITTLE_ENDIAN);
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


    public byte makerOrderID(final int index)
    {
        if (index < 0 || index >= 36)
        {
            throw new IndexOutOfBoundsException("index out of range: index=" + index);
        }

        final int pos = offset + 40 + (index * 1);

        return buffer.getByte(pos);
    }


    public static String makerOrderIDCharacterEncoding()
    {
        return java.nio.charset.StandardCharsets.US_ASCII.name();
    }

    public int getMakerOrderID(final byte[] dst, final int dstOffset)
    {
        final int length = 36;
        if (dstOffset < 0 || dstOffset > (dst.length - length))
        {
            throw new IndexOutOfBoundsException("Copy will go out of range: offset=" + dstOffset);
        }

        buffer.getBytes(offset + 40, dst, dstOffset, length);

        return length;
    }

    public String makerOrderID()
    {
        final byte[] dst = new byte[36];
        buffer.getBytes(offset + 40, dst, 0, 36);

        int end = 0;
        for (; end < 36 && dst[end] != 0; ++end);

        return new String(dst, 0, end, java.nio.charset.StandardCharsets.US_ASCII);
    }


    public int getMakerOrderID(final Appendable value)
    {
        for (int i = 0; i < 36; ++i)
        {
            final int c = buffer.getByte(offset + 40 + i) & 0xFF;
            if (c == 0)
            {
                return i;
            }

            try
            {
                value.append(c > 127 ? '?' : (char)c);
            }
            catch (final java.io.IOException ex)
            {
                throw new java.io.UncheckedIOException(ex);
            }
        }

        return 36;
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

    public long tradeID()
    {
        return buffer.getLong(offset + 76, java.nio.ByteOrder.LITTLE_ENDIAN);
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


    public byte takerOrderID(final int index)
    {
        if (index < 0 || index >= 36)
        {
            throw new IndexOutOfBoundsException("index out of range: index=" + index);
        }

        final int pos = offset + 84 + (index * 1);

        return buffer.getByte(pos);
    }


    public static String takerOrderIDCharacterEncoding()
    {
        return java.nio.charset.StandardCharsets.US_ASCII.name();
    }

    public int getTakerOrderID(final byte[] dst, final int dstOffset)
    {
        final int length = 36;
        if (dstOffset < 0 || dstOffset > (dst.length - length))
        {
            throw new IndexOutOfBoundsException("Copy will go out of range: offset=" + dstOffset);
        }

        buffer.getBytes(offset + 84, dst, dstOffset, length);

        return length;
    }

    public String takerOrderID()
    {
        final byte[] dst = new byte[36];
        buffer.getBytes(offset + 84, dst, 0, 36);

        int end = 0;
        for (; end < 36 && dst[end] != 0; ++end);

        return new String(dst, 0, end, java.nio.charset.StandardCharsets.US_ASCII);
    }


    public int getTakerOrderID(final Appendable value)
    {
        for (int i = 0; i < 36; ++i)
        {
            final int c = buffer.getByte(offset + 84 + i) & 0xFF;
            if (c == 0)
            {
                return i;
            }

            try
            {
                value.append(c > 127 ? '?' : (char)c);
            }
            catch (final java.io.IOException ex)
            {
                throw new java.io.UncheckedIOException(ex);
            }
        }

        return 36;
    }


    public String toString()
    {
        if (null == buffer)
        {
            return "";
        }

        final MatchEventDecoder decoder = new MatchEventDecoder();
        decoder.wrap(buffer, initialOffset, actingBlockLength, actingVersion);

        return decoder.appendTo(new StringBuilder()).toString();
    }

    public StringBuilder appendTo(final StringBuilder builder)
    {
        if (null == buffer)
        {
            return builder;
        }

        final int originalLimit = limit();
        limit(initialOffset + actingBlockLength);
        builder.append("[MatchEvent](sbeTemplateId=");
        builder.append(TEMPLATE_ID);
        builder.append("|sbeSchemaId=");
        builder.append(SCHEMA_ID);
        builder.append("|sbeSchemaVersion=");
        if (parentMessage.actingVersion != SCHEMA_VERSION)
        {
            builder.append(parentMessage.actingVersion);
            builder.append('/');
        }
        builder.append(SCHEMA_VERSION);
        builder.append("|sbeBlockLength=");
        if (actingBlockLength != BLOCK_LENGTH)
        {
            builder.append(actingBlockLength);
            builder.append('/');
        }
        builder.append(BLOCK_LENGTH);
        builder.append("):");
        builder.append("transactTime=");
        builder.append(transactTime());
        builder.append('|');
        builder.append("securityID=");
        builder.append(securityID());
        builder.append('|');
        builder.append("orderAction=");
        builder.append(orderAction());
        builder.append('|');
        builder.append("makerSide=");
        builder.append(makerSide());
        builder.append('|');
        builder.append("price=");
        builder.append(price());
        builder.append('|');
        builder.append("size=");
        builder.append(size());
        builder.append('|');
        builder.append("securitySequence=");
        builder.append(securitySequence());
        builder.append('|');
        builder.append("makerOrderID=");
        for (int i = 0; i < makerOrderIDLength() && makerOrderID(i) > 0; i++)
        {
            builder.append((char)makerOrderID(i));
        }
        builder.append('|');
        builder.append("tradeID=");
        builder.append(tradeID());
        builder.append('|');
        builder.append("takerOrderID=");
        for (int i = 0; i < takerOrderIDLength() && takerOrderID(i) > 0; i++)
        {
            builder.append((char)takerOrderID(i));
        }

        limit(originalLimit);

        return builder;
    }
    
    public MatchEventDecoder sbeSkip()
    {
        sbeRewind();

        return this;
    }
}
