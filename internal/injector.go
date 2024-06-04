package internal

import (
	"fmt"

	framework "example.com/on_path_robotics2/framework/dto"
)

func GetUseCase(dto framework.GetRobotsDTO) {
	fmt.Println("Use case obtained from injector file")
}

func GetComponent() {
	fmt.Println("Component obtained from injector file")
}
