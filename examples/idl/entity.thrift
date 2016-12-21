namespace go wangxingge.thrift_clientpool.examples.entity
namespace csharp wangxingge.thrift_clientpool.examples.entity

const string ServiceTag_BookService = "BookService"
const string ServiceTag_UserService = "UserService"
const string ServiceTag_KeepAliveService = "KeepAliveService"

struct Book
{
    1:required  string        BookId,
    2:required  string        BookName,
    3:required  string        Author,
    4:optional  string        Price,
    5:optional  string        Date,
    6:optional  binary        Cover
}

struct User
{
    1:required  string        UserId,
    2:required  string        UserName,
    3:optional  i64           BookNum,
    4:optional  set<string>   BookList,
    5:optional  binary        Avatar
}