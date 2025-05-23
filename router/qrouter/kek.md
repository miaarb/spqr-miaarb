## поддержать парсинг create func

```
create function f1 (i int) returns void as $$ insert into tt values(i); insert into tt values(i + 1) $$ language sql;
```

https://github.com/postgres/postgres/blob/master/src/backend/parser/gram.y#L8356
n->funcname = $4;
n->sql_body = $9;

такое же реализовать здесь
https://github.com/pg-sharding/lyx/blob/master/lyx/gram.y#L4040


## начать создавать функции на шардах
routeWithRules
	case *lyx.CreateTable:
    // todo: case *lyx.CreateFunction


## начать анализировать функции в рантайме на роутере (planByWhereClause)



https://github.com/pg-sharding/lyx

https://github.com/postgres/postgres/blob/master/src/backend/parser/gram.y#L15667
https://github.com/pg-sharding/lyx/blob/master/lyx/gram.y#L2193

`select f1(10);`

`select f1(10) from tt where i = 32 (planByWhereClause) order by `

```

switch stmt.TargetList[0].(type) {
case *lyx.FuncApplication:
    // lkwenciwejc;qs
}


&lyx.Select{
    TargetList: []lyx.Node{
        &lyx.FuncApplication {
            Name: "f1",
            Args: []lyx.Node {
                &lyx.AExprIConst {
                    Value: 10
                }
            }
        }
    }
}

```