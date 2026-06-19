# sing-box 使用分析报告

> 基于 sing-box v1.13.12 源码分析，对照官方文档验证。
> 分析日期：2026-06-01

## 1. 配置文件格式

**结论：只支持 JSON，不支持 YAML/TOML**

- 官方文档首句："sing-box uses JSON for configuration files."
- 源码使用 `json.UnmarshalExtendedContext` 解析配置
- 目录扫描时只认 `.json` 后缀文件

配置结构（JSON）：

```json
{
  "log": {},
  "dns": {},
  "ntp": {},
  "certificate": {},
  "endpoints": [],
  "inbounds": [],
  "outbounds": [],
  "route": {},
  "services": [],
  "experimental": {}
}
```

各字段说明：

| Key            | 格式/文档                                        |
|----------------|--------------------------------------------------|
| `log`          | [Log](https://sing-box.sagernet.org/configuration/log/)               |
| `dns`          | [DNS](https://sing-box.sagernet.org/configuration/dns/)               |
| `ntp`          | [NTP](https://sing-box.sagernet.org/configuration/ntp/)               |
| `certificate`  | [Certificate](https://sing-box.sagernet.org/configuration/certificate/) |
| `endpoints`    | [Endpoint](https://sing-box.sagernet.org/configuration/endpoint/)     |
| `inbounds`     | [Inbound](https://sing-box.sagernet.org/configuration/inbound/)       |
| `outbounds`    | [Outbound](https://sing-box.sagernet.org/configuration/outbound/)     |
| `route`        | [Route](https://sing-box.sagernet.org/configuration/route/)           |
| `services`     | [Service](https://sing-box.sagernet.org/configuration/experimental/)  |
| `experimental` | [Experimental](https://sing-box.sagernet.org/configuration/experimental/) |

> 注：网页版文档额外列出 `certificate_providers` 和 `http_clients` 两个字段（v1.13.x 新增），本地 docs/ markdown 源文件尚未更新。

## 2. 启动命令

```bash
sing-box run -c /path/to/config.json
```

### CLI Flag

| Flag | 短写 | 说明 |
|------|------|------|
| `--config` | `-c` | 指定配置文件路径（可多次指定，多文件自动合并） |
| `--config-directory` | `-C` | 指定配置目录（扫描目录下所有 `.json`） |
| `--directory` | `-D` | 设置工作目录 |
| `--disable-color` | | 禁用彩色输出 |

**默认行为**：不指定 `-c` 和 `-C` 时，自动读取当前目录下的 `config.json`。

### 多配置文件

可多次使用 `-c` 指定多个配置文件，sing-box 会通过 `badjson.MergeJSON` 自动合并为一个完整配置：

```bash
sing-box run -c base.json -c inbound.json -c outbound.json
```

也可用 `-C` 指定一个目录，自动扫描目录下所有 `.json` 文件合并。

## 3. 配置检查

```bash
sing-box check
sing-box check -c /path/to/config.json
```

源码逻辑：读取配置 → `box.New()` 创建实例 → 关闭实例。无错误输出即为配置合法。等同于"干跑一次配置解析 + 实例化"，不实际监听端口。

## 4. 格式化配置

```bash
# 输出到 stdout
sing-box format -c config.json

# 直接写回原文件
sing-box format -w -c config.json

# 格式化整个目录
sing-box format -w -D config_directory
```

## 5. 合并配置

```bash
sing-box merge output.json -c config.json -D config_directory
```

将多个配置文件合并输出为一个文件。

## 6. 热重载（SIGHUP）

**文档未提及，但源码支持**

运行中的 sing-box 进程收到 `SIGHUP` 信号后：
1. 调用 `check()` 重新校验配置
2. 校验通过则关闭旧实例，用新配置重新创建
3. 校验失败则打印错误，继续使用旧配置

```bash
kill -HUP $(pidof sing-box)
```

## 7. 其他实用子命令

| 命令 | 说明 |
|------|------|
| `sing-box format` | 格式化配置文件（美化 JSON） |
| `sing-box merge` | 合并多配置文件 |
| `sing-box tools fetch` | 远程获取资源 |
| `sing-box rule-set compile` | 编译规则集 |
| `sing-box rule-set decompile` | 反编译规则集 |
| `sing-box rule-set convert` | 转换规则集格式 |
| `sing-box rule-set match` | 测试规则集匹配 |
| `sing-box geoip list` | 列出 GeoIP 条目 |
| `sing-box geoip lookup` | 查询 GeoIP |
| `sing-box geosite list` | 列出 Geosite 条目 |
| `sing-box geosite lookup` | 查询 Geosite |
| `sing-box generate ech` | 生成 ECH 密钥 |
| `sing-box generate vapid` | 生成 VAPID 密钥 |
| `sing-box version` | 显示版本信息 |

---

## 信息源

### 官方文档

- 配置介绍（含 Check/Format/Merge）: https://sing-box.sagernet.org/configuration/
- Log: https://sing-box.sagernet.org/configuration/log/
- DNS: https://sing-box.sagernet.org/configuration/dns/
- NTP: https://sing-box.sagernet.org/configuration/ntp/
- Certificate: https://sing-box.sagernet.org/configuration/certificate/
- Inbound: https://sing-box.sagernet.org/configuration/inbound/
- Outbound: https://sing-box.sagernet.org/configuration/outbound/
- Route: https://sing-box.sagernet.org/configuration/route/
- Experimental/Services: https://sing-box.sagernet.org/configuration/experimental/

### 本地源码（v1.13.12）

- `cmd/sing-box/cmd.go` — CLI flag 定义（-c, -C, -D）
- `cmd/sing-box/cmd_run.go` — run 子命令、配置读取/合并、SIGHUP 热重载
- `cmd/sing-box/cmd_check.go` — check 子命令
- `cmd/sing-box/cmd_format.go` — format 子命令
- `cmd/sing-box/cmd_merge.go` — merge 子命令
- `cmd/sing-box/main.go` — 入口
- `docs/configuration/index.md` — 配置文档 markdown 源文件
