### mielofon

[![Build Status](https://travis-ci.org/bbrodriges/mielofon.svg?branch=master)](https://travis-ci.org/bbrodriges/mielofon)
[![GoDoc](https://godoc.org/github.com/bbrodriges/mielofon?status.svg)](https://godoc.org/github.com/bbrodriges/mielofon)
[![Go Report Card](https://goreportcard.com/badge/github.com/bbrodriges/mielofon)](https://goreportcard.com/report/github.com/bbrodriges/mielofon)
[![Coverage Status](https://coveralls.io/repos/github/bbrodriges/mielofon/badge.svg?branch=master)](https://coveralls.io/github/bbrodriges/mielofon?branch=master)

Простая библиотека для создания навыков для [Яндекс.Алисы](https://alice.yandex.ru).

### Содержимое

#### dialog

Пакет `dialog` содержит базовые структуры и методы для работы со входящим запросом и исходящим ответом. См. `dialog/examples` с примерами использования.

#### dialog/dialogutil

Пакет `dialog/dialogutil` содержит вспомогательные функции и полезные обертки над объектами и методами из пакета `dialog`.

#### session

Пакет `session` содержит интерфейс хранилища пользовательских сессий и 3 базовые имплементации:

- `NopStore` - несуществующее хранилище
- `MockStore` - хранилище для использования в тестах
- `MemoryStore` - in-memory хранилище на основе `sync.Map`

#### audio

Пакет `audio` содержит пакеты для использования звуков и эффектов в Text-To-Speech: `effects` и `sounds`. Пакет `effects` содержит эффекты модуляции голоса Алисы. Пакет `sounds` содержит разлчные звуки, которые можно использовать совместно со стандартной речью Алисы.

### Лицензия

MIT