package ioReadAll

import (
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
	"io"
)

func init() {
	model.HookMap["ioReadAll"] = new(IoReadAll)
}

type IoReadAll struct {
}

func (h *IoReadAll) Hook() {
	err := gohook.Hook(io.ReadAll, ReadAll, ReadAllT)
	if err != nil {
		fmt.Println(err, "IoReadAll")
	} else {
		fmt.Println("IoReadAll")
	}
}

func (h *IoReadAll) UnHook() {
	gohook.UnHook(io.ReadAll)
}