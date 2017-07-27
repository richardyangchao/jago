package jago

import "math"

func register_java_lang_Double() {
	register("java/lang/Double.doubleToRawLongBits(D)J", Java_jang_lang_Double_doubleToRawLongBits)
	register("java/lang/Double.longBitsToDouble(J)D", Java_jang_lang_Double_longBitsToDouble)
}

// public static native int floatToRawIntBits(float value)
func Java_jang_lang_Double_doubleToRawLongBits(value Double) Long {
	bits := math.Float64bits(float64(value))
	return Long(int64(bits))
}

// public static native int floatToRawIntBits(float value)
func Java_jang_lang_Double_longBitsToDouble(bits Long) Double {
	value := math.Float64frombits(uint64(bits)) // todo
	return Double(value)
}