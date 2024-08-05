package metadata

import (
	"fmt"

	"math"
	"strings"

	substrateTypes "github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/gobeam/stringy"
	"github.com/streamingfast/logging"
	"github.com/streamingfast/substreams-gear/protobuf"
	"github.com/streamingfast/substreams-gear/types"
	"go.uber.org/zap"
)

type Converter struct {
	logger        *zap.Logger
	tracer        logging.Tracer
	specVersion   string
	TypeConverter *TypeConverter
	metadata      *substrateTypes.Metadata
}

func NewMetadataConverter(metadata *substrateTypes.Metadata, specVersion string, logger *zap.Logger, tracer logging.Tracer) *Converter {
	return &Converter{
		specVersion: specVersion,
		logger:      logger,
		tracer:      tracer,
		metadata:    metadata,
	}
}

func (c *Converter) FetchMessages() []*protobuf.Message {
	outputs := make([]*protobuf.Message, 0)
	for _, seenMsg := range c.TypeConverter.messages {
		if seenMsg != nil {
			outputs = append(outputs, seenMsg)
		}
	}
	return outputs
}

func (c *Converter) FetchMetadata() *substrateTypes.Metadata {
	return c.metadata
}

func (c *Converter) Convert(knownMessage map[string]*protobuf.Message) error {
	switch c.metadata.Version {
	case 14:
		c.TypeConverter = &TypeConverter{
			specVersion:      c.specVersion,
			messages:         knownMessage,
			allMetadataTypes: []substrateTypes.PortableTypeV14{},
		}
		return c.TypeConverter.convertTypesFromv14(c.metadata.AsMetadataV14)
	default:
		fmt.Println("Unsupported metadata version", c.metadata.Version)
	}

	return nil
}

type TypeConverter struct {
	specVersion      string
	messages         map[string]*protobuf.Message
	allMetadataTypes []substrateTypes.PortableTypeV14
}

func (c *TypeConverter) convertTypesFromv14(metadata substrateTypes.MetadataV14) error {
	allMetadataTypes := metadata.Lookup.Types
	c.allMetadataTypes = allMetadataTypes

	for _, pallet := range metadata.Pallets {
		palletMessage := &protobuf.Message{
			SpecVersion: c.specVersion,
			Pallet:      "",
			Name:        string(pallet.Name) + "_pallet",
			LookupID:    pallet.Calls.Type.Int64(),
		}

		if pallet.HasCalls {
			callIdx := pallet.Calls.Type.Int64()
			palletName := string(pallet.Name)
			calls := c.ProcessPalletCalls(callIdx, palletName)

			oneOfCall := &protobuf.OneOfField{
				Name: "call",
			}

			for i, call := range calls {
				t := protobuf.BasicField{
					Pallet:      call.Pallet,
					Type:        call.Name,
					Name:        call.Name,
					LookupID:    call.LookupID,
					VariantByte: int64(i),
				}
				oneOfCall.Types = append(oneOfCall.Types, &t)
			}

			palletMessage.Fields = append(palletMessage.Fields, oneOfCall)
		}

		if _, found := c.messages[palletMessage.Hash()]; !found {
			c.messages[palletMessage.Hash()] = palletMessage
		}

		if pallet.HasEvents {
			idx := pallet.Events.Type.Int64()
			variants := c.allMetadataTypes[idx]

			events := variants.Type.Def.Variant
			for i, variant := range events.Variants {
				palletName := string(pallet.Name)
				callName := string(variant.Name)
				messageName := stringy.New(palletName).PascalCase().Get() + "_" + stringy.New(callName).PascalCase().Get() + "_Event"
				message := &protobuf.Message{
					SpecVersion: c.specVersion,
					Name:        messageName,
					IsEvent:     true,
				}

				c.ProcessPalletEvents(int64(i), palletName)
				if _, found := c.messages[palletMessage.Hash()]; !found {
					c.messages[message.Hash()] = message
				}
			}
		}
	}

	return nil
}

