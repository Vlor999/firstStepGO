GO = go
BUILD = build
SRC = src
BIN = bin
FILE = $(BIN)/main
MAIN = $(SRC)/main.go

.PHONY: all clean

all:run

$(FILE): $(MAIN)
	$(GO) $(BUILD) -o $(FILE) $(MAIN)

run: $(FILE)
	./$(FILE)

clean:
	rm -rf $(BIN)
