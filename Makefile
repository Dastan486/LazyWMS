# Определяем переменные
BINARY_NAME = fs
BINARY_DIR = bin
SOURCE_DIR = .

# Правило сборки
build:
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(SOURCE_DIR)/cmd/server

# Правило запуска
run: build
	@./$(BINARY_DIR)/$(BINARY_NAME)

# Правило тестирования
test:
	@go test ./... -v
