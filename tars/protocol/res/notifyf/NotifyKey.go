//Package notifyf comment
// This file war generated by tars2go 1.1
// Generated from NotifyF.tars
package notifyf

import (
	"fmt"

	"github.com/jslyzt/tarsgo/tars/protocol/codec"
)

//NotifyKey strcut implement
type NotifyKey struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Page int32  `json:"page"`
}

func (st *NotifyKey) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *NotifyKey) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Name, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Ip, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Page, 3, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *NotifyKey) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	have, err = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NotifyKey, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *NotifyKey) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Name, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Ip, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Page, 3)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *NotifyKey) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
