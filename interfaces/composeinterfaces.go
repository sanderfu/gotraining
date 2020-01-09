package interfaces

import (
	"bytes"
	"fmt"
	"io"
)

//PlayCompositeInterfaces tests composite interfaces
func PlayCompositeInterfaces() {
	var wc writerCloser = newBufferedWriterCloser()
	wc.write([]byte("Hello world, this is a test"))
	wc.close()

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	//Empty interface
	var myObj interface{} = newBufferedWriterCloser()

	//Empty interface must be typecast to something useful
	if wc, ok := myObj.(writerCloser); ok {
		wc.write([]byte("Hello Go world!"))
		wc.close()
	}

	k, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(k)
	} else {
		fmt.Println("Conversion failed")
	}

	//Type switches
	var i interface{} = "0"
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what I is")
	}

	var wc2 writerCloser = &myWriterCloser{}
	fmt.Println(wc2)
	/*
		Comment on this:
		When implementing an interface, if value type used,
		all methods that implement the value type has to
		have value recievers

		If implements interface with pointer, then just have
		to have methods there, regardless of wether the
		methods have value reciever or pointer reciever.

	*/

	/*
		Best practices on interfaces:
		- Use many, small interfaces
			- io.Writer, io.Reader, interface{}

		- Don't export interfaces for types
		that will be consumed by others (they will have to implement
		methods they might not need)

		- DO export interfaces for types that will be
		used by package. (GO has implicit implementation)

		- Design functions and methods to recieve interfaces
		whenever possible.

	*/
}

//writer already defined in interfaces.go

type closer interface {
	close() error
}

//To implement a composed interface, we must simplyu implement all methods of the underlying interfaces for this composite interface!
type writerCloser interface {
	writer
	closer
}

type bufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *bufferedWriterCloser) write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *bufferedWriterCloser) close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func newBufferedWriterCloser() *bufferedWriterCloser {
	return &bufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type myWriterCloser struct{}

func (mwc *myWriterCloser) write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) close() error {
	return nil
}
