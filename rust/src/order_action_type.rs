#[derive(Clone, Copy, Debug, PartialEq, Eq, PartialOrd, Ord, Hash)]
#[repr(u8)]
pub enum OrderActionType {
    NewMarketOrder = 0x1_u8, 
    NewLimitOrder = 0x2_u8, 
    OrderOpen = 0x3_u8, 
    OrderFilled = 0x8_u8, 
    OrderCancel = 0x9_u8, 
    MatchEvent = 0xc_u8, 
    OrderChange = 0xd_u8, 
    NullVal = 0xff_u8, 
}
impl Default for OrderActionType {
    #[inline]
    fn default() -> Self { OrderActionType::NullVal }
}
impl From<u8> for OrderActionType {
    #[inline]
    fn from(v: u8) -> Self {
        match v {
            0x1_u8 => Self::NewMarketOrder, 
            0x2_u8 => Self::NewLimitOrder, 
            0x3_u8 => Self::OrderOpen, 
            0x8_u8 => Self::OrderFilled, 
            0x9_u8 => Self::OrderCancel, 
            0xc_u8 => Self::MatchEvent, 
            0xd_u8 => Self::OrderChange, 
            _ => Self::NullVal,
        }
    }
}
