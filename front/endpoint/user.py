from flask import Blueprint, request
from rpc import UserService
from server import app
from server.message import Message

user = Blueprint('user', app.name)


@user.get('/sign-in')
def sign_in():
    """
    Sign in to user account

    :rtype: flask.Response
    """
    m = Message()

    username = request.args.get('user')
    password = request.args.get('pass')

    us = UserService()

    res = us.sign_in(username, password)

    m.status = res

    return m.json()
