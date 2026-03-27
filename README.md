### Template Go Project Structure

For Jetbrains IDE: Goland

##### Добавление External Tool

Settings → Tools → External Tools → +

Заполняешь:

- Name: `Init Go Structure`
- Program: `/bin/bash` <--- путь к интепретатору bash
- Arguments: `/полный/путь/init-structure.sh`
- Working directory: `$ProjectFileDir$`

##### Использование

- Создаёшь новый проект в GoLand
- IDE создаёт go.mod
- ПКМ по проекту →
External Tools → Init Go Structure

Готово!