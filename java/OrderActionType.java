/* Generated SBE (Simple Binary Encoding) message codec. */
package finslice_md;

public enum OrderActionType
{

    /**
     * New Market Order
     */
    NewMarketOrder((short)1),


    /**
     * New Limit Order
     */
    NewLimitOrder((short)2),


    /**
     * Order Open
     */
    OrderOpen((short)3),


    /**
     * Order done Filled
     */
    OrderFilled((short)8),


    /**
     * Order done Cancelled
     */
    OrderCancel((short)9),


    /**
     * Match(trade)
     */
    MatchEvent((short)12),


    /**
     * Order size changed
     */
    OrderChange((short)13),

    /**
     * To be used to represent not present or null.
     */
    NULL_VAL((short)255);

    private final short value;

    OrderActionType(final short value)
    {
        this.value = value;
    }

    /**
     * The raw encoded value in the Java type representation.
     *
     * @return the raw value encoded.
     */
    public short value()
    {
        return value;
    }

    /**
     * Lookup the enum value representing the value.
     *
     * @param value encoded to be looked up.
     * @return the enum value representing the value.
     */
    public static OrderActionType get(final short value)
    {
        switch (value)
        {
            case 1: return NewMarketOrder;
            case 2: return NewLimitOrder;
            case 3: return OrderOpen;
            case 8: return OrderFilled;
            case 9: return OrderCancel;
            case 12: return MatchEvent;
            case 13: return OrderChange;
            case 255: return NULL_VAL;
        }

        throw new IllegalArgumentException("Unknown value: " + value);
    }
}
