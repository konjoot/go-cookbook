package error_handling

import "testing"

var r error

func BenchmarkStandartErrorHandling(b *testing.B) {
	var e error
	for n := 0; n < b.N; n++ {
		e = StandartErrorHandling()
	}
	r = e
}

func BenchmarkErrorHandlingWithPanic(b *testing.B) {
	var e error

	for n := 0; n < b.N; n++ {
		e = ErrorHandlingWithPanic()
	}

	r = e
}

func BenchmarkErrorHandlingWithPanicIfHelper(b *testing.B) {
	var e error

	for n := 0; n < b.N; n++ {
		e = ErrorHandlingWithPanicIfHelper()
	}
	r = e
}

func BenchmarkErrorHandlingWithPanicIfHelperAndRescueFactory(b *testing.B) {
	var e error

	for n := 0; n < b.N; n++ {
		e = ErrorHandlingWithPanicIfHelperAndRescueFactory()
	}
	r = e
}

func BenchmarkErrorHandlingWithRunner(b *testing.B) {
	var e error

	for n := 0; n < b.N; n++ {
		e = ErrorHandlingWithRunner()
	}
	r = e
}