func (c *TypeConverter) ProcessPalletCalls(callIdx int64, palletName string) []protobuf.Message {
	var out []protobuf.Message
	variants := c.allMetadataTypes[callIdx]

	calls := variants.Type.Def.Variant
	for _, variant := range calls.Variants {
		callName := string(variant.Name)
		message := &protobuf.Message{
			SpecVersion: c.specVersion,
			Pallet:      palletName,
			Name:        callName + "_call",
			LookupID:    math.MaxInt64,
			IsCall:      true,
		}

		for _, f := range variant.Fields {
			fieldName := string(f.Name)
			field := c.ProcessField(f, palletName, callName, fieldName)
			if field != nil {
				message.Fields = append(message.Fields, field)
			}
		}

		if _, found := c.messages[message.Hash()]; !found {
			c.messages[message.Hash()] = message
		}
		out = append(out, *message)
	}
	return out
}

func (c *TypeConverter) ProcessPalletEvents(callIdx int64, palletName string) {
	variants := c.allMetadataTypes[callIdx]

	events := variants.Type.Def.Variant
	for _, variant := range events.Variants {
		eventName := string(variant.Name)
		message := &protobuf.Message{
			SpecVersion: c.specVersion,
			Pallet:      palletName,
			Name:        eventName + "_Event",
			LookupID:    math.MaxInt64,
			IsEvent:     true,
		}

		for i, f := range variant.Fields {
			fieldName := string(f.Name)
			if !f.HasName {
				fieldName = fmt.Sprintf("value%d", i)
			}

			field := c.ProcessField(f, palletName, eventName, fieldName)
			if field != nil {
				message.Fields = append(message.Fields, field)
			}
		}

		if _, found := c.messages[message.Hash()]; !found {
			c.messages[message.Hash()] = message
		}
	}
}

func (c *TypeConverter) ProcessField(f substrateTypes.Si1Field, palletName string, callName string, fieldName string) protobuf.Field {
	idx := f.Type.Int64()
	ttype := c.allMetadataTypes[idx]

	return c.FieldForType(ttype, palletName, callName, fieldName)
}

func (c *TypeConverter) FieldForPrimitive(ttype substrateTypes.PortableTypeV14, fieldName string) *protobuf.BasicField {
	typeName := types.ConvertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).GetProtoFieldName()

	if typeName == "uint8" {
		typeName = "uint32"
	}

	if typeName == "int8" {
		typeName = "int32"
	}

	return &protobuf.BasicField{
		Pallet:    "",
		Name:      fieldName,
		Type:      typeName,
		Primitive: true,
		LookupID:  ttype.ID.Int64(),
	}
}

func (c *TypeConverter) FieldForSequence(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	lookupId := ttype.Type.Def.Sequence.Type.Int64()
	lookupType := c.allMetadataTypes[lookupId]
	typeName := c.ExtractTypeName(lookupType, palletName, callName, fieldName)

	if typeName == "uint8" || typeName == "int8" {
		return &protobuf.BasicField{
			Pallet:    palletNameFromPath(lookupType.Type.Path, palletName, lookupType.Type.Def.IsPrimitive),
			Type:      "bytes",
			Name:      fieldName,
			Primitive: lookupType.Type.Def.IsPrimitive,
			LookupID:  lookupId,
		}
	}
	f := &protobuf.RepeatedField{
		Pallet:    palletNameFromPath(lookupType.Type.Path, palletName, lookupType.Type.Def.IsPrimitive),
		Type:      typeName,
		Name:      fieldName,
		Primitive: lookupType.Type.Def.IsPrimitive,
		LookupID:  lookupId,
	}

	return f
}

func (c *TypeConverter) FieldForArray(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	lookupId := ttype.Type.Def.Array.Type.Int64()
	lookupType := c.allMetadataTypes[lookupId]
	typeName := c.ExtractTypeName(lookupType, palletName, callName, fieldName)

	if typeName == "uint8" || typeName == "int8" {
		return &protobuf.BasicField{
			Pallet:    palletNameFromPath(ttype.Type.Path, palletName, lookupType.Type.Def.IsPrimitive),
			Type:      "bytes",
			Name:      fieldName,
			Primitive: lookupType.Type.Def.IsPrimitive,
			LookupID:  lookupId,
		}
	}

	return &protobuf.RepeatedField{
		Pallet:    palletNameFromPath(ttype.Type.Path, palletName, lookupType.Type.Def.IsPrimitive),
		Name:      fieldName,
		Type:      typeName,
		Primitive: lookupType.Type.Def.IsPrimitive,
		LookupID:  lookupId,
	}
}

