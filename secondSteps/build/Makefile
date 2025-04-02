GO = go
BUILD = build
CMD = cmd
BIN = bin
FILE = ../$(BIN)/main
MAIN_PACKAGE = ../$(CMD)

.PHONY: all clean run

all: run

$(FILE):
	mkdir -p ../$(BIN)
	cd $(MAIN_PACKAGE) && $(GO) $(BUILD) -o $(FILE) .

run: $(FILE)
	./$(FILE)

clean:
	rm -rf ../$(BIN)