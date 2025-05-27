package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		fmt.Println("__ Калькулятор индекса массы тела __ ")
		userWeight, userHeight := getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(IMT)
		if !breakeApp() {
			break
		}
	}

}

func outputResult(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальная масса тела")
	case IMT < 30:
		fmt.Println("У вас избыточная масса тела")
	case IMT < 35:
		fmt.Println("У вас 1-я степень ожирения")
	case IMT < 40:
		fmt.Println("У вас 2-я степень ожирения")
	default:
		fmt.Println("У вас 3-я степень ожирения")
	}
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", IMT)
	fmt.Print(result)
}

func calculateIMT(userWeight float64, userHeight float64) (IMT float64, err error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("Не указан вес или рост")
	}
	IMT = userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (userHeight float64, userWeight float64) {
	fmt.Println("Введите свой рост(см):")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес:")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}

func breakeApp() bool {
	var answerForCalc string
	fmt.Println("Хотите рассчитать индекс повторно? y/n")
	fmt.Scan(&answerForCalc)
	if answerForCalc == "y" || answerForCalc == "Y" {
		return true
	}
	return false
}
