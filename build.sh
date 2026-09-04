#!/bin/bash

function print_logo() {
    echo "      _____                   _______                   _____                   _______         "
    echo "     /\    \                 /::\    \                 /\    \                 /::\    \        "
    echo "    /::\    \               /::::\    \               /::\    \               /::::\    \       "
    echo "    \:::\    \             /::::::\    \             /::::\    \             /::::::\    \      "
    echo "     \:::\    \           /::::::::\    \           /::::::\    \           /::::::::\    \     "
    echo "      \:::\    \         /:::/~~\:::\    \         /:::/\:::\    \         /:::/~~\:::\    \    "
    echo "       \:::\    \       /:::/    \:::\    \       /:::/  \:::\    \       /:::/    \:::\    \   "
    echo "       /::::\    \     /:::/    / \:::\    \     /:::/    \:::\    \     /:::/    / \:::\    \  "
    echo "      /::::::\    \   /:::/____/   \:::\____\   /:::/    / \:::\    \   /:::/____/   \:::\____\ "
    echo "     /:::/\:::\    \ |:::|    |     |:::|    | /:::/    /   \:::\ ___\ |:::|    |     |:::|    |"
    echo "    /:::/  \:::\____\|:::|____|     |:::|    |/:::/____/  ___\:::|    ||:::|____|     |:::|    |"
    echo "   /:::/    \::/    / \:::\    \   /:::/    / \:::\    \ /\  /:::|____| \:::\    \   /:::/    / "
    echo "  /:::/    / \/____/   \:::\    \ /:::/    /   \:::\    /::\ \::/    /   \:::\    \ /:::/    /  "
    echo " /:::/    /             \:::\    /:::/    /     \:::\   \:::\ \/____/     \:::\    /:::/    /   "
    echo "/:::/    /               \:::\__/:::/    /       \:::\   \:::\____\        \:::\__/:::/    /    "
    echo "\::/    /                 \::::::::/    /         \:::\  /:::/    /         \::::::::/    /     "
    echo " \/____/                   \::::::/    /           \:::\/:::/    /           \::::::/    /      "
    echo "                            \::::/    /             \::::::/    /             \::::/    /       "
    echo "                             \::/____/               \::::/    /               \::/____/        "
    echo "                              ~~                      \::/____/                 ~~              "
    echo "                                                       ~~                                       "
    echo "                                                                                                "
}


function install_go() {
    echo "$INSTALL_GO_MESSAGE"

    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        wget https://golang.org/dl/go1.27.1.linux-amd64.tar.gz
        sudo tar -C /usr/local -xzf go1.27.1.linux-amd64.tar.gz
        echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
        source ~/.bashrc
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        wget https://golang.org/dl/go1.27.1.darwin-amd64.pkg
        sudo installer -pkg go1.27.1.darwin-amd64.pkg -target /
    elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
        curl -LO https://golang.org/dl/go1.27.1.windows-amd64.msi
        msiexec /i go1.27.1.windows-amd64.msi /quiet
        echo "setx PATH \"%PATH%;C:\\Go\\bin\"" >> ~/.bashrc
        source ~/.bashrc
    else
        echo "$UNKNOWN_OS_MESSAGE"
        exit 1
    fi

    echo "$GO_INSTALLED_MESSAGE"
}
print_logo
if ! command -v go &> /dev/null; then
    install_go
fi

lang_choice=1

if [[ "$lang_choice" == "1" ]]; then
    LANGUAGE_DIR="cmd/en"
    INSTALL_GO_MESSAGE="Go is not installed. Installing Go..."
    UNKNOWN_OS_MESSAGE="Unknown operating system. Please install Go manually."
    GO_INSTALLED_MESSAGE="Go has been successfully installed!"
    LANGUAGE_SELECTION_MESSAGE="Select the interface language:"
    LANGUAGE_PROMPT_MESSAGE="Enter the language number (1 or 2): "
    BUILD_MESSAGE="Building the project..."
    MOVE_MESSAGE="Moving the compiled file to $DESTINATION..."
    COMPLETE_MESSAGE="Build completed! You can run toGo from anywhere."
elif [[ "$lang_choice" == "2" ]]; then
    LANGUAGE_DIR="cmd/ru"
    INSTALL_GO_MESSAGE="Go не установлен. Устанавливаем Go..."
    UNKNOWN_OS_MESSAGE="Неизвестная операционная система. Пожалуйста, установите Go вручную."
    GO_INSTALLED_MESSAGE="Go успешно установлен!"
    LANGUAGE_SELECTION_MESSAGE="Выберите язык интерфейса:"
    LANGUAGE_PROMPT_MESSAGE="Введите номер языка (1 или 2): "
    BUILD_MESSAGE="Сборка проекта..."
    MOVE_MESSAGE="Перемещение скомпилированного файла в $DESTINATION..."
    COMPLETE_MESSAGE="Сборка завершена! Вы можете запустить toGo из любого места."
else
    echo "$INVALID_SELECTION_MESSAGE"
    exit 1
fi


echo "$BUILD_MESSAGE"
go build -o togo main.go

if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
    DESTINATION="/usr/local/bin/togo"
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    DESTINATION="$HOME/go/bin/togo"
else
    echo "$UNKNOWN_OS_MESSAGE"
    exit 1
fi

echo "$MOVE_MESSAGE"
if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
    sudo mv togo "$DESTINATION"
    sudo chmod +x "$DESTINATION"
    mkdir -p ~/.local/state/toGo
else
    mv toGo "$DESTINATION"
fi

echo "$COMPLETE_MESSAGE"
