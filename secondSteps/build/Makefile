GO = go
BUILD = build
CMD = cmd
BIN = bin
FILE = ../$(BIN)/main
MAIN_PACKAGE = ../$(CMD)
SRC = $(wildcard $(MAIN_PACKAGE)/*.go)

.PHONY: all clean run

all: run

$(FILE): $(SRC)
	mkdir -p ../$(BIN)
	cd $(MAIN_PACKAGE) && $(GO) $(BUILD) -o $(FILE) .

run: $(FILE)
	./$(FILE)

clean:
	rm -rf ../$(BIN)