from grpc import RpcError
from server import app
from server.message import Message


@app.errorhandler(RpcError)
def handle_grpc_error(e):
    """
    Handle grpc exception

    :param RPCError e: rpc exception
    :rtype: flask.Response
    """
    m = Message()
    m.message = e.details()

    return m.json(), 500
