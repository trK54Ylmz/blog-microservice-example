from endpoint import article, user
from endpoint.handle import handle_grpc_error
from server import app

app.register_error_handler(500, handle_grpc_error)

app.register_blueprint(article, url_prefix='/article')
app.register_blueprint(user, url_prefix='/user')

app.run(host='0.0.0.0', port=8003)
