// Generated SBE (Simple Binary Encoding) message codec

package finslice_md

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MatchEvent struct {
	TransactTime     uint64
	SecurityID       uint32
	OrderAction      OrderActionTypeEnum
	MakerSide        SideTypeEnum
	Price            float64
	Size             float64
	SecuritySequence uint64
	MakerOrderID     [36]byte
	TradeID          uint64
	TakerOrderID     [36]byte
}

func (m *MatchEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, m.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, m.SecurityID); err != nil {
		return err
	}
	if err := m.OrderAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.MakerSide.Encode(_m, _w); err != nil {
		return err
	}

	for i := 0; i < 2; i++ {
		if err := _m.WriteUint8(_w, uint8(0)); err != nil {
			return err
		}
	}
	if err := _m.WriteFloat64(_w, m.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, m.Size); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.SecuritySequence); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.MakerOrderID[:]); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.TradeID); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.TakerOrderID[:]); err != nil {
		return err
	}
	return nil
}

func (m *MatchEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.TransactTimeInActingVersion(actingVersion) {
		m.TransactTime = m.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.TransactTime); err != nil {
			return err
		}
	}
	if !m.SecurityIDInActingVersion(actingVersion) {
		m.SecurityID = m.SecurityIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &m.SecurityID); err != nil {
			return err
		}
	}
	if m.OrderActionInActingVersion(actingVersion) {
		if err := m.OrderAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.MakerSideInActingVersion(actingVersion) {
		if err := m.MakerSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	io.CopyN(ioutil.Discard, _r, 2)
	if !m.PriceInActingVersion(actingVersion) {
		m.Price = m.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &m.Price); err != nil {
			return err
		}
	}
	if !m.SizeInActingVersion(actingVersion) {
		m.Size = m.SizeNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &m.Size); err != nil {
			return err
		}
	}
	if !m.SecuritySequenceInActingVersion(actingVersion) {
		m.SecuritySequence = m.SecuritySequenceNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.SecuritySequence); err != nil {
			return err
		}
	}
	if !m.MakerOrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 36; idx++ {
			m.MakerOrderID[idx] = m.MakerOrderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.MakerOrderID[:]); err != nil {
			return err
		}
	}
	if !m.TradeIDInActingVersion(actingVersion) {
		m.TradeID = m.TradeIDNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.TradeID); err != nil {
			return err
		}
	}
	if !m.TakerOrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 36; idx++ {
			m.TakerOrderID[idx] = m.TakerOrderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.TakerOrderID[:]); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MatchEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TransactTimeInActingVersion(actingVersion) {
		if m.TransactTime < m.TransactTimeMinValue() || m.TransactTime > m.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.TransactTime (%v < %v > %v)", m.TransactTimeMinValue(), m.TransactTime, m.TransactTimeMaxValue())
		}
	}
	if m.SecurityIDInActingVersion(actingVersion) {
		if m.SecurityID < m.SecurityIDMinValue() || m.SecurityID > m.SecurityIDMaxValue() {
			return fmt.Errorf("Range check failed on m.SecurityID (%v < %v > %v)", m.SecurityIDMinValue(), m.SecurityID, m.SecurityIDMaxValue())
		}
	}
	if err := m.OrderAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.MakerSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.PriceInActingVersion(actingVersion) {
		if m.Price < m.PriceMinValue() || m.Price > m.PriceMaxValue() {
			return fmt.Errorf("Range check failed on m.Price (%v < %v > %v)", m.PriceMinValue(), m.Price, m.PriceMaxValue())
		}
	}
	if m.SizeInActingVersion(actingVersion) {
		if m.Size < m.SizeMinValue() || m.Size > m.SizeMaxValue() {
			return fmt.Errorf("Range check failed on m.Size (%v < %v > %v)", m.SizeMinValue(), m.Size, m.SizeMaxValue())
		}
	}
	if m.SecuritySequenceInActingVersion(actingVersion) {
		if m.SecuritySequence < m.SecuritySequenceMinValue() || m.SecuritySequence > m.SecuritySequenceMaxValue() {
			return fmt.Errorf("Range check failed on m.SecuritySequence (%v < %v > %v)", m.SecuritySequenceMinValue(), m.SecuritySequence, m.SecuritySequenceMaxValue())
		}
	}
	if m.MakerOrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 36; idx++ {
			if m.MakerOrderID[idx] < m.MakerOrderIDMinValue() || m.MakerOrderID[idx] > m.MakerOrderIDMaxValue() {
				return fmt.Errorf("Range check failed on m.MakerOrderID[%d] (%v < %v > %v)", idx, m.MakerOrderIDMinValue(), m.MakerOrderID[idx], m.MakerOrderIDMaxValue())
			}
		}
	}
	for idx, ch := range m.MakerOrderID {
		if ch > 127 {
			return fmt.Errorf("m.MakerOrderID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if m.TradeIDInActingVersion(actingVersion) {
		if m.TradeID < m.TradeIDMinValue() || m.TradeID > m.TradeIDMaxValue() {
			return fmt.Errorf("Range check failed on m.TradeID (%v < %v > %v)", m.TradeIDMinValue(), m.TradeID, m.TradeIDMaxValue())
		}
	}
	if m.TakerOrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 36; idx++ {
			if m.TakerOrderID[idx] < m.TakerOrderIDMinValue() || m.TakerOrderID[idx] > m.TakerOrderIDMaxValue() {
				return fmt.Errorf("Range check failed on m.TakerOrderID[%d] (%v < %v > %v)", idx, m.TakerOrderIDMinValue(), m.TakerOrderID[idx], m.TakerOrderIDMaxValue())
			}
		}
	}
	for idx, ch := range m.TakerOrderID {
		if ch > 127 {
			return fmt.Errorf("m.TakerOrderID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func MatchEventInit(m *MatchEvent) {
	return
}

func (*MatchEvent) SbeBlockLength() (blockLength uint16) {
	return 120
}

func (*MatchEvent) SbeTemplateId() (templateId uint16) {
	return 42
}

func (*MatchEvent) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MatchEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MatchEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MatchEvent) TransactTimeId() uint16 {
	return 60
}

func (*MatchEvent) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MatchEvent) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MatchEvent) TransactTimeMetaAttribute(meta int) string {
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

func (*MatchEvent) TransactTimeMinValue() uint64 {
	return 0
}

func (*MatchEvent) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MatchEvent) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MatchEvent) SecurityIDId() uint16 {
	return 48
}