func (c *TypeConverter) FieldForCompact(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string, primitive bool) *protobuf.BasicField {
	childType := c.allMetadataTypes[ttype.Type.Def.Compact.Type.Int64()]
	typeName := c.ExtractTypeName(childType, palletName, callName, fieldName)

	return &protobuf.BasicField{
		Pallet:    palletNameFromPath(childType.Type.Path, palletName, primitive),
		Name:      fieldName,
		Type:      typeName,
		LookupID:  childType.ID.Int64(),
		Compact:   true,
		Primitive: primitive,
	}
}

func (c *TypeConverter) FieldForComposite(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	name := c.ExtractTypeName(ttype, palletName, callName, fieldName)

	return &protobuf.BasicField{
		Pallet:   palletNameFromPath(ttype.Type.Path, palletName, false),
		Name:     fieldName,
		Type:     name,
		LookupID: ttype.ID.Int64(),
	}
}

func (c *TypeConverter) FieldForType(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	recursiveIdx := int64(0)

	switch c.specVersion {
	case "1420", "1410", "1400":
		recursiveIdx = 65
	case "1310", "1300", "1210", "1200", "1110", "1050", "1040", "1030":
		recursiveIdx = 64
	case "1020", "1010", "1000", "350", "340", "330", "320", "310":
		recursiveIdx = 63
	case "210", "140", "130", "120", "100":
		recursiveIdx = 60
	default:
		panic("unsupported spec version")
	}

	if ttype.ID.Int64() == recursiveIdx {
		return c.FieldForRecursiveType(ttype, palletName, callName, fieldName)
	}

	var field protobuf.Field
	var optional bool
	if len(ttype.Type.Path) != 0 {
		isOptional := ttype.Type.Path[len(ttype.Type.Path)-1]
		optional = isOptional == "Option"
	}

	if optional {
		lookupId := ttype.Type.Def.Variant.Variants[1].Fields[0].Type.Int64()
		ttype = c.allMetadataTypes[lookupId]
	}

	switch {
	case ttype.Type.Def.IsPrimitive:
		field = c.FieldForPrimitive(ttype, fieldName)
		if optional {
			field.SetOptional()
		}
		return field

	case ttype.Type.Def.IsSequence:
		field = c.FieldForSequence(ttype, palletName, callName, fieldName)
		lookupId := ttype.Type.Def.Sequence.Type.Int64()
		t := c.allMetadataTypes[lookupId]
		if t.Type.Def.IsTuple {
			ttype = t
		}

	case ttype.Type.Def.IsArray:
		field = c.FieldForArray(ttype, palletName, callName, fieldName)
		lookupId := ttype.Type.Def.Array.Type.Int64()
		t := c.allMetadataTypes[lookupId]
		if t.Type.Def.IsTuple {
			ttype = t
		}

	case ttype.Type.Def.IsTuple:
		field = c.FieldForTuple(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}

	case ttype.Type.Def.IsVariant:
		field = c.FieldForVariant(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}

	case ttype.Type.Def.IsCompact:
		idx := ttype.Type.Def.Compact.Type.Int64()
		childType := c.allMetadataTypes[idx]
		field = c.FieldForCompact(ttype, palletName, callName, fieldName, childType.Type.Def.IsPrimitive)
		if optional {
			field.SetOptional()
		}

	case ttype.Type.Def.IsComposite:
		field = c.FieldForComposite(ttype, palletName, callName, fieldName)
		if optional {
			field.SetOptional()
		}
	}

	if !ttype.Type.Def.IsVariant && !field.IsPrimitive() {
		c.MessageForType(field.GetType(), ttype, palletName, callName, fieldName)
	}

	return field
}

func (c *TypeConverter) ExtractTypeName(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) string {
	typeName := ""
	if ttype.Type.Def.IsPrimitive {
		return types.ConvertPrimitiveType(ttype.Type.Def.Primitive.Si0TypeDefPrimitive).ToGoType()
	}

	if ttype.Type.Def.IsTuple {
		typeName = c.ExtractTypeNameFromTuple(ttype.Type.Def.Tuple, palletName, callName, fieldName)
	}

	if ttype.Type.Def.IsComposite {
		typeName = c.ExtractTypeNameFromPath(ttype.Type.Path)
	}

	if ttype.Type.Def.IsVariant {
		typeName = c.ExtractTypeNameFromPath(ttype.Type.Path)
	}

	if ttype.Type.Def.IsSequence || ttype.Type.Def.IsArray {
		typeName = fmt.Sprintf("%s_%s_list", palletName, fieldName)
	}

	if ttype.Type.Def.IsCompact {
		lookupId := ttype.Type.Def.Compact.Type.Int64()
		lookupType := c.allMetadataTypes[lookupId]
		typeName = c.ExtractTypeName(lookupType, palletName, callName, fieldName)
	}

	return typeName
}

