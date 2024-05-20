# Ethernodes Fetcher

这是一个用Go编写的程序，用于从Ethernodes网站获取以太坊节点信息并输出指定格式的节点字符串。

## 功能

- 根据提供的参数，从Ethernodes网站获取节点数据。
- 支持输出 `enode` 格式和 `admin.addPeer` 格式的节点字符串。

## 安装

1. 确保你已经安装了Go环境。
2. 下载并克隆此项目到你的本地机器。

    ```sh
    git clone https://github.com/yourusername/ethernodes-fetcher.git
    cd ethernodes-fetcher
    ```

3. 编译程序。

    ```sh
    go build -o ethernodes-fetcher
    ```

## 使用方法

程序支持以下命令行参数：

- `-net`：选择网络类型，可以是 `mainnet` 或 `testnet`，默认为 `mainnet`。
- `-start`：起始索引，默认为 `0`。
- `-length`：获取的节点数量，默认为 `10`。
- `-fmt`：输出格式，可以是 `enode` 或 `addpeer`，默认为 `addpeer`。

### 示例

获取主网的前10个节点并输出 `admin.addPeer` 格式的节点字符串：

```sh
./ethernodes-fetcher -net mainnet -start 0 -length 10 -fmt addpeer
```

