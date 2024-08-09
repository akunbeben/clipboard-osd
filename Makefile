# Makefile

# Compiler and compiler flags
CC := go
CFLAGS := build

# Source file
SRC := main.go

# Output directory
OUT_DIR := ./build

# Output binary name
OUT_NAME := clipboard-osd

all: build

build:
	@mkdir -p $(OUT_DIR)
	$(CC) $(CFLAGS) -o $(OUT_DIR)/$(OUT_NAME) $(SRC)

clean:
	@rm -rf $(OUT_DIR)

run:
	@$(OUT_DIR)/$(OUT_NAME)

.PHONY: all build clean