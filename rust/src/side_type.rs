#[derive(Clone, Copy, Debug, PartialEq, Eq, PartialOrd, Ord, Hash)]
#[repr(u8)]
pub enum SideType {
    NoSide = 0x0_u8, 
    Buy = 0x1_u8, 
    Sell = 0x2_u8, 
    NullVal = 0xff_u8, 
}
impl Default for SideType {
    #[inline]
    fn default() -> Self { SideType::NullVal }
}
impl From<u8> for SideType {
    #[inline]
    fn from(v: u8) -> Self {
        match v {
            0x0_u8 => Self::NoSide, 
            0x1_u8 => Self::Buy, 
            0x2_u8 => Self::Sell, 
            _ => Self::NullVal,
        }
    }
}
