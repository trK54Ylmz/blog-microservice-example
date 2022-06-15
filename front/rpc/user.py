from .base import RpcBase
from proto import UserServiceStub, SignInRequest


class UserService(RpcBase):
    def sign_in(self, username, password) -> bool:
        """
        Sign in to user account by using username and password

        :param str username: user name
        :param str password: password
        :return: sign in status
        :rtype: bool
        """
        r = SignInRequest(username=username, password=password)

        stub = UserServiceStub(self.channel)

        # send request to grpc server
        res = stub.SignIn(r)

        return res.status
