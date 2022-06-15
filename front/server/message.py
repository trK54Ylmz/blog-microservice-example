import json
from flask import Response
from util.json import CustomEncoder


class Message:
    def __init__(self, status=False, message=None, **kwargs):
        """
        Create empty base message
        """
        self.status = status
        self.message = message

        for k, v in kwargs.items():
            self.__dict__[k] = v

    def __setattr__(self, key, value):
        """
        Add or update key

        :param str key: object key name
        :param str or float or int value: object value
        """
        self.__dict__[key] = value

    def json(self, status=200) -> Response:
        """
        Encode response as json object

        :return: json response object
        :rtype: Response
        """
        o = self.__dict__
        if self.message is None:
            o.pop('message')

        # encode json object
        return Response(
            status=status,
            response=json.dumps(o, cls=CustomEncoder, ensure_ascii=False),
            mimetype='application/json; charset=utf-8',
        )
