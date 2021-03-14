package streaMaster_test

import (
	"fmt"
	"testing"

	"github.com/JustinTimperio/streaMaster/api"
)

func TestMain(m *testing.M) {
	info, err := api.GetDevices()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(info.String())
}
