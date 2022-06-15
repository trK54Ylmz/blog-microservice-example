from endpoint import user
from server import app

app.register_blueprint(user, url_prefix='/user')

app.run(port=8003)
