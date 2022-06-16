import os
import grpc
from google.protobuf.empty_pb2 import Empty
from proto import Article, ArticleCreateResponse, ArticleListResponse, ArticleServiceStub


class ArticleService:
    def __init__(self):
        host = os.environ.get('ARTICLE_HOST', 'localhost:8002')
        channel = grpc.insecure_channel(host)

        self.stub = ArticleServiceStub(channel)

    def create(self, id, title, description, user_id) -> ArticleCreateResponse:
        """
        Create new article

        :param str id: article ID
        :param str title: article title
        :param str description: article description
        :param int user_id: user ID
        :return: article create response
        :rtype: ArticleCreateResponse
        """
        a = Article()
        a.id = id
        a.user_id = user_id
        a.title = title
        a.description = description

        # send request to grpc server
        return self.stub.Create(a)

    def list(self) -> ArticleListResponse:
        """
        Get all articles

        :return: article list response
        :rtype: ArticleListResponse
        """
        e = Empty()

        # send request to grpc server
        return self.stub.List(e)
