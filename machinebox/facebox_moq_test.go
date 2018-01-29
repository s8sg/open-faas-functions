// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package function

import (
	"github.com/machinebox/sdk-go/facebox"
	"io"
	"sync"
)

var (
	lockFaceboxCheckerMockCheck sync.RWMutex
)

// FaceboxCheckerMock is a mock implementation of FaceboxChecker.
//
//     func TestSomethingThatUsesFaceboxChecker(t *testing.T) {
//
//         // make and configure a mocked FaceboxChecker
//         mockedFaceboxChecker := &FaceboxCheckerMock{
//             CheckFunc: func(image io.Reader) ([]facebox.Face, error) {
// 	               panic("TODO: mock out the Check method")
//             },
//         }
//
//         // TODO: use mockedFaceboxChecker in code that requires FaceboxChecker
//         //       and then make assertions.
//
//     }
type FaceboxCheckerMock struct {
	// CheckFunc mocks the Check method.
	CheckFunc func(image io.Reader) ([]facebox.Face, error)

	// calls tracks calls to the methods.
	calls struct {
		// Check holds details about calls to the Check method.
		Check []struct {
			// Image is the image argument value.
			Image io.Reader
		}
	}
}

// Check calls CheckFunc.
func (mock *FaceboxCheckerMock) Check(image io.Reader) ([]facebox.Face, error) {
	if mock.CheckFunc == nil {
		panic("moq: FaceboxCheckerMock.CheckFunc is nil but FaceboxChecker.Check was just called")
	}
	callInfo := struct {
		Image io.Reader
	}{
		Image: image,
	}
	lockFaceboxCheckerMockCheck.Lock()
	mock.calls.Check = append(mock.calls.Check, callInfo)
	lockFaceboxCheckerMockCheck.Unlock()
	return mock.CheckFunc(image)
}

// CheckCalls gets all the calls that were made to Check.
// Check the length with:
//     len(mockedFaceboxChecker.CheckCalls())
func (mock *FaceboxCheckerMock) CheckCalls() []struct {
	Image io.Reader
} {
	var calls []struct {
		Image io.Reader
	}
	lockFaceboxCheckerMockCheck.RLock()
	calls = mock.calls.Check
	lockFaceboxCheckerMockCheck.RUnlock()
	return calls
}