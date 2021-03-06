package actions

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/driusan/de/demodel"
)

var ErrNoFile = errors.New("Can not save empty filename or nil buffer.")

func SaveFile(From, To demodel.Position, buff *demodel.CharBuffer) error {
	if buff == nil || buff.Filename == "" {
		buff.AppendTag("\nNo file to Save")
		return ErrNoFile
	}

	// we don't care about positions, just write the file
	err := ioutil.WriteFile(buff.Filename, buff.Buffer, 0644)
	if err != nil {
		buff.AppendTag(fmt.Sprintf("\n%v", err))
		//fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	buff.AppendTag(fmt.Sprintf("\nSaved %s", buff.Filename))
	buff.Dirty = false
	return nil

}
