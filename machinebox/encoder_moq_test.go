// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package function

import (
	"bytes"
	"image/draw"
	"io"
	"sync"
)

var (
	lockImageEncoderDecoderMockDecode sync.RWMutex
	lockImageEncoderDecoderMockEncode sync.RWMutex
)

// ImageEncoderDecoderMock is a mock implementation of ImageEncoderDecoder.
//
//     func TestSomethingThatUsesImageEncoderDecoder(t *testing.T) {
//
//         // make and configure a mocked ImageEncoderDecoder
//         mockedImageEncoderDecoder := &ImageEncoderDecoderMock{
//             DecodeFunc: func(img io.Reader) (draw.Image, error) {
// 	               panic("TODO: mock out the Decode method")
//             },
//             EncodeFunc: func(b *bytes.Buffer, img draw.Image)  {
// 	               panic("TODO: mock out the Encode method")
//             },
//         }
//
//         // TODO: use mockedImageEncoderDecoder in code that requires ImageEncoderDecoder
//         //       and then make assertions.
//
//     }
type ImageEncoderDecoderMock struct {
	// DecodeFunc mocks the Decode method.
	DecodeFunc func(img io.Reader) (draw.Image, error)

	// EncodeFunc mocks the Encode method.
	EncodeFunc func(b *bytes.Buffer, img draw.Image)

	// calls tracks calls to the methods.
	calls struct {
		// Decode holds details about calls to the Decode method.
		Decode []struct {
			// Img is the img argument value.
			Img io.Reader
		}
		// Encode holds details about calls to the Encode method.
		Encode []struct {
			// B is the b argument value.
			B *bytes.Buffer
			// Img is the img argument value.
			Img draw.Image
		}
	}
}

// Decode calls DecodeFunc.
func (mock *ImageEncoderDecoderMock) Decode(img io.Reader) (draw.Image, error) {
	if mock.DecodeFunc == nil {
		panic("moq: ImageEncoderDecoderMock.DecodeFunc is nil but ImageEncoderDecoder.Decode was just called")
	}
	callInfo := struct {
		Img io.Reader
	}{
		Img: img,
	}
	lockImageEncoderDecoderMockDecode.Lock()
	mock.calls.Decode = append(mock.calls.Decode, callInfo)
	lockImageEncoderDecoderMockDecode.Unlock()
	return mock.DecodeFunc(img)
}

// DecodeCalls gets all the calls that were made to Decode.
// Check the length with:
//     len(mockedImageEncoderDecoder.DecodeCalls())
func (mock *ImageEncoderDecoderMock) DecodeCalls() []struct {
	Img io.Reader
} {
	var calls []struct {
		Img io.Reader
	}
	lockImageEncoderDecoderMockDecode.RLock()
	calls = mock.calls.Decode
	lockImageEncoderDecoderMockDecode.RUnlock()
	return calls
}

// Encode calls EncodeFunc.
func (mock *ImageEncoderDecoderMock) Encode(b *bytes.Buffer, img draw.Image) {
	if mock.EncodeFunc == nil {
		panic("moq: ImageEncoderDecoderMock.EncodeFunc is nil but ImageEncoderDecoder.Encode was just called")
	}
	callInfo := struct {
		B   *bytes.Buffer
		Img draw.Image
	}{
		B:   b,
		Img: img,
	}
	lockImageEncoderDecoderMockEncode.Lock()
	mock.calls.Encode = append(mock.calls.Encode, callInfo)
	lockImageEncoderDecoderMockEncode.Unlock()
	mock.EncodeFunc(b, img)
}

// EncodeCalls gets all the calls that were made to Encode.
// Check the length with:
//     len(mockedImageEncoderDecoder.EncodeCalls())
func (mock *ImageEncoderDecoderMock) EncodeCalls() []struct {
	B   *bytes.Buffer
	Img draw.Image
} {
	var calls []struct {
		B   *bytes.Buffer
		Img draw.Image
	}
	lockImageEncoderDecoderMockEncode.RLock()
	calls = mock.calls.Encode
	lockImageEncoderDecoderMockEncode.RUnlock()
	return calls
}
