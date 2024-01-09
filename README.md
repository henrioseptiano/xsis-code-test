# xsis-code-test
this repository contains source code of RESTful API XSIS Test Assignment

## Explaination
Using Golang with RESTful API Go-Chi routing framework. it has contains unit test
and Dependency Injection layer.

### Tech Stacks
  <ul>
       <li>Golang Version 1.21</li>
       <li>GORM</li>
       <li>PostgreSQL</li>
       <li>Docker</li>
       <li>Make</li>
  </ul>

### How to run the program?

Simply run 

```
    make build
```

NOTE : make sure to create your own .env file. otherwise it would not work!

### How to run the unit test?

NOTE : Please install make first in order to run makefile command

Simply run 

```
    make test_cover
```

or if you want to look at html view 

```
    make test_cover_html
``` 

### how to turn off the application?

simply run 

```
    make down
```