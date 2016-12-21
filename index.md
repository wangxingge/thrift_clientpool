## Welcome to use thrift client pool.

Why wrote thrift client pool, cause of thrift does not provide client pool for golang version.

# You should provide {...}
1. A function to dail your server.
2. A function to keep alive with your server, keep alive is simple but i think you'll need customize your keep alive logic.
3. Optional function to close the connection with your server.

# You could configure {...}
1. Max pool size.
2. Initial connected pool size.
3. Intervals:
        - Keep alive interval,
        - Reconnect interval.
        - Create new interval.

### Support or Contact
You can send message to me on github, and my mail: wangxingge83@gmail.com
