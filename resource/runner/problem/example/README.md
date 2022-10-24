# 示例题目

## 文件构成

```
.
├── config.json            # 题目配置信息
├── data                   # 数据目录
│   ├── input              # 输入数据
│   │   ├── (x).input      # 第 x 组输入数据
│   │   └── ...
│   └── output             # 输出数据
│       ├── (x).output     # 第 x 组输出数据
│       └── ...
└── judge                  # 评测脚本目录
    ├── c.Makefile         # (optional) 自定义评测脚本
    ├── prebuild.Makefile  # (optional) 题目初始化脚本
    └── ...
```

## 详细说明

### 题目配置信息

```json5
{
    "Runtime": {
        // 运行时配置
        "TimeLimit": 1000, // 时间限制 (ms)
        "MemoryLimit": 16, // 内存限制 (MB)
        "NProcLimit": 1     // 进(线)程 限制
    },
    "Languages": [
        {"Lang": "c", "Type": "custom", "Script": "XYZ.Makefile", "Cmp": ""},
        {"Lang": "cpp", "Type": "default", "Script": "", "Cmp": "NCMP"}
    ], // 支持的语言
    "Tasks": [
        // 评测点信息
        {"Id": 1, "Points": 25}, // 第一个评测点，分值 25 分，使用 ./data/1.? 为测试数据
        {"Id": 2, "Points": 25},
        {"Id": 3, "Points": 25},
        {"Id": 4, "Points": 25}
    ]
}
```

### 评测脚本

见注释















