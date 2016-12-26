Welcome to use thrift client pool.

Why wrote thrift client pool, cause of thrift does not provide client pool for golang version.

You should provide {...}

A function to dail your server.
A function to keep alive with your server, keep alive is simple but i think you'll need customize your keep alive logic.
Optional function to close the connection with your server.
You could configure {...}

Max pool size.
Initial connected pool size.
Intervals: - Keep alive interval, - Reconnect interval. - Create new interval.
Support or Contact

Having trouble with Pages? Check out our documentation or contact support and we’ll help you sort it out.

You can send message to me on github, and my mail: wangxingge83@gmail.com
