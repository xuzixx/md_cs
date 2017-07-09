#####
md cs
#####

.. contents:: 目录

毛邓Comic Spider

通过配置抓取新的漫画

DB
==

::

    sqlite3 data.db

    sqlite> .databases
    
    sqlite> .tables

    # cli 设置
    sqlite>.header on
    sqlite>.mode column
    sqlite>.timer on

建表语句

::

    create table test (id INTEGER PRIMARY KEY AUTOINCREMENT, data char(50))

Swagger
=======

::

    bee run -gendoc=true -downdoc=true

http://127.0.0.1:8080/swagger/

**refer**

::

    https://beego.me/docs/advantage/docs.md
    https://my.oschina.net/astaxie/blog/284072


Version
=======

+---------------+--------+
| beego         | v1.8.4 |
+---------------+--------+
| astaxie/beego | v1.8.3 |
+---------------+--------+

TODO
====

- 指定重新下载 某话的某页

- config 的json 落地一份？

- web 分页
