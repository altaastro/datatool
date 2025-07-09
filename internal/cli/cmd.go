package cli

import (
	"fmt"
	"os"

	"asys/datatool/internal/upload"

	"github.com/spf13/cobra"
)

var ConfigPath string

var rootCmd = &cobra.Command{
	Use:   "datatool",
	Short: "CLI утилита для загрузки файлов",
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "configs/config.yaml", "Путь к YAML-конфигу")

	rootCmd.AddCommand(uploadCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Загрузить файл на сервер",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadConfig(ConfigPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка загрузки конфига:", err)
			os.Exit(1)
		}

		fmt.Printf("Загрузка: %s → %s\n", cfg.FilePath, cfg.ServerURL)

		if err := upload.SendFile(cfg.FilePath, cfg.ServerURL); err != nil {
			fmt.Fprint(os.Stderr, "Ошибка:", err)
			os.Exit(1)
		}

		fmt.Println("Файл успешно отправлен")
	},
}
