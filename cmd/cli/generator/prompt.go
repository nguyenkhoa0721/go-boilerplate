package generator

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func Run() {
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
			Label: "File name",
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
	} else if result == "SQL Repo" {
		prompt := promptui.Prompt{
			Label: "Feature",
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		err = CopyDir("scripts/dev/go-code-gen/repo", "internal/"+result+"/data", GetModuleName(), result, "")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		log.Info().Msg("Service template generated successfully")
	} else if result == "Service" {
		prompt := promptui.Prompt{
			Label: "Feature",
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		err = CopyDir("scripts/dev/go-code-gen/service", "internal/"+result+"/domain", GetModuleName(), result, "")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//create folder
		err = os.MkdirAll("internal/"+result+"/domain/interfaces", 0755)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = os.MkdirAll("internal/"+result+"/domain/validation", 0755)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		log.Info().Msg("Service template generated successfully")
	} else if result == "Handler" {
		prompt := promptui.Prompt{
			Label: "Feature",
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		prompt = promptui.Prompt{
			Label: "Domain",
		}
		domain, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		err = CopyDir("scripts/dev/go-code-gen/handler", "internal/"+strcase.ToSnake(result)+"/presenter/"+domain, GetModuleName(), result, domain)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		log.Info().Msg("Handler template generated successfully")
	}
}
