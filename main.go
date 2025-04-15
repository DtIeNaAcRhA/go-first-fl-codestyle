package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// для расчёта уровня урона,блокировки,суперсилы
func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

// тренировка атаки
func attack(name, character string) string {
	if character == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(3, 5))
	} else if character == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(5, 10))
	}
	return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randint(-3, -1))

}

// тренировка обороны
func defence(name, character string) string {
	if character == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randint(5, 10))
	} else if character == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randint(-2, 2))
	}
	return fmt.Sprintf("%s блокировал %d урона.", name, 10+randint(2, 5))

}

// тренировка суперсилы
func special(name, character string) string {
	if character == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", name, 80+25)
	} else if character == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", name, 5+40)
	}
	return fmt.Sprintf("%s применил специальное умение `Защита %d`", name, 10+30)

}

// тренировка (если класс персонажа отсутвует в системе, он не может тренироваться )
func training(name, character string) string {
	if character == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", name)
	} else if character == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", name)
	} else if character == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", name)
	} else {
		fmt.Printf("%s, у тебя неизвестный класс персонажа.\n", name)
		return "Выбери персонажа из сущетвующих и возвращайся снова"
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	i := 0
	for cmd != "skip" {
		if i > 6 {
			fmt.Println("Если ты устал, введи команду skip для завершения тренировки")
		}
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(name, character))
		} else if cmd == "defence" {
			fmt.Println(defence(name, character))
		} else if cmd == "special" {
			fmt.Println(special(name, character))
		} else {
			fmt.Println("Этот навык потренировать не получится.\nВыбери команду из существующих:")
			fmt.Println("attack — чтобы атаковать противника,")
			fmt.Println("defence — чтобы блокировать атаку противника,")
			fmt.Println("special — чтобы использовать свою суперсилу.")
			fmt.Println("Если не хочешь тренироваться, введи команду skip.")
		}
		i++
	}
	return "тренировка окончена"
}

// пользователь задаёт имя и выбирает персонажа
func choiseСharacter() string {
	var approveChoice string
	var character string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &character)
		if character == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if character == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if character == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		} else {
			fmt.Println("Такого персонажа нет:(")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return character
}

// главная функция (игра)
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Здравствуй, %s\n", name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	character := choiseСharacter()

	fmt.Println(training(name, character))
}
