# Go言語で作るインタプリタ(Writing An Interpreter In Go)
![Go言語で作るインタプリタ](./images/picture_large978-4-87311-822-2.jpeg)

## セットアップ
```
git clone https://github.com/Acasune/monkey-lang.git
cd monkey-lang
go run main.go
```

## 言語仕様
以下の機能を実装
- 変数束縛
- 整数と真偽値
- 算術式
- 組み込み関数
- 第一級の高階関数
- クロージャー
- 文字列データ型
- 配列データ型
- ハッシュデータ型
- マクロ

## 実行テスト

```
Hello acasune! This is the Monkey programming language!
Feel free to type in commands
>> 1 + 1
2
>> let age = 1;
>> let name = "Monkey";
>> let result = 10 * (20 / 2);
>> let myArray=[1 ,2 ,3 ,4 ,5];
>> let acasune = {"name": "Acasune", "age":"18"};
>> age
1
>> name
Monkey
>> result
100
>> myArray[0]
1
>> acasune["name"]
Acasune
>> let add = fn(a,b){ a + b; };
>> add(1, 2);
3
>> let fibonacci = fn(x){ if (x == 0){ 0 } else { if(x== 1){ 1 } else {fibonacci(x - 1) + fibonacci(x - 2);}}}
>> fibonacci(7)
13
>> let twice = fn(f, x){ return f(f(x));};
>> let addTwo = fn(x) { return x + 2;};
>> twice(addTwo, 2)
6
>> let unless = macro(condition, consequence, alternative) { quote(if (!(unquote(condition))) { unquote(consequence); } else { unquote(alternative); }); };
>> unless(10 > 5, puts("not greater"), puts("greater"))
greater
null
>> 
```