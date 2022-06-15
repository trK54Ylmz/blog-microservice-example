from datetime import datetime
from json import JSONEncoder


class CustomEncoder(JSONEncoder):
    def default(self, o):
        """
        Serialize and de-serialize by using custom parser

        :param any o: custom object
        """
        if isinstance(o, datetime):
            return o.strftime('%Y-%m-%d %H:%M:%S%z')

        return super(CustomEncoder, self).default(o)
