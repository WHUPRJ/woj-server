include ${TEMPLATE}/c.mk ${TEMPLATE}/Judger.mk

prebuild:
	clang++ -I$(TESTLIB) -Ofast -o $(PREFIX)/judge/gen.out $(PREFIX)/judge/gen.cpp
	@if [ ! -f $(PREFIX)/data/input/2.input ]; then \
		$(PREFIX)/judge/gen.out > $(PREFIX)/data/input/2.input; \
		python3 -c "print(sum(map(int, input().split())))" < $(PREFIX)/data/input/2.input > $(PREFIX)/data/output/2.output; \
	fi
	@if [ ! -f $(PREFIX)/data/input/4.input ]; then \
		$(PREFIX)/judge/gen.out > $(PREFIX)/data/input/4.input; \
		python3 -c "print(sum(map(int, input().split())))" < $(PREFIX)/data/input/4.input > $(PREFIX)/data/output/4.output; \
	fi
