# sing-shadowsocks

Форк [metacubex/sing-shadowsocks](https://github.com/metacubex/sing-shadowsocks) для [miray](https://github.com/rityak/miray). Наш код — в ветке `miray`.

## Что изменено

- Shadowsocks 2022, UDP: поля заголовка пакета читаются и пишутся напрямую, без `binary.Read/Write` (рефлексия и аллокация на каждое поле каждой датаграммы). Формат на проводе тот же.