func (*MatchEvent) SecurityIDSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) SecurityIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecurityIDSinceVersion()
}

func (*MatchEvent) SecurityIDDeprecated() uint16 {
	return 0
}

func (*MatchEvent) SecurityIDMetaAttribute(meta int) string {
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

func (*MatchEvent) SecurityIDMinValue() uint32 {
	return 0
}

func (*MatchEvent) SecurityIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*MatchEvent) SecurityIDNullValue() uint32 {
	return math.MaxUint32
}

func (*MatchEvent) OrderActionId() uint16 {
	return 279
}

func (*MatchEvent) OrderActionSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) OrderActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrderActionSinceVersion()
}

func (*MatchEvent) OrderActionDeprecated() uint16 {
	return 0
}

func (*MatchEvent) OrderActionMetaAttribute(meta int) string {
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

func (*MatchEvent) MakerSideId() uint16 {
	return 54
}

func (*MatchEvent) MakerSideSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) MakerSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MakerSideSinceVersion()
}

func (*MatchEvent) MakerSideDeprecated() uint16 {
	return 0
}

func (*MatchEvent) MakerSideMetaAttribute(meta int) string {
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

func (*MatchEvent) PriceId() uint16 {
	return 270
}

func (*MatchEvent) PriceSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceSinceVersion()
}

func (*MatchEvent) PriceDeprecated() uint16 {
	return 0
}

func (*MatchEvent) PriceMetaAttribute(meta int) string {
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

func (*MatchEvent) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*MatchEvent) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*MatchEvent) PriceNullValue() float64 {
	return math.NaN()
}

func (*MatchEvent) SizeId() uint16 {
	return 271
}

func (*MatchEvent) SizeSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) SizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SizeSinceVersion()
}

func (*MatchEvent) SizeDeprecated() uint16 {
	return 0
}

func (*MatchEvent) SizeMetaAttribute(meta int) string {
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

func (*MatchEvent) SizeMinValue() float64 {
	return -math.MaxFloat64
}

func (*MatchEvent) SizeMaxValue() float64 {
	return math.MaxFloat64
}

func (*MatchEvent) SizeNullValue() float64 {
	return math.NaN()
}

func (*MatchEvent) SecuritySequenceId() uint16 {
	return 83
}

func (*MatchEvent) SecuritySequenceSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) SecuritySequenceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SecuritySequenceSinceVersion()
}

func (*MatchEvent) SecuritySequenceDeprecated() uint16 {
	return 0
}

func (*MatchEvent) SecuritySequenceMetaAttribute(meta int) string {
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

func (*MatchEvent) SecuritySequenceMinValue() uint64 {
	return 0
}

func (*MatchEvent) SecuritySequenceMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MatchEvent) SecuritySequenceNullValue() uint64 {
	return math.MaxUint64
}

func (*MatchEvent) MakerOrderIDId() uint16 {
	return 37
}

func (*MatchEvent) MakerOrderIDSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) MakerOrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MakerOrderIDSinceVersion()
}

func (*MatchEvent) MakerOrderIDDeprecated() uint16 {
	return 0
}

func (*MatchEvent) MakerOrderIDMetaAttribute(meta int) string {
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

func (*MatchEvent) MakerOrderIDMinValue() byte {
	return byte(32)
}

func (*MatchEvent) MakerOrderIDMaxValue() byte {
	return byte(126)
}

func (*MatchEvent) MakerOrderIDNullValue() byte {
	return 0
}

func (m *MatchEvent) MakerOrderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*MatchEvent) TradeIDId() uint16 {
	return 76
}

func (*MatchEvent) TradeIDSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) TradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeIDSinceVersion()
}

func (*MatchEvent) TradeIDDeprecated() uint16 {
	return 0
}

func (*MatchEvent) TradeIDMetaAttribute(meta int) string {
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

func (*MatchEvent) TradeIDMinValue() uint64 {
	return 0
}

func (*MatchEvent) TradeIDMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MatchEvent) TradeIDNullValue() uint64 {
	return math.MaxUint64
}

func (*MatchEvent) TakerOrderIDId() uint16 {
	return 11
}

func (*MatchEvent) TakerOrderIDSinceVersion() uint16 {
	return 0
}

func (m *MatchEvent) TakerOrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TakerOrderIDSinceVersion()
}

func (*MatchEvent) TakerOrderIDDeprecated() uint16 {
	return 0
}

func (*MatchEvent) TakerOrderIDMetaAttribute(meta int) string {
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

func (*MatchEvent) TakerOrderIDMinValue() byte {
	return byte(32)
}

func (*MatchEvent) TakerOrderIDMaxValue() byte {
	return byte(126)
}

func (*MatchEvent) TakerOrderIDNullValue() byte {
	return 0
}

func (m *MatchEvent) TakerOrderIDCharacterEncoding() string {
	return "US-ASCII"
}
