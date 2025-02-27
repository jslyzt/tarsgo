//Package statf comment
// This file war generated by tars2go 1.1
// Generated from StatF.tars
package statf

import (
	"fmt"

	"github.com/jslyzt/tarsgo/tars/protocol/codec"
)

//StatMicMsgHead strcut implement
type StatMicMsgHead struct {
	MasterName    string `json:"masterName"`
	SlaveName     string `json:"slaveName"`
	InterfaceName string `json:"interfaceName"`
	MasterIp      string `json:"masterIp"`
	SlaveIp       string `json:"slaveIp"`
	SlavePort     int32  `json:"slavePort"`
	ReturnValue   int32  `json:"returnValue"`
	SlaveSetName  string `json:"slaveSetName"`
	SlaveSetArea  string `json:"slaveSetArea"`
	SlaveSetID    string `json:"slaveSetID"`
	TarsVersion   string `json:"tarsVersion"`
}

func (st *StatMicMsgHead) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *StatMicMsgHead) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.MasterName, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SlaveName, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.InterfaceName, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.MasterIp, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SlaveIp, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.SlavePort, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ReturnValue, 6, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SlaveSetName, 7, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SlaveSetArea, 8, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SlaveSetID, 9, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.TarsVersion, 10, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *StatMicMsgHead) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require StatMicMsgHead, but not exist. tag %d", tag)
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
func (st *StatMicMsgHead) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.MasterName, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SlaveName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.InterfaceName, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.MasterIp, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SlaveIp, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.SlavePort, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ReturnValue, 6)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SlaveSetName, 7)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SlaveSetArea, 8)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SlaveSetID, 9)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.TarsVersion, 10)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *StatMicMsgHead) WriteBlock(_os *codec.Buffer, tag byte) error {
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
