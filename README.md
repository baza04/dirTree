Утилита tree.

Выводит дерево каталогов и файлов (если указана опция -f).
Первым аргументом принимает путь, вторым флаг -f
В данном проекте использовались стандартные пакеты golang:
	```
    "fmt"
	"io"
	"io/ioutil"
	"os"
```
А также структуры и методы.

Запускать тесты через `go test -v` находясь в папке c заданием (предварительно удалив или перенеся папку .git).
После запуска вы должны увидеть такой результат:

```
$ go test -v
=== RUN   TestTreeFull
--- PASS: TestTreeFull (0.00s)
=== RUN   TestTreeDir
--- PASS: TestTreeDir (0.00s)
PASS
ok      coursera/homework/tree     0.127s
```

```
go run main.go . -f
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)
go run main.go .
└───testdata
	├───project
	├───static
	│	├───css
	│	├───html
	│	└───js
	└───zline
```