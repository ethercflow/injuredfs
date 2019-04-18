#!/usr/bin/env python

import grpc
import injure_pb2
import injure_pb2_grpc
from datetime import datetime

if __name__ == "__main__":
    methods = "write"
    errno = 0
    random = False
    pct = 100
    path = "testdir/*"
    delay = 5000000

    print "fault..."

    with grpc.insecure_channel('127.0.0.1:65534') as channel:
        stub = injure_pb2_grpc.InjureStub(channel)
        stub.SetFault(injure_pb2.Request(methods=methods.split(
            ","), errno=errno, random=random, pct=pct, path=path, delay=delay))

    b = datetime.now()
    with open('/mnt/edata/testdir/output.txt', 'w') as f:
        f.write('Hi there!')
    e = datetime.now()

    print "write cost %d seconds, do recover..." % (e - b).seconds

    with grpc.insecure_channel('127.0.0.1:65534') as channel:
        stub = injure_pb2_grpc.InjureStub(channel)
        stub.RecoverMethod(injure_pb2.Request(methods=methods.split(",")))

    b = datetime.now()
    with open('/mnt/edata/testdir/output.txt', 'w') as f:
        f.write('Hi there!')
    e = datetime.now()
    print "write cost %d seconds after recover..." % (e - b).seconds

    errno = 0x1c  # NOSPC
    with grpc.insecure_channel('127.0.0.1:65534') as channel:
        stub = injure_pb2_grpc.InjureStub(channel)
        stub.SetFault(injure_pb2.Request(methods=methods.split(
            ","), errno=errno, random=random, pct=pct, path=path, delay=0))

    try:
        with open('/mnt/edata/testdir/output.txt', 'w') as f:
            f.write('Hi there!')
    except IOError as e:
        print "I/O error({0}): {1}".format(e.errno, e.strerror)
    finally:
        with grpc.insecure_channel('127.0.0.1:65534') as channel:
            stub = injure_pb2_grpc.InjureStub(channel)
            stub.RecoverMethod(injure_pb2.Request(methods=methods.split(",")))
