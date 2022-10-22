# 默认环境变量有 PREFIX, TEMPLATE, TESTLIB
include ${TEMPLATE}/c.mk ${TEMPLATE}/Judger.mk

# 评测分四个阶段
# 1. prebuild: 用于提前生成测试数据、评测器、spj等工具，runner 只执行一次
#              只有 ./data, ./judge 目录可见
# 2. compile:  用于编译用户提交的程序
#              只有 ./user/$(USER_PROG).$(LANG) 和 ./judge 目录可见
# 3. run:      运行用户程序
#              只有 ./data/input/*.input 和 ./user/$(USER_PROG).out 可见
#              用户输出存放于 ./user/?.out.usr
#              使用 woj-sandbox 运行，等效于 $(PREFIX)/user/$(USER_PROG).out < $(PREFIX)/data/input/$(TEST_NUM).input > $(PREFIX)/user/$(TEST_NUM).out.usr
# 4. judge:    用于判定输出结果 环境变量 TEST_NUM 表示当前测试点编号
#              所有目录 ./data ./judge ./user 可见

compile:
	$(CC) $(CFLAGS) -o $(PREFIX)/user/$(USER_PROG).out $(PREFIX)/user/$(USER_PROG).$(LANG) $(PREFIX)/judge/gadget.c

judge:
	# Rename on *.out.usr or *.judge is not allowed
	sed '/gadgets/d' $(PREFIX)/user/$(TEST_NUM).out.usr > $(PREFIX)/user/$(TEST_NUM).out.usr1
	$(NCMP) $(PREFIX)/data/input/$(TEST_NUM).input $(PREFIX)/user/$(TEST_NUM).out.usr1 $(PREFIX)/data/output/$(TEST_NUM).output $(PREFIX)/user/$(TEST_NUM).judge -appes
