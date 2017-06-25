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

TODO
====


- 指定重新下载 某话的某页

- 定时任务通知新抓去的漫画

- config 的json 落地一份？

API
===


::

    curl "http://127.0.0.1:8080/v1/book" -X POST -d '{"name":"book2"}'
    curl "http://127.0.0.1:8080/v1/book/1"
