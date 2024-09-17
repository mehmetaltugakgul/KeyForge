package main

import (
	"fmt"
	"math/rand"
	"unicode"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)


var (
	includeUppercase    bool = true
	includeLowercase    bool = true
	includeNumbers      bool = true
	includeSpecialChars bool = true
	excludeAmbiguous    bool = false 
	passwordLength      int  = 12
	minOneOfEach        bool = false
	generateButton      *widget.Button
)



func atLeastOneSelected() bool {
	return includeUppercase || includeLowercase || includeNumbers || includeSpecialChars
}


func updateGenerateButtonState(generateButton *widget.Button) {
	generateButton.Disable()
	if atLeastOneSelected() {
		generateButton.Enable()
	}
}

func generatePassword(length int, uppercase, lowercase, numbers, specialChars, minOneOfEach, excludeAmbiguous bool) string {
	if !uppercase && !lowercase && !numbers && !specialChars {
		return ""
	}

	uppercasePool := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercasePool := "abcdefghijklmnopqrstuvwxyz"
	numbersPool := "0123456789"
	specialPool := "!@#$%^&*()-_+="

	if excludeAmbiguous {
		uppercasePool = "ABCDEFGHJKLMNPQRSTUVWXYZ"
		lowercasePool = "abcdefghijkmnpqrstuvwxyz"
		numbersPool = "23456789"
		specialPool = "!@#$%^&*()-_+"
	}

	characterSet := ""
	if uppercase {
		characterSet += uppercasePool
	}
	if lowercase {
		characterSet += lowercasePool
	}
	if numbers {
		characterSet += numbersPool
	}
	if specialChars {
		characterSet += specialPool
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)

	if minOneOfEach {
		requiredTypes := 0
		if uppercase {
			requiredTypes++
		}
		if lowercase {
			requiredTypes++
		}
		if numbers {
			requiredTypes++
		}
		if specialChars {
			requiredTypes++
		}

		if length < requiredTypes {
			
			minOneOfEach = false
		} else {
			remainingLength := length - requiredTypes
			currentIndex := 0

			if uppercase {
				password[currentIndex] = uppercasePool[rand.Intn(len(uppercasePool))]
				currentIndex++
			}
			if lowercase {
				password[currentIndex] = lowercasePool[rand.Intn(len(lowercasePool))]
				currentIndex++
			}
			if numbers {
				password[currentIndex] = numbersPool[rand.Intn(len(numbersPool))]
				currentIndex++
			}
			if specialChars {
				password[currentIndex] = specialPool[rand.Intn(len(specialPool))]
				currentIndex++
			}

			
			for i := 0; i < remainingLength; i++ {
				password[currentIndex] = characterSet[rand.Intn(len(characterSet))]
				currentIndex++
			}

		
			rand.Shuffle(len(password), func(i, j int) {
				password[i], password[j] = password[j], password[i]
			})
		}
	} else {
		
		for i := range password {
			password[i] = characterSet[rand.Intn(len(characterSet))]
		}
	}

	return string(password)
}

func calculatePasswordStrength(password string) float64 {
	var score float64
	length := len(password)

	if length > 8 {
		score += float64(length) * 0.5
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasUpper {
		score += 10
	}
	if hasLower {
		score += 10
	}
	if hasNumber {
		score += 10
	}
	if hasSpecial {
		score += 15
	}


	varietyCount := 0
	if hasUpper {
		varietyCount++
	}
	if hasLower {
		varietyCount++
	}
	if hasNumber {
		varietyCount++
	}
	if hasSpecial {
		varietyCount++
	}
	score += float64(varietyCount) * 10

	normalizedScore := score / 100

	if normalizedScore > 1 {
		normalizedScore = 1
	}

	return normalizedScore
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("KeyForge")

	myWindow.Hide()

	clipboard := myWindow.Clipboard()

	icon, err := fyne.LoadResourceFromPath("passwd.png")
	
	if err != nil {
		fmt.Println("Error loading icon:", err)
	}

	if desk, ok := myApp.(desktop.App); ok {
		menu := fyne.NewMenu("Password Generator",
			fyne.NewMenuItem("Show", func() {
				myWindow.Show()
			}),
			
		)
		desk.SetSystemTrayMenu(menu)
		if icon != nil {
			desk.SetSystemTrayIcon(icon)
		}
	} else {
		fmt.Println("Desktop features not supported")
	}


	lengthSlider := widget.NewSlider(4, 128)
	lengthSlider.SetValue(12)
	lengthSlider.Step = 1
	
	lengthLabel := widget.NewLabel("Password Length: 12")
	lengthSlider.OnChanged = func(value float64) {
		passwordLength = int(value)
		lengthLabel.SetText(fmt.Sprintf("Password Length: %d", passwordLength))
	}

	lettersCheck := widget.NewCheck("Include Letters", func(checked bool) {
		
		
	})
	lettersCheck.SetChecked(true)

	specialCharsCheck := widget.NewCheck("Include Special Characters", func(checked bool) {
		includeSpecialChars = checked
	})
	specialCharsCheck.SetChecked(true)

	uppercaseCheck := widget.NewCheck("Include Uppercase", func(checked bool) {
		includeUppercase = checked
		if generateButton != nil {
			updateGenerateButtonState(generateButton)
		}
	})
	uppercaseCheck.SetChecked(includeUppercase)

	lowercaseCheck := widget.NewCheck("Include Lowercase", func(checked bool) {
		includeLowercase = checked
		if generateButton != nil {
			updateGenerateButtonState(generateButton)
		}
	})
	lowercaseCheck.SetChecked(includeLowercase)

	numbersCheck := widget.NewCheck("Include Numbers", func(checked bool) {
		includeNumbers = checked
		if generateButton != nil {
			updateGenerateButtonState(generateButton)
		}
	})
	numbersCheck.SetChecked(includeNumbers)

	minOneOfEachCheck := widget.NewCheck("At Least One of Each Selected Type", func(checked bool) {
		minOneOfEach = checked
	})

	strengthProgress := widget.NewProgressBar()
	strengthProgress.Hide()

	
	strengthLabel := widget.NewLabel("")
	strengthLabel.Hide()

	passwordOutput := widget.NewEntry()
	passwordOutput.Disable()


	generateButton = widget.NewButton("Generate Password", func() {
		password := generatePassword(passwordLength, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars, minOneOfEach, excludeAmbiguous)
		passwordOutput.SetText(password)
		
		strength := calculatePasswordStrength(password)
		strengthProgress.SetValue(strength)
		strengthLabel.SetText(fmt.Sprintf("Password Strength: %.0f%%", strength*100))
	
		strengthProgress.Show()
		
	
	})


	copyButton := widget.NewButton("Copy Password", func() {
		clipboard.SetContent(passwordOutput.Text)
	})


	exitButton := widget.NewButton("Exit", func() {
		myApp.Quit()
	})


	excludeAmbiguousCheck := widget.NewCheck("Exclude Ambiguous Characters", func(checked bool) {
		excludeAmbiguous = checked
	})

	content := container.NewVBox(
		lengthLabel,
		lengthSlider,
		uppercaseCheck,
		lowercaseCheck,
		numbersCheck,
		specialCharsCheck,
		minOneOfEachCheck,
		excludeAmbiguousCheck,
		generateButton,
		widget.NewLabel("Generated Password:"),
		passwordOutput,
		strengthProgress,
		strengthLabel,
		copyButton,
		exitButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 300))

	myWindow.SetFixedSize(true)
	myWindow.CenterOnScreen()

	
	myWindow.SetCloseIntercept(func() {
		myWindow.Hide()
	})

	
	myApp.Run()
}
