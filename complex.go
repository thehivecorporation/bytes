package bytes

func Complex64ToBytes(d complex64) []byte {
	f32r := real(d)
	f32i := imag(d)

	return append(Float32ToBytes(f32r), Float32ToBytes(f32i)...)
}

func Complex128ToBytes(d complex128) []byte {
	f64i := imag(d)
	f64r := real(d)

	return append(Float64ToBytes(f64r), Float64ToBytes(f64i)...)
}
