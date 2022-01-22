# Streaming attacker

Streaming attacker is a stress testing tool for streaming server

via multiple goroutine process to simulate multi-sources pushing streaming by protocol rtmp

## Prerequisite

- Make sure the environment with `ffmpeg` cause tool execute `ffmpeg` directly

- Prepare test streaming data ready

## Usage

### Build

Build binary according to your os

Or get some help

```sh
make help
```

Example of building mac binary

```sh
make build-mac
```

### execution

```sh
./attacker -n=20 -s=stream-testing.kfs.io -t=A -f=./audio/example.mp4
```

Above command would spawn 20 goroutines, and start rtmp streaming respectively.

And naming each streaming with A1, A2, ... A20 order by order.

## Arguments

- c: `ffmpeg` path
- n: Client numbers in simulation
- t: Tag for naming，default with A
- s: Target URL
- p: Port
- a: Streaming application
- f: File path for streaming test

----------------------------------------

# Streaming attacker

這是一個用在 streaming server 上的壓力測試工具

此工具可以用來針對特定 streaming server 進行 rtmp push 端的壓力測試

透過程式模擬多個 client 進行 rtmp 的推流

## 環境確認

- 工具透過 go 執行 local 的 `ffmpeg` 指令工具

  因此必須確認執行的環境有 ffmpeg 指令安裝

- 確認測試環境要串流用的音檔

## 使用方式

### Build

根據 makefile 選擇要 compile 的平台，執行檔產生在 `bin` 資料夾中

查看 makefile 有哪些目標可以使用

```sh
make help
```

build mac 版本的 binary

```sh
make build-mac
```

### 執行

```sh
./attacker -n=20 -s=stream-testing.kfs.io -t=A -f=./audio/example.mp4
```

上述指令將依序開啟 20 個子任務，透過 ffmpeg 執行 rtmp 串流

並將 stream key 依序自動命名為 A1, A2, ... A20

## 指令參數

- c: ffmpeg 路徑
- n: 模擬的 client 數量
- t: 任務標籤，會標示在每個 stream 前面
- s: 目標 URL
- p: Port
- a: Streaming app
- f: 欲串流的檔案路徑
