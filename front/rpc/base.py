import grpc


class RpcBase:
    def __init__(self):
        self.channel = grpc.insecure_channel('localhost:8001')
