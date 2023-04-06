// Code generated from _gen/Wasm.rules using 'go generate'; DO NOT EDIT.

package ssa

import "internal/buildcfg"
import "math"
import "cmd/compile/internal/types"

func rewriteValueWasm(v *Value) bool {
	switch v.Op {
	case OpAbs:
		v.Op = OpWasmF64Abs
		return true
	case OpAdd16:
		v.Op = OpWasmI32Add
		return true
	case OpAdd32:
		v.Op = OpWasmI32Add
		return true
	case OpAdd32F:
		v.Op = OpWasmF32Add
		return true
	case OpAdd64:
		v.Op = OpWasmI64Add
		return true
	case OpAdd64F:
		v.Op = OpWasmF64Add
		return true
	case OpAdd8:
		v.Op = OpWasmI32Add
		return true
	case OpAddPtr:
		v.Op = OpWasmI32Add
		return true
	case OpAddr:
		return rewriteValueWasm_OpAddr(v)
	case OpAnd16:
		v.Op = OpWasmI32And
		return true
	case OpAnd32:
		v.Op = OpWasmI32And
		return true
	case OpAnd64:
		v.Op = OpWasmI64And
		return true
	case OpAnd8:
		v.Op = OpWasmI32And
		return true
	case OpAndB:
		v.Op = OpWasmI32And
		return true
	case OpBitLen32:
		return rewriteValueWasm_OpBitLen32(v)
	case OpBitLen64:
		return rewriteValueWasm_OpBitLen64(v)
	case OpCeil:
		v.Op = OpWasmF64Ceil
		return true
	case OpClosureCall:
		v.Op = OpWasmLoweredClosureCall
		return true
	case OpCom16:
		return rewriteValueWasm_OpCom16(v)
	case OpCom32:
		return rewriteValueWasm_OpCom32(v)
	case OpCom64:
		return rewriteValueWasm_OpCom64(v)
	case OpCom8:
		return rewriteValueWasm_OpCom8(v)
	case OpCondSelect:
		v.Op = OpWasmSelect
		return true
	case OpConst16:
		return rewriteValueWasm_OpConst16(v)
	case OpConst32:
		return rewriteValueWasm_OpConst32(v)
	case OpConst32F:
		v.Op = OpWasmF32Const
		return true
	case OpConst64:
		v.Op = OpWasmI64Const
		return true
	case OpConst64F:
		v.Op = OpWasmF64Const
		return true
	case OpConst8:
		return rewriteValueWasm_OpConst8(v)
	case OpConstBool:
		return rewriteValueWasm_OpConstBool(v)
	case OpConstNil:
		return rewriteValueWasm_OpConstNil(v)
	case OpConvert:
		v.Op = OpWasmLoweredConvert
		return true
	case OpCopysign:
		v.Op = OpWasmF64Copysign
		return true
	case OpCtz16:
		return rewriteValueWasm_OpCtz16(v)
	case OpCtz16NonZero:
		v.Op = OpWasmI32Ctz
		return true
	case OpCtz32:
		return rewriteValueWasm_OpCtz32(v)
	case OpCtz32NonZero:
		v.Op = OpWasmI32Ctz
		return true
	case OpCtz64:
		v.Op = OpWasmI64Ctz
		return true
	case OpCtz64NonZero:
		v.Op = OpWasmI64Ctz
		return true
	case OpCtz8:
		return rewriteValueWasm_OpCtz8(v)
	case OpCtz8NonZero:
		v.Op = OpWasmI32Ctz
		return true
	case OpCvt32Fto32:
		v.Op = OpWasmI32TruncSatF32S
		return true
	case OpCvt32Fto32U:
		v.Op = OpWasmI32TruncSatF32U
		return true
	case OpCvt32Fto64:
		v.Op = OpWasmI64TruncSatF32S
		return true
	case OpCvt32Fto64F:
		v.Op = OpWasmF64PromoteF32
		return true
	case OpCvt32Fto64U:
		v.Op = OpWasmI64TruncSatF32U
		return true
	case OpCvt32Uto32F:
		return rewriteValueWasm_OpCvt32Uto32F(v)
	case OpCvt32Uto64F:
		return rewriteValueWasm_OpCvt32Uto64F(v)
	case OpCvt32to32F:
		return rewriteValueWasm_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValueWasm_OpCvt32to64F(v)
	case OpCvt64Fto32:
		v.Op = OpWasmI32TruncSatF64S
		return true
	case OpCvt64Fto32F:
		v.Op = OpWasmF32DemoteF64
		return true
	case OpCvt64Fto32U:
		v.Op = OpWasmI64TruncSatF64U
		return true
	case OpCvt64Fto64:
		v.Op = OpWasmI64TruncSatF64S
		return true
	case OpCvt64Fto64U:
		v.Op = OpWasmI64TruncSatF64U
		return true
	case OpCvt64Uto32F:
		v.Op = OpWasmF32ConvertI64U
		return true
	case OpCvt64Uto64F:
		v.Op = OpWasmF64ConvertI64U
		return true
	case OpCvt64to32F:
		v.Op = OpWasmF32ConvertI64S
		return true
	case OpCvt64to64F:
		v.Op = OpWasmF64ConvertI64S
		return true
	case OpCvtBoolToUint8:
		v.Op = OpCopy
		return true
	case OpDiv16:
		return rewriteValueWasm_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueWasm_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueWasm_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpWasmF32Div
		return true
	case OpDiv32u:
		return rewriteValueWasm_OpDiv32u(v)
	case OpDiv64:
		return rewriteValueWasm_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpWasmF64Div
		return true
	case OpDiv64u:
		v.Op = OpWasmI64DivU
		return true
	case OpDiv8:
		return rewriteValueWasm_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueWasm_OpDiv8u(v)
	case OpEq16:
		return rewriteValueWasm_OpEq16(v)
	case OpEq32:
		return rewriteValueWasm_OpEq32(v)
	case OpEq32F:
		v.Op = OpWasmF32Eq
		return true
	case OpEq64:
		v.Op = OpWasmI64Eq
		return true
	case OpEq64F:
		v.Op = OpWasmF64Eq
		return true
	case OpEq8:
		return rewriteValueWasm_OpEq8(v)
	case OpEqB:
		v.Op = OpWasmI32Eq
		return true
	case OpEqPtr:
		v.Op = OpWasmI32Eq
		return true
	case OpFloor:
		v.Op = OpWasmF64Floor
		return true
	case OpGetCallerPC:
		v.Op = OpWasmLoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpWasmLoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpWasmLoweredGetClosurePtr
		return true
	case OpInterCall:
		v.Op = OpWasmLoweredInterCall
		return true
	case OpIsInBounds:
		v.Op = OpWasmI32LtU
		return true
	case OpIsNonNil:
		return rewriteValueWasm_OpIsNonNil(v)
	case OpIsSliceInBounds:
		v.Op = OpWasmI32LeU
		return true
	case OpLeq16:
		return rewriteValueWasm_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueWasm_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueWasm_OpLeq32(v)
	case OpLeq32F:
		v.Op = OpWasmF32Le
		return true
	case OpLeq32U:
		return rewriteValueWasm_OpLeq32U(v)
	case OpLeq64:
		v.Op = OpWasmI64LeS
		return true
	case OpLeq64F:
		v.Op = OpWasmF64Le
		return true
	case OpLeq64U:
		v.Op = OpWasmI64LeU
		return true
	case OpLeq8:
		return rewriteValueWasm_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueWasm_OpLeq8U(v)
	case OpLess16:
		return rewriteValueWasm_OpLess16(v)
	case OpLess16U:
		return rewriteValueWasm_OpLess16U(v)
	case OpLess32:
		return rewriteValueWasm_OpLess32(v)
	case OpLess32F:
		v.Op = OpWasmF32Lt
		return true
	case OpLess32U:
		return rewriteValueWasm_OpLess32U(v)
	case OpLess64:
		v.Op = OpWasmI64LtS
		return true
	case OpLess64F:
		v.Op = OpWasmF64Lt
		return true
	case OpLess64U:
		v.Op = OpWasmI64LtU
		return true
	case OpLess8:
		return rewriteValueWasm_OpLess8(v)
	case OpLess8U:
		return rewriteValueWasm_OpLess8U(v)
	case OpLoad:
		return rewriteValueWasm_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueWasm_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueWasm_OpLsh16x16(v)
	case OpLsh16x32:
		v.Op = OpLsh32x32
		return true
	case OpLsh16x64:
		v.Op = OpLsh64x64
		return true
	case OpLsh16x8:
		return rewriteValueWasm_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueWasm_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueWasm_OpLsh32x32(v)
	case OpLsh32x64:
		v.Op = OpLsh64x64
		return true
	case OpLsh32x8:
		return rewriteValueWasm_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueWasm_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueWasm_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueWasm_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueWasm_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueWasm_OpLsh8x16(v)
	case OpLsh8x32:
		v.Op = OpLsh32x32
		return true
	case OpLsh8x64:
		v.Op = OpLsh64x64
		return true
	case OpLsh8x8:
		return rewriteValueWasm_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueWasm_OpMod16(v)
	case OpMod16u:
		return rewriteValueWasm_OpMod16u(v)
	case OpMod32:
		return rewriteValueWasm_OpMod32(v)
	case OpMod32u:
		return rewriteValueWasm_OpMod32u(v)
	case OpMod64:
		return rewriteValueWasm_OpMod64(v)
	case OpMod64u:
		v.Op = OpWasmI64RemU
		return true
	case OpMod8:
		return rewriteValueWasm_OpMod8(v)
	case OpMod8u:
		return rewriteValueWasm_OpMod8u(v)
	case OpMove:
		return rewriteValueWasm_OpMove(v)
	case OpMul16:
		v.Op = OpWasmI32Mul
		return true
	case OpMul32:
		v.Op = OpWasmI32Mul
		return true
	case OpMul32F:
		v.Op = OpWasmF32Mul
		return true
	case OpMul64:
		v.Op = OpWasmI64Mul
		return true
	case OpMul64F:
		v.Op = OpWasmF64Mul
		return true
	case OpMul8:
		v.Op = OpWasmI32Mul
		return true
	case OpNeg16:
		return rewriteValueWasm_OpNeg16(v)
	case OpNeg32:
		return rewriteValueWasm_OpNeg32(v)
	case OpNeg32F:
		v.Op = OpWasmF32Neg
		return true
	case OpNeg64:
		return rewriteValueWasm_OpNeg64(v)
	case OpNeg64F:
		v.Op = OpWasmF64Neg
		return true
	case OpNeg8:
		return rewriteValueWasm_OpNeg8(v)
	case OpNeq16:
		return rewriteValueWasm_OpNeq16(v)
	case OpNeq32:
		return rewriteValueWasm_OpNeq32(v)
	case OpNeq32F:
		v.Op = OpWasmF32Ne
		return true
	case OpNeq64:
		v.Op = OpWasmI64Ne
		return true
	case OpNeq64F:
		v.Op = OpWasmF64Ne
		return true
	case OpNeq8:
		return rewriteValueWasm_OpNeq8(v)
	case OpNeqB:
		v.Op = OpWasmI32Ne
		return true
	case OpNeqPtr:
		v.Op = OpWasmI32Ne
		return true
	case OpNilCheck:
		v.Op = OpWasmLoweredNilCheck
		return true
	case OpNot:
		v.Op = OpWasmI32Eqz
		return true
	case OpOffPtr:
		return rewriteValueWasm_OpOffPtr(v)
	case OpOr16:
		v.Op = OpWasmI32Or
		return true
	case OpOr32:
		v.Op = OpWasmI32Or
		return true
	case OpOr64:
		v.Op = OpWasmI64Or
		return true
	case OpOr8:
		v.Op = OpWasmI32Or
		return true
	case OpOrB:
		v.Op = OpWasmI32Or
		return true
	case OpPopCount16:
		return rewriteValueWasm_OpPopCount16(v)
	case OpPopCount32:
		return rewriteValueWasm_OpPopCount32(v)
	case OpPopCount64:
		v.Op = OpWasmI64Popcnt
		return true
	case OpPopCount8:
		return rewriteValueWasm_OpPopCount8(v)
	case OpRotateLeft16:
		return rewriteValueWasm_OpRotateLeft16(v)
	case OpRotateLeft32:
		v.Op = OpWasmI32Rotl
		return true
	case OpRotateLeft64:
		v.Op = OpWasmI64Rotl
		return true
	case OpRotateLeft8:
		return rewriteValueWasm_OpRotateLeft8(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRoundToEven:
		v.Op = OpWasmF64Nearest
		return true
	case OpRsh16Ux16:
		return rewriteValueWasm_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueWasm_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueWasm_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueWasm_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueWasm_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueWasm_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueWasm_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueWasm_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueWasm_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueWasm_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueWasm_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueWasm_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueWasm_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueWasm_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueWasm_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueWasm_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueWasm_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueWasm_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueWasm_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueWasm_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueWasm_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueWasm_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueWasm_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueWasm_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueWasm_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueWasm_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueWasm_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueWasm_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueWasm_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueWasm_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueWasm_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueWasm_OpRsh8x8(v)
	case OpSignExt16to32:
		return rewriteValueWasm_OpSignExt16to32(v)
	case OpSignExt16to64:
		return rewriteValueWasm_OpSignExt16to64(v)
	case OpSignExt32to64:
		return rewriteValueWasm_OpSignExt32to64(v)
	case OpSignExt8to16:
		return rewriteValueWasm_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValueWasm_OpSignExt8to32(v)
	case OpSignExt8to64:
		return rewriteValueWasm_OpSignExt8to64(v)
	case OpSlicemask:
		return rewriteValueWasm_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpWasmF64Sqrt
		return true
	case OpSqrt32:
		v.Op = OpWasmF32Sqrt
		return true
	case OpStaticCall:
		v.Op = OpWasmLoweredStaticCall
		return true
	case OpStore:
		return rewriteValueWasm_OpStore(v)
	case OpSub16:
		v.Op = OpWasmI32Sub
		return true
	case OpSub32:
		v.Op = OpWasmI32Sub
		return true
	case OpSub32F:
		v.Op = OpWasmF32Sub
		return true
	case OpSub64:
		v.Op = OpWasmI64Sub
		return true
	case OpSub64F:
		v.Op = OpWasmF64Sub
		return true
	case OpSub8:
		v.Op = OpWasmI32Sub
		return true
	case OpSubPtr:
		v.Op = OpWasmI32Sub
		return true
	case OpTailCall:
		v.Op = OpWasmLoweredTailCall
		return true
	case OpTrunc:
		v.Op = OpWasmF64Trunc
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpTrunc64to16:
		v.Op = OpCopy
		return true
	case OpTrunc64to32:
		v.Op = OpCopy
		return true
	case OpTrunc64to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpWasmLoweredWB
		return true
	case OpWasmF32Add:
		return rewriteValueWasm_OpWasmF32Add(v)
	case OpWasmF32Mul:
		return rewriteValueWasm_OpWasmF32Mul(v)
	case OpWasmF64Add:
		return rewriteValueWasm_OpWasmF64Add(v)
	case OpWasmF64Mul:
		return rewriteValueWasm_OpWasmF64Mul(v)
	case OpWasmI32Add:
		return rewriteValueWasm_OpWasmI32Add(v)
	case OpWasmI32AddConst:
		return rewriteValueWasm_OpWasmI32AddConst(v)
	case OpWasmI32And:
		return rewriteValueWasm_OpWasmI32And(v)
	case OpWasmI32Eq:
		return rewriteValueWasm_OpWasmI32Eq(v)
	case OpWasmI32Eqz:
		return rewriteValueWasm_OpWasmI32Eqz(v)
	case OpWasmI32LeU:
		return rewriteValueWasm_OpWasmI32LeU(v)
	case OpWasmI32Load:
		return rewriteValueWasm_OpWasmI32Load(v)
	case OpWasmI32Load16S:
		return rewriteValueWasm_OpWasmI32Load16S(v)
	case OpWasmI32Load16U:
		return rewriteValueWasm_OpWasmI32Load16U(v)
	case OpWasmI32Load8S:
		return rewriteValueWasm_OpWasmI32Load8S(v)
	case OpWasmI32Load8U:
		return rewriteValueWasm_OpWasmI32Load8U(v)
	case OpWasmI32LtU:
		return rewriteValueWasm_OpWasmI32LtU(v)
	case OpWasmI32Mul:
		return rewriteValueWasm_OpWasmI32Mul(v)
	case OpWasmI32Ne:
		return rewriteValueWasm_OpWasmI32Ne(v)
	case OpWasmI32Or:
		return rewriteValueWasm_OpWasmI32Or(v)
	case OpWasmI32Shl:
		return rewriteValueWasm_OpWasmI32Shl(v)
	case OpWasmI32ShrS:
		return rewriteValueWasm_OpWasmI32ShrS(v)
	case OpWasmI32ShrU:
		return rewriteValueWasm_OpWasmI32ShrU(v)
	case OpWasmI32Store:
		return rewriteValueWasm_OpWasmI32Store(v)
	case OpWasmI32Store16:
		return rewriteValueWasm_OpWasmI32Store16(v)
	case OpWasmI32Store8:
		return rewriteValueWasm_OpWasmI32Store8(v)
	case OpWasmI32Xor:
		return rewriteValueWasm_OpWasmI32Xor(v)
	case OpWasmI64Add:
		return rewriteValueWasm_OpWasmI64Add(v)
	case OpWasmI64And:
		return rewriteValueWasm_OpWasmI64And(v)
	case OpWasmI64Eq:
		return rewriteValueWasm_OpWasmI64Eq(v)
	case OpWasmI64LeU:
		return rewriteValueWasm_OpWasmI64LeU(v)
	case OpWasmI64Load:
		return rewriteValueWasm_OpWasmI64Load(v)
	case OpWasmI64Load16S:
		return rewriteValueWasm_OpWasmI64Load16S(v)
	case OpWasmI64Load16U:
		return rewriteValueWasm_OpWasmI64Load16U(v)
	case OpWasmI64Load32S:
		return rewriteValueWasm_OpWasmI64Load32S(v)
	case OpWasmI64Load32U:
		return rewriteValueWasm_OpWasmI64Load32U(v)
	case OpWasmI64Load8S:
		return rewriteValueWasm_OpWasmI64Load8S(v)
	case OpWasmI64Load8U:
		return rewriteValueWasm_OpWasmI64Load8U(v)
	case OpWasmI64LtU:
		return rewriteValueWasm_OpWasmI64LtU(v)
	case OpWasmI64Mul:
		return rewriteValueWasm_OpWasmI64Mul(v)
	case OpWasmI64Ne:
		return rewriteValueWasm_OpWasmI64Ne(v)
	case OpWasmI64Or:
		return rewriteValueWasm_OpWasmI64Or(v)
	case OpWasmI64Shl:
		return rewriteValueWasm_OpWasmI64Shl(v)
	case OpWasmI64ShrS:
		return rewriteValueWasm_OpWasmI64ShrS(v)
	case OpWasmI64ShrU:
		return rewriteValueWasm_OpWasmI64ShrU(v)
	case OpWasmI64Store:
		return rewriteValueWasm_OpWasmI64Store(v)
	case OpWasmI64Store16:
		return rewriteValueWasm_OpWasmI64Store16(v)
	case OpWasmI64Store32:
		return rewriteValueWasm_OpWasmI64Store32(v)
	case OpWasmI64Store8:
		return rewriteValueWasm_OpWasmI64Store8(v)
	case OpWasmI64Xor:
		return rewriteValueWasm_OpWasmI64Xor(v)
	case OpXor16:
		v.Op = OpWasmI32Xor
		return true
	case OpXor32:
		v.Op = OpWasmI32Xor
		return true
	case OpXor64:
		v.Op = OpWasmI64Xor
		return true
	case OpXor8:
		v.Op = OpWasmI32Xor
		return true
	case OpZero:
		return rewriteValueWasm_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValueWasm_OpZeroExt16to32(v)
	case OpZeroExt16to64:
		return rewriteValueWasm_OpZeroExt16to64(v)
	case OpZeroExt32to64:
		return rewriteValueWasm_OpZeroExt32to64(v)
	case OpZeroExt8to16:
		return rewriteValueWasm_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValueWasm_OpZeroExt8to32(v)
	case OpZeroExt8to64:
		return rewriteValueWasm_OpZeroExt8to64(v)
	}
	return false
}
func rewriteValueWasm_OpAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (LoweredAddr {sym} [0] base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpWasmLoweredAddr)
		v.AuxInt = int32ToAuxInt(0)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueWasm_OpBitLen32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen32 x)
	// result: (I32Sub (I32Const [32]) (I32Clz x))
	for {
		x := v_0
		v.reset(OpWasmI32Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Clz, typ.Int32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpBitLen64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen64 x)
	// result: (I64Sub (I64Const [64]) (I64Clz x))
	for {
		x := v_0
		v.reset(OpWasmI64Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(64)
		v1 := b.NewValue0(v.Pos, OpWasmI64Clz, typ.Int64)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com16 x)
	// result: (I32Xor x (I32Const [-1]))
	for {
		x := v_0
		v.reset(OpWasmI32Xor)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com32 x)
	// result: (I32Xor x (I32Const [-1]))
	for {
		x := v_0
		v.reset(OpWasmI32Xor)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpCom64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com64 x)
	// result: (I64Xor x (I64Const [-1]))
	for {
		x := v_0
		v.reset(OpWasmI64Xor)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com8 x)
	// result: (I32Xor x (I32Const [-1]))
	for {
		x := v_0
		v.reset(OpWasmI32Xor)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpConst16(v *Value) bool {
	// match: (Const16 [c])
	// result: (I32Const [int32(c)])
	for {
		c := auxIntToInt16(v.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(int32(c))
		return true
	}
}
func rewriteValueWasm_OpConst32(v *Value) bool {
	// match: (Const32 [c])
	// result: (I32Const [int32(c)])
	for {
		c := auxIntToInt32(v.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(int32(c))
		return true
	}
}
func rewriteValueWasm_OpConst8(v *Value) bool {
	// match: (Const8 [c])
	// result: (I32Const [int32(c)])
	for {
		c := auxIntToInt8(v.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(int32(c))
		return true
	}
}
func rewriteValueWasm_OpConstBool(v *Value) bool {
	// match: (ConstBool [c])
	// result: (I32Const [b2i32(c)])
	for {
		c := auxIntToBool(v.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(b2i32(c))
		return true
	}
}
func rewriteValueWasm_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (I32Const [0])
	for {
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
}
func rewriteValueWasm_OpCtz16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 x)
	// result: (I32Ctz (I32Or x (I32Const [0x10000])))
	for {
		x := v_0
		v.reset(OpWasmI32Ctz)
		v0 := b.NewValue0(v.Pos, OpWasmI32Or, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(0x10000)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Ctz32 x)
	// result: (I32Ctz x)
	for {
		x := v_0
		v.reset(OpWasmI32Ctz)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCtz8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 x)
	// result: (I32Ctz (I32Or x (I32Const [0x100])))
	for {
		x := v_0
		v.reset(OpWasmI32Ctz)
		v0 := b.NewValue0(v.Pos, OpWasmI32Or, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(0x100)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt32Uto32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Uto32F x)
	// result: (F32ConvertI32U x)
	for {
		x := v_0
		v.reset(OpWasmF32ConvertI32U)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt32Uto64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Uto64F x)
	// result: (F64ConvertI64U (ZeroExt32to64 x))
	for {
		x := v_0
		v.reset(OpWasmF64ConvertI64U)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt32to32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32to32F x)
	// result: (F32ConvertI32S x)
	for {
		x := v_0
		v.reset(OpWasmF32ConvertI32S)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt32to64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to64F x)
	// result: (F64ConvertI64S (SignExt32to64 x))
	for {
		x := v_0
		v.reset(OpWasmF64ConvertI64S)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 [false] x y)
	// result: (I32DivS (SignExt16to32 x) (SignExt16to32 y))
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasmI32DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
	return false
}
func rewriteValueWasm_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (I32DivU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div32 [false] x y)
	// result: (I32DivS x y)
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasmI32DivS)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueWasm_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div32u x y)
	// result: (I32DivU x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32DivU)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64 [false] x y)
	// result: (I64DivS x y)
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasmI64DivS)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueWasm_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (I32DivS (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (I32DivU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (I32Eq (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Eq32 x y)
	// result: (I32Eq x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32Eq)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (I32Eq (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil p)
	// result: (I32Eqz (I32Eqz p))
	for {
		p := v_0
		v.reset(OpWasmI32Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI32Eqz, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (I32LeS (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (I32LeU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq32 x y)
	// result: (I32LeS x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LeS)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq32U x y)
	// result: (I32LeU x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LeU)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (I32LeS (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (I32LeU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (I32LtS (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (I32LtU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32 x y)
	// result: (I32LtS x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LtS)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32U x y)
	// result: (I64LtU x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI64LtU)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (I32LtS (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (I32LtU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (F32Load ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpWasmF32Load)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (F64Load ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpWasmF64Load)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 8
	// result: (I64Load ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 8) {
			break
		}
		v.reset(OpWasmI64Load)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 4
	// result: (I32Load ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 4) {
			break
		}
		v.reset(OpWasmI32Load)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 2 && !t.IsSigned()
	// result: (I32Load16U ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 2 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasmI32Load16U)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 2 && t.IsSigned()
	// result: (I32Load16S ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 2 && t.IsSigned()) {
			break
		}
		v.reset(OpWasmI32Load16S)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 1 && !t.IsSigned()
	// result: (I32Load8U ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 1 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasmI32Load8U)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 1 && t.IsSigned()
	// result: (I32Load8S ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 1 && t.IsSigned()) {
			break
		}
		v.reset(OpWasmI32Load8S)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpLocalAddr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (LocalAddr <t> {sym} base mem)
	// cond: t.Elem().HasPointers()
	// result: (LoweredAddr {sym} (SPanchored base mem))
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		mem := v_1
		if !(t.Elem().HasPointers()) {
			break
		}
		v.reset(OpWasmLoweredAddr)
		v.Aux = symToAux(sym)
		v0 := b.NewValue0(v.Pos, OpSPanchored, typ.Uintptr)
		v0.AddArg2(base, mem)
		v.AddArg(v0)
		return true
	}
	// match: (LocalAddr <t> {sym} base _)
	// cond: !t.Elem().HasPointers()
	// result: (LoweredAddr {sym} base)
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		if !(!t.Elem().HasPointers()) {
			break
		}
		v.reset(OpWasmLoweredAddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
	return false
}
func rewriteValueWasm_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 [c] x y)
	// result: (Lsh32x32 [c] x (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 [c] x y)
	// result: (Lsh32x32 [c] x (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 [c] x y)
	// result: (Lsh32x32 [c] x (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh32x32 [c] x y)
	// result: (Lsh32x32 [c] x y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 [c] x y)
	// result: (Lsh32x32 [c] x (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (I64Shl x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasmI64Shl)
		v.AddArg2(x, y)
		return true
	}
	// match: (Lsh64x64 x (I64Const [c]))
	// cond: uint64(c) < 64
	// result: (I64Shl x (I64Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpWasmI64Shl)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh64x64 x (I64Const [c]))
	// cond: uint64(c) >= 64
	// result: (I64Const [0])
	for {
		if v_1.Op != OpWasmI64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Lsh64x64 x y)
	// result: (Select (I64Shl x y) (I64Const [0]) (I64LtU y (I64Const [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmSelect)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpWasmI64LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(64)
		v2.AddArg2(y, v3)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueWasm_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 [c] x y)
	// result: (Lsh32x32 [c] x (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 [c] x y)
	// result: (Lsh32x32 [c] x (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 [false] x y)
	// result: (I32RemS (SignExt16to32 x) (SignExt16to32 y))
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasmI32RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
	return false
}
func rewriteValueWasm_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (I32RemU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mod32 [false] x y)
	// result: (I32RemS x y)
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasmI32RemS)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueWasm_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mod32u x y)
	// result: (I32RemU x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32RemU)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mod64 [false] x y)
	// result: (I64RemS x y)
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasmI64RemS)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueWasm_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (I32RemS (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (I32RemU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (I32Store8 dst (I32Load8U src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store8)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load8U, typ.UInt8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (I32Store16 dst (I32Load16U src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store16)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load16U, typ.UInt16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (I32Store dst (I32Load src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (I32Store8 [2] dst (I32Load8U [2] src mem) (I32Store16 dst (I32Load16U src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store8)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load8U, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store16, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasmI32Load16U, typ.UInt16)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [5] dst src mem)
	// result: (I32Store8 [4] dst (I32Load8U [4] src mem) (I32Store dst (I32Load src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 5 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store8)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load8U, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [6] dst src mem)
	// result: (I32Store16 [4] dst (I32Load16U [4] src mem) (I32Store dst (I32Load src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store16)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load16U, typ.UInt16)
		v0.AuxInt = int32ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [7] dst src mem)
	// result: (I32Store [3] dst (I32Load [3] src mem) (I32Store dst (I32Load src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 7 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(3)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [8] dst src mem)
	// result: (I32Store [4] dst (I32Load [4] src mem) (I32Store dst (I32Load src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 4 && s < 8
	// result: (I32Store [int32(s-4)] dst (I32Load [int32(s-4)] src mem) (I32Store dst (I32Load src mem) mem))
	for {
		s := auxIntToInt64(v.AuxInt)
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 4 && s < 8) {
			break
		}
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(int32(s - 4))
		v0 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(int32(s - 4))
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasmI32Load, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: logLargeCopy(v, s)
	// result: (LoweredMove [s] dst src mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		dst := v_0
		src := v_1
		mem := v_2
		if !(logLargeCopy(v, s)) {
			break
		}
		v.reset(OpWasmLoweredMove)
		v.AuxInt = int64ToAuxInt(s)
		v.AddArg3(dst, src, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpNeg16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg16 x)
	// result: (I32Sub (I32Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasmI32Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm_OpNeg32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg32 x)
	// result: (I32Sub (I32Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasmI32Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm_OpNeg64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg64 x)
	// result: (I64Sub (I64Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasmI64Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm_OpNeg8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg8 x)
	// result: (I32Sub (I32Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasmI32Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (I32Ne (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Neq32 x y)
	// result: (I32Ne x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32Ne)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueWasm_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (I32Ne (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr)
	// result: (I32AddConst [int32(off)] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		v.reset(OpWasmI32AddConst)
		v.AuxInt = int32ToAuxInt(int32(off))
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueWasm_OpPopCount16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount16 x)
	// result: (I32Popcnt (ZeroExt16to32 x))
	for {
		x := v_0
		v.reset(OpWasmI32Popcnt)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpPopCount32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (PopCount32 x)
	// result: (I32Popcnt x)
	for {
		x := v_0
		v.reset(OpWasmI32Popcnt)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpPopCount8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount8 x)
	// result: (I32Popcnt (ZeroExt8to32 x))
	for {
		x := v_0
		v.reset(OpWasmI32Popcnt)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (I32Const [c]))
	// result: (Or16 (Lsh16x32 <t> x (I32Const [c&15])) (Rsh16Ux32 <t> x (I32Const [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpWasmI32Const {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x32, t)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(c & 15)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux32, t)
		v3 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v3.AuxInt = int32ToAuxInt(-c & 15)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueWasm_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (I32Const [c]))
	// result: (Or8 (Lsh8x32 <t> x (I32Const [c&7])) (Rsh8Ux32 <t> x (I32Const [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpWasmI32Const {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x32, t)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(c & 7)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux32, t)
		v3 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v3.AuxInt = int32ToAuxInt(-c & 7)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueWasm_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 [c] x y)
	// result: (Rsh32Ux32 [c] (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 [c] x y)
	// result: (Rsh32Ux32 [c] (ZeroExt16to32 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt16to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 [c] x y)
	// result: (Rsh32Ux32 [c] (ZeroExt16to32 x) (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 [c] x y)
	// result: (Rsh32x32 [c] (SignExt16to32 x) (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 [c] x y)
	// result: (Rsh32x32 [c] (SignExt16to32 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 [c] x y)
	// result: (Rsh64x64 [c] (SignExt16to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 [c] x y)
	// result: (Rsh32x32 [c] (SignExt16to32 x) (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 [c] x y)
	// result: (Rsh32Ux32 [c] x (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 x y)
	// cond: shiftIsBounded(v)
	// result: (I32ShrU x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasmI32ShrU)
		v.AddArg2(x, y)
		return true
	}
	// match: (Rsh32Ux32 x (I32Const [c]))
	// cond: uint32(c) < 32
	// result: (I32ShrU x (I32Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpWasmI32ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32Ux32 x (I32Const [c]))
	// cond: uint32(c) >= 32
	// result: (I32Const [0])
	for {
		if v_1.Op != OpWasmI32Const {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (Rsh32Ux32 x y)
	// result: (Select (I32ShrU x y) (I32Const [0]) (I32LtU y (I32Const [32])))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmSelect)
		v0 := b.NewValue0(v.Pos, OpWasmI32ShrU, typ.Int32)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpWasmI32LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v3.AuxInt = int32ToAuxInt(32)
		v2.AddArg2(y, v3)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt32to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 [c] x y)
	// result: (Rsh32Ux32 [c] x (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 [c] x y)
	// result: (Rsh32x32 [c] x (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 x y)
	// cond: shiftIsBounded(v)
	// result: (I32ShrS x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasmI32ShrS)
		v.AddArg2(x, y)
		return true
	}
	// match: (Rsh32x32 x (I32Const [c]))
	// cond: uint32(c) < 32
	// result: (I32ShrS x (I32Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpWasmI32ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32x32 x (I32Const [c]))
	// cond: uint32(c) >= 32
	// result: (I32ShrS x (I32Const [31]))
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpWasmI32ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(31)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32x32 x y)
	// result: (I32ShrS x (Select <typ.Int32> y (I32Const [31]) (I32LtU y (I32Const [32]))))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI32ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmSelect, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(31)
		v2 := b.NewValue0(v.Pos, OpWasmI32LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v3.AuxInt = int32ToAuxInt(32)
		v2.AddArg2(y, v3)
		v0.AddArg3(y, v1, v2)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 [c] x y)
	// result: (Rsh64x64 [c] (SignExt32to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 [c] x y)
	// result: (Rsh32x32 [c] x (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 [c] x y)
	// result: (Rsh64Ux64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 [c] x y)
	// result: (Rsh64Ux64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (I64ShrU x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasmI64ShrU)
		v.AddArg2(x, y)
		return true
	}
	// match: (Rsh64Ux64 x (I64Const [c]))
	// cond: uint64(c) < 64
	// result: (I64ShrU x (I64Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpWasmI64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux64 x (I64Const [c]))
	// cond: uint64(c) >= 64
	// result: (I64Const [0])
	for {
		if v_1.Op != OpWasmI64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Rsh64Ux64 x y)
	// result: (Select (I64ShrU x y) (I64Const [0]) (I64LtU y (I64Const [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmSelect)
		v0 := b.NewValue0(v.Pos, OpWasmI64ShrU, typ.Int64)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpWasmI64LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(64)
		v2.AddArg2(y, v3)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 [c] x y)
	// result: (Rsh64Ux64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 [c] x y)
	// result: (Rsh64x64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 [c] x y)
	// result: (Rsh64x64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (I64ShrS x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasmI64ShrS)
		v.AddArg2(x, y)
		return true
	}
	// match: (Rsh64x64 x (I64Const [c]))
	// cond: uint64(c) < 64
	// result: (I64ShrS x (I64Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x64 x (I64Const [c]))
	// cond: uint64(c) >= 64
	// result: (I64ShrS x (I64Const [63]))
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(63)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x64 x y)
	// result: (I64ShrS x (Select <typ.Int64> y (I64Const [63]) (I64LtU y (I64Const [64]))))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmSelect, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(63)
		v2 := b.NewValue0(v.Pos, OpWasmI64LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(64)
		v2.AddArg2(y, v3)
		v0.AddArg3(y, v1, v2)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 [c] x y)
	// result: (Rsh64x64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 [c] x y)
	// result: (Rsh32Ux32 [c] (ZeroExt8to32 x) (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 [c] x y)
	// result: (Rsh32Ux32 [c] (ZeroExt8to32 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt8to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 [c] x y)
	// result: (Rsh32Ux32 [c] (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 [c] x y)
	// result: (Rsh32x32 [c] (SignExt8to32 x) (ZeroExt16to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 [c] x y)
	// result: (Rsh32x32 [c] (SignExt8to32 x) (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 [c] x y)
	// result: (Rsh64x64 [c] (SignExt8to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 [c] x y)
	// result: (Rsh32x32 [c] (SignExt8to32 x) (ZeroExt8to32 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh32x32)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpSignExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt16to32 x:(I32Load16S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI32Load16S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt16to32 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I32Extend16S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasmI32Extend16S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt16to32 x)
	// result: (I32ShrS (I32Shl x (I32Const [16])) (I32Const [16]))
	for {
		x := v_0
		v.reset(OpWasmI32ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI32Shl, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(16)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpSignExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt16to64 x:(I64Load16S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI64Load16S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt16to64 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend16S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasmI64Extend16S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt16to64 x)
	// result: (I64ShrS (I64Shl x (I64Const [48])) (I64Const [48]))
	for {
		x := v_0
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(48)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpSignExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt32to64 x:(I64Load32S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI64Load32S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt32to64 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend32S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasmI64Extend32S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt32to64 x)
	// result: (I64ShrS (I64Shl x (I64Const [32])) (I64Const [32]))
	for {
		x := v_0
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(32)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpSignExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt8to16 x:(I32Load8S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI32Load8S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt8to16 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I32Extend8S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasmI32Extend8S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt8to16 x)
	// result: (I32ShrS (I32Shl x (I32Const [24])) (I32Const [24]))
	for {
		x := v_0
		v.reset(OpWasmI32ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI32Shl, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(24)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpSignExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt8to32 x:(I32Load8S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI32Load8S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt8to32 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I32Extend8S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasmI32Extend8S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt8to32 x)
	// result: (I32ShrS (I32Shl x (I32Const [24])) (I32Const [24]))
	for {
		x := v_0
		v.reset(OpWasmI32ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI32Shl, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(24)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpSignExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt8to64 x:(I64Load8S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI64Load8S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt8to64 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend8S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasmI64Extend8S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt8to64 x)
	// result: (I64ShrS (I64Shl x (I64Const [56])) (I64Const [56]))
	for {
		x := v_0
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(56)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Slicemask x)
	// result: (I32ShrS (I32Sub (I32Const [0]) x) (I32Const [63]))
	for {
		x := v_0
		v.reset(OpWasmI32ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI32Sub, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v1.AuxInt = int32ToAuxInt(0)
		v0.AddArg2(v1, x)
		v2 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v2.AuxInt = int32ToAuxInt(63)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueWasm_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: is64BitFloat(t)
	// result: (F64Store ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpWasmF64Store)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: is32BitFloat(t)
	// result: (F32Store ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpWasmF32Store)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8
	// result: (I64Store ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8) {
			break
		}
		v.reset(OpWasmI64Store)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4
	// result: (I32Store ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4) {
			break
		}
		v.reset(OpWasmI32Store)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 2
	// result: (I32Store16 ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 2) {
			break
		}
		v.reset(OpWasmI32Store16)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 1
	// result: (I32Store8 ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 1) {
			break
		}
		v.reset(OpWasmI32Store8)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmF32Add(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (F32Add (F32Const [x]) y)
	// cond: y.Op != OpWasmF32Const
	// result: (F32Add y (F32Const [x]))
	for {
		if v_0.Op != OpWasmF32Const {
			break
		}
		x := auxIntToFloat32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmF32Const) {
			break
		}
		v.reset(OpWasmF32Add)
		v0 := b.NewValue0(v.Pos, OpWasmF32Const, typ.Float32)
		v0.AuxInt = float32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmF32Mul(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (F32Mul (F32Const [x]) y)
	// cond: y.Op != OpWasmF32Const
	// result: (F32Mul y (F32Const [x]))
	for {
		if v_0.Op != OpWasmF32Const {
			break
		}
		x := auxIntToFloat32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmF32Const) {
			break
		}
		v.reset(OpWasmF32Mul)
		v0 := b.NewValue0(v.Pos, OpWasmF32Const, typ.Float32)
		v0.AuxInt = float32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmF64Add(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (F64Add (F64Const [x]) (F64Const [y]))
	// result: (F64Const [x + y])
	for {
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		if v_1.Op != OpWasmF64Const {
			break
		}
		y := auxIntToFloat64(v_1.AuxInt)
		v.reset(OpWasmF64Const)
		v.AuxInt = float64ToAuxInt(x + y)
		return true
	}
	// match: (F64Add (F64Const [x]) y)
	// cond: y.Op != OpWasmF64Const
	// result: (F64Add y (F64Const [x]))
	for {
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmF64Const) {
			break
		}
		v.reset(OpWasmF64Add)
		v0 := b.NewValue0(v.Pos, OpWasmF64Const, typ.Float64)
		v0.AuxInt = float64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmF64Mul(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (F64Mul (F64Const [x]) (F64Const [y]))
	// cond: !math.IsNaN(x * y)
	// result: (F64Const [x * y])
	for {
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		if v_1.Op != OpWasmF64Const {
			break
		}
		y := auxIntToFloat64(v_1.AuxInt)
		if !(!math.IsNaN(x * y)) {
			break
		}
		v.reset(OpWasmF64Const)
		v.AuxInt = float64ToAuxInt(x * y)
		return true
	}
	// match: (F64Mul (F64Const [x]) y)
	// cond: y.Op != OpWasmF64Const
	// result: (F64Mul y (F64Const [x]))
	for {
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmF64Const) {
			break
		}
		v.reset(OpWasmF64Mul)
		v0 := b.NewValue0(v.Pos, OpWasmF64Const, typ.Float64)
		v0.AuxInt = float64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Add(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32Add (I32Const [x]) (I32Const [y]))
	// result: (I32Const [x + y])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(x + y)
		return true
	}
	// match: (I32Add (I32Const [x]) y)
	// cond: y.Op != OpWasmI32Const
	// result: (I32Add y (I32Const [x]))
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI32Const) {
			break
		}
		v.reset(OpWasmI32Add)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I32Add x (I32Const <t> [y]))
	// cond: !t.IsPtr()
	// result: (I32AddConst [y] x)
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const {
			break
		}
		t := v_1.Type
		y := auxIntToInt32(v_1.AuxInt)
		if !(!t.IsPtr()) {
			break
		}
		v.reset(OpWasmI32AddConst)
		v.AuxInt = int32ToAuxInt(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32AddConst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (I32AddConst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (I32AddConst [off] (LoweredAddr {sym} [off2] base))
	// cond: isU32Bit(int64(off+int32(off2)))
	// result: (LoweredAddr {sym} [int32(off)+off2] base)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmLoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		base := v_0.Args[0]
		if !(isU32Bit(int64(off + int32(off2)))) {
			break
		}
		v.reset(OpWasmLoweredAddr)
		v.AuxInt = int32ToAuxInt(int32(off) + off2)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
	// match: (I32AddConst [off] x:(SP))
	// cond: isU32Bit(int64(off))
	// result: (LoweredAddr [int32(off)] x)
	for {
		off := auxIntToInt32(v.AuxInt)
		x := v_0
		if x.Op != OpSP || !(isU32Bit(int64(off))) {
			break
		}
		v.reset(OpWasmLoweredAddr)
		v.AuxInt = int32ToAuxInt(int32(off))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32And(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32And (I32Const [x]) (I32Const [y]))
	// result: (I32Const [x & y])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(x & y)
		return true
	}
	// match: (I32And (I32Const [x]) y)
	// cond: y.Op != OpWasmI32Const
	// result: (I32And y (I32Const [x]))
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI32Const) {
			break
		}
		v.reset(OpWasmI32And)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Eq(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32Eq (I32Const [x]) (I32Const [y]))
	// cond: x == y
	// result: (I32Const [1])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		if !(x == y) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (I32Eq (I32Const [x]) (I32Const [y]))
	// cond: x != y
	// result: (I32Const [0])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		if !(x != y) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (I32Eq (I32Const [x]) y)
	// cond: y.Op != OpWasmI32Const
	// result: (I32Eq y (I32Const [x]))
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI32Const) {
			break
		}
		v.reset(OpWasmI32Eq)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I32Eq x (I32Const [0]))
	// result: (I32Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasmI32Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Eqz(v *Value) bool {
	v_0 := v.Args[0]
	// match: (I32Eqz (I32Eqz (I32Eqz x)))
	// result: (I32Eqz x)
	for {
		if v_0.Op != OpWasmI32Eqz {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpWasmI32Eqz {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpWasmI32Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32LeU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32LeU x (I32Const [0]))
	// result: (I32Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasmI32Eqz)
		v.AddArg(x)
		return true
	}
	// match: (I32LeU (I32Const [1]) x)
	// result: (I32Eqz (I32Eqz x))
	for {
		if v_0.Op != OpWasmI32Const || auxIntToInt32(v_0.AuxInt) != 1 {
			break
		}
		x := v_1
		v.reset(OpWasmI32Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI32Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Load(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (I32Load [off] (I32AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Load [off+off2] ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Load)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (I32Load [off] (LoweredAddr {sym} [off2] (SB)) _)
	// cond: symIsRO(sym) && isU32Bit(int64(off+int32(off2)))
	// result: (I32Const [int32(read32(sym, off+int32(off2), config.ctxt.Arch.ByteOrder))])
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmLoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSB || !(symIsRO(sym) && isU32Bit(int64(off+int32(off2)))) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(int32(read32(sym, off+int32(off2), config.ctxt.Arch.ByteOrder)))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Load16S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32Load16S [off] (I32AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Load16S [off+off2] ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Load16S)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Load16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (I32Load16U [off] (I32AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Load16U [off+off2] ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Load16U)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (I32Load16U [off] (LoweredAddr {sym} [off2] (SB)) _)
	// cond: symIsRO(sym) && isU32Bit(int64(off+int32(off2)))
	// result: (I32Const [int32(read16(sym, off+int32(off2), config.ctxt.Arch.ByteOrder))])
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmLoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSB || !(symIsRO(sym) && isU32Bit(int64(off+int32(off2)))) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(int32(read16(sym, off+int32(off2), config.ctxt.Arch.ByteOrder)))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Load8S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32Load8S [off] (I32AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Load8S [off+off2] ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Load8S)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Load8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32Load8U [off] (I32AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Load8U [off+off2] ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Load8U)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (I32Load8U [off] (LoweredAddr {sym} [off2] (SB)) _)
	// cond: symIsRO(sym) && isU32Bit(int64(off+int32(off2)))
	// result: (I32Const [int32(read8(sym, off+int32(off2)))])
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmLoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSB || !(symIsRO(sym) && isU32Bit(int64(off+int32(off2)))) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(int32(read8(sym, off+int32(off2))))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32LtU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32LtU (I32Const [0]) x)
	// result: (I32Eqz (I32Eqz x))
	for {
		if v_0.Op != OpWasmI32Const || auxIntToInt32(v_0.AuxInt) != 0 {
			break
		}
		x := v_1
		v.reset(OpWasmI32Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI32Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (I32LtU x (I32Const [1]))
	// result: (I32Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const || auxIntToInt32(v_1.AuxInt) != 1 {
			break
		}
		v.reset(OpWasmI32Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Mul(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32Mul (I32Const [x]) (I32Const [y]))
	// result: (I32Const [x * y])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(x * y)
		return true
	}
	// match: (I32Mul (I32Const [x]) y)
	// cond: y.Op != OpWasmI32Const
	// result: (I32Mul y (I32Const [x]))
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI32Const) {
			break
		}
		v.reset(OpWasmI32Mul)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Ne(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32Ne (I32Const [x]) (I32Const [y]))
	// cond: x == y
	// result: (I32Const [0])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		if !(x == y) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (I32Ne (I32Const [x]) (I32Const [y]))
	// cond: x != y
	// result: (I32Const [1])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		if !(x != y) {
			break
		}
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (I32Ne (I32Const [x]) y)
	// cond: y.Op != OpWasmI32Const
	// result: (I32Ne y (I32Const [x]))
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI32Const) {
			break
		}
		v.reset(OpWasmI32Ne)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I32Ne x (I32Const [0]))
	// result: (I32Eqz (I32Eqz x))
	for {
		x := v_0
		if v_1.Op != OpWasmI32Const || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasmI32Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI32Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Or(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32Or (I32Const [x]) (I32Const [y]))
	// result: (I32Const [x | y])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(x | y)
		return true
	}
	// match: (I32Or (I32Const [x]) y)
	// cond: y.Op != OpWasmI32Const
	// result: (I32Or y (I32Const [x]))
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI32Const) {
			break
		}
		v.reset(OpWasmI32Or)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Shl(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32Shl (I32Const [x]) (I32Const [y]))
	// result: (I32Const [x << uint32(y)])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(x << uint32(y))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32ShrS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32ShrS (I32Const [x]) (I32Const [y]))
	// result: (I32Const [x >> uint32(y)])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(x >> uint32(y))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32ShrU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32ShrU (I32Const [x]) (I32Const [y]))
	// result: (I32Const [int32(uint32(x) >> uint32(y))])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(int32(uint32(x) >> uint32(y)))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Store(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32Store [off] (I32AddConst [off2] ptr) val mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Store [off+off2] ptr val mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Store16(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32Store16 [off] (I32AddConst [off2] ptr) val mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Store16 [off+off2] ptr val mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Store16)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Store8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I32Store8 [off] (I32AddConst [off2] ptr) val mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I32Store8 [off+off2] ptr val mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpWasmI32AddConst {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI32Store8)
		v.AuxInt = int32ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI32Xor(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I32Xor (I32Const [x]) (I32Const [y]))
	// result: (I32Const [x ^ y])
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpWasmI32Const {
			break
		}
		y := auxIntToInt32(v_1.AuxInt)
		v.reset(OpWasmI32Const)
		v.AuxInt = int32ToAuxInt(x ^ y)
		return true
	}
	// match: (I32Xor (I32Const [x]) y)
	// cond: y.Op != OpWasmI32Const
	// result: (I32Xor y (I32Const [x]))
	for {
		if v_0.Op != OpWasmI32Const {
			break
		}
		x := auxIntToInt32(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI32Const) {
			break
		}
		v.reset(OpWasmI32Xor)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Add(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Add (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x + y])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(x + y)
		return true
	}
	// match: (I64Add (I64Const [x]) y)
	// cond: y.Op != OpWasmI64Const
	// result: (I64Add y (I64Const [x]))
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI64Const) {
			break
		}
		v.reset(OpWasmI64Add)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64And(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64And (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x & y])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(x & y)
		return true
	}
	// match: (I64And (I64Const [x]) y)
	// cond: y.Op != OpWasmI64Const
	// result: (I64And y (I64Const [x]))
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI64Const) {
			break
		}
		v.reset(OpWasmI64And)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Eq(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Eq (I64Const [x]) (I64Const [y]))
	// cond: x == y
	// result: (I64Const [1])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x == y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(1)
		return true
	}
	// match: (I64Eq (I64Const [x]) (I64Const [y]))
	// cond: x != y
	// result: (I64Const [0])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x != y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (I64Eq (I64Const [x]) y)
	// cond: y.Op != OpWasmI64Const
	// result: (I64Eq y (I64Const [x]))
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI64Const) {
			break
		}
		v.reset(OpWasmI64Eq)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I64Eq x (I64Const [0]))
	// result: (I64Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasmI64Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64LeU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64LeU x (I64Const [0]))
	// result: (I64Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasmI64Eqz)
		v.AddArg(x)
		return true
	}
	// match: (I64LeU (I64Const [1]) x)
	// result: (I64Eqz (I64Eqz x))
	for {
		if v_0.Op != OpWasmI64Const || auxIntToInt64(v_0.AuxInt) != 1 {
			break
		}
		x := v_1
		v.reset(OpWasmI64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI64Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Load [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Load)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load16S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load16S [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Load16S [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Load16S)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load16U [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Load16U [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Load16U)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load32S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load32S [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Load32S [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Load32S)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load32U [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Load32U [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Load32U)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load8S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load8S [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Load8S [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Load8S)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load8U [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Load8U [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Load8U)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64LtU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64LtU (I64Const [0]) x)
	// result: (I64Eqz (I64Eqz x))
	for {
		if v_0.Op != OpWasmI64Const || auxIntToInt64(v_0.AuxInt) != 0 {
			break
		}
		x := v_1
		v.reset(OpWasmI64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI64Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (I64LtU x (I64Const [1]))
	// result: (I64Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const || auxIntToInt64(v_1.AuxInt) != 1 {
			break
		}
		v.reset(OpWasmI64Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Mul(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Mul (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x * y])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(x * y)
		return true
	}
	// match: (I64Mul (I64Const [x]) y)
	// cond: y.Op != OpWasmI64Const
	// result: (I64Mul y (I64Const [x]))
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI64Const) {
			break
		}
		v.reset(OpWasmI64Mul)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Ne(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Ne (I64Const [x]) (I64Const [y]))
	// cond: x == y
	// result: (I64Const [0])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x == y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (I64Ne (I64Const [x]) (I64Const [y]))
	// cond: x != y
	// result: (I64Const [1])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x != y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(1)
		return true
	}
	// match: (I64Ne (I64Const [x]) y)
	// cond: y.Op != OpWasmI64Const
	// result: (I64Ne y (I64Const [x]))
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI64Const) {
			break
		}
		v.reset(OpWasmI64Ne)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I64Ne x (I64Const [0]))
	// result: (I64Eqz (I64Eqz x))
	for {
		x := v_0
		if v_1.Op != OpWasmI64Const || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasmI64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI64Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Or(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Or (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x | y])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(x | y)
		return true
	}
	// match: (I64Or (I64Const [x]) y)
	// cond: y.Op != OpWasmI64Const
	// result: (I64Or y (I64Const [x]))
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI64Const) {
			break
		}
		v.reset(OpWasmI64Or)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Shl(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Shl (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x << uint64(y)])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(x << uint64(y))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64ShrS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64ShrS (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x >> uint64(y)])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(x >> uint64(y))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64ShrU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64ShrU (I64Const [x]) (I64Const [y]))
	// result: (I64Const [int64(uint64(x) >> uint64(y))])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(int64(uint64(x) >> uint64(y)))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Store [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Store)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store16(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store16 [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Store16 [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Store16)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store32 [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Store32 [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Store32)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store8 [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(int64(off+off2))
	// result: (I64Store8 [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(int64(off + off2))) {
			break
		}
		v.reset(OpWasmI64Store8)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Xor(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Xor (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x ^ y])
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasmI64Const)
		v.AuxInt = int64ToAuxInt(x ^ y)
		return true
	}
	// match: (I64Xor (I64Const [x]) y)
	// cond: y.Op != OpWasmI64Const
	// result: (I64Xor y (I64Const [x]))
	for {
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasmI64Const) {
			break
		}
		v.reset(OpWasmI64Xor)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] destptr mem)
	// result: (I32Store8 destptr (I32Const [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store8)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(destptr, v0, mem)
		return true
	}
	// match: (Zero [2] destptr mem)
	// result: (I32Store16 destptr (I32Const [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store16)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(destptr, v0, mem)
		return true
	}
	// match: (Zero [4] destptr mem)
	// result: (I32Store destptr (I32Const [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(destptr, v0, mem)
		return true
	}
	// match: (Zero [3] destptr mem)
	// result: (I32Store8 [2] destptr (I32Const [0]) (I32Store16 destptr (I32Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store8)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store16, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [5] destptr mem)
	// result: (I32Store8 [4] destptr (I32Const [0]) (I32Store destptr (I32Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 5 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store8)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [6] destptr mem)
	// result: (I32Store16 [4] destptr (I32Const [0]) (I32Store destptr (I32Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store16)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [7] destptr mem)
	// result: (I32Store [3] destptr (I32Const [0]) (I32Store destptr (I32Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 7 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [8] destptr mem)
	// result: (I32Store [4] destptr (I32Const [0]) (I32Store destptr (I32Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s%4 != 0 && s > 4 && s < 16
	// result: (Zero [s-s%4] (OffPtr <destptr.Type> destptr [s%4]) (I32Store destptr (I32Const [0]) mem))
	for {
		s := auxIntToInt64(v.AuxInt)
		destptr := v_0
		mem := v_1
		if !(s%4 != 0 && s > 4 && s < 16) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = int64ToAuxInt(s - s%4)
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = int64ToAuxInt(s % 4)
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v2.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(destptr, v2, mem)
		v.AddArg2(v0, v1)
		return true
	}
	// match: (Zero [12] destptr mem)
	// result: (I32Store [8] destptr (I32Const [0]) (I32Store [4] destptr (I32Const [0]) (I32Store destptr (I32Const [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 12 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2.AddArg3(destptr, v0, mem)
		v1.AddArg3(destptr, v0, v2)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [16] destptr mem)
	// result: (I32Store [12] destptr (I32Const [0]) (I32Store [8] destptr (I32Const [0]) (I32Store [4] destptr (I32Const [0]) (I32Store destptr (I32Const [0]) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 16 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasmI32Store)
		v.AuxInt = int32ToAuxInt(12)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(8)
		v2 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(4)
		v3 := b.NewValue0(v.Pos, OpWasmI32Store, types.TypeMem)
		v3.AddArg3(destptr, v0, mem)
		v2.AddArg3(destptr, v0, v3)
		v1.AddArg3(destptr, v0, v2)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [s] destptr mem)
	// result: (LoweredZero [s] destptr mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		destptr := v_0
		mem := v_1
		v.reset(OpWasmLoweredZero)
		v.AuxInt = int64ToAuxInt(s)
		v.AddArg2(destptr, mem)
		return true
	}
}
func rewriteValueWasm_OpZeroExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt16to32 x:(I32Load16U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI32Load16U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt16to32 x)
	// result: (I32And x (I32Const [0xffff]))
	for {
		x := v_0
		v.reset(OpWasmI32And)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0xffff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpZeroExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt16to64 x:(I64Load16U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI64Load16U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt16to64 x)
	// result: (I64And x (I64Const [0xffff]))
	for {
		x := v_0
		v.reset(OpWasmI64And)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xffff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpZeroExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt32to64 x:(I64Load32U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI64Load32U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt32to64 x)
	// result: (I64And x (I64Const [0xffffffff]))
	for {
		x := v_0
		v.reset(OpWasmI64And)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xffffffff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpZeroExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt8to16 x:(I32Load8U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI32Load8U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt8to16 x)
	// result: (I32And x (I32Const [0xff]))
	for {
		x := v_0
		v.reset(OpWasmI32And)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0xff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpZeroExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt8to32 x:(I32Load8U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI32Load8U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt8to32 x)
	// result: (I32And x (I32Const [0xff]))
	for {
		x := v_0
		v.reset(OpWasmI32And)
		v0 := b.NewValue0(v.Pos, OpWasmI32Const, typ.Int32)
		v0.AuxInt = int32ToAuxInt(0xff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm_OpZeroExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt8to64 x:(I64Load8U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasmI64Load8U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt8to64 x)
	// result: (I64And x (I64Const [0xff]))
	for {
		x := v_0
		v.reset(OpWasmI64And)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteBlockWasm(b *Block) bool {
	return false
}
