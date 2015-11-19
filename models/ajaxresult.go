package models

import (
	"encoding/json"
)

type AjaxResult struct {
	Success bool
	Msg interface{}
	Error string
}

func (s *AjaxResult) ToString() string {
	str ,err := json.Marshal(&s)
	if err != nil {
		return "{Success:false, Msg:\"\", Error:\"" + err.Error() + "\"}"
	}
	return string(str);
}