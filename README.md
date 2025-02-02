# get_switchbot_device

このプロジェクトは、Switchbot APIを使用してSwitchbotデバイスと赤外線リモコンで操作できるデバイスの一覧を表示します。

## 前提条件

- Go 1.16以降
- Switchbot APIトークン

## インストール

1. リポジトリをクローンします:
    ```sh
    git clone https://github.com/yourusername/get_switchbot_device.git
    cd get_switchbot_device
    ```

2. 依存関係をインストールします:
    ```sh
    make deps
    ```

3. プロジェクトをビルドします:
    ```sh
    make build
    ```

4. オプションで、バイナリをローカルのbinにインストールします:
    ```sh
    make install
    ```

## 使用方法

1. Switchbot APIトークンの環境変数を設定します:
    ```sh
    export OPEN_TOKEN=your_open_token
    export SECRET_TOKEN=your_secret_token
    ```

2. バイナリを実行します:
    ```sh
    ./get_switchbot_devices
    ```

## Makefileターゲット

- `build`: `CGO_ENABLED=0`および`-trimpath`でプロジェクトをビルドします。
- `clean`: ビルド成果物をクリーンします。
- `test`: テストを実行します。
- `deps`: 依存関係をインストールします。
- `build-linux`: Linux用にプロジェクトをクロスコンパイルします。
- `install`: バイナリを`$(HOME)/.local/bin/`にコピーします。

origin: https://gist.github.com/nasa9084/2be5e22485ff2ad91fdb3e5a0f2a01b4
