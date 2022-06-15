from endpoint import article, user
from server import app

app.register_blueprint(article, url_prefix='/article')
app.register_blueprint(user, url_prefix='/user')

app.run(port=8003)
