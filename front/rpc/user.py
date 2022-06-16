import os
import grpc
from proto import UserServiceStub, SignInRequest, SignInResponse


class UserService:
    def __init__(self):
        host = os.environ.get('USER_HOST', 'localhost:8001')
        channel = grpc.insecure_channel(host)

        self.stub = UserServiceStub(channel)

    def sign_in(self, username, password) -> SignInResponse:
        """
        Sign in to user account by using username and password

        :param str username: user name
        :param str password: password
        :return: sign in status
        :rtype: bool
        """
        r = SignInRequest(username=username, password=password)

        # send request to grpc server
        return self.stub.SignIn(r)
