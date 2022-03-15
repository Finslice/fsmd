use crate::*;

pub use encoder::*;
pub use decoder::*;

pub const SBE_BLOCK_LENGTH: u16 = 84;
pub const SBE_TEMPLATE_ID: u16 = 50;
pub const SBE_SCHEMA_ID: u16 = 1;
pub const SBE_SCHEMA_VERSION: u16 = 1;

pub mod encoder {
    use super::*;

    #[derive(Debug, Default)]
    pub struct OrderChangeEncoder<'a> {
        buf: WriteBuf<'a>,
        initial_offset: usize,
        offset: usize,
        limit: usize,
    }

    impl<'a> Writer<'a> for OrderChangeEncoder<'a> {
        #[inline]
        fn get_buf_mut(&mut self) -> &mut WriteBuf<'a> {
            &mut self.buf
        }
    }

    impl<'a> Encoder<'a> for OrderChangeEncoder<'a> {
        #[inline]
        fn get_limit(&self) -> usize {
            self.limit
        }

        #[inline]
        fn set_limit(&mut self, limit: usize) {
            self.limit = limit;
        }
    }

    impl<'a> OrderChangeEncoder<'a> {
        pub fn wrap(mut self, buf: WriteBuf<'a>, offset: usize) -> Self {
            let limit = offset + SBE_BLOCK_LENGTH as usize;
            self.buf = buf;
            self.initial_offset = offset;
            self.offset = offset;
            self.limit = limit;
            self
        }

        #[inline]
        pub fn encoded_length(&self) -> usize {
            self.limit - self.offset
        }

        pub fn header(self, offset: usize) -> MessageHeaderEncoder<Self> {
            let mut header = MessageHeaderEncoder::default().wrap(self, offset);
            header.block_length(SBE_BLOCK_LENGTH);
            header.template_id(SBE_TEMPLATE_ID);
            header.schema_id(SBE_SCHEMA_ID);
            header.version(SBE_SCHEMA_VERSION);
            header
        }

        /// primitive field 'TransactTime'
        /// - min value: 0
        /// - max value: -2
        /// - null value: -1
        /// - characterEncoding: null
        /// - semanticType: UTCTimestamp
        /// - encodedOffset: 0
        /// - encodedLength: 8
        #[inline]
        pub fn transact_time(&mut self, value: u64) {
            let offset = self.offset;
            self.get_buf_mut().put_u64_at(offset, value);
        }

        /// primitive field 'SecurityID'
        /// - min value: 0
        /// - max value: 4294967294
        /// - null value: 4294967295
        /// - characterEncoding: null
        /// - semanticType: int
        /// - encodedOffset: 8
        /// - encodedLength: 4
        #[inline]
        pub fn security_id(&mut self, value: u32) {
            let offset = self.offset + 8;
            self.get_buf_mut().put_u32_at(offset, value);
        }

        /// REQUIRED enum
        #[inline]
        pub fn order_action(&mut self, value: OrderActionType) {
            let offset = self.offset + 12;
            self.get_buf_mut().put_u8_at(offset, value as u8)
        }

        /// REQUIRED enum
        #[inline]
        pub fn side(&mut self, value: SideType) {
            let offset = self.offset + 13;
            self.get_buf_mut().put_u8_at(offset, value as u8)
        }

        /// primitive field 'Price'
        /// - min value: 4.9E-324
        /// - max value: 1.7976931348623157E308
        /// - null value: NaN
        /// - characterEncoding: null
        /// - semanticType: double
        /// - encodedOffset: 16
        /// - encodedLength: 8
        #[inline]
        pub fn price(&mut self, value: f64) {
            let offset = self.offset + 16;
            self.get_buf_mut().put_f64_at(offset, value);
        }

        /// primitive field 'NewSize'
        /// - min value: 4.9E-324
        /// - max value: 1.7976931348623157E308
        /// - null value: NaN
        /// - characterEncoding: null
        /// - semanticType: double
        /// - encodedOffset: 24
        /// - encodedLength: 8
        #[inline]
        pub fn new_size(&mut self, value: f64) {
            let offset = self.offset + 24;
            self.get_buf_mut().put_f64_at(offset, value);
        }

        /// primitive field 'SecuritySequence'
        /// - min value: 0
        /// - max value: -2
        /// - null value: -1
        /// - characterEncoding: null
        /// - semanticType: int
        /// - encodedOffset: 32
        /// - encodedLength: 8
        #[inline]
        pub fn security_sequence(&mut self, value: u64) {
            let offset = self.offset + 32;
            self.get_buf_mut().put_u64_at(offset, value);
        }

        /// primitive array field 'OrderID'
        /// - min value: 32
        /// - max value: 126
        /// - null value: 0
        /// - characterEncoding: US-ASCII
        /// - semanticType: String
        /// - encodedOffset: 40
        /// - encodedLength: 36
        /// - version: 0
        #[inline]
        pub fn order_id(&mut self, value: [u8; 36]) {
            let offset = self.offset + 40;
            let buf = self.get_buf_mut();
            buf.put_u8_at(offset, value[0]);
            buf.put_u8_at(offset + 1, value[1]);
            buf.put_u8_at(offset + 2, value[2]);
            buf.put_u8_at(offset + 3, value[3]);
            buf.put_u8_at(offset + 4, value[4]);
            buf.put_u8_at(offset + 5, value[5]);
            buf.put_u8_at(offset + 6, value[6]);
            buf.put_u8_at(offset + 7, value[7]);
            buf.put_u8_at(offset + 8, value[8]);
            buf.put_u8_at(offset + 9, value[9]);
            buf.put_u8_at(offset + 10, value[10]);
            buf.put_u8_at(offset + 11, value[11]);
            buf.put_u8_at(offset + 12, value[12]);
            buf.put_u8_at(offset + 13, value[13]);
            buf.put_u8_at(offset + 14, value[14]);
            buf.put_u8_at(offset + 15, value[15]);
            buf.put_u8_at(offset + 16, value[16]);
            buf.put_u8_at(offset + 17, value[17]);
            buf.put_u8_at(offset + 18, value[18]);
            buf.put_u8_at(offset + 19, value[19]);
            buf.put_u8_at(offset + 20, value[20]);
            buf.put_u8_at(offset + 21, value[21]);
            buf.put_u8_at(offset + 22, value[22]);
            buf.put_u8_at(offset + 23, value[23]);
            buf.put_u8_at(offset + 24, value[24]);
            buf.put_u8_at(offset + 25, value[25]);
            buf.put_u8_at(offset + 26, value[26]);
            buf.put_u8_at(offset + 27, value[27]);
            buf.put_u8_at(offset + 28, value[28]);
            buf.put_u8_at(offset + 29, value[29]);
            buf.put_u8_at(offset + 30, value[30]);
            buf.put_u8_at(offset + 31, value[31]);
            buf.put_u8_at(offset + 32, value[32]);
            buf.put_u8_at(offset + 33, value[33]);
            buf.put_u8_at(offset + 34, value[34]);
            buf.put_u8_at(offset + 35, value[35]);
        }

        /// primitive field 'OldSize'
        /// - min value: 4.9E-324
        /// - max value: 1.7976931348623157E308
        /// - null value: NaN
        /// - characterEncoding: null
        /// - semanticType: double
        /// - encodedOffset: 76
        /// - encodedLength: 8
        #[inline]
        pub fn old_size(&mut self, value: f64) {
            let offset = self.offset + 76;
            self.get_buf_mut().put_f64_at(offset, value);
        }

    }

} // end encoder

