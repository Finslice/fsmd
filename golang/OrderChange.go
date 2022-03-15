// Generated SBE (Simple Binary Encoding) message codec

package finslice_md

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderChange struct {
	TransactTime     uint64
	SecurityID       uint32
	OrderAction      OrderActionTypeEnum
	Side             SideTypeEnum
	Price            float64
	NewSize          float64
	SecuritySequence uint64
	OrderID          [36]byte
	OldSize          float64
}

func (o *OrderChange) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteFloat64(_w, o.NewSize); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.SecuritySequence); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrderID[:]); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.OldSize); err != nil {
		return err
	}
	return nil
}

func (o *OrderChange) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !o.NewSizeInActingVersion(actingVersion) {
		o.NewSize = o.NewSizeNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.NewSize); err != nil {
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
	if !o.OldSizeInActingVersion(actingVersion) {
		o.OldSize = o.OldSizeNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.OldSize); err != nil {
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

func (o *OrderChange) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if o.NewSizeInActingVersion(actingVersion) {
		if o.NewSize < o.NewSizeMinValue() || o.NewSize > o.NewSizeMaxValue() {
			return fmt.Errorf("Range check failed on o.NewSize (%v < %v > %v)", o.NewSizeMinValue(), o.NewSize, o.NewSizeMaxValue())
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
	if o.OldSizeInActingVersion(actingVersion) {
		if o.OldSize < o.OldSizeMinValue() || o.OldSize > o.OldSizeMaxValue() {
			return fmt.Errorf("Range check failed on o.OldSize (%v < %v > %v)", o.OldSizeMinValue(), o.OldSize, o.OldSizeMaxValue())
		}
	}
	return nil
}

func OrderChangeInit(o *OrderChange) {
	return
}

func (*OrderChange) SbeBlockLength() (blockLength uint16) {
	return 84
}

func (*OrderChange) SbeTemplateId() (templateId uint16) {
	return 50
}

func (*OrderChange) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderChange) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderChange) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*OrderChange) TransactTimeId() uint16 {
	return 60
}

func (*OrderChange) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderChange) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderChange) TransactTimeMetaAttribute(meta int) string {
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

func (*OrderChange) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderChange) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderChange) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderChange) SecurityIDId() uint16 {
	return 48
}

func (*OrderChange) SecurityIDSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecurityIDSinceVersion()
}

func (*OrderChange) SecurityIDDeprecated() uint16 {
	return 0
}

func (*OrderChange) SecurityIDMetaAttribute(meta int) string {
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

func (*OrderChange) SecurityIDMinValue() uint32 {
	return 0
}

func (*OrderChange) SecurityIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*OrderChange) SecurityIDNullValue() uint32 {
	return math.MaxUint32
}

func (*OrderChange) OrderActionId() uint16 {
	return 279
}

func (*OrderChange) OrderActionSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) OrderActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderActionSinceVersion()
}

func (*OrderChange) OrderActionDeprecated() uint16 {
	return 0
}

func (*OrderChange) OrderActionMetaAttribute(meta int) string {
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

func (*OrderChange) SideId() uint16 {
	return 54
}

func (*OrderChange) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderChange) SideDeprecated() uint16 {
	return 0
}

func (*OrderChange) SideMetaAttribute(meta int) string {
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

func (*OrderChange) PriceId() uint16 {
	return 270
}

func (*OrderChange) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderChange) PriceDeprecated() uint16 {
	return 0
}

func (*OrderChange) PriceMetaAttribute(meta int) string {
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

func (*OrderChange) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderChange) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderChange) PriceNullValue() float64 {
	return math.NaN()
}

func (*OrderChange) NewSizeId() uint16 {
	return 271
}

func (*OrderChange) NewSizeSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) NewSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewSizeSinceVersion()
}

func (*OrderChange) NewSizeDeprecated() uint16 {
	return 0
}

func (*OrderChange) NewSizeMetaAttribute(meta int) string {
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

func (*OrderChange) NewSizeMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderChange) NewSizeMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderChange) NewSizeNullValue() float64 {
	return math.NaN()
}

func (*OrderChange) SecuritySequenceId() uint16 {
	return 83
}

func (*OrderChange) SecuritySequenceSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) SecuritySequenceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SecuritySequenceSinceVersion()
}

func (*OrderChange) SecuritySequenceDeprecated() uint16 {
	return 0
}

func (*OrderChange) SecuritySequenceMetaAttribute(meta int) string {
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

func (*OrderChange) SecuritySequenceMinValue() uint64 {
	return 0
}

func (*OrderChange) SecuritySequenceMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderChange) SecuritySequenceNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderChange) OrderIDId() uint16 {
	return 37
}

func (*OrderChange) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderChange) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderChange) OrderIDMetaAttribute(meta int) string {
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

func (*OrderChange) OrderIDMinValue() byte {
	return byte(32)
}

func (*OrderChange) OrderIDMaxValue() byte {
	return byte(126)
}

func (*OrderChange) OrderIDNullValue() byte {
	return 0
}

func (o *OrderChange) OrderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderChange) OldSizeId() uint16 {
	return 273
}

func (*OrderChange) OldSizeSinceVersion() uint16 {
	return 0
}

func (o *OrderChange) OldSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OldSizeSinceVersion()
}

func (*OrderChange) OldSizeDeprecated() uint16 {
	return 0
}

func (*OrderChange) OldSizeMetaAttribute(meta int) string {
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

func (*OrderChange) OldSizeMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderChange) OldSizeMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderChange) OldSizeNullValue() float64 {
	return math.NaN()
}
