# Bin Management System с Dependency Injection

Система управления bin объектами с использованием интерфейсов для dependency injection.

## Архитектура

### Структура проекта

```
├── main.go              # Точка входа приложения
├── models/              # Модели данных
│   └── models.go        # Структура Bin
├── interfaces/          # Интерфейсы для DI
│   └── interfaces.go    # Storage, FileReader, BinService
├── storage/             # Реализация хранилища
│   └── storage.go       # JSONStorage + legacy функции
├── file/                # Работа с файлами
│   └── file.go          # OSFileReader + legacy функции
├── bins/                # Бизнес-логика bins
│   └── bins.go          # BinServiceImpl + legacy функции
└── api/                 # API слой (не изменен)
    └── api.go           # BinAPI
```

### Dependency Injection

Приложение использует интерфейсы для внедрения зависимостей:

1. **Storage Interface** - для работы с хранилищем данных
2. **FileReader Interface** - для чтения файлов  
3. **BinService Interface** - для бизнес-логики bins

### Основные компоненты

#### Models
- `Bin` - основная модель данных в отдельном пакете для избежания циклических импортов

#### Interfaces
- `Storage` - методы для сохранения и загрузки bins
- `FileReader` - методы для работы с файлами
- `BinService` - методы для создания bins и генерации ID

#### Implementations
- `JSONStorage` - реализация хранилища в JSON файле
- `OSFileReader` - реализация для работы с файловой системой
- `BinServiceImpl` - реализация бизнес-логики bins

#### Legacy Support
Все пакеты сохраняют обратную совместимость со старыми функциями:
- `storage.SaveBin()` и `storage.LoadBins()`
- `file.ReadFile()` и `file.IsJSONFile()`
- `bins.NewBin()` и `bins.generateID()`

### Преимущества архитектуры

1. **Разделение ответственности** - каждый модуль отвечает за свою область
2. **Тестируемость** - легко создавать моки для тестирования
3. **Расширяемость** - легко добавлять новые реализации интерфейсов
4. **Слабая связанность** - модули зависят от интерфейсов, а не от конкретных реализаций
5. **Обратная совместимость** - старый код продолжает работать

### Использование

#### Новый подход с интерфейсами:
```go
// Инициализация зависимостей через интерфейсы
storageService := storage.NewJSONStorage()
fileReader := file.NewOSFileReader()
binService := bins.NewBinService()

// Использование через интерфейсы
bin := binService.CreateBin("test", true)
err := storageService.SaveBin(bin)
```

#### Старый подход (все еще работает):
```go
// Прямое использование функций
bin := bins.NewBin("test", true)
storage.SaveBin(bin)
```

### Демонстрация работы

Приложение демонстрирует:
- Создание bin через интерфейс BinService
- Сохранение через интерфейс Storage
- Проверку файла через интерфейс FileReader
- Загрузку всех bins через интерфейс Storage
- Обратную совместимость со старым API