func (c *TypeConverter) MessageForType(typeName string, ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) {
	msg := &protobuf.Message{
		SpecVersion: c.specVersion,
		Name:        typeName,
		LookupID:    ttype.ID.Int64(),
	}

	if ttype.Type.Def.IsTuple {
		if typeName != "Tuple_Null" {

			for i, idx := range ttype.Type.Def.Tuple {
				lookupId := idx.Int64()
				t := c.allMetadataTypes[lookupId]
				f := c.FieldForType(t, palletName, callName, fmt.Sprintf("value%d", i))
				msg.Fields = append(msg.Fields, f)
			}
		}
		msg.Pallet = palletNameFromPath(ttype.Type.Path, palletName, false)
	}

	if ttype.Type.Def.IsComposite {
		for i, f := range ttype.Type.Def.Composite.Fields {
			lookup := f.Type.Int64()
			lookupType := c.allMetadataTypes[lookup]
			fieldName := string(f.Name)
			if !f.HasName {
				fieldName = fmt.Sprintf("value%d", i)
			}
			field := c.FieldForType(lookupType, palletName, callName, fieldName)
			msg.Fields = append(msg.Fields, field)
		}
		msg.Pallet = palletNameFromPath(ttype.Type.Path, palletName, false)
	}

	if ttype.Type.Def.IsVariant {
		field := &protobuf.OneOfField{
			Pallet:   palletName,
			Name:     typeName,
			LookupID: ttype.ID.Int64(),
		}
		for _, v := range ttype.Type.Def.Variant.Variants {
			field.Types = append(field.Types, &protobuf.BasicField{
				Pallet:   palletName,
				Name:     fmt.Sprintf("%s_%s", palletName, string(v.Name)),
				Type:     string(v.Name),
				LookupID: math.MaxInt64,
			})
		}
		msg.Pallet = palletNameFromPath(ttype.Type.Path, palletName, false)

		msg.Fields = append(msg.Fields, field)
	}

	if ttype.Type.Def.IsSequence {
		lookupId := ttype.Type.Def.Sequence.Type.Int64()
		childType := c.allMetadataTypes[lookupId]

		f := c.FieldForType(childType, palletName, callName, fieldName)

		msg.Pallet = palletNameFromPath(childType.Type.Path, palletName, childType.Type.Def.IsPrimitive)
		msg.Fields = append(msg.Fields, f)
	}

	if ttype.Type.Def.IsArray {
		lookupId := ttype.Type.Def.Array.Type.Int64()
		childType := c.allMetadataTypes[lookupId]

		f := c.FieldForType(childType, palletName, callName, fieldName)

		msg.Pallet = palletNameFromPath(childType.Type.Path, palletName, childType.Type.Def.IsPrimitive)
		msg.Fields = append(msg.Fields, f)
	}

	if ttype.Type.Def.IsCompact {
		lookupId := ttype.Type.Def.Compact.Type.Int64()
		childType := c.allMetadataTypes[lookupId]
		field := c.FieldForType(childType, msg.Pallet, callName, "value")

		msg.Pallet = palletNameFromPath(childType.Type.Path, palletName, childType.Type.Def.IsPrimitive)
		msg.Fields = append(msg.Fields, field)
	}

	if _, found := c.messages[msg.Hash()]; !found {
		//if !ttype.Type.Def.IsPrimitive {
		c.messages[msg.Hash()] = msg
		//}
	}

	return
}

func (c *TypeConverter) FieldForRecursiveType(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	field := &protobuf.OneOfField{
		Pallet:   palletNameFromPath(ttype.Type.Path, palletName, false),
		Name:     fieldName,
		LookupID: ttype.ID.Int64(),
	}

	palletRefs := ttype.Type.Def.Variant.Variants
	for _, ref := range palletRefs {
		//refID := ref.Fields[0].Type.Int64()
		pName := string(ref.Name)
		//pType := c.allMetadataTypes[refID]
		//calls := pType.Type.Def.Variant.Variants

		field.Types = append(field.Types, &protobuf.BasicField{
			Pallet:      "",
			Name:        pName,
			LookupID:    math.MaxInt64,
			Type:        pName + "_Pallet",
			VariantByte: int64(ref.Index),
		})
	}

	return field
}

