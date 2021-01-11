package confreader

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

type App struct {
	Port         int32  `yaml:"port"`
	BD_url       string `yaml:"bd_url"`
	Jaeger_url   string `yaml:"jaeger_url"`
	Sentry_url   string `yaml:"sentry_url"`
	Kafka_broker string `yaml:"kafka_broker"`
	Some_app_id  string `yaml:"some_app_id"`
	Some_app_key string `yaml:"some_app_key"`
}

func ConfReader() {
	// Создаем файловый дескриптор для файла, в котором хрнаится json-конфигурация
	file, err := os.Open("./conf.yaml")
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	// Не забываем закрыть файл при выходе из функции
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	// Задаем переменную, в которую будем считывать конфиг
	app := App{}

	// Задаем декодер из файла и сразу же вызываем его
	err = yaml.NewDecoder(file).Decode(&app)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	// валлидация url Jaeger_url
	_, err = url.ParseRequestURI(app.Jaeger_url)
	if err != nil {
		log.Fatalf("Параметр Jaeger_url не соответсвует URL")
	}

	// валлидация url Sentry_url
	_, err = url.ParseRequestURI(app.Sentry_url)
	if err != nil {
		log.Fatalf("Параметр Sentry_url не соответсвует URL")
	}

	// Для проверки -  в app  доступны параметры из файла
	fmt.Printf("\n%v\n", app)
}
