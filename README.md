# Batmanlang

Batmanlang is a perfect stack based programming language (no heap == no segfault). These are its features!<br>

How to use
- git clone https://github.com/BatmanBoxer/BatmanLang
- cd interpreter-in-go
- go mod tidy
- go run cmd/main/main.go [filename].batman
  
## Features!
Unlike every other language which lack originality (i dont) the start of this programming langauge is from fun batman

```java
fun batman(){
  print("Hello Gotham");
}
```

1) No need for any syscalls for printing . you can do it with a simple print statement. this is revolution art

```java
print("Hello world")
```

2) No need for multiple ways of declaring variable because its just confusing. just use var.

```java
var i = 1;
print(i)
```

3) No comments Beause if you need comment to tell you what code does just get better
4) You can create another function and call it as well (it has scoping too cause its just basic stuff(gives a cool error when you break scoping rules) )

```java
fun test(){
    print("hello");
}

fun batman(){
    var p = 1;
    print(p);

    var a = 10;
    var b = 10;

    if(a == b){
        print("batman from if");
    }

    while(a == 19){
        var darwin = 1;
        print(darwin);
        print("this is the value of a");
        print(a);
        a = a + 1;
    }

    test();
}

```
5) Conditional Statements: Batmanlang supports simple if conditions and no else  only weak programmer use else
   
  ```java
if(condition){
      print("batman from if");
}
```

6) Since i am original my while loop will continue to loop while it is false and stop when it is true (genius)
  ```java
fun batman(){
var i = 1;
var j = 2;
while(i==j){
      print("batman from if");
  }
}
```
that above is a infinite loop