func (c *TypeConverter) MessageForVariantTypes(name string, variant substrateTypes.Si1Variant, palletName string, callName string, fieldName string) *protobuf.Message {
	msg := &protobuf.Message{
		SpecVersion: c.specVersion,
		Pallet:      palletName,
		Name:        name,
		LookupID:    math.MaxInt64,
	}

	for i, f := range variant.Fields {
		idx := f.Type.Int64()
		fieldType := c.allMetadataTypes[idx]
		fn := string(f.Name)
		if !f.HasName {
			fn = fmt.Sprintf("value%d", i)
		}

		pName := palletNameFromPath(fieldType.Type.Path, palletName, false)
		field := c.FieldForType(fieldType, pName, callName, fn)
		msg.Fields = append(msg.Fields, field)
	}

	return msg
}

func (c *TypeConverter) ExtractTypeNameFromTuple(tuple substrateTypes.Si1TypeDefTuple, palletName string, callName string, fieldName string) string {
	var name strings.Builder
	name.WriteString("Tuple_")
	if len(tuple) == 0 {
		name.WriteString("Null")
	}

	for _, item := range tuple {
		lookupId := item.Int64()
		ttype := c.allMetadataTypes[lookupId]
		name.WriteString(c.ExtractTypeName(ttype, palletName, callName, fieldName))
	}

	return name.String()
}

func (c *TypeConverter) ExtractTypeNameFromPath(path substrateTypes.Si1Path) string {
	parts := make([]string, 0)
	for _, p := range path {
		parts = append(parts, string(p))
	}
	return strings.Join(parts, "_")
}

func (c *TypeConverter) FieldForTuple(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) *protobuf.BasicField {
	if fieldName == "" {
		panic("field name is required")
	}

	field := &protobuf.BasicField{
		Pallet:   palletName,
		Name:     fieldName,
		Type:     c.ExtractTypeNameFromTuple(ttype.Type.Def.Tuple, palletName, callName, fieldName),
		LookupID: ttype.ID.Int64(),
	}

	return field
}

func (c *TypeConverter) FieldForVariant(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) protobuf.Field {
	//palletName = palletNameFromPath(ttype.Type.Path, palletName, false)
	oneOf := &protobuf.OneOfField{
		Name:     fieldName,
		Pallet:   palletName,
		LookupID: ttype.ID.Int64(),
	}

	msg := &protobuf.Message{
		SpecVersion: c.specVersion,
		Pallet:      palletName,
		Name:        fieldName,
		Fields:      []protobuf.Field{oneOf},
		LookupID:    ttype.ID.Int64(),
	}

	for _, v := range ttype.Type.Def.Variant.Variants {
		typeName := string(v.Name)

		oneOf.Types = append(oneOf.Types, &protobuf.BasicField{
			Pallet:      palletName,
			Name:        string(v.Name),
			Type:        string(v.Name),
			LookupID:    math.MaxInt64,
			VariantByte: int64(v.Index),
		})

		m := c.MessageForVariantTypes(typeName, v, palletName, callName, string(v.Name))
		if _, found := c.messages[m.Hash()]; !found {
			c.messages[m.Hash()] = m
		}
	}

	if _, found := c.messages[msg.Hash()]; !found {
		c.messages[msg.Hash()] = msg
	}

	return oneOf
}

func palletNameFromPath(path substrateTypes.Si1Path, defaultt string, primitive bool) string {
	if primitive {
		return ""
	}
	if len(path) == 0 {
		return defaultt
	}

	pallet := string(path[0])
	if strings.HasPrefix(pallet, "pallet_") {
		return strings.TrimPrefix(pallet, "pallet_")
	}

	if strings.Contains(strings.ToLower(pallet), strings.ToLower(defaultt)) {
		return defaultt
	}

	return ""
}

func (c *TypeConverter) ProcessOptionalType(ttype substrateTypes.PortableTypeV14, palletName string, callName string, fieldName string) string {
	variant := ttype.Type.Def.Variant
	some := variant.Variants[1]
	someField := some.Fields[0]
	tttype := c.allMetadataTypes[someField.Type.Int64()]

	return c.ExtractTypeName(tttype, palletName, callName, fieldName)
}
