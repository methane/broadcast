#!/usr/bin/env python
from __future__ import print_function
from tornado.options import define, options, parse_command_line
from tornado.ioloop import IOLoop
from tornado.tcpserver import TCPServer


define("port", default=5000, type=int)


class BroadCaster(TCPServer):
    connections = set()


    def handle_stream(self, stream, address):
        stream.set_nodelay(True)

        def on_close(data):
            self.connections.remove(stream)
            stream.close()

        def on_read(data):
            # broadcast
            for conn in self.connections:
                conn.write(data)

        stream.read_until_close(callback=on_close, streaming_callback=on_read)
        self.connections.add(stream)


def main():
    parse_command_line()
    server = BroadCaster()
    print("Start listening on", options.port)
    server.listen(options.port)
    IOLoop.instance().start()


if __name__ == "__main__":
    main()

