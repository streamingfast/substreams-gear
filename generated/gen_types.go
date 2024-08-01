package generated

import (
    "fmt"
    "reflect"

    "github.com/centrifuge/go-substrate-rpc-client/v4/registry"
    "github.com/centrifuge/go-substrate-rpc-client/v4/types"
    pbgearextrinsic "github.com/streamingfast/substreams-gear/pb/sf/substreams/gear/type/v1"
    pbgear "github.com/streamingfast/substreams-gear/pb/sf/substreams/gear/type/v1"
)
func To_Active(in any) *pbgear.Active {
    out := &pbgear.Active{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Expiration
    out.Expiration = To_Optional_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Any(in any) *pbgear.Any {
    out := &pbgear.Any{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Arithmetic(in any) *pbgear.Arithmetic {
    out := &pbgear.Arithmetic{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Arithmetic_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Arithmetic_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Arithmetic_Value0Underflow{
        Value0Underflow: To_Underflow(param),
        }
    case 1:
        return &pbgear.Arithmetic_Value0Overflow{
        Value0Overflow: To_Overflow(param),
        }
    case 2:
        return &pbgear.Arithmetic_Value0DivisionByZero{
        Value0DivisionByZero: To_DivisionByZero(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_AuthorityDiscoveryPallet(in any) *pbgear.AuthorityDiscoveryPallet {
    out := &pbgear.AuthorityDiscoveryPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_AuthorshipPallet(in any) *pbgear.AuthorshipPallet {
    out := &pbgear.AuthorshipPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Auto(in any) *pbgear.Auto {
    out := &pbgear.Auto{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BTreeMap(in any) *pbgear.BTreeMap {
    out := &pbgear.BTreeMap{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Value0
    out.Value0 = To_Repeated_BTreeMap_value0(v.ValueAt(0))

    return out //from message
}
    

func To_Repeated_BTreeMap_value0(in any) []*pbgear.TupleGprimitivesMessageIdgearCommonEventDispatchStatus {
    items := in.([]interface{})

    var out []*pbgear.TupleGprimitivesMessageIdgearCommonEventDispatchStatus
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_TupleGprimitivesMessageIdgearCommonEventDispatchStatus(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_BTreeSet(in any) *pbgear.BTreeSet {
    out := &pbgear.BTreeSet{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Value0
    out.Value0 = To_Repeated_BTreeSet_value0(v.ValueAt(0))

    return out //from message
}
    

func To_Repeated_BTreeSet_value0(in any) []*pbgear.GprimitivesActorId {
    items := in.([]interface{})

    var out []*pbgear.GprimitivesActorId
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_GprimitivesActorId(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_BabePallet(in any) *pbgear.BabePallet {
    out := &pbgear.BabePallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_BabePallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_BabePallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BabePallet_CallReportEquivocationCall{
        CallReportEquivocationCall: To_Babe_ReportEquivocationCall(param),
        }
    case 1:
        return &pbgear.BabePallet_CallReportEquivocationUnsignedCall{
        CallReportEquivocationUnsignedCall: To_Babe_ReportEquivocationUnsignedCall(param),
        }
    case 2:
        return &pbgear.BabePallet_CallPlanConfigChangeCall{
        CallPlanConfigChangeCall: To_Babe_PlanConfigChangeCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Babe_AllowedSlots(in any) *pbgear.Babe_AllowedSlots {
    out := &pbgear.Babe_AllowedSlots{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field AllowedSlots
    v0 := To_OneOf_Babe_AllowedSlots_allowed_slots(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("AllowedSlots").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Babe_AllowedSlots_allowed_slots (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Babe_AllowedSlots_AllowedSlotsPrimarySlots{
        AllowedSlotsPrimarySlots: To_Babe_PrimarySlots(param),
        }
    case 1:
        return &pbgear.Babe_AllowedSlots_AllowedSlotsPrimaryAndSecondaryPlainSlots{
        AllowedSlotsPrimaryAndSecondaryPlainSlots: To_Babe_PrimaryAndSecondaryPlainSlots(param),
        }
    case 2:
        return &pbgear.Babe_AllowedSlots_AllowedSlotsPrimaryAndSecondaryVrfSlots{
        AllowedSlotsPrimaryAndSecondaryVrfSlots: To_Babe_PrimaryAndSecondaryVrfSlots(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Babe_BabeTrieNodesList(in any) *pbgear.Babe_BabeTrieNodesList {
    out := &pbgear.Babe_BabeTrieNodesList{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field TrieNodes
    out.TrieNodes = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Babe_Config(in any) *pbgear.Babe_Config {
    out := &pbgear.Babe_Config{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Config
    v0 := To_OneOf_Babe_Config_config(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Config").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Babe_Config_config (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 1:
        return &pbgear.Babe_Config_ConfigV1{
        ConfigV1: To_Babe_V1(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Babe_Consensus(in any) *pbgear.Babe_Consensus {
    out := &pbgear.Babe_Consensus{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Babe_Logs(in any) *pbgear.Babe_Logs {
    out := &pbgear.Babe_Logs{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Logs
    v0 := To_OneOf_Babe_Logs_logs(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Logs").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Babe_Logs_logs (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 6:
        return &pbgear.Babe_Logs_LogsPreRuntime{
        LogsPreRuntime: To_Babe_PreRuntime(param),
        }
    case 4:
        return &pbgear.Babe_Logs_LogsConsensus{
        LogsConsensus: To_Babe_Consensus(param),
        }
    case 5:
        return &pbgear.Babe_Logs_LogsSeal{
        LogsSeal: To_Babe_Seal(param),
        }
    case 0:
        return &pbgear.Babe_Logs_LogsOther{
        LogsOther: To_Babe_Other(param),
        }
    case 8:
        return &pbgear.Babe_Logs_LogsRuntimeEnvironmentUpdated{
        LogsRuntimeEnvironmentUpdated: To_Babe_RuntimeEnvironmentUpdated(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Babe_Other(in any) *pbgear.Babe_Other {
    out := &pbgear.Babe_Other{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Babe_PlanConfigChangeCall(in any) *pbgear.Babe_PlanConfigChangeCall {
    out := &pbgear.Babe_PlanConfigChangeCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Config
    v0 := To_OneOf_Babe_PlanConfigChangeCall_config(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Config").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_Babe_PlanConfigChangeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Babe_PlanConfigChangeCall(in)
    out := &pbgearextrinsic.RuntimeCall_BabePlanConfigChangeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BabePlanConfigChangeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Babe_PlanConfigChangeCall_config (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 1:
        return &pbgear.Babe_PlanConfigChangeCall_ConfigV1{
        ConfigV1: To_Babe_V1(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Babe_PreRuntime(in any) *pbgear.Babe_PreRuntime {
    out := &pbgear.Babe_PreRuntime{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Babe_PrimaryAndSecondaryPlainSlots(in any) *pbgear.Babe_PrimaryAndSecondaryPlainSlots {
    out := &pbgear.Babe_PrimaryAndSecondaryPlainSlots{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Babe_PrimaryAndSecondaryVrfSlots(in any) *pbgear.Babe_PrimaryAndSecondaryVrfSlots {
    out := &pbgear.Babe_PrimaryAndSecondaryVrfSlots{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Babe_PrimarySlots(in any) *pbgear.Babe_PrimarySlots {
    out := &pbgear.Babe_PrimarySlots{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Babe_ReportEquivocationCall(in any) *pbgear.Babe_ReportEquivocationCall {
    out := &pbgear.Babe_ReportEquivocationCall{}
    v := in.(registry.Valuable)
    _ = v

    // field EquivocationProof To_SpConsensusSlotsEquivocationProof(w)
    out.EquivocationProof = To_SpConsensusSlotsEquivocationProof(v.ValueAt(0))
    // field KeyOwnerProof To_SpSessionMembershipProof(w)
    out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

    return out //from message
}
func To_Babe_ReportEquivocationCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Babe_ReportEquivocationCall(in)
    out := &pbgearextrinsic.RuntimeCall_BabeReportEquivocationCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BabeReportEquivocationCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_Babe_ReportEquivocationUnsignedCall(in any) *pbgear.Babe_ReportEquivocationUnsignedCall {
    out := &pbgear.Babe_ReportEquivocationUnsignedCall{}
    v := in.(registry.Valuable)
    _ = v

    // field EquivocationProof To_SpConsensusSlotsEquivocationProof(w)
    out.EquivocationProof = To_SpConsensusSlotsEquivocationProof(v.ValueAt(0))
    // field KeyOwnerProof To_SpSessionMembershipProof(w)
    out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

    return out //from message
}
func To_Babe_ReportEquivocationUnsignedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Babe_ReportEquivocationUnsignedCall(in)
    out := &pbgearextrinsic.RuntimeCall_BabeReportEquivocationUnsignedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BabeReportEquivocationUnsignedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_Babe_RuntimeEnvironmentUpdated(in any) *pbgear.Babe_RuntimeEnvironmentUpdated {
    out := &pbgear.Babe_RuntimeEnvironmentUpdated{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Babe_Seal(in any) *pbgear.Babe_Seal {
    out := &pbgear.Babe_Seal{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Babe_SpConsensusBabeAppPublic(in any) *pbgear.Babe_SpConsensusBabeAppPublic {
    out := &pbgear.Babe_SpConsensusBabeAppPublic{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreSr25519Public(w)
    out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

    return out //from message
}
    

func To_Babe_TupleUint64Uint64(in any) *pbgear.Babe_TupleUint64Uint64 {
    out := &pbgear.Babe_TupleUint64Uint64{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint64(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_uint64(v.ValueAt(1))

    return out //from message
}
    
func To_Babe_V1(in any) *pbgear.Babe_V1 {
    out := &pbgear.Babe_V1{}
    v := in.(registry.Valuable)
    _ = v

    // field C To_Babe_TupleUint64Uint64(w)
    out.C = To_Babe_TupleUint64Uint64(v.ValueAt(0))
    // oneOf field AllowedSlots
    v1 := To_OneOf_Babe_V1_allowed_slots(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("AllowedSlots").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_Babe_V1_allowed_slots (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Babe_V1_AllowedSlotsPrimarySlots{
        AllowedSlotsPrimarySlots: To_Babe_PrimarySlots(param),
        }
    case 1:
        return &pbgear.Babe_V1_AllowedSlotsPrimaryAndSecondaryPlainSlots{
        AllowedSlotsPrimaryAndSecondaryPlainSlots: To_Babe_PrimaryAndSecondaryPlainSlots(param),
        }
    case 2:
        return &pbgear.Babe_V1_AllowedSlotsPrimaryAndSecondaryVrfSlots{
        AllowedSlotsPrimaryAndSecondaryVrfSlots: To_Babe_PrimaryAndSecondaryVrfSlots(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BackendError(in any) *pbgear.BackendError {
    out := &pbgear.BackendError{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BadOrigin(in any) *pbgear.BadOrigin {
    out := &pbgear.BadOrigin{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BagsListPallet(in any) *pbgear.BagsListPallet {
    out := &pbgear.BagsListPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_BagsListPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_BagsListPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsListPallet_CallRebagCall{
        CallRebagCall: To_BagsList_RebagCall(param),
        }
    case 1:
        return &pbgear.BagsListPallet_CallPutInFrontOfCall{
        CallPutInFrontOfCall: To_BagsList_PutInFrontOfCall(param),
        }
    case 2:
        return &pbgear.BagsListPallet_CallPutInFrontOfOtherCall{
        CallPutInFrontOfOtherCall: To_BagsList_PutInFrontOfOtherCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BagsListRebaggedEvent(in any) *pbgear.BagsListRebaggedEvent {
    out := &pbgear.BagsListRebaggedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BagsListScoreUpdatedEvent(in any) *pbgear.BagsListScoreUpdatedEvent {
    out := &pbgear.BagsListScoreUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BagsList_Address20(in any) *pbgear.BagsList_Address20 {
    out := &pbgear.BagsList_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_BagsList_Address32(in any) *pbgear.BagsList_Address32 {
    out := &pbgear.BagsList_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_BagsList_Dislocated(in any) *pbgear.BagsList_Dislocated {
    out := &pbgear.BagsList_Dislocated{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Dislocated
    v0 := To_OneOf_BagsList_Dislocated_dislocated(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Dislocated").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_BagsList_Dislocated_dislocated (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsList_Dislocated_DislocatedId{
        DislocatedId: To_BagsList_Id(param),
        }
    case 1:
        return &pbgear.BagsList_Dislocated_DislocatedIndex{
        DislocatedIndex: To_BagsList_Index(param),
        }
    case 2:
        return &pbgear.BagsList_Dislocated_DislocatedRaw{
        DislocatedRaw: To_BagsList_Raw(param),
        }
    case 3:
        return &pbgear.BagsList_Dislocated_DislocatedAddress32{
        DislocatedAddress32: To_BagsList_Address32(param),
        }
    case 4:
        return &pbgear.BagsList_Dislocated_DislocatedAddress20{
        DislocatedAddress20: To_BagsList_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BagsList_Heavier(in any) *pbgear.BagsList_Heavier {
    out := &pbgear.BagsList_Heavier{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Heavier
    v0 := To_OneOf_BagsList_Heavier_heavier(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Heavier").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_BagsList_Heavier_heavier (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsList_Heavier_HeavierId{
        HeavierId: To_BagsList_Id(param),
        }
    case 1:
        return &pbgear.BagsList_Heavier_HeavierIndex{
        HeavierIndex: To_BagsList_Index(param),
        }
    case 2:
        return &pbgear.BagsList_Heavier_HeavierRaw{
        HeavierRaw: To_BagsList_Raw(param),
        }
    case 3:
        return &pbgear.BagsList_Heavier_HeavierAddress32{
        HeavierAddress32: To_BagsList_Address32(param),
        }
    case 4:
        return &pbgear.BagsList_Heavier_HeavierAddress20{
        HeavierAddress20: To_BagsList_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BagsList_Id(in any) *pbgear.BagsList_Id {
    out := &pbgear.BagsList_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_BagsList_Index(in any) *pbgear.BagsList_Index {
    out := &pbgear.BagsList_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_BagsList_TupleNull(w)
    out.Value0 = To_BagsList_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_BagsList_Lighter(in any) *pbgear.BagsList_Lighter {
    out := &pbgear.BagsList_Lighter{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Lighter
    v0 := To_OneOf_BagsList_Lighter_lighter(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Lighter").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_BagsList_Lighter_lighter (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsList_Lighter_LighterId{
        LighterId: To_BagsList_Id(param),
        }
    case 1:
        return &pbgear.BagsList_Lighter_LighterIndex{
        LighterIndex: To_BagsList_Index(param),
        }
    case 2:
        return &pbgear.BagsList_Lighter_LighterRaw{
        LighterRaw: To_BagsList_Raw(param),
        }
    case 3:
        return &pbgear.BagsList_Lighter_LighterAddress32{
        LighterAddress32: To_BagsList_Address32(param),
        }
    case 4:
        return &pbgear.BagsList_Lighter_LighterAddress20{
        LighterAddress20: To_BagsList_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BagsList_PutInFrontOfCall(in any) *pbgear.BagsList_PutInFrontOfCall {
    out := &pbgear.BagsList_PutInFrontOfCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Lighter
    v0 := To_OneOf_BagsList_PutInFrontOfCall_lighter(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Lighter").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_BagsList_PutInFrontOfCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_BagsList_PutInFrontOfCall(in)
    out := &pbgearextrinsic.RuntimeCall_BagslistPutInFrontOfCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BagslistPutInFrontOfCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_BagsList_PutInFrontOfCall_lighter (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsList_PutInFrontOfCall_LighterId{
        LighterId: To_BagsList_Id(param),
        }
    case 1:
        return &pbgear.BagsList_PutInFrontOfCall_LighterIndex{
        LighterIndex: To_BagsList_Index(param),
        }
    case 2:
        return &pbgear.BagsList_PutInFrontOfCall_LighterRaw{
        LighterRaw: To_BagsList_Raw(param),
        }
    case 3:
        return &pbgear.BagsList_PutInFrontOfCall_LighterAddress32{
        LighterAddress32: To_BagsList_Address32(param),
        }
    case 4:
        return &pbgear.BagsList_PutInFrontOfCall_LighterAddress20{
        LighterAddress20: To_BagsList_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BagsList_PutInFrontOfOtherCall(in any) *pbgear.BagsList_PutInFrontOfOtherCall {
    out := &pbgear.BagsList_PutInFrontOfOtherCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Heavier
    v0 := To_OneOf_BagsList_PutInFrontOfOtherCall_heavier(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Heavier").Set(reflect.ValueOf(v0))
    // oneOf field Lighter
    v1 := To_OneOf_BagsList_PutInFrontOfOtherCall_lighter(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Lighter").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_BagsList_PutInFrontOfOtherCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_BagsList_PutInFrontOfOtherCall(in)
    out := &pbgearextrinsic.RuntimeCall_BagslistPutInFrontOfOtherCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BagslistPutInFrontOfOtherCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_BagsList_PutInFrontOfOtherCall_heavier (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierId{
        HeavierId: To_BagsList_Id(param),
        }
    case 1:
        return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierIndex{
        HeavierIndex: To_BagsList_Index(param),
        }
    case 2:
        return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierRaw{
        HeavierRaw: To_BagsList_Raw(param),
        }
    case 3:
        return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierAddress32{
        HeavierAddress32: To_BagsList_Address32(param),
        }
    case 4:
        return &pbgear.BagsList_PutInFrontOfOtherCall_HeavierAddress20{
        HeavierAddress20: To_BagsList_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_BagsList_PutInFrontOfOtherCall_lighter (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsList_PutInFrontOfOtherCall_LighterId{
        LighterId: To_BagsList_Id(param),
        }
    case 1:
        return &pbgear.BagsList_PutInFrontOfOtherCall_LighterIndex{
        LighterIndex: To_BagsList_Index(param),
        }
    case 2:
        return &pbgear.BagsList_PutInFrontOfOtherCall_LighterRaw{
        LighterRaw: To_BagsList_Raw(param),
        }
    case 3:
        return &pbgear.BagsList_PutInFrontOfOtherCall_LighterAddress32{
        LighterAddress32: To_BagsList_Address32(param),
        }
    case 4:
        return &pbgear.BagsList_PutInFrontOfOtherCall_LighterAddress20{
        LighterAddress20: To_BagsList_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BagsList_Raw(in any) *pbgear.BagsList_Raw {
    out := &pbgear.BagsList_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_BagsList_RebagCall(in any) *pbgear.BagsList_RebagCall {
    out := &pbgear.BagsList_RebagCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Dislocated
    v0 := To_OneOf_BagsList_RebagCall_dislocated(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Dislocated").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_BagsList_RebagCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_BagsList_RebagCall(in)
    out := &pbgearextrinsic.RuntimeCall_BagslistRebagCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BagslistRebagCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_BagsList_RebagCall_dislocated (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BagsList_RebagCall_DislocatedId{
        DislocatedId: To_BagsList_Id(param),
        }
    case 1:
        return &pbgear.BagsList_RebagCall_DislocatedIndex{
        DislocatedIndex: To_BagsList_Index(param),
        }
    case 2:
        return &pbgear.BagsList_RebagCall_DislocatedRaw{
        DislocatedRaw: To_BagsList_Raw(param),
        }
    case 3:
        return &pbgear.BagsList_RebagCall_DislocatedAddress32{
        DislocatedAddress32: To_BagsList_Address32(param),
        }
    case 4:
        return &pbgear.BagsList_RebagCall_DislocatedAddress20{
        DislocatedAddress20: To_BagsList_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BagsList_TupleNull(in any) *pbgear.BagsList_TupleNull {
    out := &pbgear.BagsList_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_BalancesBalanceSetEvent(in any) *pbgear.BalancesBalanceSetEvent {
    out := &pbgear.BalancesBalanceSetEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesBurnedEvent(in any) *pbgear.BalancesBurnedEvent {
    out := &pbgear.BalancesBurnedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesDepositEvent(in any) *pbgear.BalancesDepositEvent {
    out := &pbgear.BalancesDepositEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesDustLostEvent(in any) *pbgear.BalancesDustLostEvent {
    out := &pbgear.BalancesDustLostEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesEndowedEvent(in any) *pbgear.BalancesEndowedEvent {
    out := &pbgear.BalancesEndowedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesFrozenEvent(in any) *pbgear.BalancesFrozenEvent {
    out := &pbgear.BalancesFrozenEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesIssuedEvent(in any) *pbgear.BalancesIssuedEvent {
    out := &pbgear.BalancesIssuedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesLockedEvent(in any) *pbgear.BalancesLockedEvent {
    out := &pbgear.BalancesLockedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesMintedEvent(in any) *pbgear.BalancesMintedEvent {
    out := &pbgear.BalancesMintedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesPallet(in any) *pbgear.BalancesPallet {
    out := &pbgear.BalancesPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_BalancesPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_BalancesPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BalancesPallet_CallTransferAllowDeathCall{
        CallTransferAllowDeathCall: To_Balances_TransferAllowDeathCall(param),
        }
    case 1:
        return &pbgear.BalancesPallet_CallForceTransferCall{
        CallForceTransferCall: To_Balances_ForceTransferCall(param),
        }
    case 2:
        return &pbgear.BalancesPallet_CallTransferKeepAliveCall{
        CallTransferKeepAliveCall: To_Balances_TransferKeepAliveCall(param),
        }
    case 3:
        return &pbgear.BalancesPallet_CallTransferAllCall{
        CallTransferAllCall: To_Balances_TransferAllCall(param),
        }
    case 4:
        return &pbgear.BalancesPallet_CallForceUnreserveCall{
        CallForceUnreserveCall: To_Balances_ForceUnreserveCall(param),
        }
    case 5:
        return &pbgear.BalancesPallet_CallUpgradeAccountsCall{
        CallUpgradeAccountsCall: To_Balances_UpgradeAccountsCall(param),
        }
    case 6:
        return &pbgear.BalancesPallet_CallForceSetBalanceCall{
        CallForceSetBalanceCall: To_Balances_ForceSetBalanceCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_BalancesRescindedEvent(in any) *pbgear.BalancesRescindedEvent {
    out := &pbgear.BalancesRescindedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesReserveRepatriatedEvent(in any) *pbgear.BalancesReserveRepatriatedEvent {
    out := &pbgear.BalancesReserveRepatriatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesReservedEvent(in any) *pbgear.BalancesReservedEvent {
    out := &pbgear.BalancesReservedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesRestoredEvent(in any) *pbgear.BalancesRestoredEvent {
    out := &pbgear.BalancesRestoredEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesSlashedEvent(in any) *pbgear.BalancesSlashedEvent {
    out := &pbgear.BalancesSlashedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesSuspendedEvent(in any) *pbgear.BalancesSuspendedEvent {
    out := &pbgear.BalancesSuspendedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesThawedEvent(in any) *pbgear.BalancesThawedEvent {
    out := &pbgear.BalancesThawedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesTransferEvent(in any) *pbgear.BalancesTransferEvent {
    out := &pbgear.BalancesTransferEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesUnlockedEvent(in any) *pbgear.BalancesUnlockedEvent {
    out := &pbgear.BalancesUnlockedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesUnreservedEvent(in any) *pbgear.BalancesUnreservedEvent {
    out := &pbgear.BalancesUnreservedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesUpgradedEvent(in any) *pbgear.BalancesUpgradedEvent {
    out := &pbgear.BalancesUpgradedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BalancesWithdrawEvent(in any) *pbgear.BalancesWithdrawEvent {
    out := &pbgear.BalancesWithdrawEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_Added(in any) *pbgear.Balances_Added {
    out := &pbgear.Balances_Added{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field ChildIndex
    out.ChildIndex = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_Address20(in any) *pbgear.Balances_Address20 {
    out := &pbgear.Balances_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Address32(in any) *pbgear.Balances_Address32 {
    out := &pbgear.Balances_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_AllGood(in any) *pbgear.Balances_AllGood {
    out := &pbgear.Balances_AllGood{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_Announced(in any) *pbgear.Balances_Announced {
    out := &pbgear.Balances_Announced{}
    v := in.(registry.Valuable)
    _ = v

    // field Real To_SpCoreCryptoAccountId32(w)
    out.Real = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Proxy To_SpCoreCryptoAccountId32(w)
    out.Proxy = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(2))

    return out //from message
}
    



func To_Balances_Approved(in any) *pbgear.Balances_Approved {
    out := &pbgear.Balances_Approved{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_AssetSpendApproved(in any) *pbgear.Balances_AssetSpendApproved {
    out := &pbgear.Balances_AssetSpendApproved{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field AssetKind To_Balances_TupleNull(w)
    out.AssetKind = To_Balances_TupleNull(v.ValueAt(1))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(2))
    // field Beneficiary To_SpCoreCryptoAccountId32(w)
    out.Beneficiary = To_SpCoreCryptoAccountId32(v.ValueAt(3))
    // primitive field ValidFrom
    out.ValidFrom = To_uint32(v.ValueAt(4))
    // primitive field ExpireAt
    out.ExpireAt = To_uint32(v.ValueAt(5))

    return out //from message
}
    


func To_Balances_AssetSpendVoided(in any) *pbgear.Balances_AssetSpendVoided {
    out := &pbgear.Balances_AssetSpendVoided{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Awarded(in any) *pbgear.Balances_Awarded {
    out := &pbgear.Balances_Awarded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ProposalIndex
    out.ProposalIndex = To_uint32(v.ValueAt(0))
    // primitive field Award
    out.Award = To_string(v.ValueAt(1))
    // field Account To_SpCoreCryptoAccountId32(w)
    out.Account = To_SpCoreCryptoAccountId32(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_BagsListEvent(in any) *pbgear.Balances_BagsListEvent {
    out := &pbgear.Balances_BagsListEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_BagsListEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_BagsListEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_BagsListEvent_Value0Rebagged{
        Value0Rebagged: To_Balances_Rebagged(param),
        }
    case 1:
        return &pbgear.Balances_BagsListEvent_Value0ScoreUpdated{
        Value0ScoreUpdated: To_Balances_ScoreUpdated(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_BalanceSet(in any) *pbgear.Balances_BalanceSet {
    out := &pbgear.Balances_BalanceSet{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Free
    out.Free = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_BalancesEvent(in any) *pbgear.Balances_BalancesEvent {
    out := &pbgear.Balances_BalancesEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_BalancesEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_BalancesEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_BalancesEvent_Value0Endowed{
        Value0Endowed: To_Balances_Endowed(param),
        }
    case 1:
        return &pbgear.Balances_BalancesEvent_Value0DustLost{
        Value0DustLost: To_Balances_DustLost(param),
        }
    case 2:
        return &pbgear.Balances_BalancesEvent_Value0Transfer{
        Value0Transfer: To_Balances_Transfer(param),
        }
    case 3:
        return &pbgear.Balances_BalancesEvent_Value0BalanceSet{
        Value0BalanceSet: To_Balances_BalanceSet(param),
        }
    case 4:
        return &pbgear.Balances_BalancesEvent_Value0Reserved{
        Value0Reserved: To_Balances_Reserved(param),
        }
    case 5:
        return &pbgear.Balances_BalancesEvent_Value0Unreserved{
        Value0Unreserved: To_Balances_Unreserved(param),
        }
    case 6:
        return &pbgear.Balances_BalancesEvent_Value0ReserveRepatriated{
        Value0ReserveRepatriated: To_Balances_ReserveRepatriated(param),
        }
    case 7:
        return &pbgear.Balances_BalancesEvent_Value0Deposit{
        Value0Deposit: To_Balances_Deposit(param),
        }
    case 8:
        return &pbgear.Balances_BalancesEvent_Value0Withdraw{
        Value0Withdraw: To_Balances_Withdraw(param),
        }
    case 9:
        return &pbgear.Balances_BalancesEvent_Value0Slashed{
        Value0Slashed: To_Balances_Slashed(param),
        }
    case 10:
        return &pbgear.Balances_BalancesEvent_Value0Minted{
        Value0Minted: To_Balances_Minted(param),
        }
    case 11:
        return &pbgear.Balances_BalancesEvent_Value0Burned{
        Value0Burned: To_Balances_Burned(param),
        }
    case 12:
        return &pbgear.Balances_BalancesEvent_Value0Suspended{
        Value0Suspended: To_Balances_Suspended(param),
        }
    case 13:
        return &pbgear.Balances_BalancesEvent_Value0Restored{
        Value0Restored: To_Balances_Restored(param),
        }
    case 14:
        return &pbgear.Balances_BalancesEvent_Value0Upgraded{
        Value0Upgraded: To_Balances_Upgraded(param),
        }
    case 15:
        return &pbgear.Balances_BalancesEvent_Value0Issued{
        Value0Issued: To_Balances_Issued(param),
        }
    case 16:
        return &pbgear.Balances_BalancesEvent_Value0Rescinded{
        Value0Rescinded: To_Balances_Rescinded(param),
        }
    case 17:
        return &pbgear.Balances_BalancesEvent_Value0Locked{
        Value0Locked: To_Balances_Locked(param),
        }
    case 18:
        return &pbgear.Balances_BalancesEvent_Value0Unlocked{
        Value0Unlocked: To_Balances_Unlocked(param),
        }
    case 19:
        return &pbgear.Balances_BalancesEvent_Value0Frozen{
        Value0Frozen: To_Balances_Frozen(param),
        }
    case 20:
        return &pbgear.Balances_BalancesEvent_Value0Thawed{
        Value0Thawed: To_Balances_Thawed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_BatchCompleted(in any) *pbgear.Balances_BatchCompleted {
    out := &pbgear.Balances_BatchCompleted{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_BatchCompletedWithErrors(in any) *pbgear.Balances_BatchCompletedWithErrors {
    out := &pbgear.Balances_BatchCompletedWithErrors{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_BatchInterrupted(in any) *pbgear.Balances_BatchInterrupted {
    out := &pbgear.Balances_BatchInterrupted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // oneOf field Error
    v1 := To_OneOf_Balances_BatchInterrupted_error(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Error").Set(reflect.ValueOf(v1))

    return out //from message
}
    
func To_OneOf_Balances_BatchInterrupted_error (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_BatchInterrupted_ErrorOther{
        ErrorOther: To_Other(param),
        }
    case 1:
        return &pbgear.Balances_BatchInterrupted_ErrorCannotLookup{
        ErrorCannotLookup: To_CannotLookup(param),
        }
    case 2:
        return &pbgear.Balances_BatchInterrupted_ErrorBadOrigin{
        ErrorBadOrigin: To_BadOrigin(param),
        }
    case 3:
        return &pbgear.Balances_BatchInterrupted_ErrorModule{
        ErrorModule: To_Module(param),
        }
    case 4:
        return &pbgear.Balances_BatchInterrupted_ErrorConsumerRemaining{
        ErrorConsumerRemaining: To_ConsumerRemaining(param),
        }
    case 5:
        return &pbgear.Balances_BatchInterrupted_ErrorNoProviders{
        ErrorNoProviders: To_NoProviders(param),
        }
    case 6:
        return &pbgear.Balances_BatchInterrupted_ErrorTooManyConsumers{
        ErrorTooManyConsumers: To_TooManyConsumers(param),
        }
    case 7:
        return &pbgear.Balances_BatchInterrupted_ErrorToken{
        ErrorToken: To_Token(param),
        }
    case 8:
        return &pbgear.Balances_BatchInterrupted_ErrorArithmetic{
        ErrorArithmetic: To_Arithmetic(param),
        }
    case 9:
        return &pbgear.Balances_BatchInterrupted_ErrorTransactional{
        ErrorTransactional: To_Transactional(param),
        }
    case 10:
        return &pbgear.Balances_BatchInterrupted_ErrorExhausted{
        ErrorExhausted: To_Exhausted(param),
        }
    case 11:
        return &pbgear.Balances_BatchInterrupted_ErrorCorruption{
        ErrorCorruption: To_Corruption(param),
        }
    case 12:
        return &pbgear.Balances_BatchInterrupted_ErrorUnavailable{
        ErrorUnavailable: To_Unavailable(param),
        }
    case 13:
        return &pbgear.Balances_BatchInterrupted_ErrorRootNotAllowed{
        ErrorRootNotAllowed: To_RootNotAllowed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Bonded(in any) *pbgear.Balances_Bonded {
    out := &pbgear.Balances_Bonded{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_BountiesEvent(in any) *pbgear.Balances_BountiesEvent {
    out := &pbgear.Balances_BountiesEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_BountiesEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_BountiesEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_BountiesEvent_Value0BountyProposed{
        Value0BountyProposed: To_Balances_BountyProposed(param),
        }
    case 1:
        return &pbgear.Balances_BountiesEvent_Value0BountyRejected{
        Value0BountyRejected: To_Balances_BountyRejected(param),
        }
    case 2:
        return &pbgear.Balances_BountiesEvent_Value0BountyBecameActive{
        Value0BountyBecameActive: To_Balances_BountyBecameActive(param),
        }
    case 3:
        return &pbgear.Balances_BountiesEvent_Value0BountyAwarded{
        Value0BountyAwarded: To_Balances_BountyAwarded(param),
        }
    case 4:
        return &pbgear.Balances_BountiesEvent_Value0BountyClaimed{
        Value0BountyClaimed: To_Balances_BountyClaimed(param),
        }
    case 5:
        return &pbgear.Balances_BountiesEvent_Value0BountyCanceled{
        Value0BountyCanceled: To_Balances_BountyCanceled(param),
        }
    case 6:
        return &pbgear.Balances_BountiesEvent_Value0BountyExtended{
        Value0BountyExtended: To_Balances_BountyExtended(param),
        }
    case 7:
        return &pbgear.Balances_BountiesEvent_Value0BountyApproved{
        Value0BountyApproved: To_Balances_BountyApproved(param),
        }
    case 8:
        return &pbgear.Balances_BountiesEvent_Value0CuratorProposed{
        Value0CuratorProposed: To_Balances_CuratorProposed(param),
        }
    case 9:
        return &pbgear.Balances_BountiesEvent_Value0CuratorUnassigned{
        Value0CuratorUnassigned: To_Balances_CuratorUnassigned(param),
        }
    case 10:
        return &pbgear.Balances_BountiesEvent_Value0CuratorAccepted{
        Value0CuratorAccepted: To_Balances_CuratorAccepted(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_BountyApproved(in any) *pbgear.Balances_BountyApproved {
    out := &pbgear.Balances_BountyApproved{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_BountyAwarded(in any) *pbgear.Balances_BountyAwarded {
    out := &pbgear.Balances_BountyAwarded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Beneficiary To_SpCoreCryptoAccountId32(w)
    out.Beneficiary = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_BountyBecameActive(in any) *pbgear.Balances_BountyBecameActive {
    out := &pbgear.Balances_BountyBecameActive{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_BountyCanceled(in any) *pbgear.Balances_BountyCanceled {
    out := &pbgear.Balances_BountyCanceled{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_BountyClaimed(in any) *pbgear.Balances_BountyClaimed {
    out := &pbgear.Balances_BountyClaimed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field Payout
    out.Payout = To_string(v.ValueAt(1))
    // field Beneficiary To_SpCoreCryptoAccountId32(w)
    out.Beneficiary = To_SpCoreCryptoAccountId32(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_BountyExtended(in any) *pbgear.Balances_BountyExtended {
    out := &pbgear.Balances_BountyExtended{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_BountyProposed(in any) *pbgear.Balances_BountyProposed {
    out := &pbgear.Balances_BountyProposed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_BountyRejected(in any) *pbgear.Balances_BountyRejected {
    out := &pbgear.Balances_BountyRejected{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field Bond
    out.Bond = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_Burned(in any) *pbgear.Balances_Burned {
    out := &pbgear.Balances_Burned{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Burnt(in any) *pbgear.Balances_Burnt {
    out := &pbgear.Balances_Burnt{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BurntFunds
    out.BurntFunds = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_CallUnavailable(in any) *pbgear.Balances_CallUnavailable {
    out := &pbgear.Balances_CallUnavailable{}
    v := in.(registry.Valuable)
    _ = v

    // field Task To_Balances_TupleUint32Uint32(w)
    out.Task = To_Balances_TupleUint32Uint32(v.ValueAt(0))
    // primitive field Id
    out.Id = To_bytes(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_CallWhitelisted(in any) *pbgear.Balances_CallWhitelisted {
    out := &pbgear.Balances_CallWhitelisted{}
    v := in.(registry.Valuable)
    _ = v

    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Canceled(in any) *pbgear.Balances_Canceled {
    out := &pbgear.Balances_Canceled{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field When
    out.When = To_uint32(v.ValueAt(0))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_Cancelled(in any) *pbgear.Balances_Cancelled {
    out := &pbgear.Balances_Cancelled{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Tally To_ConvictionVoting_PalletConvictionVotingTypesTally(w)
    out.Tally = To_ConvictionVoting_PalletConvictionVotingTypesTally(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_ChildBountiesEvent(in any) *pbgear.Balances_ChildBountiesEvent {
    out := &pbgear.Balances_ChildBountiesEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_ChildBountiesEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ChildBountiesEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ChildBountiesEvent_Value0Added{
        Value0Added: To_Balances_Added(param),
        }
    case 1:
        return &pbgear.Balances_ChildBountiesEvent_Value0Awarded{
        Value0Awarded: To_Balances_Awarded(param),
        }
    case 2:
        return &pbgear.Balances_ChildBountiesEvent_Value0Claimed{
        Value0Claimed: To_Balances_Claimed(param),
        }
    case 3:
        return &pbgear.Balances_ChildBountiesEvent_Value0Canceled{
        Value0Canceled: To_Balances_Canceled(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Chilled(in any) *pbgear.Balances_Chilled {
    out := &pbgear.Balances_Chilled{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Claimed(in any) *pbgear.Balances_Claimed {
    out := &pbgear.Balances_Claimed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field ChildIndex
    out.ChildIndex = To_uint32(v.ValueAt(1))
    // primitive field Payout
    out.Payout = To_string(v.ValueAt(2))
    // field Beneficiary To_SpCoreCryptoAccountId32(w)
    out.Beneficiary = To_SpCoreCryptoAccountId32(v.ValueAt(3))

    return out //from message
}
    

func To_Balances_Cleared(in any) *pbgear.Balances_Cleared {
    out := &pbgear.Balances_Cleared{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_CodeChanged(in any) *pbgear.Balances_CodeChanged {
    out := &pbgear.Balances_CodeChanged{}
    v := in.(registry.Valuable)
    _ = v

    // field Id To_GprimitivesCodeId(w)
    out.Id = To_GprimitivesCodeId(v.ValueAt(0))
    // oneOf field Change
    v1 := To_OneOf_Balances_CodeChanged_change(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Change").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_Balances_CodeChanged_change (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_CodeChanged_ChangeActive{
        ChangeActive: To_Active(param),
        }
    case 1:
        return &pbgear.Balances_CodeChanged_ChangeInactive{
        ChangeInactive: To_Inactive(param),
        }
    case 2:
        return &pbgear.Balances_CodeChanged_ChangeReinstrumented{
        ChangeReinstrumented: To_Reinstrumented(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_CodeUpdated(in any) *pbgear.Balances_CodeUpdated {
    out := &pbgear.Balances_CodeUpdated{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_ConfirmAborted(in any) *pbgear.Balances_ConfirmAborted {
    out := &pbgear.Balances_ConfirmAborted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_ConfirmStarted(in any) *pbgear.Balances_ConfirmStarted {
    out := &pbgear.Balances_ConfirmStarted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Confirmed(in any) *pbgear.Balances_Confirmed {
    out := &pbgear.Balances_Confirmed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Tally To_ConvictionVoting_PalletConvictionVotingTypesTally(w)
    out.Tally = To_ConvictionVoting_PalletConvictionVotingTypesTally(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_ConsensusEvent(in any) *pbgear.Balances_ConsensusEvent {
    out := &pbgear.Balances_ConsensusEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_ConvictionVotingEvent(in any) *pbgear.Balances_ConvictionVotingEvent {
    out := &pbgear.Balances_ConvictionVotingEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_ConvictionVotingEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ConvictionVotingEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ConvictionVotingEvent_Value0Delegated{
        Value0Delegated: To_Balances_Delegated(param),
        }
    case 1:
        return &pbgear.Balances_ConvictionVotingEvent_Value0Undelegated{
        Value0Undelegated: To_Balances_Undelegated(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Created(in any) *pbgear.Balances_Created {
    out := &pbgear.Balances_Created{}
    v := in.(registry.Valuable)
    _ = v

    // field Depositor To_SpCoreCryptoAccountId32(w)
    out.Depositor = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_CuratorAccepted(in any) *pbgear.Balances_CuratorAccepted {
    out := &pbgear.Balances_CuratorAccepted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))
    // field Curator To_SpCoreCryptoAccountId32(w)
    out.Curator = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_CuratorProposed(in any) *pbgear.Balances_CuratorProposed {
    out := &pbgear.Balances_CuratorProposed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))
    // field Curator To_SpCoreCryptoAccountId32(w)
    out.Curator = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_CuratorUnassigned(in any) *pbgear.Balances_CuratorUnassigned {
    out := &pbgear.Balances_CuratorUnassigned{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_DecisionDepositPlaced(in any) *pbgear.Balances_DecisionDepositPlaced {
    out := &pbgear.Balances_DecisionDepositPlaced{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_DecisionDepositRefunded(in any) *pbgear.Balances_DecisionDepositRefunded {
    out := &pbgear.Balances_DecisionDepositRefunded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_DecisionStarted(in any) *pbgear.Balances_DecisionStarted {
    out := &pbgear.Balances_DecisionStarted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field Track
    out.Track = To_uint32(v.ValueAt(1))
    // oneOf field Proposal
    v2 := To_OneOf_Balances_DecisionStarted_proposal(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v2))
    // field Tally To_ConvictionVoting_PalletConvictionVotingTypesTally(w)
    out.Tally = To_ConvictionVoting_PalletConvictionVotingTypesTally(v.ValueAt(3))

    return out //from message
}
    
func To_OneOf_Balances_DecisionStarted_proposal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_DecisionStarted_ProposalLegacy{
        ProposalLegacy: To_Legacy(param),
        }
    case 1:
        return &pbgear.Balances_DecisionStarted_ProposalInline{
        ProposalInline: To_Inline(param),
        }
    case 2:
        return &pbgear.Balances_DecisionStarted_ProposalLookup{
        ProposalLookup: To_Lookup(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Balances_Delegated(in any) *pbgear.Balances_Delegated {
    out := &pbgear.Balances_Delegated{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Value1 To_SpCoreCryptoAccountId32(w)
    out.Value1 = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    


func To_Balances_Deposit(in any) *pbgear.Balances_Deposit {
    out := &pbgear.Balances_Deposit{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_DepositSlashed(in any) *pbgear.Balances_DepositSlashed {
    out := &pbgear.Balances_DepositSlashed{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Deposited(in any) *pbgear.Balances_Deposited {
    out := &pbgear.Balances_Deposited{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Amount
    out.Amount = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Dest(in any) *pbgear.Balances_Dest {
    out := &pbgear.Balances_Dest{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Dest
    v0 := To_OneOf_Balances_Dest_dest(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_Dest_dest (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Dest_DestId{
        DestId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_Dest_DestIndex{
        DestIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_Dest_DestRaw{
        DestRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_Dest_DestAddress32{
        DestAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_Dest_DestAddress20{
        DestAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Destroyed(in any) *pbgear.Balances_Destroyed {
    out := &pbgear.Balances_Destroyed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Dispatched(in any) *pbgear.Balances_Dispatched {
    out := &pbgear.Balances_Dispatched{}
    v := in.(registry.Valuable)
    _ = v

    // field Task To_Balances_TupleUint32Uint32(w)
    out.Task = To_Balances_TupleUint32Uint32(v.ValueAt(0))
    // primitive field Id
    out.Id = To_bytes(v.ValueAt(1))
    // oneOf field Result
    v2 := To_OneOf_Balances_Dispatched_result(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Result").Set(reflect.ValueOf(v2))

    return out //from message
}
    

func To_OneOf_Balances_Dispatched_result (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Dispatched_ResultOk{
        ResultOk: To_Ok(param),
        }
    case 1:
        return &pbgear.Balances_Dispatched_ResultErr{
        ResultErr: To_Err(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_DispatchedAs(in any) *pbgear.Balances_DispatchedAs {
    out := &pbgear.Balances_DispatchedAs{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Result
    v0 := To_OneOf_Balances_DispatchedAs_result(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Result").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_DispatchedAs_result (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_DispatchedAs_ResultOk{
        ResultOk: To_Ok(param),
        }
    case 1:
        return &pbgear.Balances_DispatchedAs_ResultErr{
        ResultErr: To_Err(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_DustLost(in any) *pbgear.Balances_DustLost {
    out := &pbgear.Balances_DustLost{}
    v := in.(registry.Valuable)
    _ = v

    // field Account To_SpCoreCryptoAccountId32(w)
    out.Account = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_ElectionFailed(in any) *pbgear.Balances_ElectionFailed {
    out := &pbgear.Balances_ElectionFailed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_ElectionFinalized(in any) *pbgear.Balances_ElectionFinalized {
    out := &pbgear.Balances_ElectionFinalized{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Compute
    v0 := To_OneOf_Balances_ElectionFinalized_compute(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Compute").Set(reflect.ValueOf(v0))
    // field Score To_SpNposElectionsElectionScore(w)
    out.Score = To_SpNposElectionsElectionScore(v.ValueAt(1))

    return out //from message
}
    
func To_OneOf_Balances_ElectionFinalized_compute (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ElectionFinalized_ComputeOnChain{
        ComputeOnChain: To_ElectionProviderMultiPhase_OnChain(param),
        }
    case 1:
        return &pbgear.Balances_ElectionFinalized_ComputeSigned{
        ComputeSigned: To_ElectionProviderMultiPhase_Signed(param),
        }
    case 2:
        return &pbgear.Balances_ElectionFinalized_ComputeUnsigned{
        ComputeUnsigned: To_ElectionProviderMultiPhase_Unsigned(param),
        }
    case 3:
        return &pbgear.Balances_ElectionFinalized_ComputeFallback{
        ComputeFallback: To_ElectionProviderMultiPhase_Fallback(param),
        }
    case 4:
        return &pbgear.Balances_ElectionFinalized_ComputeEmergency{
        ComputeEmergency: To_ElectionProviderMultiPhase_Emergency(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Balances_ElectionProviderMultiPhaseEvent(in any) *pbgear.Balances_ElectionProviderMultiPhaseEvent {
    out := &pbgear.Balances_ElectionProviderMultiPhaseEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_ElectionProviderMultiPhaseEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ElectionProviderMultiPhaseEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ElectionProviderMultiPhaseEvent_Value0SolutionStored{
        Value0SolutionStored: To_Balances_SolutionStored(param),
        }
    case 1:
        return &pbgear.Balances_ElectionProviderMultiPhaseEvent_Value0ElectionFinalized{
        Value0ElectionFinalized: To_Balances_ElectionFinalized(param),
        }
    case 2:
        return &pbgear.Balances_ElectionProviderMultiPhaseEvent_Value0ElectionFailed{
        Value0ElectionFailed: To_Balances_ElectionFailed(param),
        }
    case 3:
        return &pbgear.Balances_ElectionProviderMultiPhaseEvent_Value0Rewarded{
        Value0Rewarded: To_Balances_Rewarded(param),
        }
    case 4:
        return &pbgear.Balances_ElectionProviderMultiPhaseEvent_Value0Slashed{
        Value0Slashed: To_Balances_Slashed(param),
        }
    case 5:
        return &pbgear.Balances_ElectionProviderMultiPhaseEvent_Value0PhaseTransitioned{
        Value0PhaseTransitioned: To_Balances_PhaseTransitioned(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Endowed(in any) *pbgear.Balances_Endowed {
    out := &pbgear.Balances_Endowed{}
    v := in.(registry.Valuable)
    _ = v

    // field Account To_SpCoreCryptoAccountId32(w)
    out.Account = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field FreeBalance
    out.FreeBalance = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_EraPaid(in any) *pbgear.Balances_EraPaid {
    out := &pbgear.Balances_EraPaid{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field EraIndex
    out.EraIndex = To_uint32(v.ValueAt(0))
    // primitive field ValidatorPayout
    out.ValidatorPayout = To_string(v.ValueAt(1))
    // primitive field Remainder
    out.Remainder = To_string(v.ValueAt(2))

    return out //from message
}
    
func To_Balances_ExtrinsicFailed(in any) *pbgear.Balances_ExtrinsicFailed {
    out := &pbgear.Balances_ExtrinsicFailed{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field DispatchError
    v0 := To_OneOf_Balances_ExtrinsicFailed_dispatch_error(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("DispatchError").Set(reflect.ValueOf(v0))
    // field DispatchInfo To_FrameSupportDispatchDispatchInfo(w)
    out.DispatchInfo = To_FrameSupportDispatchDispatchInfo(v.ValueAt(1))

    return out //from message
}
    
func To_OneOf_Balances_ExtrinsicFailed_dispatch_error (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorOther{
        DispatchErrorOther: To_Other(param),
        }
    case 1:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorCannotLookup{
        DispatchErrorCannotLookup: To_CannotLookup(param),
        }
    case 2:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorBadOrigin{
        DispatchErrorBadOrigin: To_BadOrigin(param),
        }
    case 3:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorModule{
        DispatchErrorModule: To_Module(param),
        }
    case 4:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorConsumerRemaining{
        DispatchErrorConsumerRemaining: To_ConsumerRemaining(param),
        }
    case 5:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorNoProviders{
        DispatchErrorNoProviders: To_NoProviders(param),
        }
    case 6:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorTooManyConsumers{
        DispatchErrorTooManyConsumers: To_TooManyConsumers(param),
        }
    case 7:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorToken{
        DispatchErrorToken: To_Token(param),
        }
    case 8:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorArithmetic{
        DispatchErrorArithmetic: To_Arithmetic(param),
        }
    case 9:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorTransactional{
        DispatchErrorTransactional: To_Transactional(param),
        }
    case 10:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorExhausted{
        DispatchErrorExhausted: To_Exhausted(param),
        }
    case 11:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorCorruption{
        DispatchErrorCorruption: To_Corruption(param),
        }
    case 12:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorUnavailable{
        DispatchErrorUnavailable: To_Unavailable(param),
        }
    case 13:
        return &pbgear.Balances_ExtrinsicFailed_DispatchErrorRootNotAllowed{
        DispatchErrorRootNotAllowed: To_RootNotAllowed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Balances_ExtrinsicSuccess(in any) *pbgear.Balances_ExtrinsicSuccess {
    out := &pbgear.Balances_ExtrinsicSuccess{}
    v := in.(registry.Valuable)
    _ = v

    // field DispatchInfo To_FrameSupportDispatchDispatchInfo(w)
    out.DispatchInfo = To_FrameSupportDispatchDispatchInfo(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_FellowshipCollectiveEvent(in any) *pbgear.Balances_FellowshipCollectiveEvent {
    out := &pbgear.Balances_FellowshipCollectiveEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_FellowshipCollectiveEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_FellowshipCollectiveEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_FellowshipCollectiveEvent_Value0MemberAdded{
        Value0MemberAdded: To_Balances_MemberAdded(param),
        }
    case 1:
        return &pbgear.Balances_FellowshipCollectiveEvent_Value0RankChanged{
        Value0RankChanged: To_Balances_RankChanged(param),
        }
    case 2:
        return &pbgear.Balances_FellowshipCollectiveEvent_Value0MemberRemoved{
        Value0MemberRemoved: To_Balances_MemberRemoved(param),
        }
    case 3:
        return &pbgear.Balances_FellowshipCollectiveEvent_Value0Voted{
        Value0Voted: To_Balances_Voted(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_FellowshipReferendaEvent(in any) *pbgear.Balances_FellowshipReferendaEvent {
    out := &pbgear.Balances_FellowshipReferendaEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_FellowshipReferendaEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_FellowshipReferendaEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0Submitted{
        Value0Submitted: To_Balances_Submitted(param),
        }
    case 1:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0DecisionDepositPlaced{
        Value0DecisionDepositPlaced: To_Balances_DecisionDepositPlaced(param),
        }
    case 2:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0DecisionDepositRefunded{
        Value0DecisionDepositRefunded: To_Balances_DecisionDepositRefunded(param),
        }
    case 3:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0DepositSlashed{
        Value0DepositSlashed: To_Balances_DepositSlashed(param),
        }
    case 4:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0DecisionStarted{
        Value0DecisionStarted: To_Balances_DecisionStarted(param),
        }
    case 5:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0ConfirmStarted{
        Value0ConfirmStarted: To_Balances_ConfirmStarted(param),
        }
    case 6:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0ConfirmAborted{
        Value0ConfirmAborted: To_Balances_ConfirmAborted(param),
        }
    case 7:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0Confirmed{
        Value0Confirmed: To_Balances_Confirmed(param),
        }
    case 8:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0Approved{
        Value0Approved: To_Balances_Approved(param),
        }
    case 9:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0Rejected{
        Value0Rejected: To_Balances_Rejected(param),
        }
    case 10:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0TimedOut{
        Value0TimedOut: To_Balances_TimedOut(param),
        }
    case 11:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0Cancelled{
        Value0Cancelled: To_Balances_Cancelled(param),
        }
    case 12:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0Killed{
        Value0Killed: To_Balances_Killed(param),
        }
    case 13:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0SubmissionDepositRefunded{
        Value0SubmissionDepositRefunded: To_Balances_SubmissionDepositRefunded(param),
        }
    case 14:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0MetadataSet{
        Value0MetadataSet: To_Balances_MetadataSet(param),
        }
    case 15:
        return &pbgear.Balances_FellowshipReferendaEvent_Value0MetadataCleared{
        Value0MetadataCleared: To_Balances_MetadataCleared(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ForceEra(in any) *pbgear.Balances_ForceEra {
    out := &pbgear.Balances_ForceEra{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Mode
    v0 := To_OneOf_Balances_ForceEra_mode(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Mode").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ForceEra_mode (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ForceEra_ModeNotForcing{
        ModeNotForcing: To_Staking_NotForcing(param),
        }
    case 1:
        return &pbgear.Balances_ForceEra_ModeForceNew{
        ModeForceNew: To_Staking_ForceNew(param),
        }
    case 2:
        return &pbgear.Balances_ForceEra_ModeForceNone{
        ModeForceNone: To_Staking_ForceNone(param),
        }
    case 3:
        return &pbgear.Balances_ForceEra_ModeForceAlways{
        ModeForceAlways: To_Staking_ForceAlways(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ForceSetBalanceCall(in any) *pbgear.Balances_ForceSetBalanceCall {
    out := &pbgear.Balances_ForceSetBalanceCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_Balances_ForceSetBalanceCall_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))
    // primitive field NewFree
    out.NewFree = To_string(v.ValueAt(1))

    return out //from message
}
func To_Balances_ForceSetBalanceCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Balances_ForceSetBalanceCall(in)
    out := &pbgearextrinsic.RuntimeCall_BalancesForceSetBalanceCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BalancesForceSetBalanceCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Balances_ForceSetBalanceCall_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ForceSetBalanceCall_WhoId{
        WhoId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_ForceSetBalanceCall_WhoIndex{
        WhoIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_ForceSetBalanceCall_WhoRaw{
        WhoRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_ForceSetBalanceCall_WhoAddress32{
        WhoAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_ForceSetBalanceCall_WhoAddress20{
        WhoAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ForceTransferCall(in any) *pbgear.Balances_ForceTransferCall {
    out := &pbgear.Balances_ForceTransferCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Source
    v0 := To_OneOf_Balances_ForceTransferCall_source(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))
    // oneOf field Dest
    v1 := To_OneOf_Balances_ForceTransferCall_dest(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v1))
    // primitive field Value
    out.Value = To_string(v.ValueAt(2))

    return out //from message
}
func To_Balances_ForceTransferCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Balances_ForceTransferCall(in)
    out := &pbgearextrinsic.RuntimeCall_BalancesForceTransferCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BalancesForceTransferCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Balances_ForceTransferCall_source (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ForceTransferCall_SourceId{
        SourceId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_ForceTransferCall_SourceIndex{
        SourceIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_ForceTransferCall_SourceRaw{
        SourceRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_ForceTransferCall_SourceAddress32{
        SourceAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_ForceTransferCall_SourceAddress20{
        SourceAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Balances_ForceTransferCall_dest (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ForceTransferCall_DestId{
        DestId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_ForceTransferCall_DestIndex{
        DestIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_ForceTransferCall_DestRaw{
        DestRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_ForceTransferCall_DestAddress32{
        DestAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_ForceTransferCall_DestAddress20{
        DestAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ForceUnreserveCall(in any) *pbgear.Balances_ForceUnreserveCall {
    out := &pbgear.Balances_ForceUnreserveCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_Balances_ForceUnreserveCall_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
func To_Balances_ForceUnreserveCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Balances_ForceUnreserveCall(in)
    out := &pbgearextrinsic.RuntimeCall_BalancesForceUnreserveCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BalancesForceUnreserveCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Balances_ForceUnreserveCall_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ForceUnreserveCall_WhoId{
        WhoId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_ForceUnreserveCall_WhoIndex{
        WhoIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_ForceUnreserveCall_WhoRaw{
        WhoRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_ForceUnreserveCall_WhoAddress32{
        WhoAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_ForceUnreserveCall_WhoAddress20{
        WhoAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Frozen(in any) *pbgear.Balances_Frozen {
    out := &pbgear.Balances_Frozen{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_GearEvent(in any) *pbgear.Balances_GearEvent {
    out := &pbgear.Balances_GearEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_GearEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_GearEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_GearEvent_Value0MessageQueued{
        Value0MessageQueued: To_Balances_MessageQueued(param),
        }
    case 1:
        return &pbgear.Balances_GearEvent_Value0UserMessageSent{
        Value0UserMessageSent: To_Balances_UserMessageSent(param),
        }
    case 2:
        return &pbgear.Balances_GearEvent_Value0UserMessageRead{
        Value0UserMessageRead: To_Balances_UserMessageRead(param),
        }
    case 3:
        return &pbgear.Balances_GearEvent_Value0MessagesDispatched{
        Value0MessagesDispatched: To_Balances_MessagesDispatched(param),
        }
    case 4:
        return &pbgear.Balances_GearEvent_Value0MessageWaited{
        Value0MessageWaited: To_Balances_MessageWaited(param),
        }
    case 5:
        return &pbgear.Balances_GearEvent_Value0MessageWoken{
        Value0MessageWoken: To_Balances_MessageWoken(param),
        }
    case 6:
        return &pbgear.Balances_GearEvent_Value0CodeChanged{
        Value0CodeChanged: To_Balances_CodeChanged(param),
        }
    case 7:
        return &pbgear.Balances_GearEvent_Value0ProgramChanged{
        Value0ProgramChanged: To_Balances_ProgramChanged(param),
        }
    case 8:
        return &pbgear.Balances_GearEvent_Value0QueueNotProcessed{
        Value0QueueNotProcessed: To_Balances_QueueNotProcessed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_GearVoucherEvent(in any) *pbgear.Balances_GearVoucherEvent {
    out := &pbgear.Balances_GearVoucherEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_GearVoucherEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_GearVoucherEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_GearVoucherEvent_Value0VoucherIssued{
        Value0VoucherIssued: To_Balances_VoucherIssued(param),
        }
    case 1:
        return &pbgear.Balances_GearVoucherEvent_Value0VoucherRevoked{
        Value0VoucherRevoked: To_Balances_VoucherRevoked(param),
        }
    case 2:
        return &pbgear.Balances_GearVoucherEvent_Value0VoucherUpdated{
        Value0VoucherUpdated: To_Balances_VoucherUpdated(param),
        }
    case 3:
        return &pbgear.Balances_GearVoucherEvent_Value0VoucherDeclined{
        Value0VoucherDeclined: To_Balances_VoucherDeclined(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_GrandpaEvent(in any) *pbgear.Balances_GrandpaEvent {
    out := &pbgear.Balances_GrandpaEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_GrandpaEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_GrandpaEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_GrandpaEvent_Value0NewAuthorities{
        Value0NewAuthorities: To_Balances_NewAuthorities(param),
        }
    case 1:
        return &pbgear.Balances_GrandpaEvent_Value0Paused{
        Value0Paused: To_Balances_Paused(param),
        }
    case 2:
        return &pbgear.Balances_GrandpaEvent_Value0Resumed{
        Value0Resumed: To_Balances_Resumed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_HeartbeatReceived(in any) *pbgear.Balances_HeartbeatReceived {
    out := &pbgear.Balances_HeartbeatReceived{}
    v := in.(registry.Valuable)
    _ = v

    // field AuthorityId To_ImOnline_PalletImOnlineSr25519AppSr25519Public(w)
    out.AuthorityId = To_ImOnline_PalletImOnlineSr25519AppSr25519Public(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Id(in any) *pbgear.Balances_Id {
    out := &pbgear.Balances_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_IdentityCleared(in any) *pbgear.Balances_IdentityCleared {
    out := &pbgear.Balances_IdentityCleared{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Deposit
    out.Deposit = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_IdentityEvent(in any) *pbgear.Balances_IdentityEvent {
    out := &pbgear.Balances_IdentityEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_IdentityEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_IdentityEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_IdentityEvent_Value0IdentitySet{
        Value0IdentitySet: To_Balances_IdentitySet(param),
        }
    case 1:
        return &pbgear.Balances_IdentityEvent_Value0IdentityCleared{
        Value0IdentityCleared: To_Balances_IdentityCleared(param),
        }
    case 2:
        return &pbgear.Balances_IdentityEvent_Value0IdentityKilled{
        Value0IdentityKilled: To_Balances_IdentityKilled(param),
        }
    case 3:
        return &pbgear.Balances_IdentityEvent_Value0JudgementRequested{
        Value0JudgementRequested: To_Balances_JudgementRequested(param),
        }
    case 4:
        return &pbgear.Balances_IdentityEvent_Value0JudgementUnrequested{
        Value0JudgementUnrequested: To_Balances_JudgementUnrequested(param),
        }
    case 5:
        return &pbgear.Balances_IdentityEvent_Value0JudgementGiven{
        Value0JudgementGiven: To_Balances_JudgementGiven(param),
        }
    case 6:
        return &pbgear.Balances_IdentityEvent_Value0RegistrarAdded{
        Value0RegistrarAdded: To_Balances_RegistrarAdded(param),
        }
    case 7:
        return &pbgear.Balances_IdentityEvent_Value0SubIdentityAdded{
        Value0SubIdentityAdded: To_Balances_SubIdentityAdded(param),
        }
    case 8:
        return &pbgear.Balances_IdentityEvent_Value0SubIdentityRemoved{
        Value0SubIdentityRemoved: To_Balances_SubIdentityRemoved(param),
        }
    case 9:
        return &pbgear.Balances_IdentityEvent_Value0SubIdentityRevoked{
        Value0SubIdentityRevoked: To_Balances_SubIdentityRevoked(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_IdentityKilled(in any) *pbgear.Balances_IdentityKilled {
    out := &pbgear.Balances_IdentityKilled{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Deposit
    out.Deposit = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_IdentitySet(in any) *pbgear.Balances_IdentitySet {
    out := &pbgear.Balances_IdentitySet{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_ImOnlineEvent(in any) *pbgear.Balances_ImOnlineEvent {
    out := &pbgear.Balances_ImOnlineEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_ImOnlineEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ImOnlineEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ImOnlineEvent_Value0HeartbeatReceived{
        Value0HeartbeatReceived: To_Balances_HeartbeatReceived(param),
        }
    case 1:
        return &pbgear.Balances_ImOnlineEvent_Value0AllGood{
        Value0AllGood: To_Balances_AllGood(param),
        }
    case 2:
        return &pbgear.Balances_ImOnlineEvent_Value0SomeOffline{
        Value0SomeOffline: To_Balances_SomeOffline(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Index(in any) *pbgear.Balances_Index {
    out := &pbgear.Balances_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Balances_TupleNull(w)
    out.Value0 = To_Balances_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Issued(in any) *pbgear.Balances_Issued {
    out := &pbgear.Balances_Issued{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Amount
    out.Amount = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_ItemCompleted(in any) *pbgear.Balances_ItemCompleted {
    out := &pbgear.Balances_ItemCompleted{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_ItemFailed(in any) *pbgear.Balances_ItemFailed {
    out := &pbgear.Balances_ItemFailed{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Error
    v0 := To_OneOf_Balances_ItemFailed_error(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Error").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ItemFailed_error (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ItemFailed_ErrorOther{
        ErrorOther: To_Other(param),
        }
    case 1:
        return &pbgear.Balances_ItemFailed_ErrorCannotLookup{
        ErrorCannotLookup: To_CannotLookup(param),
        }
    case 2:
        return &pbgear.Balances_ItemFailed_ErrorBadOrigin{
        ErrorBadOrigin: To_BadOrigin(param),
        }
    case 3:
        return &pbgear.Balances_ItemFailed_ErrorModule{
        ErrorModule: To_Module(param),
        }
    case 4:
        return &pbgear.Balances_ItemFailed_ErrorConsumerRemaining{
        ErrorConsumerRemaining: To_ConsumerRemaining(param),
        }
    case 5:
        return &pbgear.Balances_ItemFailed_ErrorNoProviders{
        ErrorNoProviders: To_NoProviders(param),
        }
    case 6:
        return &pbgear.Balances_ItemFailed_ErrorTooManyConsumers{
        ErrorTooManyConsumers: To_TooManyConsumers(param),
        }
    case 7:
        return &pbgear.Balances_ItemFailed_ErrorToken{
        ErrorToken: To_Token(param),
        }
    case 8:
        return &pbgear.Balances_ItemFailed_ErrorArithmetic{
        ErrorArithmetic: To_Arithmetic(param),
        }
    case 9:
        return &pbgear.Balances_ItemFailed_ErrorTransactional{
        ErrorTransactional: To_Transactional(param),
        }
    case 10:
        return &pbgear.Balances_ItemFailed_ErrorExhausted{
        ErrorExhausted: To_Exhausted(param),
        }
    case 11:
        return &pbgear.Balances_ItemFailed_ErrorCorruption{
        ErrorCorruption: To_Corruption(param),
        }
    case 12:
        return &pbgear.Balances_ItemFailed_ErrorUnavailable{
        ErrorUnavailable: To_Unavailable(param),
        }
    case 13:
        return &pbgear.Balances_ItemFailed_ErrorRootNotAllowed{
        ErrorRootNotAllowed: To_RootNotAllowed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_JudgementGiven(in any) *pbgear.Balances_JudgementGiven {
    out := &pbgear.Balances_JudgementGiven{}
    v := in.(registry.Valuable)
    _ = v

    // field Target To_SpCoreCryptoAccountId32(w)
    out.Target = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field RegistrarIndex
    out.RegistrarIndex = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_JudgementRequested(in any) *pbgear.Balances_JudgementRequested {
    out := &pbgear.Balances_JudgementRequested{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field RegistrarIndex
    out.RegistrarIndex = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_JudgementUnrequested(in any) *pbgear.Balances_JudgementUnrequested {
    out := &pbgear.Balances_JudgementUnrequested{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field RegistrarIndex
    out.RegistrarIndex = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Kicked(in any) *pbgear.Balances_Kicked {
    out := &pbgear.Balances_Kicked{}
    v := in.(registry.Valuable)
    _ = v

    // field Nominator To_SpCoreCryptoAccountId32(w)
    out.Nominator = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    


func To_Balances_Killed(in any) *pbgear.Balances_Killed {
    out := &pbgear.Balances_Killed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Tally To_ConvictionVoting_PalletConvictionVotingTypesTally(w)
    out.Tally = To_ConvictionVoting_PalletConvictionVotingTypesTally(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_KilledAccount(in any) *pbgear.Balances_KilledAccount {
    out := &pbgear.Balances_KilledAccount{}
    v := in.(registry.Valuable)
    _ = v

    // field Account To_SpCoreCryptoAccountId32(w)
    out.Account = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Locked(in any) *pbgear.Balances_Locked {
    out := &pbgear.Balances_Locked{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_MemberAdded(in any) *pbgear.Balances_MemberAdded {
    out := &pbgear.Balances_MemberAdded{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_MemberRemoved(in any) *pbgear.Balances_MemberRemoved {
    out := &pbgear.Balances_MemberRemoved{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Rank
    out.Rank = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_MessageQueued(in any) *pbgear.Balances_MessageQueued {
    out := &pbgear.Balances_MessageQueued{}
    v := in.(registry.Valuable)
    _ = v

    // field Id To_GprimitivesMessageId(w)
    out.Id = To_GprimitivesMessageId(v.ValueAt(0))
    // field Source To_SpCoreCryptoAccountId32(w)
    out.Source = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // field Destination To_GprimitivesActorId(w)
    out.Destination = To_GprimitivesActorId(v.ValueAt(2))
    // oneOf field Entry
    v3 := To_OneOf_Balances_MessageQueued_entry(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Entry").Set(reflect.ValueOf(v3))

    return out //from message
}
    



func To_OneOf_Balances_MessageQueued_entry (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_MessageQueued_EntryInit{
        EntryInit: To_Init(param),
        }
    case 1:
        return &pbgear.Balances_MessageQueued_EntryHandle{
        EntryHandle: To_Handle(param),
        }
    case 2:
        return &pbgear.Balances_MessageQueued_EntryReply{
        EntryReply: To_Reply(param),
        }
    case 3:
        return &pbgear.Balances_MessageQueued_EntrySignal{
        EntrySignal: To_Signal(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_MessageWaited(in any) *pbgear.Balances_MessageWaited {
    out := &pbgear.Balances_MessageWaited{}
    v := in.(registry.Valuable)
    _ = v

    // field Id To_GprimitivesMessageId(w)
    out.Id = To_GprimitivesMessageId(v.ValueAt(0))
    // oneOf field Origin
    v1 := To_OneOf_Balances_MessageWaited_origin(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Origin").Set(reflect.ValueOf(v1))
    // oneOf field Reason
    v2 := To_OneOf_Balances_MessageWaited_reason(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Reason").Set(reflect.ValueOf(v2))
    // primitive field Expiration
    out.Expiration = To_uint32(v.ValueAt(3))

    return out //from message
}
    

func To_OneOf_Balances_MessageWaited_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_MessageWaited_OriginNode{
        OriginNode: To_Node(param),
        }
    case 1:
        return &pbgear.Balances_MessageWaited_OriginReservation{
        OriginReservation: To_Reservation(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Balances_MessageWaited_reason (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_MessageWaited_ReasonRuntime{
        ReasonRuntime: To_Runtime(param),
        }
    case 1:
        return &pbgear.Balances_MessageWaited_ReasonSystem{
        ReasonSystem: To_System(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_MessageWoken(in any) *pbgear.Balances_MessageWoken {
    out := &pbgear.Balances_MessageWoken{}
    v := in.(registry.Valuable)
    _ = v

    // field Id To_GprimitivesMessageId(w)
    out.Id = To_GprimitivesMessageId(v.ValueAt(0))
    // oneOf field Reason
    v1 := To_OneOf_Balances_MessageWoken_reason(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Reason").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_Balances_MessageWoken_reason (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_MessageWoken_ReasonRuntime{
        ReasonRuntime: To_Runtime(param),
        }
    case 1:
        return &pbgear.Balances_MessageWoken_ReasonSystem{
        ReasonSystem: To_System(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_MessagesDispatched(in any) *pbgear.Balances_MessagesDispatched {
    out := &pbgear.Balances_MessagesDispatched{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Total
    out.Total = To_uint32(v.ValueAt(0))
    // field Statuses To_BTreeMap(w)
    out.Statuses = To_BTreeMap(v.ValueAt(1))
    // field StateChanges To_BTreeSet(w)
    out.StateChanges = To_BTreeSet(v.ValueAt(2))

    return out //from message
}
    


func To_Balances_MetadataCleared(in any) *pbgear.Balances_MetadataCleared {
    out := &pbgear.Balances_MetadataCleared{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_MetadataSet(in any) *pbgear.Balances_MetadataSet {
    out := &pbgear.Balances_MetadataSet{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_MinBalanceDeficitAdjusted(in any) *pbgear.Balances_MinBalanceDeficitAdjusted {
    out := &pbgear.Balances_MinBalanceDeficitAdjusted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_MinBalanceExcessAdjusted(in any) *pbgear.Balances_MinBalanceExcessAdjusted {
    out := &pbgear.Balances_MinBalanceExcessAdjusted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_Minted(in any) *pbgear.Balances_Minted {
    out := &pbgear.Balances_Minted{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_MultisigApproval(in any) *pbgear.Balances_MultisigApproval {
    out := &pbgear.Balances_MultisigApproval{}
    v := in.(registry.Valuable)
    _ = v

    // field Approving To_SpCoreCryptoAccountId32(w)
    out.Approving = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Timepoint To_Multisig_PalletMultisigTimepoint(w)
    out.Timepoint = To_Multisig_PalletMultisigTimepoint(v.ValueAt(1))
    // field Multisig To_SpCoreCryptoAccountId32(w)
    out.Multisig = To_SpCoreCryptoAccountId32(v.ValueAt(2))
    // primitive field CallHash
    out.CallHash = To_bytes(v.ValueAt(3))

    return out //from message
}
    



func To_Balances_MultisigCancelled(in any) *pbgear.Balances_MultisigCancelled {
    out := &pbgear.Balances_MultisigCancelled{}
    v := in.(registry.Valuable)
    _ = v

    // field Cancelling To_SpCoreCryptoAccountId32(w)
    out.Cancelling = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Timepoint To_Multisig_PalletMultisigTimepoint(w)
    out.Timepoint = To_Multisig_PalletMultisigTimepoint(v.ValueAt(1))
    // field Multisig To_SpCoreCryptoAccountId32(w)
    out.Multisig = To_SpCoreCryptoAccountId32(v.ValueAt(2))
    // primitive field CallHash
    out.CallHash = To_bytes(v.ValueAt(3))

    return out //from message
}
    



func To_Balances_MultisigEvent(in any) *pbgear.Balances_MultisigEvent {
    out := &pbgear.Balances_MultisigEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_MultisigEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_MultisigEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_MultisigEvent_Value0NewMultisig{
        Value0NewMultisig: To_Balances_NewMultisig(param),
        }
    case 1:
        return &pbgear.Balances_MultisigEvent_Value0MultisigApproval{
        Value0MultisigApproval: To_Balances_MultisigApproval(param),
        }
    case 2:
        return &pbgear.Balances_MultisigEvent_Value0MultisigExecuted{
        Value0MultisigExecuted: To_Balances_MultisigExecuted(param),
        }
    case 3:
        return &pbgear.Balances_MultisigEvent_Value0MultisigCancelled{
        Value0MultisigCancelled: To_Balances_MultisigCancelled(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_MultisigExecuted(in any) *pbgear.Balances_MultisigExecuted {
    out := &pbgear.Balances_MultisigExecuted{}
    v := in.(registry.Valuable)
    _ = v

    // field Approving To_SpCoreCryptoAccountId32(w)
    out.Approving = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Timepoint To_Multisig_PalletMultisigTimepoint(w)
    out.Timepoint = To_Multisig_PalletMultisigTimepoint(v.ValueAt(1))
    // field Multisig To_SpCoreCryptoAccountId32(w)
    out.Multisig = To_SpCoreCryptoAccountId32(v.ValueAt(2))
    // primitive field CallHash
    out.CallHash = To_bytes(v.ValueAt(3))
    // oneOf field Result
    v4 := To_OneOf_Balances_MultisigExecuted_result(v.ValueAt(4))
    reflect.ValueOf(out).Elem().FieldByName("Result").Set(reflect.ValueOf(v4))

    return out //from message
}
    



func To_OneOf_Balances_MultisigExecuted_result (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_MultisigExecuted_ResultOk{
        ResultOk: To_Ok(param),
        }
    case 1:
        return &pbgear.Balances_MultisigExecuted_ResultErr{
        ResultErr: To_Err(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_NewAccount(in any) *pbgear.Balances_NewAccount {
    out := &pbgear.Balances_NewAccount{}
    v := in.(registry.Valuable)
    _ = v

    // field Account To_SpCoreCryptoAccountId32(w)
    out.Account = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_NewAuthorities(in any) *pbgear.Balances_NewAuthorities {
    out := &pbgear.Balances_NewAuthorities{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field AuthoritySet
    out.AuthoritySet = To_Repeated_Balances_NewAuthorities_authority_set(v.ValueAt(0))

    return out //from message
}
    

func To_Repeated_Balances_NewAuthorities_authority_set(in any) []*pbgear.Balances_TupleSpConsensusGrandpaAppPublicuint64 {
    items := in.([]interface{})

    var out []*pbgear.Balances_TupleSpConsensusGrandpaAppPublicuint64
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_Balances_TupleSpConsensusGrandpaAppPublicuint64(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Balances_NewMultisig(in any) *pbgear.Balances_NewMultisig {
    out := &pbgear.Balances_NewMultisig{}
    v := in.(registry.Valuable)
    _ = v

    // field Approving To_SpCoreCryptoAccountId32(w)
    out.Approving = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Multisig To_SpCoreCryptoAccountId32(w)
    out.Multisig = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field CallHash
    out.CallHash = To_bytes(v.ValueAt(2))

    return out //from message
}
    


func To_Balances_NewSession(in any) *pbgear.Balances_NewSession {
    out := &pbgear.Balances_NewSession{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field SessionIndex
    out.SessionIndex = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_NominationPoolsEvent(in any) *pbgear.Balances_NominationPoolsEvent {
    out := &pbgear.Balances_NominationPoolsEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_NominationPoolsEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_NominationPoolsEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_NominationPoolsEvent_Value0Created{
        Value0Created: To_Balances_Created(param),
        }
    case 1:
        return &pbgear.Balances_NominationPoolsEvent_Value0Bonded{
        Value0Bonded: To_Balances_Bonded(param),
        }
    case 2:
        return &pbgear.Balances_NominationPoolsEvent_Value0PaidOut{
        Value0PaidOut: To_Balances_PaidOut(param),
        }
    case 3:
        return &pbgear.Balances_NominationPoolsEvent_Value0Unbonded{
        Value0Unbonded: To_Balances_Unbonded(param),
        }
    case 4:
        return &pbgear.Balances_NominationPoolsEvent_Value0Withdrawn{
        Value0Withdrawn: To_Balances_Withdrawn(param),
        }
    case 5:
        return &pbgear.Balances_NominationPoolsEvent_Value0Destroyed{
        Value0Destroyed: To_Balances_Destroyed(param),
        }
    case 6:
        return &pbgear.Balances_NominationPoolsEvent_Value0StateChanged{
        Value0StateChanged: To_Balances_StateChanged(param),
        }
    case 7:
        return &pbgear.Balances_NominationPoolsEvent_Value0MemberRemoved{
        Value0MemberRemoved: To_Balances_MemberRemoved(param),
        }
    case 8:
        return &pbgear.Balances_NominationPoolsEvent_Value0RolesUpdated{
        Value0RolesUpdated: To_Balances_RolesUpdated(param),
        }
    case 9:
        return &pbgear.Balances_NominationPoolsEvent_Value0PoolSlashed{
        Value0PoolSlashed: To_Balances_PoolSlashed(param),
        }
    case 10:
        return &pbgear.Balances_NominationPoolsEvent_Value0UnbondingPoolSlashed{
        Value0UnbondingPoolSlashed: To_Balances_UnbondingPoolSlashed(param),
        }
    case 11:
        return &pbgear.Balances_NominationPoolsEvent_Value0PoolCommissionUpdated{
        Value0PoolCommissionUpdated: To_Balances_PoolCommissionUpdated(param),
        }
    case 12:
        return &pbgear.Balances_NominationPoolsEvent_Value0PoolMaxCommissionUpdated{
        Value0PoolMaxCommissionUpdated: To_Balances_PoolMaxCommissionUpdated(param),
        }
    case 13:
        return &pbgear.Balances_NominationPoolsEvent_Value0PoolCommissionChangeRateUpdated{
        Value0PoolCommissionChangeRateUpdated: To_Balances_PoolCommissionChangeRateUpdated(param),
        }
    case 14:
        return &pbgear.Balances_NominationPoolsEvent_Value0PoolCommissionClaimed{
        Value0PoolCommissionClaimed: To_Balances_PoolCommissionClaimed(param),
        }
    case 15:
        return &pbgear.Balances_NominationPoolsEvent_Value0MinBalanceDeficitAdjusted{
        Value0MinBalanceDeficitAdjusted: To_Balances_MinBalanceDeficitAdjusted(param),
        }
    case 16:
        return &pbgear.Balances_NominationPoolsEvent_Value0MinBalanceExcessAdjusted{
        Value0MinBalanceExcessAdjusted: To_Balances_MinBalanceExcessAdjusted(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Noted(in any) *pbgear.Balances_Noted {
    out := &pbgear.Balances_Noted{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Offence(in any) *pbgear.Balances_Offence {
    out := &pbgear.Balances_Offence{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Kind
    out.Kind = To_bytes(v.ValueAt(0))
    // primitive field Timeslot
    out.Timeslot = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_OffencesEvent(in any) *pbgear.Balances_OffencesEvent {
    out := &pbgear.Balances_OffencesEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_OffencesEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_OffencesEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_OffencesEvent_Value0Offence{
        Value0Offence: To_Balances_Offence(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_OldSlashingReportDiscarded(in any) *pbgear.Balances_OldSlashingReportDiscarded {
    out := &pbgear.Balances_OldSlashingReportDiscarded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field SessionIndex
    out.SessionIndex = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_OtherEvent(in any) *pbgear.Balances_OtherEvent {
    out := &pbgear.Balances_OtherEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Paid(in any) *pbgear.Balances_Paid {
    out := &pbgear.Balances_Paid{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field PaymentId To_Balances_TupleNull(w)
    out.PaymentId = To_Balances_TupleNull(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_PaidOut(in any) *pbgear.Balances_PaidOut {
    out := &pbgear.Balances_PaidOut{}
    v := in.(registry.Valuable)
    _ = v

    // field Member To_SpCoreCryptoAccountId32(w)
    out.Member = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(1))
    // primitive field Payout
    out.Payout = To_string(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_Paused(in any) *pbgear.Balances_Paused {
    out := &pbgear.Balances_Paused{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_PaymentFailed(in any) *pbgear.Balances_PaymentFailed {
    out := &pbgear.Balances_PaymentFailed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field PaymentId To_Balances_TupleNull(w)
    out.PaymentId = To_Balances_TupleNull(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_PayoutStarted(in any) *pbgear.Balances_PayoutStarted {
    out := &pbgear.Balances_PayoutStarted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field EraIndex
    out.EraIndex = To_uint32(v.ValueAt(0))
    // field ValidatorStash To_SpCoreCryptoAccountId32(w)
    out.ValidatorStash = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_PeriodicFailed(in any) *pbgear.Balances_PeriodicFailed {
    out := &pbgear.Balances_PeriodicFailed{}
    v := in.(registry.Valuable)
    _ = v

    // field Task To_Balances_TupleUint32Uint32(w)
    out.Task = To_Balances_TupleUint32Uint32(v.ValueAt(0))
    // primitive field Id
    out.Id = To_bytes(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_PermanentlyOverweight(in any) *pbgear.Balances_PermanentlyOverweight {
    out := &pbgear.Balances_PermanentlyOverweight{}
    v := in.(registry.Valuable)
    _ = v

    // field Task To_Balances_TupleUint32Uint32(w)
    out.Task = To_Balances_TupleUint32Uint32(v.ValueAt(0))
    // primitive field Id
    out.Id = To_bytes(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_PhaseTransitioned(in any) *pbgear.Balances_PhaseTransitioned {
    out := &pbgear.Balances_PhaseTransitioned{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field From
    v0 := To_OneOf_Balances_PhaseTransitioned_from(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("From").Set(reflect.ValueOf(v0))
    // oneOf field To
    v1 := To_OneOf_Balances_PhaseTransitioned_to(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v1))
    // primitive field Round
    out.Round = To_uint32(v.ValueAt(2))

    return out //from message
}
    
func To_OneOf_Balances_PhaseTransitioned_from (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_PhaseTransitioned_FromOff{
        FromOff: To_ElectionProviderMultiPhase_Off(param),
        }
    case 1:
        return &pbgear.Balances_PhaseTransitioned_FromSigned{
        FromSigned: To_ElectionProviderMultiPhase_Signed(param),
        }
    case 2:
        return &pbgear.Balances_PhaseTransitioned_FromUnsigned{
        FromUnsigned: To_ElectionProviderMultiPhase_Unsigned(param),
        }
    case 3:
        return &pbgear.Balances_PhaseTransitioned_FromEmergency{
        FromEmergency: To_ElectionProviderMultiPhase_Emergency(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Balances_PhaseTransitioned_to (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_PhaseTransitioned_ToOff{
        ToOff: To_ElectionProviderMultiPhase_Off(param),
        }
    case 1:
        return &pbgear.Balances_PhaseTransitioned_ToSigned{
        ToSigned: To_ElectionProviderMultiPhase_Signed(param),
        }
    case 2:
        return &pbgear.Balances_PhaseTransitioned_ToUnsigned{
        ToUnsigned: To_ElectionProviderMultiPhase_Unsigned(param),
        }
    case 3:
        return &pbgear.Balances_PhaseTransitioned_ToEmergency{
        ToEmergency: To_ElectionProviderMultiPhase_Emergency(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_PoolCommissionChangeRateUpdated(in any) *pbgear.Balances_PoolCommissionChangeRateUpdated {
    out := &pbgear.Balances_PoolCommissionChangeRateUpdated{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // field ChangeRate To_NominationPools_PalletNominationPoolsCommissionChangeRate(w)
    out.ChangeRate = To_NominationPools_PalletNominationPoolsCommissionChangeRate(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_PoolCommissionClaimed(in any) *pbgear.Balances_PoolCommissionClaimed {
    out := &pbgear.Balances_PoolCommissionClaimed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // primitive field Commission
    out.Commission = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_PoolCommissionUpdated(in any) *pbgear.Balances_PoolCommissionUpdated {
    out := &pbgear.Balances_PoolCommissionUpdated{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // optional field Current
    if v.HasValue() {
        out.Current = To_Optional_Balances_PoolCommissionUpdated_current(v.ValueAt(1))
    }

    return out //from message
}
    
func To_Optional_Balances_PoolCommissionUpdated_current(in any) *pbgear.TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_Balances_PoolMaxCommissionUpdated(in any) *pbgear.Balances_PoolMaxCommissionUpdated {
    out := &pbgear.Balances_PoolMaxCommissionUpdated{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // field MaxCommission To_SpArithmeticPerThingsPerbill(w)
    out.MaxCommission = To_SpArithmeticPerThingsPerbill(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_PoolSlashed(in any) *pbgear.Balances_PoolSlashed {
    out := &pbgear.Balances_PoolSlashed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // primitive field Balance
    out.Balance = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_PreRuntimeEvent(in any) *pbgear.Balances_PreRuntimeEvent {
    out := &pbgear.Balances_PreRuntimeEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_PreimageEvent(in any) *pbgear.Balances_PreimageEvent {
    out := &pbgear.Balances_PreimageEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_PreimageEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_PreimageEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_PreimageEvent_Value0Noted{
        Value0Noted: To_Balances_Noted(param),
        }
    case 1:
        return &pbgear.Balances_PreimageEvent_Value0Requested{
        Value0Requested: To_Balances_Requested(param),
        }
    case 2:
        return &pbgear.Balances_PreimageEvent_Value0Cleared{
        Value0Cleared: To_Balances_Cleared(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ProgramChanged(in any) *pbgear.Balances_ProgramChanged {
    out := &pbgear.Balances_ProgramChanged{}
    v := in.(registry.Valuable)
    _ = v

    // field Id To_GprimitivesActorId(w)
    out.Id = To_GprimitivesActorId(v.ValueAt(0))
    // oneOf field Change
    v1 := To_OneOf_Balances_ProgramChanged_change(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Change").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_Balances_ProgramChanged_change (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ProgramChanged_ChangeActive{
        ChangeActive: To_Active(param),
        }
    case 1:
        return &pbgear.Balances_ProgramChanged_ChangeInactive{
        ChangeInactive: To_Inactive(param),
        }
    case 2:
        return &pbgear.Balances_ProgramChanged_ChangePaused{
        ChangePaused: To_Paused(param),
        }
    case 3:
        return &pbgear.Balances_ProgramChanged_ChangeTerminated{
        ChangeTerminated: To_Terminated(param),
        }
    case 4:
        return &pbgear.Balances_ProgramChanged_ChangeExpirationChanged{
        ChangeExpirationChanged: To_ExpirationChanged(param),
        }
    case 5:
        return &pbgear.Balances_ProgramChanged_ChangeProgramSet{
        ChangeProgramSet: To_ProgramSet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Proposed(in any) *pbgear.Balances_Proposed {
    out := &pbgear.Balances_Proposed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ProposalIndex
    out.ProposalIndex = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_ProxyAdded(in any) *pbgear.Balances_ProxyAdded {
    out := &pbgear.Balances_ProxyAdded{}
    v := in.(registry.Valuable)
    _ = v

    // field Delegator To_SpCoreCryptoAccountId32(w)
    out.Delegator = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Delegatee To_SpCoreCryptoAccountId32(w)
    out.Delegatee = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // oneOf field ProxyType
    v2 := To_OneOf_Balances_ProxyAdded_proxy_type(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v2))
    // primitive field Delay
    out.Delay = To_uint32(v.ValueAt(3))

    return out //from message
}
    


func To_OneOf_Balances_ProxyAdded_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ProxyAdded_ProxyTypeAny{
        ProxyTypeAny: To_Any(param),
        }
    case 1:
        return &pbgear.Balances_ProxyAdded_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_NonTransfer(param),
        }
    case 2:
        return &pbgear.Balances_ProxyAdded_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Governance(param),
        }
    case 3:
        return &pbgear.Balances_ProxyAdded_ProxyTypeStaking{
        ProxyTypeStaking: To_Staking(param),
        }
    case 4:
        return &pbgear.Balances_ProxyAdded_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Balances_ProxyAdded_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ProxyEvent(in any) *pbgear.Balances_ProxyEvent {
    out := &pbgear.Balances_ProxyEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_ProxyEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ProxyEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ProxyEvent_Value0ProxyExecuted{
        Value0ProxyExecuted: To_Balances_ProxyExecuted(param),
        }
    case 1:
        return &pbgear.Balances_ProxyEvent_Value0PureCreated{
        Value0PureCreated: To_Balances_PureCreated(param),
        }
    case 2:
        return &pbgear.Balances_ProxyEvent_Value0Announced{
        Value0Announced: To_Balances_Announced(param),
        }
    case 3:
        return &pbgear.Balances_ProxyEvent_Value0ProxyAdded{
        Value0ProxyAdded: To_Balances_ProxyAdded(param),
        }
    case 4:
        return &pbgear.Balances_ProxyEvent_Value0ProxyRemoved{
        Value0ProxyRemoved: To_Balances_ProxyRemoved(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ProxyExecuted(in any) *pbgear.Balances_ProxyExecuted {
    out := &pbgear.Balances_ProxyExecuted{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Result
    v0 := To_OneOf_Balances_ProxyExecuted_result(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Result").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ProxyExecuted_result (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ProxyExecuted_ResultOk{
        ResultOk: To_Ok(param),
        }
    case 1:
        return &pbgear.Balances_ProxyExecuted_ResultErr{
        ResultErr: To_Err(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ProxyRemoved(in any) *pbgear.Balances_ProxyRemoved {
    out := &pbgear.Balances_ProxyRemoved{}
    v := in.(registry.Valuable)
    _ = v

    // field Delegator To_SpCoreCryptoAccountId32(w)
    out.Delegator = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Delegatee To_SpCoreCryptoAccountId32(w)
    out.Delegatee = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // oneOf field ProxyType
    v2 := To_OneOf_Balances_ProxyRemoved_proxy_type(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v2))
    // primitive field Delay
    out.Delay = To_uint32(v.ValueAt(3))

    return out //from message
}
    


func To_OneOf_Balances_ProxyRemoved_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ProxyRemoved_ProxyTypeAny{
        ProxyTypeAny: To_Any(param),
        }
    case 1:
        return &pbgear.Balances_ProxyRemoved_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_NonTransfer(param),
        }
    case 2:
        return &pbgear.Balances_ProxyRemoved_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Governance(param),
        }
    case 3:
        return &pbgear.Balances_ProxyRemoved_ProxyTypeStaking{
        ProxyTypeStaking: To_Staking(param),
        }
    case 4:
        return &pbgear.Balances_ProxyRemoved_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Balances_ProxyRemoved_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_PureCreated(in any) *pbgear.Balances_PureCreated {
    out := &pbgear.Balances_PureCreated{}
    v := in.(registry.Valuable)
    _ = v

    // field Pure To_SpCoreCryptoAccountId32(w)
    out.Pure = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // oneOf field ProxyType
    v2 := To_OneOf_Balances_PureCreated_proxy_type(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v2))
    // primitive field DisambiguationIndex
    out.DisambiguationIndex = To_uint32(v.ValueAt(3))

    return out //from message
}
    


func To_OneOf_Balances_PureCreated_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_PureCreated_ProxyTypeAny{
        ProxyTypeAny: To_Any(param),
        }
    case 1:
        return &pbgear.Balances_PureCreated_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_NonTransfer(param),
        }
    case 2:
        return &pbgear.Balances_PureCreated_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Governance(param),
        }
    case 3:
        return &pbgear.Balances_PureCreated_ProxyTypeStaking{
        ProxyTypeStaking: To_Staking(param),
        }
    case 4:
        return &pbgear.Balances_PureCreated_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Balances_PureCreated_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_QueueNotProcessed(in any) *pbgear.Balances_QueueNotProcessed {
    out := &pbgear.Balances_QueueNotProcessed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_RankChanged(in any) *pbgear.Balances_RankChanged {
    out := &pbgear.Balances_RankChanged{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Rank
    out.Rank = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Raw(in any) *pbgear.Balances_Raw {
    out := &pbgear.Balances_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Rebagged(in any) *pbgear.Balances_Rebagged {
    out := &pbgear.Balances_Rebagged{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field From
    out.From = To_uint64(v.ValueAt(1))
    // primitive field To
    out.To = To_uint64(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_ReferendaEvent(in any) *pbgear.Balances_ReferendaEvent {
    out := &pbgear.Balances_ReferendaEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_ReferendaEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_ReferendaEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ReferendaEvent_Value0Submitted{
        Value0Submitted: To_Balances_Submitted(param),
        }
    case 1:
        return &pbgear.Balances_ReferendaEvent_Value0DecisionDepositPlaced{
        Value0DecisionDepositPlaced: To_Balances_DecisionDepositPlaced(param),
        }
    case 2:
        return &pbgear.Balances_ReferendaEvent_Value0DecisionDepositRefunded{
        Value0DecisionDepositRefunded: To_Balances_DecisionDepositRefunded(param),
        }
    case 3:
        return &pbgear.Balances_ReferendaEvent_Value0DepositSlashed{
        Value0DepositSlashed: To_Balances_DepositSlashed(param),
        }
    case 4:
        return &pbgear.Balances_ReferendaEvent_Value0DecisionStarted{
        Value0DecisionStarted: To_Balances_DecisionStarted(param),
        }
    case 5:
        return &pbgear.Balances_ReferendaEvent_Value0ConfirmStarted{
        Value0ConfirmStarted: To_Balances_ConfirmStarted(param),
        }
    case 6:
        return &pbgear.Balances_ReferendaEvent_Value0ConfirmAborted{
        Value0ConfirmAborted: To_Balances_ConfirmAborted(param),
        }
    case 7:
        return &pbgear.Balances_ReferendaEvent_Value0Confirmed{
        Value0Confirmed: To_Balances_Confirmed(param),
        }
    case 8:
        return &pbgear.Balances_ReferendaEvent_Value0Approved{
        Value0Approved: To_Balances_Approved(param),
        }
    case 9:
        return &pbgear.Balances_ReferendaEvent_Value0Rejected{
        Value0Rejected: To_Balances_Rejected(param),
        }
    case 10:
        return &pbgear.Balances_ReferendaEvent_Value0TimedOut{
        Value0TimedOut: To_Balances_TimedOut(param),
        }
    case 11:
        return &pbgear.Balances_ReferendaEvent_Value0Cancelled{
        Value0Cancelled: To_Balances_Cancelled(param),
        }
    case 12:
        return &pbgear.Balances_ReferendaEvent_Value0Killed{
        Value0Killed: To_Balances_Killed(param),
        }
    case 13:
        return &pbgear.Balances_ReferendaEvent_Value0SubmissionDepositRefunded{
        Value0SubmissionDepositRefunded: To_Balances_SubmissionDepositRefunded(param),
        }
    case 14:
        return &pbgear.Balances_ReferendaEvent_Value0MetadataSet{
        Value0MetadataSet: To_Balances_MetadataSet(param),
        }
    case 15:
        return &pbgear.Balances_ReferendaEvent_Value0MetadataCleared{
        Value0MetadataCleared: To_Balances_MetadataCleared(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_RegistrarAdded(in any) *pbgear.Balances_RegistrarAdded {
    out := &pbgear.Balances_RegistrarAdded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field RegistrarIndex
    out.RegistrarIndex = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Rejected(in any) *pbgear.Balances_Rejected {
    out := &pbgear.Balances_Rejected{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ProposalIndex
    out.ProposalIndex = To_uint32(v.ValueAt(0))
    // primitive field Slashed
    out.Slashed = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_Remarked(in any) *pbgear.Balances_Remarked {
    out := &pbgear.Balances_Remarked{}
    v := in.(registry.Valuable)
    _ = v

    // field Sender To_SpCoreCryptoAccountId32(w)
    out.Sender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(1))

    return out //from message
}
    


func To_Balances_Requested(in any) *pbgear.Balances_Requested {
    out := &pbgear.Balances_Requested{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Rescinded(in any) *pbgear.Balances_Rescinded {
    out := &pbgear.Balances_Rescinded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Amount
    out.Amount = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_ReserveRepatriated(in any) *pbgear.Balances_ReserveRepatriated {
    out := &pbgear.Balances_ReserveRepatriated{}
    v := in.(registry.Valuable)
    _ = v

    // field From To_SpCoreCryptoAccountId32(w)
    out.From = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field To To_SpCoreCryptoAccountId32(w)
    out.To = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(2))
    // oneOf field DestinationStatus
    v3 := To_OneOf_Balances_ReserveRepatriated_destination_status(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("DestinationStatus").Set(reflect.ValueOf(v3))

    return out //from message
}
    


func To_OneOf_Balances_ReserveRepatriated_destination_status (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_ReserveRepatriated_DestinationStatusFree{
        DestinationStatusFree: To_Free(param),
        }
    case 1:
        return &pbgear.Balances_ReserveRepatriated_DestinationStatusReserved{
        DestinationStatusReserved: To_Reserved(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Reserved(in any) *pbgear.Balances_Reserved {
    out := &pbgear.Balances_Reserved{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Restored(in any) *pbgear.Balances_Restored {
    out := &pbgear.Balances_Restored{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Resumed(in any) *pbgear.Balances_Resumed {
    out := &pbgear.Balances_Resumed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_Rewarded(in any) *pbgear.Balances_Rewarded {
    out := &pbgear.Balances_Rewarded{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // oneOf field Dest
    v1 := To_OneOf_Balances_Rewarded_dest(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v1))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(2))

    return out //from message
}
    

func To_OneOf_Balances_Rewarded_dest (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Rewarded_DestStaked{
        DestStaked: To_Staking_Staked(param),
        }
    case 1:
        return &pbgear.Balances_Rewarded_DestStash{
        DestStash: To_Staking_Stash(param),
        }
    case 2:
        return &pbgear.Balances_Rewarded_DestController{
        DestController: To_Staking_Controller(param),
        }
    case 3:
        return &pbgear.Balances_Rewarded_DestAccount{
        DestAccount: To_Staking_Account(param),
        }
    case 4:
        return &pbgear.Balances_Rewarded_DestNone{
        DestNone: To_Staking_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_RolesUpdated(in any) *pbgear.Balances_RolesUpdated {
    out := &pbgear.Balances_RolesUpdated{}
    v := in.(registry.Valuable)
    _ = v

    // optional field Root
    if v.HasValue() {
        out.Root = To_Optional_Balances_RolesUpdated_root(v.ValueAt(0))
    }
    // optional field Bouncer
    if v.HasValue() {
        out.Bouncer = To_Optional_Balances_RolesUpdated_bouncer(v.ValueAt(1))
    }
    // optional field Nominator
    if v.HasValue() {
        out.Nominator = To_Optional_Balances_RolesUpdated_nominator(v.ValueAt(2))
    }

    return out //from message
}
    
func To_Optional_Balances_RolesUpdated_root(in any) *pbgear.SpCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_Optional_Balances_RolesUpdated_bouncer(in any) *pbgear.SpCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_Optional_Balances_RolesUpdated_nominator(in any) *pbgear.SpCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_Balances_Rollover(in any) *pbgear.Balances_Rollover {
    out := &pbgear.Balances_Rollover{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field RolloverBalance
    out.RolloverBalance = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_RuntimeEnvironmentUpdatedEvent(in any) *pbgear.Balances_RuntimeEnvironmentUpdatedEvent {
    out := &pbgear.Balances_RuntimeEnvironmentUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_Scheduled(in any) *pbgear.Balances_Scheduled {
    out := &pbgear.Balances_Scheduled{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field When
    out.When = To_uint32(v.ValueAt(0))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_SchedulerEvent(in any) *pbgear.Balances_SchedulerEvent {
    out := &pbgear.Balances_SchedulerEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_SchedulerEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_SchedulerEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_SchedulerEvent_Value0Scheduled{
        Value0Scheduled: To_Balances_Scheduled(param),
        }
    case 1:
        return &pbgear.Balances_SchedulerEvent_Value0Canceled{
        Value0Canceled: To_Balances_Canceled(param),
        }
    case 2:
        return &pbgear.Balances_SchedulerEvent_Value0Dispatched{
        Value0Dispatched: To_Balances_Dispatched(param),
        }
    case 3:
        return &pbgear.Balances_SchedulerEvent_Value0CallUnavailable{
        Value0CallUnavailable: To_Balances_CallUnavailable(param),
        }
    case 4:
        return &pbgear.Balances_SchedulerEvent_Value0PeriodicFailed{
        Value0PeriodicFailed: To_Balances_PeriodicFailed(param),
        }
    case 5:
        return &pbgear.Balances_SchedulerEvent_Value0PermanentlyOverweight{
        Value0PermanentlyOverweight: To_Balances_PermanentlyOverweight(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ScoreUpdated(in any) *pbgear.Balances_ScoreUpdated {
    out := &pbgear.Balances_ScoreUpdated{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field NewScore
    out.NewScore = To_uint64(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_SealEvent(in any) *pbgear.Balances_SealEvent {
    out := &pbgear.Balances_SealEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_SessionEvent(in any) *pbgear.Balances_SessionEvent {
    out := &pbgear.Balances_SessionEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_SessionEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_SessionEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_SessionEvent_Value0NewSession{
        Value0NewSession: To_Balances_NewSession(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_SlashReported(in any) *pbgear.Balances_SlashReported {
    out := &pbgear.Balances_SlashReported{}
    v := in.(registry.Valuable)
    _ = v

    // field Validator To_SpCoreCryptoAccountId32(w)
    out.Validator = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Fraction To_SpArithmeticPerThingsPerbill(w)
    out.Fraction = To_SpArithmeticPerThingsPerbill(v.ValueAt(1))
    // primitive field SlashEra
    out.SlashEra = To_uint32(v.ValueAt(2))

    return out //from message
}
    


func To_Balances_Slashed(in any) *pbgear.Balances_Slashed {
    out := &pbgear.Balances_Slashed{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_SnapshotTargetsSizeExceeded(in any) *pbgear.Balances_SnapshotTargetsSizeExceeded {
    out := &pbgear.Balances_SnapshotTargetsSizeExceeded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Size
    out.Size = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_SnapshotVotersSizeExceeded(in any) *pbgear.Balances_SnapshotVotersSizeExceeded {
    out := &pbgear.Balances_SnapshotVotersSizeExceeded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Size
    out.Size = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_SolutionStored(in any) *pbgear.Balances_SolutionStored {
    out := &pbgear.Balances_SolutionStored{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Compute
    v0 := To_OneOf_Balances_SolutionStored_compute(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Compute").Set(reflect.ValueOf(v0))
    // optional field Origin
    if v.HasValue() {
        out.Origin = To_Optional_Balances_SolutionStored_origin(v.ValueAt(1))
    }
    // primitive field PrevEjected
    out.PrevEjected = To_bool(v.ValueAt(2))

    return out //from message
}
    
func To_OneOf_Balances_SolutionStored_compute (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_SolutionStored_ComputeOnChain{
        ComputeOnChain: To_ElectionProviderMultiPhase_OnChain(param),
        }
    case 1:
        return &pbgear.Balances_SolutionStored_ComputeSigned{
        ComputeSigned: To_ElectionProviderMultiPhase_Signed(param),
        }
    case 2:
        return &pbgear.Balances_SolutionStored_ComputeUnsigned{
        ComputeUnsigned: To_ElectionProviderMultiPhase_Unsigned(param),
        }
    case 3:
        return &pbgear.Balances_SolutionStored_ComputeFallback{
        ComputeFallback: To_ElectionProviderMultiPhase_Fallback(param),
        }
    case 4:
        return &pbgear.Balances_SolutionStored_ComputeEmergency{
        ComputeEmergency: To_ElectionProviderMultiPhase_Emergency(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Optional_Balances_SolutionStored_origin(in any) *pbgear.SpCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_Balances_SomeOffline(in any) *pbgear.Balances_SomeOffline {
    out := &pbgear.Balances_SomeOffline{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Offline
    out.Offline = To_Repeated_Balances_SomeOffline_offline(v.ValueAt(0))

    return out //from message
}
    

func To_Repeated_Balances_SomeOffline_offline(in any) []*pbgear.Balances_TupleSpCoreCryptoAccountId32PalletStakingExposure {
    items := in.([]interface{})

    var out []*pbgear.Balances_TupleSpCoreCryptoAccountId32PalletStakingExposure
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_Balances_TupleSpCoreCryptoAccountId32PalletStakingExposure(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Balances_Source(in any) *pbgear.Balances_Source {
    out := &pbgear.Balances_Source{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Source
    v0 := To_OneOf_Balances_Source_source(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_Source_source (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Source_SourceId{
        SourceId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_Source_SourceIndex{
        SourceIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_Source_SourceRaw{
        SourceRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_Source_SourceAddress32{
        SourceAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_Source_SourceAddress20{
        SourceAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_SpendApproved(in any) *pbgear.Balances_SpendApproved {
    out := &pbgear.Balances_SpendApproved{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ProposalIndex
    out.ProposalIndex = To_uint32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))
    // field Beneficiary To_SpCoreCryptoAccountId32(w)
    out.Beneficiary = To_SpCoreCryptoAccountId32(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_SpendProcessed(in any) *pbgear.Balances_SpendProcessed {
    out := &pbgear.Balances_SpendProcessed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_Spending(in any) *pbgear.Balances_Spending {
    out := &pbgear.Balances_Spending{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BudgetRemaining
    out.BudgetRemaining = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Balances_StakersElected(in any) *pbgear.Balances_StakersElected {
    out := &pbgear.Balances_StakersElected{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_StakingElectionFailed(in any) *pbgear.Balances_StakingElectionFailed {
    out := &pbgear.Balances_StakingElectionFailed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Balances_StakingEvent(in any) *pbgear.Balances_StakingEvent {
    out := &pbgear.Balances_StakingEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_StakingEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_StakingEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_StakingEvent_Value0EraPaid{
        Value0EraPaid: To_Balances_EraPaid(param),
        }
    case 1:
        return &pbgear.Balances_StakingEvent_Value0Rewarded{
        Value0Rewarded: To_Balances_Rewarded(param),
        }
    case 2:
        return &pbgear.Balances_StakingEvent_Value0Slashed{
        Value0Slashed: To_Balances_Slashed(param),
        }
    case 3:
        return &pbgear.Balances_StakingEvent_Value0SlashReported{
        Value0SlashReported: To_Balances_SlashReported(param),
        }
    case 4:
        return &pbgear.Balances_StakingEvent_Value0OldSlashingReportDiscarded{
        Value0OldSlashingReportDiscarded: To_Balances_OldSlashingReportDiscarded(param),
        }
    case 5:
        return &pbgear.Balances_StakingEvent_Value0StakersElected{
        Value0StakersElected: To_Balances_StakersElected(param),
        }
    case 6:
        return &pbgear.Balances_StakingEvent_Value0Bonded{
        Value0Bonded: To_Balances_Bonded(param),
        }
    case 7:
        return &pbgear.Balances_StakingEvent_Value0Unbonded{
        Value0Unbonded: To_Balances_Unbonded(param),
        }
    case 8:
        return &pbgear.Balances_StakingEvent_Value0Withdrawn{
        Value0Withdrawn: To_Balances_Withdrawn(param),
        }
    case 9:
        return &pbgear.Balances_StakingEvent_Value0Kicked{
        Value0Kicked: To_Balances_Kicked(param),
        }
    case 10:
        return &pbgear.Balances_StakingEvent_Value0StakingElectionFailed{
        Value0StakingElectionFailed: To_Balances_StakingElectionFailed(param),
        }
    case 11:
        return &pbgear.Balances_StakingEvent_Value0Chilled{
        Value0Chilled: To_Balances_Chilled(param),
        }
    case 12:
        return &pbgear.Balances_StakingEvent_Value0PayoutStarted{
        Value0PayoutStarted: To_Balances_PayoutStarted(param),
        }
    case 13:
        return &pbgear.Balances_StakingEvent_Value0ValidatorPrefsSet{
        Value0ValidatorPrefsSet: To_Balances_ValidatorPrefsSet(param),
        }
    case 14:
        return &pbgear.Balances_StakingEvent_Value0SnapshotVotersSizeExceeded{
        Value0SnapshotVotersSizeExceeded: To_Balances_SnapshotVotersSizeExceeded(param),
        }
    case 15:
        return &pbgear.Balances_StakingEvent_Value0SnapshotTargetsSizeExceeded{
        Value0SnapshotTargetsSizeExceeded: To_Balances_SnapshotTargetsSizeExceeded(param),
        }
    case 16:
        return &pbgear.Balances_StakingEvent_Value0ForceEra{
        Value0ForceEra: To_Balances_ForceEra(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_StakingRewardsEvent(in any) *pbgear.Balances_StakingRewardsEvent {
    out := &pbgear.Balances_StakingRewardsEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_StakingRewardsEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_StakingRewardsEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_StakingRewardsEvent_Value0Deposited{
        Value0Deposited: To_Balances_Deposited(param),
        }
    case 1:
        return &pbgear.Balances_StakingRewardsEvent_Value0Withdrawn{
        Value0Withdrawn: To_Balances_Withdrawn(param),
        }
    case 2:
        return &pbgear.Balances_StakingRewardsEvent_Value0Burned{
        Value0Burned: To_Balances_Burned(param),
        }
    case 3:
        return &pbgear.Balances_StakingRewardsEvent_Value0Minted{
        Value0Minted: To_Balances_Minted(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_StateChanged(in any) *pbgear.Balances_StateChanged {
    out := &pbgear.Balances_StateChanged{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // oneOf field NewState
    v1 := To_OneOf_Balances_StateChanged_new_state(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("NewState").Set(reflect.ValueOf(v1))

    return out //from message
}
    
func To_OneOf_Balances_StateChanged_new_state (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_StateChanged_NewStateOpen{
        NewStateOpen: To_NominationPools_Open(param),
        }
    case 1:
        return &pbgear.Balances_StateChanged_NewStateBlocked{
        NewStateBlocked: To_NominationPools_Blocked(param),
        }
    case 2:
        return &pbgear.Balances_StateChanged_NewStateDestroying{
        NewStateDestroying: To_NominationPools_Destroying(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_SubIdentityAdded(in any) *pbgear.Balances_SubIdentityAdded {
    out := &pbgear.Balances_SubIdentityAdded{}
    v := in.(registry.Valuable)
    _ = v

    // field Sub To_SpCoreCryptoAccountId32(w)
    out.Sub = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Main To_SpCoreCryptoAccountId32(w)
    out.Main = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Deposit
    out.Deposit = To_string(v.ValueAt(2))

    return out //from message
}
    


func To_Balances_SubIdentityRemoved(in any) *pbgear.Balances_SubIdentityRemoved {
    out := &pbgear.Balances_SubIdentityRemoved{}
    v := in.(registry.Valuable)
    _ = v

    // field Sub To_SpCoreCryptoAccountId32(w)
    out.Sub = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Main To_SpCoreCryptoAccountId32(w)
    out.Main = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Deposit
    out.Deposit = To_string(v.ValueAt(2))

    return out //from message
}
    


func To_Balances_SubIdentityRevoked(in any) *pbgear.Balances_SubIdentityRevoked {
    out := &pbgear.Balances_SubIdentityRevoked{}
    v := in.(registry.Valuable)
    _ = v

    // field Sub To_SpCoreCryptoAccountId32(w)
    out.Sub = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Main To_SpCoreCryptoAccountId32(w)
    out.Main = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Deposit
    out.Deposit = To_string(v.ValueAt(2))

    return out //from message
}
    


func To_Balances_SubmissionDepositRefunded(in any) *pbgear.Balances_SubmissionDepositRefunded {
    out := &pbgear.Balances_SubmissionDepositRefunded{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_Submitted(in any) *pbgear.Balances_Submitted {
    out := &pbgear.Balances_Submitted{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field Track
    out.Track = To_uint32(v.ValueAt(1))
    // oneOf field Proposal
    v2 := To_OneOf_Balances_Submitted_proposal(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v2))

    return out //from message
}
    
func To_OneOf_Balances_Submitted_proposal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Submitted_ProposalLegacy{
        ProposalLegacy: To_Legacy(param),
        }
    case 1:
        return &pbgear.Balances_Submitted_ProposalInline{
        ProposalInline: To_Inline(param),
        }
    case 2:
        return &pbgear.Balances_Submitted_ProposalLookup{
        ProposalLookup: To_Lookup(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Suspended(in any) *pbgear.Balances_Suspended {
    out := &pbgear.Balances_Suspended{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_SystemEvent(in any) *pbgear.Balances_SystemEvent {
    out := &pbgear.Balances_SystemEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_SystemEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_SystemEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_SystemEvent_Value0ExtrinsicSuccess{
        Value0ExtrinsicSuccess: To_Balances_ExtrinsicSuccess(param),
        }
    case 1:
        return &pbgear.Balances_SystemEvent_Value0ExtrinsicFailed{
        Value0ExtrinsicFailed: To_Balances_ExtrinsicFailed(param),
        }
    case 2:
        return &pbgear.Balances_SystemEvent_Value0CodeUpdated{
        Value0CodeUpdated: To_Balances_CodeUpdated(param),
        }
    case 3:
        return &pbgear.Balances_SystemEvent_Value0NewAccount{
        Value0NewAccount: To_Balances_NewAccount(param),
        }
    case 4:
        return &pbgear.Balances_SystemEvent_Value0KilledAccount{
        Value0KilledAccount: To_Balances_KilledAccount(param),
        }
    case 5:
        return &pbgear.Balances_SystemEvent_Value0Remarked{
        Value0Remarked: To_Balances_Remarked(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Thawed(in any) *pbgear.Balances_Thawed {
    out := &pbgear.Balances_Thawed{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_TimedOut(in any) *pbgear.Balances_TimedOut {
    out := &pbgear.Balances_TimedOut{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Tally To_ConvictionVoting_PalletConvictionVotingTypesTally(w)
    out.Tally = To_ConvictionVoting_PalletConvictionVotingTypesTally(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_TransactionFeePaid(in any) *pbgear.Balances_TransactionFeePaid {
    out := &pbgear.Balances_TransactionFeePaid{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field ActualFee
    out.ActualFee = To_string(v.ValueAt(1))
    // primitive field Tip
    out.Tip = To_string(v.ValueAt(2))

    return out //from message
}
    

func To_Balances_TransactionPaymentEvent(in any) *pbgear.Balances_TransactionPaymentEvent {
    out := &pbgear.Balances_TransactionPaymentEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_TransactionPaymentEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_TransactionPaymentEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_TransactionPaymentEvent_Value0TransactionFeePaid{
        Value0TransactionFeePaid: To_Balances_TransactionFeePaid(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Transfer(in any) *pbgear.Balances_Transfer {
    out := &pbgear.Balances_Transfer{}
    v := in.(registry.Valuable)
    _ = v

    // field From To_SpCoreCryptoAccountId32(w)
    out.From = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field To To_SpCoreCryptoAccountId32(w)
    out.To = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(2))

    return out //from message
}
    


func To_Balances_TransferAllCall(in any) *pbgear.Balances_TransferAllCall {
    out := &pbgear.Balances_TransferAllCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Dest
    v0 := To_OneOf_Balances_TransferAllCall_dest(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))
    // primitive field KeepAlive
    out.KeepAlive = To_bool(v.ValueAt(1))

    return out //from message
}
func To_Balances_TransferAllCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Balances_TransferAllCall(in)
    out := &pbgearextrinsic.RuntimeCall_BalancesTransferAllCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BalancesTransferAllCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Balances_TransferAllCall_dest (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_TransferAllCall_DestId{
        DestId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_TransferAllCall_DestIndex{
        DestIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_TransferAllCall_DestRaw{
        DestRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_TransferAllCall_DestAddress32{
        DestAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_TransferAllCall_DestAddress20{
        DestAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_TransferAllowDeathCall(in any) *pbgear.Balances_TransferAllowDeathCall {
    out := &pbgear.Balances_TransferAllowDeathCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Dest
    v0 := To_OneOf_Balances_TransferAllowDeathCall_dest(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))
    // primitive field Value
    out.Value = To_string(v.ValueAt(1))

    return out //from message
}
func To_Balances_TransferAllowDeathCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Balances_TransferAllowDeathCall(in)
    out := &pbgearextrinsic.RuntimeCall_BalancesTransferAllowDeathCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BalancesTransferAllowDeathCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Balances_TransferAllowDeathCall_dest (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_TransferAllowDeathCall_DestId{
        DestId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_TransferAllowDeathCall_DestIndex{
        DestIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_TransferAllowDeathCall_DestRaw{
        DestRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_TransferAllowDeathCall_DestAddress32{
        DestAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_TransferAllowDeathCall_DestAddress20{
        DestAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_TransferKeepAliveCall(in any) *pbgear.Balances_TransferKeepAliveCall {
    out := &pbgear.Balances_TransferKeepAliveCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Dest
    v0 := To_OneOf_Balances_TransferKeepAliveCall_dest(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))
    // primitive field Value
    out.Value = To_string(v.ValueAt(1))

    return out //from message
}
func To_Balances_TransferKeepAliveCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Balances_TransferKeepAliveCall(in)
    out := &pbgearextrinsic.RuntimeCall_BalancesTransferKeepAliveCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BalancesTransferKeepAliveCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Balances_TransferKeepAliveCall_dest (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_TransferKeepAliveCall_DestId{
        DestId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_TransferKeepAliveCall_DestIndex{
        DestIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_TransferKeepAliveCall_DestRaw{
        DestRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_TransferKeepAliveCall_DestAddress32{
        DestAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_TransferKeepAliveCall_DestAddress20{
        DestAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_TreasuryEvent(in any) *pbgear.Balances_TreasuryEvent {
    out := &pbgear.Balances_TreasuryEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_TreasuryEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_TreasuryEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_TreasuryEvent_Value0Proposed{
        Value0Proposed: To_Balances_Proposed(param),
        }
    case 1:
        return &pbgear.Balances_TreasuryEvent_Value0Spending{
        Value0Spending: To_Balances_Spending(param),
        }
    case 2:
        return &pbgear.Balances_TreasuryEvent_Value0Awarded{
        Value0Awarded: To_Balances_Awarded(param),
        }
    case 3:
        return &pbgear.Balances_TreasuryEvent_Value0Rejected{
        Value0Rejected: To_Balances_Rejected(param),
        }
    case 4:
        return &pbgear.Balances_TreasuryEvent_Value0Burnt{
        Value0Burnt: To_Balances_Burnt(param),
        }
    case 5:
        return &pbgear.Balances_TreasuryEvent_Value0Rollover{
        Value0Rollover: To_Balances_Rollover(param),
        }
    case 6:
        return &pbgear.Balances_TreasuryEvent_Value0Deposit{
        Value0Deposit: To_Balances_Deposit(param),
        }
    case 7:
        return &pbgear.Balances_TreasuryEvent_Value0SpendApproved{
        Value0SpendApproved: To_Balances_SpendApproved(param),
        }
    case 8:
        return &pbgear.Balances_TreasuryEvent_Value0UpdatedInactive{
        Value0UpdatedInactive: To_Balances_UpdatedInactive(param),
        }
    case 9:
        return &pbgear.Balances_TreasuryEvent_Value0AssetSpendApproved{
        Value0AssetSpendApproved: To_Balances_AssetSpendApproved(param),
        }
    case 10:
        return &pbgear.Balances_TreasuryEvent_Value0AssetSpendVoided{
        Value0AssetSpendVoided: To_Balances_AssetSpendVoided(param),
        }
    case 11:
        return &pbgear.Balances_TreasuryEvent_Value0Paid{
        Value0Paid: To_Balances_Paid(param),
        }
    case 12:
        return &pbgear.Balances_TreasuryEvent_Value0PaymentFailed{
        Value0PaymentFailed: To_Balances_PaymentFailed(param),
        }
    case 13:
        return &pbgear.Balances_TreasuryEvent_Value0SpendProcessed{
        Value0SpendProcessed: To_Balances_SpendProcessed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_TupleNull(in any) *pbgear.Balances_TupleNull {
    out := &pbgear.Balances_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_TupleSpConsensusGrandpaAppPublicuint64(in any) *pbgear.Balances_TupleSpConsensusGrandpaAppPublicuint64 {
    out := &pbgear.Balances_TupleSpConsensusGrandpaAppPublicuint64{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpConsensusGrandpaAppPublic(w)
    out.Value0 = To_SpConsensusGrandpaAppPublic(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_uint64(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_TupleSpCoreCryptoAccountId32PalletStakingExposure(in any) *pbgear.Balances_TupleSpCoreCryptoAccountId32PalletStakingExposure {
    out := &pbgear.Balances_TupleSpCoreCryptoAccountId32PalletStakingExposure{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Value1 To_Staking_PalletStakingExposure(w)
    out.Value1 = To_Staking_PalletStakingExposure(v.ValueAt(1))

    return out //from message
}
    


func To_Balances_TupleUint32Uint32(in any) *pbgear.Balances_TupleUint32Uint32 {
    out := &pbgear.Balances_TupleUint32Uint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_Unbonded(in any) *pbgear.Balances_Unbonded {
    out := &pbgear.Balances_Unbonded{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_UnbondingPoolSlashed(in any) *pbgear.Balances_UnbondingPoolSlashed {
    out := &pbgear.Balances_UnbondingPoolSlashed{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // primitive field Era
    out.Era = To_uint32(v.ValueAt(1))
    // primitive field Balance
    out.Balance = To_string(v.ValueAt(2))

    return out //from message
}
    
func To_Balances_Undelegated(in any) *pbgear.Balances_Undelegated {
    out := &pbgear.Balances_Undelegated{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Unlocked(in any) *pbgear.Balances_Unlocked {
    out := &pbgear.Balances_Unlocked{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Unreserved(in any) *pbgear.Balances_Unreserved {
    out := &pbgear.Balances_Unreserved{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_UpdatedInactive(in any) *pbgear.Balances_UpdatedInactive {
    out := &pbgear.Balances_UpdatedInactive{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Reactivated
    out.Reactivated = To_string(v.ValueAt(0))
    // primitive field Deactivated
    out.Deactivated = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_Balances_UpgradeAccountsCall(in any) *pbgear.Balances_UpgradeAccountsCall {
    out := &pbgear.Balances_UpgradeAccountsCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Who
    out.Who = To_Repeated_Balances_UpgradeAccountsCall_who(v.ValueAt(0))

    return out //from message
}
func To_Balances_UpgradeAccountsCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Balances_UpgradeAccountsCall(in)
    out := &pbgearextrinsic.RuntimeCall_BalancesUpgradeAccountsCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BalancesUpgradeAccountsCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Balances_UpgradeAccountsCall_who(in any) []*pbgear.SpCoreCryptoAccountId32 {
    items := in.([]interface{})

    var out []*pbgear.SpCoreCryptoAccountId32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpCoreCryptoAccountId32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Balances_Upgraded(in any) *pbgear.Balances_Upgraded {
    out := &pbgear.Balances_Upgraded{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_UserMessageRead(in any) *pbgear.Balances_UserMessageRead {
    out := &pbgear.Balances_UserMessageRead{}
    v := in.(registry.Valuable)
    _ = v

    // field Id To_GprimitivesMessageId(w)
    out.Id = To_GprimitivesMessageId(v.ValueAt(0))
    // oneOf field Reason
    v1 := To_OneOf_Balances_UserMessageRead_reason(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Reason").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_Balances_UserMessageRead_reason (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_UserMessageRead_ReasonRuntime{
        ReasonRuntime: To_Runtime(param),
        }
    case 1:
        return &pbgear.Balances_UserMessageRead_ReasonSystem{
        ReasonSystem: To_System(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_UserMessageSent(in any) *pbgear.Balances_UserMessageSent {
    out := &pbgear.Balances_UserMessageSent{}
    v := in.(registry.Valuable)
    _ = v

    // field Message To_GearCoreMessageUserUserMessage(w)
    out.Message = To_GearCoreMessageUserUserMessage(v.ValueAt(0))
    // primitive field Expiration
    out.Expiration = To_Optional_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_UtilityEvent(in any) *pbgear.Balances_UtilityEvent {
    out := &pbgear.Balances_UtilityEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_UtilityEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_UtilityEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_UtilityEvent_Value0BatchInterrupted{
        Value0BatchInterrupted: To_Balances_BatchInterrupted(param),
        }
    case 1:
        return &pbgear.Balances_UtilityEvent_Value0BatchCompleted{
        Value0BatchCompleted: To_Balances_BatchCompleted(param),
        }
    case 2:
        return &pbgear.Balances_UtilityEvent_Value0BatchCompletedWithErrors{
        Value0BatchCompletedWithErrors: To_Balances_BatchCompletedWithErrors(param),
        }
    case 3:
        return &pbgear.Balances_UtilityEvent_Value0ItemCompleted{
        Value0ItemCompleted: To_Balances_ItemCompleted(param),
        }
    case 4:
        return &pbgear.Balances_UtilityEvent_Value0ItemFailed{
        Value0ItemFailed: To_Balances_ItemFailed(param),
        }
    case 5:
        return &pbgear.Balances_UtilityEvent_Value0DispatchedAs{
        Value0DispatchedAs: To_Balances_DispatchedAs(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_ValidatorPrefsSet(in any) *pbgear.Balances_ValidatorPrefsSet {
    out := &pbgear.Balances_ValidatorPrefsSet{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Prefs To_Staking_PalletStakingValidatorPrefs(w)
    out.Prefs = To_Staking_PalletStakingValidatorPrefs(v.ValueAt(1))

    return out //from message
}
    


func To_Balances_Value0(in any) *pbgear.Balances_Value0 {
    out := &pbgear.Balances_Value0{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_Value0_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_Value0_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Value0_Value0ExtrinsicSuccess{
        Value0ExtrinsicSuccess: To_Balances_ExtrinsicSuccess(param),
        }
    case 1:
        return &pbgear.Balances_Value0_Value0ExtrinsicFailed{
        Value0ExtrinsicFailed: To_Balances_ExtrinsicFailed(param),
        }
    case 2:
        return &pbgear.Balances_Value0_Value0CodeUpdated{
        Value0CodeUpdated: To_Balances_CodeUpdated(param),
        }
    case 3:
        return &pbgear.Balances_Value0_Value0NewAccount{
        Value0NewAccount: To_Balances_NewAccount(param),
        }
    case 4:
        return &pbgear.Balances_Value0_Value0KilledAccount{
        Value0KilledAccount: To_Balances_KilledAccount(param),
        }
    case 5:
        return &pbgear.Balances_Value0_Value0Remarked{
        Value0Remarked: To_Balances_Remarked(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_VestingCompleted(in any) *pbgear.Balances_VestingCompleted {
    out := &pbgear.Balances_VestingCompleted{}
    v := in.(registry.Valuable)
    _ = v

    // field Account To_SpCoreCryptoAccountId32(w)
    out.Account = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_VestingEvent(in any) *pbgear.Balances_VestingEvent {
    out := &pbgear.Balances_VestingEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_VestingEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_VestingEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_VestingEvent_Value0VestingUpdated{
        Value0VestingUpdated: To_Balances_VestingUpdated(param),
        }
    case 1:
        return &pbgear.Balances_VestingEvent_Value0VestingCompleted{
        Value0VestingCompleted: To_Balances_VestingCompleted(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_VestingUpdated(in any) *pbgear.Balances_VestingUpdated {
    out := &pbgear.Balances_VestingUpdated{}
    v := in.(registry.Valuable)
    _ = v

    // field Account To_SpCoreCryptoAccountId32(w)
    out.Account = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Unvested
    out.Unvested = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Voted(in any) *pbgear.Balances_Voted {
    out := &pbgear.Balances_Voted{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Poll
    out.Poll = To_uint32(v.ValueAt(1))
    // oneOf field Vote
    v2 := To_OneOf_Balances_Voted_vote(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Vote").Set(reflect.ValueOf(v2))
    // field Tally To_RankedCollective_PalletRankedCollectiveTally(w)
    out.Tally = To_RankedCollective_PalletRankedCollectiveTally(v.ValueAt(3))

    return out //from message
}
    

func To_OneOf_Balances_Voted_vote (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Voted_VoteAye{
        VoteAye: To_RankedCollective_Aye(param),
        }
    case 1:
        return &pbgear.Balances_Voted_VoteNay{
        VoteNay: To_RankedCollective_Nay(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Balances_VoucherDeclined(in any) *pbgear.Balances_VoucherDeclined {
    out := &pbgear.Balances_VoucherDeclined{}
    v := in.(registry.Valuable)
    _ = v

    // field Spender To_SpCoreCryptoAccountId32(w)
    out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(1))

    return out //from message
}
    


func To_Balances_VoucherIssued(in any) *pbgear.Balances_VoucherIssued {
    out := &pbgear.Balances_VoucherIssued{}
    v := in.(registry.Valuable)
    _ = v

    // field Owner To_SpCoreCryptoAccountId32(w)
    out.Owner = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Spender To_SpCoreCryptoAccountId32(w)
    out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(1))
    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(2))

    return out //from message
}
    



func To_Balances_VoucherRevoked(in any) *pbgear.Balances_VoucherRevoked {
    out := &pbgear.Balances_VoucherRevoked{}
    v := in.(registry.Valuable)
    _ = v

    // field Spender To_SpCoreCryptoAccountId32(w)
    out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(1))

    return out //from message
}
    


func To_Balances_VoucherUpdated(in any) *pbgear.Balances_VoucherUpdated {
    out := &pbgear.Balances_VoucherUpdated{}
    v := in.(registry.Valuable)
    _ = v

    // field Spender To_SpCoreCryptoAccountId32(w)
    out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(1))
    // optional field NewOwner
    if v.HasValue() {
        out.NewOwner = To_Optional_Balances_VoucherUpdated_new_owner(v.ValueAt(2))
    }

    return out //from message
}
    


func To_Optional_Balances_VoucherUpdated_new_owner(in any) *pbgear.SpCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_Balances_WhitelistEvent(in any) *pbgear.Balances_WhitelistEvent {
    out := &pbgear.Balances_WhitelistEvent{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Balances_WhitelistEvent_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_WhitelistEvent_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_WhitelistEvent_Value0CallWhitelisted{
        Value0CallWhitelisted: To_Balances_CallWhitelisted(param),
        }
    case 1:
        return &pbgear.Balances_WhitelistEvent_Value0WhitelistedCallRemoved{
        Value0WhitelistedCallRemoved: To_Balances_WhitelistedCallRemoved(param),
        }
    case 2:
        return &pbgear.Balances_WhitelistEvent_Value0WhitelistedCallDispatched{
        Value0WhitelistedCallDispatched: To_Balances_WhitelistedCallDispatched(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_WhitelistedCallDispatched(in any) *pbgear.Balances_WhitelistedCallDispatched {
    out := &pbgear.Balances_WhitelistedCallDispatched{}
    v := in.(registry.Valuable)
    _ = v

    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))
    // oneOf field Result
    v1 := To_OneOf_Balances_WhitelistedCallDispatched_result(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Result").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_Balances_WhitelistedCallDispatched_result (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_WhitelistedCallDispatched_ResultOk{
        ResultOk: To_Ok(param),
        }
    case 1:
        return &pbgear.Balances_WhitelistedCallDispatched_ResultErr{
        ResultErr: To_Err(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_WhitelistedCallRemoved(in any) *pbgear.Balances_WhitelistedCallRemoved {
    out := &pbgear.Balances_WhitelistedCallRemoved{}
    v := in.(registry.Valuable)
    _ = v

    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_Balances_Who(in any) *pbgear.Balances_Who {
    out := &pbgear.Balances_Who{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_Balances_Who_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Balances_Who_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Balances_Who_WhoId{
        WhoId: To_Balances_Id(param),
        }
    case 1:
        return &pbgear.Balances_Who_WhoIndex{
        WhoIndex: To_Balances_Index(param),
        }
    case 2:
        return &pbgear.Balances_Who_WhoRaw{
        WhoRaw: To_Balances_Raw(param),
        }
    case 3:
        return &pbgear.Balances_Who_WhoAddress32{
        WhoAddress32: To_Balances_Address32(param),
        }
    case 4:
        return &pbgear.Balances_Who_WhoAddress20{
        WhoAddress20: To_Balances_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Balances_Withdraw(in any) *pbgear.Balances_Withdraw {
    out := &pbgear.Balances_Withdraw{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Balances_Withdrawn(in any) *pbgear.Balances_Withdrawn {
    out := &pbgear.Balances_Withdrawn{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_BelowMinimum(in any) *pbgear.BelowMinimum {
    out := &pbgear.BelowMinimum{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BigSpender(in any) *pbgear.BigSpender {
    out := &pbgear.BigSpender{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BigTipper(in any) *pbgear.BigTipper {
    out := &pbgear.BigTipper{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Blocked(in any) *pbgear.Blocked {
    out := &pbgear.Blocked{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BoundedCollectionsBoundedVecBoundedVec(in any) *pbgear.BoundedCollectionsBoundedVecBoundedVec {
    out := &pbgear.BoundedCollectionsBoundedVecBoundedVec{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_BountiesBountyApprovedEvent(in any) *pbgear.BountiesBountyApprovedEvent {
    out := &pbgear.BountiesBountyApprovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesBountyAwardedEvent(in any) *pbgear.BountiesBountyAwardedEvent {
    out := &pbgear.BountiesBountyAwardedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesBountyBecameActiveEvent(in any) *pbgear.BountiesBountyBecameActiveEvent {
    out := &pbgear.BountiesBountyBecameActiveEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesBountyCanceledEvent(in any) *pbgear.BountiesBountyCanceledEvent {
    out := &pbgear.BountiesBountyCanceledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesBountyClaimedEvent(in any) *pbgear.BountiesBountyClaimedEvent {
    out := &pbgear.BountiesBountyClaimedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesBountyExtendedEvent(in any) *pbgear.BountiesBountyExtendedEvent {
    out := &pbgear.BountiesBountyExtendedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesBountyProposedEvent(in any) *pbgear.BountiesBountyProposedEvent {
    out := &pbgear.BountiesBountyProposedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesBountyRejectedEvent(in any) *pbgear.BountiesBountyRejectedEvent {
    out := &pbgear.BountiesBountyRejectedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesCuratorAcceptedEvent(in any) *pbgear.BountiesCuratorAcceptedEvent {
    out := &pbgear.BountiesCuratorAcceptedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesCuratorProposedEvent(in any) *pbgear.BountiesCuratorProposedEvent {
    out := &pbgear.BountiesCuratorProposedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesCuratorUnassignedEvent(in any) *pbgear.BountiesCuratorUnassignedEvent {
    out := &pbgear.BountiesCuratorUnassignedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_BountiesPallet(in any) *pbgear.BountiesPallet {
    out := &pbgear.BountiesPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_BountiesPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_BountiesPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.BountiesPallet_CallProposeBountyCall{
        CallProposeBountyCall: To_Bounties_ProposeBountyCall(param),
        }
    case 1:
        return &pbgear.BountiesPallet_CallApproveBountyCall{
        CallApproveBountyCall: To_Bounties_ApproveBountyCall(param),
        }
    case 2:
        return &pbgear.BountiesPallet_CallProposeCuratorCall{
        CallProposeCuratorCall: To_Bounties_ProposeCuratorCall(param),
        }
    case 3:
        return &pbgear.BountiesPallet_CallUnassignCuratorCall{
        CallUnassignCuratorCall: To_Bounties_UnassignCuratorCall(param),
        }
    case 4:
        return &pbgear.BountiesPallet_CallAcceptCuratorCall{
        CallAcceptCuratorCall: To_Bounties_AcceptCuratorCall(param),
        }
    case 5:
        return &pbgear.BountiesPallet_CallAwardBountyCall{
        CallAwardBountyCall: To_Bounties_AwardBountyCall(param),
        }
    case 6:
        return &pbgear.BountiesPallet_CallClaimBountyCall{
        CallClaimBountyCall: To_Bounties_ClaimBountyCall(param),
        }
    case 7:
        return &pbgear.BountiesPallet_CallCloseBountyCall{
        CallCloseBountyCall: To_Bounties_CloseBountyCall(param),
        }
    case 8:
        return &pbgear.BountiesPallet_CallExtendBountyExpiryCall{
        CallExtendBountyExpiryCall: To_Bounties_ExtendBountyExpiryCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Bounties_AcceptCuratorCall(in any) *pbgear.Bounties_AcceptCuratorCall {
    out := &pbgear.Bounties_AcceptCuratorCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Bounties_AcceptCuratorCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_AcceptCuratorCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesAcceptCuratorCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesAcceptCuratorCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Bounties_Address20(in any) *pbgear.Bounties_Address20 {
    out := &pbgear.Bounties_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Bounties_Address32(in any) *pbgear.Bounties_Address32 {
    out := &pbgear.Bounties_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Bounties_ApproveBountyCall(in any) *pbgear.Bounties_ApproveBountyCall {
    out := &pbgear.Bounties_ApproveBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Bounties_ApproveBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_ApproveBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesApproveBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesApproveBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Bounties_AwardBountyCall(in any) *pbgear.Bounties_AwardBountyCall {
    out := &pbgear.Bounties_AwardBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))
    // oneOf field Beneficiary
    v1 := To_OneOf_Bounties_AwardBountyCall_beneficiary(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Bounties_AwardBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_AwardBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesAwardBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesAwardBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Bounties_AwardBountyCall_beneficiary (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Bounties_AwardBountyCall_BeneficiaryId{
        BeneficiaryId: To_Bounties_Id(param),
        }
    case 1:
        return &pbgear.Bounties_AwardBountyCall_BeneficiaryIndex{
        BeneficiaryIndex: To_Bounties_Index(param),
        }
    case 2:
        return &pbgear.Bounties_AwardBountyCall_BeneficiaryRaw{
        BeneficiaryRaw: To_Bounties_Raw(param),
        }
    case 3:
        return &pbgear.Bounties_AwardBountyCall_BeneficiaryAddress32{
        BeneficiaryAddress32: To_Bounties_Address32(param),
        }
    case 4:
        return &pbgear.Bounties_AwardBountyCall_BeneficiaryAddress20{
        BeneficiaryAddress20: To_Bounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Bounties_Beneficiary(in any) *pbgear.Bounties_Beneficiary {
    out := &pbgear.Bounties_Beneficiary{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Beneficiary
    v0 := To_OneOf_Bounties_Beneficiary_beneficiary(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Bounties_Beneficiary_beneficiary (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Bounties_Beneficiary_BeneficiaryId{
        BeneficiaryId: To_Bounties_Id(param),
        }
    case 1:
        return &pbgear.Bounties_Beneficiary_BeneficiaryIndex{
        BeneficiaryIndex: To_Bounties_Index(param),
        }
    case 2:
        return &pbgear.Bounties_Beneficiary_BeneficiaryRaw{
        BeneficiaryRaw: To_Bounties_Raw(param),
        }
    case 3:
        return &pbgear.Bounties_Beneficiary_BeneficiaryAddress32{
        BeneficiaryAddress32: To_Bounties_Address32(param),
        }
    case 4:
        return &pbgear.Bounties_Beneficiary_BeneficiaryAddress20{
        BeneficiaryAddress20: To_Bounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Bounties_ClaimBountyCall(in any) *pbgear.Bounties_ClaimBountyCall {
    out := &pbgear.Bounties_ClaimBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Bounties_ClaimBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_ClaimBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesClaimBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesClaimBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Bounties_CloseBountyCall(in any) *pbgear.Bounties_CloseBountyCall {
    out := &pbgear.Bounties_CloseBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Bounties_CloseBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_CloseBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesCloseBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesCloseBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Bounties_Curator(in any) *pbgear.Bounties_Curator {
    out := &pbgear.Bounties_Curator{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Curator
    v0 := To_OneOf_Bounties_Curator_curator(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Bounties_Curator_curator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Bounties_Curator_CuratorId{
        CuratorId: To_Bounties_Id(param),
        }
    case 1:
        return &pbgear.Bounties_Curator_CuratorIndex{
        CuratorIndex: To_Bounties_Index(param),
        }
    case 2:
        return &pbgear.Bounties_Curator_CuratorRaw{
        CuratorRaw: To_Bounties_Raw(param),
        }
    case 3:
        return &pbgear.Bounties_Curator_CuratorAddress32{
        CuratorAddress32: To_Bounties_Address32(param),
        }
    case 4:
        return &pbgear.Bounties_Curator_CuratorAddress20{
        CuratorAddress20: To_Bounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Bounties_ExtendBountyExpiryCall(in any) *pbgear.Bounties_ExtendBountyExpiryCall {
    out := &pbgear.Bounties_ExtendBountyExpiryCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))
    // primitive field Remark
    out.Remark = To_bytes(v.ValueAt(1))

    return out //from message
}
func To_Bounties_ExtendBountyExpiryCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_ExtendBountyExpiryCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesExtendBountyExpiryCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesExtendBountyExpiryCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Bounties_Id(in any) *pbgear.Bounties_Id {
    out := &pbgear.Bounties_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Bounties_Index(in any) *pbgear.Bounties_Index {
    out := &pbgear.Bounties_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Bounties_TupleNull(w)
    out.Value0 = To_Bounties_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Bounties_ProposeBountyCall(in any) *pbgear.Bounties_ProposeBountyCall {
    out := &pbgear.Bounties_ProposeBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value
    out.Value = To_string(v.ValueAt(0))
    // primitive field Description
    out.Description = To_bytes(v.ValueAt(1))

    return out //from message
}
func To_Bounties_ProposeBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_ProposeBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesProposeBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesProposeBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Bounties_ProposeCuratorCall(in any) *pbgear.Bounties_ProposeCuratorCall {
    out := &pbgear.Bounties_ProposeCuratorCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))
    // oneOf field Curator
    v1 := To_OneOf_Bounties_ProposeCuratorCall_curator(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v1))
    // primitive field Fee
    out.Fee = To_string(v.ValueAt(2))

    return out //from message
}
func To_Bounties_ProposeCuratorCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_ProposeCuratorCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesProposeCuratorCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesProposeCuratorCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Bounties_ProposeCuratorCall_curator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Bounties_ProposeCuratorCall_CuratorId{
        CuratorId: To_Bounties_Id(param),
        }
    case 1:
        return &pbgear.Bounties_ProposeCuratorCall_CuratorIndex{
        CuratorIndex: To_Bounties_Index(param),
        }
    case 2:
        return &pbgear.Bounties_ProposeCuratorCall_CuratorRaw{
        CuratorRaw: To_Bounties_Raw(param),
        }
    case 3:
        return &pbgear.Bounties_ProposeCuratorCall_CuratorAddress32{
        CuratorAddress32: To_Bounties_Address32(param),
        }
    case 4:
        return &pbgear.Bounties_ProposeCuratorCall_CuratorAddress20{
        CuratorAddress20: To_Bounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Bounties_Raw(in any) *pbgear.Bounties_Raw {
    out := &pbgear.Bounties_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Bounties_TupleNull(in any) *pbgear.Bounties_TupleNull {
    out := &pbgear.Bounties_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Bounties_UnassignCuratorCall(in any) *pbgear.Bounties_UnassignCuratorCall {
    out := &pbgear.Bounties_UnassignCuratorCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BountyId
    out.BountyId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Bounties_UnassignCuratorCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Bounties_UnassignCuratorCall(in)
    out := &pbgearextrinsic.RuntimeCall_BountiesUnassignCuratorCall{ }
    reflect.ValueOf(out).Elem().FieldByName("BountiesUnassignCuratorCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_CancelProxy(in any) *pbgear.CancelProxy {
    out := &pbgear.CancelProxy{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_CannotCreate(in any) *pbgear.CannotCreate {
    out := &pbgear.CannotCreate{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_CannotCreateHold(in any) *pbgear.CannotCreateHold {
    out := &pbgear.CannotCreateHold{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_CannotLookup(in any) *pbgear.CannotLookup {
    out := &pbgear.CannotLookup{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Change(in any) *pbgear.Change {
    out := &pbgear.Change{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Change
    v0 := To_OneOf_Change_change(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Change").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Change_change (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Change_ChangeActive{
        ChangeActive: To_Active(param),
        }
    case 1:
        return &pbgear.Change_ChangeInactive{
        ChangeInactive: To_Inactive(param),
        }
    case 2:
        return &pbgear.Change_ChangeReinstrumented{
        ChangeReinstrumented: To_Reinstrumented(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ChildBountiesAddedEvent(in any) *pbgear.ChildBountiesAddedEvent {
    out := &pbgear.ChildBountiesAddedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ChildBountiesAwardedEvent(in any) *pbgear.ChildBountiesAwardedEvent {
    out := &pbgear.ChildBountiesAwardedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ChildBountiesCanceledEvent(in any) *pbgear.ChildBountiesCanceledEvent {
    out := &pbgear.ChildBountiesCanceledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ChildBountiesClaimedEvent(in any) *pbgear.ChildBountiesClaimedEvent {
    out := &pbgear.ChildBountiesClaimedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ChildBountiesPallet(in any) *pbgear.ChildBountiesPallet {
    out := &pbgear.ChildBountiesPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_ChildBountiesPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ChildBountiesPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ChildBountiesPallet_CallAddChildBountyCall{
        CallAddChildBountyCall: To_ChildBounties_AddChildBountyCall(param),
        }
    case 1:
        return &pbgear.ChildBountiesPallet_CallProposeCuratorCall{
        CallProposeCuratorCall: To_ChildBounties_ProposeCuratorCall(param),
        }
    case 2:
        return &pbgear.ChildBountiesPallet_CallAcceptCuratorCall{
        CallAcceptCuratorCall: To_ChildBounties_AcceptCuratorCall(param),
        }
    case 3:
        return &pbgear.ChildBountiesPallet_CallUnassignCuratorCall{
        CallUnassignCuratorCall: To_ChildBounties_UnassignCuratorCall(param),
        }
    case 4:
        return &pbgear.ChildBountiesPallet_CallAwardChildBountyCall{
        CallAwardChildBountyCall: To_ChildBounties_AwardChildBountyCall(param),
        }
    case 5:
        return &pbgear.ChildBountiesPallet_CallClaimChildBountyCall{
        CallClaimChildBountyCall: To_ChildBounties_ClaimChildBountyCall(param),
        }
    case 6:
        return &pbgear.ChildBountiesPallet_CallCloseChildBountyCall{
        CallCloseChildBountyCall: To_ChildBounties_CloseChildBountyCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ChildBounties_AcceptCuratorCall(in any) *pbgear.ChildBounties_AcceptCuratorCall {
    out := &pbgear.ChildBounties_AcceptCuratorCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ParentBountyId
    out.ParentBountyId = To_uint32(v.ValueAt(0))
    // primitive field ChildBountyId
    out.ChildBountyId = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_ChildBounties_AcceptCuratorCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ChildBounties_AcceptCuratorCall(in)
    out := &pbgearextrinsic.RuntimeCall_ChildbountiesAcceptCuratorCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ChildbountiesAcceptCuratorCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_ChildBounties_AddChildBountyCall(in any) *pbgear.ChildBounties_AddChildBountyCall {
    out := &pbgear.ChildBounties_AddChildBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ParentBountyId
    out.ParentBountyId = To_uint32(v.ValueAt(0))
    // primitive field Value
    out.Value = To_string(v.ValueAt(1))
    // primitive field Description
    out.Description = To_bytes(v.ValueAt(2))

    return out //from message
}
func To_ChildBounties_AddChildBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ChildBounties_AddChildBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_ChildbountiesAddChildBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ChildbountiesAddChildBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_ChildBounties_Address20(in any) *pbgear.ChildBounties_Address20 {
    out := &pbgear.ChildBounties_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_ChildBounties_Address32(in any) *pbgear.ChildBounties_Address32 {
    out := &pbgear.ChildBounties_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_ChildBounties_AwardChildBountyCall(in any) *pbgear.ChildBounties_AwardChildBountyCall {
    out := &pbgear.ChildBounties_AwardChildBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ParentBountyId
    out.ParentBountyId = To_uint32(v.ValueAt(0))
    // primitive field ChildBountyId
    out.ChildBountyId = To_uint32(v.ValueAt(1))
    // oneOf field Beneficiary
    v2 := To_OneOf_ChildBounties_AwardChildBountyCall_beneficiary(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v2))

    return out //from message
}
func To_ChildBounties_AwardChildBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ChildBounties_AwardChildBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_ChildbountiesAwardChildBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ChildbountiesAwardChildBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_ChildBounties_AwardChildBountyCall_beneficiary (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryId{
        BeneficiaryId: To_ChildBounties_Id(param),
        }
    case 1:
        return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryIndex{
        BeneficiaryIndex: To_ChildBounties_Index(param),
        }
    case 2:
        return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryRaw{
        BeneficiaryRaw: To_ChildBounties_Raw(param),
        }
    case 3:
        return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryAddress32{
        BeneficiaryAddress32: To_ChildBounties_Address32(param),
        }
    case 4:
        return &pbgear.ChildBounties_AwardChildBountyCall_BeneficiaryAddress20{
        BeneficiaryAddress20: To_ChildBounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ChildBounties_Beneficiary(in any) *pbgear.ChildBounties_Beneficiary {
    out := &pbgear.ChildBounties_Beneficiary{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Beneficiary
    v0 := To_OneOf_ChildBounties_Beneficiary_beneficiary(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ChildBounties_Beneficiary_beneficiary (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ChildBounties_Beneficiary_BeneficiaryId{
        BeneficiaryId: To_ChildBounties_Id(param),
        }
    case 1:
        return &pbgear.ChildBounties_Beneficiary_BeneficiaryIndex{
        BeneficiaryIndex: To_ChildBounties_Index(param),
        }
    case 2:
        return &pbgear.ChildBounties_Beneficiary_BeneficiaryRaw{
        BeneficiaryRaw: To_ChildBounties_Raw(param),
        }
    case 3:
        return &pbgear.ChildBounties_Beneficiary_BeneficiaryAddress32{
        BeneficiaryAddress32: To_ChildBounties_Address32(param),
        }
    case 4:
        return &pbgear.ChildBounties_Beneficiary_BeneficiaryAddress20{
        BeneficiaryAddress20: To_ChildBounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ChildBounties_ClaimChildBountyCall(in any) *pbgear.ChildBounties_ClaimChildBountyCall {
    out := &pbgear.ChildBounties_ClaimChildBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ParentBountyId
    out.ParentBountyId = To_uint32(v.ValueAt(0))
    // primitive field ChildBountyId
    out.ChildBountyId = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_ChildBounties_ClaimChildBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ChildBounties_ClaimChildBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_ChildbountiesClaimChildBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ChildbountiesClaimChildBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_ChildBounties_CloseChildBountyCall(in any) *pbgear.ChildBounties_CloseChildBountyCall {
    out := &pbgear.ChildBounties_CloseChildBountyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ParentBountyId
    out.ParentBountyId = To_uint32(v.ValueAt(0))
    // primitive field ChildBountyId
    out.ChildBountyId = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_ChildBounties_CloseChildBountyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ChildBounties_CloseChildBountyCall(in)
    out := &pbgearextrinsic.RuntimeCall_ChildbountiesCloseChildBountyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ChildbountiesCloseChildBountyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_ChildBounties_Curator(in any) *pbgear.ChildBounties_Curator {
    out := &pbgear.ChildBounties_Curator{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Curator
    v0 := To_OneOf_ChildBounties_Curator_curator(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ChildBounties_Curator_curator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ChildBounties_Curator_CuratorId{
        CuratorId: To_ChildBounties_Id(param),
        }
    case 1:
        return &pbgear.ChildBounties_Curator_CuratorIndex{
        CuratorIndex: To_ChildBounties_Index(param),
        }
    case 2:
        return &pbgear.ChildBounties_Curator_CuratorRaw{
        CuratorRaw: To_ChildBounties_Raw(param),
        }
    case 3:
        return &pbgear.ChildBounties_Curator_CuratorAddress32{
        CuratorAddress32: To_ChildBounties_Address32(param),
        }
    case 4:
        return &pbgear.ChildBounties_Curator_CuratorAddress20{
        CuratorAddress20: To_ChildBounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ChildBounties_Id(in any) *pbgear.ChildBounties_Id {
    out := &pbgear.ChildBounties_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_ChildBounties_Index(in any) *pbgear.ChildBounties_Index {
    out := &pbgear.ChildBounties_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_ChildBounties_TupleNull(w)
    out.Value0 = To_ChildBounties_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_ChildBounties_ProposeCuratorCall(in any) *pbgear.ChildBounties_ProposeCuratorCall {
    out := &pbgear.ChildBounties_ProposeCuratorCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ParentBountyId
    out.ParentBountyId = To_uint32(v.ValueAt(0))
    // primitive field ChildBountyId
    out.ChildBountyId = To_uint32(v.ValueAt(1))
    // oneOf field Curator
    v2 := To_OneOf_ChildBounties_ProposeCuratorCall_curator(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Curator").Set(reflect.ValueOf(v2))
    // primitive field Fee
    out.Fee = To_string(v.ValueAt(3))

    return out //from message
}
func To_ChildBounties_ProposeCuratorCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ChildBounties_ProposeCuratorCall(in)
    out := &pbgearextrinsic.RuntimeCall_ChildbountiesProposeCuratorCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ChildbountiesProposeCuratorCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_ChildBounties_ProposeCuratorCall_curator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ChildBounties_ProposeCuratorCall_CuratorId{
        CuratorId: To_ChildBounties_Id(param),
        }
    case 1:
        return &pbgear.ChildBounties_ProposeCuratorCall_CuratorIndex{
        CuratorIndex: To_ChildBounties_Index(param),
        }
    case 2:
        return &pbgear.ChildBounties_ProposeCuratorCall_CuratorRaw{
        CuratorRaw: To_ChildBounties_Raw(param),
        }
    case 3:
        return &pbgear.ChildBounties_ProposeCuratorCall_CuratorAddress32{
        CuratorAddress32: To_ChildBounties_Address32(param),
        }
    case 4:
        return &pbgear.ChildBounties_ProposeCuratorCall_CuratorAddress20{
        CuratorAddress20: To_ChildBounties_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ChildBounties_Raw(in any) *pbgear.ChildBounties_Raw {
    out := &pbgear.ChildBounties_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_ChildBounties_TupleNull(in any) *pbgear.ChildBounties_TupleNull {
    out := &pbgear.ChildBounties_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_ChildBounties_UnassignCuratorCall(in any) *pbgear.ChildBounties_UnassignCuratorCall {
    out := &pbgear.ChildBounties_UnassignCuratorCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ParentBountyId
    out.ParentBountyId = To_uint32(v.ValueAt(0))
    // primitive field ChildBountyId
    out.ChildBountyId = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_ChildBounties_UnassignCuratorCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ChildBounties_UnassignCuratorCall(in)
    out := &pbgearextrinsic.RuntimeCall_ChildbountiesUnassignCuratorCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ChildbountiesUnassignCuratorCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Class(in any) *pbgear.Class {
    out := &pbgear.Class{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Class
    v0 := To_OneOf_Class_class(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Class").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Class_class (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Class_ClassNormal{
        ClassNormal: To_Normal(param),
        }
    case 1:
        return &pbgear.Class_ClassOperational{
        ClassOperational: To_Operational(param),
        }
    case 2:
        return &pbgear.Class_ClassMandatory{
        ClassMandatory: To_Mandatory(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Code(in any) *pbgear.Code {
    out := &pbgear.Code{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Code
    v0 := To_OneOf_Code_code(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Code").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Code_code (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Code_CodeSuccess{
        CodeSuccess: To_Success(param),
        }
    case 1:
        return &pbgear.Code_CodeError{
        CodeError: To_Error(param),
        }
    case 255:
        return &pbgear.Code_CodeUnsupported{
        CodeUnsupported: To_Unsupported(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_CodeNotExists(in any) *pbgear.CodeNotExists {
    out := &pbgear.CodeNotExists{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConsumerRemaining(in any) *pbgear.ConsumerRemaining {
    out := &pbgear.ConsumerRemaining{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVotingDelegatedEvent(in any) *pbgear.ConvictionVotingDelegatedEvent {
    out := &pbgear.ConvictionVotingDelegatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVotingPallet(in any) *pbgear.ConvictionVotingPallet {
    out := &pbgear.ConvictionVotingPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_ConvictionVotingPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ConvictionVotingPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVotingPallet_CallVoteCall{
        CallVoteCall: To_ConvictionVoting_VoteCall(param),
        }
    case 1:
        return &pbgear.ConvictionVotingPallet_CallDelegateCall{
        CallDelegateCall: To_ConvictionVoting_DelegateCall(param),
        }
    case 2:
        return &pbgear.ConvictionVotingPallet_CallUndelegateCall{
        CallUndelegateCall: To_ConvictionVoting_UndelegateCall(param),
        }
    case 3:
        return &pbgear.ConvictionVotingPallet_CallUnlockCall{
        CallUnlockCall: To_ConvictionVoting_UnlockCall(param),
        }
    case 4:
        return &pbgear.ConvictionVotingPallet_CallRemoveVoteCall{
        CallRemoveVoteCall: To_ConvictionVoting_RemoveVoteCall(param),
        }
    case 5:
        return &pbgear.ConvictionVotingPallet_CallRemoveOtherVoteCall{
        CallRemoveOtherVoteCall: To_ConvictionVoting_RemoveOtherVoteCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVotingUndelegatedEvent(in any) *pbgear.ConvictionVotingUndelegatedEvent {
    out := &pbgear.ConvictionVotingUndelegatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_Address20(in any) *pbgear.ConvictionVoting_Address20 {
    out := &pbgear.ConvictionVoting_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_ConvictionVoting_Address32(in any) *pbgear.ConvictionVoting_Address32 {
    out := &pbgear.ConvictionVoting_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_ConvictionVoting_Conviction(in any) *pbgear.ConvictionVoting_Conviction {
    out := &pbgear.ConvictionVoting_Conviction{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Conviction
    v0 := To_OneOf_ConvictionVoting_Conviction_conviction(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Conviction").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ConvictionVoting_Conviction_conviction (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_Conviction_ConvictionNone{
        ConvictionNone: To_ConvictionVoting_None(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_Conviction_ConvictionLocked1X{
        ConvictionLocked1X: To_ConvictionVoting_Locked1X(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_Conviction_ConvictionLocked2X{
        ConvictionLocked2X: To_ConvictionVoting_Locked2X(param),
        }
    case 3:
        return &pbgear.ConvictionVoting_Conviction_ConvictionLocked3X{
        ConvictionLocked3X: To_ConvictionVoting_Locked3X(param),
        }
    case 4:
        return &pbgear.ConvictionVoting_Conviction_ConvictionLocked4X{
        ConvictionLocked4X: To_ConvictionVoting_Locked4X(param),
        }
    case 5:
        return &pbgear.ConvictionVoting_Conviction_ConvictionLocked5X{
        ConvictionLocked5X: To_ConvictionVoting_Locked5X(param),
        }
    case 6:
        return &pbgear.ConvictionVoting_Conviction_ConvictionLocked6X{
        ConvictionLocked6X: To_ConvictionVoting_Locked6X(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVoting_DelegateCall(in any) *pbgear.ConvictionVoting_DelegateCall {
    out := &pbgear.ConvictionVoting_DelegateCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Class
    out.Class = To_uint32(v.ValueAt(0))
    // oneOf field To
    v1 := To_OneOf_ConvictionVoting_DelegateCall_to(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v1))
    // oneOf field Conviction
    v2 := To_OneOf_ConvictionVoting_DelegateCall_conviction(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Conviction").Set(reflect.ValueOf(v2))
    // primitive field Balance
    out.Balance = To_string(v.ValueAt(3))

    return out //from message
}
func To_ConvictionVoting_DelegateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ConvictionVoting_DelegateCall(in)
    out := &pbgearextrinsic.RuntimeCall_ConvictionvotingDelegateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ConvictionvotingDelegateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_ConvictionVoting_DelegateCall_to (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_DelegateCall_ToId{
        ToId: To_ConvictionVoting_Id(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_DelegateCall_ToIndex{
        ToIndex: To_ConvictionVoting_Index(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_DelegateCall_ToRaw{
        ToRaw: To_ConvictionVoting_Raw(param),
        }
    case 3:
        return &pbgear.ConvictionVoting_DelegateCall_ToAddress32{
        ToAddress32: To_ConvictionVoting_Address32(param),
        }
    case 4:
        return &pbgear.ConvictionVoting_DelegateCall_ToAddress20{
        ToAddress20: To_ConvictionVoting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_ConvictionVoting_DelegateCall_conviction (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_DelegateCall_ConvictionNone{
        ConvictionNone: To_ConvictionVoting_None(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked1X{
        ConvictionLocked1X: To_ConvictionVoting_Locked1X(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked2X{
        ConvictionLocked2X: To_ConvictionVoting_Locked2X(param),
        }
    case 3:
        return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked3X{
        ConvictionLocked3X: To_ConvictionVoting_Locked3X(param),
        }
    case 4:
        return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked4X{
        ConvictionLocked4X: To_ConvictionVoting_Locked4X(param),
        }
    case 5:
        return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked5X{
        ConvictionLocked5X: To_ConvictionVoting_Locked5X(param),
        }
    case 6:
        return &pbgear.ConvictionVoting_DelegateCall_ConvictionLocked6X{
        ConvictionLocked6X: To_ConvictionVoting_Locked6X(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVoting_Id(in any) *pbgear.ConvictionVoting_Id {
    out := &pbgear.ConvictionVoting_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_ConvictionVoting_Index(in any) *pbgear.ConvictionVoting_Index {
    out := &pbgear.ConvictionVoting_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_ConvictionVoting_TupleNull(w)
    out.Value0 = To_ConvictionVoting_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_ConvictionVoting_Locked1X(in any) *pbgear.ConvictionVoting_Locked1X {
    out := &pbgear.ConvictionVoting_Locked1X{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_Locked2X(in any) *pbgear.ConvictionVoting_Locked2X {
    out := &pbgear.ConvictionVoting_Locked2X{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_Locked3X(in any) *pbgear.ConvictionVoting_Locked3X {
    out := &pbgear.ConvictionVoting_Locked3X{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_Locked4X(in any) *pbgear.ConvictionVoting_Locked4X {
    out := &pbgear.ConvictionVoting_Locked4X{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_Locked5X(in any) *pbgear.ConvictionVoting_Locked5X {
    out := &pbgear.ConvictionVoting_Locked5X{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_Locked6X(in any) *pbgear.ConvictionVoting_Locked6X {
    out := &pbgear.ConvictionVoting_Locked6X{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_None(in any) *pbgear.ConvictionVoting_None {
    out := &pbgear.ConvictionVoting_None{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ConvictionVoting_PalletConvictionVotingTypesTally(in any) *pbgear.ConvictionVoting_PalletConvictionVotingTypesTally {
    out := &pbgear.ConvictionVoting_PalletConvictionVotingTypesTally{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Ayes
    out.Ayes = To_string(v.ValueAt(0))
    // primitive field Nays
    out.Nays = To_string(v.ValueAt(1))
    // primitive field Support
    out.Support = To_string(v.ValueAt(2))

    return out //from message
}
    
func To_ConvictionVoting_PalletConvictionVotingVoteVote(in any) *pbgear.ConvictionVoting_PalletConvictionVotingVoteVote {
    out := &pbgear.ConvictionVoting_PalletConvictionVotingVoteVote{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_ConvictionVoting_Raw(in any) *pbgear.ConvictionVoting_Raw {
    out := &pbgear.ConvictionVoting_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_ConvictionVoting_RemoveOtherVoteCall(in any) *pbgear.ConvictionVoting_RemoveOtherVoteCall {
    out := &pbgear.ConvictionVoting_RemoveOtherVoteCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Target
    v0 := To_OneOf_ConvictionVoting_RemoveOtherVoteCall_target(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))
    // primitive field Class
    out.Class = To_uint32(v.ValueAt(1))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(2))

    return out //from message
}
func To_ConvictionVoting_RemoveOtherVoteCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ConvictionVoting_RemoveOtherVoteCall(in)
    out := &pbgearextrinsic.RuntimeCall_ConvictionvotingRemoveOtherVoteCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ConvictionvotingRemoveOtherVoteCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_ConvictionVoting_RemoveOtherVoteCall_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetId{
        TargetId: To_ConvictionVoting_Id(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetIndex{
        TargetIndex: To_ConvictionVoting_Index(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetRaw{
        TargetRaw: To_ConvictionVoting_Raw(param),
        }
    case 3:
        return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetAddress32{
        TargetAddress32: To_ConvictionVoting_Address32(param),
        }
    case 4:
        return &pbgear.ConvictionVoting_RemoveOtherVoteCall_TargetAddress20{
        TargetAddress20: To_ConvictionVoting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVoting_RemoveVoteCall(in any) *pbgear.ConvictionVoting_RemoveVoteCall {
    out := &pbgear.ConvictionVoting_RemoveVoteCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Class
    out.Class = To_Optional_uint32(v.ValueAt(0))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_ConvictionVoting_RemoveVoteCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ConvictionVoting_RemoveVoteCall(in)
    out := &pbgearextrinsic.RuntimeCall_ConvictionvotingRemoveVoteCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ConvictionvotingRemoveVoteCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_ConvictionVoting_Split(in any) *pbgear.ConvictionVoting_Split {
    out := &pbgear.ConvictionVoting_Split{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Aye
    out.Aye = To_string(v.ValueAt(0))
    // primitive field Nay
    out.Nay = To_string(v.ValueAt(1))

    return out //from message
}
    
func To_ConvictionVoting_SplitAbstain(in any) *pbgear.ConvictionVoting_SplitAbstain {
    out := &pbgear.ConvictionVoting_SplitAbstain{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Aye
    out.Aye = To_string(v.ValueAt(0))
    // primitive field Nay
    out.Nay = To_string(v.ValueAt(1))
    // primitive field Abstain
    out.Abstain = To_string(v.ValueAt(2))

    return out //from message
}
    
func To_ConvictionVoting_Standard(in any) *pbgear.ConvictionVoting_Standard {
    out := &pbgear.ConvictionVoting_Standard{}
    v := in.(registry.Valuable)
    _ = v

    // field Vote To_ConvictionVoting_PalletConvictionVotingVoteVote(w)
    out.Vote = To_ConvictionVoting_PalletConvictionVotingVoteVote(v.ValueAt(0))
    // primitive field Balance
    out.Balance = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_ConvictionVoting_Target(in any) *pbgear.ConvictionVoting_Target {
    out := &pbgear.ConvictionVoting_Target{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Target
    v0 := To_OneOf_ConvictionVoting_Target_target(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ConvictionVoting_Target_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_Target_TargetId{
        TargetId: To_ConvictionVoting_Id(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_Target_TargetIndex{
        TargetIndex: To_ConvictionVoting_Index(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_Target_TargetRaw{
        TargetRaw: To_ConvictionVoting_Raw(param),
        }
    case 3:
        return &pbgear.ConvictionVoting_Target_TargetAddress32{
        TargetAddress32: To_ConvictionVoting_Address32(param),
        }
    case 4:
        return &pbgear.ConvictionVoting_Target_TargetAddress20{
        TargetAddress20: To_ConvictionVoting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVoting_To(in any) *pbgear.ConvictionVoting_To {
    out := &pbgear.ConvictionVoting_To{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field To
    v0 := To_OneOf_ConvictionVoting_To_to(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ConvictionVoting_To_to (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_To_ToId{
        ToId: To_ConvictionVoting_Id(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_To_ToIndex{
        ToIndex: To_ConvictionVoting_Index(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_To_ToRaw{
        ToRaw: To_ConvictionVoting_Raw(param),
        }
    case 3:
        return &pbgear.ConvictionVoting_To_ToAddress32{
        ToAddress32: To_ConvictionVoting_Address32(param),
        }
    case 4:
        return &pbgear.ConvictionVoting_To_ToAddress20{
        ToAddress20: To_ConvictionVoting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVoting_TupleNull(in any) *pbgear.ConvictionVoting_TupleNull {
    out := &pbgear.ConvictionVoting_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_ConvictionVoting_UndelegateCall(in any) *pbgear.ConvictionVoting_UndelegateCall {
    out := &pbgear.ConvictionVoting_UndelegateCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Class
    out.Class = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_ConvictionVoting_UndelegateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ConvictionVoting_UndelegateCall(in)
    out := &pbgearextrinsic.RuntimeCall_ConvictionvotingUndelegateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ConvictionvotingUndelegateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_ConvictionVoting_UnlockCall(in any) *pbgear.ConvictionVoting_UnlockCall {
    out := &pbgear.ConvictionVoting_UnlockCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Class
    out.Class = To_uint32(v.ValueAt(0))
    // oneOf field Target
    v1 := To_OneOf_ConvictionVoting_UnlockCall_target(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_ConvictionVoting_UnlockCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ConvictionVoting_UnlockCall(in)
    out := &pbgearextrinsic.RuntimeCall_ConvictionvotingUnlockCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ConvictionvotingUnlockCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_ConvictionVoting_UnlockCall_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_UnlockCall_TargetId{
        TargetId: To_ConvictionVoting_Id(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_UnlockCall_TargetIndex{
        TargetIndex: To_ConvictionVoting_Index(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_UnlockCall_TargetRaw{
        TargetRaw: To_ConvictionVoting_Raw(param),
        }
    case 3:
        return &pbgear.ConvictionVoting_UnlockCall_TargetAddress32{
        TargetAddress32: To_ConvictionVoting_Address32(param),
        }
    case 4:
        return &pbgear.ConvictionVoting_UnlockCall_TargetAddress20{
        TargetAddress20: To_ConvictionVoting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVoting_Vote(in any) *pbgear.ConvictionVoting_Vote {
    out := &pbgear.ConvictionVoting_Vote{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Vote
    v0 := To_OneOf_ConvictionVoting_Vote_vote(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Vote").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ConvictionVoting_Vote_vote (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_Vote_VoteStandard{
        VoteStandard: To_ConvictionVoting_Standard(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_Vote_VoteSplit{
        VoteSplit: To_ConvictionVoting_Split(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_Vote_VoteSplitAbstain{
        VoteSplitAbstain: To_ConvictionVoting_SplitAbstain(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ConvictionVoting_VoteCall(in any) *pbgear.ConvictionVoting_VoteCall {
    out := &pbgear.ConvictionVoting_VoteCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PollIndex
    out.PollIndex = To_uint32(v.ValueAt(0))
    // oneOf field Vote
    v1 := To_OneOf_ConvictionVoting_VoteCall_vote(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Vote").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_ConvictionVoting_VoteCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ConvictionVoting_VoteCall(in)
    out := &pbgearextrinsic.RuntimeCall_ConvictionvotingVoteCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ConvictionvotingVoteCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_ConvictionVoting_VoteCall_vote (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ConvictionVoting_VoteCall_VoteStandard{
        VoteStandard: To_ConvictionVoting_Standard(param),
        }
    case 1:
        return &pbgear.ConvictionVoting_VoteCall_VoteSplit{
        VoteSplit: To_ConvictionVoting_Split(param),
        }
    case 2:
        return &pbgear.ConvictionVoting_VoteCall_VoteSplitAbstain{
        VoteSplitAbstain: To_ConvictionVoting_SplitAbstain(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Corruption(in any) *pbgear.Corruption {
    out := &pbgear.Corruption{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_DestinationStatus(in any) *pbgear.DestinationStatus {
    out := &pbgear.DestinationStatus{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field DestinationStatus
    v0 := To_OneOf_DestinationStatus_destination_status(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("DestinationStatus").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_DestinationStatus_destination_status (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.DestinationStatus_DestinationStatusFree{
        DestinationStatusFree: To_Free(param),
        }
    case 1:
        return &pbgear.DestinationStatus_DestinationStatusReserved{
        DestinationStatusReserved: To_Reserved(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_DispatchError(in any) *pbgear.DispatchError {
    out := &pbgear.DispatchError{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field DispatchError
    v0 := To_OneOf_DispatchError_dispatch_error(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("DispatchError").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_DispatchError_dispatch_error (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.DispatchError_DispatchErrorOther{
        DispatchErrorOther: To_Other(param),
        }
    case 1:
        return &pbgear.DispatchError_DispatchErrorCannotLookup{
        DispatchErrorCannotLookup: To_CannotLookup(param),
        }
    case 2:
        return &pbgear.DispatchError_DispatchErrorBadOrigin{
        DispatchErrorBadOrigin: To_BadOrigin(param),
        }
    case 3:
        return &pbgear.DispatchError_DispatchErrorModule{
        DispatchErrorModule: To_Module(param),
        }
    case 4:
        return &pbgear.DispatchError_DispatchErrorConsumerRemaining{
        DispatchErrorConsumerRemaining: To_ConsumerRemaining(param),
        }
    case 5:
        return &pbgear.DispatchError_DispatchErrorNoProviders{
        DispatchErrorNoProviders: To_NoProviders(param),
        }
    case 6:
        return &pbgear.DispatchError_DispatchErrorTooManyConsumers{
        DispatchErrorTooManyConsumers: To_TooManyConsumers(param),
        }
    case 7:
        return &pbgear.DispatchError_DispatchErrorToken{
        DispatchErrorToken: To_Token(param),
        }
    case 8:
        return &pbgear.DispatchError_DispatchErrorArithmetic{
        DispatchErrorArithmetic: To_Arithmetic(param),
        }
    case 9:
        return &pbgear.DispatchError_DispatchErrorTransactional{
        DispatchErrorTransactional: To_Transactional(param),
        }
    case 10:
        return &pbgear.DispatchError_DispatchErrorExhausted{
        DispatchErrorExhausted: To_Exhausted(param),
        }
    case 11:
        return &pbgear.DispatchError_DispatchErrorCorruption{
        DispatchErrorCorruption: To_Corruption(param),
        }
    case 12:
        return &pbgear.DispatchError_DispatchErrorUnavailable{
        DispatchErrorUnavailable: To_Unavailable(param),
        }
    case 13:
        return &pbgear.DispatchError_DispatchErrorRootNotAllowed{
        DispatchErrorRootNotAllowed: To_RootNotAllowed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_DivisionByZero(in any) *pbgear.DivisionByZero {
    out := &pbgear.DivisionByZero{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhaseElectionFailedEvent(in any) *pbgear.ElectionProviderMultiPhaseElectionFailedEvent {
    out := &pbgear.ElectionProviderMultiPhaseElectionFailedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhaseElectionFinalizedEvent(in any) *pbgear.ElectionProviderMultiPhaseElectionFinalizedEvent {
    out := &pbgear.ElectionProviderMultiPhaseElectionFinalizedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhasePallet(in any) *pbgear.ElectionProviderMultiPhasePallet {
    out := &pbgear.ElectionProviderMultiPhasePallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_ElectionProviderMultiPhasePallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ElectionProviderMultiPhasePallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ElectionProviderMultiPhasePallet_CallSubmitUnsignedCall{
        CallSubmitUnsignedCall: To_ElectionProviderMultiPhase_SubmitUnsignedCall(param),
        }
    case 1:
        return &pbgear.ElectionProviderMultiPhasePallet_CallSetMinimumUntrustedScoreCall{
        CallSetMinimumUntrustedScoreCall: To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(param),
        }
    case 2:
        return &pbgear.ElectionProviderMultiPhasePallet_CallSetEmergencyElectionResultCall{
        CallSetEmergencyElectionResultCall: To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(param),
        }
    case 3:
        return &pbgear.ElectionProviderMultiPhasePallet_CallSubmitCall{
        CallSubmitCall: To_ElectionProviderMultiPhase_SubmitCall(param),
        }
    case 4:
        return &pbgear.ElectionProviderMultiPhasePallet_CallGovernanceFallbackCall{
        CallGovernanceFallbackCall: To_ElectionProviderMultiPhase_GovernanceFallbackCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ElectionProviderMultiPhasePhaseTransitionedEvent(in any) *pbgear.ElectionProviderMultiPhasePhaseTransitionedEvent {
    out := &pbgear.ElectionProviderMultiPhasePhaseTransitionedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhaseRewardedEvent(in any) *pbgear.ElectionProviderMultiPhaseRewardedEvent {
    out := &pbgear.ElectionProviderMultiPhaseRewardedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhaseSlashedEvent(in any) *pbgear.ElectionProviderMultiPhaseSlashedEvent {
    out := &pbgear.ElectionProviderMultiPhaseSlashedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhaseSolutionStoredEvent(in any) *pbgear.ElectionProviderMultiPhaseSolutionStoredEvent {
    out := &pbgear.ElectionProviderMultiPhaseSolutionStoredEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhase_Compute(in any) *pbgear.ElectionProviderMultiPhase_Compute {
    out := &pbgear.ElectionProviderMultiPhase_Compute{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Compute
    v0 := To_OneOf_ElectionProviderMultiPhase_Compute_compute(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Compute").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ElectionProviderMultiPhase_Compute_compute (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ElectionProviderMultiPhase_Compute_ComputeOnChain{
        ComputeOnChain: To_ElectionProviderMultiPhase_OnChain(param),
        }
    case 1:
        return &pbgear.ElectionProviderMultiPhase_Compute_ComputeSigned{
        ComputeSigned: To_ElectionProviderMultiPhase_Signed(param),
        }
    case 2:
        return &pbgear.ElectionProviderMultiPhase_Compute_ComputeUnsigned{
        ComputeUnsigned: To_ElectionProviderMultiPhase_Unsigned(param),
        }
    case 3:
        return &pbgear.ElectionProviderMultiPhase_Compute_ComputeFallback{
        ComputeFallback: To_ElectionProviderMultiPhase_Fallback(param),
        }
    case 4:
        return &pbgear.ElectionProviderMultiPhase_Compute_ComputeEmergency{
        ComputeEmergency: To_ElectionProviderMultiPhase_Emergency(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ElectionProviderMultiPhase_Emergency(in any) *pbgear.ElectionProviderMultiPhase_Emergency {
    out := &pbgear.ElectionProviderMultiPhase_Emergency{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhase_Fallback(in any) *pbgear.ElectionProviderMultiPhase_Fallback {
    out := &pbgear.ElectionProviderMultiPhase_Fallback{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhase_From(in any) *pbgear.ElectionProviderMultiPhase_From {
    out := &pbgear.ElectionProviderMultiPhase_From{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field From
    v0 := To_OneOf_ElectionProviderMultiPhase_From_from(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("From").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ElectionProviderMultiPhase_From_from (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ElectionProviderMultiPhase_From_FromOff{
        FromOff: To_ElectionProviderMultiPhase_Off(param),
        }
    case 1:
        return &pbgear.ElectionProviderMultiPhase_From_FromSigned{
        FromSigned: To_ElectionProviderMultiPhase_Signed(param),
        }
    case 2:
        return &pbgear.ElectionProviderMultiPhase_From_FromUnsigned{
        FromUnsigned: To_ElectionProviderMultiPhase_Unsigned(param),
        }
    case 3:
        return &pbgear.ElectionProviderMultiPhase_From_FromEmergency{
        FromEmergency: To_ElectionProviderMultiPhase_Emergency(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ElectionProviderMultiPhase_GovernanceFallbackCall(in any) *pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall {
    out := &pbgear.ElectionProviderMultiPhase_GovernanceFallbackCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field MaybeMaxVoters
    out.MaybeMaxVoters = To_Optional_uint32(v.ValueAt(0))
    // primitive field MaybeMaxTargets
    out.MaybeMaxTargets = To_Optional_uint32(v.ValueAt(1))

    return out //from message
}
func To_ElectionProviderMultiPhase_GovernanceFallbackCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ElectionProviderMultiPhase_GovernanceFallbackCall(in)
    out := &pbgearextrinsic.RuntimeCall_ElectionprovidermultiphaseGovernanceFallbackCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ElectionprovidermultiphaseGovernanceFallbackCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_ElectionProviderMultiPhase_Off(in any) *pbgear.ElectionProviderMultiPhase_Off {
    out := &pbgear.ElectionProviderMultiPhase_Off{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhase_OnChain(in any) *pbgear.ElectionProviderMultiPhase_OnChain {
    out := &pbgear.ElectionProviderMultiPhase_OnChain{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution {
    out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution{}
    v := in.(registry.Valuable)
    _ = v

    // field Solution To_VaraRuntimeNposSolution16(w)
    out.Solution = To_VaraRuntimeNposSolution16(v.ValueAt(0))
    // field Score To_SpNposElectionsElectionScore(w)
    out.Score = To_SpNposElectionsElectionScore(v.ValueAt(1))
    // primitive field Round
    out.Round = To_uint32(v.ValueAt(2))

    return out //from message
}
    


func To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(in any) *pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize {
    out := &pbgear.ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Voters
    out.Voters = To_uint32(v.ValueAt(0))
    // primitive field Targets
    out.Targets = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(in any) *pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall {
    out := &pbgear.ElectionProviderMultiPhase_SetEmergencyElectionResultCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Supports
    out.Supports = To_Repeated_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_supports(v.ValueAt(0))

    return out //from message
}
func To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall(in)
    out := &pbgearextrinsic.RuntimeCall_ElectionprovidermultiphaseSetEmergencyElectionResultCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ElectionprovidermultiphaseSetEmergencyElectionResultCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_supports(in any) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(in any) *pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall {
    out := &pbgear.ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall{}
    v := in.(registry.Valuable)
    _ = v

    // optional field MaybeNextScore
    if v.HasValue() {
        out.MaybeNextScore = To_Optional_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_maybe_next_score(v.ValueAt(0))
    }

    return out //from message
}
func To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall(in)
    out := &pbgearextrinsic.RuntimeCall_ElectionprovidermultiphaseSetMinimumUntrustedScoreCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ElectionprovidermultiphaseSetMinimumUntrustedScoreCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_maybe_next_score(in any) *pbgear.SpNposElectionsElectionScore {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_ElectionProviderMultiPhase_Signed(in any) *pbgear.ElectionProviderMultiPhase_Signed {
    out := &pbgear.ElectionProviderMultiPhase_Signed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ElectionProviderMultiPhase_SubmitCall(in any) *pbgear.ElectionProviderMultiPhase_SubmitCall {
    out := &pbgear.ElectionProviderMultiPhase_SubmitCall{}
    v := in.(registry.Valuable)
    _ = v

    // field RawSolution To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(w)
    out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(v.ValueAt(0))

    return out //from message
}
func To_ElectionProviderMultiPhase_SubmitCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ElectionProviderMultiPhase_SubmitCall(in)
    out := &pbgearextrinsic.RuntimeCall_ElectionprovidermultiphaseSubmitCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ElectionprovidermultiphaseSubmitCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_ElectionProviderMultiPhase_SubmitUnsignedCall(in any) *pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall {
    out := &pbgear.ElectionProviderMultiPhase_SubmitUnsignedCall{}
    v := in.(registry.Valuable)
    _ = v

    // field RawSolution To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(w)
    out.RawSolution = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseRawSolution(v.ValueAt(0))
    // field Witness To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(w)
    out.Witness = To_ElectionProviderMultiPhase_PalletElectionProviderMultiPhaseSolutionOrSnapshotSize(v.ValueAt(1))

    return out //from message
}
func To_ElectionProviderMultiPhase_SubmitUnsignedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ElectionProviderMultiPhase_SubmitUnsignedCall(in)
    out := &pbgearextrinsic.RuntimeCall_ElectionprovidermultiphaseSubmitUnsignedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ElectionprovidermultiphaseSubmitUnsignedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_ElectionProviderMultiPhase_To(in any) *pbgear.ElectionProviderMultiPhase_To {
    out := &pbgear.ElectionProviderMultiPhase_To{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field To
    v0 := To_OneOf_ElectionProviderMultiPhase_To_to(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ElectionProviderMultiPhase_To_to (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ElectionProviderMultiPhase_To_ToOff{
        ToOff: To_ElectionProviderMultiPhase_Off(param),
        }
    case 1:
        return &pbgear.ElectionProviderMultiPhase_To_ToSigned{
        ToSigned: To_ElectionProviderMultiPhase_Signed(param),
        }
    case 2:
        return &pbgear.ElectionProviderMultiPhase_To_ToUnsigned{
        ToUnsigned: To_ElectionProviderMultiPhase_Unsigned(param),
        }
    case 3:
        return &pbgear.ElectionProviderMultiPhase_To_ToEmergency{
        ToEmergency: To_ElectionProviderMultiPhase_Emergency(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ElectionProviderMultiPhase_TupleBooluint32(in any) *pbgear.ElectionProviderMultiPhase_TupleBooluint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleBooluint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bool(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport(in any) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport {
    out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32SpNposElectionsSupport{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field Value1 To_SpNposElectionsSupport(w)
    out.Value1 = To_SpNposElectionsSupport(v.ValueAt(1))

    return out //from message
}
    


func To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(in any) *pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
    out := &pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // repeated field Value1
    out.Value1 = To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32_value1(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32_value1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // field Value1 To_SpArithmeticPerThingsPerU16(w)
    out.Value1 = To_SpArithmeticPerThingsPerU16(v.ValueAt(1))

    return out //from message
}
    

func To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // field Value1 To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(w)
    out.Value1 = To_ElectionProviderMultiPhase_TupleUint32SpArithmeticPerThingsPerU16(v.ValueAt(1))
    // primitive field Value2
    out.Value2 = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_ElectionProviderMultiPhase_TupleUint32Uint32(in any) *pbgear.ElectionProviderMultiPhase_TupleUint32Uint32 {
    out := &pbgear.ElectionProviderMultiPhase_TupleUint32Uint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_ElectionProviderMultiPhase_Unsigned(in any) *pbgear.ElectionProviderMultiPhase_Unsigned {
    out := &pbgear.ElectionProviderMultiPhase_Unsigned{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Entry(in any) *pbgear.Entry {
    out := &pbgear.Entry{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Entry
    v0 := To_OneOf_Entry_entry(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Entry").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Entry_entry (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Entry_EntryInit{
        EntryInit: To_Init(param),
        }
    case 1:
        return &pbgear.Entry_EntryHandle{
        EntryHandle: To_Handle(param),
        }
    case 2:
        return &pbgear.Entry_EntryReply{
        EntryReply: To_Reply(param),
        }
    case 3:
        return &pbgear.Entry_EntrySignal{
        EntrySignal: To_Signal(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Err(in any) *pbgear.Err {
    out := &pbgear.Err{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Err_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Err_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Err_Value0Other{
        Value0Other: To_Other(param),
        }
    case 1:
        return &pbgear.Err_Value0CannotLookup{
        Value0CannotLookup: To_CannotLookup(param),
        }
    case 2:
        return &pbgear.Err_Value0BadOrigin{
        Value0BadOrigin: To_BadOrigin(param),
        }
    case 3:
        return &pbgear.Err_Value0Module{
        Value0Module: To_Module(param),
        }
    case 4:
        return &pbgear.Err_Value0ConsumerRemaining{
        Value0ConsumerRemaining: To_ConsumerRemaining(param),
        }
    case 5:
        return &pbgear.Err_Value0NoProviders{
        Value0NoProviders: To_NoProviders(param),
        }
    case 6:
        return &pbgear.Err_Value0TooManyConsumers{
        Value0TooManyConsumers: To_TooManyConsumers(param),
        }
    case 7:
        return &pbgear.Err_Value0Token{
        Value0Token: To_Token(param),
        }
    case 8:
        return &pbgear.Err_Value0Arithmetic{
        Value0Arithmetic: To_Arithmetic(param),
        }
    case 9:
        return &pbgear.Err_Value0Transactional{
        Value0Transactional: To_Transactional(param),
        }
    case 10:
        return &pbgear.Err_Value0Exhausted{
        Value0Exhausted: To_Exhausted(param),
        }
    case 11:
        return &pbgear.Err_Value0Corruption{
        Value0Corruption: To_Corruption(param),
        }
    case 12:
        return &pbgear.Err_Value0Unavailable{
        Value0Unavailable: To_Unavailable(param),
        }
    case 13:
        return &pbgear.Err_Value0RootNotAllowed{
        Value0RootNotAllowed: To_RootNotAllowed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Error(in any) *pbgear.Error {
    out := &pbgear.Error{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Error
    v0 := To_OneOf_Error_error(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Error").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Error_error (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Error_ErrorOther{
        ErrorOther: To_Other(param),
        }
    case 1:
        return &pbgear.Error_ErrorCannotLookup{
        ErrorCannotLookup: To_CannotLookup(param),
        }
    case 2:
        return &pbgear.Error_ErrorBadOrigin{
        ErrorBadOrigin: To_BadOrigin(param),
        }
    case 3:
        return &pbgear.Error_ErrorModule{
        ErrorModule: To_Module(param),
        }
    case 4:
        return &pbgear.Error_ErrorConsumerRemaining{
        ErrorConsumerRemaining: To_ConsumerRemaining(param),
        }
    case 5:
        return &pbgear.Error_ErrorNoProviders{
        ErrorNoProviders: To_NoProviders(param),
        }
    case 6:
        return &pbgear.Error_ErrorTooManyConsumers{
        ErrorTooManyConsumers: To_TooManyConsumers(param),
        }
    case 7:
        return &pbgear.Error_ErrorToken{
        ErrorToken: To_Token(param),
        }
    case 8:
        return &pbgear.Error_ErrorArithmetic{
        ErrorArithmetic: To_Arithmetic(param),
        }
    case 9:
        return &pbgear.Error_ErrorTransactional{
        ErrorTransactional: To_Transactional(param),
        }
    case 10:
        return &pbgear.Error_ErrorExhausted{
        ErrorExhausted: To_Exhausted(param),
        }
    case 11:
        return &pbgear.Error_ErrorCorruption{
        ErrorCorruption: To_Corruption(param),
        }
    case 12:
        return &pbgear.Error_ErrorUnavailable{
        ErrorUnavailable: To_Unavailable(param),
        }
    case 13:
        return &pbgear.Error_ErrorRootNotAllowed{
        ErrorRootNotAllowed: To_RootNotAllowed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Execution(in any) *pbgear.Execution {
    out := &pbgear.Execution{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Execution_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Execution_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Execution_Value0RanOutOfGas{
        Value0RanOutOfGas: To_RanOutOfGas(param),
        }
    case 1:
        return &pbgear.Execution_Value0MemoryOverflow{
        Value0MemoryOverflow: To_MemoryOverflow(param),
        }
    case 2:
        return &pbgear.Execution_Value0BackendError{
        Value0BackendError: To_BackendError(param),
        }
    case 3:
        return &pbgear.Execution_Value0UserspacePanic{
        Value0UserspacePanic: To_UserspacePanic(param),
        }
    case 4:
        return &pbgear.Execution_Value0UnreachableInstruction{
        Value0UnreachableInstruction: To_UnreachableInstruction(param),
        }
    case 5:
        return &pbgear.Execution_Value0StackLimitExceeded{
        Value0StackLimitExceeded: To_StackLimitExceeded(param),
        }
    case 255:
        return &pbgear.Execution_Value0Unsupported{
        Value0Unsupported: To_Unsupported(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Exhausted(in any) *pbgear.Exhausted {
    out := &pbgear.Exhausted{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ExpirationChanged(in any) *pbgear.ExpirationChanged {
    out := &pbgear.ExpirationChanged{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Expiration
    out.Expiration = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Failed(in any) *pbgear.Failed {
    out := &pbgear.Failed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FailedToCreateProgram(in any) *pbgear.FailedToCreateProgram {
    out := &pbgear.FailedToCreateProgram{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_FailedToCreateProgram_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FailedToCreateProgram_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FailedToCreateProgram_Value0CodeNotExists{
        Value0CodeNotExists: To_CodeNotExists(param),
        }
    case 255:
        return &pbgear.FailedToCreateProgram_Value0Unsupported{
        Value0Unsupported: To_Unsupported(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Fellows(in any) *pbgear.Fellows {
    out := &pbgear.Fellows{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship1Dan(in any) *pbgear.Fellowship1Dan {
    out := &pbgear.Fellowship1Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship2Dan(in any) *pbgear.Fellowship2Dan {
    out := &pbgear.Fellowship2Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship3Dan(in any) *pbgear.Fellowship3Dan {
    out := &pbgear.Fellowship3Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship4Dan(in any) *pbgear.Fellowship4Dan {
    out := &pbgear.Fellowship4Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship5Dan(in any) *pbgear.Fellowship5Dan {
    out := &pbgear.Fellowship5Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship6Dan(in any) *pbgear.Fellowship6Dan {
    out := &pbgear.Fellowship6Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship7Dan(in any) *pbgear.Fellowship7Dan {
    out := &pbgear.Fellowship7Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship8Dan(in any) *pbgear.Fellowship8Dan {
    out := &pbgear.Fellowship8Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Fellowship9Dan(in any) *pbgear.Fellowship9Dan {
    out := &pbgear.Fellowship9Dan{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipAdmin(in any) *pbgear.FellowshipAdmin {
    out := &pbgear.FellowshipAdmin{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipCollectiveMemberAddedEvent(in any) *pbgear.FellowshipCollectiveMemberAddedEvent {
    out := &pbgear.FellowshipCollectiveMemberAddedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipCollectiveMemberRemovedEvent(in any) *pbgear.FellowshipCollectiveMemberRemovedEvent {
    out := &pbgear.FellowshipCollectiveMemberRemovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipCollectivePallet(in any) *pbgear.FellowshipCollectivePallet {
    out := &pbgear.FellowshipCollectivePallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_FellowshipCollectivePallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipCollectivePallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipCollectivePallet_CallAddMemberCall{
        CallAddMemberCall: To_FellowshipCollective_AddMemberCall(param),
        }
    case 1:
        return &pbgear.FellowshipCollectivePallet_CallPromoteMemberCall{
        CallPromoteMemberCall: To_FellowshipCollective_PromoteMemberCall(param),
        }
    case 2:
        return &pbgear.FellowshipCollectivePallet_CallDemoteMemberCall{
        CallDemoteMemberCall: To_FellowshipCollective_DemoteMemberCall(param),
        }
    case 3:
        return &pbgear.FellowshipCollectivePallet_CallRemoveMemberCall{
        CallRemoveMemberCall: To_FellowshipCollective_RemoveMemberCall(param),
        }
    case 4:
        return &pbgear.FellowshipCollectivePallet_CallVoteCall{
        CallVoteCall: To_FellowshipCollective_VoteCall(param),
        }
    case 5:
        return &pbgear.FellowshipCollectivePallet_CallCleanupPollCall{
        CallCleanupPollCall: To_FellowshipCollective_CleanupPollCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipCollectiveRankChangedEvent(in any) *pbgear.FellowshipCollectiveRankChangedEvent {
    out := &pbgear.FellowshipCollectiveRankChangedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipCollectiveVotedEvent(in any) *pbgear.FellowshipCollectiveVotedEvent {
    out := &pbgear.FellowshipCollectiveVotedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipCollective_AddMemberCall(in any) *pbgear.FellowshipCollective_AddMemberCall {
    out := &pbgear.FellowshipCollective_AddMemberCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_FellowshipCollective_AddMemberCall_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_FellowshipCollective_AddMemberCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipCollective_AddMemberCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipcollectiveAddMemberCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipcollectiveAddMemberCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_FellowshipCollective_AddMemberCall_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipCollective_AddMemberCall_WhoId{
        WhoId: To_FellowshipCollective_Id(param),
        }
    case 1:
        return &pbgear.FellowshipCollective_AddMemberCall_WhoIndex{
        WhoIndex: To_FellowshipCollective_Index(param),
        }
    case 2:
        return &pbgear.FellowshipCollective_AddMemberCall_WhoRaw{
        WhoRaw: To_FellowshipCollective_Raw(param),
        }
    case 3:
        return &pbgear.FellowshipCollective_AddMemberCall_WhoAddress32{
        WhoAddress32: To_FellowshipCollective_Address32(param),
        }
    case 4:
        return &pbgear.FellowshipCollective_AddMemberCall_WhoAddress20{
        WhoAddress20: To_FellowshipCollective_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipCollective_Address20(in any) *pbgear.FellowshipCollective_Address20 {
    out := &pbgear.FellowshipCollective_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_FellowshipCollective_Address32(in any) *pbgear.FellowshipCollective_Address32 {
    out := &pbgear.FellowshipCollective_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_FellowshipCollective_CleanupPollCall(in any) *pbgear.FellowshipCollective_CleanupPollCall {
    out := &pbgear.FellowshipCollective_CleanupPollCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PollIndex
    out.PollIndex = To_uint32(v.ValueAt(0))
    // primitive field Max
    out.Max = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_FellowshipCollective_CleanupPollCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipCollective_CleanupPollCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipcollectiveCleanupPollCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipcollectiveCleanupPollCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipCollective_DemoteMemberCall(in any) *pbgear.FellowshipCollective_DemoteMemberCall {
    out := &pbgear.FellowshipCollective_DemoteMemberCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_FellowshipCollective_DemoteMemberCall_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_FellowshipCollective_DemoteMemberCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipCollective_DemoteMemberCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipcollectiveDemoteMemberCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipcollectiveDemoteMemberCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_FellowshipCollective_DemoteMemberCall_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipCollective_DemoteMemberCall_WhoId{
        WhoId: To_FellowshipCollective_Id(param),
        }
    case 1:
        return &pbgear.FellowshipCollective_DemoteMemberCall_WhoIndex{
        WhoIndex: To_FellowshipCollective_Index(param),
        }
    case 2:
        return &pbgear.FellowshipCollective_DemoteMemberCall_WhoRaw{
        WhoRaw: To_FellowshipCollective_Raw(param),
        }
    case 3:
        return &pbgear.FellowshipCollective_DemoteMemberCall_WhoAddress32{
        WhoAddress32: To_FellowshipCollective_Address32(param),
        }
    case 4:
        return &pbgear.FellowshipCollective_DemoteMemberCall_WhoAddress20{
        WhoAddress20: To_FellowshipCollective_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipCollective_Id(in any) *pbgear.FellowshipCollective_Id {
    out := &pbgear.FellowshipCollective_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_FellowshipCollective_Index(in any) *pbgear.FellowshipCollective_Index {
    out := &pbgear.FellowshipCollective_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_FellowshipCollective_TupleNull(w)
    out.Value0 = To_FellowshipCollective_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_FellowshipCollective_PromoteMemberCall(in any) *pbgear.FellowshipCollective_PromoteMemberCall {
    out := &pbgear.FellowshipCollective_PromoteMemberCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_FellowshipCollective_PromoteMemberCall_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_FellowshipCollective_PromoteMemberCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipCollective_PromoteMemberCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipcollectivePromoteMemberCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipcollectivePromoteMemberCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_FellowshipCollective_PromoteMemberCall_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipCollective_PromoteMemberCall_WhoId{
        WhoId: To_FellowshipCollective_Id(param),
        }
    case 1:
        return &pbgear.FellowshipCollective_PromoteMemberCall_WhoIndex{
        WhoIndex: To_FellowshipCollective_Index(param),
        }
    case 2:
        return &pbgear.FellowshipCollective_PromoteMemberCall_WhoRaw{
        WhoRaw: To_FellowshipCollective_Raw(param),
        }
    case 3:
        return &pbgear.FellowshipCollective_PromoteMemberCall_WhoAddress32{
        WhoAddress32: To_FellowshipCollective_Address32(param),
        }
    case 4:
        return &pbgear.FellowshipCollective_PromoteMemberCall_WhoAddress20{
        WhoAddress20: To_FellowshipCollective_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipCollective_Raw(in any) *pbgear.FellowshipCollective_Raw {
    out := &pbgear.FellowshipCollective_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_FellowshipCollective_RemoveMemberCall(in any) *pbgear.FellowshipCollective_RemoveMemberCall {
    out := &pbgear.FellowshipCollective_RemoveMemberCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_FellowshipCollective_RemoveMemberCall_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))
    // primitive field MinRank
    out.MinRank = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_FellowshipCollective_RemoveMemberCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipCollective_RemoveMemberCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipcollectiveRemoveMemberCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipcollectiveRemoveMemberCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_FellowshipCollective_RemoveMemberCall_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipCollective_RemoveMemberCall_WhoId{
        WhoId: To_FellowshipCollective_Id(param),
        }
    case 1:
        return &pbgear.FellowshipCollective_RemoveMemberCall_WhoIndex{
        WhoIndex: To_FellowshipCollective_Index(param),
        }
    case 2:
        return &pbgear.FellowshipCollective_RemoveMemberCall_WhoRaw{
        WhoRaw: To_FellowshipCollective_Raw(param),
        }
    case 3:
        return &pbgear.FellowshipCollective_RemoveMemberCall_WhoAddress32{
        WhoAddress32: To_FellowshipCollective_Address32(param),
        }
    case 4:
        return &pbgear.FellowshipCollective_RemoveMemberCall_WhoAddress20{
        WhoAddress20: To_FellowshipCollective_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipCollective_TupleNull(in any) *pbgear.FellowshipCollective_TupleNull {
    out := &pbgear.FellowshipCollective_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_FellowshipCollective_VoteCall(in any) *pbgear.FellowshipCollective_VoteCall {
    out := &pbgear.FellowshipCollective_VoteCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Poll
    out.Poll = To_uint32(v.ValueAt(0))
    // primitive field Aye
    out.Aye = To_bool(v.ValueAt(1))

    return out //from message
}
func To_FellowshipCollective_VoteCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipCollective_VoteCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipcollectiveVoteCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipcollectiveVoteCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipCollective_Who(in any) *pbgear.FellowshipCollective_Who {
    out := &pbgear.FellowshipCollective_Who{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_FellowshipCollective_Who_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipCollective_Who_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipCollective_Who_WhoId{
        WhoId: To_FellowshipCollective_Id(param),
        }
    case 1:
        return &pbgear.FellowshipCollective_Who_WhoIndex{
        WhoIndex: To_FellowshipCollective_Index(param),
        }
    case 2:
        return &pbgear.FellowshipCollective_Who_WhoRaw{
        WhoRaw: To_FellowshipCollective_Raw(param),
        }
    case 3:
        return &pbgear.FellowshipCollective_Who_WhoAddress32{
        WhoAddress32: To_FellowshipCollective_Address32(param),
        }
    case 4:
        return &pbgear.FellowshipCollective_Who_WhoAddress20{
        WhoAddress20: To_FellowshipCollective_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipExperts(in any) *pbgear.FellowshipExperts {
    out := &pbgear.FellowshipExperts{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipInitiates(in any) *pbgear.FellowshipInitiates {
    out := &pbgear.FellowshipInitiates{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipMasters(in any) *pbgear.FellowshipMasters {
    out := &pbgear.FellowshipMasters{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaApprovedEvent(in any) *pbgear.FellowshipReferendaApprovedEvent {
    out := &pbgear.FellowshipReferendaApprovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaCancelledEvent(in any) *pbgear.FellowshipReferendaCancelledEvent {
    out := &pbgear.FellowshipReferendaCancelledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaConfirmAbortedEvent(in any) *pbgear.FellowshipReferendaConfirmAbortedEvent {
    out := &pbgear.FellowshipReferendaConfirmAbortedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaConfirmStartedEvent(in any) *pbgear.FellowshipReferendaConfirmStartedEvent {
    out := &pbgear.FellowshipReferendaConfirmStartedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaConfirmedEvent(in any) *pbgear.FellowshipReferendaConfirmedEvent {
    out := &pbgear.FellowshipReferendaConfirmedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaDecisionDepositPlacedEvent(in any) *pbgear.FellowshipReferendaDecisionDepositPlacedEvent {
    out := &pbgear.FellowshipReferendaDecisionDepositPlacedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaDecisionDepositRefundedEvent(in any) *pbgear.FellowshipReferendaDecisionDepositRefundedEvent {
    out := &pbgear.FellowshipReferendaDecisionDepositRefundedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaDecisionStartedEvent(in any) *pbgear.FellowshipReferendaDecisionStartedEvent {
    out := &pbgear.FellowshipReferendaDecisionStartedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaDepositSlashedEvent(in any) *pbgear.FellowshipReferendaDepositSlashedEvent {
    out := &pbgear.FellowshipReferendaDepositSlashedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaKilledEvent(in any) *pbgear.FellowshipReferendaKilledEvent {
    out := &pbgear.FellowshipReferendaKilledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaMetadataClearedEvent(in any) *pbgear.FellowshipReferendaMetadataClearedEvent {
    out := &pbgear.FellowshipReferendaMetadataClearedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaMetadataSetEvent(in any) *pbgear.FellowshipReferendaMetadataSetEvent {
    out := &pbgear.FellowshipReferendaMetadataSetEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaPallet(in any) *pbgear.FellowshipReferendaPallet {
    out := &pbgear.FellowshipReferendaPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_FellowshipReferendaPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipReferendaPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferendaPallet_CallSubmitCall{
        CallSubmitCall: To_FellowshipReferenda_SubmitCall(param),
        }
    case 1:
        return &pbgear.FellowshipReferendaPallet_CallPlaceDecisionDepositCall{
        CallPlaceDecisionDepositCall: To_FellowshipReferenda_PlaceDecisionDepositCall(param),
        }
    case 2:
        return &pbgear.FellowshipReferendaPallet_CallRefundDecisionDepositCall{
        CallRefundDecisionDepositCall: To_FellowshipReferenda_RefundDecisionDepositCall(param),
        }
    case 3:
        return &pbgear.FellowshipReferendaPallet_CallCancelCall{
        CallCancelCall: To_FellowshipReferenda_CancelCall(param),
        }
    case 4:
        return &pbgear.FellowshipReferendaPallet_CallKillCall{
        CallKillCall: To_FellowshipReferenda_KillCall(param),
        }
    case 5:
        return &pbgear.FellowshipReferendaPallet_CallNudgeReferendumCall{
        CallNudgeReferendumCall: To_FellowshipReferenda_NudgeReferendumCall(param),
        }
    case 6:
        return &pbgear.FellowshipReferendaPallet_CallOneFewerDecidingCall{
        CallOneFewerDecidingCall: To_FellowshipReferenda_OneFewerDecidingCall(param),
        }
    case 7:
        return &pbgear.FellowshipReferendaPallet_CallRefundSubmissionDepositCall{
        CallRefundSubmissionDepositCall: To_FellowshipReferenda_RefundSubmissionDepositCall(param),
        }
    case 8:
        return &pbgear.FellowshipReferendaPallet_CallSetMetadataCall{
        CallSetMetadataCall: To_FellowshipReferenda_SetMetadataCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipReferendaRejectedEvent(in any) *pbgear.FellowshipReferendaRejectedEvent {
    out := &pbgear.FellowshipReferendaRejectedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaSubmissionDepositRefundedEvent(in any) *pbgear.FellowshipReferendaSubmissionDepositRefundedEvent {
    out := &pbgear.FellowshipReferendaSubmissionDepositRefundedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaSubmittedEvent(in any) *pbgear.FellowshipReferendaSubmittedEvent {
    out := &pbgear.FellowshipReferendaSubmittedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferendaTimedOutEvent(in any) *pbgear.FellowshipReferendaTimedOutEvent {
    out := &pbgear.FellowshipReferendaTimedOutEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FellowshipReferenda_After(in any) *pbgear.FellowshipReferenda_After {
    out := &pbgear.FellowshipReferenda_After{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_FellowshipReferenda_At(in any) *pbgear.FellowshipReferenda_At {
    out := &pbgear.FellowshipReferenda_At{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_FellowshipReferenda_CancelCall(in any) *pbgear.FellowshipReferenda_CancelCall {
    out := &pbgear.FellowshipReferenda_CancelCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_FellowshipReferenda_CancelCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_CancelCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaCancelCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaCancelCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipReferenda_EnactmentMoment(in any) *pbgear.FellowshipReferenda_EnactmentMoment {
    out := &pbgear.FellowshipReferenda_EnactmentMoment{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field EnactmentMoment
    v0 := To_OneOf_FellowshipReferenda_EnactmentMoment_enactment_moment(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipReferenda_EnactmentMoment_enactment_moment (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_EnactmentMoment_EnactmentMomentAt{
        EnactmentMomentAt: To_FellowshipReferenda_At(param),
        }
    case 1:
        return &pbgear.FellowshipReferenda_EnactmentMoment_EnactmentMomentAfter{
        EnactmentMomentAfter: To_FellowshipReferenda_After(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipReferenda_Inline(in any) *pbgear.FellowshipReferenda_Inline {
    out := &pbgear.FellowshipReferenda_Inline{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_BoundedCollectionsBoundedVecBoundedVec(w)
    out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(v.ValueAt(0))

    return out //from message
}
    

func To_FellowshipReferenda_KillCall(in any) *pbgear.FellowshipReferenda_KillCall {
    out := &pbgear.FellowshipReferenda_KillCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_FellowshipReferenda_KillCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_KillCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaKillCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaKillCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipReferenda_Legacy(in any) *pbgear.FellowshipReferenda_Legacy {
    out := &pbgear.FellowshipReferenda_Legacy{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_FellowshipReferenda_Lookup(in any) *pbgear.FellowshipReferenda_Lookup {
    out := &pbgear.FellowshipReferenda_Lookup{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))
    // primitive field Len
    out.Len = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_FellowshipReferenda_NudgeReferendumCall(in any) *pbgear.FellowshipReferenda_NudgeReferendumCall {
    out := &pbgear.FellowshipReferenda_NudgeReferendumCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_FellowshipReferenda_NudgeReferendumCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_NudgeReferendumCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaNudgeReferendumCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaNudgeReferendumCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipReferenda_OneFewerDecidingCall(in any) *pbgear.FellowshipReferenda_OneFewerDecidingCall {
    out := &pbgear.FellowshipReferenda_OneFewerDecidingCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Track
    out.Track = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_FellowshipReferenda_OneFewerDecidingCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_OneFewerDecidingCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaOneFewerDecidingCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaOneFewerDecidingCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipReferenda_Origins(in any) *pbgear.FellowshipReferenda_Origins {
    out := &pbgear.FellowshipReferenda_Origins{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_FellowshipReferenda_Origins_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipReferenda_Origins_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_Origins_Value0StakingAdmin{
        Value0StakingAdmin: To_StakingAdmin(param),
        }
    case 1:
        return &pbgear.FellowshipReferenda_Origins_Value0Treasurer{
        Value0Treasurer: To_Treasurer(param),
        }
    case 2:
        return &pbgear.FellowshipReferenda_Origins_Value0FellowshipAdmin{
        Value0FellowshipAdmin: To_FellowshipAdmin(param),
        }
    case 3:
        return &pbgear.FellowshipReferenda_Origins_Value0GeneralAdmin{
        Value0GeneralAdmin: To_GeneralAdmin(param),
        }
    case 4:
        return &pbgear.FellowshipReferenda_Origins_Value0ReferendumCanceller{
        Value0ReferendumCanceller: To_ReferendumCanceller(param),
        }
    case 5:
        return &pbgear.FellowshipReferenda_Origins_Value0ReferendumKiller{
        Value0ReferendumKiller: To_ReferendumKiller(param),
        }
    case 6:
        return &pbgear.FellowshipReferenda_Origins_Value0SmallTipper{
        Value0SmallTipper: To_SmallTipper(param),
        }
    case 7:
        return &pbgear.FellowshipReferenda_Origins_Value0BigTipper{
        Value0BigTipper: To_BigTipper(param),
        }
    case 8:
        return &pbgear.FellowshipReferenda_Origins_Value0SmallSpender{
        Value0SmallSpender: To_SmallSpender(param),
        }
    case 9:
        return &pbgear.FellowshipReferenda_Origins_Value0MediumSpender{
        Value0MediumSpender: To_MediumSpender(param),
        }
    case 10:
        return &pbgear.FellowshipReferenda_Origins_Value0BigSpender{
        Value0BigSpender: To_BigSpender(param),
        }
    case 11:
        return &pbgear.FellowshipReferenda_Origins_Value0WhitelistedCaller{
        Value0WhitelistedCaller: To_WhitelistedCaller(param),
        }
    case 12:
        return &pbgear.FellowshipReferenda_Origins_Value0FellowshipInitiates{
        Value0FellowshipInitiates: To_FellowshipInitiates(param),
        }
    case 13:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellows{
        Value0Fellows: To_Fellows(param),
        }
    case 14:
        return &pbgear.FellowshipReferenda_Origins_Value0FellowshipExperts{
        Value0FellowshipExperts: To_FellowshipExperts(param),
        }
    case 15:
        return &pbgear.FellowshipReferenda_Origins_Value0FellowshipMasters{
        Value0FellowshipMasters: To_FellowshipMasters(param),
        }
    case 16:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship1Dan{
        Value0Fellowship1Dan: To_Fellowship1Dan(param),
        }
    case 17:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship2Dan{
        Value0Fellowship2Dan: To_Fellowship2Dan(param),
        }
    case 18:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship3Dan{
        Value0Fellowship3Dan: To_Fellowship3Dan(param),
        }
    case 19:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship4Dan{
        Value0Fellowship4Dan: To_Fellowship4Dan(param),
        }
    case 20:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship5Dan{
        Value0Fellowship5Dan: To_Fellowship5Dan(param),
        }
    case 21:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship6Dan{
        Value0Fellowship6Dan: To_Fellowship6Dan(param),
        }
    case 22:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship7Dan{
        Value0Fellowship7Dan: To_Fellowship7Dan(param),
        }
    case 23:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship8Dan{
        Value0Fellowship8Dan: To_Fellowship8Dan(param),
        }
    case 24:
        return &pbgear.FellowshipReferenda_Origins_Value0Fellowship9Dan{
        Value0Fellowship9Dan: To_Fellowship9Dan(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipReferenda_PlaceDecisionDepositCall(in any) *pbgear.FellowshipReferenda_PlaceDecisionDepositCall {
    out := &pbgear.FellowshipReferenda_PlaceDecisionDepositCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_FellowshipReferenda_PlaceDecisionDepositCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_PlaceDecisionDepositCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaPlaceDecisionDepositCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaPlaceDecisionDepositCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipReferenda_Proposal(in any) *pbgear.FellowshipReferenda_Proposal {
    out := &pbgear.FellowshipReferenda_Proposal{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Proposal
    v0 := To_OneOf_FellowshipReferenda_Proposal_proposal(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipReferenda_Proposal_proposal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_Proposal_ProposalLegacy{
        ProposalLegacy: To_FellowshipReferenda_Legacy(param),
        }
    case 1:
        return &pbgear.FellowshipReferenda_Proposal_ProposalInline{
        ProposalInline: To_FellowshipReferenda_Inline(param),
        }
    case 2:
        return &pbgear.FellowshipReferenda_Proposal_ProposalLookup{
        ProposalLookup: To_FellowshipReferenda_Lookup(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipReferenda_ProposalOrigin(in any) *pbgear.FellowshipReferenda_ProposalOrigin {
    out := &pbgear.FellowshipReferenda_ProposalOrigin{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ProposalOrigin
    v0 := To_OneOf_FellowshipReferenda_ProposalOrigin_proposal_origin(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipReferenda_ProposalOrigin_proposal_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_ProposalOrigin_ProposalOriginSystem{
        ProposalOriginSystem: To_FellowshipReferenda_System(param),
        }
    case 20:
        return &pbgear.FellowshipReferenda_ProposalOrigin_ProposalOriginOrigins{
        ProposalOriginOrigins: To_FellowshipReferenda_Origins(param),
        }
    case 2:
        return &pbgear.FellowshipReferenda_ProposalOrigin_ProposalOriginVoid{
        ProposalOriginVoid: To_FellowshipReferenda_Void(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipReferenda_RefundDecisionDepositCall(in any) *pbgear.FellowshipReferenda_RefundDecisionDepositCall {
    out := &pbgear.FellowshipReferenda_RefundDecisionDepositCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_FellowshipReferenda_RefundDecisionDepositCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_RefundDecisionDepositCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaRefundDecisionDepositCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaRefundDecisionDepositCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipReferenda_RefundSubmissionDepositCall(in any) *pbgear.FellowshipReferenda_RefundSubmissionDepositCall {
    out := &pbgear.FellowshipReferenda_RefundSubmissionDepositCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_FellowshipReferenda_RefundSubmissionDepositCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_RefundSubmissionDepositCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaRefundSubmissionDepositCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaRefundSubmissionDepositCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_FellowshipReferenda_SetMetadataCall(in any) *pbgear.FellowshipReferenda_SetMetadataCall {
    out := &pbgear.FellowshipReferenda_SetMetadataCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // optional field MaybeHash
    if v.HasValue() {
        out.MaybeHash = To_Optional_FellowshipReferenda_SetMetadataCall_maybe_hash(v.ValueAt(1))
    }

    return out //from message
}
func To_FellowshipReferenda_SetMetadataCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_SetMetadataCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaSetMetadataCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaSetMetadataCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_FellowshipReferenda_SetMetadataCall_maybe_hash(in any) *pbgear.PrimitiveTypesH256 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_FellowshipReferenda_SubmitCall(in any) *pbgear.FellowshipReferenda_SubmitCall {
    out := &pbgear.FellowshipReferenda_SubmitCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ProposalOrigin
    v0 := To_OneOf_FellowshipReferenda_SubmitCall_proposal_origin(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))
    // oneOf field Proposal
    v1 := To_OneOf_FellowshipReferenda_SubmitCall_proposal(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v1))
    // oneOf field EnactmentMoment
    v2 := To_OneOf_FellowshipReferenda_SubmitCall_enactment_moment(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v2))

    return out //from message
}
func To_FellowshipReferenda_SubmitCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_FellowshipReferenda_SubmitCall(in)
    out := &pbgearextrinsic.RuntimeCall_FellowshipreferendaSubmitCall{ }
    reflect.ValueOf(out).Elem().FieldByName("FellowshipreferendaSubmitCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_FellowshipReferenda_SubmitCall_proposal_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_SubmitCall_ProposalOriginSystem{
        ProposalOriginSystem: To_FellowshipReferenda_System(param),
        }
    case 20:
        return &pbgear.FellowshipReferenda_SubmitCall_ProposalOriginOrigins{
        ProposalOriginOrigins: To_FellowshipReferenda_Origins(param),
        }
    case 2:
        return &pbgear.FellowshipReferenda_SubmitCall_ProposalOriginVoid{
        ProposalOriginVoid: To_FellowshipReferenda_Void(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_FellowshipReferenda_SubmitCall_proposal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_SubmitCall_ProposalLegacy{
        ProposalLegacy: To_FellowshipReferenda_Legacy(param),
        }
    case 1:
        return &pbgear.FellowshipReferenda_SubmitCall_ProposalInline{
        ProposalInline: To_FellowshipReferenda_Inline(param),
        }
    case 2:
        return &pbgear.FellowshipReferenda_SubmitCall_ProposalLookup{
        ProposalLookup: To_FellowshipReferenda_Lookup(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_FellowshipReferenda_SubmitCall_enactment_moment (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_SubmitCall_EnactmentMomentAt{
        EnactmentMomentAt: To_FellowshipReferenda_At(param),
        }
    case 1:
        return &pbgear.FellowshipReferenda_SubmitCall_EnactmentMomentAfter{
        EnactmentMomentAfter: To_FellowshipReferenda_After(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipReferenda_System(in any) *pbgear.FellowshipReferenda_System {
    out := &pbgear.FellowshipReferenda_System{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_FellowshipReferenda_System_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipReferenda_System_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FellowshipReferenda_System_Value0Root{
        Value0Root: To_Root(param),
        }
    case 1:
        return &pbgear.FellowshipReferenda_System_Value0Signed{
        Value0Signed: To_Signed(param),
        }
    case 2:
        return &pbgear.FellowshipReferenda_System_Value0None{
        Value0None: To_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FellowshipReferenda_Void(in any) *pbgear.FellowshipReferenda_Void {
    out := &pbgear.FellowshipReferenda_Void{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_FellowshipReferenda_Void_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_FellowshipReferenda_Void_value0 (in any) any  {
    return nil
}
func To_FrameSupportDispatchDispatchInfo(in any) *pbgear.FrameSupportDispatchDispatchInfo {
    out := &pbgear.FrameSupportDispatchDispatchInfo{}
    v := in.(registry.Valuable)
    _ = v

    // field Weight To_SpWeightsWeightV2Weight(w)
    out.Weight = To_SpWeightsWeightV2Weight(v.ValueAt(0))
    // oneOf field Class
    v1 := To_OneOf_FrameSupportDispatchDispatchInfo_class(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Class").Set(reflect.ValueOf(v1))
    // oneOf field PaysFee
    v2 := To_OneOf_FrameSupportDispatchDispatchInfo_pays_fee(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("PaysFee").Set(reflect.ValueOf(v2))

    return out //from message
}
    

func To_OneOf_FrameSupportDispatchDispatchInfo_class (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FrameSupportDispatchDispatchInfo_ClassNormal{
        ClassNormal: To_Normal(param),
        }
    case 1:
        return &pbgear.FrameSupportDispatchDispatchInfo_ClassOperational{
        ClassOperational: To_Operational(param),
        }
    case 2:
        return &pbgear.FrameSupportDispatchDispatchInfo_ClassMandatory{
        ClassMandatory: To_Mandatory(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_FrameSupportDispatchDispatchInfo_pays_fee (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FrameSupportDispatchDispatchInfo_PaysFeeYes{
        PaysFeeYes: To_Yes(param),
        }
    case 1:
        return &pbgear.FrameSupportDispatchDispatchInfo_PaysFeeNo{
        PaysFeeNo: To_No(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_FrameSupportDispatchPostDispatchInfo(in any) *pbgear.FrameSupportDispatchPostDispatchInfo {
    out := &pbgear.FrameSupportDispatchPostDispatchInfo{}
    v := in.(registry.Valuable)
    _ = v

    // optional field ActualWeight
    if v.HasValue() {
        out.ActualWeight = To_Optional_FrameSupportDispatchPostDispatchInfo_actual_weight(v.ValueAt(0))
    }
    // oneOf field PaysFee
    v1 := To_OneOf_FrameSupportDispatchPostDispatchInfo_pays_fee(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("PaysFee").Set(reflect.ValueOf(v1))

    return out //from message
}
    
func To_Optional_FrameSupportDispatchPostDispatchInfo_actual_weight(in any) *pbgear.SpWeightsWeightV2Weight {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_OneOf_FrameSupportDispatchPostDispatchInfo_pays_fee (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.FrameSupportDispatchPostDispatchInfo_PaysFeeYes{
        PaysFeeYes: To_Yes(param),
        }
    case 1:
        return &pbgear.FrameSupportDispatchPostDispatchInfo_PaysFeeNo{
        PaysFeeNo: To_No(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Free(in any) *pbgear.Free {
    out := &pbgear.Free{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Frozen(in any) *pbgear.Frozen {
    out := &pbgear.Frozen{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_FundsUnavailable(in any) *pbgear.FundsUnavailable {
    out := &pbgear.FundsUnavailable{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearBankPallet(in any) *pbgear.GearBankPallet {
    out := &pbgear.GearBankPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearBuiltinPallet(in any) *pbgear.GearBuiltinPallet {
    out := &pbgear.GearBuiltinPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearCodeChangedEvent(in any) *pbgear.GearCodeChangedEvent {
    out := &pbgear.GearCodeChangedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearCoreBufferLimitedVec(in any) *pbgear.GearCoreBufferLimitedVec {
    out := &pbgear.GearCoreBufferLimitedVec{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_GearCoreMessageCommonReplyDetails(in any) *pbgear.GearCoreMessageCommonReplyDetails {
    out := &pbgear.GearCoreMessageCommonReplyDetails{}
    v := in.(registry.Valuable)
    _ = v

    // field To To_GprimitivesMessageId(w)
    out.To = To_GprimitivesMessageId(v.ValueAt(0))
    // oneOf field Code
    v1 := To_OneOf_GearCoreMessageCommonReplyDetails_code(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Code").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_GearCoreMessageCommonReplyDetails_code (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearCoreMessageCommonReplyDetails_CodeSuccess{
        CodeSuccess: To_Success(param),
        }
    case 1:
        return &pbgear.GearCoreMessageCommonReplyDetails_CodeError{
        CodeError: To_Error(param),
        }
    case 255:
        return &pbgear.GearCoreMessageCommonReplyDetails_CodeUnsupported{
        CodeUnsupported: To_Unsupported(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearCoreMessageUserUserMessage(in any) *pbgear.GearCoreMessageUserUserMessage {
    out := &pbgear.GearCoreMessageUserUserMessage{}
    v := in.(registry.Valuable)
    _ = v

    // field Id To_GprimitivesMessageId(w)
    out.Id = To_GprimitivesMessageId(v.ValueAt(0))
    // field Source To_GprimitivesActorId(w)
    out.Source = To_GprimitivesActorId(v.ValueAt(1))
    // field Destination To_GprimitivesActorId(w)
    out.Destination = To_GprimitivesActorId(v.ValueAt(2))
    // field Payload To_GearCoreBufferLimitedVec(w)
    out.Payload = To_GearCoreBufferLimitedVec(v.ValueAt(3))
    // primitive field Value
    out.Value = To_string(v.ValueAt(4))
    // optional field Details
    if v.HasValue() {
        out.Details = To_Optional_GearCoreMessageUserUserMessage_details(v.ValueAt(5))
    }

    return out //from message
}
    




func To_Optional_GearCoreMessageUserUserMessage_details(in any) *pbgear.GearCoreMessageCommonReplyDetails {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_GearGasPallet(in any) *pbgear.GearGasPallet {
    out := &pbgear.GearGasPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearMessageQueuedEvent(in any) *pbgear.GearMessageQueuedEvent {
    out := &pbgear.GearMessageQueuedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearMessageWaitedEvent(in any) *pbgear.GearMessageWaitedEvent {
    out := &pbgear.GearMessageWaitedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearMessageWokenEvent(in any) *pbgear.GearMessageWokenEvent {
    out := &pbgear.GearMessageWokenEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearMessagesDispatchedEvent(in any) *pbgear.GearMessagesDispatchedEvent {
    out := &pbgear.GearMessagesDispatchedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearMessengerPallet(in any) *pbgear.GearMessengerPallet {
    out := &pbgear.GearMessengerPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearPallet(in any) *pbgear.GearPallet {
    out := &pbgear.GearPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_GearPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_GearPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearPallet_CallUploadCodeCall{
        CallUploadCodeCall: To_Gear_UploadCodeCall(param),
        }
    case 1:
        return &pbgear.GearPallet_CallUploadProgramCall{
        CallUploadProgramCall: To_Gear_UploadProgramCall(param),
        }
    case 2:
        return &pbgear.GearPallet_CallCreateProgramCall{
        CallCreateProgramCall: To_Gear_CreateProgramCall(param),
        }
    case 3:
        return &pbgear.GearPallet_CallSendMessageCall{
        CallSendMessageCall: To_Gear_SendMessageCall(param),
        }
    case 4:
        return &pbgear.GearPallet_CallSendReplyCall{
        CallSendReplyCall: To_Gear_SendReplyCall(param),
        }
    case 5:
        return &pbgear.GearPallet_CallClaimValueCall{
        CallClaimValueCall: To_Gear_ClaimValueCall(param),
        }
    case 6:
        return &pbgear.GearPallet_CallRunCall{
        CallRunCall: To_Gear_RunCall(param),
        }
    case 7:
        return &pbgear.GearPallet_CallSetExecuteInherentCall{
        CallSetExecuteInherentCall: To_Gear_SetExecuteInherentCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearPaymentPallet(in any) *pbgear.GearPaymentPallet {
    out := &pbgear.GearPaymentPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearProgramChangedEvent(in any) *pbgear.GearProgramChangedEvent {
    out := &pbgear.GearProgramChangedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearProgramPallet(in any) *pbgear.GearProgramPallet {
    out := &pbgear.GearProgramPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearQueueNotProcessedEvent(in any) *pbgear.GearQueueNotProcessedEvent {
    out := &pbgear.GearQueueNotProcessedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearSchedulerPallet(in any) *pbgear.GearSchedulerPallet {
    out := &pbgear.GearSchedulerPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearUserMessageReadEvent(in any) *pbgear.GearUserMessageReadEvent {
    out := &pbgear.GearUserMessageReadEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearUserMessageSentEvent(in any) *pbgear.GearUserMessageSentEvent {
    out := &pbgear.GearUserMessageSentEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearVoucherPallet(in any) *pbgear.GearVoucherPallet {
    out := &pbgear.GearVoucherPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_GearVoucherPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_GearVoucherPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearVoucherPallet_CallIssueCall{
        CallIssueCall: To_GearVoucher_IssueCall(param),
        }
    case 1:
        return &pbgear.GearVoucherPallet_CallCallCall{
        CallCallCall: To_GearVoucher_CallCall(param),
        }
    case 2:
        return &pbgear.GearVoucherPallet_CallRevokeCall{
        CallRevokeCall: To_GearVoucher_RevokeCall(param),
        }
    case 3:
        return &pbgear.GearVoucherPallet_CallUpdateCall{
        CallUpdateCall: To_GearVoucher_UpdateCall(param),
        }
    case 4:
        return &pbgear.GearVoucherPallet_CallCallDeprecatedCall{
        CallCallDeprecatedCall: To_GearVoucher_CallDeprecatedCall(param),
        }
    case 5:
        return &pbgear.GearVoucherPallet_CallDeclineCall{
        CallDeclineCall: To_GearVoucher_DeclineCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearVoucherVoucherDeclinedEvent(in any) *pbgear.GearVoucherVoucherDeclinedEvent {
    out := &pbgear.GearVoucherVoucherDeclinedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearVoucherVoucherIssuedEvent(in any) *pbgear.GearVoucherVoucherIssuedEvent {
    out := &pbgear.GearVoucherVoucherIssuedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearVoucherVoucherRevokedEvent(in any) *pbgear.GearVoucherVoucherRevokedEvent {
    out := &pbgear.GearVoucherVoucherRevokedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearVoucherVoucherUpdatedEvent(in any) *pbgear.GearVoucherVoucherUpdatedEvent {
    out := &pbgear.GearVoucherVoucherUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearVoucher_AppendPrograms(in any) *pbgear.GearVoucher_AppendPrograms {
    out := &pbgear.GearVoucher_AppendPrograms{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field AppendPrograms
    v0 := To_OneOf_GearVoucher_AppendPrograms_append_programs(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("AppendPrograms").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_GearVoucher_AppendPrograms_append_programs (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearVoucher_AppendPrograms_AppendProgramsNone{
        AppendProgramsNone: To_GearVoucher_None(param),
        }
    case 1:
        return &pbgear.GearVoucher_AppendPrograms_AppendProgramsSome{
        AppendProgramsSome: To_GearVoucher_Some(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearVoucher_Call(in any) *pbgear.GearVoucher_Call {
    out := &pbgear.GearVoucher_Call{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_GearVoucher_Call_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_GearVoucher_Call_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearVoucher_Call_CallSendMessage{
        CallSendMessage: To_GearVoucher_SendMessage(param),
        }
    case 1:
        return &pbgear.GearVoucher_Call_CallSendReply{
        CallSendReply: To_GearVoucher_SendReply(param),
        }
    case 2:
        return &pbgear.GearVoucher_Call_CallUploadCode{
        CallUploadCode: To_GearVoucher_UploadCode(param),
        }
    case 3:
        return &pbgear.GearVoucher_Call_CallDeclineVoucher{
        CallDeclineVoucher: To_GearVoucher_DeclineVoucher(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearVoucher_CallCall(in any) *pbgear.GearVoucher_CallCall {
    out := &pbgear.GearVoucher_CallCall{}
    v := in.(registry.Valuable)
    _ = v

    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(0))
    // oneOf field Call
    v1 := To_OneOf_GearVoucher_CallCall_call(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_GearVoucher_CallCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_GearVoucher_CallCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearvoucherCallCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearvoucherCallCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_OneOf_GearVoucher_CallCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearVoucher_CallCall_CallSendMessage{
        CallSendMessage: To_GearVoucher_SendMessage(param),
        }
    case 1:
        return &pbgear.GearVoucher_CallCall_CallSendReply{
        CallSendReply: To_GearVoucher_SendReply(param),
        }
    case 2:
        return &pbgear.GearVoucher_CallCall_CallUploadCode{
        CallUploadCode: To_GearVoucher_UploadCode(param),
        }
    case 3:
        return &pbgear.GearVoucher_CallCall_CallDeclineVoucher{
        CallDeclineVoucher: To_GearVoucher_DeclineVoucher(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearVoucher_CallDeprecatedCall(in any) *pbgear.GearVoucher_CallDeprecatedCall {
    out := &pbgear.GearVoucher_CallDeprecatedCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_GearVoucher_CallDeprecatedCall_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_GearVoucher_CallDeprecatedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_GearVoucher_CallDeprecatedCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearvoucherCallDeprecatedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearvoucherCallDeprecatedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_GearVoucher_CallDeprecatedCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearVoucher_CallDeprecatedCall_CallSendMessage{
        CallSendMessage: To_GearVoucher_SendMessage(param),
        }
    case 1:
        return &pbgear.GearVoucher_CallDeprecatedCall_CallSendReply{
        CallSendReply: To_GearVoucher_SendReply(param),
        }
    case 2:
        return &pbgear.GearVoucher_CallDeprecatedCall_CallUploadCode{
        CallUploadCode: To_GearVoucher_UploadCode(param),
        }
    case 3:
        return &pbgear.GearVoucher_CallDeprecatedCall_CallDeclineVoucher{
        CallDeclineVoucher: To_GearVoucher_DeclineVoucher(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearVoucher_DeclineCall(in any) *pbgear.GearVoucher_DeclineCall {
    out := &pbgear.GearVoucher_DeclineCall{}
    v := in.(registry.Valuable)
    _ = v

    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(0))

    return out //from message
}
func To_GearVoucher_DeclineCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_GearVoucher_DeclineCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearvoucherDeclineCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearvoucherDeclineCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_GearVoucher_DeclineVoucher(in any) *pbgear.GearVoucher_DeclineVoucher {
    out := &pbgear.GearVoucher_DeclineVoucher{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearVoucher_IssueCall(in any) *pbgear.GearVoucher_IssueCall {
    out := &pbgear.GearVoucher_IssueCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Spender To_SpCoreCryptoAccountId32(w)
    out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Balance
    out.Balance = To_string(v.ValueAt(1))
    // optional field Programs
    if v.HasValue() {
        out.Programs = To_Optional_GearVoucher_IssueCall_programs(v.ValueAt(2))
    }
    // primitive field CodeUploading
    out.CodeUploading = To_bool(v.ValueAt(3))
    // primitive field Duration
    out.Duration = To_uint32(v.ValueAt(4))

    return out //from message
}
func To_GearVoucher_IssueCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_GearVoucher_IssueCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearvoucherIssueCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearvoucherIssueCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Optional_GearVoucher_IssueCall_programs(in any) *pbgear.BTreeSet {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_GearVoucher_None(in any) *pbgear.GearVoucher_None {
    out := &pbgear.GearVoucher_None{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GearVoucher_PalletGearVoucherInternalVoucherId(in any) *pbgear.GearVoucher_PalletGearVoucherInternalVoucherId {
    out := &pbgear.GearVoucher_PalletGearVoucherInternalVoucherId{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_GearVoucher_RevokeCall(in any) *pbgear.GearVoucher_RevokeCall {
    out := &pbgear.GearVoucher_RevokeCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Spender To_SpCoreCryptoAccountId32(w)
    out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(1))

    return out //from message
}
func To_GearVoucher_RevokeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_GearVoucher_RevokeCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearvoucherRevokeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearvoucherRevokeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_GearVoucher_SendMessage(in any) *pbgear.GearVoucher_SendMessage {
    out := &pbgear.GearVoucher_SendMessage{}
    v := in.(registry.Valuable)
    _ = v

    // field Destination To_GprimitivesActorId(w)
    out.Destination = To_GprimitivesActorId(v.ValueAt(0))
    // primitive field Payload
    out.Payload = To_bytes(v.ValueAt(1))
    // primitive field GasLimit
    out.GasLimit = To_uint64(v.ValueAt(2))
    // primitive field Value
    out.Value = To_string(v.ValueAt(3))
    // primitive field KeepAlive
    out.KeepAlive = To_bool(v.ValueAt(4))

    return out //from message
}
    

func To_GearVoucher_SendReply(in any) *pbgear.GearVoucher_SendReply {
    out := &pbgear.GearVoucher_SendReply{}
    v := in.(registry.Valuable)
    _ = v

    // field ReplyToId To_GprimitivesMessageId(w)
    out.ReplyToId = To_GprimitivesMessageId(v.ValueAt(0))
    // primitive field Payload
    out.Payload = To_bytes(v.ValueAt(1))
    // primitive field GasLimit
    out.GasLimit = To_uint64(v.ValueAt(2))
    // primitive field Value
    out.Value = To_string(v.ValueAt(3))
    // primitive field KeepAlive
    out.KeepAlive = To_bool(v.ValueAt(4))

    return out //from message
}
    

func To_GearVoucher_Some(in any) *pbgear.GearVoucher_Some {
    out := &pbgear.GearVoucher_Some{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_BTreeSet(w)
    out.Value0 = To_BTreeSet(v.ValueAt(0))

    return out //from message
}
    

func To_GearVoucher_UpdateCall(in any) *pbgear.GearVoucher_UpdateCall {
    out := &pbgear.GearVoucher_UpdateCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Spender To_SpCoreCryptoAccountId32(w)
    out.Spender = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // field VoucherId To_GearVoucher_PalletGearVoucherInternalVoucherId(w)
    out.VoucherId = To_GearVoucher_PalletGearVoucherInternalVoucherId(v.ValueAt(1))
    // optional field MoveOwnership
    if v.HasValue() {
        out.MoveOwnership = To_Optional_GearVoucher_UpdateCall_move_ownership(v.ValueAt(2))
    }
    // primitive field BalanceTopUp
    out.BalanceTopUp = To_Optional_string(v.ValueAt(3))
    // oneOf field AppendPrograms
    v4 := To_OneOf_GearVoucher_UpdateCall_append_programs(v.ValueAt(4))
    reflect.ValueOf(out).Elem().FieldByName("AppendPrograms").Set(reflect.ValueOf(v4))
    // primitive field CodeUploading
    out.CodeUploading = To_Optional_bool(v.ValueAt(5))
    // primitive field ProlongDuration
    out.ProlongDuration = To_Optional_uint32(v.ValueAt(6))

    return out //from message
}
func To_GearVoucher_UpdateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_GearVoucher_UpdateCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearvoucherUpdateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearvoucherUpdateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_Optional_GearVoucher_UpdateCall_move_ownership(in any) *pbgear.SpCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_OneOf_GearVoucher_UpdateCall_append_programs (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GearVoucher_UpdateCall_AppendProgramsNone{
        AppendProgramsNone: To_GearVoucher_None(param),
        }
    case 1:
        return &pbgear.GearVoucher_UpdateCall_AppendProgramsSome{
        AppendProgramsSome: To_GearVoucher_Some(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GearVoucher_UploadCode(in any) *pbgear.GearVoucher_UploadCode {
    out := &pbgear.GearVoucher_UploadCode{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Code
    out.Code = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Gear_ClaimValueCall(in any) *pbgear.Gear_ClaimValueCall {
    out := &pbgear.Gear_ClaimValueCall{}
    v := in.(registry.Valuable)
    _ = v

    // field MessageId To_GprimitivesMessageId(w)
    out.MessageId = To_GprimitivesMessageId(v.ValueAt(0))

    return out //from message
}
func To_Gear_ClaimValueCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_ClaimValueCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearClaimValueCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearClaimValueCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Gear_CreateProgramCall(in any) *pbgear.Gear_CreateProgramCall {
    out := &pbgear.Gear_CreateProgramCall{}
    v := in.(registry.Valuable)
    _ = v

    // field CodeId To_GprimitivesCodeId(w)
    out.CodeId = To_GprimitivesCodeId(v.ValueAt(0))
    // primitive field Salt
    out.Salt = To_bytes(v.ValueAt(1))
    // primitive field InitPayload
    out.InitPayload = To_bytes(v.ValueAt(2))
    // primitive field GasLimit
    out.GasLimit = To_uint64(v.ValueAt(3))
    // primitive field Value
    out.Value = To_string(v.ValueAt(4))
    // primitive field KeepAlive
    out.KeepAlive = To_bool(v.ValueAt(5))

    return out //from message
}
func To_Gear_CreateProgramCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_CreateProgramCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearCreateProgramCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearCreateProgramCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Gear_RunCall(in any) *pbgear.Gear_RunCall {
    out := &pbgear.Gear_RunCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field MaxGas
    out.MaxGas = To_Optional_uint64(v.ValueAt(0))

    return out //from message
}
func To_Gear_RunCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_RunCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearRunCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearRunCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Gear_SendMessageCall(in any) *pbgear.Gear_SendMessageCall {
    out := &pbgear.Gear_SendMessageCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Destination To_GprimitivesActorId(w)
    out.Destination = To_GprimitivesActorId(v.ValueAt(0))
    // primitive field Payload
    out.Payload = To_bytes(v.ValueAt(1))
    // primitive field GasLimit
    out.GasLimit = To_uint64(v.ValueAt(2))
    // primitive field Value
    out.Value = To_string(v.ValueAt(3))
    // primitive field KeepAlive
    out.KeepAlive = To_bool(v.ValueAt(4))

    return out //from message
}
func To_Gear_SendMessageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_SendMessageCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearSendMessageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearSendMessageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Gear_SendReplyCall(in any) *pbgear.Gear_SendReplyCall {
    out := &pbgear.Gear_SendReplyCall{}
    v := in.(registry.Valuable)
    _ = v

    // field ReplyToId To_GprimitivesMessageId(w)
    out.ReplyToId = To_GprimitivesMessageId(v.ValueAt(0))
    // primitive field Payload
    out.Payload = To_bytes(v.ValueAt(1))
    // primitive field GasLimit
    out.GasLimit = To_uint64(v.ValueAt(2))
    // primitive field Value
    out.Value = To_string(v.ValueAt(3))
    // primitive field KeepAlive
    out.KeepAlive = To_bool(v.ValueAt(4))

    return out //from message
}
func To_Gear_SendReplyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_SendReplyCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearSendReplyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearSendReplyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Gear_SetExecuteInherentCall(in any) *pbgear.Gear_SetExecuteInherentCall {
    out := &pbgear.Gear_SetExecuteInherentCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value
    out.Value = To_bool(v.ValueAt(0))

    return out //from message
}
func To_Gear_SetExecuteInherentCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_SetExecuteInherentCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearSetExecuteInherentCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearSetExecuteInherentCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Gear_UploadCodeCall(in any) *pbgear.Gear_UploadCodeCall {
    out := &pbgear.Gear_UploadCodeCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Code
    out.Code = To_bytes(v.ValueAt(0))

    return out //from message
}
func To_Gear_UploadCodeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_UploadCodeCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearUploadCodeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearUploadCodeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Gear_UploadProgramCall(in any) *pbgear.Gear_UploadProgramCall {
    out := &pbgear.Gear_UploadProgramCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Code
    out.Code = To_bytes(v.ValueAt(0))
    // primitive field Salt
    out.Salt = To_bytes(v.ValueAt(1))
    // primitive field InitPayload
    out.InitPayload = To_bytes(v.ValueAt(2))
    // primitive field GasLimit
    out.GasLimit = To_uint64(v.ValueAt(3))
    // primitive field Value
    out.Value = To_string(v.ValueAt(4))
    // primitive field KeepAlive
    out.KeepAlive = To_bool(v.ValueAt(5))

    return out //from message
}
func To_Gear_UploadProgramCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Gear_UploadProgramCall(in)
    out := &pbgearextrinsic.RuntimeCall_GearUploadProgramCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GearUploadProgramCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_GeneralAdmin(in any) *pbgear.GeneralAdmin {
    out := &pbgear.GeneralAdmin{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Governance(in any) *pbgear.Governance {
    out := &pbgear.Governance{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GprimitivesActorId(in any) *pbgear.GprimitivesActorId {
    out := &pbgear.GprimitivesActorId{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_GprimitivesCodeId(in any) *pbgear.GprimitivesCodeId {
    out := &pbgear.GprimitivesCodeId{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_GprimitivesMessageId(in any) *pbgear.GprimitivesMessageId {
    out := &pbgear.GprimitivesMessageId{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_GprimitivesReservationId(in any) *pbgear.GprimitivesReservationId {
    out := &pbgear.GprimitivesReservationId{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_GrandpaNewAuthoritiesEvent(in any) *pbgear.GrandpaNewAuthoritiesEvent {
    out := &pbgear.GrandpaNewAuthoritiesEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GrandpaPallet(in any) *pbgear.GrandpaPallet {
    out := &pbgear.GrandpaPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_GrandpaPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_GrandpaPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.GrandpaPallet_CallReportEquivocationCall{
        CallReportEquivocationCall: To_Grandpa_ReportEquivocationCall(param),
        }
    case 1:
        return &pbgear.GrandpaPallet_CallReportEquivocationUnsignedCall{
        CallReportEquivocationUnsignedCall: To_Grandpa_ReportEquivocationUnsignedCall(param),
        }
    case 2:
        return &pbgear.GrandpaPallet_CallNoteStalledCall{
        CallNoteStalledCall: To_Grandpa_NoteStalledCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_GrandpaPausedEvent(in any) *pbgear.GrandpaPausedEvent {
    out := &pbgear.GrandpaPausedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_GrandpaResumedEvent(in any) *pbgear.GrandpaResumedEvent {
    out := &pbgear.GrandpaResumedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Grandpa_Equivocation(in any) *pbgear.Grandpa_Equivocation {
    out := &pbgear.Grandpa_Equivocation{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Equivocation
    v0 := To_OneOf_Grandpa_Equivocation_equivocation(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Equivocation").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Grandpa_Equivocation_equivocation (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Grandpa_Equivocation_EquivocationPrevote{
        EquivocationPrevote: To_Grandpa_Prevote(param),
        }
    case 1:
        return &pbgear.Grandpa_Equivocation_EquivocationPrecommit{
        EquivocationPrecommit: To_Grandpa_Precommit(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Grandpa_FinalityGrandpaEquivocation(in any) *pbgear.Grandpa_FinalityGrandpaEquivocation {
    out := &pbgear.Grandpa_FinalityGrandpaEquivocation{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field RoundNumber
    out.RoundNumber = To_uint64(v.ValueAt(0))
    // field Identity To_Grandpa_SpConsensusGrandpaAppPublic(w)
    out.Identity = To_Grandpa_SpConsensusGrandpaAppPublic(v.ValueAt(1))
    // field First To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(w)
    out.First = To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(v.ValueAt(2))
    // field Second To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(w)
    out.Second = To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(v.ValueAt(3))

    return out //from message
}
    



func To_Grandpa_FinalityGrandpaPrecommit(in any) *pbgear.Grandpa_FinalityGrandpaPrecommit {
    out := &pbgear.Grandpa_FinalityGrandpaPrecommit{}
    v := in.(registry.Valuable)
    _ = v

    // field TargetHash To_PrimitiveTypesH256(w)
    out.TargetHash = To_PrimitiveTypesH256(v.ValueAt(0))
    // primitive field TargetNumber
    out.TargetNumber = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Grandpa_FinalityGrandpaPrevote(in any) *pbgear.Grandpa_FinalityGrandpaPrevote {
    out := &pbgear.Grandpa_FinalityGrandpaPrevote{}
    v := in.(registry.Valuable)
    _ = v

    // field TargetHash To_PrimitiveTypesH256(w)
    out.TargetHash = To_PrimitiveTypesH256(v.ValueAt(0))
    // primitive field TargetNumber
    out.TargetNumber = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Grandpa_GrandpaTrieNodesList(in any) *pbgear.Grandpa_GrandpaTrieNodesList {
    out := &pbgear.Grandpa_GrandpaTrieNodesList{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field TrieNodes
    out.TrieNodes = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Grandpa_NoteStalledCall(in any) *pbgear.Grandpa_NoteStalledCall {
    out := &pbgear.Grandpa_NoteStalledCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Delay
    out.Delay = To_uint32(v.ValueAt(0))
    // primitive field BestFinalizedBlockNumber
    out.BestFinalizedBlockNumber = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_Grandpa_NoteStalledCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Grandpa_NoteStalledCall(in)
    out := &pbgearextrinsic.RuntimeCall_GrandpaNoteStalledCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GrandpaNoteStalledCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Grandpa_Precommit(in any) *pbgear.Grandpa_Precommit {
    out := &pbgear.Grandpa_Precommit{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Grandpa_FinalityGrandpaEquivocation(w)
    out.Value0 = To_Grandpa_FinalityGrandpaEquivocation(v.ValueAt(0))

    return out //from message
}
    

func To_Grandpa_Prevote(in any) *pbgear.Grandpa_Prevote {
    out := &pbgear.Grandpa_Prevote{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Grandpa_FinalityGrandpaEquivocation(w)
    out.Value0 = To_Grandpa_FinalityGrandpaEquivocation(v.ValueAt(0))

    return out //from message
}
    

func To_Grandpa_ReportEquivocationCall(in any) *pbgear.Grandpa_ReportEquivocationCall {
    out := &pbgear.Grandpa_ReportEquivocationCall{}
    v := in.(registry.Valuable)
    _ = v

    // field EquivocationProof To_Grandpa_SpConsensusGrandpaEquivocationProof(w)
    out.EquivocationProof = To_Grandpa_SpConsensusGrandpaEquivocationProof(v.ValueAt(0))
    // field KeyOwnerProof To_SpSessionMembershipProof(w)
    out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

    return out //from message
}
func To_Grandpa_ReportEquivocationCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Grandpa_ReportEquivocationCall(in)
    out := &pbgearextrinsic.RuntimeCall_GrandpaReportEquivocationCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GrandpaReportEquivocationCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_Grandpa_ReportEquivocationUnsignedCall(in any) *pbgear.Grandpa_ReportEquivocationUnsignedCall {
    out := &pbgear.Grandpa_ReportEquivocationUnsignedCall{}
    v := in.(registry.Valuable)
    _ = v

    // field EquivocationProof To_Grandpa_SpConsensusGrandpaEquivocationProof(w)
    out.EquivocationProof = To_Grandpa_SpConsensusGrandpaEquivocationProof(v.ValueAt(0))
    // field KeyOwnerProof To_SpSessionMembershipProof(w)
    out.KeyOwnerProof = To_SpSessionMembershipProof(v.ValueAt(1))

    return out //from message
}
func To_Grandpa_ReportEquivocationUnsignedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Grandpa_ReportEquivocationUnsignedCall(in)
    out := &pbgearextrinsic.RuntimeCall_GrandpaReportEquivocationUnsignedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("GrandpaReportEquivocationUnsignedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_Grandpa_SpConsensusGrandpaAppPublic(in any) *pbgear.Grandpa_SpConsensusGrandpaAppPublic {
    out := &pbgear.Grandpa_SpConsensusGrandpaAppPublic{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreEd25519Public(w)
    out.Value0 = To_SpCoreEd25519Public(v.ValueAt(0))

    return out //from message
}
    

func To_Grandpa_SpConsensusGrandpaAppSignature(in any) *pbgear.Grandpa_SpConsensusGrandpaAppSignature {
    out := &pbgear.Grandpa_SpConsensusGrandpaAppSignature{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreEd25519Signature(w)
    out.Value0 = To_SpCoreEd25519Signature(v.ValueAt(0))

    return out //from message
}
    

func To_Grandpa_SpConsensusGrandpaEquivocationProof(in any) *pbgear.Grandpa_SpConsensusGrandpaEquivocationProof {
    out := &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field SetId
    out.SetId = To_uint64(v.ValueAt(0))
    // oneOf field Equivocation
    v1 := To_OneOf_Grandpa_SpConsensusGrandpaEquivocationProof_equivocation(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Equivocation").Set(reflect.ValueOf(v1))

    return out //from message
}
    
func To_OneOf_Grandpa_SpConsensusGrandpaEquivocationProof_equivocation (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof_EquivocationPrevote{
        EquivocationPrevote: To_Grandpa_Prevote(param),
        }
    case 1:
        return &pbgear.Grandpa_SpConsensusGrandpaEquivocationProof_EquivocationPrecommit{
        EquivocationPrecommit: To_Grandpa_Precommit(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature(in any) *pbgear.Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature {
    out := &pbgear.Grandpa_TupleFinalityGrandpaPrecommitspConsensusGrandpaAppSignature{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Grandpa_FinalityGrandpaPrecommit(w)
    out.Value0 = To_Grandpa_FinalityGrandpaPrecommit(v.ValueAt(0))
    // field Value1 To_Grandpa_SpConsensusGrandpaAppSignature(w)
    out.Value1 = To_Grandpa_SpConsensusGrandpaAppSignature(v.ValueAt(1))

    return out //from message
}
    


func To_Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature(in any) *pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature {
    out := &pbgear.Grandpa_TupleFinalityGrandpaPrevotespConsensusGrandpaAppSignature{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Grandpa_FinalityGrandpaPrevote(w)
    out.Value0 = To_Grandpa_FinalityGrandpaPrevote(v.ValueAt(0))
    // field Value1 To_Grandpa_SpConsensusGrandpaAppSignature(w)
    out.Value1 = To_Grandpa_SpConsensusGrandpaAppSignature(v.ValueAt(1))

    return out //from message
}
    


func To_Handle(in any) *pbgear.Handle {
    out := &pbgear.Handle{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_HistoricalPallet(in any) *pbgear.HistoricalPallet {
    out := &pbgear.HistoricalPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityIdentityClearedEvent(in any) *pbgear.IdentityIdentityClearedEvent {
    out := &pbgear.IdentityIdentityClearedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityIdentityKilledEvent(in any) *pbgear.IdentityIdentityKilledEvent {
    out := &pbgear.IdentityIdentityKilledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityIdentitySetEvent(in any) *pbgear.IdentityIdentitySetEvent {
    out := &pbgear.IdentityIdentitySetEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityJudgement(in any) *pbgear.IdentityJudgement {
    out := &pbgear.IdentityJudgement{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityJudgementGivenEvent(in any) *pbgear.IdentityJudgementGivenEvent {
    out := &pbgear.IdentityJudgementGivenEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityJudgementRequestedEvent(in any) *pbgear.IdentityJudgementRequestedEvent {
    out := &pbgear.IdentityJudgementRequestedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityJudgementUnrequestedEvent(in any) *pbgear.IdentityJudgementUnrequestedEvent {
    out := &pbgear.IdentityJudgementUnrequestedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentityPallet(in any) *pbgear.IdentityPallet {
    out := &pbgear.IdentityPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_IdentityPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_IdentityPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.IdentityPallet_CallAddRegistrarCall{
        CallAddRegistrarCall: To_Identity_AddRegistrarCall(param),
        }
    case 1:
        return &pbgear.IdentityPallet_CallSetIdentityCall{
        CallSetIdentityCall: To_Identity_SetIdentityCall(param),
        }
    case 2:
        return &pbgear.IdentityPallet_CallSetSubsCall{
        CallSetSubsCall: To_Identity_SetSubsCall(param),
        }
    case 3:
        return &pbgear.IdentityPallet_CallClearIdentityCall{
        CallClearIdentityCall: To_Identity_ClearIdentityCall(param),
        }
    case 4:
        return &pbgear.IdentityPallet_CallRequestJudgementCall{
        CallRequestJudgementCall: To_Identity_RequestJudgementCall(param),
        }
    case 5:
        return &pbgear.IdentityPallet_CallCancelRequestCall{
        CallCancelRequestCall: To_Identity_CancelRequestCall(param),
        }
    case 6:
        return &pbgear.IdentityPallet_CallSetFeeCall{
        CallSetFeeCall: To_Identity_SetFeeCall(param),
        }
    case 7:
        return &pbgear.IdentityPallet_CallSetAccountIdCall{
        CallSetAccountIdCall: To_Identity_SetAccountIdCall(param),
        }
    case 8:
        return &pbgear.IdentityPallet_CallSetFieldsCall{
        CallSetFieldsCall: To_Identity_SetFieldsCall(param),
        }
    case 9:
        return &pbgear.IdentityPallet_CallProvideJudgementCall{
        CallProvideJudgementCall: To_Identity_ProvideJudgementCall(param),
        }
    case 10:
        return &pbgear.IdentityPallet_CallKillIdentityCall{
        CallKillIdentityCall: To_Identity_KillIdentityCall(param),
        }
    case 11:
        return &pbgear.IdentityPallet_CallAddSubCall{
        CallAddSubCall: To_Identity_AddSubCall(param),
        }
    case 12:
        return &pbgear.IdentityPallet_CallRenameSubCall{
        CallRenameSubCall: To_Identity_RenameSubCall(param),
        }
    case 13:
        return &pbgear.IdentityPallet_CallRemoveSubCall{
        CallRemoveSubCall: To_Identity_RemoveSubCall(param),
        }
    case 14:
        return &pbgear.IdentityPallet_CallQuitSubCall{
        CallQuitSubCall: To_Identity_QuitSubCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_IdentityRegistrarAddedEvent(in any) *pbgear.IdentityRegistrarAddedEvent {
    out := &pbgear.IdentityRegistrarAddedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentitySubIdentityAddedEvent(in any) *pbgear.IdentitySubIdentityAddedEvent {
    out := &pbgear.IdentitySubIdentityAddedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentitySubIdentityRemovedEvent(in any) *pbgear.IdentitySubIdentityRemovedEvent {
    out := &pbgear.IdentitySubIdentityRemovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_IdentitySubIdentityRevokedEvent(in any) *pbgear.IdentitySubIdentityRevokedEvent {
    out := &pbgear.IdentitySubIdentityRevokedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_Account(in any) *pbgear.Identity_Account {
    out := &pbgear.Identity_Account{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Account
    v0 := To_OneOf_Identity_Account_account(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Account").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Account_account (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Account_AccountId{
        AccountId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_Account_AccountIndex{
        AccountIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_Account_AccountRaw{
        AccountRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_Account_AccountAddress32{
        AccountAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_Account_AccountAddress20{
        AccountAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_AddRegistrarCall(in any) *pbgear.Identity_AddRegistrarCall {
    out := &pbgear.Identity_AddRegistrarCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Account
    v0 := To_OneOf_Identity_AddRegistrarCall_account(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Account").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_Identity_AddRegistrarCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_AddRegistrarCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityAddRegistrarCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityAddRegistrarCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Identity_AddRegistrarCall_account (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_AddRegistrarCall_AccountId{
        AccountId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_AddRegistrarCall_AccountIndex{
        AccountIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_AddRegistrarCall_AccountRaw{
        AccountRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_AddRegistrarCall_AccountAddress32{
        AccountAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_AddRegistrarCall_AccountAddress20{
        AccountAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_AddSubCall(in any) *pbgear.Identity_AddSubCall {
    out := &pbgear.Identity_AddSubCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Sub
    v0 := To_OneOf_Identity_AddSubCall_sub(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))
    // oneOf field Data
    v1 := To_OneOf_Identity_AddSubCall_data(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Data").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Identity_AddSubCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_AddSubCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityAddSubCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityAddSubCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Identity_AddSubCall_sub (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_AddSubCall_SubId{
        SubId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_AddSubCall_SubIndex{
        SubIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_AddSubCall_SubRaw{
        SubRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_AddSubCall_SubAddress32{
        SubAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_AddSubCall_SubAddress20{
        SubAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_AddSubCall_data (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_AddSubCall_DataNone{
        DataNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_AddSubCall_DataRaw0{
        DataRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_AddSubCall_DataRaw1{
        DataRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_AddSubCall_DataRaw2{
        DataRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_AddSubCall_DataRaw3{
        DataRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_AddSubCall_DataRaw4{
        DataRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_AddSubCall_DataRaw5{
        DataRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_AddSubCall_DataRaw6{
        DataRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_AddSubCall_DataRaw7{
        DataRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_AddSubCall_DataRaw8{
        DataRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_AddSubCall_DataRaw9{
        DataRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_AddSubCall_DataRaw10{
        DataRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_AddSubCall_DataRaw11{
        DataRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_AddSubCall_DataRaw12{
        DataRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_AddSubCall_DataRaw13{
        DataRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_AddSubCall_DataRaw14{
        DataRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_AddSubCall_DataRaw15{
        DataRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_AddSubCall_DataRaw16{
        DataRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_AddSubCall_DataRaw17{
        DataRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_AddSubCall_DataRaw18{
        DataRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_AddSubCall_DataRaw19{
        DataRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_AddSubCall_DataRaw20{
        DataRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_AddSubCall_DataRaw21{
        DataRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_AddSubCall_DataRaw22{
        DataRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_AddSubCall_DataRaw23{
        DataRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_AddSubCall_DataRaw24{
        DataRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_AddSubCall_DataRaw25{
        DataRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_AddSubCall_DataRaw26{
        DataRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_AddSubCall_DataRaw27{
        DataRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_AddSubCall_DataRaw28{
        DataRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_AddSubCall_DataRaw29{
        DataRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_AddSubCall_DataRaw30{
        DataRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_AddSubCall_DataRaw31{
        DataRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_AddSubCall_DataRaw32{
        DataRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_AddSubCall_DataBlakeTwo256{
        DataBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_AddSubCall_DataSha256{
        DataSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_AddSubCall_DataKeccak256{
        DataKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_AddSubCall_DataShaThree256{
        DataShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Address20(in any) *pbgear.Identity_Address20 {
    out := &pbgear.Identity_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Address32(in any) *pbgear.Identity_Address32 {
    out := &pbgear.Identity_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_BlakeTwo256(in any) *pbgear.Identity_BlakeTwo256 {
    out := &pbgear.Identity_BlakeTwo256{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_CancelRequestCall(in any) *pbgear.Identity_CancelRequestCall {
    out := &pbgear.Identity_CancelRequestCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field RegIndex
    out.RegIndex = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Identity_CancelRequestCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_CancelRequestCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityCancelRequestCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityCancelRequestCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Identity_ClearIdentityCall(in any) *pbgear.Identity_ClearIdentityCall {
    out := &pbgear.Identity_ClearIdentityCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_Data(in any) *pbgear.Identity_Data {
    out := &pbgear.Identity_Data{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Data
    v0 := To_OneOf_Identity_Data_data(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Data").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Data_data (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Data_DataNone{
        DataNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Data_DataRaw0{
        DataRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Data_DataRaw1{
        DataRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Data_DataRaw2{
        DataRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Data_DataRaw3{
        DataRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Data_DataRaw4{
        DataRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Data_DataRaw5{
        DataRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Data_DataRaw6{
        DataRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Data_DataRaw7{
        DataRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Data_DataRaw8{
        DataRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Data_DataRaw9{
        DataRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Data_DataRaw10{
        DataRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Data_DataRaw11{
        DataRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Data_DataRaw12{
        DataRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Data_DataRaw13{
        DataRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Data_DataRaw14{
        DataRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Data_DataRaw15{
        DataRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Data_DataRaw16{
        DataRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Data_DataRaw17{
        DataRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Data_DataRaw18{
        DataRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Data_DataRaw19{
        DataRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Data_DataRaw20{
        DataRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Data_DataRaw21{
        DataRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Data_DataRaw22{
        DataRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Data_DataRaw23{
        DataRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Data_DataRaw24{
        DataRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Data_DataRaw25{
        DataRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Data_DataRaw26{
        DataRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Data_DataRaw27{
        DataRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Data_DataRaw28{
        DataRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Data_DataRaw29{
        DataRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Data_DataRaw30{
        DataRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Data_DataRaw31{
        DataRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Data_DataRaw32{
        DataRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Data_DataBlakeTwo256{
        DataBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Data_DataSha256{
        DataSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Data_DataKeccak256{
        DataKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Data_DataShaThree256{
        DataShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Display(in any) *pbgear.Identity_Display {
    out := &pbgear.Identity_Display{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Display
    v0 := To_OneOf_Identity_Display_display(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Display").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Display_display (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Display_DisplayNone{
        DisplayNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Display_DisplayRaw0{
        DisplayRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Display_DisplayRaw1{
        DisplayRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Display_DisplayRaw2{
        DisplayRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Display_DisplayRaw3{
        DisplayRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Display_DisplayRaw4{
        DisplayRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Display_DisplayRaw5{
        DisplayRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Display_DisplayRaw6{
        DisplayRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Display_DisplayRaw7{
        DisplayRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Display_DisplayRaw8{
        DisplayRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Display_DisplayRaw9{
        DisplayRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Display_DisplayRaw10{
        DisplayRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Display_DisplayRaw11{
        DisplayRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Display_DisplayRaw12{
        DisplayRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Display_DisplayRaw13{
        DisplayRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Display_DisplayRaw14{
        DisplayRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Display_DisplayRaw15{
        DisplayRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Display_DisplayRaw16{
        DisplayRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Display_DisplayRaw17{
        DisplayRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Display_DisplayRaw18{
        DisplayRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Display_DisplayRaw19{
        DisplayRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Display_DisplayRaw20{
        DisplayRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Display_DisplayRaw21{
        DisplayRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Display_DisplayRaw22{
        DisplayRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Display_DisplayRaw23{
        DisplayRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Display_DisplayRaw24{
        DisplayRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Display_DisplayRaw25{
        DisplayRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Display_DisplayRaw26{
        DisplayRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Display_DisplayRaw27{
        DisplayRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Display_DisplayRaw28{
        DisplayRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Display_DisplayRaw29{
        DisplayRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Display_DisplayRaw30{
        DisplayRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Display_DisplayRaw31{
        DisplayRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Display_DisplayRaw32{
        DisplayRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Display_DisplayBlakeTwo256{
        DisplayBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Display_DisplaySha256{
        DisplaySha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Display_DisplayKeccak256{
        DisplayKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Display_DisplayShaThree256{
        DisplayShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Email(in any) *pbgear.Identity_Email {
    out := &pbgear.Identity_Email{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Email
    v0 := To_OneOf_Identity_Email_email(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Email").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Email_email (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Email_EmailNone{
        EmailNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Email_EmailRaw0{
        EmailRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Email_EmailRaw1{
        EmailRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Email_EmailRaw2{
        EmailRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Email_EmailRaw3{
        EmailRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Email_EmailRaw4{
        EmailRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Email_EmailRaw5{
        EmailRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Email_EmailRaw6{
        EmailRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Email_EmailRaw7{
        EmailRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Email_EmailRaw8{
        EmailRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Email_EmailRaw9{
        EmailRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Email_EmailRaw10{
        EmailRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Email_EmailRaw11{
        EmailRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Email_EmailRaw12{
        EmailRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Email_EmailRaw13{
        EmailRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Email_EmailRaw14{
        EmailRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Email_EmailRaw15{
        EmailRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Email_EmailRaw16{
        EmailRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Email_EmailRaw17{
        EmailRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Email_EmailRaw18{
        EmailRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Email_EmailRaw19{
        EmailRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Email_EmailRaw20{
        EmailRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Email_EmailRaw21{
        EmailRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Email_EmailRaw22{
        EmailRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Email_EmailRaw23{
        EmailRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Email_EmailRaw24{
        EmailRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Email_EmailRaw25{
        EmailRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Email_EmailRaw26{
        EmailRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Email_EmailRaw27{
        EmailRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Email_EmailRaw28{
        EmailRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Email_EmailRaw29{
        EmailRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Email_EmailRaw30{
        EmailRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Email_EmailRaw31{
        EmailRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Email_EmailRaw32{
        EmailRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Email_EmailBlakeTwo256{
        EmailBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Email_EmailSha256{
        EmailSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Email_EmailKeccak256{
        EmailKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Email_EmailShaThree256{
        EmailShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Erroneous(in any) *pbgear.Identity_Erroneous {
    out := &pbgear.Identity_Erroneous{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_FeePaid(in any) *pbgear.Identity_FeePaid {
    out := &pbgear.Identity_FeePaid{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Id(in any) *pbgear.Identity_Id {
    out := &pbgear.Identity_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Identity_Image(in any) *pbgear.Identity_Image {
    out := &pbgear.Identity_Image{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Image
    v0 := To_OneOf_Identity_Image_image(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Image").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Image_image (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Image_ImageNone{
        ImageNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Image_ImageRaw0{
        ImageRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Image_ImageRaw1{
        ImageRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Image_ImageRaw2{
        ImageRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Image_ImageRaw3{
        ImageRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Image_ImageRaw4{
        ImageRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Image_ImageRaw5{
        ImageRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Image_ImageRaw6{
        ImageRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Image_ImageRaw7{
        ImageRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Image_ImageRaw8{
        ImageRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Image_ImageRaw9{
        ImageRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Image_ImageRaw10{
        ImageRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Image_ImageRaw11{
        ImageRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Image_ImageRaw12{
        ImageRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Image_ImageRaw13{
        ImageRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Image_ImageRaw14{
        ImageRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Image_ImageRaw15{
        ImageRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Image_ImageRaw16{
        ImageRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Image_ImageRaw17{
        ImageRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Image_ImageRaw18{
        ImageRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Image_ImageRaw19{
        ImageRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Image_ImageRaw20{
        ImageRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Image_ImageRaw21{
        ImageRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Image_ImageRaw22{
        ImageRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Image_ImageRaw23{
        ImageRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Image_ImageRaw24{
        ImageRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Image_ImageRaw25{
        ImageRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Image_ImageRaw26{
        ImageRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Image_ImageRaw27{
        ImageRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Image_ImageRaw28{
        ImageRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Image_ImageRaw29{
        ImageRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Image_ImageRaw30{
        ImageRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Image_ImageRaw31{
        ImageRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Image_ImageRaw32{
        ImageRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Image_ImageBlakeTwo256{
        ImageBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Image_ImageSha256{
        ImageSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Image_ImageKeccak256{
        ImageKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Image_ImageShaThree256{
        ImageShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Index(in any) *pbgear.Identity_Index {
    out := &pbgear.Identity_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Identity_TupleNull(w)
    out.Value0 = To_Identity_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Identity_Judgement(in any) *pbgear.Identity_Judgement {
    out := &pbgear.Identity_Judgement{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Judgement
    v0 := To_OneOf_Identity_Judgement_judgement(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Judgement").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Judgement_judgement (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Judgement_JudgementUnknown{
        JudgementUnknown: To_Identity_Unknown(param),
        }
    case 1:
        return &pbgear.Identity_Judgement_JudgementFeePaid{
        JudgementFeePaid: To_Identity_FeePaid(param),
        }
    case 2:
        return &pbgear.Identity_Judgement_JudgementReasonable{
        JudgementReasonable: To_Identity_Reasonable(param),
        }
    case 3:
        return &pbgear.Identity_Judgement_JudgementKnownGood{
        JudgementKnownGood: To_Identity_KnownGood(param),
        }
    case 4:
        return &pbgear.Identity_Judgement_JudgementOutOfDate{
        JudgementOutOfDate: To_Identity_OutOfDate(param),
        }
    case 5:
        return &pbgear.Identity_Judgement_JudgementLowQuality{
        JudgementLowQuality: To_Identity_LowQuality(param),
        }
    case 6:
        return &pbgear.Identity_Judgement_JudgementErroneous{
        JudgementErroneous: To_Identity_Erroneous(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Keccak256(in any) *pbgear.Identity_Keccak256 {
    out := &pbgear.Identity_Keccak256{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_KillIdentityCall(in any) *pbgear.Identity_KillIdentityCall {
    out := &pbgear.Identity_KillIdentityCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Target
    v0 := To_OneOf_Identity_KillIdentityCall_target(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_Identity_KillIdentityCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_KillIdentityCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityKillIdentityCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityKillIdentityCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Identity_KillIdentityCall_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_KillIdentityCall_TargetId{
        TargetId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_KillIdentityCall_TargetIndex{
        TargetIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_KillIdentityCall_TargetRaw{
        TargetRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_KillIdentityCall_TargetAddress32{
        TargetAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_KillIdentityCall_TargetAddress20{
        TargetAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_KnownGood(in any) *pbgear.Identity_KnownGood {
    out := &pbgear.Identity_KnownGood{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_Legal(in any) *pbgear.Identity_Legal {
    out := &pbgear.Identity_Legal{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Legal
    v0 := To_OneOf_Identity_Legal_legal(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Legal").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Legal_legal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Legal_LegalNone{
        LegalNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Legal_LegalRaw0{
        LegalRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Legal_LegalRaw1{
        LegalRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Legal_LegalRaw2{
        LegalRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Legal_LegalRaw3{
        LegalRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Legal_LegalRaw4{
        LegalRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Legal_LegalRaw5{
        LegalRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Legal_LegalRaw6{
        LegalRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Legal_LegalRaw7{
        LegalRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Legal_LegalRaw8{
        LegalRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Legal_LegalRaw9{
        LegalRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Legal_LegalRaw10{
        LegalRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Legal_LegalRaw11{
        LegalRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Legal_LegalRaw12{
        LegalRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Legal_LegalRaw13{
        LegalRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Legal_LegalRaw14{
        LegalRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Legal_LegalRaw15{
        LegalRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Legal_LegalRaw16{
        LegalRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Legal_LegalRaw17{
        LegalRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Legal_LegalRaw18{
        LegalRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Legal_LegalRaw19{
        LegalRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Legal_LegalRaw20{
        LegalRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Legal_LegalRaw21{
        LegalRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Legal_LegalRaw22{
        LegalRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Legal_LegalRaw23{
        LegalRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Legal_LegalRaw24{
        LegalRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Legal_LegalRaw25{
        LegalRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Legal_LegalRaw26{
        LegalRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Legal_LegalRaw27{
        LegalRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Legal_LegalRaw28{
        LegalRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Legal_LegalRaw29{
        LegalRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Legal_LegalRaw30{
        LegalRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Legal_LegalRaw31{
        LegalRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Legal_LegalRaw32{
        LegalRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Legal_LegalBlakeTwo256{
        LegalBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Legal_LegalSha256{
        LegalSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Legal_LegalKeccak256{
        LegalKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Legal_LegalShaThree256{
        LegalShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_LowQuality(in any) *pbgear.Identity_LowQuality {
    out := &pbgear.Identity_LowQuality{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_New(in any) *pbgear.Identity_New {
    out := &pbgear.Identity_New{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field New
    v0 := To_OneOf_Identity_New_new(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("New").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_New_new (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_New_NewId{
        NewId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_New_NewIndex{
        NewIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_New_NewRaw{
        NewRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_New_NewAddress32{
        NewAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_New_NewAddress20{
        NewAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_None(in any) *pbgear.Identity_None {
    out := &pbgear.Identity_None{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_OutOfDate(in any) *pbgear.Identity_OutOfDate {
    out := &pbgear.Identity_OutOfDate{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_PalletIdentitySimpleIdentityInfo(in any) *pbgear.Identity_PalletIdentitySimpleIdentityInfo {
    out := &pbgear.Identity_PalletIdentitySimpleIdentityInfo{}
    v := in.(registry.Valuable)
    _ = v

    // field Additional To_BoundedCollectionsBoundedVecBoundedVec(w)
    out.Additional = To_BoundedCollectionsBoundedVecBoundedVec(v.ValueAt(0))
    // oneOf field Display
    v1 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_display(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Display").Set(reflect.ValueOf(v1))
    // oneOf field Legal
    v2 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_legal(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Legal").Set(reflect.ValueOf(v2))
    // oneOf field Web
    v3 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_web(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Web").Set(reflect.ValueOf(v3))
    // oneOf field Riot
    v4 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_riot(v.ValueAt(4))
    reflect.ValueOf(out).Elem().FieldByName("Riot").Set(reflect.ValueOf(v4))
    // oneOf field Email
    v5 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_email(v.ValueAt(5))
    reflect.ValueOf(out).Elem().FieldByName("Email").Set(reflect.ValueOf(v5))
    // primitive field PgpFingerprint
    out.PgpFingerprint = To_bytes(v.ValueAt(6))
    // oneOf field Image
    v7 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_image(v.ValueAt(7))
    reflect.ValueOf(out).Elem().FieldByName("Image").Set(reflect.ValueOf(v7))
    // oneOf field Twitter
    v8 := To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_twitter(v.ValueAt(8))
    reflect.ValueOf(out).Elem().FieldByName("Twitter").Set(reflect.ValueOf(v8))

    return out //from message
}
    

func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_display (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayNone{
        DisplayNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw0{
        DisplayRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw1{
        DisplayRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw2{
        DisplayRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw3{
        DisplayRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw4{
        DisplayRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw5{
        DisplayRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw6{
        DisplayRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw7{
        DisplayRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw8{
        DisplayRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw9{
        DisplayRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw10{
        DisplayRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw11{
        DisplayRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw12{
        DisplayRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw13{
        DisplayRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw14{
        DisplayRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw15{
        DisplayRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw16{
        DisplayRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw17{
        DisplayRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw18{
        DisplayRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw19{
        DisplayRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw20{
        DisplayRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw21{
        DisplayRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw22{
        DisplayRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw23{
        DisplayRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw24{
        DisplayRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw25{
        DisplayRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw26{
        DisplayRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw27{
        DisplayRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw28{
        DisplayRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw29{
        DisplayRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw30{
        DisplayRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw31{
        DisplayRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayRaw32{
        DisplayRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayBlakeTwo256{
        DisplayBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplaySha256{
        DisplaySha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayKeccak256{
        DisplayKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_DisplayShaThree256{
        DisplayShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_legal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalNone{
        LegalNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw0{
        LegalRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw1{
        LegalRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw2{
        LegalRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw3{
        LegalRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw4{
        LegalRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw5{
        LegalRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw6{
        LegalRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw7{
        LegalRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw8{
        LegalRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw9{
        LegalRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw10{
        LegalRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw11{
        LegalRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw12{
        LegalRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw13{
        LegalRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw14{
        LegalRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw15{
        LegalRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw16{
        LegalRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw17{
        LegalRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw18{
        LegalRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw19{
        LegalRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw20{
        LegalRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw21{
        LegalRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw22{
        LegalRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw23{
        LegalRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw24{
        LegalRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw25{
        LegalRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw26{
        LegalRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw27{
        LegalRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw28{
        LegalRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw29{
        LegalRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw30{
        LegalRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw31{
        LegalRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalRaw32{
        LegalRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalBlakeTwo256{
        LegalBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalSha256{
        LegalSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalKeccak256{
        LegalKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_LegalShaThree256{
        LegalShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_web (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebNone{
        WebNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw0{
        WebRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw1{
        WebRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw2{
        WebRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw3{
        WebRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw4{
        WebRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw5{
        WebRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw6{
        WebRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw7{
        WebRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw8{
        WebRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw9{
        WebRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw10{
        WebRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw11{
        WebRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw12{
        WebRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw13{
        WebRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw14{
        WebRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw15{
        WebRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw16{
        WebRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw17{
        WebRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw18{
        WebRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw19{
        WebRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw20{
        WebRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw21{
        WebRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw22{
        WebRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw23{
        WebRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw24{
        WebRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw25{
        WebRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw26{
        WebRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw27{
        WebRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw28{
        WebRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw29{
        WebRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw30{
        WebRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw31{
        WebRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebRaw32{
        WebRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebBlakeTwo256{
        WebBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebSha256{
        WebSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebKeccak256{
        WebKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_WebShaThree256{
        WebShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_riot (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotNone{
        RiotNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw0{
        RiotRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw1{
        RiotRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw2{
        RiotRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw3{
        RiotRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw4{
        RiotRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw5{
        RiotRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw6{
        RiotRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw7{
        RiotRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw8{
        RiotRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw9{
        RiotRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw10{
        RiotRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw11{
        RiotRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw12{
        RiotRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw13{
        RiotRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw14{
        RiotRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw15{
        RiotRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw16{
        RiotRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw17{
        RiotRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw18{
        RiotRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw19{
        RiotRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw20{
        RiotRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw21{
        RiotRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw22{
        RiotRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw23{
        RiotRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw24{
        RiotRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw25{
        RiotRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw26{
        RiotRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw27{
        RiotRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw28{
        RiotRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw29{
        RiotRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw30{
        RiotRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw31{
        RiotRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotRaw32{
        RiotRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotBlakeTwo256{
        RiotBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotSha256{
        RiotSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotKeccak256{
        RiotKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_RiotShaThree256{
        RiotShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_email (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailNone{
        EmailNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw0{
        EmailRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw1{
        EmailRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw2{
        EmailRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw3{
        EmailRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw4{
        EmailRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw5{
        EmailRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw6{
        EmailRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw7{
        EmailRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw8{
        EmailRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw9{
        EmailRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw10{
        EmailRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw11{
        EmailRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw12{
        EmailRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw13{
        EmailRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw14{
        EmailRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw15{
        EmailRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw16{
        EmailRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw17{
        EmailRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw18{
        EmailRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw19{
        EmailRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw20{
        EmailRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw21{
        EmailRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw22{
        EmailRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw23{
        EmailRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw24{
        EmailRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw25{
        EmailRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw26{
        EmailRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw27{
        EmailRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw28{
        EmailRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw29{
        EmailRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw30{
        EmailRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw31{
        EmailRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailRaw32{
        EmailRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailBlakeTwo256{
        EmailBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailSha256{
        EmailSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailKeccak256{
        EmailKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_EmailShaThree256{
        EmailShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_image (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageNone{
        ImageNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw0{
        ImageRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw1{
        ImageRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw2{
        ImageRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw3{
        ImageRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw4{
        ImageRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw5{
        ImageRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw6{
        ImageRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw7{
        ImageRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw8{
        ImageRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw9{
        ImageRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw10{
        ImageRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw11{
        ImageRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw12{
        ImageRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw13{
        ImageRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw14{
        ImageRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw15{
        ImageRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw16{
        ImageRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw17{
        ImageRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw18{
        ImageRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw19{
        ImageRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw20{
        ImageRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw21{
        ImageRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw22{
        ImageRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw23{
        ImageRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw24{
        ImageRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw25{
        ImageRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw26{
        ImageRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw27{
        ImageRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw28{
        ImageRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw29{
        ImageRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw30{
        ImageRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw31{
        ImageRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageRaw32{
        ImageRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageBlakeTwo256{
        ImageBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageSha256{
        ImageSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageKeccak256{
        ImageKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_ImageShaThree256{
        ImageShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_PalletIdentitySimpleIdentityInfo_twitter (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterNone{
        TwitterNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw0{
        TwitterRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw1{
        TwitterRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw2{
        TwitterRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw3{
        TwitterRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw4{
        TwitterRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw5{
        TwitterRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw6{
        TwitterRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw7{
        TwitterRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw8{
        TwitterRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw9{
        TwitterRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw10{
        TwitterRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw11{
        TwitterRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw12{
        TwitterRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw13{
        TwitterRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw14{
        TwitterRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw15{
        TwitterRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw16{
        TwitterRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw17{
        TwitterRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw18{
        TwitterRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw19{
        TwitterRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw20{
        TwitterRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw21{
        TwitterRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw22{
        TwitterRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw23{
        TwitterRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw24{
        TwitterRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw25{
        TwitterRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw26{
        TwitterRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw27{
        TwitterRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw28{
        TwitterRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw29{
        TwitterRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw30{
        TwitterRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw31{
        TwitterRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterRaw32{
        TwitterRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterBlakeTwo256{
        TwitterBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterSha256{
        TwitterSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterKeccak256{
        TwitterKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_PalletIdentitySimpleIdentityInfo_TwitterShaThree256{
        TwitterShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_PalletIdentityTypesBitFlags(in any) *pbgear.Identity_PalletIdentityTypesBitFlags {
    out := &pbgear.Identity_PalletIdentityTypesBitFlags{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint64(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_ProvideJudgementCall(in any) *pbgear.Identity_ProvideJudgementCall {
    out := &pbgear.Identity_ProvideJudgementCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field RegIndex
    out.RegIndex = To_uint32(v.ValueAt(0))
    // oneOf field Target
    v1 := To_OneOf_Identity_ProvideJudgementCall_target(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v1))
    // oneOf field Judgement
    v2 := To_OneOf_Identity_ProvideJudgementCall_judgement(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Judgement").Set(reflect.ValueOf(v2))
    // field Identity To_PrimitiveTypesH256(w)
    out.Identity = To_PrimitiveTypesH256(v.ValueAt(3))

    return out //from message
}
func To_Identity_ProvideJudgementCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_ProvideJudgementCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityProvideJudgementCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityProvideJudgementCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Identity_ProvideJudgementCall_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_ProvideJudgementCall_TargetId{
        TargetId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_ProvideJudgementCall_TargetIndex{
        TargetIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_ProvideJudgementCall_TargetRaw{
        TargetRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_ProvideJudgementCall_TargetAddress32{
        TargetAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_ProvideJudgementCall_TargetAddress20{
        TargetAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_ProvideJudgementCall_judgement (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_ProvideJudgementCall_JudgementUnknown{
        JudgementUnknown: To_Identity_Unknown(param),
        }
    case 1:
        return &pbgear.Identity_ProvideJudgementCall_JudgementFeePaid{
        JudgementFeePaid: To_Identity_FeePaid(param),
        }
    case 2:
        return &pbgear.Identity_ProvideJudgementCall_JudgementReasonable{
        JudgementReasonable: To_Identity_Reasonable(param),
        }
    case 3:
        return &pbgear.Identity_ProvideJudgementCall_JudgementKnownGood{
        JudgementKnownGood: To_Identity_KnownGood(param),
        }
    case 4:
        return &pbgear.Identity_ProvideJudgementCall_JudgementOutOfDate{
        JudgementOutOfDate: To_Identity_OutOfDate(param),
        }
    case 5:
        return &pbgear.Identity_ProvideJudgementCall_JudgementLowQuality{
        JudgementLowQuality: To_Identity_LowQuality(param),
        }
    case 6:
        return &pbgear.Identity_ProvideJudgementCall_JudgementErroneous{
        JudgementErroneous: To_Identity_Erroneous(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Identity_QuitSubCall(in any) *pbgear.Identity_QuitSubCall {
    out := &pbgear.Identity_QuitSubCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_Raw(in any) *pbgear.Identity_Raw {
    out := &pbgear.Identity_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw0(in any) *pbgear.Identity_Raw0 {
    out := &pbgear.Identity_Raw0{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw1(in any) *pbgear.Identity_Raw1 {
    out := &pbgear.Identity_Raw1{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw10(in any) *pbgear.Identity_Raw10 {
    out := &pbgear.Identity_Raw10{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw11(in any) *pbgear.Identity_Raw11 {
    out := &pbgear.Identity_Raw11{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw12(in any) *pbgear.Identity_Raw12 {
    out := &pbgear.Identity_Raw12{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw13(in any) *pbgear.Identity_Raw13 {
    out := &pbgear.Identity_Raw13{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw14(in any) *pbgear.Identity_Raw14 {
    out := &pbgear.Identity_Raw14{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw15(in any) *pbgear.Identity_Raw15 {
    out := &pbgear.Identity_Raw15{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw16(in any) *pbgear.Identity_Raw16 {
    out := &pbgear.Identity_Raw16{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw17(in any) *pbgear.Identity_Raw17 {
    out := &pbgear.Identity_Raw17{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw18(in any) *pbgear.Identity_Raw18 {
    out := &pbgear.Identity_Raw18{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw19(in any) *pbgear.Identity_Raw19 {
    out := &pbgear.Identity_Raw19{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw2(in any) *pbgear.Identity_Raw2 {
    out := &pbgear.Identity_Raw2{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw20(in any) *pbgear.Identity_Raw20 {
    out := &pbgear.Identity_Raw20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw21(in any) *pbgear.Identity_Raw21 {
    out := &pbgear.Identity_Raw21{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw22(in any) *pbgear.Identity_Raw22 {
    out := &pbgear.Identity_Raw22{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw23(in any) *pbgear.Identity_Raw23 {
    out := &pbgear.Identity_Raw23{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw24(in any) *pbgear.Identity_Raw24 {
    out := &pbgear.Identity_Raw24{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw25(in any) *pbgear.Identity_Raw25 {
    out := &pbgear.Identity_Raw25{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw26(in any) *pbgear.Identity_Raw26 {
    out := &pbgear.Identity_Raw26{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw27(in any) *pbgear.Identity_Raw27 {
    out := &pbgear.Identity_Raw27{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw28(in any) *pbgear.Identity_Raw28 {
    out := &pbgear.Identity_Raw28{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw29(in any) *pbgear.Identity_Raw29 {
    out := &pbgear.Identity_Raw29{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw3(in any) *pbgear.Identity_Raw3 {
    out := &pbgear.Identity_Raw3{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw30(in any) *pbgear.Identity_Raw30 {
    out := &pbgear.Identity_Raw30{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw31(in any) *pbgear.Identity_Raw31 {
    out := &pbgear.Identity_Raw31{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw32(in any) *pbgear.Identity_Raw32 {
    out := &pbgear.Identity_Raw32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw4(in any) *pbgear.Identity_Raw4 {
    out := &pbgear.Identity_Raw4{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw5(in any) *pbgear.Identity_Raw5 {
    out := &pbgear.Identity_Raw5{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw6(in any) *pbgear.Identity_Raw6 {
    out := &pbgear.Identity_Raw6{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw7(in any) *pbgear.Identity_Raw7 {
    out := &pbgear.Identity_Raw7{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw8(in any) *pbgear.Identity_Raw8 {
    out := &pbgear.Identity_Raw8{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Raw9(in any) *pbgear.Identity_Raw9 {
    out := &pbgear.Identity_Raw9{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Reasonable(in any) *pbgear.Identity_Reasonable {
    out := &pbgear.Identity_Reasonable{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_RemoveSubCall(in any) *pbgear.Identity_RemoveSubCall {
    out := &pbgear.Identity_RemoveSubCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Sub
    v0 := To_OneOf_Identity_RemoveSubCall_sub(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_Identity_RemoveSubCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_RemoveSubCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityRemoveSubCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityRemoveSubCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Identity_RemoveSubCall_sub (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_RemoveSubCall_SubId{
        SubId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_RemoveSubCall_SubIndex{
        SubIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_RemoveSubCall_SubRaw{
        SubRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_RemoveSubCall_SubAddress32{
        SubAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_RemoveSubCall_SubAddress20{
        SubAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_RenameSubCall(in any) *pbgear.Identity_RenameSubCall {
    out := &pbgear.Identity_RenameSubCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Sub
    v0 := To_OneOf_Identity_RenameSubCall_sub(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))
    // oneOf field Data
    v1 := To_OneOf_Identity_RenameSubCall_data(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Data").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Identity_RenameSubCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_RenameSubCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityRenameSubCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityRenameSubCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Identity_RenameSubCall_sub (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_RenameSubCall_SubId{
        SubId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_RenameSubCall_SubIndex{
        SubIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_RenameSubCall_SubRaw{
        SubRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_RenameSubCall_SubAddress32{
        SubAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_RenameSubCall_SubAddress20{
        SubAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_RenameSubCall_data (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_RenameSubCall_DataNone{
        DataNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_RenameSubCall_DataRaw0{
        DataRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_RenameSubCall_DataRaw1{
        DataRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_RenameSubCall_DataRaw2{
        DataRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_RenameSubCall_DataRaw3{
        DataRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_RenameSubCall_DataRaw4{
        DataRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_RenameSubCall_DataRaw5{
        DataRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_RenameSubCall_DataRaw6{
        DataRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_RenameSubCall_DataRaw7{
        DataRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_RenameSubCall_DataRaw8{
        DataRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_RenameSubCall_DataRaw9{
        DataRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_RenameSubCall_DataRaw10{
        DataRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_RenameSubCall_DataRaw11{
        DataRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_RenameSubCall_DataRaw12{
        DataRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_RenameSubCall_DataRaw13{
        DataRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_RenameSubCall_DataRaw14{
        DataRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_RenameSubCall_DataRaw15{
        DataRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_RenameSubCall_DataRaw16{
        DataRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_RenameSubCall_DataRaw17{
        DataRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_RenameSubCall_DataRaw18{
        DataRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_RenameSubCall_DataRaw19{
        DataRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_RenameSubCall_DataRaw20{
        DataRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_RenameSubCall_DataRaw21{
        DataRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_RenameSubCall_DataRaw22{
        DataRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_RenameSubCall_DataRaw23{
        DataRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_RenameSubCall_DataRaw24{
        DataRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_RenameSubCall_DataRaw25{
        DataRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_RenameSubCall_DataRaw26{
        DataRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_RenameSubCall_DataRaw27{
        DataRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_RenameSubCall_DataRaw28{
        DataRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_RenameSubCall_DataRaw29{
        DataRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_RenameSubCall_DataRaw30{
        DataRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_RenameSubCall_DataRaw31{
        DataRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_RenameSubCall_DataRaw32{
        DataRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_RenameSubCall_DataBlakeTwo256{
        DataBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_RenameSubCall_DataSha256{
        DataSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_RenameSubCall_DataKeccak256{
        DataKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_RenameSubCall_DataShaThree256{
        DataShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_RequestJudgementCall(in any) *pbgear.Identity_RequestJudgementCall {
    out := &pbgear.Identity_RequestJudgementCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field RegIndex
    out.RegIndex = To_uint32(v.ValueAt(0))
    // primitive field MaxFee
    out.MaxFee = To_string(v.ValueAt(1))

    return out //from message
}
func To_Identity_RequestJudgementCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_RequestJudgementCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentityRequestJudgementCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentityRequestJudgementCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Identity_Riot(in any) *pbgear.Identity_Riot {
    out := &pbgear.Identity_Riot{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Riot
    v0 := To_OneOf_Identity_Riot_riot(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Riot").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Riot_riot (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Riot_RiotNone{
        RiotNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Riot_RiotRaw0{
        RiotRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Riot_RiotRaw1{
        RiotRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Riot_RiotRaw2{
        RiotRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Riot_RiotRaw3{
        RiotRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Riot_RiotRaw4{
        RiotRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Riot_RiotRaw5{
        RiotRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Riot_RiotRaw6{
        RiotRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Riot_RiotRaw7{
        RiotRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Riot_RiotRaw8{
        RiotRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Riot_RiotRaw9{
        RiotRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Riot_RiotRaw10{
        RiotRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Riot_RiotRaw11{
        RiotRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Riot_RiotRaw12{
        RiotRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Riot_RiotRaw13{
        RiotRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Riot_RiotRaw14{
        RiotRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Riot_RiotRaw15{
        RiotRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Riot_RiotRaw16{
        RiotRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Riot_RiotRaw17{
        RiotRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Riot_RiotRaw18{
        RiotRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Riot_RiotRaw19{
        RiotRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Riot_RiotRaw20{
        RiotRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Riot_RiotRaw21{
        RiotRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Riot_RiotRaw22{
        RiotRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Riot_RiotRaw23{
        RiotRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Riot_RiotRaw24{
        RiotRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Riot_RiotRaw25{
        RiotRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Riot_RiotRaw26{
        RiotRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Riot_RiotRaw27{
        RiotRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Riot_RiotRaw28{
        RiotRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Riot_RiotRaw29{
        RiotRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Riot_RiotRaw30{
        RiotRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Riot_RiotRaw31{
        RiotRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Riot_RiotRaw32{
        RiotRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Riot_RiotBlakeTwo256{
        RiotBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Riot_RiotSha256{
        RiotSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Riot_RiotKeccak256{
        RiotKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Riot_RiotShaThree256{
        RiotShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_SetAccountIdCall(in any) *pbgear.Identity_SetAccountIdCall {
    out := &pbgear.Identity_SetAccountIdCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // oneOf field New
    v1 := To_OneOf_Identity_SetAccountIdCall_new(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("New").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Identity_SetAccountIdCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_SetAccountIdCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentitySetAccountIdCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentitySetAccountIdCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Identity_SetAccountIdCall_new (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_SetAccountIdCall_NewId{
        NewId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_SetAccountIdCall_NewIndex{
        NewIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_SetAccountIdCall_NewRaw{
        NewRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_SetAccountIdCall_NewAddress32{
        NewAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_SetAccountIdCall_NewAddress20{
        NewAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_SetFeeCall(in any) *pbgear.Identity_SetFeeCall {
    out := &pbgear.Identity_SetFeeCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field Fee
    out.Fee = To_string(v.ValueAt(1))

    return out //from message
}
func To_Identity_SetFeeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_SetFeeCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentitySetFeeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentitySetFeeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Identity_SetFieldsCall(in any) *pbgear.Identity_SetFieldsCall {
    out := &pbgear.Identity_SetFieldsCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // field Fields To_Identity_PalletIdentityTypesBitFlags(w)
    out.Fields = To_Identity_PalletIdentityTypesBitFlags(v.ValueAt(1))

    return out //from message
}
func To_Identity_SetFieldsCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_SetFieldsCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentitySetFieldsCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentitySetFieldsCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Identity_SetIdentityCall(in any) *pbgear.Identity_SetIdentityCall {
    out := &pbgear.Identity_SetIdentityCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Info To_Identity_PalletIdentitySimpleIdentityInfo(w)
    out.Info = To_Identity_PalletIdentitySimpleIdentityInfo(v.ValueAt(0))

    return out //from message
}
func To_Identity_SetIdentityCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_SetIdentityCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentitySetIdentityCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentitySetIdentityCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Identity_SetSubsCall(in any) *pbgear.Identity_SetSubsCall {
    out := &pbgear.Identity_SetSubsCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Subs
    out.Subs = To_Repeated_Identity_SetSubsCall_subs(v.ValueAt(0))

    return out //from message
}
func To_Identity_SetSubsCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Identity_SetSubsCall(in)
    out := &pbgearextrinsic.RuntimeCall_IdentitySetSubsCall{ }
    reflect.ValueOf(out).Elem().FieldByName("IdentitySetSubsCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Identity_SetSubsCall_subs(in any) []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
    items := in.([]interface{})

    var out []*pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Identity_Sha256(in any) *pbgear.Identity_Sha256 {
    out := &pbgear.Identity_Sha256{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_ShaThree256(in any) *pbgear.Identity_ShaThree256 {
    out := &pbgear.Identity_ShaThree256{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Identity_Sub(in any) *pbgear.Identity_Sub {
    out := &pbgear.Identity_Sub{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Sub
    v0 := To_OneOf_Identity_Sub_sub(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Sub").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Sub_sub (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Sub_SubId{
        SubId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_Sub_SubIndex{
        SubIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_Sub_SubRaw{
        SubRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_Sub_SubAddress32{
        SubAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_Sub_SubAddress20{
        SubAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Target(in any) *pbgear.Identity_Target {
    out := &pbgear.Identity_Target{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Target
    v0 := To_OneOf_Identity_Target_target(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Target_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Target_TargetId{
        TargetId: To_Identity_Id(param),
        }
    case 1:
        return &pbgear.Identity_Target_TargetIndex{
        TargetIndex: To_Identity_Index(param),
        }
    case 2:
        return &pbgear.Identity_Target_TargetRaw{
        TargetRaw: To_Identity_Raw(param),
        }
    case 3:
        return &pbgear.Identity_Target_TargetAddress32{
        TargetAddress32: To_Identity_Address32(param),
        }
    case 4:
        return &pbgear.Identity_Target_TargetAddress20{
        TargetAddress20: To_Identity_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_TupleNull(in any) *pbgear.Identity_TupleNull {
    out := &pbgear.Identity_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData(in any) *pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData {
    out := &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))
    // oneOf field Value1
    v1 := To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value1(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v1))

    return out //from message
}
    
func To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0None{
        Value0None: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw0{
        Value0Raw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw1{
        Value0Raw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw2{
        Value0Raw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw3{
        Value0Raw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw4{
        Value0Raw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw5{
        Value0Raw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw6{
        Value0Raw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw7{
        Value0Raw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw8{
        Value0Raw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw9{
        Value0Raw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw10{
        Value0Raw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw11{
        Value0Raw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw12{
        Value0Raw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw13{
        Value0Raw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw14{
        Value0Raw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw15{
        Value0Raw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw16{
        Value0Raw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw17{
        Value0Raw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw18{
        Value0Raw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw19{
        Value0Raw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw20{
        Value0Raw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw21{
        Value0Raw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw22{
        Value0Raw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw23{
        Value0Raw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw24{
        Value0Raw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw25{
        Value0Raw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw26{
        Value0Raw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw27{
        Value0Raw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw28{
        Value0Raw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw29{
        Value0Raw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw30{
        Value0Raw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw31{
        Value0Raw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Raw32{
        Value0Raw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0BlakeTwo256{
        Value0BlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Sha256{
        Value0Sha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0Keccak256{
        Value0Keccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value0ShaThree256{
        Value0ShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_value1 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1None{
        Value1None: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw0{
        Value1Raw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw1{
        Value1Raw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw2{
        Value1Raw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw3{
        Value1Raw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw4{
        Value1Raw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw5{
        Value1Raw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw6{
        Value1Raw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw7{
        Value1Raw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw8{
        Value1Raw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw9{
        Value1Raw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw10{
        Value1Raw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw11{
        Value1Raw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw12{
        Value1Raw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw13{
        Value1Raw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw14{
        Value1Raw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw15{
        Value1Raw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw16{
        Value1Raw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw17{
        Value1Raw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw18{
        Value1Raw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw19{
        Value1Raw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw20{
        Value1Raw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw21{
        Value1Raw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw22{
        Value1Raw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw23{
        Value1Raw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw24{
        Value1Raw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw25{
        Value1Raw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw26{
        Value1Raw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw27{
        Value1Raw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw28{
        Value1Raw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw29{
        Value1Raw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw30{
        Value1Raw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw31{
        Value1Raw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Raw32{
        Value1Raw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1BlakeTwo256{
        Value1BlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Sha256{
        Value1Sha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1Keccak256{
        Value1Keccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_TuplePalletIdentityTypesDatapalletIdentityTypesData_Value1ShaThree256{
        Value1ShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData(in any) *pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData {
    out := &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // oneOf field Value1
    v1 := To_OneOf_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value1(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_value1 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1None{
        Value1None: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw0{
        Value1Raw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw1{
        Value1Raw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw2{
        Value1Raw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw3{
        Value1Raw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw4{
        Value1Raw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw5{
        Value1Raw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw6{
        Value1Raw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw7{
        Value1Raw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw8{
        Value1Raw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw9{
        Value1Raw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw10{
        Value1Raw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw11{
        Value1Raw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw12{
        Value1Raw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw13{
        Value1Raw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw14{
        Value1Raw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw15{
        Value1Raw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw16{
        Value1Raw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw17{
        Value1Raw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw18{
        Value1Raw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw19{
        Value1Raw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw20{
        Value1Raw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw21{
        Value1Raw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw22{
        Value1Raw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw23{
        Value1Raw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw24{
        Value1Raw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw25{
        Value1Raw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw26{
        Value1Raw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw27{
        Value1Raw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw28{
        Value1Raw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw29{
        Value1Raw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw30{
        Value1Raw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw31{
        Value1Raw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Raw32{
        Value1Raw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1BlakeTwo256{
        Value1BlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Sha256{
        Value1Sha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1Keccak256{
        Value1Keccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_TupleSpCoreCryptoAccountId32PalletIdentityTypesData_Value1ShaThree256{
        Value1ShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Twitter(in any) *pbgear.Identity_Twitter {
    out := &pbgear.Identity_Twitter{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Twitter
    v0 := To_OneOf_Identity_Twitter_twitter(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Twitter").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Twitter_twitter (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Twitter_TwitterNone{
        TwitterNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Twitter_TwitterRaw0{
        TwitterRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Twitter_TwitterRaw1{
        TwitterRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Twitter_TwitterRaw2{
        TwitterRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Twitter_TwitterRaw3{
        TwitterRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Twitter_TwitterRaw4{
        TwitterRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Twitter_TwitterRaw5{
        TwitterRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Twitter_TwitterRaw6{
        TwitterRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Twitter_TwitterRaw7{
        TwitterRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Twitter_TwitterRaw8{
        TwitterRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Twitter_TwitterRaw9{
        TwitterRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Twitter_TwitterRaw10{
        TwitterRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Twitter_TwitterRaw11{
        TwitterRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Twitter_TwitterRaw12{
        TwitterRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Twitter_TwitterRaw13{
        TwitterRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Twitter_TwitterRaw14{
        TwitterRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Twitter_TwitterRaw15{
        TwitterRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Twitter_TwitterRaw16{
        TwitterRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Twitter_TwitterRaw17{
        TwitterRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Twitter_TwitterRaw18{
        TwitterRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Twitter_TwitterRaw19{
        TwitterRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Twitter_TwitterRaw20{
        TwitterRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Twitter_TwitterRaw21{
        TwitterRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Twitter_TwitterRaw22{
        TwitterRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Twitter_TwitterRaw23{
        TwitterRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Twitter_TwitterRaw24{
        TwitterRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Twitter_TwitterRaw25{
        TwitterRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Twitter_TwitterRaw26{
        TwitterRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Twitter_TwitterRaw27{
        TwitterRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Twitter_TwitterRaw28{
        TwitterRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Twitter_TwitterRaw29{
        TwitterRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Twitter_TwitterRaw30{
        TwitterRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Twitter_TwitterRaw31{
        TwitterRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Twitter_TwitterRaw32{
        TwitterRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Twitter_TwitterBlakeTwo256{
        TwitterBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Twitter_TwitterSha256{
        TwitterSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Twitter_TwitterKeccak256{
        TwitterKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Twitter_TwitterShaThree256{
        TwitterShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Unknown(in any) *pbgear.Identity_Unknown {
    out := &pbgear.Identity_Unknown{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Identity_Value0(in any) *pbgear.Identity_Value0 {
    out := &pbgear.Identity_Value0{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Identity_Value0_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Value0_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Value0_Value0None{
        Value0None: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Value0_Value0Raw0{
        Value0Raw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Value0_Value0Raw1{
        Value0Raw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Value0_Value0Raw2{
        Value0Raw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Value0_Value0Raw3{
        Value0Raw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Value0_Value0Raw4{
        Value0Raw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Value0_Value0Raw5{
        Value0Raw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Value0_Value0Raw6{
        Value0Raw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Value0_Value0Raw7{
        Value0Raw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Value0_Value0Raw8{
        Value0Raw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Value0_Value0Raw9{
        Value0Raw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Value0_Value0Raw10{
        Value0Raw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Value0_Value0Raw11{
        Value0Raw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Value0_Value0Raw12{
        Value0Raw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Value0_Value0Raw13{
        Value0Raw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Value0_Value0Raw14{
        Value0Raw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Value0_Value0Raw15{
        Value0Raw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Value0_Value0Raw16{
        Value0Raw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Value0_Value0Raw17{
        Value0Raw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Value0_Value0Raw18{
        Value0Raw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Value0_Value0Raw19{
        Value0Raw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Value0_Value0Raw20{
        Value0Raw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Value0_Value0Raw21{
        Value0Raw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Value0_Value0Raw22{
        Value0Raw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Value0_Value0Raw23{
        Value0Raw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Value0_Value0Raw24{
        Value0Raw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Value0_Value0Raw25{
        Value0Raw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Value0_Value0Raw26{
        Value0Raw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Value0_Value0Raw27{
        Value0Raw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Value0_Value0Raw28{
        Value0Raw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Value0_Value0Raw29{
        Value0Raw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Value0_Value0Raw30{
        Value0Raw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Value0_Value0Raw31{
        Value0Raw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Value0_Value0Raw32{
        Value0Raw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Value0_Value0BlakeTwo256{
        Value0BlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Value0_Value0Sha256{
        Value0Sha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Value0_Value0Keccak256{
        Value0Keccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Value0_Value0ShaThree256{
        Value0ShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Value1(in any) *pbgear.Identity_Value1 {
    out := &pbgear.Identity_Value1{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value1
    v0 := To_OneOf_Identity_Value1_value1(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Value1_value1 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Value1_Value1None{
        Value1None: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Value1_Value1Raw0{
        Value1Raw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Value1_Value1Raw1{
        Value1Raw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Value1_Value1Raw2{
        Value1Raw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Value1_Value1Raw3{
        Value1Raw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Value1_Value1Raw4{
        Value1Raw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Value1_Value1Raw5{
        Value1Raw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Value1_Value1Raw6{
        Value1Raw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Value1_Value1Raw7{
        Value1Raw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Value1_Value1Raw8{
        Value1Raw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Value1_Value1Raw9{
        Value1Raw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Value1_Value1Raw10{
        Value1Raw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Value1_Value1Raw11{
        Value1Raw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Value1_Value1Raw12{
        Value1Raw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Value1_Value1Raw13{
        Value1Raw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Value1_Value1Raw14{
        Value1Raw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Value1_Value1Raw15{
        Value1Raw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Value1_Value1Raw16{
        Value1Raw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Value1_Value1Raw17{
        Value1Raw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Value1_Value1Raw18{
        Value1Raw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Value1_Value1Raw19{
        Value1Raw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Value1_Value1Raw20{
        Value1Raw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Value1_Value1Raw21{
        Value1Raw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Value1_Value1Raw22{
        Value1Raw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Value1_Value1Raw23{
        Value1Raw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Value1_Value1Raw24{
        Value1Raw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Value1_Value1Raw25{
        Value1Raw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Value1_Value1Raw26{
        Value1Raw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Value1_Value1Raw27{
        Value1Raw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Value1_Value1Raw28{
        Value1Raw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Value1_Value1Raw29{
        Value1Raw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Value1_Value1Raw30{
        Value1Raw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Value1_Value1Raw31{
        Value1Raw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Value1_Value1Raw32{
        Value1Raw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Value1_Value1BlakeTwo256{
        Value1BlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Value1_Value1Sha256{
        Value1Sha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Value1_Value1Keccak256{
        Value1Keccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Value1_Value1ShaThree256{
        Value1ShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Identity_Web(in any) *pbgear.Identity_Web {
    out := &pbgear.Identity_Web{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Web
    v0 := To_OneOf_Identity_Web_web(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Web").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Identity_Web_web (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Identity_Web_WebNone{
        WebNone: To_Identity_None(param),
        }
    case 1:
        return &pbgear.Identity_Web_WebRaw0{
        WebRaw0: To_Identity_Raw0(param),
        }
    case 2:
        return &pbgear.Identity_Web_WebRaw1{
        WebRaw1: To_Identity_Raw1(param),
        }
    case 3:
        return &pbgear.Identity_Web_WebRaw2{
        WebRaw2: To_Identity_Raw2(param),
        }
    case 4:
        return &pbgear.Identity_Web_WebRaw3{
        WebRaw3: To_Identity_Raw3(param),
        }
    case 5:
        return &pbgear.Identity_Web_WebRaw4{
        WebRaw4: To_Identity_Raw4(param),
        }
    case 6:
        return &pbgear.Identity_Web_WebRaw5{
        WebRaw5: To_Identity_Raw5(param),
        }
    case 7:
        return &pbgear.Identity_Web_WebRaw6{
        WebRaw6: To_Identity_Raw6(param),
        }
    case 8:
        return &pbgear.Identity_Web_WebRaw7{
        WebRaw7: To_Identity_Raw7(param),
        }
    case 9:
        return &pbgear.Identity_Web_WebRaw8{
        WebRaw8: To_Identity_Raw8(param),
        }
    case 10:
        return &pbgear.Identity_Web_WebRaw9{
        WebRaw9: To_Identity_Raw9(param),
        }
    case 11:
        return &pbgear.Identity_Web_WebRaw10{
        WebRaw10: To_Identity_Raw10(param),
        }
    case 12:
        return &pbgear.Identity_Web_WebRaw11{
        WebRaw11: To_Identity_Raw11(param),
        }
    case 13:
        return &pbgear.Identity_Web_WebRaw12{
        WebRaw12: To_Identity_Raw12(param),
        }
    case 14:
        return &pbgear.Identity_Web_WebRaw13{
        WebRaw13: To_Identity_Raw13(param),
        }
    case 15:
        return &pbgear.Identity_Web_WebRaw14{
        WebRaw14: To_Identity_Raw14(param),
        }
    case 16:
        return &pbgear.Identity_Web_WebRaw15{
        WebRaw15: To_Identity_Raw15(param),
        }
    case 17:
        return &pbgear.Identity_Web_WebRaw16{
        WebRaw16: To_Identity_Raw16(param),
        }
    case 18:
        return &pbgear.Identity_Web_WebRaw17{
        WebRaw17: To_Identity_Raw17(param),
        }
    case 19:
        return &pbgear.Identity_Web_WebRaw18{
        WebRaw18: To_Identity_Raw18(param),
        }
    case 20:
        return &pbgear.Identity_Web_WebRaw19{
        WebRaw19: To_Identity_Raw19(param),
        }
    case 21:
        return &pbgear.Identity_Web_WebRaw20{
        WebRaw20: To_Identity_Raw20(param),
        }
    case 22:
        return &pbgear.Identity_Web_WebRaw21{
        WebRaw21: To_Identity_Raw21(param),
        }
    case 23:
        return &pbgear.Identity_Web_WebRaw22{
        WebRaw22: To_Identity_Raw22(param),
        }
    case 24:
        return &pbgear.Identity_Web_WebRaw23{
        WebRaw23: To_Identity_Raw23(param),
        }
    case 25:
        return &pbgear.Identity_Web_WebRaw24{
        WebRaw24: To_Identity_Raw24(param),
        }
    case 26:
        return &pbgear.Identity_Web_WebRaw25{
        WebRaw25: To_Identity_Raw25(param),
        }
    case 27:
        return &pbgear.Identity_Web_WebRaw26{
        WebRaw26: To_Identity_Raw26(param),
        }
    case 28:
        return &pbgear.Identity_Web_WebRaw27{
        WebRaw27: To_Identity_Raw27(param),
        }
    case 29:
        return &pbgear.Identity_Web_WebRaw28{
        WebRaw28: To_Identity_Raw28(param),
        }
    case 30:
        return &pbgear.Identity_Web_WebRaw29{
        WebRaw29: To_Identity_Raw29(param),
        }
    case 31:
        return &pbgear.Identity_Web_WebRaw30{
        WebRaw30: To_Identity_Raw30(param),
        }
    case 32:
        return &pbgear.Identity_Web_WebRaw31{
        WebRaw31: To_Identity_Raw31(param),
        }
    case 33:
        return &pbgear.Identity_Web_WebRaw32{
        WebRaw32: To_Identity_Raw32(param),
        }
    case 34:
        return &pbgear.Identity_Web_WebBlakeTwo256{
        WebBlakeTwo256: To_Identity_BlakeTwo256(param),
        }
    case 35:
        return &pbgear.Identity_Web_WebSha256{
        WebSha256: To_Identity_Sha256(param),
        }
    case 36:
        return &pbgear.Identity_Web_WebKeccak256{
        WebKeccak256: To_Identity_Keccak256(param),
        }
    case 37:
        return &pbgear.Identity_Web_WebShaThree256{
        WebShaThree256: To_Identity_ShaThree256(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ImOnlineAllGoodEvent(in any) *pbgear.ImOnlineAllGoodEvent {
    out := &pbgear.ImOnlineAllGoodEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ImOnlineHeartbeatReceivedEvent(in any) *pbgear.ImOnlineHeartbeatReceivedEvent {
    out := &pbgear.ImOnlineHeartbeatReceivedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ImOnlinePallet(in any) *pbgear.ImOnlinePallet {
    out := &pbgear.ImOnlinePallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_ImOnlinePallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ImOnlinePallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ImOnlinePallet_CallHeartbeatCall{
        CallHeartbeatCall: To_ImOnline_HeartbeatCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ImOnlineSomeOfflineEvent(in any) *pbgear.ImOnlineSomeOfflineEvent {
    out := &pbgear.ImOnlineSomeOfflineEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ImOnline_HeartbeatCall(in any) *pbgear.ImOnline_HeartbeatCall {
    out := &pbgear.ImOnline_HeartbeatCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Heartbeat To_ImOnline_PalletImOnlineHeartbeat(w)
    out.Heartbeat = To_ImOnline_PalletImOnlineHeartbeat(v.ValueAt(0))
    // field Signature To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(w)
    out.Signature = To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(v.ValueAt(1))

    return out //from message
}
func To_ImOnline_HeartbeatCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_ImOnline_HeartbeatCall(in)
    out := &pbgearextrinsic.RuntimeCall_ImonlineHeartbeatCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ImonlineHeartbeatCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_ImOnline_PalletImOnlineHeartbeat(in any) *pbgear.ImOnline_PalletImOnlineHeartbeat {
    out := &pbgear.ImOnline_PalletImOnlineHeartbeat{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BlockNumber
    out.BlockNumber = To_uint32(v.ValueAt(0))
    // primitive field SessionIndex
    out.SessionIndex = To_uint32(v.ValueAt(1))
    // primitive field AuthorityIndex
    out.AuthorityIndex = To_uint32(v.ValueAt(2))
    // primitive field ValidatorsLen
    out.ValidatorsLen = To_uint32(v.ValueAt(3))

    return out //from message
}
    
func To_ImOnline_PalletImOnlineSr25519AppSr25519Public(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public {
    out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Public{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreSr25519Public(w)
    out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

    return out //from message
}
    

func To_ImOnline_PalletImOnlineSr25519AppSr25519Signature(in any) *pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature {
    out := &pbgear.ImOnline_PalletImOnlineSr25519AppSr25519Signature{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreSr25519Signature(w)
    out.Value0 = To_SpCoreSr25519Signature(v.ValueAt(0))

    return out //from message
}
    

func To_Inactive(in any) *pbgear.Inactive {
    out := &pbgear.Inactive{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_InactiveActor(in any) *pbgear.InactiveActor {
    out := &pbgear.InactiveActor{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Init(in any) *pbgear.Init {
    out := &pbgear.Init{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Inline(in any) *pbgear.Inline {
    out := &pbgear.Inline{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_BoundedCollectionsBoundedVecBoundedVec(w)
    out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(v.ValueAt(0))

    return out //from message
}
    

func To_Legacy(in any) *pbgear.Legacy {
    out := &pbgear.Legacy{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_LimitReached(in any) *pbgear.LimitReached {
    out := &pbgear.LimitReached{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Lookup(in any) *pbgear.Lookup {
    out := &pbgear.Lookup{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))
    // primitive field Len
    out.Len = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Mandatory(in any) *pbgear.Mandatory {
    out := &pbgear.Mandatory{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Manual(in any) *pbgear.Manual {
    out := &pbgear.Manual{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MediumSpender(in any) *pbgear.MediumSpender {
    out := &pbgear.MediumSpender{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MemoryOverflow(in any) *pbgear.MemoryOverflow {
    out := &pbgear.MemoryOverflow{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MessageClaimed(in any) *pbgear.MessageClaimed {
    out := &pbgear.MessageClaimed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MessageReplied(in any) *pbgear.MessageReplied {
    out := &pbgear.MessageReplied{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Module(in any) *pbgear.Module {
    out := &pbgear.Module{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpRuntimeModuleError(w)
    out.Value0 = To_SpRuntimeModuleError(v.ValueAt(0))

    return out //from message
}
    

func To_MultisigMultisigApprovalEvent(in any) *pbgear.MultisigMultisigApprovalEvent {
    out := &pbgear.MultisigMultisigApprovalEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MultisigMultisigCancelledEvent(in any) *pbgear.MultisigMultisigCancelledEvent {
    out := &pbgear.MultisigMultisigCancelledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MultisigMultisigExecutedEvent(in any) *pbgear.MultisigMultisigExecutedEvent {
    out := &pbgear.MultisigMultisigExecutedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MultisigNewMultisigEvent(in any) *pbgear.MultisigNewMultisigEvent {
    out := &pbgear.MultisigNewMultisigEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_MultisigPallet(in any) *pbgear.MultisigPallet {
    out := &pbgear.MultisigPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_MultisigPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_MultisigPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.MultisigPallet_CallAsMultiThreshold1Call{
        CallAsMultiThreshold1Call: To_Multisig_AsMultiThreshold1Call(param),
        }
    case 1:
        return &pbgear.MultisigPallet_CallAsMultiCall{
        CallAsMultiCall: To_Multisig_AsMultiCall(param),
        }
    case 2:
        return &pbgear.MultisigPallet_CallApproveAsMultiCall{
        CallApproveAsMultiCall: To_Multisig_ApproveAsMultiCall(param),
        }
    case 3:
        return &pbgear.MultisigPallet_CallCancelAsMultiCall{
        CallCancelAsMultiCall: To_Multisig_CancelAsMultiCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Multisig_ApproveAsMultiCall(in any) *pbgear.Multisig_ApproveAsMultiCall {
    out := &pbgear.Multisig_ApproveAsMultiCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Threshold
    out.Threshold = To_uint32(v.ValueAt(0))
    // repeated field OtherSignatories
    out.OtherSignatories = To_Repeated_Multisig_ApproveAsMultiCall_other_signatories(v.ValueAt(1))
    // optional field MaybeTimepoint
    if v.HasValue() {
        out.MaybeTimepoint = To_Optional_Multisig_ApproveAsMultiCall_maybe_timepoint(v.ValueAt(2))
    }
    // primitive field CallHash
    out.CallHash = To_bytes(v.ValueAt(3))
    // field MaxWeight To_SpWeightsWeightV2Weight(w)
    out.MaxWeight = To_SpWeightsWeightV2Weight(v.ValueAt(4))

    return out //from message
}
func To_Multisig_ApproveAsMultiCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Multisig_ApproveAsMultiCall(in)
    out := &pbgearextrinsic.RuntimeCall_MultisigApproveAsMultiCall{ }
    reflect.ValueOf(out).Elem().FieldByName("MultisigApproveAsMultiCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Multisig_ApproveAsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
    items := in.([]interface{})

    var out []*pbgear.SpCoreCryptoAccountId32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpCoreCryptoAccountId32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Optional_Multisig_ApproveAsMultiCall_maybe_timepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
    panic("Not implemented")
    return nil //funcForOptionalField
}

func To_Multisig_AsMultiCall(in any) *pbgear.Multisig_AsMultiCall {
    out := &pbgear.Multisig_AsMultiCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Threshold
    out.Threshold = To_uint32(v.ValueAt(0))
    // repeated field OtherSignatories
    out.OtherSignatories = To_Repeated_Multisig_AsMultiCall_other_signatories(v.ValueAt(1))
    // optional field MaybeTimepoint
    if v.HasValue() {
        out.MaybeTimepoint = To_Optional_Multisig_AsMultiCall_maybe_timepoint(v.ValueAt(2))
    }
    // oneOf field Call
    v3 := To_OneOf_Multisig_AsMultiCall_call(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))
    // field MaxWeight To_SpWeightsWeightV2Weight(w)
    out.MaxWeight = To_SpWeightsWeightV2Weight(v.ValueAt(4))

    return out //from message
}
func To_Multisig_AsMultiCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Multisig_AsMultiCall(in)
    out := &pbgearextrinsic.RuntimeCall_MultisigAsMultiCall{ }
    reflect.ValueOf(out).Elem().FieldByName("MultisigAsMultiCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Multisig_AsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
    items := in.([]interface{})

    var out []*pbgear.SpCoreCryptoAccountId32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpCoreCryptoAccountId32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Optional_Multisig_AsMultiCall_maybe_timepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_OneOf_Multisig_AsMultiCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Multisig_AsMultiCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Multisig_AsMultiCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Multisig_AsMultiCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Multisig_AsMultiCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Multisig_AsMultiCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Multisig_AsMultiCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Multisig_AsMultiCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Multisig_AsMultiCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Multisig_AsMultiCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Multisig_AsMultiCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Multisig_AsMultiCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Multisig_AsMultiCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Multisig_AsMultiCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Multisig_AsMultiCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Multisig_AsMultiCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Multisig_AsMultiCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Multisig_AsMultiCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Multisig_AsMultiCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Multisig_AsMultiCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Multisig_AsMultiCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Multisig_AsMultiCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Multisig_AsMultiCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Multisig_AsMultiCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Multisig_AsMultiCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Multisig_AsMultiCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Multisig_AsMultiCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Multisig_AsMultiCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Multisig_AsMultiCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Multisig_AsMultiCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Multisig_AsMultiThreshold1Call(in any) *pbgear.Multisig_AsMultiThreshold1Call {
    out := &pbgear.Multisig_AsMultiThreshold1Call{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field OtherSignatories
    out.OtherSignatories = To_Repeated_Multisig_AsMultiThreshold1Call_other_signatories(v.ValueAt(0))
    // oneOf field Call
    v1 := To_OneOf_Multisig_AsMultiThreshold1Call_call(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Multisig_AsMultiThreshold1Call_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Multisig_AsMultiThreshold1Call(in)
    out := &pbgearextrinsic.RuntimeCall_MultisigAsMultiThreshold_1Call{ }
    reflect.ValueOf(out).Elem().FieldByName("MultisigAsMultiThreshold_1Call").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Multisig_AsMultiThreshold1Call_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
    items := in.([]interface{})

    var out []*pbgear.SpCoreCryptoAccountId32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpCoreCryptoAccountId32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_OneOf_Multisig_AsMultiThreshold1Call_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Multisig_AsMultiThreshold1Call_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Multisig_CancelAsMultiCall(in any) *pbgear.Multisig_CancelAsMultiCall {
    out := &pbgear.Multisig_CancelAsMultiCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Threshold
    out.Threshold = To_uint32(v.ValueAt(0))
    // repeated field OtherSignatories
    out.OtherSignatories = To_Repeated_Multisig_CancelAsMultiCall_other_signatories(v.ValueAt(1))
    // field Timepoint To_Multisig_PalletMultisigTimepoint(w)
    out.Timepoint = To_Multisig_PalletMultisigTimepoint(v.ValueAt(2))
    // primitive field CallHash
    out.CallHash = To_bytes(v.ValueAt(3))

    return out //from message
}
func To_Multisig_CancelAsMultiCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Multisig_CancelAsMultiCall(in)
    out := &pbgearextrinsic.RuntimeCall_MultisigCancelAsMultiCall{ }
    reflect.ValueOf(out).Elem().FieldByName("MultisigCancelAsMultiCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Multisig_CancelAsMultiCall_other_signatories(in any) []*pbgear.SpCoreCryptoAccountId32 {
    items := in.([]interface{})

    var out []*pbgear.SpCoreCryptoAccountId32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpCoreCryptoAccountId32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Multisig_PalletMultisigTimepoint(in any) *pbgear.Multisig_PalletMultisigTimepoint {
    out := &pbgear.Multisig_PalletMultisigTimepoint{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Height
    out.Height = To_uint32(v.ValueAt(0))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_No(in any) *pbgear.No {
    out := &pbgear.No{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NoLayer(in any) *pbgear.NoLayer {
    out := &pbgear.NoLayer{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NoProviders(in any) *pbgear.NoProviders {
    out := &pbgear.NoProviders{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Node(in any) *pbgear.Node {
    out := &pbgear.Node{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_GprimitivesMessageId(w)
    out.Value0 = To_GprimitivesMessageId(v.ValueAt(0))

    return out //from message
}
    

func To_NominationPoolsBondedEvent(in any) *pbgear.NominationPoolsBondedEvent {
    out := &pbgear.NominationPoolsBondedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsCreatedEvent(in any) *pbgear.NominationPoolsCreatedEvent {
    out := &pbgear.NominationPoolsCreatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsDestroyedEvent(in any) *pbgear.NominationPoolsDestroyedEvent {
    out := &pbgear.NominationPoolsDestroyedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsMemberRemovedEvent(in any) *pbgear.NominationPoolsMemberRemovedEvent {
    out := &pbgear.NominationPoolsMemberRemovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsMinBalanceDeficitAdjustedEvent(in any) *pbgear.NominationPoolsMinBalanceDeficitAdjustedEvent {
    out := &pbgear.NominationPoolsMinBalanceDeficitAdjustedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsMinBalanceExcessAdjustedEvent(in any) *pbgear.NominationPoolsMinBalanceExcessAdjustedEvent {
    out := &pbgear.NominationPoolsMinBalanceExcessAdjustedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsPaidOutEvent(in any) *pbgear.NominationPoolsPaidOutEvent {
    out := &pbgear.NominationPoolsPaidOutEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsPallet(in any) *pbgear.NominationPoolsPallet {
    out := &pbgear.NominationPoolsPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_NominationPoolsPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPoolsPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPoolsPallet_CallJoinCall{
        CallJoinCall: To_NominationPools_JoinCall(param),
        }
    case 1:
        return &pbgear.NominationPoolsPallet_CallBondExtraCall{
        CallBondExtraCall: To_NominationPools_BondExtraCall(param),
        }
    case 2:
        return &pbgear.NominationPoolsPallet_CallClaimPayoutCall{
        CallClaimPayoutCall: To_NominationPools_ClaimPayoutCall(param),
        }
    case 3:
        return &pbgear.NominationPoolsPallet_CallUnbondCall{
        CallUnbondCall: To_NominationPools_UnbondCall(param),
        }
    case 4:
        return &pbgear.NominationPoolsPallet_CallPoolWithdrawUnbondedCall{
        CallPoolWithdrawUnbondedCall: To_NominationPools_PoolWithdrawUnbondedCall(param),
        }
    case 5:
        return &pbgear.NominationPoolsPallet_CallWithdrawUnbondedCall{
        CallWithdrawUnbondedCall: To_NominationPools_WithdrawUnbondedCall(param),
        }
    case 6:
        return &pbgear.NominationPoolsPallet_CallCreateCall{
        CallCreateCall: To_NominationPools_CreateCall(param),
        }
    case 7:
        return &pbgear.NominationPoolsPallet_CallCreateWithPoolIdCall{
        CallCreateWithPoolIdCall: To_NominationPools_CreateWithPoolIdCall(param),
        }
    case 8:
        return &pbgear.NominationPoolsPallet_CallNominateCall{
        CallNominateCall: To_NominationPools_NominateCall(param),
        }
    case 9:
        return &pbgear.NominationPoolsPallet_CallSetStateCall{
        CallSetStateCall: To_NominationPools_SetStateCall(param),
        }
    case 10:
        return &pbgear.NominationPoolsPallet_CallSetMetadataCall{
        CallSetMetadataCall: To_NominationPools_SetMetadataCall(param),
        }
    case 11:
        return &pbgear.NominationPoolsPallet_CallSetConfigsCall{
        CallSetConfigsCall: To_NominationPools_SetConfigsCall(param),
        }
    case 12:
        return &pbgear.NominationPoolsPallet_CallUpdateRolesCall{
        CallUpdateRolesCall: To_NominationPools_UpdateRolesCall(param),
        }
    case 13:
        return &pbgear.NominationPoolsPallet_CallChillCall{
        CallChillCall: To_NominationPools_ChillCall(param),
        }
    case 14:
        return &pbgear.NominationPoolsPallet_CallBondExtraOtherCall{
        CallBondExtraOtherCall: To_NominationPools_BondExtraOtherCall(param),
        }
    case 15:
        return &pbgear.NominationPoolsPallet_CallSetClaimPermissionCall{
        CallSetClaimPermissionCall: To_NominationPools_SetClaimPermissionCall(param),
        }
    case 16:
        return &pbgear.NominationPoolsPallet_CallClaimPayoutOtherCall{
        CallClaimPayoutOtherCall: To_NominationPools_ClaimPayoutOtherCall(param),
        }
    case 17:
        return &pbgear.NominationPoolsPallet_CallSetCommissionCall{
        CallSetCommissionCall: To_NominationPools_SetCommissionCall(param),
        }
    case 18:
        return &pbgear.NominationPoolsPallet_CallSetCommissionMaxCall{
        CallSetCommissionMaxCall: To_NominationPools_SetCommissionMaxCall(param),
        }
    case 19:
        return &pbgear.NominationPoolsPallet_CallSetCommissionChangeRateCall{
        CallSetCommissionChangeRateCall: To_NominationPools_SetCommissionChangeRateCall(param),
        }
    case 20:
        return &pbgear.NominationPoolsPallet_CallClaimCommissionCall{
        CallClaimCommissionCall: To_NominationPools_ClaimCommissionCall(param),
        }
    case 21:
        return &pbgear.NominationPoolsPallet_CallAdjustPoolDepositCall{
        CallAdjustPoolDepositCall: To_NominationPools_AdjustPoolDepositCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPoolsPoolCommissionChangeRateUpdatedEvent(in any) *pbgear.NominationPoolsPoolCommissionChangeRateUpdatedEvent {
    out := &pbgear.NominationPoolsPoolCommissionChangeRateUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsPoolCommissionClaimedEvent(in any) *pbgear.NominationPoolsPoolCommissionClaimedEvent {
    out := &pbgear.NominationPoolsPoolCommissionClaimedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsPoolCommissionUpdatedEvent(in any) *pbgear.NominationPoolsPoolCommissionUpdatedEvent {
    out := &pbgear.NominationPoolsPoolCommissionUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsPoolMaxCommissionUpdatedEvent(in any) *pbgear.NominationPoolsPoolMaxCommissionUpdatedEvent {
    out := &pbgear.NominationPoolsPoolMaxCommissionUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsPoolSlashedEvent(in any) *pbgear.NominationPoolsPoolSlashedEvent {
    out := &pbgear.NominationPoolsPoolSlashedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsRolesUpdatedEvent(in any) *pbgear.NominationPoolsRolesUpdatedEvent {
    out := &pbgear.NominationPoolsRolesUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsStateChangedEvent(in any) *pbgear.NominationPoolsStateChangedEvent {
    out := &pbgear.NominationPoolsStateChangedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsUnbondedEvent(in any) *pbgear.NominationPoolsUnbondedEvent {
    out := &pbgear.NominationPoolsUnbondedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsUnbondingPoolSlashedEvent(in any) *pbgear.NominationPoolsUnbondingPoolSlashedEvent {
    out := &pbgear.NominationPoolsUnbondingPoolSlashedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPoolsWithdrawnEvent(in any) *pbgear.NominationPoolsWithdrawnEvent {
    out := &pbgear.NominationPoolsWithdrawnEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_Address20(in any) *pbgear.NominationPools_Address20 {
    out := &pbgear.NominationPools_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_NominationPools_Address32(in any) *pbgear.NominationPools_Address32 {
    out := &pbgear.NominationPools_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_NominationPools_AdjustPoolDepositCall(in any) *pbgear.NominationPools_AdjustPoolDepositCall {
    out := &pbgear.NominationPools_AdjustPoolDepositCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_NominationPools_AdjustPoolDepositCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_AdjustPoolDepositCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsAdjustPoolDepositCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsAdjustPoolDepositCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_NominationPools_Blocked(in any) *pbgear.NominationPools_Blocked {
    out := &pbgear.NominationPools_Blocked{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_BondExtraCall(in any) *pbgear.NominationPools_BondExtraCall {
    out := &pbgear.NominationPools_BondExtraCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Extra
    v0 := To_OneOf_NominationPools_BondExtraCall_extra(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Extra").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_NominationPools_BondExtraCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_BondExtraCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsBondExtraCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsBondExtraCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_BondExtraCall_extra (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_BondExtraCall_ExtraFreeBalance{
        ExtraFreeBalance: To_NominationPools_FreeBalance(param),
        }
    case 1:
        return &pbgear.NominationPools_BondExtraCall_ExtraRewards{
        ExtraRewards: To_NominationPools_Rewards(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_BondExtraOtherCall(in any) *pbgear.NominationPools_BondExtraOtherCall {
    out := &pbgear.NominationPools_BondExtraOtherCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Member
    v0 := To_OneOf_NominationPools_BondExtraOtherCall_member(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Member").Set(reflect.ValueOf(v0))
    // oneOf field Extra
    v1 := To_OneOf_NominationPools_BondExtraOtherCall_extra(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Extra").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_NominationPools_BondExtraOtherCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_BondExtraOtherCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsBondExtraOtherCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsBondExtraOtherCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_BondExtraOtherCall_member (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_BondExtraOtherCall_MemberId{
        MemberId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_BondExtraOtherCall_MemberIndex{
        MemberIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_BondExtraOtherCall_MemberRaw{
        MemberRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_BondExtraOtherCall_MemberAddress32{
        MemberAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_BondExtraOtherCall_MemberAddress20{
        MemberAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_BondExtraOtherCall_extra (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_BondExtraOtherCall_ExtraFreeBalance{
        ExtraFreeBalance: To_NominationPools_FreeBalance(param),
        }
    case 1:
        return &pbgear.NominationPools_BondExtraOtherCall_ExtraRewards{
        ExtraRewards: To_NominationPools_Rewards(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_Bouncer(in any) *pbgear.NominationPools_Bouncer {
    out := &pbgear.NominationPools_Bouncer{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Bouncer
    v0 := To_OneOf_NominationPools_Bouncer_bouncer(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Bouncer").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_Bouncer_bouncer (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_Bouncer_BouncerId{
        BouncerId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_Bouncer_BouncerIndex{
        BouncerIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_Bouncer_BouncerRaw{
        BouncerRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_Bouncer_BouncerAddress32{
        BouncerAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_Bouncer_BouncerAddress20{
        BouncerAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_ChillCall(in any) *pbgear.NominationPools_ChillCall {
    out := &pbgear.NominationPools_ChillCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_NominationPools_ChillCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_ChillCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsChillCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsChillCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_NominationPools_ClaimCommissionCall(in any) *pbgear.NominationPools_ClaimCommissionCall {
    out := &pbgear.NominationPools_ClaimCommissionCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_NominationPools_ClaimCommissionCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_ClaimCommissionCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsClaimCommissionCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsClaimCommissionCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_NominationPools_ClaimPayoutCall(in any) *pbgear.NominationPools_ClaimPayoutCall {
    out := &pbgear.NominationPools_ClaimPayoutCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_ClaimPayoutOtherCall(in any) *pbgear.NominationPools_ClaimPayoutOtherCall {
    out := &pbgear.NominationPools_ClaimPayoutOtherCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Other To_SpCoreCryptoAccountId32(w)
    out.Other = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
func To_NominationPools_ClaimPayoutOtherCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_ClaimPayoutOtherCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsClaimPayoutOtherCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsClaimPayoutOtherCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_NominationPools_ConsensusEvent(in any) *pbgear.NominationPools_ConsensusEvent {
    out := &pbgear.NominationPools_ConsensusEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_NominationPools_CreateCall(in any) *pbgear.NominationPools_CreateCall {
    out := &pbgear.NominationPools_CreateCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Amount
    out.Amount = To_string(v.ValueAt(0))
    // oneOf field Root
    v1 := To_OneOf_NominationPools_CreateCall_root(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Root").Set(reflect.ValueOf(v1))
    // oneOf field Nominator
    v2 := To_OneOf_NominationPools_CreateCall_nominator(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Nominator").Set(reflect.ValueOf(v2))
    // oneOf field Bouncer
    v3 := To_OneOf_NominationPools_CreateCall_bouncer(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Bouncer").Set(reflect.ValueOf(v3))

    return out //from message
}
func To_NominationPools_CreateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_CreateCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsCreateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsCreateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_CreateCall_root (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_CreateCall_RootId{
        RootId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_CreateCall_RootIndex{
        RootIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_CreateCall_RootRaw{
        RootRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_CreateCall_RootAddress32{
        RootAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_CreateCall_RootAddress20{
        RootAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_CreateCall_nominator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_CreateCall_NominatorId{
        NominatorId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_CreateCall_NominatorIndex{
        NominatorIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_CreateCall_NominatorRaw{
        NominatorRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_CreateCall_NominatorAddress32{
        NominatorAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_CreateCall_NominatorAddress20{
        NominatorAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_CreateCall_bouncer (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_CreateCall_BouncerId{
        BouncerId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_CreateCall_BouncerIndex{
        BouncerIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_CreateCall_BouncerRaw{
        BouncerRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_CreateCall_BouncerAddress32{
        BouncerAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_CreateCall_BouncerAddress20{
        BouncerAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_CreateWithPoolIdCall(in any) *pbgear.NominationPools_CreateWithPoolIdCall {
    out := &pbgear.NominationPools_CreateWithPoolIdCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Amount
    out.Amount = To_string(v.ValueAt(0))
    // oneOf field Root
    v1 := To_OneOf_NominationPools_CreateWithPoolIdCall_root(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Root").Set(reflect.ValueOf(v1))
    // oneOf field Nominator
    v2 := To_OneOf_NominationPools_CreateWithPoolIdCall_nominator(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Nominator").Set(reflect.ValueOf(v2))
    // oneOf field Bouncer
    v3 := To_OneOf_NominationPools_CreateWithPoolIdCall_bouncer(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Bouncer").Set(reflect.ValueOf(v3))
    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(4))

    return out //from message
}
func To_NominationPools_CreateWithPoolIdCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_CreateWithPoolIdCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsCreateWithPoolIdCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsCreateWithPoolIdCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_CreateWithPoolIdCall_root (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_CreateWithPoolIdCall_RootId{
        RootId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_CreateWithPoolIdCall_RootIndex{
        RootIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_CreateWithPoolIdCall_RootRaw{
        RootRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_CreateWithPoolIdCall_RootAddress32{
        RootAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_CreateWithPoolIdCall_RootAddress20{
        RootAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_CreateWithPoolIdCall_nominator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorId{
        NominatorId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorIndex{
        NominatorIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorRaw{
        NominatorRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorAddress32{
        NominatorAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_CreateWithPoolIdCall_NominatorAddress20{
        NominatorAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_CreateWithPoolIdCall_bouncer (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerId{
        BouncerId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerIndex{
        BouncerIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerRaw{
        BouncerRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerAddress32{
        BouncerAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_CreateWithPoolIdCall_BouncerAddress20{
        BouncerAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_Destroying(in any) *pbgear.NominationPools_Destroying {
    out := &pbgear.NominationPools_Destroying{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_Extra(in any) *pbgear.NominationPools_Extra {
    out := &pbgear.NominationPools_Extra{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Extra
    v0 := To_OneOf_NominationPools_Extra_extra(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Extra").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_Extra_extra (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_Extra_ExtraFreeBalance{
        ExtraFreeBalance: To_NominationPools_FreeBalance(param),
        }
    case 1:
        return &pbgear.NominationPools_Extra_ExtraRewards{
        ExtraRewards: To_NominationPools_Rewards(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_FreeBalance(in any) *pbgear.NominationPools_FreeBalance {
    out := &pbgear.NominationPools_FreeBalance{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_NominationPools_GlobalMaxCommission(in any) *pbgear.NominationPools_GlobalMaxCommission {
    out := &pbgear.NominationPools_GlobalMaxCommission{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field GlobalMaxCommission
    v0 := To_OneOf_NominationPools_GlobalMaxCommission_global_max_commission(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("GlobalMaxCommission").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_GlobalMaxCommission_global_max_commission (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_GlobalMaxCommission_GlobalMaxCommissionNoop{
        GlobalMaxCommissionNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_GlobalMaxCommission_GlobalMaxCommissionSet{
        GlobalMaxCommissionSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_GlobalMaxCommission_GlobalMaxCommissionRemove{
        GlobalMaxCommissionRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_Id(in any) *pbgear.NominationPools_Id {
    out := &pbgear.NominationPools_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_NominationPools_Index(in any) *pbgear.NominationPools_Index {
    out := &pbgear.NominationPools_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_NominationPools_TupleNull(w)
    out.Value0 = To_NominationPools_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_NominationPools_JoinCall(in any) *pbgear.NominationPools_JoinCall {
    out := &pbgear.NominationPools_JoinCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Amount
    out.Amount = To_string(v.ValueAt(0))
    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_JoinCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_JoinCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsJoinCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsJoinCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_NominationPools_MaxMembers(in any) *pbgear.NominationPools_MaxMembers {
    out := &pbgear.NominationPools_MaxMembers{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MaxMembers
    v0 := To_OneOf_NominationPools_MaxMembers_max_members(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MaxMembers").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_MaxMembers_max_members (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_MaxMembers_MaxMembersNoop{
        MaxMembersNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_MaxMembers_MaxMembersSet{
        MaxMembersSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_MaxMembers_MaxMembersRemove{
        MaxMembersRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_MaxMembersPerPool(in any) *pbgear.NominationPools_MaxMembersPerPool {
    out := &pbgear.NominationPools_MaxMembersPerPool{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MaxMembersPerPool
    v0 := To_OneOf_NominationPools_MaxMembersPerPool_max_members_per_pool(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MaxMembersPerPool").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_MaxMembersPerPool_max_members_per_pool (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_MaxMembersPerPool_MaxMembersPerPoolNoop{
        MaxMembersPerPoolNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_MaxMembersPerPool_MaxMembersPerPoolSet{
        MaxMembersPerPoolSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_MaxMembersPerPool_MaxMembersPerPoolRemove{
        MaxMembersPerPoolRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_MaxPools(in any) *pbgear.NominationPools_MaxPools {
    out := &pbgear.NominationPools_MaxPools{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MaxPools
    v0 := To_OneOf_NominationPools_MaxPools_max_pools(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MaxPools").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_MaxPools_max_pools (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_MaxPools_MaxPoolsNoop{
        MaxPoolsNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_MaxPools_MaxPoolsSet{
        MaxPoolsSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_MaxPools_MaxPoolsRemove{
        MaxPoolsRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_Member(in any) *pbgear.NominationPools_Member {
    out := &pbgear.NominationPools_Member{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Member
    v0 := To_OneOf_NominationPools_Member_member(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Member").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_Member_member (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_Member_MemberId{
        MemberId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_Member_MemberIndex{
        MemberIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_Member_MemberRaw{
        MemberRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_Member_MemberAddress32{
        MemberAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_Member_MemberAddress20{
        MemberAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_MemberAccount(in any) *pbgear.NominationPools_MemberAccount {
    out := &pbgear.NominationPools_MemberAccount{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MemberAccount
    v0 := To_OneOf_NominationPools_MemberAccount_member_account(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MemberAccount").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_MemberAccount_member_account (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_MemberAccount_MemberAccountId{
        MemberAccountId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_MemberAccount_MemberAccountIndex{
        MemberAccountIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_MemberAccount_MemberAccountRaw{
        MemberAccountRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_MemberAccount_MemberAccountAddress32{
        MemberAccountAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_MemberAccount_MemberAccountAddress20{
        MemberAccountAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_MinCreateBond(in any) *pbgear.NominationPools_MinCreateBond {
    out := &pbgear.NominationPools_MinCreateBond{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MinCreateBond
    v0 := To_OneOf_NominationPools_MinCreateBond_min_create_bond(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MinCreateBond").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_MinCreateBond_min_create_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_MinCreateBond_MinCreateBondNoop{
        MinCreateBondNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_MinCreateBond_MinCreateBondSet{
        MinCreateBondSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_MinCreateBond_MinCreateBondRemove{
        MinCreateBondRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_MinJoinBond(in any) *pbgear.NominationPools_MinJoinBond {
    out := &pbgear.NominationPools_MinJoinBond{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MinJoinBond
    v0 := To_OneOf_NominationPools_MinJoinBond_min_join_bond(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MinJoinBond").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_MinJoinBond_min_join_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_MinJoinBond_MinJoinBondNoop{
        MinJoinBondNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_MinJoinBond_MinJoinBondSet{
        MinJoinBondSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_MinJoinBond_MinJoinBondRemove{
        MinJoinBondRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_NewBouncer(in any) *pbgear.NominationPools_NewBouncer {
    out := &pbgear.NominationPools_NewBouncer{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field NewBouncer
    v0 := To_OneOf_NominationPools_NewBouncer_new_bouncer(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("NewBouncer").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_NewBouncer_new_bouncer (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_NewBouncer_NewBouncerNoop{
        NewBouncerNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_NewBouncer_NewBouncerSet{
        NewBouncerSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_NewBouncer_NewBouncerRemove{
        NewBouncerRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_NewNominator(in any) *pbgear.NominationPools_NewNominator {
    out := &pbgear.NominationPools_NewNominator{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field NewNominator
    v0 := To_OneOf_NominationPools_NewNominator_new_nominator(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("NewNominator").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_NewNominator_new_nominator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_NewNominator_NewNominatorNoop{
        NewNominatorNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_NewNominator_NewNominatorSet{
        NewNominatorSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_NewNominator_NewNominatorRemove{
        NewNominatorRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_NewRoot(in any) *pbgear.NominationPools_NewRoot {
    out := &pbgear.NominationPools_NewRoot{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field NewRoot
    v0 := To_OneOf_NominationPools_NewRoot_new_root(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("NewRoot").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_NewRoot_new_root (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_NewRoot_NewRootNoop{
        NewRootNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_NewRoot_NewRootSet{
        NewRootSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_NewRoot_NewRootRemove{
        NewRootRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_NewState(in any) *pbgear.NominationPools_NewState {
    out := &pbgear.NominationPools_NewState{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field NewState
    v0 := To_OneOf_NominationPools_NewState_new_state(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("NewState").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_NewState_new_state (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_NewState_NewStateOpen{
        NewStateOpen: To_NominationPools_Open(param),
        }
    case 1:
        return &pbgear.NominationPools_NewState_NewStateBlocked{
        NewStateBlocked: To_NominationPools_Blocked(param),
        }
    case 2:
        return &pbgear.NominationPools_NewState_NewStateDestroying{
        NewStateDestroying: To_NominationPools_Destroying(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_NominateCall(in any) *pbgear.NominationPools_NominateCall {
    out := &pbgear.NominationPools_NominateCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // repeated field Validators
    out.Validators = To_Repeated_NominationPools_NominateCall_validators(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_NominateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_NominateCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsNominateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsNominateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_NominationPools_NominateCall_validators(in any) []*pbgear.SpCoreCryptoAccountId32 {
    items := in.([]interface{})

    var out []*pbgear.SpCoreCryptoAccountId32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpCoreCryptoAccountId32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_NominationPools_Nominator(in any) *pbgear.NominationPools_Nominator {
    out := &pbgear.NominationPools_Nominator{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Nominator
    v0 := To_OneOf_NominationPools_Nominator_nominator(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Nominator").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_Nominator_nominator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_Nominator_NominatorId{
        NominatorId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_Nominator_NominatorIndex{
        NominatorIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_Nominator_NominatorRaw{
        NominatorRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_Nominator_NominatorAddress32{
        NominatorAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_Nominator_NominatorAddress20{
        NominatorAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_Noop(in any) *pbgear.NominationPools_Noop {
    out := &pbgear.NominationPools_Noop{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_Open(in any) *pbgear.NominationPools_Open {
    out := &pbgear.NominationPools_Open{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_OtherEvent(in any) *pbgear.NominationPools_OtherEvent {
    out := &pbgear.NominationPools_OtherEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_NominationPools_PalletNominationPoolsCommissionChangeRate(in any) *pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate {
    out := &pbgear.NominationPools_PalletNominationPoolsCommissionChangeRate{}
    v := in.(registry.Valuable)
    _ = v

    // field MaxIncrease To_SpArithmeticPerThingsPerbill(w)
    out.MaxIncrease = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))
    // primitive field MinDelay
    out.MinDelay = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_NominationPools_Permission(in any) *pbgear.NominationPools_Permission {
    out := &pbgear.NominationPools_Permission{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Permission
    v0 := To_OneOf_NominationPools_Permission_permission(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Permission").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_Permission_permission (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_Permission_PermissionPermissioned{
        PermissionPermissioned: To_NominationPools_Permissioned(param),
        }
    case 1:
        return &pbgear.NominationPools_Permission_PermissionPermissionlessCompound{
        PermissionPermissionlessCompound: To_NominationPools_PermissionlessCompound(param),
        }
    case 2:
        return &pbgear.NominationPools_Permission_PermissionPermissionlessWithdraw{
        PermissionPermissionlessWithdraw: To_NominationPools_PermissionlessWithdraw(param),
        }
    case 3:
        return &pbgear.NominationPools_Permission_PermissionPermissionlessAll{
        PermissionPermissionlessAll: To_NominationPools_PermissionlessAll(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_Permissioned(in any) *pbgear.NominationPools_Permissioned {
    out := &pbgear.NominationPools_Permissioned{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_PermissionlessAll(in any) *pbgear.NominationPools_PermissionlessAll {
    out := &pbgear.NominationPools_PermissionlessAll{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_PermissionlessCompound(in any) *pbgear.NominationPools_PermissionlessCompound {
    out := &pbgear.NominationPools_PermissionlessCompound{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_PermissionlessWithdraw(in any) *pbgear.NominationPools_PermissionlessWithdraw {
    out := &pbgear.NominationPools_PermissionlessWithdraw{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_PoolWithdrawUnbondedCall(in any) *pbgear.NominationPools_PoolWithdrawUnbondedCall {
    out := &pbgear.NominationPools_PoolWithdrawUnbondedCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // primitive field NumSlashingSpans
    out.NumSlashingSpans = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_PoolWithdrawUnbondedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_PoolWithdrawUnbondedCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsPoolWithdrawUnbondedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsPoolWithdrawUnbondedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_NominationPools_PreRuntimeEvent(in any) *pbgear.NominationPools_PreRuntimeEvent {
    out := &pbgear.NominationPools_PreRuntimeEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_NominationPools_Raw(in any) *pbgear.NominationPools_Raw {
    out := &pbgear.NominationPools_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_NominationPools_Remove(in any) *pbgear.NominationPools_Remove {
    out := &pbgear.NominationPools_Remove{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_Rewards(in any) *pbgear.NominationPools_Rewards {
    out := &pbgear.NominationPools_Rewards{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_Root(in any) *pbgear.NominationPools_Root {
    out := &pbgear.NominationPools_Root{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Root
    v0 := To_OneOf_NominationPools_Root_root(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Root").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_Root_root (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_Root_RootId{
        RootId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_Root_RootIndex{
        RootIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_Root_RootRaw{
        RootRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_Root_RootAddress32{
        RootAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_Root_RootAddress20{
        RootAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_RuntimeEnvironmentUpdatedEvent(in any) *pbgear.NominationPools_RuntimeEnvironmentUpdatedEvent {
    out := &pbgear.NominationPools_RuntimeEnvironmentUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NominationPools_SealEvent(in any) *pbgear.NominationPools_SealEvent {
    out := &pbgear.NominationPools_SealEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_NominationPools_Set(in any) *pbgear.NominationPools_Set {
    out := &pbgear.NominationPools_Set{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_NominationPools_SetClaimPermissionCall(in any) *pbgear.NominationPools_SetClaimPermissionCall {
    out := &pbgear.NominationPools_SetClaimPermissionCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Permission
    v0 := To_OneOf_NominationPools_SetClaimPermissionCall_permission(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Permission").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_NominationPools_SetClaimPermissionCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_SetClaimPermissionCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsSetClaimPermissionCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsSetClaimPermissionCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_SetClaimPermissionCall_permission (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissioned{
        PermissionPermissioned: To_NominationPools_Permissioned(param),
        }
    case 1:
        return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissionlessCompound{
        PermissionPermissionlessCompound: To_NominationPools_PermissionlessCompound(param),
        }
    case 2:
        return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissionlessWithdraw{
        PermissionPermissionlessWithdraw: To_NominationPools_PermissionlessWithdraw(param),
        }
    case 3:
        return &pbgear.NominationPools_SetClaimPermissionCall_PermissionPermissionlessAll{
        PermissionPermissionlessAll: To_NominationPools_PermissionlessAll(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_SetCommissionCall(in any) *pbgear.NominationPools_SetCommissionCall {
    out := &pbgear.NominationPools_SetCommissionCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // optional field NewCommission
    if v.HasValue() {
        out.NewCommission = To_Optional_NominationPools_SetCommissionCall_new_commission(v.ValueAt(1))
    }

    return out //from message
}
func To_NominationPools_SetCommissionCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_SetCommissionCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsSetCommissionCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsSetCommissionCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_NominationPools_SetCommissionCall_new_commission(in any) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_NominationPools_SetCommissionChangeRateCall(in any) *pbgear.NominationPools_SetCommissionChangeRateCall {
    out := &pbgear.NominationPools_SetCommissionChangeRateCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // field ChangeRate To_NominationPools_PalletNominationPoolsCommissionChangeRate(w)
    out.ChangeRate = To_NominationPools_PalletNominationPoolsCommissionChangeRate(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_SetCommissionChangeRateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_SetCommissionChangeRateCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsSetCommissionChangeRateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsSetCommissionChangeRateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_NominationPools_SetCommissionMaxCall(in any) *pbgear.NominationPools_SetCommissionMaxCall {
    out := &pbgear.NominationPools_SetCommissionMaxCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // field MaxCommission To_SpArithmeticPerThingsPerbill(w)
    out.MaxCommission = To_SpArithmeticPerThingsPerbill(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_SetCommissionMaxCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_SetCommissionMaxCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsSetCommissionMaxCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsSetCommissionMaxCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_NominationPools_SetConfigsCall(in any) *pbgear.NominationPools_SetConfigsCall {
    out := &pbgear.NominationPools_SetConfigsCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MinJoinBond
    v0 := To_OneOf_NominationPools_SetConfigsCall_min_join_bond(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MinJoinBond").Set(reflect.ValueOf(v0))
    // oneOf field MinCreateBond
    v1 := To_OneOf_NominationPools_SetConfigsCall_min_create_bond(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("MinCreateBond").Set(reflect.ValueOf(v1))
    // oneOf field MaxPools
    v2 := To_OneOf_NominationPools_SetConfigsCall_max_pools(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("MaxPools").Set(reflect.ValueOf(v2))
    // oneOf field MaxMembers
    v3 := To_OneOf_NominationPools_SetConfigsCall_max_members(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("MaxMembers").Set(reflect.ValueOf(v3))
    // oneOf field MaxMembersPerPool
    v4 := To_OneOf_NominationPools_SetConfigsCall_max_members_per_pool(v.ValueAt(4))
    reflect.ValueOf(out).Elem().FieldByName("MaxMembersPerPool").Set(reflect.ValueOf(v4))
    // oneOf field GlobalMaxCommission
    v5 := To_OneOf_NominationPools_SetConfigsCall_global_max_commission(v.ValueAt(5))
    reflect.ValueOf(out).Elem().FieldByName("GlobalMaxCommission").Set(reflect.ValueOf(v5))

    return out //from message
}
func To_NominationPools_SetConfigsCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_SetConfigsCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsSetConfigsCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsSetConfigsCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_SetConfigsCall_min_join_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetConfigsCall_MinJoinBondNoop{
        MinJoinBondNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_SetConfigsCall_MinJoinBondSet{
        MinJoinBondSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_SetConfigsCall_MinJoinBondRemove{
        MinJoinBondRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_SetConfigsCall_min_create_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetConfigsCall_MinCreateBondNoop{
        MinCreateBondNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_SetConfigsCall_MinCreateBondSet{
        MinCreateBondSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_SetConfigsCall_MinCreateBondRemove{
        MinCreateBondRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_SetConfigsCall_max_pools (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetConfigsCall_MaxPoolsNoop{
        MaxPoolsNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_SetConfigsCall_MaxPoolsSet{
        MaxPoolsSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_SetConfigsCall_MaxPoolsRemove{
        MaxPoolsRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_SetConfigsCall_max_members (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetConfigsCall_MaxMembersNoop{
        MaxMembersNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_SetConfigsCall_MaxMembersSet{
        MaxMembersSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_SetConfigsCall_MaxMembersRemove{
        MaxMembersRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_SetConfigsCall_max_members_per_pool (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetConfigsCall_MaxMembersPerPoolNoop{
        MaxMembersPerPoolNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_SetConfigsCall_MaxMembersPerPoolSet{
        MaxMembersPerPoolSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_SetConfigsCall_MaxMembersPerPoolRemove{
        MaxMembersPerPoolRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_SetConfigsCall_global_max_commission (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetConfigsCall_GlobalMaxCommissionNoop{
        GlobalMaxCommissionNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_SetConfigsCall_GlobalMaxCommissionSet{
        GlobalMaxCommissionSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_SetConfigsCall_GlobalMaxCommissionRemove{
        GlobalMaxCommissionRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_SetMetadataCall(in any) *pbgear.NominationPools_SetMetadataCall {
    out := &pbgear.NominationPools_SetMetadataCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // primitive field Metadata
    out.Metadata = To_bytes(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_SetMetadataCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_SetMetadataCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsSetMetadataCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsSetMetadataCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_NominationPools_SetStateCall(in any) *pbgear.NominationPools_SetStateCall {
    out := &pbgear.NominationPools_SetStateCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // oneOf field State
    v1 := To_OneOf_NominationPools_SetStateCall_state(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("State").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_NominationPools_SetStateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_SetStateCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsSetStateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsSetStateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_SetStateCall_state (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_SetStateCall_StateOpen{
        StateOpen: To_NominationPools_Open(param),
        }
    case 1:
        return &pbgear.NominationPools_SetStateCall_StateBlocked{
        StateBlocked: To_NominationPools_Blocked(param),
        }
    case 2:
        return &pbgear.NominationPools_SetStateCall_StateDestroying{
        StateDestroying: To_NominationPools_Destroying(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_State(in any) *pbgear.NominationPools_State {
    out := &pbgear.NominationPools_State{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field State
    v0 := To_OneOf_NominationPools_State_state(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("State").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_NominationPools_State_state (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_State_StateOpen{
        StateOpen: To_NominationPools_Open(param),
        }
    case 1:
        return &pbgear.NominationPools_State_StateBlocked{
        StateBlocked: To_NominationPools_Blocked(param),
        }
    case 2:
        return &pbgear.NominationPools_State_StateDestroying{
        StateDestroying: To_NominationPools_Destroying(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_TupleNull(in any) *pbgear.NominationPools_TupleNull {
    out := &pbgear.NominationPools_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(in any) *pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
    out := &pbgear.NominationPools_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpArithmeticPerThingsPerbill(w)
    out.Value0 = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))
    // field Value1 To_SpCoreCryptoAccountId32(w)
    out.Value1 = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    


func To_NominationPools_UnbondCall(in any) *pbgear.NominationPools_UnbondCall {
    out := &pbgear.NominationPools_UnbondCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MemberAccount
    v0 := To_OneOf_NominationPools_UnbondCall_member_account(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MemberAccount").Set(reflect.ValueOf(v0))
    // primitive field UnbondingPoints
    out.UnbondingPoints = To_string(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_UnbondCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_UnbondCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsUnbondCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsUnbondCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_UnbondCall_member_account (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_UnbondCall_MemberAccountId{
        MemberAccountId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_UnbondCall_MemberAccountIndex{
        MemberAccountIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_UnbondCall_MemberAccountRaw{
        MemberAccountRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_UnbondCall_MemberAccountAddress32{
        MemberAccountAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_UnbondCall_MemberAccountAddress20{
        MemberAccountAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_UpdateRolesCall(in any) *pbgear.NominationPools_UpdateRolesCall {
    out := &pbgear.NominationPools_UpdateRolesCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field PoolId
    out.PoolId = To_uint32(v.ValueAt(0))
    // oneOf field NewRoot
    v1 := To_OneOf_NominationPools_UpdateRolesCall_new_root(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("NewRoot").Set(reflect.ValueOf(v1))
    // oneOf field NewNominator
    v2 := To_OneOf_NominationPools_UpdateRolesCall_new_nominator(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("NewNominator").Set(reflect.ValueOf(v2))
    // oneOf field NewBouncer
    v3 := To_OneOf_NominationPools_UpdateRolesCall_new_bouncer(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("NewBouncer").Set(reflect.ValueOf(v3))

    return out //from message
}
func To_NominationPools_UpdateRolesCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_UpdateRolesCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsUpdateRolesCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsUpdateRolesCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_UpdateRolesCall_new_root (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_UpdateRolesCall_NewRootNoop{
        NewRootNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_UpdateRolesCall_NewRootSet{
        NewRootSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_UpdateRolesCall_NewRootRemove{
        NewRootRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_UpdateRolesCall_new_nominator (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_UpdateRolesCall_NewNominatorNoop{
        NewNominatorNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_UpdateRolesCall_NewNominatorSet{
        NewNominatorSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_UpdateRolesCall_NewNominatorRemove{
        NewNominatorRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_NominationPools_UpdateRolesCall_new_bouncer (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_UpdateRolesCall_NewBouncerNoop{
        NewBouncerNoop: To_NominationPools_Noop(param),
        }
    case 1:
        return &pbgear.NominationPools_UpdateRolesCall_NewBouncerSet{
        NewBouncerSet: To_NominationPools_Set(param),
        }
    case 2:
        return &pbgear.NominationPools_UpdateRolesCall_NewBouncerRemove{
        NewBouncerRemove: To_NominationPools_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NominationPools_WithdrawUnbondedCall(in any) *pbgear.NominationPools_WithdrawUnbondedCall {
    out := &pbgear.NominationPools_WithdrawUnbondedCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MemberAccount
    v0 := To_OneOf_NominationPools_WithdrawUnbondedCall_member_account(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MemberAccount").Set(reflect.ValueOf(v0))
    // primitive field NumSlashingSpans
    out.NumSlashingSpans = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_NominationPools_WithdrawUnbondedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_NominationPools_WithdrawUnbondedCall(in)
    out := &pbgearextrinsic.RuntimeCall_NominationpoolsWithdrawUnbondedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("NominationpoolsWithdrawUnbondedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_NominationPools_WithdrawUnbondedCall_member_account (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountId{
        MemberAccountId: To_NominationPools_Id(param),
        }
    case 1:
        return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountIndex{
        MemberAccountIndex: To_NominationPools_Index(param),
        }
    case 2:
        return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountRaw{
        MemberAccountRaw: To_NominationPools_Raw(param),
        }
    case 3:
        return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountAddress32{
        MemberAccountAddress32: To_NominationPools_Address32(param),
        }
    case 4:
        return &pbgear.NominationPools_WithdrawUnbondedCall_MemberAccountAddress20{
        MemberAccountAddress20: To_NominationPools_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_NonTransfer(in any) *pbgear.NonTransfer {
    out := &pbgear.NonTransfer{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_None(in any) *pbgear.None {
    out := &pbgear.None{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Normal(in any) *pbgear.Normal {
    out := &pbgear.Normal{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NotExecuted(in any) *pbgear.NotExecuted {
    out := &pbgear.NotExecuted{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_NotExpendable(in any) *pbgear.NotExpendable {
    out := &pbgear.NotExpendable{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_OffencesOffenceEvent(in any) *pbgear.OffencesOffenceEvent {
    out := &pbgear.OffencesOffenceEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_OffencesPallet(in any) *pbgear.OffencesPallet {
    out := &pbgear.OffencesPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Ok(in any) *pbgear.Ok {
    out := &pbgear.Ok{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_TupleNull(w)
    out.Value0 = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_OnlyProvider(in any) *pbgear.OnlyProvider {
    out := &pbgear.OnlyProvider{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Operational(in any) *pbgear.Operational {
    out := &pbgear.Operational{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Origin(in any) *pbgear.Origin {
    out := &pbgear.Origin{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Origin
    v0 := To_OneOf_Origin_origin(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Origin").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Origin_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Origin_OriginNode{
        OriginNode: To_Node(param),
        }
    case 1:
        return &pbgear.Origin_OriginReservation{
        OriginReservation: To_Reservation(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OriginsPallet(in any) *pbgear.OriginsPallet {
    out := &pbgear.OriginsPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Other(in any) *pbgear.Other {
    out := &pbgear.Other{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_OutOfRent(in any) *pbgear.OutOfRent {
    out := &pbgear.OutOfRent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Overflow(in any) *pbgear.Overflow {
    out := &pbgear.Overflow{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Paused(in any) *pbgear.Paused {
    out := &pbgear.Paused{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_PaysFee(in any) *pbgear.PaysFee {
    out := &pbgear.PaysFee{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field PaysFee
    v0 := To_OneOf_PaysFee_pays_fee(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("PaysFee").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_PaysFee_pays_fee (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.PaysFee_PaysFeeYes{
        PaysFeeYes: To_Yes(param),
        }
    case 1:
        return &pbgear.PaysFee_PaysFeeNo{
        PaysFeeNo: To_No(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_PreimageClearedEvent(in any) *pbgear.PreimageClearedEvent {
    out := &pbgear.PreimageClearedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_PreimageNotedEvent(in any) *pbgear.PreimageNotedEvent {
    out := &pbgear.PreimageNotedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_PreimagePallet(in any) *pbgear.PreimagePallet {
    out := &pbgear.PreimagePallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_PreimagePallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_PreimagePallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.PreimagePallet_CallNotePreimageCall{
        CallNotePreimageCall: To_Preimage_NotePreimageCall(param),
        }
    case 1:
        return &pbgear.PreimagePallet_CallUnnotePreimageCall{
        CallUnnotePreimageCall: To_Preimage_UnnotePreimageCall(param),
        }
    case 2:
        return &pbgear.PreimagePallet_CallRequestPreimageCall{
        CallRequestPreimageCall: To_Preimage_RequestPreimageCall(param),
        }
    case 3:
        return &pbgear.PreimagePallet_CallUnrequestPreimageCall{
        CallUnrequestPreimageCall: To_Preimage_UnrequestPreimageCall(param),
        }
    case 4:
        return &pbgear.PreimagePallet_CallEnsureUpdatedCall{
        CallEnsureUpdatedCall: To_Preimage_EnsureUpdatedCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_PreimageRequestedEvent(in any) *pbgear.PreimageRequestedEvent {
    out := &pbgear.PreimageRequestedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Preimage_EnsureUpdatedCall(in any) *pbgear.Preimage_EnsureUpdatedCall {
    out := &pbgear.Preimage_EnsureUpdatedCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Hashes
    out.Hashes = To_Repeated_Preimage_EnsureUpdatedCall_hashes(v.ValueAt(0))

    return out //from message
}
func To_Preimage_EnsureUpdatedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Preimage_EnsureUpdatedCall(in)
    out := &pbgearextrinsic.RuntimeCall_PreimageEnsureUpdatedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("PreimageEnsureUpdatedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Preimage_EnsureUpdatedCall_hashes(in any) []*pbgear.PrimitiveTypesH256 {
    items := in.([]interface{})

    var out []*pbgear.PrimitiveTypesH256
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_PrimitiveTypesH256(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Preimage_NotePreimageCall(in any) *pbgear.Preimage_NotePreimageCall {
    out := &pbgear.Preimage_NotePreimageCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Bytes
    out.Bytes = To_bytes(v.ValueAt(0))

    return out //from message
}
func To_Preimage_NotePreimageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Preimage_NotePreimageCall(in)
    out := &pbgearextrinsic.RuntimeCall_PreimageNotePreimageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("PreimageNotePreimageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Preimage_RequestPreimageCall(in any) *pbgear.Preimage_RequestPreimageCall {
    out := &pbgear.Preimage_RequestPreimageCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
func To_Preimage_RequestPreimageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Preimage_RequestPreimageCall(in)
    out := &pbgearextrinsic.RuntimeCall_PreimageRequestPreimageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("PreimageRequestPreimageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Preimage_UnnotePreimageCall(in any) *pbgear.Preimage_UnnotePreimageCall {
    out := &pbgear.Preimage_UnnotePreimageCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
func To_Preimage_UnnotePreimageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Preimage_UnnotePreimageCall(in)
    out := &pbgearextrinsic.RuntimeCall_PreimageUnnotePreimageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("PreimageUnnotePreimageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Preimage_UnrequestPreimageCall(in any) *pbgear.Preimage_UnrequestPreimageCall {
    out := &pbgear.Preimage_UnrequestPreimageCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
func To_Preimage_UnrequestPreimageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Preimage_UnrequestPreimageCall(in)
    out := &pbgearextrinsic.RuntimeCall_PreimageUnrequestPreimageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("PreimageUnrequestPreimageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_PrimitiveTypesH256(in any) *pbgear.PrimitiveTypesH256 {
    out := &pbgear.PrimitiveTypesH256{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_ProgramGotInitialized(in any) *pbgear.ProgramGotInitialized {
    out := &pbgear.ProgramGotInitialized{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ProgramSet(in any) *pbgear.ProgramSet {
    out := &pbgear.ProgramSet{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Expiration
    out.Expiration = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Proposal(in any) *pbgear.Proposal {
    out := &pbgear.Proposal{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Proposal
    v0 := To_OneOf_Proposal_proposal(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Proposal_proposal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proposal_ProposalLegacy{
        ProposalLegacy: To_Legacy(param),
        }
    case 1:
        return &pbgear.Proposal_ProposalInline{
        ProposalInline: To_Inline(param),
        }
    case 2:
        return &pbgear.Proposal_ProposalLookup{
        ProposalLookup: To_Lookup(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ProxyAnnouncedEvent(in any) *pbgear.ProxyAnnouncedEvent {
    out := &pbgear.ProxyAnnouncedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ProxyPallet(in any) *pbgear.ProxyPallet {
    out := &pbgear.ProxyPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_ProxyPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ProxyPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ProxyPallet_CallProxyCall{
        CallProxyCall: To_Proxy_ProxyCall(param),
        }
    case 1:
        return &pbgear.ProxyPallet_CallAddProxyCall{
        CallAddProxyCall: To_Proxy_AddProxyCall(param),
        }
    case 2:
        return &pbgear.ProxyPallet_CallRemoveProxyCall{
        CallRemoveProxyCall: To_Proxy_RemoveProxyCall(param),
        }
    case 3:
        return &pbgear.ProxyPallet_CallRemoveProxiesCall{
        CallRemoveProxiesCall: To_Proxy_RemoveProxiesCall(param),
        }
    case 4:
        return &pbgear.ProxyPallet_CallCreatePureCall{
        CallCreatePureCall: To_Proxy_CreatePureCall(param),
        }
    case 5:
        return &pbgear.ProxyPallet_CallKillPureCall{
        CallKillPureCall: To_Proxy_KillPureCall(param),
        }
    case 6:
        return &pbgear.ProxyPallet_CallAnnounceCall{
        CallAnnounceCall: To_Proxy_AnnounceCall(param),
        }
    case 7:
        return &pbgear.ProxyPallet_CallRemoveAnnouncementCall{
        CallRemoveAnnouncementCall: To_Proxy_RemoveAnnouncementCall(param),
        }
    case 8:
        return &pbgear.ProxyPallet_CallRejectAnnouncementCall{
        CallRejectAnnouncementCall: To_Proxy_RejectAnnouncementCall(param),
        }
    case 9:
        return &pbgear.ProxyPallet_CallProxyAnnouncedCall{
        CallProxyAnnouncedCall: To_Proxy_ProxyAnnouncedCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ProxyProxyAddedEvent(in any) *pbgear.ProxyProxyAddedEvent {
    out := &pbgear.ProxyProxyAddedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ProxyProxyExecutedEvent(in any) *pbgear.ProxyProxyExecutedEvent {
    out := &pbgear.ProxyProxyExecutedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ProxyProxyRemovedEvent(in any) *pbgear.ProxyProxyRemovedEvent {
    out := &pbgear.ProxyProxyRemovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ProxyPureCreatedEvent(in any) *pbgear.ProxyPureCreatedEvent {
    out := &pbgear.ProxyPureCreatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ProxyType(in any) *pbgear.ProxyType {
    out := &pbgear.ProxyType{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ProxyType
    v0 := To_OneOf_ProxyType_proxy_type(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ProxyType_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ProxyType_ProxyTypeAny{
        ProxyTypeAny: To_Any(param),
        }
    case 1:
        return &pbgear.ProxyType_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_NonTransfer(param),
        }
    case 2:
        return &pbgear.ProxyType_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Governance(param),
        }
    case 3:
        return &pbgear.ProxyType_ProxyTypeStaking{
        ProxyTypeStaking: To_Staking(param),
        }
    case 4:
        return &pbgear.ProxyType_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.ProxyType_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_AddProxyCall(in any) *pbgear.Proxy_AddProxyCall {
    out := &pbgear.Proxy_AddProxyCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Delegate
    v0 := To_OneOf_Proxy_AddProxyCall_delegate(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
    // oneOf field ProxyType
    v1 := To_OneOf_Proxy_AddProxyCall_proxy_type(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v1))
    // primitive field Delay
    out.Delay = To_uint32(v.ValueAt(2))

    return out //from message
}
func To_Proxy_AddProxyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_AddProxyCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyAddProxyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyAddProxyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_AddProxyCall_delegate (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_AddProxyCall_DelegateId{
        DelegateId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_AddProxyCall_DelegateIndex{
        DelegateIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_AddProxyCall_DelegateRaw{
        DelegateRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_AddProxyCall_DelegateAddress32{
        DelegateAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_AddProxyCall_DelegateAddress20{
        DelegateAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_AddProxyCall_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_AddProxyCall_ProxyTypeAny{
        ProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_AddProxyCall_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_AddProxyCall_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_AddProxyCall_ProxyTypeStaking{
        ProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_AddProxyCall_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_AddProxyCall_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_Address20(in any) *pbgear.Proxy_Address20 {
    out := &pbgear.Proxy_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Proxy_Address32(in any) *pbgear.Proxy_Address32 {
    out := &pbgear.Proxy_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Proxy_AnnounceCall(in any) *pbgear.Proxy_AnnounceCall {
    out := &pbgear.Proxy_AnnounceCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Real
    v0 := To_OneOf_Proxy_AnnounceCall_real(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))
    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(1))

    return out //from message
}
func To_Proxy_AnnounceCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_AnnounceCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyAnnounceCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyAnnounceCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_AnnounceCall_real (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_AnnounceCall_RealId{
        RealId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_AnnounceCall_RealIndex{
        RealIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_AnnounceCall_RealRaw{
        RealRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_AnnounceCall_RealAddress32{
        RealAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_AnnounceCall_RealAddress20{
        RealAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Proxy_Any(in any) *pbgear.Proxy_Any {
    out := &pbgear.Proxy_Any{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Proxy_CancelProxy(in any) *pbgear.Proxy_CancelProxy {
    out := &pbgear.Proxy_CancelProxy{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Proxy_CreatePureCall(in any) *pbgear.Proxy_CreatePureCall {
    out := &pbgear.Proxy_CreatePureCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ProxyType
    v0 := To_OneOf_Proxy_CreatePureCall_proxy_type(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v0))
    // primitive field Delay
    out.Delay = To_uint32(v.ValueAt(1))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(2))

    return out //from message
}
func To_Proxy_CreatePureCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_CreatePureCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyCreatePureCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyCreatePureCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_CreatePureCall_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_CreatePureCall_ProxyTypeAny{
        ProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_CreatePureCall_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_CreatePureCall_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_CreatePureCall_ProxyTypeStaking{
        ProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_CreatePureCall_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_CreatePureCall_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_Delegate(in any) *pbgear.Proxy_Delegate {
    out := &pbgear.Proxy_Delegate{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Delegate
    v0 := To_OneOf_Proxy_Delegate_delegate(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Proxy_Delegate_delegate (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_Delegate_DelegateId{
        DelegateId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_Delegate_DelegateIndex{
        DelegateIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_Delegate_DelegateRaw{
        DelegateRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_Delegate_DelegateAddress32{
        DelegateAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_Delegate_DelegateAddress20{
        DelegateAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_ForceProxyType(in any) *pbgear.Proxy_ForceProxyType {
    out := &pbgear.Proxy_ForceProxyType{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ForceProxyType
    v0 := To_OneOf_Proxy_ForceProxyType_force_proxy_type(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ForceProxyType").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Proxy_ForceProxyType_force_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ForceProxyType_ForceProxyTypeAny{
        ForceProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_ForceProxyType_ForceProxyTypeNonTransfer{
        ForceProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_ForceProxyType_ForceProxyTypeGovernance{
        ForceProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_ForceProxyType_ForceProxyTypeStaking{
        ForceProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_ForceProxyType_ForceProxyTypeIdentityJudgement{
        ForceProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_ForceProxyType_ForceProxyTypeCancelProxy{
        ForceProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_Governance(in any) *pbgear.Proxy_Governance {
    out := &pbgear.Proxy_Governance{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Proxy_Id(in any) *pbgear.Proxy_Id {
    out := &pbgear.Proxy_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Proxy_IdentityJudgement(in any) *pbgear.Proxy_IdentityJudgement {
    out := &pbgear.Proxy_IdentityJudgement{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Proxy_Index(in any) *pbgear.Proxy_Index {
    out := &pbgear.Proxy_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Proxy_TupleNull(w)
    out.Value0 = To_Proxy_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Proxy_KillPureCall(in any) *pbgear.Proxy_KillPureCall {
    out := &pbgear.Proxy_KillPureCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Spawner
    v0 := To_OneOf_Proxy_KillPureCall_spawner(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Spawner").Set(reflect.ValueOf(v0))
    // oneOf field ProxyType
    v1 := To_OneOf_Proxy_KillPureCall_proxy_type(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v1))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(2))
    // primitive field Height
    out.Height = To_uint32(v.ValueAt(3))
    // primitive field ExtIndex
    out.ExtIndex = To_uint32(v.ValueAt(4))

    return out //from message
}
func To_Proxy_KillPureCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_KillPureCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyKillPureCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyKillPureCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_KillPureCall_spawner (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_KillPureCall_SpawnerId{
        SpawnerId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_KillPureCall_SpawnerIndex{
        SpawnerIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_KillPureCall_SpawnerRaw{
        SpawnerRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_KillPureCall_SpawnerAddress32{
        SpawnerAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_KillPureCall_SpawnerAddress20{
        SpawnerAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_KillPureCall_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_KillPureCall_ProxyTypeAny{
        ProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_KillPureCall_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_KillPureCall_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_KillPureCall_ProxyTypeStaking{
        ProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_KillPureCall_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_KillPureCall_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_NonTransfer(in any) *pbgear.Proxy_NonTransfer {
    out := &pbgear.Proxy_NonTransfer{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Proxy_ProxyAnnouncedCall(in any) *pbgear.Proxy_ProxyAnnouncedCall {
    out := &pbgear.Proxy_ProxyAnnouncedCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Delegate
    v0 := To_OneOf_Proxy_ProxyAnnouncedCall_delegate(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
    // oneOf field Real
    v1 := To_OneOf_Proxy_ProxyAnnouncedCall_real(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v1))
    // oneOf field ForceProxyType
    v2 := To_OneOf_Proxy_ProxyAnnouncedCall_force_proxy_type(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("ForceProxyType").Set(reflect.ValueOf(v2))
    // oneOf field Call
    v3 := To_OneOf_Proxy_ProxyAnnouncedCall_call(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))

    return out //from message
}
func To_Proxy_ProxyAnnouncedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_ProxyAnnouncedCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyProxyAnnouncedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyProxyAnnouncedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_ProxyAnnouncedCall_delegate (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyAnnouncedCall_DelegateId{
        DelegateId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyAnnouncedCall_DelegateIndex{
        DelegateIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_ProxyAnnouncedCall_DelegateRaw{
        DelegateRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyAnnouncedCall_DelegateAddress32{
        DelegateAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyAnnouncedCall_DelegateAddress20{
        DelegateAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_ProxyAnnouncedCall_real (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyAnnouncedCall_RealId{
        RealId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyAnnouncedCall_RealIndex{
        RealIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_ProxyAnnouncedCall_RealRaw{
        RealRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyAnnouncedCall_RealAddress32{
        RealAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyAnnouncedCall_RealAddress20{
        RealAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_ProxyAnnouncedCall_force_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeAny{
        ForceProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeNonTransfer{
        ForceProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeGovernance{
        ForceProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeStaking{
        ForceProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeIdentityJudgement{
        ForceProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_ProxyAnnouncedCall_ForceProxyTypeCancelProxy{
        ForceProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_ProxyAnnouncedCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Proxy_ProxyAnnouncedCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_ProxyCall(in any) *pbgear.Proxy_ProxyCall {
    out := &pbgear.Proxy_ProxyCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Real
    v0 := To_OneOf_Proxy_ProxyCall_real(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))
    // oneOf field ForceProxyType
    v1 := To_OneOf_Proxy_ProxyCall_force_proxy_type(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("ForceProxyType").Set(reflect.ValueOf(v1))
    // oneOf field Call
    v2 := To_OneOf_Proxy_ProxyCall_call(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v2))

    return out //from message
}
func To_Proxy_ProxyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_ProxyCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyProxyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyProxyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_ProxyCall_real (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyCall_RealId{
        RealId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyCall_RealIndex{
        RealIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_ProxyCall_RealRaw{
        RealRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyCall_RealAddress32{
        RealAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyCall_RealAddress20{
        RealAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_ProxyCall_force_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyCall_ForceProxyTypeAny{
        ForceProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyCall_ForceProxyTypeNonTransfer{
        ForceProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_ProxyCall_ForceProxyTypeGovernance{
        ForceProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyCall_ForceProxyTypeStaking{
        ForceProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyCall_ForceProxyTypeIdentityJudgement{
        ForceProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_ProxyCall_ForceProxyTypeCancelProxy{
        ForceProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_ProxyCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Proxy_ProxyCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Proxy_ProxyCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Proxy_ProxyCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Proxy_ProxyCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Proxy_ProxyCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Proxy_ProxyCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Proxy_ProxyCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Proxy_ProxyCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Proxy_ProxyCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Proxy_ProxyCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Proxy_ProxyCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Proxy_ProxyCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Proxy_ProxyCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Proxy_ProxyCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Proxy_ProxyCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Proxy_ProxyCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Proxy_ProxyCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Proxy_ProxyCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Proxy_ProxyCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Proxy_ProxyCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Proxy_ProxyCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Proxy_ProxyCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Proxy_ProxyCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Proxy_ProxyCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Proxy_ProxyCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_ProxyType(in any) *pbgear.Proxy_ProxyType {
    out := &pbgear.Proxy_ProxyType{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ProxyType
    v0 := To_OneOf_Proxy_ProxyType_proxy_type(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Proxy_ProxyType_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_ProxyType_ProxyTypeAny{
        ProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_ProxyType_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_ProxyType_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_ProxyType_ProxyTypeStaking{
        ProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_ProxyType_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_ProxyType_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_Raw(in any) *pbgear.Proxy_Raw {
    out := &pbgear.Proxy_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Proxy_Real(in any) *pbgear.Proxy_Real {
    out := &pbgear.Proxy_Real{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Real
    v0 := To_OneOf_Proxy_Real_real(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Proxy_Real_real (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_Real_RealId{
        RealId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_Real_RealIndex{
        RealIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_Real_RealRaw{
        RealRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_Real_RealAddress32{
        RealAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_Real_RealAddress20{
        RealAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_RejectAnnouncementCall(in any) *pbgear.Proxy_RejectAnnouncementCall {
    out := &pbgear.Proxy_RejectAnnouncementCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Delegate
    v0 := To_OneOf_Proxy_RejectAnnouncementCall_delegate(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(1))

    return out //from message
}
func To_Proxy_RejectAnnouncementCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_RejectAnnouncementCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyRejectAnnouncementCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyRejectAnnouncementCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_RejectAnnouncementCall_delegate (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_RejectAnnouncementCall_DelegateId{
        DelegateId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_RejectAnnouncementCall_DelegateIndex{
        DelegateIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_RejectAnnouncementCall_DelegateRaw{
        DelegateRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_RejectAnnouncementCall_DelegateAddress32{
        DelegateAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_RejectAnnouncementCall_DelegateAddress20{
        DelegateAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Proxy_RemoveAnnouncementCall(in any) *pbgear.Proxy_RemoveAnnouncementCall {
    out := &pbgear.Proxy_RemoveAnnouncementCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Real
    v0 := To_OneOf_Proxy_RemoveAnnouncementCall_real(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Real").Set(reflect.ValueOf(v0))
    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(1))

    return out //from message
}
func To_Proxy_RemoveAnnouncementCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_RemoveAnnouncementCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyRemoveAnnouncementCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyRemoveAnnouncementCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_RemoveAnnouncementCall_real (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_RemoveAnnouncementCall_RealId{
        RealId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_RemoveAnnouncementCall_RealIndex{
        RealIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_RemoveAnnouncementCall_RealRaw{
        RealRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_RemoveAnnouncementCall_RealAddress32{
        RealAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_RemoveAnnouncementCall_RealAddress20{
        RealAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Proxy_RemoveProxiesCall(in any) *pbgear.Proxy_RemoveProxiesCall {
    out := &pbgear.Proxy_RemoveProxiesCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Proxy_RemoveProxyCall(in any) *pbgear.Proxy_RemoveProxyCall {
    out := &pbgear.Proxy_RemoveProxyCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Delegate
    v0 := To_OneOf_Proxy_RemoveProxyCall_delegate(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Delegate").Set(reflect.ValueOf(v0))
    // oneOf field ProxyType
    v1 := To_OneOf_Proxy_RemoveProxyCall_proxy_type(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("ProxyType").Set(reflect.ValueOf(v1))
    // primitive field Delay
    out.Delay = To_uint32(v.ValueAt(2))

    return out //from message
}
func To_Proxy_RemoveProxyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Proxy_RemoveProxyCall(in)
    out := &pbgearextrinsic.RuntimeCall_ProxyRemoveProxyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ProxyRemoveProxyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Proxy_RemoveProxyCall_delegate (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_RemoveProxyCall_DelegateId{
        DelegateId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_RemoveProxyCall_DelegateIndex{
        DelegateIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_RemoveProxyCall_DelegateRaw{
        DelegateRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_RemoveProxyCall_DelegateAddress32{
        DelegateAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_RemoveProxyCall_DelegateAddress20{
        DelegateAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Proxy_RemoveProxyCall_proxy_type (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_RemoveProxyCall_ProxyTypeAny{
        ProxyTypeAny: To_Proxy_Any(param),
        }
    case 1:
        return &pbgear.Proxy_RemoveProxyCall_ProxyTypeNonTransfer{
        ProxyTypeNonTransfer: To_Proxy_NonTransfer(param),
        }
    case 2:
        return &pbgear.Proxy_RemoveProxyCall_ProxyTypeGovernance{
        ProxyTypeGovernance: To_Proxy_Governance(param),
        }
    case 3:
        return &pbgear.Proxy_RemoveProxyCall_ProxyTypeStaking{
        ProxyTypeStaking: To_Proxy_Staking(param),
        }
    case 4:
        return &pbgear.Proxy_RemoveProxyCall_ProxyTypeIdentityJudgement{
        ProxyTypeIdentityJudgement: To_Proxy_IdentityJudgement(param),
        }
    case 5:
        return &pbgear.Proxy_RemoveProxyCall_ProxyTypeCancelProxy{
        ProxyTypeCancelProxy: To_Proxy_CancelProxy(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_Spawner(in any) *pbgear.Proxy_Spawner {
    out := &pbgear.Proxy_Spawner{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Spawner
    v0 := To_OneOf_Proxy_Spawner_spawner(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Spawner").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Proxy_Spawner_spawner (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Proxy_Spawner_SpawnerId{
        SpawnerId: To_Proxy_Id(param),
        }
    case 1:
        return &pbgear.Proxy_Spawner_SpawnerIndex{
        SpawnerIndex: To_Proxy_Index(param),
        }
    case 2:
        return &pbgear.Proxy_Spawner_SpawnerRaw{
        SpawnerRaw: To_Proxy_Raw(param),
        }
    case 3:
        return &pbgear.Proxy_Spawner_SpawnerAddress32{
        SpawnerAddress32: To_Proxy_Address32(param),
        }
    case 4:
        return &pbgear.Proxy_Spawner_SpawnerAddress20{
        SpawnerAddress20: To_Proxy_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Proxy_Staking(in any) *pbgear.Proxy_Staking {
    out := &pbgear.Proxy_Staking{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Proxy_TupleNull(in any) *pbgear.Proxy_TupleNull {
    out := &pbgear.Proxy_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_RanOutOfGas(in any) *pbgear.RanOutOfGas {
    out := &pbgear.RanOutOfGas{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_RankedCollective_Aye(in any) *pbgear.RankedCollective_Aye {
    out := &pbgear.RankedCollective_Aye{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_RankedCollective_Nay(in any) *pbgear.RankedCollective_Nay {
    out := &pbgear.RankedCollective_Nay{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_RankedCollective_PalletRankedCollectiveTally(in any) *pbgear.RankedCollective_PalletRankedCollectiveTally {
    out := &pbgear.RankedCollective_PalletRankedCollectiveTally{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field BareAyes
    out.BareAyes = To_uint32(v.ValueAt(0))
    // primitive field Ayes
    out.Ayes = To_uint32(v.ValueAt(1))
    // primitive field Nays
    out.Nays = To_uint32(v.ValueAt(2))

    return out //from message
}
    
func To_RankedCollective_Vote(in any) *pbgear.RankedCollective_Vote {
    out := &pbgear.RankedCollective_Vote{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Vote
    v0 := To_OneOf_RankedCollective_Vote_vote(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Vote").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_RankedCollective_Vote_vote (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.RankedCollective_Vote_VoteAye{
        VoteAye: To_RankedCollective_Aye(param),
        }
    case 1:
        return &pbgear.RankedCollective_Vote_VoteNay{
        VoteNay: To_RankedCollective_Nay(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Reason(in any) *pbgear.Reason {
    out := &pbgear.Reason{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Reason
    v0 := To_OneOf_Reason_reason(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Reason").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Reason_reason (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Reason_ReasonRuntime{
        ReasonRuntime: To_Runtime(param),
        }
    case 1:
        return &pbgear.Reason_ReasonSystem{
        ReasonSystem: To_System(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ReferendaApprovedEvent(in any) *pbgear.ReferendaApprovedEvent {
    out := &pbgear.ReferendaApprovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaCancelledEvent(in any) *pbgear.ReferendaCancelledEvent {
    out := &pbgear.ReferendaCancelledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaConfirmAbortedEvent(in any) *pbgear.ReferendaConfirmAbortedEvent {
    out := &pbgear.ReferendaConfirmAbortedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaConfirmStartedEvent(in any) *pbgear.ReferendaConfirmStartedEvent {
    out := &pbgear.ReferendaConfirmStartedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaConfirmedEvent(in any) *pbgear.ReferendaConfirmedEvent {
    out := &pbgear.ReferendaConfirmedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaDecisionDepositPlacedEvent(in any) *pbgear.ReferendaDecisionDepositPlacedEvent {
    out := &pbgear.ReferendaDecisionDepositPlacedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaDecisionDepositRefundedEvent(in any) *pbgear.ReferendaDecisionDepositRefundedEvent {
    out := &pbgear.ReferendaDecisionDepositRefundedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaDecisionStartedEvent(in any) *pbgear.ReferendaDecisionStartedEvent {
    out := &pbgear.ReferendaDecisionStartedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaDepositSlashedEvent(in any) *pbgear.ReferendaDepositSlashedEvent {
    out := &pbgear.ReferendaDepositSlashedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaKilledEvent(in any) *pbgear.ReferendaKilledEvent {
    out := &pbgear.ReferendaKilledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaMetadataClearedEvent(in any) *pbgear.ReferendaMetadataClearedEvent {
    out := &pbgear.ReferendaMetadataClearedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaMetadataSetEvent(in any) *pbgear.ReferendaMetadataSetEvent {
    out := &pbgear.ReferendaMetadataSetEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaPallet(in any) *pbgear.ReferendaPallet {
    out := &pbgear.ReferendaPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_ReferendaPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_ReferendaPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.ReferendaPallet_CallSubmitCall{
        CallSubmitCall: To_Referenda_SubmitCall(param),
        }
    case 1:
        return &pbgear.ReferendaPallet_CallPlaceDecisionDepositCall{
        CallPlaceDecisionDepositCall: To_Referenda_PlaceDecisionDepositCall(param),
        }
    case 2:
        return &pbgear.ReferendaPallet_CallRefundDecisionDepositCall{
        CallRefundDecisionDepositCall: To_Referenda_RefundDecisionDepositCall(param),
        }
    case 3:
        return &pbgear.ReferendaPallet_CallCancelCall{
        CallCancelCall: To_Referenda_CancelCall(param),
        }
    case 4:
        return &pbgear.ReferendaPallet_CallKillCall{
        CallKillCall: To_Referenda_KillCall(param),
        }
    case 5:
        return &pbgear.ReferendaPallet_CallNudgeReferendumCall{
        CallNudgeReferendumCall: To_Referenda_NudgeReferendumCall(param),
        }
    case 6:
        return &pbgear.ReferendaPallet_CallOneFewerDecidingCall{
        CallOneFewerDecidingCall: To_Referenda_OneFewerDecidingCall(param),
        }
    case 7:
        return &pbgear.ReferendaPallet_CallRefundSubmissionDepositCall{
        CallRefundSubmissionDepositCall: To_Referenda_RefundSubmissionDepositCall(param),
        }
    case 8:
        return &pbgear.ReferendaPallet_CallSetMetadataCall{
        CallSetMetadataCall: To_Referenda_SetMetadataCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_ReferendaRejectedEvent(in any) *pbgear.ReferendaRejectedEvent {
    out := &pbgear.ReferendaRejectedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaSubmissionDepositRefundedEvent(in any) *pbgear.ReferendaSubmissionDepositRefundedEvent {
    out := &pbgear.ReferendaSubmissionDepositRefundedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaSubmittedEvent(in any) *pbgear.ReferendaSubmittedEvent {
    out := &pbgear.ReferendaSubmittedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendaTimedOutEvent(in any) *pbgear.ReferendaTimedOutEvent {
    out := &pbgear.ReferendaTimedOutEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Referenda_After(in any) *pbgear.Referenda_After {
    out := &pbgear.Referenda_After{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Referenda_At(in any) *pbgear.Referenda_At {
    out := &pbgear.Referenda_At{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_Referenda_CancelCall(in any) *pbgear.Referenda_CancelCall {
    out := &pbgear.Referenda_CancelCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Referenda_CancelCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_CancelCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaCancelCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaCancelCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Referenda_EnactmentMoment(in any) *pbgear.Referenda_EnactmentMoment {
    out := &pbgear.Referenda_EnactmentMoment{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field EnactmentMoment
    v0 := To_OneOf_Referenda_EnactmentMoment_enactment_moment(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Referenda_EnactmentMoment_enactment_moment (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_EnactmentMoment_EnactmentMomentAt{
        EnactmentMomentAt: To_Referenda_At(param),
        }
    case 1:
        return &pbgear.Referenda_EnactmentMoment_EnactmentMomentAfter{
        EnactmentMomentAfter: To_Referenda_After(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Referenda_Inline(in any) *pbgear.Referenda_Inline {
    out := &pbgear.Referenda_Inline{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_BoundedCollectionsBoundedVecBoundedVec(w)
    out.Value0 = To_BoundedCollectionsBoundedVecBoundedVec(v.ValueAt(0))

    return out //from message
}
    

func To_Referenda_KillCall(in any) *pbgear.Referenda_KillCall {
    out := &pbgear.Referenda_KillCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Referenda_KillCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_KillCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaKillCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaKillCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Referenda_Legacy(in any) *pbgear.Referenda_Legacy {
    out := &pbgear.Referenda_Legacy{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
    

func To_Referenda_Lookup(in any) *pbgear.Referenda_Lookup {
    out := &pbgear.Referenda_Lookup{}
    v := in.(registry.Valuable)
    _ = v

    // field Hash To_PrimitiveTypesH256(w)
    out.Hash = To_PrimitiveTypesH256(v.ValueAt(0))
    // primitive field Len
    out.Len = To_uint32(v.ValueAt(1))

    return out //from message
}
    

func To_Referenda_NudgeReferendumCall(in any) *pbgear.Referenda_NudgeReferendumCall {
    out := &pbgear.Referenda_NudgeReferendumCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Referenda_NudgeReferendumCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_NudgeReferendumCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaNudgeReferendumCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaNudgeReferendumCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Referenda_OneFewerDecidingCall(in any) *pbgear.Referenda_OneFewerDecidingCall {
    out := &pbgear.Referenda_OneFewerDecidingCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Track
    out.Track = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Referenda_OneFewerDecidingCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_OneFewerDecidingCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaOneFewerDecidingCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaOneFewerDecidingCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Referenda_Origins(in any) *pbgear.Referenda_Origins {
    out := &pbgear.Referenda_Origins{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Referenda_Origins_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Referenda_Origins_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_Origins_Value0StakingAdmin{
        Value0StakingAdmin: To_StakingAdmin(param),
        }
    case 1:
        return &pbgear.Referenda_Origins_Value0Treasurer{
        Value0Treasurer: To_Treasurer(param),
        }
    case 2:
        return &pbgear.Referenda_Origins_Value0FellowshipAdmin{
        Value0FellowshipAdmin: To_FellowshipAdmin(param),
        }
    case 3:
        return &pbgear.Referenda_Origins_Value0GeneralAdmin{
        Value0GeneralAdmin: To_GeneralAdmin(param),
        }
    case 4:
        return &pbgear.Referenda_Origins_Value0ReferendumCanceller{
        Value0ReferendumCanceller: To_ReferendumCanceller(param),
        }
    case 5:
        return &pbgear.Referenda_Origins_Value0ReferendumKiller{
        Value0ReferendumKiller: To_ReferendumKiller(param),
        }
    case 6:
        return &pbgear.Referenda_Origins_Value0SmallTipper{
        Value0SmallTipper: To_SmallTipper(param),
        }
    case 7:
        return &pbgear.Referenda_Origins_Value0BigTipper{
        Value0BigTipper: To_BigTipper(param),
        }
    case 8:
        return &pbgear.Referenda_Origins_Value0SmallSpender{
        Value0SmallSpender: To_SmallSpender(param),
        }
    case 9:
        return &pbgear.Referenda_Origins_Value0MediumSpender{
        Value0MediumSpender: To_MediumSpender(param),
        }
    case 10:
        return &pbgear.Referenda_Origins_Value0BigSpender{
        Value0BigSpender: To_BigSpender(param),
        }
    case 11:
        return &pbgear.Referenda_Origins_Value0WhitelistedCaller{
        Value0WhitelistedCaller: To_WhitelistedCaller(param),
        }
    case 12:
        return &pbgear.Referenda_Origins_Value0FellowshipInitiates{
        Value0FellowshipInitiates: To_FellowshipInitiates(param),
        }
    case 13:
        return &pbgear.Referenda_Origins_Value0Fellows{
        Value0Fellows: To_Fellows(param),
        }
    case 14:
        return &pbgear.Referenda_Origins_Value0FellowshipExperts{
        Value0FellowshipExperts: To_FellowshipExperts(param),
        }
    case 15:
        return &pbgear.Referenda_Origins_Value0FellowshipMasters{
        Value0FellowshipMasters: To_FellowshipMasters(param),
        }
    case 16:
        return &pbgear.Referenda_Origins_Value0Fellowship1Dan{
        Value0Fellowship1Dan: To_Fellowship1Dan(param),
        }
    case 17:
        return &pbgear.Referenda_Origins_Value0Fellowship2Dan{
        Value0Fellowship2Dan: To_Fellowship2Dan(param),
        }
    case 18:
        return &pbgear.Referenda_Origins_Value0Fellowship3Dan{
        Value0Fellowship3Dan: To_Fellowship3Dan(param),
        }
    case 19:
        return &pbgear.Referenda_Origins_Value0Fellowship4Dan{
        Value0Fellowship4Dan: To_Fellowship4Dan(param),
        }
    case 20:
        return &pbgear.Referenda_Origins_Value0Fellowship5Dan{
        Value0Fellowship5Dan: To_Fellowship5Dan(param),
        }
    case 21:
        return &pbgear.Referenda_Origins_Value0Fellowship6Dan{
        Value0Fellowship6Dan: To_Fellowship6Dan(param),
        }
    case 22:
        return &pbgear.Referenda_Origins_Value0Fellowship7Dan{
        Value0Fellowship7Dan: To_Fellowship7Dan(param),
        }
    case 23:
        return &pbgear.Referenda_Origins_Value0Fellowship8Dan{
        Value0Fellowship8Dan: To_Fellowship8Dan(param),
        }
    case 24:
        return &pbgear.Referenda_Origins_Value0Fellowship9Dan{
        Value0Fellowship9Dan: To_Fellowship9Dan(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Referenda_PlaceDecisionDepositCall(in any) *pbgear.Referenda_PlaceDecisionDepositCall {
    out := &pbgear.Referenda_PlaceDecisionDepositCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Referenda_PlaceDecisionDepositCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_PlaceDecisionDepositCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaPlaceDecisionDepositCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaPlaceDecisionDepositCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Referenda_Proposal(in any) *pbgear.Referenda_Proposal {
    out := &pbgear.Referenda_Proposal{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Proposal
    v0 := To_OneOf_Referenda_Proposal_proposal(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Referenda_Proposal_proposal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_Proposal_ProposalLegacy{
        ProposalLegacy: To_Referenda_Legacy(param),
        }
    case 1:
        return &pbgear.Referenda_Proposal_ProposalInline{
        ProposalInline: To_Referenda_Inline(param),
        }
    case 2:
        return &pbgear.Referenda_Proposal_ProposalLookup{
        ProposalLookup: To_Referenda_Lookup(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Referenda_ProposalOrigin(in any) *pbgear.Referenda_ProposalOrigin {
    out := &pbgear.Referenda_ProposalOrigin{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ProposalOrigin
    v0 := To_OneOf_Referenda_ProposalOrigin_proposal_origin(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Referenda_ProposalOrigin_proposal_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_ProposalOrigin_ProposalOriginSystem{
        ProposalOriginSystem: To_Referenda_System(param),
        }
    case 20:
        return &pbgear.Referenda_ProposalOrigin_ProposalOriginOrigins{
        ProposalOriginOrigins: To_Referenda_Origins(param),
        }
    case 2:
        return &pbgear.Referenda_ProposalOrigin_ProposalOriginVoid{
        ProposalOriginVoid: To_Referenda_Void(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Referenda_RefundDecisionDepositCall(in any) *pbgear.Referenda_RefundDecisionDepositCall {
    out := &pbgear.Referenda_RefundDecisionDepositCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Referenda_RefundDecisionDepositCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_RefundDecisionDepositCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaRefundDecisionDepositCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaRefundDecisionDepositCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Referenda_RefundSubmissionDepositCall(in any) *pbgear.Referenda_RefundSubmissionDepositCall {
    out := &pbgear.Referenda_RefundSubmissionDepositCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Referenda_RefundSubmissionDepositCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_RefundSubmissionDepositCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaRefundSubmissionDepositCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaRefundSubmissionDepositCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Referenda_SetMetadataCall(in any) *pbgear.Referenda_SetMetadataCall {
    out := &pbgear.Referenda_SetMetadataCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // optional field MaybeHash
    if v.HasValue() {
        out.MaybeHash = To_Optional_Referenda_SetMetadataCall_maybe_hash(v.ValueAt(1))
    }

    return out //from message
}
func To_Referenda_SetMetadataCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_SetMetadataCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaSetMetadataCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaSetMetadataCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_Referenda_SetMetadataCall_maybe_hash(in any) *pbgear.PrimitiveTypesH256 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_Referenda_SubmitCall(in any) *pbgear.Referenda_SubmitCall {
    out := &pbgear.Referenda_SubmitCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ProposalOrigin
    v0 := To_OneOf_Referenda_SubmitCall_proposal_origin(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ProposalOrigin").Set(reflect.ValueOf(v0))
    // oneOf field Proposal
    v1 := To_OneOf_Referenda_SubmitCall_proposal(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Proposal").Set(reflect.ValueOf(v1))
    // oneOf field EnactmentMoment
    v2 := To_OneOf_Referenda_SubmitCall_enactment_moment(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("EnactmentMoment").Set(reflect.ValueOf(v2))

    return out //from message
}
func To_Referenda_SubmitCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Referenda_SubmitCall(in)
    out := &pbgearextrinsic.RuntimeCall_ReferendaSubmitCall{ }
    reflect.ValueOf(out).Elem().FieldByName("ReferendaSubmitCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Referenda_SubmitCall_proposal_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_SubmitCall_ProposalOriginSystem{
        ProposalOriginSystem: To_Referenda_System(param),
        }
    case 20:
        return &pbgear.Referenda_SubmitCall_ProposalOriginOrigins{
        ProposalOriginOrigins: To_Referenda_Origins(param),
        }
    case 2:
        return &pbgear.Referenda_SubmitCall_ProposalOriginVoid{
        ProposalOriginVoid: To_Referenda_Void(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Referenda_SubmitCall_proposal (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_SubmitCall_ProposalLegacy{
        ProposalLegacy: To_Referenda_Legacy(param),
        }
    case 1:
        return &pbgear.Referenda_SubmitCall_ProposalInline{
        ProposalInline: To_Referenda_Inline(param),
        }
    case 2:
        return &pbgear.Referenda_SubmitCall_ProposalLookup{
        ProposalLookup: To_Referenda_Lookup(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Referenda_SubmitCall_enactment_moment (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_SubmitCall_EnactmentMomentAt{
        EnactmentMomentAt: To_Referenda_At(param),
        }
    case 1:
        return &pbgear.Referenda_SubmitCall_EnactmentMomentAfter{
        EnactmentMomentAfter: To_Referenda_After(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Referenda_System(in any) *pbgear.Referenda_System {
    out := &pbgear.Referenda_System{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Referenda_System_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Referenda_System_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Referenda_System_Value0Root{
        Value0Root: To_Root(param),
        }
    case 1:
        return &pbgear.Referenda_System_Value0Signed{
        Value0Signed: To_Signed(param),
        }
    case 2:
        return &pbgear.Referenda_System_Value0None{
        Value0None: To_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Referenda_Void(in any) *pbgear.Referenda_Void {
    out := &pbgear.Referenda_Void{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Referenda_Void_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Referenda_Void_value0 (in any) any  {
    return nil
}
func To_ReferendumCanceller(in any) *pbgear.ReferendumCanceller {
    out := &pbgear.ReferendumCanceller{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReferendumKiller(in any) *pbgear.ReferendumKiller {
    out := &pbgear.ReferendumKiller{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_ReinstrumentationFailure(in any) *pbgear.ReinstrumentationFailure {
    out := &pbgear.ReinstrumentationFailure{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Reinstrumented(in any) *pbgear.Reinstrumented {
    out := &pbgear.Reinstrumented{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_RemovedFromWaitlist(in any) *pbgear.RemovedFromWaitlist {
    out := &pbgear.RemovedFromWaitlist{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Reply(in any) *pbgear.Reply {
    out := &pbgear.Reply{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_GprimitivesMessageId(w)
    out.Value0 = To_GprimitivesMessageId(v.ValueAt(0))

    return out //from message
}
    

func To_Reservation(in any) *pbgear.Reservation {
    out := &pbgear.Reservation{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_GprimitivesReservationId(w)
    out.Value0 = To_GprimitivesReservationId(v.ValueAt(0))

    return out //from message
}
    

func To_Reserved(in any) *pbgear.Reserved {
    out := &pbgear.Reserved{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Result(in any) *pbgear.Result {
    out := &pbgear.Result{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Result
    v0 := To_OneOf_Result_result(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Result").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Result_result (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Result_ResultOk{
        ResultOk: To_Ok(param),
        }
    case 1:
        return &pbgear.Result_ResultErr{
        ResultErr: To_Err(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Root(in any) *pbgear.Root {
    out := &pbgear.Root{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_RootNotAllowed(in any) *pbgear.RootNotAllowed {
    out := &pbgear.RootNotAllowed{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Runtime(in any) *pbgear.Runtime {
    out := &pbgear.Runtime{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Runtime_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Runtime_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Runtime_Value0MessageReplied{
        Value0MessageReplied: To_MessageReplied(param),
        }
    case 1:
        return &pbgear.Runtime_Value0MessageClaimed{
        Value0MessageClaimed: To_MessageClaimed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_SchedulerCallUnavailableEvent(in any) *pbgear.SchedulerCallUnavailableEvent {
    out := &pbgear.SchedulerCallUnavailableEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SchedulerCanceledEvent(in any) *pbgear.SchedulerCanceledEvent {
    out := &pbgear.SchedulerCanceledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SchedulerDispatchedEvent(in any) *pbgear.SchedulerDispatchedEvent {
    out := &pbgear.SchedulerDispatchedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SchedulerPallet(in any) *pbgear.SchedulerPallet {
    out := &pbgear.SchedulerPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_SchedulerPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_SchedulerPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.SchedulerPallet_CallScheduleCall{
        CallScheduleCall: To_Scheduler_ScheduleCall(param),
        }
    case 1:
        return &pbgear.SchedulerPallet_CallCancelCall{
        CallCancelCall: To_Scheduler_CancelCall(param),
        }
    case 2:
        return &pbgear.SchedulerPallet_CallScheduleNamedCall{
        CallScheduleNamedCall: To_Scheduler_ScheduleNamedCall(param),
        }
    case 3:
        return &pbgear.SchedulerPallet_CallCancelNamedCall{
        CallCancelNamedCall: To_Scheduler_CancelNamedCall(param),
        }
    case 4:
        return &pbgear.SchedulerPallet_CallScheduleAfterCall{
        CallScheduleAfterCall: To_Scheduler_ScheduleAfterCall(param),
        }
    case 5:
        return &pbgear.SchedulerPallet_CallScheduleNamedAfterCall{
        CallScheduleNamedAfterCall: To_Scheduler_ScheduleNamedAfterCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_SchedulerPeriodicFailedEvent(in any) *pbgear.SchedulerPeriodicFailedEvent {
    out := &pbgear.SchedulerPeriodicFailedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SchedulerPermanentlyOverweightEvent(in any) *pbgear.SchedulerPermanentlyOverweightEvent {
    out := &pbgear.SchedulerPermanentlyOverweightEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SchedulerScheduledEvent(in any) *pbgear.SchedulerScheduledEvent {
    out := &pbgear.SchedulerScheduledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Scheduler_CancelCall(in any) *pbgear.Scheduler_CancelCall {
    out := &pbgear.Scheduler_CancelCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field When
    out.When = To_uint32(v.ValueAt(0))
    // primitive field Index
    out.Index = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_Scheduler_CancelCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Scheduler_CancelCall(in)
    out := &pbgearextrinsic.RuntimeCall_SchedulerCancelCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SchedulerCancelCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Scheduler_CancelNamedCall(in any) *pbgear.Scheduler_CancelNamedCall {
    out := &pbgear.Scheduler_CancelNamedCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Id
    out.Id = To_bytes(v.ValueAt(0))

    return out //from message
}
func To_Scheduler_CancelNamedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Scheduler_CancelNamedCall(in)
    out := &pbgearextrinsic.RuntimeCall_SchedulerCancelNamedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SchedulerCancelNamedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Scheduler_ScheduleAfterCall(in any) *pbgear.Scheduler_ScheduleAfterCall {
    out := &pbgear.Scheduler_ScheduleAfterCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field After
    out.After = To_uint32(v.ValueAt(0))
    // optional field MaybePeriodic
    if v.HasValue() {
        out.MaybePeriodic = To_Optional_Scheduler_ScheduleAfterCall_maybe_periodic(v.ValueAt(1))
    }
    // primitive field Priority
    out.Priority = To_uint32(v.ValueAt(2))
    // oneOf field Call
    v3 := To_OneOf_Scheduler_ScheduleAfterCall_call(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))

    return out //from message
}
func To_Scheduler_ScheduleAfterCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Scheduler_ScheduleAfterCall(in)
    out := &pbgearextrinsic.RuntimeCall_SchedulerScheduleAfterCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SchedulerScheduleAfterCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_Scheduler_ScheduleAfterCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleAfterCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Scheduler_ScheduleAfterCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Scheduler_ScheduleAfterCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Scheduler_ScheduleAfterCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Scheduler_ScheduleAfterCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Scheduler_ScheduleAfterCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Scheduler_ScheduleAfterCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Scheduler_ScheduleAfterCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Scheduler_ScheduleAfterCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Scheduler_ScheduleAfterCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Scheduler_ScheduleAfterCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Scheduler_ScheduleAfterCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Scheduler_ScheduleAfterCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Scheduler_ScheduleAfterCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Scheduler_ScheduleAfterCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Scheduler_ScheduleAfterCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Scheduler_ScheduleAfterCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Scheduler_ScheduleAfterCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Scheduler_ScheduleAfterCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Scheduler_ScheduleAfterCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Scheduler_ScheduleAfterCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Scheduler_ScheduleAfterCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Scheduler_ScheduleAfterCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Scheduler_ScheduleAfterCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Scheduler_ScheduleAfterCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Scheduler_ScheduleAfterCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Scheduler_ScheduleAfterCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Scheduler_ScheduleAfterCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Scheduler_ScheduleAfterCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Scheduler_ScheduleAfterCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Scheduler_ScheduleCall(in any) *pbgear.Scheduler_ScheduleCall {
    out := &pbgear.Scheduler_ScheduleCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field When
    out.When = To_uint32(v.ValueAt(0))
    // optional field MaybePeriodic
    if v.HasValue() {
        out.MaybePeriodic = To_Optional_Scheduler_ScheduleCall_maybe_periodic(v.ValueAt(1))
    }
    // primitive field Priority
    out.Priority = To_uint32(v.ValueAt(2))
    // oneOf field Call
    v3 := To_OneOf_Scheduler_ScheduleCall_call(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v3))

    return out //from message
}
func To_Scheduler_ScheduleCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Scheduler_ScheduleCall(in)
    out := &pbgearextrinsic.RuntimeCall_SchedulerScheduleCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SchedulerScheduleCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_Scheduler_ScheduleCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Scheduler_ScheduleCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Scheduler_ScheduleCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Scheduler_ScheduleCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Scheduler_ScheduleCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Scheduler_ScheduleCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Scheduler_ScheduleCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Scheduler_ScheduleCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Scheduler_ScheduleCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Scheduler_ScheduleCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Scheduler_ScheduleCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Scheduler_ScheduleCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Scheduler_ScheduleCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Scheduler_ScheduleCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Scheduler_ScheduleCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Scheduler_ScheduleCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Scheduler_ScheduleCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Scheduler_ScheduleCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Scheduler_ScheduleCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Scheduler_ScheduleCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Scheduler_ScheduleCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Scheduler_ScheduleCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Scheduler_ScheduleCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Scheduler_ScheduleCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Scheduler_ScheduleCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Scheduler_ScheduleCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Scheduler_ScheduleCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Scheduler_ScheduleCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Scheduler_ScheduleCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Scheduler_ScheduleCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Scheduler_ScheduleNamedAfterCall(in any) *pbgear.Scheduler_ScheduleNamedAfterCall {
    out := &pbgear.Scheduler_ScheduleNamedAfterCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Id
    out.Id = To_bytes(v.ValueAt(0))
    // primitive field After
    out.After = To_uint32(v.ValueAt(1))
    // optional field MaybePeriodic
    if v.HasValue() {
        out.MaybePeriodic = To_Optional_Scheduler_ScheduleNamedAfterCall_maybe_periodic(v.ValueAt(2))
    }
    // primitive field Priority
    out.Priority = To_uint32(v.ValueAt(3))
    // oneOf field Call
    v4 := To_OneOf_Scheduler_ScheduleNamedAfterCall_call(v.ValueAt(4))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v4))

    return out //from message
}
func To_Scheduler_ScheduleNamedAfterCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Scheduler_ScheduleNamedAfterCall(in)
    out := &pbgearextrinsic.RuntimeCall_SchedulerScheduleNamedAfterCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SchedulerScheduleNamedAfterCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_Scheduler_ScheduleNamedAfterCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleNamedAfterCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Scheduler_ScheduleNamedAfterCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Scheduler_ScheduleNamedCall(in any) *pbgear.Scheduler_ScheduleNamedCall {
    out := &pbgear.Scheduler_ScheduleNamedCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Id
    out.Id = To_bytes(v.ValueAt(0))
    // primitive field When
    out.When = To_uint32(v.ValueAt(1))
    // optional field MaybePeriodic
    if v.HasValue() {
        out.MaybePeriodic = To_Optional_Scheduler_ScheduleNamedCall_maybe_periodic(v.ValueAt(2))
    }
    // primitive field Priority
    out.Priority = To_uint32(v.ValueAt(3))
    // oneOf field Call
    v4 := To_OneOf_Scheduler_ScheduleNamedCall_call(v.ValueAt(4))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v4))

    return out //from message
}
func To_Scheduler_ScheduleNamedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Scheduler_ScheduleNamedCall(in)
    out := &pbgearextrinsic.RuntimeCall_SchedulerScheduleNamedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SchedulerScheduleNamedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Optional_Scheduler_ScheduleNamedCall_maybe_periodic(in any) *pbgear.Scheduler_TupleUint32Uint32 {
    panic("Not implemented")
    return nil //funcForOptionalField
}
func To_OneOf_Scheduler_ScheduleNamedCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Scheduler_ScheduleNamedCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Scheduler_ScheduleNamedCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Scheduler_ScheduleNamedCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Scheduler_ScheduleNamedCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Scheduler_ScheduleNamedCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Scheduler_ScheduleNamedCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Scheduler_ScheduleNamedCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Scheduler_ScheduleNamedCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Scheduler_ScheduleNamedCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Scheduler_ScheduleNamedCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Scheduler_ScheduleNamedCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Scheduler_ScheduleNamedCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Scheduler_ScheduleNamedCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Scheduler_ScheduleNamedCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Scheduler_ScheduleNamedCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Scheduler_ScheduleNamedCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Scheduler_ScheduleNamedCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Scheduler_ScheduleNamedCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Scheduler_ScheduleNamedCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Scheduler_ScheduleNamedCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Scheduler_ScheduleNamedCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Scheduler_ScheduleNamedCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Scheduler_ScheduleNamedCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Scheduler_ScheduleNamedCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Scheduler_ScheduleNamedCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Scheduler_ScheduleNamedCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Scheduler_ScheduleNamedCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Scheduler_ScheduleNamedCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Scheduler_ScheduleNamedCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Scheduler_TupleUint32Uint32(in any) *pbgear.Scheduler_TupleUint32Uint32 {
    out := &pbgear.Scheduler_TupleUint32Uint32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_uint32(v.ValueAt(1))

    return out //from message
}
    
func To_SessionNewSessionEvent(in any) *pbgear.SessionNewSessionEvent {
    out := &pbgear.SessionNewSessionEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SessionPallet(in any) *pbgear.SessionPallet {
    out := &pbgear.SessionPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_SessionPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_SessionPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.SessionPallet_CallSetKeysCall{
        CallSetKeysCall: To_Session_SetKeysCall(param),
        }
    case 1:
        return &pbgear.SessionPallet_CallPurgeKeysCall{
        CallPurgeKeysCall: To_Session_PurgeKeysCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Session_PurgeKeysCall(in any) *pbgear.Session_PurgeKeysCall {
    out := &pbgear.Session_PurgeKeysCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Session_SetKeysCall(in any) *pbgear.Session_SetKeysCall {
    out := &pbgear.Session_SetKeysCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Keys To_VaraRuntimeSessionKeys(w)
    out.Keys = To_VaraRuntimeSessionKeys(v.ValueAt(0))
    // primitive field Proof
    out.Proof = To_bytes(v.ValueAt(1))

    return out //from message
}
func To_Session_SetKeysCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Session_SetKeysCall(in)
    out := &pbgearextrinsic.RuntimeCall_SessionSetKeysCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SessionSetKeysCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Signal(in any) *pbgear.Signal {
    out := &pbgear.Signal{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Signed(in any) *pbgear.Signed {
    out := &pbgear.Signed{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_SmallSpender(in any) *pbgear.SmallSpender {
    out := &pbgear.SmallSpender{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SmallTipper(in any) *pbgear.SmallTipper {
    out := &pbgear.SmallTipper{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SpArithmeticPerThingsPerU16(in any) *pbgear.SpArithmeticPerThingsPerU16 {
    out := &pbgear.SpArithmeticPerThingsPerU16{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_SpArithmeticPerThingsPerbill(in any) *pbgear.SpArithmeticPerThingsPerbill {
    out := &pbgear.SpArithmeticPerThingsPerbill{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_SpArithmeticPerThingsPercent(in any) *pbgear.SpArithmeticPerThingsPercent {
    out := &pbgear.SpArithmeticPerThingsPercent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint32(v.ValueAt(0))

    return out //from message
}
    
func To_SpAuthorityDiscoveryAppPublic(in any) *pbgear.SpAuthorityDiscoveryAppPublic {
    out := &pbgear.SpAuthorityDiscoveryAppPublic{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreSr25519Public(w)
    out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

    return out //from message
}
    

func To_SpConsensusBabeAppPublic(in any) *pbgear.SpConsensusBabeAppPublic {
    out := &pbgear.SpConsensusBabeAppPublic{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreSr25519Public(w)
    out.Value0 = To_SpCoreSr25519Public(v.ValueAt(0))

    return out //from message
}
    

func To_SpConsensusGrandpaAppPublic(in any) *pbgear.SpConsensusGrandpaAppPublic {
    out := &pbgear.SpConsensusGrandpaAppPublic{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreEd25519Public(w)
    out.Value0 = To_SpCoreEd25519Public(v.ValueAt(0))

    return out //from message
}
    

func To_SpConsensusSlotsEquivocationProof(in any) *pbgear.SpConsensusSlotsEquivocationProof {
    out := &pbgear.SpConsensusSlotsEquivocationProof{}
    v := in.(registry.Valuable)
    _ = v

    // field Offender To_Babe_SpConsensusBabeAppPublic(w)
    out.Offender = To_Babe_SpConsensusBabeAppPublic(v.ValueAt(0))
    // field Slot To_SpConsensusSlotsSlot(w)
    out.Slot = To_SpConsensusSlotsSlot(v.ValueAt(1))
    // field FirstHeader To_SpRuntimeGenericHeaderHeader(w)
    out.FirstHeader = To_SpRuntimeGenericHeaderHeader(v.ValueAt(2))
    // field SecondHeader To_SpRuntimeGenericHeaderHeader(w)
    out.SecondHeader = To_SpRuntimeGenericHeaderHeader(v.ValueAt(3))

    return out //from message
}
    




func To_SpConsensusSlotsSlot(in any) *pbgear.SpConsensusSlotsSlot {
    out := &pbgear.SpConsensusSlotsSlot{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_uint64(v.ValueAt(0))

    return out //from message
}
    
func To_SpCoreCryptoAccountId32(in any) *pbgear.SpCoreCryptoAccountId32 {
    out := &pbgear.SpCoreCryptoAccountId32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_SpCoreEd25519Public(in any) *pbgear.SpCoreEd25519Public {
    out := &pbgear.SpCoreEd25519Public{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_SpCoreEd25519Signature(in any) *pbgear.SpCoreEd25519Signature {
    out := &pbgear.SpCoreEd25519Signature{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_SpCoreSr25519Public(in any) *pbgear.SpCoreSr25519Public {
    out := &pbgear.SpCoreSr25519Public{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_SpCoreSr25519Signature(in any) *pbgear.SpCoreSr25519Signature {
    out := &pbgear.SpCoreSr25519Signature{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_SpNposElectionsElectionScore(in any) *pbgear.SpNposElectionsElectionScore {
    out := &pbgear.SpNposElectionsElectionScore{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field MinimalStake
    out.MinimalStake = To_string(v.ValueAt(0))
    // primitive field SumStake
    out.SumStake = To_string(v.ValueAt(1))
    // primitive field SumStakeSquared
    out.SumStakeSquared = To_string(v.ValueAt(2))

    return out //from message
}
    
func To_SpNposElectionsSupport(in any) *pbgear.SpNposElectionsSupport {
    out := &pbgear.SpNposElectionsSupport{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Total
    out.Total = To_string(v.ValueAt(0))
    // repeated field Voters
    out.Voters = To_Repeated_SpNposElectionsSupport_voters(v.ValueAt(1))

    return out //from message
}
    

func To_Repeated_SpNposElectionsSupport_voters(in any) []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleSpCoreCryptoAccountId32String(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_SpRuntimeDispatchErrorWithPostInfo(in any) *pbgear.SpRuntimeDispatchErrorWithPostInfo {
    out := &pbgear.SpRuntimeDispatchErrorWithPostInfo{}
    v := in.(registry.Valuable)
    _ = v

    // field PostInfo To_FrameSupportDispatchPostDispatchInfo(w)
    out.PostInfo = To_FrameSupportDispatchPostDispatchInfo(v.ValueAt(0))
    // oneOf field Error
    v1 := To_OneOf_SpRuntimeDispatchErrorWithPostInfo_error(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Error").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_SpRuntimeDispatchErrorWithPostInfo_error (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorOther{
        ErrorOther: To_Other(param),
        }
    case 1:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorCannotLookup{
        ErrorCannotLookup: To_CannotLookup(param),
        }
    case 2:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorBadOrigin{
        ErrorBadOrigin: To_BadOrigin(param),
        }
    case 3:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorModule{
        ErrorModule: To_Module(param),
        }
    case 4:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorConsumerRemaining{
        ErrorConsumerRemaining: To_ConsumerRemaining(param),
        }
    case 5:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorNoProviders{
        ErrorNoProviders: To_NoProviders(param),
        }
    case 6:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorTooManyConsumers{
        ErrorTooManyConsumers: To_TooManyConsumers(param),
        }
    case 7:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorToken{
        ErrorToken: To_Token(param),
        }
    case 8:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorArithmetic{
        ErrorArithmetic: To_Arithmetic(param),
        }
    case 9:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorTransactional{
        ErrorTransactional: To_Transactional(param),
        }
    case 10:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorExhausted{
        ErrorExhausted: To_Exhausted(param),
        }
    case 11:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorCorruption{
        ErrorCorruption: To_Corruption(param),
        }
    case 12:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorUnavailable{
        ErrorUnavailable: To_Unavailable(param),
        }
    case 13:
        return &pbgear.SpRuntimeDispatchErrorWithPostInfo_ErrorRootNotAllowed{
        ErrorRootNotAllowed: To_RootNotAllowed(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_SpRuntimeGenericDigestDigest(in any) *pbgear.SpRuntimeGenericDigestDigest {
    out := &pbgear.SpRuntimeGenericDigestDigest{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Logs
    out.Logs = To_Repeated_SpRuntimeGenericDigestDigest_logs(v.ValueAt(0))

    return out //from message
}
    

func To_Repeated_SpRuntimeGenericDigestDigest_logs(in any) []*pbgear.SpRuntimeGenericDigestDigestItem {
    items := in.([]interface{})

    var out []*pbgear.SpRuntimeGenericDigestDigestItem
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpRuntimeGenericDigestDigestItem(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_SpRuntimeGenericDigestDigestItem(in any) *pbgear.SpRuntimeGenericDigestDigestItem {
    out := &pbgear.SpRuntimeGenericDigestDigestItem{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Logs
    v0 := To_OneOf_SpRuntimeGenericDigestDigestItem_logs(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Logs").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_SpRuntimeGenericDigestDigestItem_logs (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 6:
        return &pbgear.SpRuntimeGenericDigestDigestItem_LogsPreRuntime{
        LogsPreRuntime: To_Babe_PreRuntime(param),
        }
    case 4:
        return &pbgear.SpRuntimeGenericDigestDigestItem_LogsConsensus{
        LogsConsensus: To_Babe_Consensus(param),
        }
    case 5:
        return &pbgear.SpRuntimeGenericDigestDigestItem_LogsSeal{
        LogsSeal: To_Babe_Seal(param),
        }
    case 0:
        return &pbgear.SpRuntimeGenericDigestDigestItem_LogsOther{
        LogsOther: To_Babe_Other(param),
        }
    case 8:
        return &pbgear.SpRuntimeGenericDigestDigestItem_LogsRuntimeEnvironmentUpdated{
        LogsRuntimeEnvironmentUpdated: To_Babe_RuntimeEnvironmentUpdated(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_SpRuntimeGenericHeaderHeader(in any) *pbgear.SpRuntimeGenericHeaderHeader {
    out := &pbgear.SpRuntimeGenericHeaderHeader{}
    v := in.(registry.Valuable)
    _ = v

    // field ParentHash To_PrimitiveTypesH256(w)
    out.ParentHash = To_PrimitiveTypesH256(v.ValueAt(0))
    // primitive field Number
    out.Number = To_uint32(v.ValueAt(1))
    // field StateRoot To_PrimitiveTypesH256(w)
    out.StateRoot = To_PrimitiveTypesH256(v.ValueAt(2))
    // field ExtrinsicsRoot To_PrimitiveTypesH256(w)
    out.ExtrinsicsRoot = To_PrimitiveTypesH256(v.ValueAt(3))
    // field Digest To_SpRuntimeGenericDigestDigest(w)
    out.Digest = To_SpRuntimeGenericDigestDigest(v.ValueAt(4))

    return out //from message
}
    




func To_SpRuntimeModuleError(in any) *pbgear.SpRuntimeModuleError {
    out := &pbgear.SpRuntimeModuleError{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // primitive field Error
    out.Error = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_SpRuntimeMultiaddressMultiAddress(in any) *pbgear.SpRuntimeMultiaddressMultiAddress {
    out := &pbgear.SpRuntimeMultiaddressMultiAddress{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Targets
    v0 := To_OneOf_SpRuntimeMultiaddressMultiAddress_targets(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Targets").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_SpRuntimeMultiaddressMultiAddress_targets (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsId{
        TargetsId: To_Staking_Id(param),
        }
    case 1:
        return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsIndex{
        TargetsIndex: To_Staking_Index(param),
        }
    case 2:
        return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsRaw{
        TargetsRaw: To_Staking_Raw(param),
        }
    case 3:
        return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsAddress32{
        TargetsAddress32: To_Staking_Address32(param),
        }
    case 4:
        return &pbgear.SpRuntimeMultiaddressMultiAddress_TargetsAddress20{
        TargetsAddress20: To_Staking_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_SpSessionMembershipProof(in any) *pbgear.SpSessionMembershipProof {
    out := &pbgear.SpSessionMembershipProof{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Session
    out.Session = To_uint32(v.ValueAt(0))
    // repeated field TrieNodes
    out.TrieNodes = To_Repeated_SpSessionMembershipProof_trie_nodes(v.ValueAt(1))
    // primitive field ValidatorCount
    out.ValidatorCount = To_uint32(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_SpSessionMembershipProof_trie_nodes(in any) []*pbgear.Babe_BabeTrieNodesList {
    items := in.([]interface{})

    var out []*pbgear.Babe_BabeTrieNodesList
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_Babe_BabeTrieNodesList(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_SpWeightsWeightV2Weight(in any) *pbgear.SpWeightsWeightV2Weight {
    out := &pbgear.SpWeightsWeightV2Weight{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field RefTime
    out.RefTime = To_uint64(v.ValueAt(0))
    // primitive field ProofSize
    out.ProofSize = To_uint64(v.ValueAt(1))

    return out //from message
}
    
func To_StackLimitExceeded(in any) *pbgear.StackLimitExceeded {
    out := &pbgear.StackLimitExceeded{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking(in any) *pbgear.Staking {
    out := &pbgear.Staking{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingAdmin(in any) *pbgear.StakingAdmin {
    out := &pbgear.StakingAdmin{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingBondedEvent(in any) *pbgear.StakingBondedEvent {
    out := &pbgear.StakingBondedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingChilledEvent(in any) *pbgear.StakingChilledEvent {
    out := &pbgear.StakingChilledEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingEraPaidEvent(in any) *pbgear.StakingEraPaidEvent {
    out := &pbgear.StakingEraPaidEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingForceEraEvent(in any) *pbgear.StakingForceEraEvent {
    out := &pbgear.StakingForceEraEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingKickedEvent(in any) *pbgear.StakingKickedEvent {
    out := &pbgear.StakingKickedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingOldSlashingReportDiscardedEvent(in any) *pbgear.StakingOldSlashingReportDiscardedEvent {
    out := &pbgear.StakingOldSlashingReportDiscardedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingPallet(in any) *pbgear.StakingPallet {
    out := &pbgear.StakingPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_StakingPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_StakingPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.StakingPallet_CallBondCall{
        CallBondCall: To_Staking_BondCall(param),
        }
    case 1:
        return &pbgear.StakingPallet_CallBondExtraCall{
        CallBondExtraCall: To_Staking_BondExtraCall(param),
        }
    case 2:
        return &pbgear.StakingPallet_CallUnbondCall{
        CallUnbondCall: To_Staking_UnbondCall(param),
        }
    case 3:
        return &pbgear.StakingPallet_CallWithdrawUnbondedCall{
        CallWithdrawUnbondedCall: To_Staking_WithdrawUnbondedCall(param),
        }
    case 4:
        return &pbgear.StakingPallet_CallValidateCall{
        CallValidateCall: To_Staking_ValidateCall(param),
        }
    case 5:
        return &pbgear.StakingPallet_CallNominateCall{
        CallNominateCall: To_Staking_NominateCall(param),
        }
    case 6:
        return &pbgear.StakingPallet_CallChillCall{
        CallChillCall: To_Staking_ChillCall(param),
        }
    case 7:
        return &pbgear.StakingPallet_CallSetPayeeCall{
        CallSetPayeeCall: To_Staking_SetPayeeCall(param),
        }
    case 8:
        return &pbgear.StakingPallet_CallSetControllerCall{
        CallSetControllerCall: To_Staking_SetControllerCall(param),
        }
    case 9:
        return &pbgear.StakingPallet_CallSetValidatorCountCall{
        CallSetValidatorCountCall: To_Staking_SetValidatorCountCall(param),
        }
    case 10:
        return &pbgear.StakingPallet_CallIncreaseValidatorCountCall{
        CallIncreaseValidatorCountCall: To_Staking_IncreaseValidatorCountCall(param),
        }
    case 11:
        return &pbgear.StakingPallet_CallScaleValidatorCountCall{
        CallScaleValidatorCountCall: To_Staking_ScaleValidatorCountCall(param),
        }
    case 12:
        return &pbgear.StakingPallet_CallForceNoErasCall{
        CallForceNoErasCall: To_Staking_ForceNoErasCall(param),
        }
    case 13:
        return &pbgear.StakingPallet_CallForceNewEraCall{
        CallForceNewEraCall: To_Staking_ForceNewEraCall(param),
        }
    case 14:
        return &pbgear.StakingPallet_CallSetInvulnerablesCall{
        CallSetInvulnerablesCall: To_Staking_SetInvulnerablesCall(param),
        }
    case 15:
        return &pbgear.StakingPallet_CallForceUnstakeCall{
        CallForceUnstakeCall: To_Staking_ForceUnstakeCall(param),
        }
    case 16:
        return &pbgear.StakingPallet_CallForceNewEraAlwaysCall{
        CallForceNewEraAlwaysCall: To_Staking_ForceNewEraAlwaysCall(param),
        }
    case 17:
        return &pbgear.StakingPallet_CallCancelDeferredSlashCall{
        CallCancelDeferredSlashCall: To_Staking_CancelDeferredSlashCall(param),
        }
    case 18:
        return &pbgear.StakingPallet_CallPayoutStakersCall{
        CallPayoutStakersCall: To_Staking_PayoutStakersCall(param),
        }
    case 19:
        return &pbgear.StakingPallet_CallRebondCall{
        CallRebondCall: To_Staking_RebondCall(param),
        }
    case 20:
        return &pbgear.StakingPallet_CallReapStashCall{
        CallReapStashCall: To_Staking_ReapStashCall(param),
        }
    case 21:
        return &pbgear.StakingPallet_CallKickCall{
        CallKickCall: To_Staking_KickCall(param),
        }
    case 22:
        return &pbgear.StakingPallet_CallSetStakingConfigsCall{
        CallSetStakingConfigsCall: To_Staking_SetStakingConfigsCall(param),
        }
    case 23:
        return &pbgear.StakingPallet_CallChillOtherCall{
        CallChillOtherCall: To_Staking_ChillOtherCall(param),
        }
    case 24:
        return &pbgear.StakingPallet_CallForceApplyMinCommissionCall{
        CallForceApplyMinCommissionCall: To_Staking_ForceApplyMinCommissionCall(param),
        }
    case 25:
        return &pbgear.StakingPallet_CallSetMinCommissionCall{
        CallSetMinCommissionCall: To_Staking_SetMinCommissionCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_StakingPayoutStartedEvent(in any) *pbgear.StakingPayoutStartedEvent {
    out := &pbgear.StakingPayoutStartedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingRewardedEvent(in any) *pbgear.StakingRewardedEvent {
    out := &pbgear.StakingRewardedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingRewardsBurnedEvent(in any) *pbgear.StakingRewardsBurnedEvent {
    out := &pbgear.StakingRewardsBurnedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingRewardsDepositedEvent(in any) *pbgear.StakingRewardsDepositedEvent {
    out := &pbgear.StakingRewardsDepositedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingRewardsMintedEvent(in any) *pbgear.StakingRewardsMintedEvent {
    out := &pbgear.StakingRewardsMintedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingRewardsPallet(in any) *pbgear.StakingRewardsPallet {
    out := &pbgear.StakingRewardsPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_StakingRewardsPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_StakingRewardsPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.StakingRewardsPallet_CallRefillCall{
        CallRefillCall: To_StakingRewards_RefillCall(param),
        }
    case 1:
        return &pbgear.StakingRewardsPallet_CallForceRefillCall{
        CallForceRefillCall: To_StakingRewards_ForceRefillCall(param),
        }
    case 2:
        return &pbgear.StakingRewardsPallet_CallWithdrawCall{
        CallWithdrawCall: To_StakingRewards_WithdrawCall(param),
        }
    case 3:
        return &pbgear.StakingRewardsPallet_CallAlignSupplyCall{
        CallAlignSupplyCall: To_StakingRewards_AlignSupplyCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_StakingRewardsWithdrawnEvent(in any) *pbgear.StakingRewardsWithdrawnEvent {
    out := &pbgear.StakingRewardsWithdrawnEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingRewards_Address20(in any) *pbgear.StakingRewards_Address20 {
    out := &pbgear.StakingRewards_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_StakingRewards_Address32(in any) *pbgear.StakingRewards_Address32 {
    out := &pbgear.StakingRewards_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_StakingRewards_AlignSupplyCall(in any) *pbgear.StakingRewards_AlignSupplyCall {
    out := &pbgear.StakingRewards_AlignSupplyCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Target
    out.Target = To_string(v.ValueAt(0))

    return out //from message
}
func To_StakingRewards_AlignSupplyCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_StakingRewards_AlignSupplyCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingrewardsAlignSupplyCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingrewardsAlignSupplyCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_StakingRewards_ForceRefillCall(in any) *pbgear.StakingRewards_ForceRefillCall {
    out := &pbgear.StakingRewards_ForceRefillCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field From
    v0 := To_OneOf_StakingRewards_ForceRefillCall_from(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("From").Set(reflect.ValueOf(v0))
    // primitive field Value
    out.Value = To_string(v.ValueAt(1))

    return out //from message
}
func To_StakingRewards_ForceRefillCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_StakingRewards_ForceRefillCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingrewardsForceRefillCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingrewardsForceRefillCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_StakingRewards_ForceRefillCall_from (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.StakingRewards_ForceRefillCall_FromId{
        FromId: To_StakingRewards_Id(param),
        }
    case 1:
        return &pbgear.StakingRewards_ForceRefillCall_FromIndex{
        FromIndex: To_StakingRewards_Index(param),
        }
    case 2:
        return &pbgear.StakingRewards_ForceRefillCall_FromRaw{
        FromRaw: To_StakingRewards_Raw(param),
        }
    case 3:
        return &pbgear.StakingRewards_ForceRefillCall_FromAddress32{
        FromAddress32: To_StakingRewards_Address32(param),
        }
    case 4:
        return &pbgear.StakingRewards_ForceRefillCall_FromAddress20{
        FromAddress20: To_StakingRewards_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_StakingRewards_From(in any) *pbgear.StakingRewards_From {
    out := &pbgear.StakingRewards_From{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field From
    v0 := To_OneOf_StakingRewards_From_from(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("From").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_StakingRewards_From_from (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.StakingRewards_From_FromId{
        FromId: To_StakingRewards_Id(param),
        }
    case 1:
        return &pbgear.StakingRewards_From_FromIndex{
        FromIndex: To_StakingRewards_Index(param),
        }
    case 2:
        return &pbgear.StakingRewards_From_FromRaw{
        FromRaw: To_StakingRewards_Raw(param),
        }
    case 3:
        return &pbgear.StakingRewards_From_FromAddress32{
        FromAddress32: To_StakingRewards_Address32(param),
        }
    case 4:
        return &pbgear.StakingRewards_From_FromAddress20{
        FromAddress20: To_StakingRewards_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_StakingRewards_Id(in any) *pbgear.StakingRewards_Id {
    out := &pbgear.StakingRewards_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_StakingRewards_Index(in any) *pbgear.StakingRewards_Index {
    out := &pbgear.StakingRewards_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_StakingRewards_TupleNull(w)
    out.Value0 = To_StakingRewards_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_StakingRewards_Raw(in any) *pbgear.StakingRewards_Raw {
    out := &pbgear.StakingRewards_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_StakingRewards_RefillCall(in any) *pbgear.StakingRewards_RefillCall {
    out := &pbgear.StakingRewards_RefillCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value
    out.Value = To_string(v.ValueAt(0))

    return out //from message
}
func To_StakingRewards_RefillCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_StakingRewards_RefillCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingrewardsRefillCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingrewardsRefillCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_StakingRewards_To(in any) *pbgear.StakingRewards_To {
    out := &pbgear.StakingRewards_To{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field To
    v0 := To_OneOf_StakingRewards_To_to(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_StakingRewards_To_to (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.StakingRewards_To_ToId{
        ToId: To_StakingRewards_Id(param),
        }
    case 1:
        return &pbgear.StakingRewards_To_ToIndex{
        ToIndex: To_StakingRewards_Index(param),
        }
    case 2:
        return &pbgear.StakingRewards_To_ToRaw{
        ToRaw: To_StakingRewards_Raw(param),
        }
    case 3:
        return &pbgear.StakingRewards_To_ToAddress32{
        ToAddress32: To_StakingRewards_Address32(param),
        }
    case 4:
        return &pbgear.StakingRewards_To_ToAddress20{
        ToAddress20: To_StakingRewards_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_StakingRewards_TupleNull(in any) *pbgear.StakingRewards_TupleNull {
    out := &pbgear.StakingRewards_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_StakingRewards_WithdrawCall(in any) *pbgear.StakingRewards_WithdrawCall {
    out := &pbgear.StakingRewards_WithdrawCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field To
    v0 := To_OneOf_StakingRewards_WithdrawCall_to(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("To").Set(reflect.ValueOf(v0))
    // primitive field Value
    out.Value = To_string(v.ValueAt(1))

    return out //from message
}
func To_StakingRewards_WithdrawCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_StakingRewards_WithdrawCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingrewardsWithdrawCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingrewardsWithdrawCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_StakingRewards_WithdrawCall_to (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.StakingRewards_WithdrawCall_ToId{
        ToId: To_StakingRewards_Id(param),
        }
    case 1:
        return &pbgear.StakingRewards_WithdrawCall_ToIndex{
        ToIndex: To_StakingRewards_Index(param),
        }
    case 2:
        return &pbgear.StakingRewards_WithdrawCall_ToRaw{
        ToRaw: To_StakingRewards_Raw(param),
        }
    case 3:
        return &pbgear.StakingRewards_WithdrawCall_ToAddress32{
        ToAddress32: To_StakingRewards_Address32(param),
        }
    case 4:
        return &pbgear.StakingRewards_WithdrawCall_ToAddress20{
        ToAddress20: To_StakingRewards_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_StakingSlashReportedEvent(in any) *pbgear.StakingSlashReportedEvent {
    out := &pbgear.StakingSlashReportedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingSlashedEvent(in any) *pbgear.StakingSlashedEvent {
    out := &pbgear.StakingSlashedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingSnapshotTargetsSizeExceededEvent(in any) *pbgear.StakingSnapshotTargetsSizeExceededEvent {
    out := &pbgear.StakingSnapshotTargetsSizeExceededEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingSnapshotVotersSizeExceededEvent(in any) *pbgear.StakingSnapshotVotersSizeExceededEvent {
    out := &pbgear.StakingSnapshotVotersSizeExceededEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingStakersElectedEvent(in any) *pbgear.StakingStakersElectedEvent {
    out := &pbgear.StakingStakersElectedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingStakingElectionFailedEvent(in any) *pbgear.StakingStakingElectionFailedEvent {
    out := &pbgear.StakingStakingElectionFailedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingUnbondedEvent(in any) *pbgear.StakingUnbondedEvent {
    out := &pbgear.StakingUnbondedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingValidatorPrefsSetEvent(in any) *pbgear.StakingValidatorPrefsSetEvent {
    out := &pbgear.StakingValidatorPrefsSetEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_StakingWithdrawnEvent(in any) *pbgear.StakingWithdrawnEvent {
    out := &pbgear.StakingWithdrawnEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_Account(in any) *pbgear.Staking_Account {
    out := &pbgear.Staking_Account{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Staking_Address20(in any) *pbgear.Staking_Address20 {
    out := &pbgear.Staking_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Staking_Address32(in any) *pbgear.Staking_Address32 {
    out := &pbgear.Staking_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Staking_BondCall(in any) *pbgear.Staking_BondCall {
    out := &pbgear.Staking_BondCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value
    out.Value = To_string(v.ValueAt(0))
    // oneOf field Payee
    v1 := To_OneOf_Staking_BondCall_payee(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Payee").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Staking_BondCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_BondCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingBondCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingBondCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Staking_BondCall_payee (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_BondCall_PayeeStaked{
        PayeeStaked: To_Staking_Staked(param),
        }
    case 1:
        return &pbgear.Staking_BondCall_PayeeStash{
        PayeeStash: To_Staking_Stash(param),
        }
    case 2:
        return &pbgear.Staking_BondCall_PayeeController{
        PayeeController: To_Staking_Controller(param),
        }
    case 3:
        return &pbgear.Staking_BondCall_PayeeAccount{
        PayeeAccount: To_Staking_Account(param),
        }
    case 4:
        return &pbgear.Staking_BondCall_PayeeNone{
        PayeeNone: To_Staking_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_BondExtraCall(in any) *pbgear.Staking_BondExtraCall {
    out := &pbgear.Staking_BondExtraCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field MaxAdditional
    out.MaxAdditional = To_string(v.ValueAt(0))

    return out //from message
}
func To_Staking_BondExtraCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_BondExtraCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingBondExtraCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingBondExtraCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Staking_CancelDeferredSlashCall(in any) *pbgear.Staking_CancelDeferredSlashCall {
    out := &pbgear.Staking_CancelDeferredSlashCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Era
    out.Era = To_uint32(v.ValueAt(0))
    // primitive field SlashIndices
    out.SlashIndices = To_Repeated_uint32(v.ValueAt(1))

    return out //from message
}
func To_Staking_CancelDeferredSlashCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_CancelDeferredSlashCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingCancelDeferredSlashCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingCancelDeferredSlashCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_ChillCall(in any) *pbgear.Staking_ChillCall {
    out := &pbgear.Staking_ChillCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ChillOtherCall(in any) *pbgear.Staking_ChillOtherCall {
    out := &pbgear.Staking_ChillOtherCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Controller To_SpCoreCryptoAccountId32(w)
    out.Controller = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
func To_Staking_ChillOtherCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_ChillOtherCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingChillOtherCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingChillOtherCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_ChillThreshold(in any) *pbgear.Staking_ChillThreshold {
    out := &pbgear.Staking_ChillThreshold{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field ChillThreshold
    v0 := To_OneOf_Staking_ChillThreshold_chill_threshold(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("ChillThreshold").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_ChillThreshold_chill_threshold (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_ChillThreshold_ChillThresholdNoop{
        ChillThresholdNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_ChillThreshold_ChillThresholdSet{
        ChillThresholdSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_ChillThreshold_ChillThresholdRemove{
        ChillThresholdRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_ConsensusEvent(in any) *pbgear.Staking_ConsensusEvent {
    out := &pbgear.Staking_ConsensusEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Staking_Controller(in any) *pbgear.Staking_Controller {
    out := &pbgear.Staking_Controller{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_Dest(in any) *pbgear.Staking_Dest {
    out := &pbgear.Staking_Dest{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Dest
    v0 := To_OneOf_Staking_Dest_dest(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Dest").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_Dest_dest (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_Dest_DestStaked{
        DestStaked: To_Staking_Staked(param),
        }
    case 1:
        return &pbgear.Staking_Dest_DestStash{
        DestStash: To_Staking_Stash(param),
        }
    case 2:
        return &pbgear.Staking_Dest_DestController{
        DestController: To_Staking_Controller(param),
        }
    case 3:
        return &pbgear.Staking_Dest_DestAccount{
        DestAccount: To_Staking_Account(param),
        }
    case 4:
        return &pbgear.Staking_Dest_DestNone{
        DestNone: To_Staking_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_ForceAlways(in any) *pbgear.Staking_ForceAlways {
    out := &pbgear.Staking_ForceAlways{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ForceApplyMinCommissionCall(in any) *pbgear.Staking_ForceApplyMinCommissionCall {
    out := &pbgear.Staking_ForceApplyMinCommissionCall{}
    v := in.(registry.Valuable)
    _ = v

    // field ValidatorStash To_SpCoreCryptoAccountId32(w)
    out.ValidatorStash = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
func To_Staking_ForceApplyMinCommissionCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_ForceApplyMinCommissionCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingForceApplyMinCommissionCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingForceApplyMinCommissionCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_ForceNew(in any) *pbgear.Staking_ForceNew {
    out := &pbgear.Staking_ForceNew{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ForceNewEraAlwaysCall(in any) *pbgear.Staking_ForceNewEraAlwaysCall {
    out := &pbgear.Staking_ForceNewEraAlwaysCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ForceNewEraCall(in any) *pbgear.Staking_ForceNewEraCall {
    out := &pbgear.Staking_ForceNewEraCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ForceNoErasCall(in any) *pbgear.Staking_ForceNoErasCall {
    out := &pbgear.Staking_ForceNoErasCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ForceNone(in any) *pbgear.Staking_ForceNone {
    out := &pbgear.Staking_ForceNone{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ForceUnstakeCall(in any) *pbgear.Staking_ForceUnstakeCall {
    out := &pbgear.Staking_ForceUnstakeCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field NumSlashingSpans
    out.NumSlashingSpans = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_Staking_ForceUnstakeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_ForceUnstakeCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingForceUnstakeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingForceUnstakeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_Id(in any) *pbgear.Staking_Id {
    out := &pbgear.Staking_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Staking_IncreaseValidatorCountCall(in any) *pbgear.Staking_IncreaseValidatorCountCall {
    out := &pbgear.Staking_IncreaseValidatorCountCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Additional
    out.Additional = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Staking_IncreaseValidatorCountCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_IncreaseValidatorCountCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingIncreaseValidatorCountCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingIncreaseValidatorCountCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Staking_Index(in any) *pbgear.Staking_Index {
    out := &pbgear.Staking_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Staking_TupleNull(w)
    out.Value0 = To_Staking_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Staking_KickCall(in any) *pbgear.Staking_KickCall {
    out := &pbgear.Staking_KickCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Who
    out.Who = To_Repeated_Staking_KickCall_who(v.ValueAt(0))

    return out //from message
}
func To_Staking_KickCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_KickCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingKickCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingKickCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Staking_KickCall_who(in any) []*pbgear.SpRuntimeMultiaddressMultiAddress {
    items := in.([]interface{})

    var out []*pbgear.SpRuntimeMultiaddressMultiAddress
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpRuntimeMultiaddressMultiAddress(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Staking_MaxNominatorCount(in any) *pbgear.Staking_MaxNominatorCount {
    out := &pbgear.Staking_MaxNominatorCount{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MaxNominatorCount
    v0 := To_OneOf_Staking_MaxNominatorCount_max_nominator_count(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MaxNominatorCount").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_MaxNominatorCount_max_nominator_count (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_MaxNominatorCount_MaxNominatorCountNoop{
        MaxNominatorCountNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_MaxNominatorCount_MaxNominatorCountSet{
        MaxNominatorCountSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_MaxNominatorCount_MaxNominatorCountRemove{
        MaxNominatorCountRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_MaxValidatorCount(in any) *pbgear.Staking_MaxValidatorCount {
    out := &pbgear.Staking_MaxValidatorCount{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MaxValidatorCount
    v0 := To_OneOf_Staking_MaxValidatorCount_max_validator_count(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MaxValidatorCount").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_MaxValidatorCount_max_validator_count (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_MaxValidatorCount_MaxValidatorCountNoop{
        MaxValidatorCountNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_MaxValidatorCount_MaxValidatorCountSet{
        MaxValidatorCountSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_MaxValidatorCount_MaxValidatorCountRemove{
        MaxValidatorCountRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_MinCommission(in any) *pbgear.Staking_MinCommission {
    out := &pbgear.Staking_MinCommission{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MinCommission
    v0 := To_OneOf_Staking_MinCommission_min_commission(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MinCommission").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_MinCommission_min_commission (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_MinCommission_MinCommissionNoop{
        MinCommissionNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_MinCommission_MinCommissionSet{
        MinCommissionSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_MinCommission_MinCommissionRemove{
        MinCommissionRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_MinNominatorBond(in any) *pbgear.Staking_MinNominatorBond {
    out := &pbgear.Staking_MinNominatorBond{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MinNominatorBond
    v0 := To_OneOf_Staking_MinNominatorBond_min_nominator_bond(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MinNominatorBond").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_MinNominatorBond_min_nominator_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_MinNominatorBond_MinNominatorBondNoop{
        MinNominatorBondNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_MinNominatorBond_MinNominatorBondSet{
        MinNominatorBondSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_MinNominatorBond_MinNominatorBondRemove{
        MinNominatorBondRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_MinValidatorBond(in any) *pbgear.Staking_MinValidatorBond {
    out := &pbgear.Staking_MinValidatorBond{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MinValidatorBond
    v0 := To_OneOf_Staking_MinValidatorBond_min_validator_bond(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MinValidatorBond").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_MinValidatorBond_min_validator_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_MinValidatorBond_MinValidatorBondNoop{
        MinValidatorBondNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_MinValidatorBond_MinValidatorBondSet{
        MinValidatorBondSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_MinValidatorBond_MinValidatorBondRemove{
        MinValidatorBondRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_Mode(in any) *pbgear.Staking_Mode {
    out := &pbgear.Staking_Mode{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Mode
    v0 := To_OneOf_Staking_Mode_mode(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Mode").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_Mode_mode (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_Mode_ModeNotForcing{
        ModeNotForcing: To_Staking_NotForcing(param),
        }
    case 1:
        return &pbgear.Staking_Mode_ModeForceNew{
        ModeForceNew: To_Staking_ForceNew(param),
        }
    case 2:
        return &pbgear.Staking_Mode_ModeForceNone{
        ModeForceNone: To_Staking_ForceNone(param),
        }
    case 3:
        return &pbgear.Staking_Mode_ModeForceAlways{
        ModeForceAlways: To_Staking_ForceAlways(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_NominateCall(in any) *pbgear.Staking_NominateCall {
    out := &pbgear.Staking_NominateCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Targets
    out.Targets = To_Repeated_Staking_NominateCall_targets(v.ValueAt(0))

    return out //from message
}
func To_Staking_NominateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_NominateCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingNominateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingNominateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Staking_NominateCall_targets(in any) []*pbgear.SpRuntimeMultiaddressMultiAddress {
    items := in.([]interface{})

    var out []*pbgear.SpRuntimeMultiaddressMultiAddress
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpRuntimeMultiaddressMultiAddress(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Staking_None(in any) *pbgear.Staking_None {
    out := &pbgear.Staking_None{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_Noop(in any) *pbgear.Staking_Noop {
    out := &pbgear.Staking_Noop{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_NotForcing(in any) *pbgear.Staking_NotForcing {
    out := &pbgear.Staking_NotForcing{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_OtherEvent(in any) *pbgear.Staking_OtherEvent {
    out := &pbgear.Staking_OtherEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Staking_PalletStakingExposure(in any) *pbgear.Staking_PalletStakingExposure {
    out := &pbgear.Staking_PalletStakingExposure{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Total
    out.Total = To_string(v.ValueAt(0))
    // primitive field Own
    out.Own = To_string(v.ValueAt(1))
    // repeated field Others
    out.Others = To_Repeated_Staking_PalletStakingExposure_others(v.ValueAt(2))

    return out //from message
}
    

func To_Repeated_Staking_PalletStakingExposure_others(in any) []*pbgear.Staking_PalletStakingIndividualExposure {
    items := in.([]interface{})

    var out []*pbgear.Staking_PalletStakingIndividualExposure
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_Staking_PalletStakingIndividualExposure(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Staking_PalletStakingIndividualExposure(in any) *pbgear.Staking_PalletStakingIndividualExposure {
    out := &pbgear.Staking_PalletStakingIndividualExposure{}
    v := in.(registry.Valuable)
    _ = v

    // field Who To_SpCoreCryptoAccountId32(w)
    out.Who = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Value
    out.Value = To_string(v.ValueAt(1))

    return out //from message
}
    

func To_Staking_PalletStakingValidatorPrefs(in any) *pbgear.Staking_PalletStakingValidatorPrefs {
    out := &pbgear.Staking_PalletStakingValidatorPrefs{}
    v := in.(registry.Valuable)
    _ = v

    // field Commission To_SpArithmeticPerThingsPerbill(w)
    out.Commission = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))
    // primitive field Blocked
    out.Blocked = To_bool(v.ValueAt(1))

    return out //from message
}
    

func To_Staking_Payee(in any) *pbgear.Staking_Payee {
    out := &pbgear.Staking_Payee{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Payee
    v0 := To_OneOf_Staking_Payee_payee(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Payee").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_Payee_payee (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_Payee_PayeeStaked{
        PayeeStaked: To_Staking_Staked(param),
        }
    case 1:
        return &pbgear.Staking_Payee_PayeeStash{
        PayeeStash: To_Staking_Stash(param),
        }
    case 2:
        return &pbgear.Staking_Payee_PayeeController{
        PayeeController: To_Staking_Controller(param),
        }
    case 3:
        return &pbgear.Staking_Payee_PayeeAccount{
        PayeeAccount: To_Staking_Account(param),
        }
    case 4:
        return &pbgear.Staking_Payee_PayeeNone{
        PayeeNone: To_Staking_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_PayoutStakersCall(in any) *pbgear.Staking_PayoutStakersCall {
    out := &pbgear.Staking_PayoutStakersCall{}
    v := in.(registry.Valuable)
    _ = v

    // field ValidatorStash To_SpCoreCryptoAccountId32(w)
    out.ValidatorStash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field Era
    out.Era = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_Staking_PayoutStakersCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_PayoutStakersCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingPayoutStakersCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingPayoutStakersCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_PreRuntimeEvent(in any) *pbgear.Staking_PreRuntimeEvent {
    out := &pbgear.Staking_PreRuntimeEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Staking_Raw(in any) *pbgear.Staking_Raw {
    out := &pbgear.Staking_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Staking_ReapStashCall(in any) *pbgear.Staking_ReapStashCall {
    out := &pbgear.Staking_ReapStashCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Stash To_SpCoreCryptoAccountId32(w)
    out.Stash = To_SpCoreCryptoAccountId32(v.ValueAt(0))
    // primitive field NumSlashingSpans
    out.NumSlashingSpans = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_Staking_ReapStashCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_ReapStashCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingReapStashCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingReapStashCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_RebondCall(in any) *pbgear.Staking_RebondCall {
    out := &pbgear.Staking_RebondCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value
    out.Value = To_string(v.ValueAt(0))

    return out //from message
}
func To_Staking_RebondCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_RebondCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingRebondCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingRebondCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Staking_Remove(in any) *pbgear.Staking_Remove {
    out := &pbgear.Staking_Remove{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_RuntimeEnvironmentUpdatedEvent(in any) *pbgear.Staking_RuntimeEnvironmentUpdatedEvent {
    out := &pbgear.Staking_RuntimeEnvironmentUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_ScaleValidatorCountCall(in any) *pbgear.Staking_ScaleValidatorCountCall {
    out := &pbgear.Staking_ScaleValidatorCountCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Factor To_SpArithmeticPerThingsPercent(w)
    out.Factor = To_SpArithmeticPerThingsPercent(v.ValueAt(0))

    return out //from message
}
func To_Staking_ScaleValidatorCountCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_ScaleValidatorCountCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingScaleValidatorCountCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingScaleValidatorCountCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_SealEvent(in any) *pbgear.Staking_SealEvent {
    out := &pbgear.Staking_SealEvent{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Staking_Set(in any) *pbgear.Staking_Set {
    out := &pbgear.Staking_Set{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_string(v.ValueAt(0))

    return out //from message
}
    
func To_Staking_SetControllerCall(in any) *pbgear.Staking_SetControllerCall {
    out := &pbgear.Staking_SetControllerCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_SetInvulnerablesCall(in any) *pbgear.Staking_SetInvulnerablesCall {
    out := &pbgear.Staking_SetInvulnerablesCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Invulnerables
    out.Invulnerables = To_Repeated_Staking_SetInvulnerablesCall_invulnerables(v.ValueAt(0))

    return out //from message
}
func To_Staking_SetInvulnerablesCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_SetInvulnerablesCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingSetInvulnerablesCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingSetInvulnerablesCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Staking_SetInvulnerablesCall_invulnerables(in any) []*pbgear.SpCoreCryptoAccountId32 {
    items := in.([]interface{})

    var out []*pbgear.SpCoreCryptoAccountId32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_SpCoreCryptoAccountId32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Staking_SetMinCommissionCall(in any) *pbgear.Staking_SetMinCommissionCall {
    out := &pbgear.Staking_SetMinCommissionCall{}
    v := in.(registry.Valuable)
    _ = v

    // field New To_SpArithmeticPerThingsPerbill(w)
    out.New = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))

    return out //from message
}
func To_Staking_SetMinCommissionCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_SetMinCommissionCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingSetMinCommissionCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingSetMinCommissionCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_SetPayeeCall(in any) *pbgear.Staking_SetPayeeCall {
    out := &pbgear.Staking_SetPayeeCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Payee
    v0 := To_OneOf_Staking_SetPayeeCall_payee(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Payee").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_Staking_SetPayeeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_SetPayeeCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingSetPayeeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingSetPayeeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Staking_SetPayeeCall_payee (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_SetPayeeCall_PayeeStaked{
        PayeeStaked: To_Staking_Staked(param),
        }
    case 1:
        return &pbgear.Staking_SetPayeeCall_PayeeStash{
        PayeeStash: To_Staking_Stash(param),
        }
    case 2:
        return &pbgear.Staking_SetPayeeCall_PayeeController{
        PayeeController: To_Staking_Controller(param),
        }
    case 3:
        return &pbgear.Staking_SetPayeeCall_PayeeAccount{
        PayeeAccount: To_Staking_Account(param),
        }
    case 4:
        return &pbgear.Staking_SetPayeeCall_PayeeNone{
        PayeeNone: To_Staking_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_SetStakingConfigsCall(in any) *pbgear.Staking_SetStakingConfigsCall {
    out := &pbgear.Staking_SetStakingConfigsCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field MinNominatorBond
    v0 := To_OneOf_Staking_SetStakingConfigsCall_min_nominator_bond(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("MinNominatorBond").Set(reflect.ValueOf(v0))
    // oneOf field MinValidatorBond
    v1 := To_OneOf_Staking_SetStakingConfigsCall_min_validator_bond(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("MinValidatorBond").Set(reflect.ValueOf(v1))
    // oneOf field MaxNominatorCount
    v2 := To_OneOf_Staking_SetStakingConfigsCall_max_nominator_count(v.ValueAt(2))
    reflect.ValueOf(out).Elem().FieldByName("MaxNominatorCount").Set(reflect.ValueOf(v2))
    // oneOf field MaxValidatorCount
    v3 := To_OneOf_Staking_SetStakingConfigsCall_max_validator_count(v.ValueAt(3))
    reflect.ValueOf(out).Elem().FieldByName("MaxValidatorCount").Set(reflect.ValueOf(v3))
    // oneOf field ChillThreshold
    v4 := To_OneOf_Staking_SetStakingConfigsCall_chill_threshold(v.ValueAt(4))
    reflect.ValueOf(out).Elem().FieldByName("ChillThreshold").Set(reflect.ValueOf(v4))
    // oneOf field MinCommission
    v5 := To_OneOf_Staking_SetStakingConfigsCall_min_commission(v.ValueAt(5))
    reflect.ValueOf(out).Elem().FieldByName("MinCommission").Set(reflect.ValueOf(v5))

    return out //from message
}
func To_Staking_SetStakingConfigsCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_SetStakingConfigsCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingSetStakingConfigsCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingSetStakingConfigsCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Staking_SetStakingConfigsCall_min_nominator_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_SetStakingConfigsCall_MinNominatorBondNoop{
        MinNominatorBondNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_SetStakingConfigsCall_MinNominatorBondSet{
        MinNominatorBondSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_SetStakingConfigsCall_MinNominatorBondRemove{
        MinNominatorBondRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Staking_SetStakingConfigsCall_min_validator_bond (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_SetStakingConfigsCall_MinValidatorBondNoop{
        MinValidatorBondNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_SetStakingConfigsCall_MinValidatorBondSet{
        MinValidatorBondSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_SetStakingConfigsCall_MinValidatorBondRemove{
        MinValidatorBondRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Staking_SetStakingConfigsCall_max_nominator_count (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_SetStakingConfigsCall_MaxNominatorCountNoop{
        MaxNominatorCountNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_SetStakingConfigsCall_MaxNominatorCountSet{
        MaxNominatorCountSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_SetStakingConfigsCall_MaxNominatorCountRemove{
        MaxNominatorCountRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Staking_SetStakingConfigsCall_max_validator_count (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_SetStakingConfigsCall_MaxValidatorCountNoop{
        MaxValidatorCountNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_SetStakingConfigsCall_MaxValidatorCountSet{
        MaxValidatorCountSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_SetStakingConfigsCall_MaxValidatorCountRemove{
        MaxValidatorCountRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Staking_SetStakingConfigsCall_chill_threshold (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_SetStakingConfigsCall_ChillThresholdNoop{
        ChillThresholdNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_SetStakingConfigsCall_ChillThresholdSet{
        ChillThresholdSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_SetStakingConfigsCall_ChillThresholdRemove{
        ChillThresholdRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Staking_SetStakingConfigsCall_min_commission (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_SetStakingConfigsCall_MinCommissionNoop{
        MinCommissionNoop: To_Staking_Noop(param),
        }
    case 1:
        return &pbgear.Staking_SetStakingConfigsCall_MinCommissionSet{
        MinCommissionSet: To_Staking_Set(param),
        }
    case 2:
        return &pbgear.Staking_SetStakingConfigsCall_MinCommissionRemove{
        MinCommissionRemove: To_Staking_Remove(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_SetValidatorCountCall(in any) *pbgear.Staking_SetValidatorCountCall {
    out := &pbgear.Staking_SetValidatorCountCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field New
    out.New = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Staking_SetValidatorCountCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_SetValidatorCountCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingSetValidatorCountCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingSetValidatorCountCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Staking_Staked(in any) *pbgear.Staking_Staked {
    out := &pbgear.Staking_Staked{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_Stash(in any) *pbgear.Staking_Stash {
    out := &pbgear.Staking_Stash{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Staking_Targets(in any) *pbgear.Staking_Targets {
    out := &pbgear.Staking_Targets{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Targets
    v0 := To_OneOf_Staking_Targets_targets(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Targets").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_Targets_targets (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_Targets_TargetsId{
        TargetsId: To_Staking_Id(param),
        }
    case 1:
        return &pbgear.Staking_Targets_TargetsIndex{
        TargetsIndex: To_Staking_Index(param),
        }
    case 2:
        return &pbgear.Staking_Targets_TargetsRaw{
        TargetsRaw: To_Staking_Raw(param),
        }
    case 3:
        return &pbgear.Staking_Targets_TargetsAddress32{
        TargetsAddress32: To_Staking_Address32(param),
        }
    case 4:
        return &pbgear.Staking_Targets_TargetsAddress20{
        TargetsAddress20: To_Staking_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_TupleNull(in any) *pbgear.Staking_TupleNull {
    out := &pbgear.Staking_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Staking_UnbondCall(in any) *pbgear.Staking_UnbondCall {
    out := &pbgear.Staking_UnbondCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value
    out.Value = To_string(v.ValueAt(0))

    return out //from message
}
func To_Staking_UnbondCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_UnbondCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingUnbondCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingUnbondCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Staking_ValidateCall(in any) *pbgear.Staking_ValidateCall {
    out := &pbgear.Staking_ValidateCall{}
    v := in.(registry.Valuable)
    _ = v

    // field Prefs To_Staking_PalletStakingValidatorPrefs(w)
    out.Prefs = To_Staking_PalletStakingValidatorPrefs(v.ValueAt(0))

    return out //from message
}
func To_Staking_ValidateCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_ValidateCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingValidateCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingValidateCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Staking_Who(in any) *pbgear.Staking_Who {
    out := &pbgear.Staking_Who{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Who
    v0 := To_OneOf_Staking_Who_who(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Who").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Staking_Who_who (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Staking_Who_WhoId{
        WhoId: To_Staking_Id(param),
        }
    case 1:
        return &pbgear.Staking_Who_WhoIndex{
        WhoIndex: To_Staking_Index(param),
        }
    case 2:
        return &pbgear.Staking_Who_WhoRaw{
        WhoRaw: To_Staking_Raw(param),
        }
    case 3:
        return &pbgear.Staking_Who_WhoAddress32{
        WhoAddress32: To_Staking_Address32(param),
        }
    case 4:
        return &pbgear.Staking_Who_WhoAddress20{
        WhoAddress20: To_Staking_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Staking_WithdrawUnbondedCall(in any) *pbgear.Staking_WithdrawUnbondedCall {
    out := &pbgear.Staking_WithdrawUnbondedCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field NumSlashingSpans
    out.NumSlashingSpans = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Staking_WithdrawUnbondedCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Staking_WithdrawUnbondedCall(in)
    out := &pbgearextrinsic.RuntimeCall_StakingWithdrawUnbondedCall{ }
    reflect.ValueOf(out).Elem().FieldByName("StakingWithdrawUnbondedCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Success(in any) *pbgear.Success {
    out := &pbgear.Success{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Success_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Success_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Success_Value0Auto{
        Value0Auto: To_Auto(param),
        }
    case 1:
        return &pbgear.Success_Value0Manual{
        Value0Manual: To_Manual(param),
        }
    case 255:
        return &pbgear.Success_Value0Unsupported{
        Value0Unsupported: To_Unsupported(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_System(in any) *pbgear.System {
    out := &pbgear.System{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_System_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_System_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.System_Value0OutOfRent{
        Value0OutOfRent: To_OutOfRent(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_SystemCodeUpdatedEvent(in any) *pbgear.SystemCodeUpdatedEvent {
    out := &pbgear.SystemCodeUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SystemExtrinsicFailedEvent(in any) *pbgear.SystemExtrinsicFailedEvent {
    out := &pbgear.SystemExtrinsicFailedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SystemExtrinsicSuccessEvent(in any) *pbgear.SystemExtrinsicSuccessEvent {
    out := &pbgear.SystemExtrinsicSuccessEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SystemKilledAccountEvent(in any) *pbgear.SystemKilledAccountEvent {
    out := &pbgear.SystemKilledAccountEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SystemNewAccountEvent(in any) *pbgear.SystemNewAccountEvent {
    out := &pbgear.SystemNewAccountEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_SystemPallet(in any) *pbgear.SystemPallet {
    out := &pbgear.SystemPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_SystemPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_SystemPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.SystemPallet_CallRemarkCall{
        CallRemarkCall: To_System_RemarkCall(param),
        }
    case 1:
        return &pbgear.SystemPallet_CallSetHeapPagesCall{
        CallSetHeapPagesCall: To_System_SetHeapPagesCall(param),
        }
    case 2:
        return &pbgear.SystemPallet_CallSetCodeCall{
        CallSetCodeCall: To_System_SetCodeCall(param),
        }
    case 3:
        return &pbgear.SystemPallet_CallSetCodeWithoutChecksCall{
        CallSetCodeWithoutChecksCall: To_System_SetCodeWithoutChecksCall(param),
        }
    case 4:
        return &pbgear.SystemPallet_CallSetStorageCall{
        CallSetStorageCall: To_System_SetStorageCall(param),
        }
    case 5:
        return &pbgear.SystemPallet_CallKillStorageCall{
        CallKillStorageCall: To_System_KillStorageCall(param),
        }
    case 6:
        return &pbgear.SystemPallet_CallKillPrefixCall{
        CallKillPrefixCall: To_System_KillPrefixCall(param),
        }
    case 7:
        return &pbgear.SystemPallet_CallRemarkWithEventCall{
        CallRemarkWithEventCall: To_System_RemarkWithEventCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_SystemRemarkedEvent(in any) *pbgear.SystemRemarkedEvent {
    out := &pbgear.SystemRemarkedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_System_KillPrefixCall(in any) *pbgear.System_KillPrefixCall {
    out := &pbgear.System_KillPrefixCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Prefix
    out.Prefix = To_bytes(v.ValueAt(0))
    // primitive field Subkeys
    out.Subkeys = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_System_KillPrefixCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_KillPrefixCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemKillPrefixCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemKillPrefixCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_System_KillStorageCall(in any) *pbgear.System_KillStorageCall {
    out := &pbgear.System_KillStorageCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Keys
    out.Keys = To_Repeated_System_KillStorageCall_keys(v.ValueAt(0))

    return out //from message
}
func To_System_KillStorageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_KillStorageCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemKillStorageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemKillStorageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_System_KillStorageCall_keys(in any) []*pbgear.System_SystemKeysList {
    items := in.([]interface{})

    var out []*pbgear.System_SystemKeysList
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_System_SystemKeysList(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_System_RemarkCall(in any) *pbgear.System_RemarkCall {
    out := &pbgear.System_RemarkCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Remark
    out.Remark = To_bytes(v.ValueAt(0))

    return out //from message
}
func To_System_RemarkCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_RemarkCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemRemarkCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemRemarkCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_System_RemarkWithEventCall(in any) *pbgear.System_RemarkWithEventCall {
    out := &pbgear.System_RemarkWithEventCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Remark
    out.Remark = To_bytes(v.ValueAt(0))

    return out //from message
}
func To_System_RemarkWithEventCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_RemarkWithEventCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemRemarkWithEventCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemRemarkWithEventCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_System_SetCodeCall(in any) *pbgear.System_SetCodeCall {
    out := &pbgear.System_SetCodeCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Code
    out.Code = To_bytes(v.ValueAt(0))

    return out //from message
}
func To_System_SetCodeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_SetCodeCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemSetCodeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemSetCodeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_System_SetCodeWithoutChecksCall(in any) *pbgear.System_SetCodeWithoutChecksCall {
    out := &pbgear.System_SetCodeWithoutChecksCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Code
    out.Code = To_bytes(v.ValueAt(0))

    return out //from message
}
func To_System_SetCodeWithoutChecksCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_SetCodeWithoutChecksCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemSetCodeWithoutChecksCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemSetCodeWithoutChecksCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_System_SetHeapPagesCall(in any) *pbgear.System_SetHeapPagesCall {
    out := &pbgear.System_SetHeapPagesCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Pages
    out.Pages = To_uint64(v.ValueAt(0))

    return out //from message
}
func To_System_SetHeapPagesCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_SetHeapPagesCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemSetHeapPagesCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemSetHeapPagesCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_System_SetStorageCall(in any) *pbgear.System_SetStorageCall {
    out := &pbgear.System_SetStorageCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Items
    out.Items = To_Repeated_System_SetStorageCall_items(v.ValueAt(0))

    return out //from message
}
func To_System_SetStorageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_System_SetStorageCall(in)
    out := &pbgearextrinsic.RuntimeCall_SystemSetStorageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("SystemSetStorageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_System_SetStorageCall_items(in any) []*pbgear.System_TupleSystemItemsListSystemItemsList {
    items := in.([]interface{})

    var out []*pbgear.System_TupleSystemItemsListSystemItemsList
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_System_TupleSystemItemsListSystemItemsList(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_System_SystemKeysList(in any) *pbgear.System_SystemKeysList {
    out := &pbgear.System_SystemKeysList{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Keys
    out.Keys = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_System_TupleSystemItemsListSystemItemsList(in any) *pbgear.System_TupleSystemItemsListSystemItemsList {
    out := &pbgear.System_TupleSystemItemsListSystemItemsList{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))
    // primitive field Value1
    out.Value1 = To_bytes(v.ValueAt(1))

    return out //from message
}
    
func To_Terminated(in any) *pbgear.Terminated {
    out := &pbgear.Terminated{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TimeoutHasCome(in any) *pbgear.TimeoutHasCome {
    out := &pbgear.TimeoutHasCome{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TimestampPallet(in any) *pbgear.TimestampPallet {
    out := &pbgear.TimestampPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_TimestampPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_TimestampPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.TimestampPallet_CallSetCall{
        CallSetCall: To_Timestamp_SetCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Timestamp_SetCall(in any) *pbgear.Timestamp_SetCall {
    out := &pbgear.Timestamp_SetCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Now
    out.Now = To_uint64(v.ValueAt(0))

    return out //from message
}
func To_Timestamp_SetCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Timestamp_SetCall(in)
    out := &pbgearextrinsic.RuntimeCall_TimestampSetCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TimestampSetCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Token(in any) *pbgear.Token {
    out := &pbgear.Token{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Token_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Token_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Token_Value0FundsUnavailable{
        Value0FundsUnavailable: To_FundsUnavailable(param),
        }
    case 1:
        return &pbgear.Token_Value0OnlyProvider{
        Value0OnlyProvider: To_OnlyProvider(param),
        }
    case 2:
        return &pbgear.Token_Value0BelowMinimum{
        Value0BelowMinimum: To_BelowMinimum(param),
        }
    case 3:
        return &pbgear.Token_Value0CannotCreate{
        Value0CannotCreate: To_CannotCreate(param),
        }
    case 4:
        return &pbgear.Token_Value0UnknownAsset{
        Value0UnknownAsset: To_UnknownAsset(param),
        }
    case 5:
        return &pbgear.Token_Value0Frozen{
        Value0Frozen: To_Frozen(param),
        }
    case 6:
        return &pbgear.Token_Value0Unsupported{
        Value0Unsupported: To_Unsupported(param),
        }
    case 7:
        return &pbgear.Token_Value0CannotCreateHold{
        Value0CannotCreateHold: To_CannotCreateHold(param),
        }
    case 8:
        return &pbgear.Token_Value0NotExpendable{
        Value0NotExpendable: To_NotExpendable(param),
        }
    case 9:
        return &pbgear.Token_Value0Blocked{
        Value0Blocked: To_Blocked(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_TooManyConsumers(in any) *pbgear.TooManyConsumers {
    out := &pbgear.TooManyConsumers{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TransactionPaymentPallet(in any) *pbgear.TransactionPaymentPallet {
    out := &pbgear.TransactionPaymentPallet{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TransactionPaymentTransactionFeePaidEvent(in any) *pbgear.TransactionPaymentTransactionFeePaidEvent {
    out := &pbgear.TransactionPaymentTransactionFeePaidEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Transactional(in any) *pbgear.Transactional {
    out := &pbgear.Transactional{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Transactional_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Transactional_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Transactional_Value0LimitReached{
        Value0LimitReached: To_LimitReached(param),
        }
    case 1:
        return &pbgear.Transactional_Value0NoLayer{
        Value0NoLayer: To_NoLayer(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Treasurer(in any) *pbgear.Treasurer {
    out := &pbgear.Treasurer{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryAssetSpendApprovedEvent(in any) *pbgear.TreasuryAssetSpendApprovedEvent {
    out := &pbgear.TreasuryAssetSpendApprovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryAssetSpendVoidedEvent(in any) *pbgear.TreasuryAssetSpendVoidedEvent {
    out := &pbgear.TreasuryAssetSpendVoidedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryAwardedEvent(in any) *pbgear.TreasuryAwardedEvent {
    out := &pbgear.TreasuryAwardedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryBurntEvent(in any) *pbgear.TreasuryBurntEvent {
    out := &pbgear.TreasuryBurntEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryDepositEvent(in any) *pbgear.TreasuryDepositEvent {
    out := &pbgear.TreasuryDepositEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryPaidEvent(in any) *pbgear.TreasuryPaidEvent {
    out := &pbgear.TreasuryPaidEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryPallet(in any) *pbgear.TreasuryPallet {
    out := &pbgear.TreasuryPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_TreasuryPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_TreasuryPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.TreasuryPallet_CallProposeSpendCall{
        CallProposeSpendCall: To_Treasury_ProposeSpendCall(param),
        }
    case 1:
        return &pbgear.TreasuryPallet_CallRejectProposalCall{
        CallRejectProposalCall: To_Treasury_RejectProposalCall(param),
        }
    case 2:
        return &pbgear.TreasuryPallet_CallApproveProposalCall{
        CallApproveProposalCall: To_Treasury_ApproveProposalCall(param),
        }
    case 3:
        return &pbgear.TreasuryPallet_CallSpendLocalCall{
        CallSpendLocalCall: To_Treasury_SpendLocalCall(param),
        }
    case 4:
        return &pbgear.TreasuryPallet_CallRemoveApprovalCall{
        CallRemoveApprovalCall: To_Treasury_RemoveApprovalCall(param),
        }
    case 5:
        return &pbgear.TreasuryPallet_CallSpendCall{
        CallSpendCall: To_Treasury_SpendCall(param),
        }
    case 6:
        return &pbgear.TreasuryPallet_CallPayoutCall{
        CallPayoutCall: To_Treasury_PayoutCall(param),
        }
    case 7:
        return &pbgear.TreasuryPallet_CallCheckStatusCall{
        CallCheckStatusCall: To_Treasury_CheckStatusCall(param),
        }
    case 8:
        return &pbgear.TreasuryPallet_CallVoidSpendCall{
        CallVoidSpendCall: To_Treasury_VoidSpendCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_TreasuryPaymentFailedEvent(in any) *pbgear.TreasuryPaymentFailedEvent {
    out := &pbgear.TreasuryPaymentFailedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryProposedEvent(in any) *pbgear.TreasuryProposedEvent {
    out := &pbgear.TreasuryProposedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryRejectedEvent(in any) *pbgear.TreasuryRejectedEvent {
    out := &pbgear.TreasuryRejectedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryRolloverEvent(in any) *pbgear.TreasuryRolloverEvent {
    out := &pbgear.TreasuryRolloverEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasurySpendApprovedEvent(in any) *pbgear.TreasurySpendApprovedEvent {
    out := &pbgear.TreasurySpendApprovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasurySpendProcessedEvent(in any) *pbgear.TreasurySpendProcessedEvent {
    out := &pbgear.TreasurySpendProcessedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasurySpendingEvent(in any) *pbgear.TreasurySpendingEvent {
    out := &pbgear.TreasurySpendingEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TreasuryUpdatedInactiveEvent(in any) *pbgear.TreasuryUpdatedInactiveEvent {
    out := &pbgear.TreasuryUpdatedInactiveEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Treasury_Address20(in any) *pbgear.Treasury_Address20 {
    out := &pbgear.Treasury_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Treasury_Address32(in any) *pbgear.Treasury_Address32 {
    out := &pbgear.Treasury_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Treasury_ApproveProposalCall(in any) *pbgear.Treasury_ApproveProposalCall {
    out := &pbgear.Treasury_ApproveProposalCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ProposalId
    out.ProposalId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Treasury_ApproveProposalCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_ApproveProposalCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasuryApproveProposalCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasuryApproveProposalCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Treasury_Beneficiary(in any) *pbgear.Treasury_Beneficiary {
    out := &pbgear.Treasury_Beneficiary{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Beneficiary
    v0 := To_OneOf_Treasury_Beneficiary_beneficiary(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Treasury_Beneficiary_beneficiary (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Treasury_Beneficiary_BeneficiaryId{
        BeneficiaryId: To_Treasury_Id(param),
        }
    case 1:
        return &pbgear.Treasury_Beneficiary_BeneficiaryIndex{
        BeneficiaryIndex: To_Treasury_Index(param),
        }
    case 2:
        return &pbgear.Treasury_Beneficiary_BeneficiaryRaw{
        BeneficiaryRaw: To_Treasury_Raw(param),
        }
    case 3:
        return &pbgear.Treasury_Beneficiary_BeneficiaryAddress32{
        BeneficiaryAddress32: To_Treasury_Address32(param),
        }
    case 4:
        return &pbgear.Treasury_Beneficiary_BeneficiaryAddress20{
        BeneficiaryAddress20: To_Treasury_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Treasury_CheckStatusCall(in any) *pbgear.Treasury_CheckStatusCall {
    out := &pbgear.Treasury_CheckStatusCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Treasury_CheckStatusCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_CheckStatusCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasuryCheckStatusCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasuryCheckStatusCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Treasury_Id(in any) *pbgear.Treasury_Id {
    out := &pbgear.Treasury_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Treasury_Index(in any) *pbgear.Treasury_Index {
    out := &pbgear.Treasury_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Treasury_TupleNull(w)
    out.Value0 = To_Treasury_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Treasury_PayoutCall(in any) *pbgear.Treasury_PayoutCall {
    out := &pbgear.Treasury_PayoutCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Treasury_PayoutCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_PayoutCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasuryPayoutCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasuryPayoutCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Treasury_ProposeSpendCall(in any) *pbgear.Treasury_ProposeSpendCall {
    out := &pbgear.Treasury_ProposeSpendCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value
    out.Value = To_string(v.ValueAt(0))
    // oneOf field Beneficiary
    v1 := To_OneOf_Treasury_ProposeSpendCall_beneficiary(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Treasury_ProposeSpendCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_ProposeSpendCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasuryProposeSpendCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasuryProposeSpendCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Treasury_ProposeSpendCall_beneficiary (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Treasury_ProposeSpendCall_BeneficiaryId{
        BeneficiaryId: To_Treasury_Id(param),
        }
    case 1:
        return &pbgear.Treasury_ProposeSpendCall_BeneficiaryIndex{
        BeneficiaryIndex: To_Treasury_Index(param),
        }
    case 2:
        return &pbgear.Treasury_ProposeSpendCall_BeneficiaryRaw{
        BeneficiaryRaw: To_Treasury_Raw(param),
        }
    case 3:
        return &pbgear.Treasury_ProposeSpendCall_BeneficiaryAddress32{
        BeneficiaryAddress32: To_Treasury_Address32(param),
        }
    case 4:
        return &pbgear.Treasury_ProposeSpendCall_BeneficiaryAddress20{
        BeneficiaryAddress20: To_Treasury_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Treasury_Raw(in any) *pbgear.Treasury_Raw {
    out := &pbgear.Treasury_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Treasury_RejectProposalCall(in any) *pbgear.Treasury_RejectProposalCall {
    out := &pbgear.Treasury_RejectProposalCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ProposalId
    out.ProposalId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Treasury_RejectProposalCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_RejectProposalCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasuryRejectProposalCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasuryRejectProposalCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Treasury_RemoveApprovalCall(in any) *pbgear.Treasury_RemoveApprovalCall {
    out := &pbgear.Treasury_RemoveApprovalCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field ProposalId
    out.ProposalId = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Treasury_RemoveApprovalCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_RemoveApprovalCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasuryRemoveApprovalCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasuryRemoveApprovalCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Treasury_SpendCall(in any) *pbgear.Treasury_SpendCall {
    out := &pbgear.Treasury_SpendCall{}
    v := in.(registry.Valuable)
    _ = v

    // field AssetKind To_Treasury_TupleNull(w)
    out.AssetKind = To_Treasury_TupleNull(v.ValueAt(0))
    // primitive field Amount
    out.Amount = To_string(v.ValueAt(1))
    // field Beneficiary To_SpCoreCryptoAccountId32(w)
    out.Beneficiary = To_SpCoreCryptoAccountId32(v.ValueAt(2))
    // primitive field ValidFrom
    out.ValidFrom = To_Optional_uint32(v.ValueAt(3))

    return out //from message
}
func To_Treasury_SpendCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_SpendCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasurySpendCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasurySpendCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_Treasury_SpendLocalCall(in any) *pbgear.Treasury_SpendLocalCall {
    out := &pbgear.Treasury_SpendLocalCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Amount
    out.Amount = To_string(v.ValueAt(0))
    // oneOf field Beneficiary
    v1 := To_OneOf_Treasury_SpendLocalCall_beneficiary(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Beneficiary").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Treasury_SpendLocalCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_SpendLocalCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasurySpendLocalCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasurySpendLocalCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Treasury_SpendLocalCall_beneficiary (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Treasury_SpendLocalCall_BeneficiaryId{
        BeneficiaryId: To_Treasury_Id(param),
        }
    case 1:
        return &pbgear.Treasury_SpendLocalCall_BeneficiaryIndex{
        BeneficiaryIndex: To_Treasury_Index(param),
        }
    case 2:
        return &pbgear.Treasury_SpendLocalCall_BeneficiaryRaw{
        BeneficiaryRaw: To_Treasury_Raw(param),
        }
    case 3:
        return &pbgear.Treasury_SpendLocalCall_BeneficiaryAddress32{
        BeneficiaryAddress32: To_Treasury_Address32(param),
        }
    case 4:
        return &pbgear.Treasury_SpendLocalCall_BeneficiaryAddress20{
        BeneficiaryAddress20: To_Treasury_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Treasury_TupleNull(in any) *pbgear.Treasury_TupleNull {
    out := &pbgear.Treasury_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Treasury_VoidSpendCall(in any) *pbgear.Treasury_VoidSpendCall {
    out := &pbgear.Treasury_VoidSpendCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))

    return out //from message
}
func To_Treasury_VoidSpendCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Treasury_VoidSpendCall(in)
    out := &pbgearextrinsic.RuntimeCall_TreasuryVoidSpendCall{ }
    reflect.ValueOf(out).Elem().FieldByName("TreasuryVoidSpendCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_TupleGprimitivesMessageIdgearCommonEventDispatchStatus(in any) *pbgear.TupleGprimitivesMessageIdgearCommonEventDispatchStatus {
    out := &pbgear.TupleGprimitivesMessageIdgearCommonEventDispatchStatus{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_GprimitivesMessageId(w)
    out.Value0 = To_GprimitivesMessageId(v.ValueAt(0))
    // oneOf field Value1
    v1 := To_OneOf_TupleGprimitivesMessageIdgearCommonEventDispatchStatus_value1(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v1))

    return out //from message
}
    

func To_OneOf_TupleGprimitivesMessageIdgearCommonEventDispatchStatus_value1 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.TupleGprimitivesMessageIdgearCommonEventDispatchStatus_Value1Success{
        Value1Success: To_Success(param),
        }
    case 1:
        return &pbgear.TupleGprimitivesMessageIdgearCommonEventDispatchStatus_Value1Failed{
        Value1Failed: To_Failed(param),
        }
    case 2:
        return &pbgear.TupleGprimitivesMessageIdgearCommonEventDispatchStatus_Value1NotExecuted{
        Value1NotExecuted: To_NotExecuted(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_TupleNull(in any) *pbgear.TupleNull {
    out := &pbgear.TupleNull{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32(in any) *pbgear.TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32 {
    out := &pbgear.TupleSpArithmeticPerThingsPerbillspCoreCryptoAccountId32{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpArithmeticPerThingsPerbill(w)
    out.Value0 = To_SpArithmeticPerThingsPerbill(v.ValueAt(0))
    // field Value1 To_SpCoreCryptoAccountId32(w)
    out.Value1 = To_SpCoreCryptoAccountId32(v.ValueAt(1))

    return out //from message
}
    


func To_Unavailable(in any) *pbgear.Unavailable {
    out := &pbgear.Unavailable{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Underflow(in any) *pbgear.Underflow {
    out := &pbgear.Underflow{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UnknownAsset(in any) *pbgear.UnknownAsset {
    out := &pbgear.UnknownAsset{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UnreachableInstruction(in any) *pbgear.UnreachableInstruction {
    out := &pbgear.UnreachableInstruction{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Unsupported(in any) *pbgear.Unsupported {
    out := &pbgear.Unsupported{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UserspacePanic(in any) *pbgear.UserspacePanic {
    out := &pbgear.UserspacePanic{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UtilityBatchCompletedEvent(in any) *pbgear.UtilityBatchCompletedEvent {
    out := &pbgear.UtilityBatchCompletedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UtilityBatchCompletedWithErrorsEvent(in any) *pbgear.UtilityBatchCompletedWithErrorsEvent {
    out := &pbgear.UtilityBatchCompletedWithErrorsEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UtilityBatchInterruptedEvent(in any) *pbgear.UtilityBatchInterruptedEvent {
    out := &pbgear.UtilityBatchInterruptedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UtilityDispatchedAsEvent(in any) *pbgear.UtilityDispatchedAsEvent {
    out := &pbgear.UtilityDispatchedAsEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UtilityItemCompletedEvent(in any) *pbgear.UtilityItemCompletedEvent {
    out := &pbgear.UtilityItemCompletedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UtilityItemFailedEvent(in any) *pbgear.UtilityItemFailedEvent {
    out := &pbgear.UtilityItemFailedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_UtilityPallet(in any) *pbgear.UtilityPallet {
    out := &pbgear.UtilityPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_UtilityPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_UtilityPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.UtilityPallet_CallBatchCall{
        CallBatchCall: To_Utility_BatchCall(param),
        }
    case 1:
        return &pbgear.UtilityPallet_CallAsDerivativeCall{
        CallAsDerivativeCall: To_Utility_AsDerivativeCall(param),
        }
    case 2:
        return &pbgear.UtilityPallet_CallBatchAllCall{
        CallBatchAllCall: To_Utility_BatchAllCall(param),
        }
    case 3:
        return &pbgear.UtilityPallet_CallDispatchAsCall{
        CallDispatchAsCall: To_Utility_DispatchAsCall(param),
        }
    case 4:
        return &pbgear.UtilityPallet_CallForceBatchCall{
        CallForceBatchCall: To_Utility_ForceBatchCall(param),
        }
    case 5:
        return &pbgear.UtilityPallet_CallWithWeightCall{
        CallWithWeightCall: To_Utility_WithWeightCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Utility_AsDerivativeCall(in any) *pbgear.Utility_AsDerivativeCall {
    out := &pbgear.Utility_AsDerivativeCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Index
    out.Index = To_uint32(v.ValueAt(0))
    // oneOf field Call
    v1 := To_OneOf_Utility_AsDerivativeCall_call(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Utility_AsDerivativeCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Utility_AsDerivativeCall(in)
    out := &pbgearextrinsic.RuntimeCall_UtilityAsDerivativeCall{ }
    reflect.ValueOf(out).Elem().FieldByName("UtilityAsDerivativeCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Utility_AsDerivativeCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Utility_AsDerivativeCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Utility_AsDerivativeCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Utility_AsDerivativeCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Utility_AsDerivativeCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Utility_AsDerivativeCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Utility_AsDerivativeCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Utility_AsDerivativeCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Utility_AsDerivativeCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Utility_AsDerivativeCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Utility_AsDerivativeCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Utility_AsDerivativeCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Utility_AsDerivativeCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Utility_AsDerivativeCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Utility_AsDerivativeCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Utility_AsDerivativeCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Utility_AsDerivativeCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Utility_AsDerivativeCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Utility_AsDerivativeCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Utility_AsDerivativeCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Utility_AsDerivativeCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Utility_AsDerivativeCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Utility_AsDerivativeCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Utility_AsDerivativeCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Utility_AsDerivativeCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Utility_AsDerivativeCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Utility_AsDerivativeCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Utility_AsDerivativeCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Utility_AsDerivativeCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Utility_AsDerivativeCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Utility_AsOrigin(in any) *pbgear.Utility_AsOrigin {
    out := &pbgear.Utility_AsOrigin{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field AsOrigin
    v0 := To_OneOf_Utility_AsOrigin_as_origin(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("AsOrigin").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Utility_AsOrigin_as_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Utility_AsOrigin_AsOriginSystem{
        AsOriginSystem: To_Utility_System(param),
        }
    case 20:
        return &pbgear.Utility_AsOrigin_AsOriginOrigins{
        AsOriginOrigins: To_Utility_Origins(param),
        }
    case 2:
        return &pbgear.Utility_AsOrigin_AsOriginVoid{
        AsOriginVoid: To_Utility_Void(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Utility_BatchAllCall(in any) *pbgear.Utility_BatchAllCall {
    out := &pbgear.Utility_BatchAllCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Calls
    out.Calls = To_Repeated_Utility_BatchAllCall_calls(v.ValueAt(0))

    return out //from message
}
func To_Utility_BatchAllCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Utility_BatchAllCall(in)
    out := &pbgearextrinsic.RuntimeCall_UtilityBatchAllCall{ }
    reflect.ValueOf(out).Elem().FieldByName("UtilityBatchAllCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Utility_BatchAllCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
    items := in.([]interface{})

    var out []*pbgear.VaraRuntimeRuntimeCall
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_VaraRuntimeRuntimeCall(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Utility_BatchCall(in any) *pbgear.Utility_BatchCall {
    out := &pbgear.Utility_BatchCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Calls
    out.Calls = To_Repeated_Utility_BatchCall_calls(v.ValueAt(0))

    return out //from message
}
func To_Utility_BatchCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Utility_BatchCall(in)
    out := &pbgearextrinsic.RuntimeCall_UtilityBatchCall{ }
    reflect.ValueOf(out).Elem().FieldByName("UtilityBatchCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Utility_BatchCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
    items := in.([]interface{})

    var out []*pbgear.VaraRuntimeRuntimeCall
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_VaraRuntimeRuntimeCall(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Utility_DispatchAsCall(in any) *pbgear.Utility_DispatchAsCall {
    out := &pbgear.Utility_DispatchAsCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field AsOrigin
    v0 := To_OneOf_Utility_DispatchAsCall_as_origin(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("AsOrigin").Set(reflect.ValueOf(v0))
    // oneOf field Call
    v1 := To_OneOf_Utility_DispatchAsCall_call(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v1))

    return out //from message
}
func To_Utility_DispatchAsCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Utility_DispatchAsCall(in)
    out := &pbgearextrinsic.RuntimeCall_UtilityDispatchAsCall{ }
    reflect.ValueOf(out).Elem().FieldByName("UtilityDispatchAsCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Utility_DispatchAsCall_as_origin (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Utility_DispatchAsCall_AsOriginSystem{
        AsOriginSystem: To_Utility_System(param),
        }
    case 20:
        return &pbgear.Utility_DispatchAsCall_AsOriginOrigins{
        AsOriginOrigins: To_Utility_Origins(param),
        }
    case 2:
        return &pbgear.Utility_DispatchAsCall_AsOriginVoid{
        AsOriginVoid: To_Utility_Void(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Utility_DispatchAsCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Utility_DispatchAsCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Utility_DispatchAsCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Utility_DispatchAsCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Utility_DispatchAsCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Utility_DispatchAsCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Utility_DispatchAsCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Utility_DispatchAsCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Utility_DispatchAsCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Utility_DispatchAsCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Utility_DispatchAsCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Utility_DispatchAsCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Utility_DispatchAsCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Utility_DispatchAsCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Utility_DispatchAsCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Utility_DispatchAsCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Utility_DispatchAsCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Utility_DispatchAsCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Utility_DispatchAsCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Utility_DispatchAsCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Utility_DispatchAsCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Utility_DispatchAsCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Utility_DispatchAsCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Utility_DispatchAsCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Utility_DispatchAsCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Utility_DispatchAsCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Utility_DispatchAsCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Utility_DispatchAsCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Utility_DispatchAsCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Utility_DispatchAsCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Utility_ForceBatchCall(in any) *pbgear.Utility_ForceBatchCall {
    out := &pbgear.Utility_ForceBatchCall{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Calls
    out.Calls = To_Repeated_Utility_ForceBatchCall_calls(v.ValueAt(0))

    return out //from message
}
func To_Utility_ForceBatchCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Utility_ForceBatchCall(in)
    out := &pbgearextrinsic.RuntimeCall_UtilityForceBatchCall{ }
    reflect.ValueOf(out).Elem().FieldByName("UtilityForceBatchCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Repeated_Utility_ForceBatchCall_calls(in any) []*pbgear.VaraRuntimeRuntimeCall {
    items := in.([]interface{})

    var out []*pbgear.VaraRuntimeRuntimeCall
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_VaraRuntimeRuntimeCall(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_Utility_Origins(in any) *pbgear.Utility_Origins {
    out := &pbgear.Utility_Origins{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Utility_Origins_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Utility_Origins_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Utility_Origins_Value0StakingAdmin{
        Value0StakingAdmin: To_StakingAdmin(param),
        }
    case 1:
        return &pbgear.Utility_Origins_Value0Treasurer{
        Value0Treasurer: To_Treasurer(param),
        }
    case 2:
        return &pbgear.Utility_Origins_Value0FellowshipAdmin{
        Value0FellowshipAdmin: To_FellowshipAdmin(param),
        }
    case 3:
        return &pbgear.Utility_Origins_Value0GeneralAdmin{
        Value0GeneralAdmin: To_GeneralAdmin(param),
        }
    case 4:
        return &pbgear.Utility_Origins_Value0ReferendumCanceller{
        Value0ReferendumCanceller: To_ReferendumCanceller(param),
        }
    case 5:
        return &pbgear.Utility_Origins_Value0ReferendumKiller{
        Value0ReferendumKiller: To_ReferendumKiller(param),
        }
    case 6:
        return &pbgear.Utility_Origins_Value0SmallTipper{
        Value0SmallTipper: To_SmallTipper(param),
        }
    case 7:
        return &pbgear.Utility_Origins_Value0BigTipper{
        Value0BigTipper: To_BigTipper(param),
        }
    case 8:
        return &pbgear.Utility_Origins_Value0SmallSpender{
        Value0SmallSpender: To_SmallSpender(param),
        }
    case 9:
        return &pbgear.Utility_Origins_Value0MediumSpender{
        Value0MediumSpender: To_MediumSpender(param),
        }
    case 10:
        return &pbgear.Utility_Origins_Value0BigSpender{
        Value0BigSpender: To_BigSpender(param),
        }
    case 11:
        return &pbgear.Utility_Origins_Value0WhitelistedCaller{
        Value0WhitelistedCaller: To_WhitelistedCaller(param),
        }
    case 12:
        return &pbgear.Utility_Origins_Value0FellowshipInitiates{
        Value0FellowshipInitiates: To_FellowshipInitiates(param),
        }
    case 13:
        return &pbgear.Utility_Origins_Value0Fellows{
        Value0Fellows: To_Fellows(param),
        }
    case 14:
        return &pbgear.Utility_Origins_Value0FellowshipExperts{
        Value0FellowshipExperts: To_FellowshipExperts(param),
        }
    case 15:
        return &pbgear.Utility_Origins_Value0FellowshipMasters{
        Value0FellowshipMasters: To_FellowshipMasters(param),
        }
    case 16:
        return &pbgear.Utility_Origins_Value0Fellowship1Dan{
        Value0Fellowship1Dan: To_Fellowship1Dan(param),
        }
    case 17:
        return &pbgear.Utility_Origins_Value0Fellowship2Dan{
        Value0Fellowship2Dan: To_Fellowship2Dan(param),
        }
    case 18:
        return &pbgear.Utility_Origins_Value0Fellowship3Dan{
        Value0Fellowship3Dan: To_Fellowship3Dan(param),
        }
    case 19:
        return &pbgear.Utility_Origins_Value0Fellowship4Dan{
        Value0Fellowship4Dan: To_Fellowship4Dan(param),
        }
    case 20:
        return &pbgear.Utility_Origins_Value0Fellowship5Dan{
        Value0Fellowship5Dan: To_Fellowship5Dan(param),
        }
    case 21:
        return &pbgear.Utility_Origins_Value0Fellowship6Dan{
        Value0Fellowship6Dan: To_Fellowship6Dan(param),
        }
    case 22:
        return &pbgear.Utility_Origins_Value0Fellowship7Dan{
        Value0Fellowship7Dan: To_Fellowship7Dan(param),
        }
    case 23:
        return &pbgear.Utility_Origins_Value0Fellowship8Dan{
        Value0Fellowship8Dan: To_Fellowship8Dan(param),
        }
    case 24:
        return &pbgear.Utility_Origins_Value0Fellowship9Dan{
        Value0Fellowship9Dan: To_Fellowship9Dan(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Utility_System(in any) *pbgear.Utility_System {
    out := &pbgear.Utility_System{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Utility_System_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Utility_System_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Utility_System_Value0Root{
        Value0Root: To_Root(param),
        }
    case 1:
        return &pbgear.Utility_System_Value0Signed{
        Value0Signed: To_Signed(param),
        }
    case 2:
        return &pbgear.Utility_System_Value0None{
        Value0None: To_None(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Utility_Void(in any) *pbgear.Utility_Void {
    out := &pbgear.Utility_Void{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Utility_Void_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Utility_Void_value0 (in any) any  {
    return nil
}
func To_Utility_WithWeightCall(in any) *pbgear.Utility_WithWeightCall {
    out := &pbgear.Utility_WithWeightCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_Utility_WithWeightCall_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))
    // field Weight To_SpWeightsWeightV2Weight(w)
    out.Weight = To_SpWeightsWeightV2Weight(v.ValueAt(1))

    return out //from message
}
func To_Utility_WithWeightCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Utility_WithWeightCall(in)
    out := &pbgearextrinsic.RuntimeCall_UtilityWithWeightCall{ }
    reflect.ValueOf(out).Elem().FieldByName("UtilityWithWeightCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Utility_WithWeightCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Utility_WithWeightCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Utility_WithWeightCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Utility_WithWeightCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Utility_WithWeightCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Utility_WithWeightCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Utility_WithWeightCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Utility_WithWeightCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Utility_WithWeightCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Utility_WithWeightCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Utility_WithWeightCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Utility_WithWeightCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Utility_WithWeightCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Utility_WithWeightCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Utility_WithWeightCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Utility_WithWeightCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Utility_WithWeightCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Utility_WithWeightCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Utility_WithWeightCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Utility_WithWeightCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Utility_WithWeightCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Utility_WithWeightCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Utility_WithWeightCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Utility_WithWeightCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Utility_WithWeightCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Utility_WithWeightCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Utility_WithWeightCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Utility_WithWeightCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Utility_WithWeightCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Utility_WithWeightCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Value0(in any) *pbgear.Value0 {
    out := &pbgear.Value0{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value0
    v0 := To_OneOf_Value0_value0(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value0").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Value0_value0 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Value0_Value0FundsUnavailable{
        Value0FundsUnavailable: To_FundsUnavailable(param),
        }
    case 1:
        return &pbgear.Value0_Value0OnlyProvider{
        Value0OnlyProvider: To_OnlyProvider(param),
        }
    case 2:
        return &pbgear.Value0_Value0BelowMinimum{
        Value0BelowMinimum: To_BelowMinimum(param),
        }
    case 3:
        return &pbgear.Value0_Value0CannotCreate{
        Value0CannotCreate: To_CannotCreate(param),
        }
    case 4:
        return &pbgear.Value0_Value0UnknownAsset{
        Value0UnknownAsset: To_UnknownAsset(param),
        }
    case 5:
        return &pbgear.Value0_Value0Frozen{
        Value0Frozen: To_Frozen(param),
        }
    case 6:
        return &pbgear.Value0_Value0Unsupported{
        Value0Unsupported: To_Unsupported(param),
        }
    case 7:
        return &pbgear.Value0_Value0CannotCreateHold{
        Value0CannotCreateHold: To_CannotCreateHold(param),
        }
    case 8:
        return &pbgear.Value0_Value0NotExpendable{
        Value0NotExpendable: To_NotExpendable(param),
        }
    case 9:
        return &pbgear.Value0_Value0Blocked{
        Value0Blocked: To_Blocked(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Value1(in any) *pbgear.Value1 {
    out := &pbgear.Value1{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Value1
    v0 := To_OneOf_Value1_value1(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Value1").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Value1_value1 (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Value1_Value1Success{
        Value1Success: To_Success(param),
        }
    case 1:
        return &pbgear.Value1_Value1Failed{
        Value1Failed: To_Failed(param),
        }
    case 2:
        return &pbgear.Value1_Value1NotExecuted{
        Value1NotExecuted: To_NotExecuted(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_VaraRuntimeNposSolution16(in any) *pbgear.VaraRuntimeNposSolution16 {
    out := &pbgear.VaraRuntimeNposSolution16{}
    v := in.(registry.Valuable)
    _ = v

    // repeated field Votes1
    out.Votes1 = To_Repeated_VaraRuntimeNposSolution16_votes1(v.ValueAt(0))
    // repeated field Votes2
    out.Votes2 = To_Repeated_VaraRuntimeNposSolution16_votes2(v.ValueAt(1))
    // repeated field Votes3
    out.Votes3 = To_Repeated_VaraRuntimeNposSolution16_votes3(v.ValueAt(2))
    // repeated field Votes4
    out.Votes4 = To_Repeated_VaraRuntimeNposSolution16_votes4(v.ValueAt(3))
    // repeated field Votes5
    out.Votes5 = To_Repeated_VaraRuntimeNposSolution16_votes5(v.ValueAt(4))
    // repeated field Votes6
    out.Votes6 = To_Repeated_VaraRuntimeNposSolution16_votes6(v.ValueAt(5))
    // repeated field Votes7
    out.Votes7 = To_Repeated_VaraRuntimeNposSolution16_votes7(v.ValueAt(6))
    // repeated field Votes8
    out.Votes8 = To_Repeated_VaraRuntimeNposSolution16_votes8(v.ValueAt(7))
    // repeated field Votes9
    out.Votes9 = To_Repeated_VaraRuntimeNposSolution16_votes9(v.ValueAt(8))
    // repeated field Votes10
    out.Votes10 = To_Repeated_VaraRuntimeNposSolution16_votes10(v.ValueAt(9))
    // repeated field Votes11
    out.Votes11 = To_Repeated_VaraRuntimeNposSolution16_votes11(v.ValueAt(10))
    // repeated field Votes12
    out.Votes12 = To_Repeated_VaraRuntimeNposSolution16_votes12(v.ValueAt(11))
    // repeated field Votes13
    out.Votes13 = To_Repeated_VaraRuntimeNposSolution16_votes13(v.ValueAt(12))
    // repeated field Votes14
    out.Votes14 = To_Repeated_VaraRuntimeNposSolution16_votes14(v.ValueAt(13))
    // repeated field Votes15
    out.Votes15 = To_Repeated_VaraRuntimeNposSolution16_votes15(v.ValueAt(14))
    // repeated field Votes16
    out.Votes16 = To_Repeated_VaraRuntimeNposSolution16_votes16(v.ValueAt(15))

    return out //from message
}
    

func To_Repeated_VaraRuntimeNposSolution16_votes1(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32Uint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32Uint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32Uint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes2(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32TupleUint32SpArithmeticPerThingsPerU16Uint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes3(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes3Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes4(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes4Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes5(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes5Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes6(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes6Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes7(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes7Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes8(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes8Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes9(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes9Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes10(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes10Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes11(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes11Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes12(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes12Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes13(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes13Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes14(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes14Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes15(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes15Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}

func To_Repeated_VaraRuntimeNposSolution16_votes16(in any) []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32 {
    items := in.([]interface{})

    var out []*pbgear.ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32
    for _, item := range items {
		w := item
		if i, ok := w.(registry.Valuable); ok {//ugly ack the will bite later
			w = &Wrap{Value: i}
		}
        o := To_ElectionProviderMultiPhase_TupleUint32ElectionProviderMultiPhaseVotes16Listuint32(w)
        out = append(out, o)
    }
    return nil //funcForRepeatedField
}
func To_VaraRuntimeRuntimeCall(in any) *pbgear.VaraRuntimeRuntimeCall {
    out := &pbgear.VaraRuntimeRuntimeCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Calls
    v0 := To_OneOf_VaraRuntimeRuntimeCall_calls(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Calls").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_VaraRuntimeRuntimeCall_calls (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.VaraRuntimeRuntimeCall_CallsSystem{
        CallsSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.VaraRuntimeRuntimeCall_CallsTimestamp{
        CallsTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.VaraRuntimeRuntimeCall_CallsBabe{
        CallsBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.VaraRuntimeRuntimeCall_CallsGrandpa{
        CallsGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.VaraRuntimeRuntimeCall_CallsBalances{
        CallsBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.VaraRuntimeRuntimeCall_CallsVesting{
        CallsVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.VaraRuntimeRuntimeCall_CallsBagsList{
        CallsBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.VaraRuntimeRuntimeCall_CallsImOnline{
        CallsImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.VaraRuntimeRuntimeCall_CallsStaking{
        CallsStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.VaraRuntimeRuntimeCall_CallsSession{
        CallsSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.VaraRuntimeRuntimeCall_CallsTreasury{
        CallsTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.VaraRuntimeRuntimeCall_CallsUtility{
        CallsUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.VaraRuntimeRuntimeCall_CallsConvictionVoting{
        CallsConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.VaraRuntimeRuntimeCall_CallsReferenda{
        CallsReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.VaraRuntimeRuntimeCall_CallsFellowshipCollective{
        CallsFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.VaraRuntimeRuntimeCall_CallsFellowshipReferenda{
        CallsFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.VaraRuntimeRuntimeCall_CallsWhitelist{
        CallsWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.VaraRuntimeRuntimeCall_CallsScheduler{
        CallsScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.VaraRuntimeRuntimeCall_CallsPreimage{
        CallsPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.VaraRuntimeRuntimeCall_CallsIdentity{
        CallsIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.VaraRuntimeRuntimeCall_CallsProxy{
        CallsProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.VaraRuntimeRuntimeCall_CallsMultisig{
        CallsMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.VaraRuntimeRuntimeCall_CallsElectionProviderMultiPhase{
        CallsElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.VaraRuntimeRuntimeCall_CallsBounties{
        CallsBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.VaraRuntimeRuntimeCall_CallsChildBounties{
        CallsChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.VaraRuntimeRuntimeCall_CallsNominationPools{
        CallsNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.VaraRuntimeRuntimeCall_CallsGear{
        CallsGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.VaraRuntimeRuntimeCall_CallsStakingRewards{
        CallsStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.VaraRuntimeRuntimeCall_CallsGearVoucher{
        CallsGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_VaraRuntimeSessionKeys(in any) *pbgear.VaraRuntimeSessionKeys {
    out := &pbgear.VaraRuntimeSessionKeys{}
    v := in.(registry.Valuable)
    _ = v

    // field Babe To_SpConsensusBabeAppPublic(w)
    out.Babe = To_SpConsensusBabeAppPublic(v.ValueAt(0))
    // field Grandpa To_SpConsensusGrandpaAppPublic(w)
    out.Grandpa = To_SpConsensusGrandpaAppPublic(v.ValueAt(1))
    // field ImOnline To_ImOnline_PalletImOnlineSr25519AppSr25519Public(w)
    out.ImOnline = To_ImOnline_PalletImOnlineSr25519AppSr25519Public(v.ValueAt(2))
    // field AuthorityDiscovery To_SpAuthorityDiscoveryAppPublic(w)
    out.AuthorityDiscovery = To_SpAuthorityDiscoveryAppPublic(v.ValueAt(3))

    return out //from message
}
    




func To_VestingPallet(in any) *pbgear.VestingPallet {
    out := &pbgear.VestingPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_VestingPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_VestingPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.VestingPallet_CallVestCall{
        CallVestCall: To_Vesting_VestCall(param),
        }
    case 1:
        return &pbgear.VestingPallet_CallVestOtherCall{
        CallVestOtherCall: To_Vesting_VestOtherCall(param),
        }
    case 2:
        return &pbgear.VestingPallet_CallVestedTransferCall{
        CallVestedTransferCall: To_Vesting_VestedTransferCall(param),
        }
    case 3:
        return &pbgear.VestingPallet_CallForceVestedTransferCall{
        CallForceVestedTransferCall: To_Vesting_ForceVestedTransferCall(param),
        }
    case 4:
        return &pbgear.VestingPallet_CallMergeSchedulesCall{
        CallMergeSchedulesCall: To_Vesting_MergeSchedulesCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_VestingVestingCompletedEvent(in any) *pbgear.VestingVestingCompletedEvent {
    out := &pbgear.VestingVestingCompletedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_VestingVestingUpdatedEvent(in any) *pbgear.VestingVestingUpdatedEvent {
    out := &pbgear.VestingVestingUpdatedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Vesting_Address20(in any) *pbgear.Vesting_Address20 {
    out := &pbgear.Vesting_Address20{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Vesting_Address32(in any) *pbgear.Vesting_Address32 {
    out := &pbgear.Vesting_Address32{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Vesting_ForceVestedTransferCall(in any) *pbgear.Vesting_ForceVestedTransferCall {
    out := &pbgear.Vesting_ForceVestedTransferCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Source
    v0 := To_OneOf_Vesting_ForceVestedTransferCall_source(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))
    // oneOf field Target
    v1 := To_OneOf_Vesting_ForceVestedTransferCall_target(v.ValueAt(1))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v1))
    // field Schedule To_Vesting_PalletVestingVestingInfoVestingInfo(w)
    out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(v.ValueAt(2))

    return out //from message
}
func To_Vesting_ForceVestedTransferCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Vesting_ForceVestedTransferCall(in)
    out := &pbgearextrinsic.RuntimeCall_VestingForceVestedTransferCall{ }
    reflect.ValueOf(out).Elem().FieldByName("VestingForceVestedTransferCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Vesting_ForceVestedTransferCall_source (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Vesting_ForceVestedTransferCall_SourceId{
        SourceId: To_Vesting_Id(param),
        }
    case 1:
        return &pbgear.Vesting_ForceVestedTransferCall_SourceIndex{
        SourceIndex: To_Vesting_Index(param),
        }
    case 2:
        return &pbgear.Vesting_ForceVestedTransferCall_SourceRaw{
        SourceRaw: To_Vesting_Raw(param),
        }
    case 3:
        return &pbgear.Vesting_ForceVestedTransferCall_SourceAddress32{
        SourceAddress32: To_Vesting_Address32(param),
        }
    case 4:
        return &pbgear.Vesting_ForceVestedTransferCall_SourceAddress20{
        SourceAddress20: To_Vesting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_OneOf_Vesting_ForceVestedTransferCall_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Vesting_ForceVestedTransferCall_TargetId{
        TargetId: To_Vesting_Id(param),
        }
    case 1:
        return &pbgear.Vesting_ForceVestedTransferCall_TargetIndex{
        TargetIndex: To_Vesting_Index(param),
        }
    case 2:
        return &pbgear.Vesting_ForceVestedTransferCall_TargetRaw{
        TargetRaw: To_Vesting_Raw(param),
        }
    case 3:
        return &pbgear.Vesting_ForceVestedTransferCall_TargetAddress32{
        TargetAddress32: To_Vesting_Address32(param),
        }
    case 4:
        return &pbgear.Vesting_ForceVestedTransferCall_TargetAddress20{
        TargetAddress20: To_Vesting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_Vesting_Id(in any) *pbgear.Vesting_Id {
    out := &pbgear.Vesting_Id{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_SpCoreCryptoAccountId32(w)
    out.Value0 = To_SpCoreCryptoAccountId32(v.ValueAt(0))

    return out //from message
}
    

func To_Vesting_Index(in any) *pbgear.Vesting_Index {
    out := &pbgear.Vesting_Index{}
    v := in.(registry.Valuable)
    _ = v

    // field Value0 To_Vesting_TupleNull(w)
    out.Value0 = To_Vesting_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Vesting_MergeSchedulesCall(in any) *pbgear.Vesting_MergeSchedulesCall {
    out := &pbgear.Vesting_MergeSchedulesCall{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Schedule1Index
    out.Schedule1Index = To_uint32(v.ValueAt(0))
    // primitive field Schedule2Index
    out.Schedule2Index = To_uint32(v.ValueAt(1))

    return out //from message
}
func To_Vesting_MergeSchedulesCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Vesting_MergeSchedulesCall(in)
    out := &pbgearextrinsic.RuntimeCall_VestingMergeSchedulesCall{ }
    reflect.ValueOf(out).Elem().FieldByName("VestingMergeSchedulesCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_Vesting_PalletVestingVestingInfoVestingInfo(in any) *pbgear.Vesting_PalletVestingVestingInfoVestingInfo {
    out := &pbgear.Vesting_PalletVestingVestingInfoVestingInfo{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Locked
    out.Locked = To_string(v.ValueAt(0))
    // primitive field PerBlock
    out.PerBlock = To_string(v.ValueAt(1))
    // primitive field StartingBlock
    out.StartingBlock = To_uint32(v.ValueAt(2))

    return out //from message
}
    
func To_Vesting_Raw(in any) *pbgear.Vesting_Raw {
    out := &pbgear.Vesting_Raw{}
    v := in.(registry.Valuable)
    _ = v

    // primitive field Value0
    out.Value0 = To_bytes(v.ValueAt(0))

    return out //from message
}
    
func To_Vesting_Source(in any) *pbgear.Vesting_Source {
    out := &pbgear.Vesting_Source{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Source
    v0 := To_OneOf_Vesting_Source_source(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Source").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Vesting_Source_source (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Vesting_Source_SourceId{
        SourceId: To_Vesting_Id(param),
        }
    case 1:
        return &pbgear.Vesting_Source_SourceIndex{
        SourceIndex: To_Vesting_Index(param),
        }
    case 2:
        return &pbgear.Vesting_Source_SourceRaw{
        SourceRaw: To_Vesting_Raw(param),
        }
    case 3:
        return &pbgear.Vesting_Source_SourceAddress32{
        SourceAddress32: To_Vesting_Address32(param),
        }
    case 4:
        return &pbgear.Vesting_Source_SourceAddress20{
        SourceAddress20: To_Vesting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Vesting_Target(in any) *pbgear.Vesting_Target {
    out := &pbgear.Vesting_Target{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Target
    v0 := To_OneOf_Vesting_Target_target(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_Vesting_Target_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Vesting_Target_TargetId{
        TargetId: To_Vesting_Id(param),
        }
    case 1:
        return &pbgear.Vesting_Target_TargetIndex{
        TargetIndex: To_Vesting_Index(param),
        }
    case 2:
        return &pbgear.Vesting_Target_TargetRaw{
        TargetRaw: To_Vesting_Raw(param),
        }
    case 3:
        return &pbgear.Vesting_Target_TargetAddress32{
        TargetAddress32: To_Vesting_Address32(param),
        }
    case 4:
        return &pbgear.Vesting_Target_TargetAddress20{
        TargetAddress20: To_Vesting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Vesting_TupleNull(in any) *pbgear.Vesting_TupleNull {
    out := &pbgear.Vesting_TupleNull{}
    v := in.(registry.Valuable)
    _ = v

    // field Value To_TupleNull(w)
    out.Value = To_TupleNull(v.ValueAt(0))

    return out //from message
}
    

func To_Vesting_VestCall(in any) *pbgear.Vesting_VestCall {
    out := &pbgear.Vesting_VestCall{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Vesting_VestOtherCall(in any) *pbgear.Vesting_VestOtherCall {
    out := &pbgear.Vesting_VestOtherCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Target
    v0 := To_OneOf_Vesting_VestOtherCall_target(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_Vesting_VestOtherCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Vesting_VestOtherCall(in)
    out := &pbgearextrinsic.RuntimeCall_VestingVestOtherCall{ }
    reflect.ValueOf(out).Elem().FieldByName("VestingVestOtherCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Vesting_VestOtherCall_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Vesting_VestOtherCall_TargetId{
        TargetId: To_Vesting_Id(param),
        }
    case 1:
        return &pbgear.Vesting_VestOtherCall_TargetIndex{
        TargetIndex: To_Vesting_Index(param),
        }
    case 2:
        return &pbgear.Vesting_VestOtherCall_TargetRaw{
        TargetRaw: To_Vesting_Raw(param),
        }
    case 3:
        return &pbgear.Vesting_VestOtherCall_TargetAddress32{
        TargetAddress32: To_Vesting_Address32(param),
        }
    case 4:
        return &pbgear.Vesting_VestOtherCall_TargetAddress20{
        TargetAddress20: To_Vesting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Vesting_VestedTransferCall(in any) *pbgear.Vesting_VestedTransferCall {
    out := &pbgear.Vesting_VestedTransferCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Target
    v0 := To_OneOf_Vesting_VestedTransferCall_target(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Target").Set(reflect.ValueOf(v0))
    // field Schedule To_Vesting_PalletVestingVestingInfoVestingInfo(w)
    out.Schedule = To_Vesting_PalletVestingVestingInfoVestingInfo(v.ValueAt(1))

    return out //from message
}
func To_Vesting_VestedTransferCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Vesting_VestedTransferCall(in)
    out := &pbgearextrinsic.RuntimeCall_VestingVestedTransferCall{ }
    reflect.ValueOf(out).Elem().FieldByName("VestingVestedTransferCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Vesting_VestedTransferCall_target (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Vesting_VestedTransferCall_TargetId{
        TargetId: To_Vesting_Id(param),
        }
    case 1:
        return &pbgear.Vesting_VestedTransferCall_TargetIndex{
        TargetIndex: To_Vesting_Index(param),
        }
    case 2:
        return &pbgear.Vesting_VestedTransferCall_TargetRaw{
        TargetRaw: To_Vesting_Raw(param),
        }
    case 3:
        return &pbgear.Vesting_VestedTransferCall_TargetAddress32{
        TargetAddress32: To_Vesting_Address32(param),
        }
    case 4:
        return &pbgear.Vesting_VestedTransferCall_TargetAddress20{
        TargetAddress20: To_Vesting_Address20(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}

func To_WaitCalled(in any) *pbgear.WaitCalled {
    out := &pbgear.WaitCalled{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_WaitForCalled(in any) *pbgear.WaitForCalled {
    out := &pbgear.WaitForCalled{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_WaitUpToCalled(in any) *pbgear.WaitUpToCalled {
    out := &pbgear.WaitUpToCalled{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_WaitUpToCalledFull(in any) *pbgear.WaitUpToCalledFull {
    out := &pbgear.WaitUpToCalledFull{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_WakeCalled(in any) *pbgear.WakeCalled {
    out := &pbgear.WakeCalled{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_WhitelistCallWhitelistedEvent(in any) *pbgear.WhitelistCallWhitelistedEvent {
    out := &pbgear.WhitelistCallWhitelistedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_WhitelistPallet(in any) *pbgear.WhitelistPallet {
    out := &pbgear.WhitelistPallet{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_WhitelistPallet_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
    
func To_OneOf_WhitelistPallet_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.WhitelistPallet_CallWhitelistCallCall{
        CallWhitelistCallCall: To_Whitelist_WhitelistCallCall(param),
        }
    case 1:
        return &pbgear.WhitelistPallet_CallRemoveWhitelistedCallCall{
        CallRemoveWhitelistedCallCall: To_Whitelist_RemoveWhitelistedCallCall(param),
        }
    case 2:
        return &pbgear.WhitelistPallet_CallDispatchWhitelistedCallCall{
        CallDispatchWhitelistedCallCall: To_Whitelist_DispatchWhitelistedCallCall(param),
        }
    case 3:
        return &pbgear.WhitelistPallet_CallDispatchWhitelistedCallWithPreimageCall{
        CallDispatchWhitelistedCallWithPreimageCall: To_Whitelist_DispatchWhitelistedCallWithPreimageCall(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_WhitelistWhitelistedCallDispatchedEvent(in any) *pbgear.WhitelistWhitelistedCallDispatchedEvent {
    out := &pbgear.WhitelistWhitelistedCallDispatchedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_WhitelistWhitelistedCallRemovedEvent(in any) *pbgear.WhitelistWhitelistedCallRemovedEvent {
    out := &pbgear.WhitelistWhitelistedCallRemovedEvent{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Whitelist_DispatchWhitelistedCallCall(in any) *pbgear.Whitelist_DispatchWhitelistedCallCall {
    out := &pbgear.Whitelist_DispatchWhitelistedCallCall{}
    v := in.(registry.Valuable)
    _ = v

    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))
    // primitive field CallEncodedLen
    out.CallEncodedLen = To_uint32(v.ValueAt(1))
    // field CallWeightWitness To_SpWeightsWeightV2Weight(w)
    out.CallWeightWitness = To_SpWeightsWeightV2Weight(v.ValueAt(2))

    return out //from message
}
func To_Whitelist_DispatchWhitelistedCallCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Whitelist_DispatchWhitelistedCallCall(in)
    out := &pbgearextrinsic.RuntimeCall_WhitelistDispatchWhitelistedCallCall{ }
    reflect.ValueOf(out).Elem().FieldByName("WhitelistDispatchWhitelistedCallCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    


func To_Whitelist_DispatchWhitelistedCallWithPreimageCall(in any) *pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall {
    out := &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall{}
    v := in.(registry.Valuable)
    _ = v

    // oneOf field Call
    v0 := To_OneOf_Whitelist_DispatchWhitelistedCallWithPreimageCall_call(v.ValueAt(0))
    reflect.ValueOf(out).Elem().FieldByName("Call").Set(reflect.ValueOf(v0))

    return out //from message
}
func To_Whitelist_DispatchWhitelistedCallWithPreimageCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Whitelist_DispatchWhitelistedCallWithPreimageCall(in)
    out := &pbgearextrinsic.RuntimeCall_WhitelistDispatchWhitelistedCallWithPreimageCall{ }
    reflect.ValueOf(out).Elem().FieldByName("WhitelistDispatchWhitelistedCallWithPreimageCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    
func To_OneOf_Whitelist_DispatchWhitelistedCallWithPreimageCall_call (in any) any  {
    variantIn := in.(*registry.VariantWTF)
    param := variantIn.Value
    switch variantIn.VariantByte {
    case 0:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallSystem{
        CallSystem: To_SystemPallet(param),
        }
    case 1:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallTimestamp{
        CallTimestamp: To_TimestampPallet(param),
        }
    case 3:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBabe{
        CallBabe: To_BabePallet(param),
        }
    case 4:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallGrandpa{
        CallGrandpa: To_GrandpaPallet(param),
        }
    case 5:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBalances{
        CallBalances: To_BalancesPallet(param),
        }
    case 10:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallVesting{
        CallVesting: To_VestingPallet(param),
        }
    case 11:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBagsList{
        CallBagsList: To_BagsListPallet(param),
        }
    case 12:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallImOnline{
        CallImOnline: To_ImOnlinePallet(param),
        }
    case 13:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallStaking{
        CallStaking: To_StakingPallet(param),
        }
    case 7:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallSession{
        CallSession: To_SessionPallet(param),
        }
    case 14:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallTreasury{
        CallTreasury: To_TreasuryPallet(param),
        }
    case 8:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallUtility{
        CallUtility: To_UtilityPallet(param),
        }
    case 16:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallConvictionVoting{
        CallConvictionVoting: To_ConvictionVotingPallet(param),
        }
    case 17:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallReferenda{
        CallReferenda: To_ReferendaPallet(param),
        }
    case 18:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallFellowshipCollective{
        CallFellowshipCollective: To_FellowshipCollectivePallet(param),
        }
    case 19:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallFellowshipReferenda{
        CallFellowshipReferenda: To_FellowshipReferendaPallet(param),
        }
    case 21:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallWhitelist{
        CallWhitelist: To_WhitelistPallet(param),
        }
    case 22:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallScheduler{
        CallScheduler: To_SchedulerPallet(param),
        }
    case 23:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallPreimage{
        CallPreimage: To_PreimagePallet(param),
        }
    case 24:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallIdentity{
        CallIdentity: To_IdentityPallet(param),
        }
    case 25:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallProxy{
        CallProxy: To_ProxyPallet(param),
        }
    case 26:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallMultisig{
        CallMultisig: To_MultisigPallet(param),
        }
    case 27:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallElectionProviderMultiPhase{
        CallElectionProviderMultiPhase: To_ElectionProviderMultiPhasePallet(param),
        }
    case 29:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallBounties{
        CallBounties: To_BountiesPallet(param),
        }
    case 30:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallChildBounties{
        CallChildBounties: To_ChildBountiesPallet(param),
        }
    case 31:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallNominationPools{
        CallNominationPools: To_NominationPoolsPallet(param),
        }
    case 104:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallGear{
        CallGear: To_GearPallet(param),
        }
    case 106:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallStakingRewards{
        CallStakingRewards: To_StakingRewardsPallet(param),
        }
    case 107:
        return &pbgear.Whitelist_DispatchWhitelistedCallWithPreimageCall_CallGearVoucher{
        CallGearVoucher: To_GearVoucherPallet(param),
        }
    default:
        panic(fmt.Sprintf("Unknown variant byte: %d", variantIn.VariantByte))
    }
}
func To_Whitelist_RemoveWhitelistedCallCall(in any) *pbgear.Whitelist_RemoveWhitelistedCallCall {
    out := &pbgear.Whitelist_RemoveWhitelistedCallCall{}
    v := in.(registry.Valuable)
    _ = v

    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
func To_Whitelist_RemoveWhitelistedCallCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Whitelist_RemoveWhitelistedCallCall(in)
    out := &pbgearextrinsic.RuntimeCall_WhitelistRemoveWhitelistedCallCall{ }
    reflect.ValueOf(out).Elem().FieldByName("WhitelistRemoveWhitelistedCallCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_Whitelist_WhitelistCallCall(in any) *pbgear.Whitelist_WhitelistCallCall {
    out := &pbgear.Whitelist_WhitelistCallCall{}
    v := in.(registry.Valuable)
    _ = v

    // field CallHash To_PrimitiveTypesH256(w)
    out.CallHash = To_PrimitiveTypesH256(v.ValueAt(0))

    return out //from message
}
func To_Whitelist_WhitelistCallCall_wrap(in any) *pbgearextrinsic.RuntimeCall{
    res := To_Whitelist_WhitelistCallCall(in)
    out := &pbgearextrinsic.RuntimeCall_WhitelistWhitelistCallCall{ }
    reflect.ValueOf(out).Elem().FieldByName("WhitelistWhitelistCallCall").Set(reflect.ValueOf(res))
    return &pbgearextrinsic.RuntimeCall{
        Call: out,
    }
}
    

func To_WhitelistedCaller(in any) *pbgear.WhitelistedCaller {
    out := &pbgear.WhitelistedCaller{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    
func To_Yes(in any) *pbgear.Yes {
    out := &pbgear.Yes{}
    v := in.(registry.Valuable)
    _ = v


    return out //from message
}
    













type Wrap struct {
	Value registry.Valuable
}

func (w Wrap) ValueAt(index int) any {
    return w.Value
}

func To_bytes(a any) []byte {
	var items []any
	switch in := a.(type) {
	case registry.DecodedFields:
		fields := []*registry.DecodedField(in)
		items = fields[0].Value.([]any)
	case []any:
		items = in
	default:
		panic(fmt.Sprintf("Unsupported type %T", in))
	}
	var out []byte
	for _, i := range items {
		u := uint8(i.(types.U8))
		out = append(out, byte(u))
	}
	return out
}

func To_string(i any) string {
	switch v := i.(type) {
	case types.UCompact:
		return fmt.Sprintf("%d", v.Int64())
	case types.Text:
		return string(v)
	case types.U128:
		return v.String()
	case types.U256:
		return v.String()
	case types.I128:
		return v.String()
	case types.I256:
		return v.String()
	default:
		panic("Unknown type")
	}
}

func To_uint32(i any) uint32 {
    switch v := i.(type) {
    case types.UCompact:
        return uint32(v.Int64())
    case types.U8:
        return uint32(v)
    case types.U16:
        return uint32(v)
    case types.U32:
        return uint32(v)
    default:
        panic("Unknown type")
    }
}

func To_Optional_uint32(i any) *uint32 {
    o := To_uint32(i)
    return &o
}

func To_uint64(a any) uint64 {
    switch v := a.(type) {
    case types.UCompact:
        return uint64(v.Int64())
    case types.U8:
        return uint64(v)
    case types.U16:
        return uint64(v)
    case types.U32:
        return uint64(v)
    case types.U64:
        return uint64(v)
    case uint8:
        return uint64(v)
    case uint16:
        return uint64(v)
    case uint32:
        return uint64(v)
    case uint64:
        return uint64(v)
    default:
        panic("Unknown type")
    }
}

func To_Optional_uint64(i any) *uint64 {
    o := To_uint64(i)
    return &o
}

func To_Optional_string(a any) *string {
    s := To_string(a)
    return &s
}

func To_bool(b any) bool {
    return b.(bool)
}

func To_Optional_bool(b any) *bool {
    o :=  To_bool(b)
    return &o
}

func To_Repeated_uint32(a any) []uint32 {
    panic("To_repeated_uint32")
}

var FuncMap map[string]reflect.Value

func init(){
    FuncMap = make(map[string]reflect.Value)
    FuncMap["To_Babe_PlanConfigChangeCall"] = reflect.ValueOf(To_Babe_PlanConfigChangeCall_wrap)
    FuncMap["To_Babe_ReportEquivocationCall"] = reflect.ValueOf(To_Babe_ReportEquivocationCall_wrap)
    FuncMap["To_Babe_ReportEquivocationUnsignedCall"] = reflect.ValueOf(To_Babe_ReportEquivocationUnsignedCall_wrap)
    FuncMap["To_BagsList_PutInFrontOfCall"] = reflect.ValueOf(To_BagsList_PutInFrontOfCall_wrap)
    FuncMap["To_BagsList_PutInFrontOfOtherCall"] = reflect.ValueOf(To_BagsList_PutInFrontOfOtherCall_wrap)
    FuncMap["To_BagsList_RebagCall"] = reflect.ValueOf(To_BagsList_RebagCall_wrap)
    FuncMap["To_Balances_ForceSetBalanceCall"] = reflect.ValueOf(To_Balances_ForceSetBalanceCall_wrap)
    FuncMap["To_Balances_ForceTransferCall"] = reflect.ValueOf(To_Balances_ForceTransferCall_wrap)
    FuncMap["To_Balances_ForceUnreserveCall"] = reflect.ValueOf(To_Balances_ForceUnreserveCall_wrap)
    FuncMap["To_Balances_TransferAllCall"] = reflect.ValueOf(To_Balances_TransferAllCall_wrap)
    FuncMap["To_Balances_TransferAllowDeathCall"] = reflect.ValueOf(To_Balances_TransferAllowDeathCall_wrap)
    FuncMap["To_Balances_TransferKeepAliveCall"] = reflect.ValueOf(To_Balances_TransferKeepAliveCall_wrap)
    FuncMap["To_Balances_UpgradeAccountsCall"] = reflect.ValueOf(To_Balances_UpgradeAccountsCall_wrap)
    FuncMap["To_Bounties_AcceptCuratorCall"] = reflect.ValueOf(To_Bounties_AcceptCuratorCall_wrap)
    FuncMap["To_Bounties_ApproveBountyCall"] = reflect.ValueOf(To_Bounties_ApproveBountyCall_wrap)
    FuncMap["To_Bounties_AwardBountyCall"] = reflect.ValueOf(To_Bounties_AwardBountyCall_wrap)
    FuncMap["To_Bounties_ClaimBountyCall"] = reflect.ValueOf(To_Bounties_ClaimBountyCall_wrap)
    FuncMap["To_Bounties_CloseBountyCall"] = reflect.ValueOf(To_Bounties_CloseBountyCall_wrap)
    FuncMap["To_Bounties_ExtendBountyExpiryCall"] = reflect.ValueOf(To_Bounties_ExtendBountyExpiryCall_wrap)
    FuncMap["To_Bounties_ProposeBountyCall"] = reflect.ValueOf(To_Bounties_ProposeBountyCall_wrap)
    FuncMap["To_Bounties_ProposeCuratorCall"] = reflect.ValueOf(To_Bounties_ProposeCuratorCall_wrap)
    FuncMap["To_Bounties_UnassignCuratorCall"] = reflect.ValueOf(To_Bounties_UnassignCuratorCall_wrap)
    FuncMap["To_ChildBounties_AcceptCuratorCall"] = reflect.ValueOf(To_ChildBounties_AcceptCuratorCall_wrap)
    FuncMap["To_ChildBounties_AddChildBountyCall"] = reflect.ValueOf(To_ChildBounties_AddChildBountyCall_wrap)
    FuncMap["To_ChildBounties_AwardChildBountyCall"] = reflect.ValueOf(To_ChildBounties_AwardChildBountyCall_wrap)
    FuncMap["To_ChildBounties_ClaimChildBountyCall"] = reflect.ValueOf(To_ChildBounties_ClaimChildBountyCall_wrap)
    FuncMap["To_ChildBounties_CloseChildBountyCall"] = reflect.ValueOf(To_ChildBounties_CloseChildBountyCall_wrap)
    FuncMap["To_ChildBounties_ProposeCuratorCall"] = reflect.ValueOf(To_ChildBounties_ProposeCuratorCall_wrap)
    FuncMap["To_ChildBounties_UnassignCuratorCall"] = reflect.ValueOf(To_ChildBounties_UnassignCuratorCall_wrap)
    FuncMap["To_ConvictionVoting_DelegateCall"] = reflect.ValueOf(To_ConvictionVoting_DelegateCall_wrap)
    FuncMap["To_ConvictionVoting_RemoveOtherVoteCall"] = reflect.ValueOf(To_ConvictionVoting_RemoveOtherVoteCall_wrap)
    FuncMap["To_ConvictionVoting_RemoveVoteCall"] = reflect.ValueOf(To_ConvictionVoting_RemoveVoteCall_wrap)
    FuncMap["To_ConvictionVoting_UndelegateCall"] = reflect.ValueOf(To_ConvictionVoting_UndelegateCall_wrap)
    FuncMap["To_ConvictionVoting_UnlockCall"] = reflect.ValueOf(To_ConvictionVoting_UnlockCall_wrap)
    FuncMap["To_ConvictionVoting_VoteCall"] = reflect.ValueOf(To_ConvictionVoting_VoteCall_wrap)
    FuncMap["To_ElectionProviderMultiPhase_GovernanceFallbackCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_GovernanceFallbackCall_wrap)
    FuncMap["To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SetEmergencyElectionResultCall_wrap)
    FuncMap["To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SetMinimumUntrustedScoreCall_wrap)
    FuncMap["To_ElectionProviderMultiPhase_SubmitCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SubmitCall_wrap)
    FuncMap["To_ElectionProviderMultiPhase_SubmitUnsignedCall"] = reflect.ValueOf(To_ElectionProviderMultiPhase_SubmitUnsignedCall_wrap)
    FuncMap["To_FellowshipCollective_AddMemberCall"] = reflect.ValueOf(To_FellowshipCollective_AddMemberCall_wrap)
    FuncMap["To_FellowshipCollective_CleanupPollCall"] = reflect.ValueOf(To_FellowshipCollective_CleanupPollCall_wrap)
    FuncMap["To_FellowshipCollective_DemoteMemberCall"] = reflect.ValueOf(To_FellowshipCollective_DemoteMemberCall_wrap)
    FuncMap["To_FellowshipCollective_PromoteMemberCall"] = reflect.ValueOf(To_FellowshipCollective_PromoteMemberCall_wrap)
    FuncMap["To_FellowshipCollective_RemoveMemberCall"] = reflect.ValueOf(To_FellowshipCollective_RemoveMemberCall_wrap)
    FuncMap["To_FellowshipCollective_VoteCall"] = reflect.ValueOf(To_FellowshipCollective_VoteCall_wrap)
    FuncMap["To_FellowshipReferenda_CancelCall"] = reflect.ValueOf(To_FellowshipReferenda_CancelCall_wrap)
    FuncMap["To_FellowshipReferenda_KillCall"] = reflect.ValueOf(To_FellowshipReferenda_KillCall_wrap)
    FuncMap["To_FellowshipReferenda_NudgeReferendumCall"] = reflect.ValueOf(To_FellowshipReferenda_NudgeReferendumCall_wrap)
    FuncMap["To_FellowshipReferenda_OneFewerDecidingCall"] = reflect.ValueOf(To_FellowshipReferenda_OneFewerDecidingCall_wrap)
    FuncMap["To_FellowshipReferenda_PlaceDecisionDepositCall"] = reflect.ValueOf(To_FellowshipReferenda_PlaceDecisionDepositCall_wrap)
    FuncMap["To_FellowshipReferenda_RefundDecisionDepositCall"] = reflect.ValueOf(To_FellowshipReferenda_RefundDecisionDepositCall_wrap)
    FuncMap["To_FellowshipReferenda_RefundSubmissionDepositCall"] = reflect.ValueOf(To_FellowshipReferenda_RefundSubmissionDepositCall_wrap)
    FuncMap["To_FellowshipReferenda_SetMetadataCall"] = reflect.ValueOf(To_FellowshipReferenda_SetMetadataCall_wrap)
    FuncMap["To_FellowshipReferenda_SubmitCall"] = reflect.ValueOf(To_FellowshipReferenda_SubmitCall_wrap)
    FuncMap["To_GearVoucher_CallCall"] = reflect.ValueOf(To_GearVoucher_CallCall_wrap)
    FuncMap["To_GearVoucher_CallDeprecatedCall"] = reflect.ValueOf(To_GearVoucher_CallDeprecatedCall_wrap)
    FuncMap["To_GearVoucher_DeclineCall"] = reflect.ValueOf(To_GearVoucher_DeclineCall_wrap)
    FuncMap["To_GearVoucher_IssueCall"] = reflect.ValueOf(To_GearVoucher_IssueCall_wrap)
    FuncMap["To_GearVoucher_RevokeCall"] = reflect.ValueOf(To_GearVoucher_RevokeCall_wrap)
    FuncMap["To_GearVoucher_UpdateCall"] = reflect.ValueOf(To_GearVoucher_UpdateCall_wrap)
    FuncMap["To_Gear_ClaimValueCall"] = reflect.ValueOf(To_Gear_ClaimValueCall_wrap)
    FuncMap["To_Gear_CreateProgramCall"] = reflect.ValueOf(To_Gear_CreateProgramCall_wrap)
    FuncMap["To_Gear_RunCall"] = reflect.ValueOf(To_Gear_RunCall_wrap)
    FuncMap["To_Gear_SendMessageCall"] = reflect.ValueOf(To_Gear_SendMessageCall_wrap)
    FuncMap["To_Gear_SendReplyCall"] = reflect.ValueOf(To_Gear_SendReplyCall_wrap)
    FuncMap["To_Gear_SetExecuteInherentCall"] = reflect.ValueOf(To_Gear_SetExecuteInherentCall_wrap)
    FuncMap["To_Gear_UploadCodeCall"] = reflect.ValueOf(To_Gear_UploadCodeCall_wrap)
    FuncMap["To_Gear_UploadProgramCall"] = reflect.ValueOf(To_Gear_UploadProgramCall_wrap)
    FuncMap["To_Grandpa_NoteStalledCall"] = reflect.ValueOf(To_Grandpa_NoteStalledCall_wrap)
    FuncMap["To_Grandpa_ReportEquivocationCall"] = reflect.ValueOf(To_Grandpa_ReportEquivocationCall_wrap)
    FuncMap["To_Grandpa_ReportEquivocationUnsignedCall"] = reflect.ValueOf(To_Grandpa_ReportEquivocationUnsignedCall_wrap)
    FuncMap["To_Identity_AddRegistrarCall"] = reflect.ValueOf(To_Identity_AddRegistrarCall_wrap)
    FuncMap["To_Identity_AddSubCall"] = reflect.ValueOf(To_Identity_AddSubCall_wrap)
    FuncMap["To_Identity_CancelRequestCall"] = reflect.ValueOf(To_Identity_CancelRequestCall_wrap)
    FuncMap["To_Identity_KillIdentityCall"] = reflect.ValueOf(To_Identity_KillIdentityCall_wrap)
    FuncMap["To_Identity_ProvideJudgementCall"] = reflect.ValueOf(To_Identity_ProvideJudgementCall_wrap)
    FuncMap["To_Identity_RemoveSubCall"] = reflect.ValueOf(To_Identity_RemoveSubCall_wrap)
    FuncMap["To_Identity_RenameSubCall"] = reflect.ValueOf(To_Identity_RenameSubCall_wrap)
    FuncMap["To_Identity_RequestJudgementCall"] = reflect.ValueOf(To_Identity_RequestJudgementCall_wrap)
    FuncMap["To_Identity_SetAccountIdCall"] = reflect.ValueOf(To_Identity_SetAccountIdCall_wrap)
    FuncMap["To_Identity_SetFeeCall"] = reflect.ValueOf(To_Identity_SetFeeCall_wrap)
    FuncMap["To_Identity_SetFieldsCall"] = reflect.ValueOf(To_Identity_SetFieldsCall_wrap)
    FuncMap["To_Identity_SetIdentityCall"] = reflect.ValueOf(To_Identity_SetIdentityCall_wrap)
    FuncMap["To_Identity_SetSubsCall"] = reflect.ValueOf(To_Identity_SetSubsCall_wrap)
    FuncMap["To_ImOnline_HeartbeatCall"] = reflect.ValueOf(To_ImOnline_HeartbeatCall_wrap)
    FuncMap["To_Multisig_ApproveAsMultiCall"] = reflect.ValueOf(To_Multisig_ApproveAsMultiCall_wrap)
    FuncMap["To_Multisig_AsMultiCall"] = reflect.ValueOf(To_Multisig_AsMultiCall_wrap)
    FuncMap["To_Multisig_AsMultiThreshold1Call"] = reflect.ValueOf(To_Multisig_AsMultiThreshold1Call_wrap)
    FuncMap["To_Multisig_CancelAsMultiCall"] = reflect.ValueOf(To_Multisig_CancelAsMultiCall_wrap)
    FuncMap["To_NominationPools_AdjustPoolDepositCall"] = reflect.ValueOf(To_NominationPools_AdjustPoolDepositCall_wrap)
    FuncMap["To_NominationPools_BondExtraCall"] = reflect.ValueOf(To_NominationPools_BondExtraCall_wrap)
    FuncMap["To_NominationPools_BondExtraOtherCall"] = reflect.ValueOf(To_NominationPools_BondExtraOtherCall_wrap)
    FuncMap["To_NominationPools_ChillCall"] = reflect.ValueOf(To_NominationPools_ChillCall_wrap)
    FuncMap["To_NominationPools_ClaimCommissionCall"] = reflect.ValueOf(To_NominationPools_ClaimCommissionCall_wrap)
    FuncMap["To_NominationPools_ClaimPayoutOtherCall"] = reflect.ValueOf(To_NominationPools_ClaimPayoutOtherCall_wrap)
    FuncMap["To_NominationPools_CreateCall"] = reflect.ValueOf(To_NominationPools_CreateCall_wrap)
    FuncMap["To_NominationPools_CreateWithPoolIdCall"] = reflect.ValueOf(To_NominationPools_CreateWithPoolIdCall_wrap)
    FuncMap["To_NominationPools_JoinCall"] = reflect.ValueOf(To_NominationPools_JoinCall_wrap)
    FuncMap["To_NominationPools_NominateCall"] = reflect.ValueOf(To_NominationPools_NominateCall_wrap)
    FuncMap["To_NominationPools_PoolWithdrawUnbondedCall"] = reflect.ValueOf(To_NominationPools_PoolWithdrawUnbondedCall_wrap)
    FuncMap["To_NominationPools_SetClaimPermissionCall"] = reflect.ValueOf(To_NominationPools_SetClaimPermissionCall_wrap)
    FuncMap["To_NominationPools_SetCommissionCall"] = reflect.ValueOf(To_NominationPools_SetCommissionCall_wrap)
    FuncMap["To_NominationPools_SetCommissionChangeRateCall"] = reflect.ValueOf(To_NominationPools_SetCommissionChangeRateCall_wrap)
    FuncMap["To_NominationPools_SetCommissionMaxCall"] = reflect.ValueOf(To_NominationPools_SetCommissionMaxCall_wrap)
    FuncMap["To_NominationPools_SetConfigsCall"] = reflect.ValueOf(To_NominationPools_SetConfigsCall_wrap)
    FuncMap["To_NominationPools_SetMetadataCall"] = reflect.ValueOf(To_NominationPools_SetMetadataCall_wrap)
    FuncMap["To_NominationPools_SetStateCall"] = reflect.ValueOf(To_NominationPools_SetStateCall_wrap)
    FuncMap["To_NominationPools_UnbondCall"] = reflect.ValueOf(To_NominationPools_UnbondCall_wrap)
    FuncMap["To_NominationPools_UpdateRolesCall"] = reflect.ValueOf(To_NominationPools_UpdateRolesCall_wrap)
    FuncMap["To_NominationPools_WithdrawUnbondedCall"] = reflect.ValueOf(To_NominationPools_WithdrawUnbondedCall_wrap)
    FuncMap["To_Preimage_EnsureUpdatedCall"] = reflect.ValueOf(To_Preimage_EnsureUpdatedCall_wrap)
    FuncMap["To_Preimage_NotePreimageCall"] = reflect.ValueOf(To_Preimage_NotePreimageCall_wrap)
    FuncMap["To_Preimage_RequestPreimageCall"] = reflect.ValueOf(To_Preimage_RequestPreimageCall_wrap)
    FuncMap["To_Preimage_UnnotePreimageCall"] = reflect.ValueOf(To_Preimage_UnnotePreimageCall_wrap)
    FuncMap["To_Preimage_UnrequestPreimageCall"] = reflect.ValueOf(To_Preimage_UnrequestPreimageCall_wrap)
    FuncMap["To_Proxy_AddProxyCall"] = reflect.ValueOf(To_Proxy_AddProxyCall_wrap)
    FuncMap["To_Proxy_AnnounceCall"] = reflect.ValueOf(To_Proxy_AnnounceCall_wrap)
    FuncMap["To_Proxy_CreatePureCall"] = reflect.ValueOf(To_Proxy_CreatePureCall_wrap)
    FuncMap["To_Proxy_KillPureCall"] = reflect.ValueOf(To_Proxy_KillPureCall_wrap)
    FuncMap["To_Proxy_ProxyAnnouncedCall"] = reflect.ValueOf(To_Proxy_ProxyAnnouncedCall_wrap)
    FuncMap["To_Proxy_ProxyCall"] = reflect.ValueOf(To_Proxy_ProxyCall_wrap)
    FuncMap["To_Proxy_RejectAnnouncementCall"] = reflect.ValueOf(To_Proxy_RejectAnnouncementCall_wrap)
    FuncMap["To_Proxy_RemoveAnnouncementCall"] = reflect.ValueOf(To_Proxy_RemoveAnnouncementCall_wrap)
    FuncMap["To_Proxy_RemoveProxyCall"] = reflect.ValueOf(To_Proxy_RemoveProxyCall_wrap)
    FuncMap["To_Referenda_CancelCall"] = reflect.ValueOf(To_Referenda_CancelCall_wrap)
    FuncMap["To_Referenda_KillCall"] = reflect.ValueOf(To_Referenda_KillCall_wrap)
    FuncMap["To_Referenda_NudgeReferendumCall"] = reflect.ValueOf(To_Referenda_NudgeReferendumCall_wrap)
    FuncMap["To_Referenda_OneFewerDecidingCall"] = reflect.ValueOf(To_Referenda_OneFewerDecidingCall_wrap)
    FuncMap["To_Referenda_PlaceDecisionDepositCall"] = reflect.ValueOf(To_Referenda_PlaceDecisionDepositCall_wrap)
    FuncMap["To_Referenda_RefundDecisionDepositCall"] = reflect.ValueOf(To_Referenda_RefundDecisionDepositCall_wrap)
    FuncMap["To_Referenda_RefundSubmissionDepositCall"] = reflect.ValueOf(To_Referenda_RefundSubmissionDepositCall_wrap)
    FuncMap["To_Referenda_SetMetadataCall"] = reflect.ValueOf(To_Referenda_SetMetadataCall_wrap)
    FuncMap["To_Referenda_SubmitCall"] = reflect.ValueOf(To_Referenda_SubmitCall_wrap)
    FuncMap["To_Scheduler_CancelCall"] = reflect.ValueOf(To_Scheduler_CancelCall_wrap)
    FuncMap["To_Scheduler_CancelNamedCall"] = reflect.ValueOf(To_Scheduler_CancelNamedCall_wrap)
    FuncMap["To_Scheduler_ScheduleAfterCall"] = reflect.ValueOf(To_Scheduler_ScheduleAfterCall_wrap)
    FuncMap["To_Scheduler_ScheduleCall"] = reflect.ValueOf(To_Scheduler_ScheduleCall_wrap)
    FuncMap["To_Scheduler_ScheduleNamedAfterCall"] = reflect.ValueOf(To_Scheduler_ScheduleNamedAfterCall_wrap)
    FuncMap["To_Scheduler_ScheduleNamedCall"] = reflect.ValueOf(To_Scheduler_ScheduleNamedCall_wrap)
    FuncMap["To_Session_SetKeysCall"] = reflect.ValueOf(To_Session_SetKeysCall_wrap)
    FuncMap["To_StakingRewards_AlignSupplyCall"] = reflect.ValueOf(To_StakingRewards_AlignSupplyCall_wrap)
    FuncMap["To_StakingRewards_ForceRefillCall"] = reflect.ValueOf(To_StakingRewards_ForceRefillCall_wrap)
    FuncMap["To_StakingRewards_RefillCall"] = reflect.ValueOf(To_StakingRewards_RefillCall_wrap)
    FuncMap["To_StakingRewards_WithdrawCall"] = reflect.ValueOf(To_StakingRewards_WithdrawCall_wrap)
    FuncMap["To_Staking_BondCall"] = reflect.ValueOf(To_Staking_BondCall_wrap)
    FuncMap["To_Staking_BondExtraCall"] = reflect.ValueOf(To_Staking_BondExtraCall_wrap)
    FuncMap["To_Staking_CancelDeferredSlashCall"] = reflect.ValueOf(To_Staking_CancelDeferredSlashCall_wrap)
    FuncMap["To_Staking_ChillOtherCall"] = reflect.ValueOf(To_Staking_ChillOtherCall_wrap)
    FuncMap["To_Staking_ForceApplyMinCommissionCall"] = reflect.ValueOf(To_Staking_ForceApplyMinCommissionCall_wrap)
    FuncMap["To_Staking_ForceUnstakeCall"] = reflect.ValueOf(To_Staking_ForceUnstakeCall_wrap)
    FuncMap["To_Staking_IncreaseValidatorCountCall"] = reflect.ValueOf(To_Staking_IncreaseValidatorCountCall_wrap)
    FuncMap["To_Staking_KickCall"] = reflect.ValueOf(To_Staking_KickCall_wrap)
    FuncMap["To_Staking_NominateCall"] = reflect.ValueOf(To_Staking_NominateCall_wrap)
    FuncMap["To_Staking_PayoutStakersCall"] = reflect.ValueOf(To_Staking_PayoutStakersCall_wrap)
    FuncMap["To_Staking_ReapStashCall"] = reflect.ValueOf(To_Staking_ReapStashCall_wrap)
    FuncMap["To_Staking_RebondCall"] = reflect.ValueOf(To_Staking_RebondCall_wrap)
    FuncMap["To_Staking_ScaleValidatorCountCall"] = reflect.ValueOf(To_Staking_ScaleValidatorCountCall_wrap)
    FuncMap["To_Staking_SetInvulnerablesCall"] = reflect.ValueOf(To_Staking_SetInvulnerablesCall_wrap)
    FuncMap["To_Staking_SetMinCommissionCall"] = reflect.ValueOf(To_Staking_SetMinCommissionCall_wrap)
    FuncMap["To_Staking_SetPayeeCall"] = reflect.ValueOf(To_Staking_SetPayeeCall_wrap)
    FuncMap["To_Staking_SetStakingConfigsCall"] = reflect.ValueOf(To_Staking_SetStakingConfigsCall_wrap)
    FuncMap["To_Staking_SetValidatorCountCall"] = reflect.ValueOf(To_Staking_SetValidatorCountCall_wrap)
    FuncMap["To_Staking_UnbondCall"] = reflect.ValueOf(To_Staking_UnbondCall_wrap)
    FuncMap["To_Staking_ValidateCall"] = reflect.ValueOf(To_Staking_ValidateCall_wrap)
    FuncMap["To_Staking_WithdrawUnbondedCall"] = reflect.ValueOf(To_Staking_WithdrawUnbondedCall_wrap)
    FuncMap["To_System_KillPrefixCall"] = reflect.ValueOf(To_System_KillPrefixCall_wrap)
    FuncMap["To_System_KillStorageCall"] = reflect.ValueOf(To_System_KillStorageCall_wrap)
    FuncMap["To_System_RemarkCall"] = reflect.ValueOf(To_System_RemarkCall_wrap)
    FuncMap["To_System_RemarkWithEventCall"] = reflect.ValueOf(To_System_RemarkWithEventCall_wrap)
    FuncMap["To_System_SetCodeCall"] = reflect.ValueOf(To_System_SetCodeCall_wrap)
    FuncMap["To_System_SetCodeWithoutChecksCall"] = reflect.ValueOf(To_System_SetCodeWithoutChecksCall_wrap)
    FuncMap["To_System_SetHeapPagesCall"] = reflect.ValueOf(To_System_SetHeapPagesCall_wrap)
    FuncMap["To_System_SetStorageCall"] = reflect.ValueOf(To_System_SetStorageCall_wrap)
    FuncMap["To_Timestamp_SetCall"] = reflect.ValueOf(To_Timestamp_SetCall_wrap)
    FuncMap["To_Treasury_ApproveProposalCall"] = reflect.ValueOf(To_Treasury_ApproveProposalCall_wrap)
    FuncMap["To_Treasury_CheckStatusCall"] = reflect.ValueOf(To_Treasury_CheckStatusCall_wrap)
    FuncMap["To_Treasury_PayoutCall"] = reflect.ValueOf(To_Treasury_PayoutCall_wrap)
    FuncMap["To_Treasury_ProposeSpendCall"] = reflect.ValueOf(To_Treasury_ProposeSpendCall_wrap)
    FuncMap["To_Treasury_RejectProposalCall"] = reflect.ValueOf(To_Treasury_RejectProposalCall_wrap)
    FuncMap["To_Treasury_RemoveApprovalCall"] = reflect.ValueOf(To_Treasury_RemoveApprovalCall_wrap)
    FuncMap["To_Treasury_SpendCall"] = reflect.ValueOf(To_Treasury_SpendCall_wrap)
    FuncMap["To_Treasury_SpendLocalCall"] = reflect.ValueOf(To_Treasury_SpendLocalCall_wrap)
    FuncMap["To_Treasury_VoidSpendCall"] = reflect.ValueOf(To_Treasury_VoidSpendCall_wrap)
    FuncMap["To_Utility_AsDerivativeCall"] = reflect.ValueOf(To_Utility_AsDerivativeCall_wrap)
    FuncMap["To_Utility_BatchAllCall"] = reflect.ValueOf(To_Utility_BatchAllCall_wrap)
    FuncMap["To_Utility_BatchCall"] = reflect.ValueOf(To_Utility_BatchCall_wrap)
    FuncMap["To_Utility_DispatchAsCall"] = reflect.ValueOf(To_Utility_DispatchAsCall_wrap)
    FuncMap["To_Utility_ForceBatchCall"] = reflect.ValueOf(To_Utility_ForceBatchCall_wrap)
    FuncMap["To_Utility_WithWeightCall"] = reflect.ValueOf(To_Utility_WithWeightCall_wrap)
    FuncMap["To_Vesting_ForceVestedTransferCall"] = reflect.ValueOf(To_Vesting_ForceVestedTransferCall_wrap)
    FuncMap["To_Vesting_MergeSchedulesCall"] = reflect.ValueOf(To_Vesting_MergeSchedulesCall_wrap)
    FuncMap["To_Vesting_VestOtherCall"] = reflect.ValueOf(To_Vesting_VestOtherCall_wrap)
    FuncMap["To_Vesting_VestedTransferCall"] = reflect.ValueOf(To_Vesting_VestedTransferCall_wrap)
    FuncMap["To_Whitelist_DispatchWhitelistedCallCall"] = reflect.ValueOf(To_Whitelist_DispatchWhitelistedCallCall_wrap)
    FuncMap["To_Whitelist_DispatchWhitelistedCallWithPreimageCall"] = reflect.ValueOf(To_Whitelist_DispatchWhitelistedCallWithPreimageCall_wrap)
    FuncMap["To_Whitelist_RemoveWhitelistedCallCall"] = reflect.ValueOf(To_Whitelist_RemoveWhitelistedCallCall_wrap)
    FuncMap["To_Whitelist_WhitelistCallCall"] = reflect.ValueOf(To_Whitelist_WhitelistCallCall_wrap)
}
