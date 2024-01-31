package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"time"
)

func main() {
	prompt := promptui.Select{
		Label: "Select",
		Items: []string{"SQL Migration", "SQL Repo", "Service", "Handler", "Lib"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "SQL Migration" {
		prompt := promptui.Prompt{
			Label: "File name: ",
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		now := time.Now()
		year := fmt.Sprint(now.Year())
		month := "0" + fmt.Sprint(int(now.Month()))
		if len(month) > 2 {
			month = month[1:]
		}
		day := "0" + fmt.Sprint(now.Day())
		if len(day) > 2 {
			day = day[1:]
		}
		hour := "0" + fmt.Sprint(now.Hour())
		if len(hour) > 2 {
			hour = hour[1:]
		}
		minute := "0" + fmt.Sprint(now.Minute())
		if len(minute) > 2 {
			minute = minute[1:]
		}
		second := "0" + fmt.Sprint(now.Second())
		if len(second) > 2 {
			second = second[1:]
		}

		_, err = os.Create(fmt.Sprintf("scripts/migration/%s%s%s%s%s%s_%s.up.sql", year, month, day, hour, minute, second, result))
		_, err = os.Create(fmt.Sprintf("scripts/migration/%s%s%s%s%s%s_%s.down.sql", year, month, day, hour, minute, second, result))
		if err != nil {
			fmt.Printf("File creation failed %v\n", err)
			return
		}
	}
}
