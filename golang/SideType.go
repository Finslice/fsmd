// Generated SBE (Simple Binary Encoding) message codec

package finslice_md

import (
	"fmt"
	"io"
	"reflect"
)

type SideTypeEnum uint8
type SideTypeValues struct {
	NoSide    SideTypeEnum
	Buy       SideTypeEnum
	Sell      SideTypeEnum
	NullValue SideTypeEnum
}

var SideType = SideTypeValues{0, 1, 2, 255}

func (s SideTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SideTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SideTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SideType)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SideType, unknown enumeration value %d", s)
}

func (*SideTypeEnum) EncodedLength() int64 {
	return 1
}

func (*SideTypeEnum) NoSideSinceVersion() uint16 {
	return 0
}

func (s *SideTypeEnum) NoSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NoSideSinceVersion()
}

func (*SideTypeEnum) NoSideDeprecated() uint16 {
	return 0
}

func (*SideTypeEnum) BuySinceVersion() uint16 {
	return 0
}

func (s *SideTypeEnum) BuyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BuySinceVersion()
}

func (*SideTypeEnum) BuyDeprecated() uint16 {
	return 0
}

func (*SideTypeEnum) SellSinceVersion() uint16 {
	return 0
}

func (s *SideTypeEnum) SellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SellSinceVersion()
}

func (*SideTypeEnum) SellDeprecated() uint16 {
	return 0
}
