// Generated SBE (Simple Binary Encoding) message codec

package finslice_md

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderUpdate struct {
	TransactTime     uint64
	SecurityID       uint32
	OrderAction      OrderActionTypeEnum
	Side             SideTypeEnum
	Price            float64
	Size             float64
	SecuritySequence uint64
	OrderID          [36]byte
}

func (o *OrderUpdate) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, o.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, o.SecurityID); err != nil {
		return err
	}
	if err := o.OrderAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := _m.WriteFloat64(_w, o.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.Size); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SecuritySequence); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrderID[:]); err != nil {
		return err
	}
	return nil
}

func (o *OrderUpdate) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.TransactTimeInActingVersion(actingVersion) {
		o.TransactTime = o.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.TransactTime); err != nil {
			return err
		}
	}
	if !o.SecurityIDInActingVersion(actingVersion) {
		o.SecurityID = o.SecurityIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &o.SecurityID); err != nil {
			return err
		}
	}
	if o.OrderActionInActingVersion(actingVersion) {
		if err := o.OrderAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 2)
	if !o.PriceInActingVersion(actingVersion) {
		o.Price = o.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.Price); err != nil {
			return err
		}
	}
	if !o.SizeInActingVersion(actingVersion) {
		o.Size = o.SizeNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.Size); err != nil {
			return err
		}
	}
	if !o.SecuritySequenceInActingVersion(actingVersion) {
		o.SecuritySequence = o.SecuritySequenceNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.SecuritySequence); err != nil {
			return err
		}
	}
	if !o.OrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 36; idx++ {
			o.OrderID[idx] = o.OrderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.OrderID[:]); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderUpdate) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.TransactTimeInActingVersion(actingVersion) {
		if o.TransactTime < o.TransactTimeMinValue() || o.TransactTime > o.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.TransactTime (%v < %v > %v)", o.TransactTimeMinValue(), o.TransactTime, o.TransactTimeMaxValue())
		}
	}
	if o.SecurityIDInActingVersion(actingVersion) {
		if o.SecurityID < o.SecurityIDMinValue() || o.SecurityID > o.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on o.SecurityID (%v < %v > %v)", o.SecurityIDMinValue(), o.SecurityID, o.SecurityIDMaxValue())
		}
	}
	if err := o.OrderAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.PriceInActingVersion(actingVersion) {
		if o.Price < o.PriceMinValue() || o.Price > o.PriceMaxValue() {
			return fmt.Errorf("Range check failed on o.Price (%v < %v > %v)", o.PriceMinValue(), o.Price, o.PriceMaxValue())
		}
	}
	if o.SizeInActingVersion(actingVersion) {
		if o.Size < o.SizeMinValue() || o.Size > o.SizeMaxValue() {
			return fmt.Errorf("Range check failed on o.Size (%v < %v > %v)", o.SizeMinValue(), o.Size, o.SizeMaxValue())
		}
	}
	if o.SecuritySequenceInActingVersion(actingVersion) {
		if o.SecuritySequence < o.SecuritySequenceMinValue() || o.SecuritySequence > o.SecuritySequenceMaxValue() {
			return fmt.Errorf("Range check failed on o.SecuritySequence (%v < %v > %v)", o.SecuritySequenceMinValue(), o.SecuritySequence, o.SecuritySequenceMaxValue())
		}
	}
	if o.OrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 36; idx++ {
			if o.OrderID[idx] < o.OrderIDMinValue() || o.OrderID[idx] > o.OrderIDMaxValue() {
				return fmt.Errorf("Range check failed on o.OrderID[%d] (%v < %v > %v)", idx, o.OrderIDMinValue(), o.OrderID[idx], o.OrderIDMaxValue())
			}
		}
	}
	for idx, ch := range o.OrderID {
		if ch > 127 {
			return fmt.Errorf("o.OrderID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func OrderUpdateInit(o *OrderUpdate) {
	return
}

func (*OrderUpdate) SbeBlockLength() (blockLength uint16) {
	return 76
}

func (*OrderUpdate) SbeTemplateId() (templateId uint16) {
	return 20
}

func (*OrderUpdate) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderUpdate) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderUpdate) SbeSemanticType() (semanticType []byte) {
	return []byte("D")
}

func (*OrderUpdate) TransactTimeId() uint16 {
	return 60
}

func (*OrderUpdate) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderUpdate) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "UTCTimestamp"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderUpdate) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderUpdate) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderUpdate) SecurityIDId() uint16 {
	return 48
}

func (*OrderUpdate) SecurityIDSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityIDSinceVersion()
}

func (*OrderUpdate) SecurityIDDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) SecurityIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) SecurityIDMinValue() uint32 {
	return 0
}

func (*OrderUpdate) SecurityIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderUpdate) SecurityIDNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderUpdate) OrderActionId() uint16 {
	return 279
}

func (*OrderUpdate) OrderActionSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) OrderActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderActionSinceVersion()
}

func (*OrderUpdate) OrderActionDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) OrderActionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) SideId() uint16 {
	return 54
}

func (*OrderUpdate) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderUpdate) SideDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) SideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) PriceId() uint16 {
	return 270
}

func (*OrderUpdate) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderUpdate) PriceDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) PriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "double"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderUpdate) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderUpdate) PriceNullValue() float64 {
	return math.NaN()
}

func (*OrderUpdate) SizeId() uint16 {
	return 271
}

func (*OrderUpdate) SizeSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) SizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SizeSinceVersion()
}

func (*OrderUpdate) SizeDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) SizeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "double"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) SizeMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderUpdate) SizeMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderUpdate) SizeNullValue() float64 {
	return math.NaN()
}

func (*OrderUpdate) SecuritySequenceId() uint16 {
	return 83
}

func (*OrderUpdate) SecuritySequenceSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) SecuritySequenceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecuritySequenceSinceVersion()
}

func (*OrderUpdate) SecuritySequenceDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) SecuritySequenceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "int"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) SecuritySequenceMinValue() uint64 {
	return 0
}

func (*OrderUpdate) SecuritySequenceMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderUpdate) SecuritySequenceNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderUpdate) OrderIDId() uint16 {
	return 37
}

func (*OrderUpdate) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderUpdate) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderUpdate) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderUpdate) OrderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return "String"
	case 4:
		return "required"
	}
	return ""
}

func (*OrderUpdate) OrderIDMinValue() byte {
	return byte(32)
}

func (*OrderUpdate) OrderIDMaxValue() byte {
	return byte(126)
}

func (*OrderUpdate) OrderIDNullValue() byte {
	return 0
}

func (o *OrderUpdate) OrderIDCharacterEncoding() string {
	return "US-ASCII"
}
