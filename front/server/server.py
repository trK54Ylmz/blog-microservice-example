from flask import Flask

app = Flask('stater', template_folder='template', static_folder='./asset')

# update flask configuration
app.config['JSON_AS_ASCII'] = False
