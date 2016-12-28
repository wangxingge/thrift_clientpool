namespace go wangxingge.thrift_clientpool.examples.bookservice
namespace csharp wangxingge.thrift_clientpool.examples.bookservice

service CommonService
{
    bool                DefaultKeepAlive(1:string clientId)
}