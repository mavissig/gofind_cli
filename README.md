# gofind_cli

Разработка сервиса работы с json файлами

Author : [Egor Kondratov(mavissig)](https://github.com/mavissig)

---

## Contents

1. [Preamble](#preamble)
2. [Introduction](#introduction)
3. [Information](#information)
4. [About the program](#about-the-program)

## Preamble

---

<img src="misc/docs/readme/preamble.jpg" width="100%">

В рамках учебного проекта написал утилиту, которая работает подобно утилите Find в Linux/Unix системах, но 
расширена  дополнительным функционалом

## introduction

---

Программа может работать в следующих режимах:

- Find - обход директорий и вывод их содержимого в stdout 


## Information

Далее приведена подробная информация по режимаи программы

---

### Find

```bash
gofind [-f [-ext 'ext'] | -d | -sl] path
```

Вызов программы без флагов запустит программу активировав по умолчанию следующие флаги и опции: `-f` `-d` `-sl`

Описание флагов и опций:

- `-f` рекурсивно обходит вложенные директории и выводит в stdout названия файлов

    ```bash
    gofind -f ~/go/testdir/testdata
  
    === output ===
    
    /Users/admin/go/testdir/testdata/src/src.go
    /Users/admin/go/testdir/testdata/text.txt
    ```

- `-ext` рекурсивно обходит вложенные директории и выводит в stdout названия файлов с расширением указанным после 
флага `ext`, например `-ext 'go'` выведет только файлы с расширением _.go_

    > флаг -ext может использоваться только после флага -f

    ```bash
    gofind -f -ext 'go' ~/go/testdir/testdata
    
    === output ===
    
    /Users/admin/go/testdir/testdata/src/src.go
    ```

- `-d` рекурсивно обходит вложенные директории и выводит в stdout названия директорий
    
    ```bash
    gofind -d ~/go/testdir/testdata
    
    === output ===
    
    /Users/admin/go/testdir/testdata
    /Users/admin/go/testdir/testdata/src
    ```

- `-sl` рекурсивно обходит вложенные директории и выводит в stdout символические ссылки и путь к файлу, на который они указывают

    ```
    gofind -sl ~/go/testdir/testdata
    
    === output ===
    
    /Users/admin/go/testdir/testdata/sl_Pictures -> /Users/admin/Pictures/
    ```


## About the program

--- 

- Программа разработана на языке _Go(Golang) 1.21.1_
- Программы написана согласно принципам "Чистой архитектуры" и не нарушает принципы SOLID