# Biz

Biz层仅处理数据输入和输出，不包含具体业务逻辑代码。

Repository声明数据库访问接口，接口的方法在data层实现(具体的数据库CRUD操作逻辑)

UseCase是利用Repository生成的用例类型，需要实现可供service调用的多种业务逻辑方法，方法内部调用适当的Repository方法操作数据库（也可能是和其他service交互）

举例：

Repository接口的方法包含Create、R、U、D等简单的数据库操作方法

service想调用server.useCase.add()方法，实现插入数据到数据库

这就要useCase实现add()方法，内部调用useCase.Repo.Create()方法。

service业务复杂时，各种各样名称的操作，都可以用Repository中简单的数据库操作方法组合而成，service中只需要调用对应的某个方法即可。这样可以避免在service中重复编写相同的复杂查询过程。
