# Makefile

.PHONY: all cl clean

all: cl

now_lesson := $(shell cat now-lesson.txt)

cl:
	@echo "Creating lesson $(now_lesson)"
	mkdir -p lesson$(shell expr $(now_lesson) + 1)
	echo "package main\n\nfunc main() {\n}" > lesson$(shell expr $(now_lesson) + 1)/main.go
	echo $(shell expr $(now_lesson) + 1) > now-lesson.txt

clean:
	rm -rf lesson*