pub mod decoder {
    use super::*;

    #[derive(Debug, Default)]
    pub struct OrderChangeDecoder<'a> {
        buf: ReadBuf<'a>,
        initial_offset: usize,
        offset: usize,
        limit: usize,
        pub acting_block_length: u16,
        pub acting_version: u16,
    }

    impl<'a> Reader<'a> for OrderChangeDecoder<'a> {
        #[inline]
        fn get_buf(&self) -> &ReadBuf<'a> {
            &self.buf
        }
    }

    impl<'a> Decoder<'a> for OrderChangeDecoder<'a> {
        #[inline]
        fn get_limit(&self) -> usize {
            self.limit
        }

        #[inline]
        fn set_limit(&mut self, limit: usize) {
            self.limit = limit;
        }
    }

    impl<'a> OrderChangeDecoder<'a> {
        pub fn wrap(
            mut self,
            buf: ReadBuf<'a>,
            offset: usize,
            acting_block_length: u16,
            acting_version: u16,
        ) -> Self {
            let limit = offset + acting_block_length as usize;
            self.buf = buf;
            self.initial_offset = offset;
            self.offset = offset;
            self.limit = limit;
            self.acting_block_length = acting_block_length;
            self.acting_version = acting_version;
            self
        }

        #[inline]
        pub fn encoded_length(&self) -> usize {
            self.limit - self.offset
        }

        pub fn header(self, mut header: MessageHeaderDecoder<ReadBuf<'a>>) -> Self {
            debug_assert_eq!(SBE_TEMPLATE_ID, header.template_id());
            let acting_block_length = header.block_length();
            let acting_version = header.version();

            self.wrap(
                header.parent().unwrap(),
                message_header_codec::ENCODED_LENGTH,
                acting_block_length,
                acting_version,
            )
        }

        /// primitive field - 'REQUIRED'
        #[inline]
        pub fn transact_time(&self) -> u64 {
            self.get_buf().get_u64_at(self.offset)
        }

        /// primitive field - 'REQUIRED'
        #[inline]
        pub fn security_id(&self) -> u32 {
            self.get_buf().get_u32_at(self.offset + 8)
        }

        /// REQUIRED enum
        #[inline]
        pub fn order_action(&self) -> OrderActionType {
            self.get_buf().get_u8_at(self.offset + 12).into()
        }

        /// REQUIRED enum
        #[inline]
        pub fn side(&self) -> SideType {
            self.get_buf().get_u8_at(self.offset + 13).into()
        }

        /// primitive field - 'REQUIRED'
        #[inline]
        pub fn price(&self) -> f64 {
            self.get_buf().get_f64_at(self.offset + 16)
        }

        /// primitive field - 'REQUIRED'
        #[inline]
        pub fn new_size(&self) -> f64 {
            self.get_buf().get_f64_at(self.offset + 24)
        }

        /// primitive field - 'REQUIRED'
        #[inline]
        pub fn security_sequence(&self) -> u64 {
            self.get_buf().get_u64_at(self.offset + 32)
        }

        #[inline]
        pub fn order_id(&self) -> [u8; 36] {
            let buf = self.get_buf();
            [
                buf.get_u8_at(self.offset + 40),
                buf.get_u8_at(self.offset + 40 + 1),
                buf.get_u8_at(self.offset + 40 + 2),
                buf.get_u8_at(self.offset + 40 + 3),
                buf.get_u8_at(self.offset + 40 + 4),
                buf.get_u8_at(self.offset + 40 + 5),
                buf.get_u8_at(self.offset + 40 + 6),
                buf.get_u8_at(self.offset + 40 + 7),
                buf.get_u8_at(self.offset + 40 + 8),
                buf.get_u8_at(self.offset + 40 + 9),
                buf.get_u8_at(self.offset + 40 + 10),
                buf.get_u8_at(self.offset + 40 + 11),
                buf.get_u8_at(self.offset + 40 + 12),
                buf.get_u8_at(self.offset + 40 + 13),
                buf.get_u8_at(self.offset + 40 + 14),
                buf.get_u8_at(self.offset + 40 + 15),
                buf.get_u8_at(self.offset + 40 + 16),
                buf.get_u8_at(self.offset + 40 + 17),
                buf.get_u8_at(self.offset + 40 + 18),
                buf.get_u8_at(self.offset + 40 + 19),
                buf.get_u8_at(self.offset + 40 + 20),
                buf.get_u8_at(self.offset + 40 + 21),
                buf.get_u8_at(self.offset + 40 + 22),
                buf.get_u8_at(self.offset + 40 + 23),
                buf.get_u8_at(self.offset + 40 + 24),
                buf.get_u8_at(self.offset + 40 + 25),
                buf.get_u8_at(self.offset + 40 + 26),
                buf.get_u8_at(self.offset + 40 + 27),
                buf.get_u8_at(self.offset + 40 + 28),
                buf.get_u8_at(self.offset + 40 + 29),
                buf.get_u8_at(self.offset + 40 + 30),
                buf.get_u8_at(self.offset + 40 + 31),
                buf.get_u8_at(self.offset + 40 + 32),
                buf.get_u8_at(self.offset + 40 + 33),
                buf.get_u8_at(self.offset + 40 + 34),
                buf.get_u8_at(self.offset + 40 + 35),
            ]
        }

        /// primitive field - 'REQUIRED'
        #[inline]
        pub fn old_size(&self) -> f64 {
            self.get_buf().get_f64_at(self.offset + 76)
        }

    }

} // end decoder

