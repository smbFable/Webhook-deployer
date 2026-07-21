package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func Runner(cloneURL string, path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка получения доступа к домашней директории: ", err)
	}
	targetPath := filepath.Join(homeDir, "Projects", "golang", path)

	gitPath := filepath.Join(targetPath, ".git")
	_, err = os.Stat(gitPath)

	if os.IsNotExist(err) {
		cmd := exec.CommandContext(ctx, "git", "clone", cloneURL, targetPath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Лог ошибки: ", string(output))
			return err
		}
		fmt.Println("Клонирование завершено")
		return nil

	} else if err == nil {
		fetchcmd := exec.CommandContext(ctx, "git", "fetch", "--all")
		fetchcmd.Dir = targetPath
		fetchOutput, err := fetchcmd.CombinedOutput()
		if err != nil {
			fmt.Println("Лог ошибки: ", string(fetchOutput))
			return err
		}

		resetCmd := exec.CommandContext(ctx, "git", "reset", "--hard", "origin/main")
		resetCmd.Dir = targetPath
		if output, err := resetCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("Ошибка reset: %v\nЛоги: %s", err, string(output))
		}

		cleanCmd := exec.CommandContext(ctx, "git", "clean", "-fd")
		cleanCmd.Dir = targetPath
		if output, err := cleanCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("Ошибка clean: %v\nЛоги: %s", err, string(output))
		}

	} else {
		return fmt.Errorf("ошибка проверки директории: %w", err)
	}
	fmt.Println("Успешно!!!")
	return nil
}

// 1: Запуск сервера
// 2: Отправка и лбработка вебхука от гит
// 3: Валидация
// 4: Пулл кода в папку
// 5: Сборка докер образа
// 6: Запуск контейнера
