include ${TEMPLATE}/cpp.mk ${TEMPLATE}/Judger.mk

compile:
	@$(CXX) $(CFLAGS) -o $(PREFIX)/user/$(USER_PROG).out $(PREFIX)/user/$(USER_PROG).$(LANG)

judge:
	$($(CMP)) $(PREFIX)/data/input/$(TEST_NUM).input $(PREFIX)/user/$(TEST_NUM).out.usr $(PREFIX)/data/output/$(TEST_NUM).output $(PREFIX)/user/$(TEST_NUM).judge -appes
