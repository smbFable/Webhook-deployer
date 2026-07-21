package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func Runner() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "git", "pull")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Произошла ошибка при пулле: ", err)
	}
}

// 1: Запуск сервера
// 2: Отправка и лбработка вебхука от гит
// 3: Валидация
// 4: Пулл кода в папку
// 5: Сборка докер образа
// 6: Запуск контейнера
