package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)

type Location struct {

}

func (o *Location) OnRead(r ResourceModel, resourceId int) ResourceValue {
    return core.NewEmptyValue()
}
