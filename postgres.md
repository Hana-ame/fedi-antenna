
follow 

https://medium.com/@2018.itsuki/postgresql-with-next-js-and-prisma-44f66a05378a#eb05

```sh
npm install prisma @prisma/client
```

```sh
npx prisma init
```

```md
✔ Your Prisma schema was created at prisma/schema.prisma
  You can now open it in your favorite editor.

warn You already have a .gitignore file. Don't forget to add `.env` in it to not commit any private information.

Next steps:
1. Set the DATABASE_URL in the .env file to point to your existing database. If your database has no tables yet, read https://pris.ly/d/getting-started
2. Set the provider of the datasource block in schema.prisma to match your database: postgresql, mysql, sqlite, sqlserver, mongodb or cockroachdb.
3. Run `prisma db pull` to turn your database schema into a Prisma schema.
4. Run prisma generate to generate the Prisma Client. You can then start querying your database.

More information in our documentation:
https://pris.ly/d/getting-started
```

1. `.env`

```conf
DATABASE_URL="postgresql://postgres:password@localhost:5432/mydb?schema=public"
```

2. 要先写 Model

这个是建好数据库的表格之后自动更新的
use command `npx prisma db pull`

~~`prisma/schema.prisma`~~
```conf
model User {
  id Int        @id
  name String
}
```

3. `npx prisma db pull`

要先新建对应的数据库(这里是`mydb`)

去DBeaver或者console

`npx prisma db push`



4. run `npx prisma generate`


无法使用import

重命名为`seed.mjs`

字段类型为date

`date(value)`

字段类型为uuid

`uuid(value)`

`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

[quickstart](https://www.prisma.io/docs/getting-started/quickstart)

