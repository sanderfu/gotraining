package main

import (
	"fmt"

	"github.com/sanderfu/gotraining/channels"
	"github.com/sanderfu/gotraining/goroutines"

	"github.com/sanderfu/gotraining/interfaces"

	"github.com/sanderfu/gotraining/functions"

	"github.com/sanderfu/gotraining/pointers"

	"github.com/sanderfu/gotraining/looping"

	"github.com/sanderfu/gotraining/flowcontrol"

	"github.com/sanderfu/gotraining/mapsandstructs"
)

func main() {
	fmt.Println("Hello Go!")
	mapsandstructs.MapPlay1()
	mapsandstructs.StructPlay1()
	mapsandstructs.StructPlay2()
	flowcontrol.FlowPlay1()
	looping.LoopPlay1()
	flowcontrol.DeferPlay1()
	flowcontrol.DeferPlay2()
	flowcontrol.DeferPlay3()
	flowcontrol.DeferPlay4()
	//flowcontrol.PanicPlay1()
	//flowcontrol.PanicPlay2()
	//flowcontrol.PanicPlay3()
	flowcontrol.PanicPlay4()
	fmt.Println("Proceeded from error")
	pointers.PointerPlay1()
	pointers.PointerPlay2()
	pointers.PointerPlay3()
	functions.PlayFunctions()
	interfaces.PlayInterfaces()
	interfaces.PlayCompositeInterfaces()
	goroutines.PlayGoroutines()
	channels.PlayChannels1()
	fmt.Println()
	channels.PlayChannels2()
	fmt.Println()
	channels.PlayChannels3()
	fmt.Println()
	channels.PlayChannels4()
	fmt.Println()
	channels.PlayChannels5()
	fmt.Println()
	channels.PlayChannels6()
	fmt.Println()
	channels.PlayChannels7()
}
