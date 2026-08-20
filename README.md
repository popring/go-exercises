# 命令行 Todo（Go 阶段一验收项目）

> 覆盖：struct / slice / map / 函数 / 指针 / 方法 / 包 / 标准库（除并发）。
> 做出来 = 阶段一真的过关。

## 需求

### 数据结构
```go
Todo{ID int; Title string; Done bool; CreatedAt time.Time}
```
用 slice 存所有 Todo。

### 方法（挂在一个管理者类型上，注意 receiver 用值还是指针）
- `Add(title)` — 新增
- `Complete(id)` — 标记完成
- `Delete(id)` — 删除
- `List()` — 列出全部

### CLI 交互
- `bufio.Scanner` 读 stdin
- `switch` 分发命令：`add xxx` / `done 1` / `del 1` / `list` / `quit`

### 持久化
- 退出时 `json.Marshal` 写 `todos.json`
- 启动时读回来（文件不存在要能正常启动，别 panic）

### 结构
- 拆成 `main` 包 + `todo` 包（注意首字母大写才能导出）

## 验收标准
- [ ] add / done / del / list 全部能用
- [ ] quit 后重启，数据还在
- [ ] 错误分支真的处理了（id 不存在、JSON 读写失败）——不许只写 happy path

## 自己容易踩的坑（练习时暴露过）
1. **错误分支被吞**：写 `if err != nil { 处理; return }`，happy path 靠左，不要 `if err == nil { ... }`
2. **`for range` 是复印件**：改元素要 `todos[i].Done = true`，不能 `for _, t := range` 里改 `t`
3. `error` 不要当变量名，用 `err`

## 跑法
```bash
go run .
```
