// Code generated by convert_test-gen.
// DO NOT EDIT!

package convert_test

import (
"testing"
"github.com/katydid/katydid/tests"
)


func BenchmarkProtoKeyAndNameTelephoneProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyAndNameTelephone"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyBridgePepperProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyBridgePepper"]["proto"]
	bench(b, v.Grammar, tests.RandomPuddingMilkshakeProtoParser)
}

func BenchmarkProtoKeyBridgePepperAndFountainTargetProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyBridgePepperAndFountainTarget"]["proto"]
	bench(b, v.Grammar, tests.RandomPuddingMilkshakeProtoParser)
}

func BenchmarkProtoKeyContextPersonProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyContextPerson"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyCorrectNotNameProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyCorrectNotName"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyEmptyOrNilProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyEmptyOrNil"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyIncorrectNotNameProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyIncorrectNotName"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyLenNameProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyLenName"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyListIndexAddressProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyListIndexAddress"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyNilNameProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyNilName"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyOrNameTelephoneProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyOrNameTelephone"]["proto"]
	bench(b, v.Grammar, tests.RandomPersonProtoParser)
}

func BenchmarkProtoKeyRecursiveSrcTreeProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyRecursiveSrcTree"]["proto"]
	bench(b, v.Grammar, tests.RandomSrcTreeProtoParser)
}

func BenchmarkProtoKeyTypewriterPrisonDaisySledProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyTypewriterPrisonDaisySled"]["proto"]
	bench(b, v.Grammar, tests.RandomTypewriterPrisonProtoParser)
}

func BenchmarkProtoKeyTypewriterPrisonMapSharkProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyTypewriterPrisonMapShark"]["proto"]
	bench(b, v.Grammar, tests.RandomTypewriterPrisonProtoParser)
}

func BenchmarkProtoKeyTypewriterPrisonMenuPaperclipProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyTypewriterPrisonMenuPaperclip"]["proto"]
	bench(b, v.Grammar, tests.RandomTypewriterPrisonProtoParser)
}

func BenchmarkProtoKeyTypewriterPrisonScarBusStopProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyTypewriterPrisonScarBusStop"]["proto"]
	bench(b, v.Grammar, tests.RandomTypewriterPrisonProtoParser)
}

func BenchmarkProtoKeyTypewriterPrisonSmileLetterProto(b *testing.B) {
	v := tests.BenchValidators["ProtoKeyTypewriterPrisonSmileLetter"]["proto"]
	bench(b, v.Grammar, tests.RandomTypewriterPrisonProtoParser)
}
