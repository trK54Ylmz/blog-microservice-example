import grpc
from proto import UserServiceStub, SignInRequest, SignInResponse


class UserService:
    def __init__(self):
        channel = grpc.insecure_channel('localhost:8001')

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
