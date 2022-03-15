// Generated SBE (Simple Binary Encoding) message codec

package finslice_md

import (
	"fmt"
	"io"
	"reflect"
)

type OrderActionTypeEnum uint8
type OrderActionTypeValues struct {
	NewMarketOrder OrderActionTypeEnum
	NewLimitOrder  OrderActionTypeEnum
	OrderOpen      OrderActionTypeEnum
	OrderFilled    OrderActionTypeEnum
	OrderCancel    OrderActionTypeEnum
	MatchEvent     OrderActionTypeEnum
	OrderChange    OrderActionTypeEnum
	NullValue      OrderActionTypeEnum
}

var OrderActionType = OrderActionTypeValues{1, 2, 3, 8, 9, 12, 13, 255}

func (o OrderActionTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderActionTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderActionTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderActionType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderActionType, unknown enumeration value %d", o)
}

func (*OrderActionTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrderActionTypeEnum) NewMarketOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderActionTypeEnum) NewMarketOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewMarketOrderSinceVersion()
}

func (*OrderActionTypeEnum) NewMarketOrderDeprecated() uint16 {
	return 0
}

func (*OrderActionTypeEnum) NewLimitOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderActionTypeEnum) NewLimitOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewLimitOrderSinceVersion()
}

func (*OrderActionTypeEnum) NewLimitOrderDeprecated() uint16 {
	return 0
}

func (*OrderActionTypeEnum) OrderOpenSinceVersion() uint16 {
	return 0
}

func (o *OrderActionTypeEnum) OrderOpenInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderOpenSinceVersion()
}

func (*OrderActionTypeEnum) OrderOpenDeprecated() uint16 {
	return 0
}

func (*OrderActionTypeEnum) OrderFilledSinceVersion() uint16 {
	return 0
}

func (o *OrderActionTypeEnum) OrderFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderFilledSinceVersion()
}

func (*OrderActionTypeEnum) OrderFilledDeprecated() uint16 {
	return 0
}

func (*OrderActionTypeEnum) OrderCancelSinceVersion() uint16 {
	return 0
}

func (o *OrderActionTypeEnum) OrderCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderCancelSinceVersion()
}

func (*OrderActionTypeEnum) OrderCancelDeprecated() uint16 {
	return 0
}

func (*OrderActionTypeEnum) MatchEventSinceVersion() uint16 {
	return 0
}

func (o *OrderActionTypeEnum) MatchEventInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MatchEventSinceVersion()
}

func (*OrderActionTypeEnum) MatchEventDeprecated() uint16 {
	return 0
}

func (*OrderActionTypeEnum) OrderChangeSinceVersion() uint16 {
	return 0
}

func (o *OrderActionTypeEnum) OrderChangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderChangeSinceVersion()
}

func (*OrderActionTypeEnum) OrderChangeDeprecated() uint16 {
	return 0
}
