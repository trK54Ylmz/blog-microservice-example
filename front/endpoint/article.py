from flask import Blueprint, request
from rpc import ArticleService
from server import app
from server.message import Message

article = Blueprint('article', app.name)


@article.get('/create')
def create():
    """
    Create new article

    :rtype: flask.Response
    """
    m = Message()

    id = request.args.get('id')
    title = request.args.get('title')
    description = request.args.get('description')
    user_id = int(request.args.get('user_id'))

    ars = ArticleService()

    # create new article
    res = ars.create(id, title, description, user_id)

    m.status = res.status

    return m.json()


@article.get('/list')
def list():
    """
    Get all articles

    :rtype: flask.Response
    """
    m = Message()

    ars = ArticleService()

    # create new article
    res = ars.list()

    m.status = res.status

    if res.status:
        m.articles = res.articles

    return m.json()
