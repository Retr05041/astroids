APP_NAME := Astroids
SRC_DIR := .
BUILD_DIR := ./bin

all: build

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

run: build
	$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)

.PHONY: all build run clean
