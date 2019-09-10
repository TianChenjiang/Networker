from flask import Flask
from flask_restplus import Api
from werkzeug.middleware.proxy_fix import ProxyFix
from config import config
from os import getenv
import tensorflow as tf
from flask_cors import CORS


from resources.predict_risk import api as ns_predict


APP_ENV = getenv('APP_ENV', 'dev')
app = Flask(__name__)
CORS(app)
app.config.from_object(config[APP_ENV])
app.wsgi_app = ProxyFix(app.wsgi_app)

tf.compat.v1.logging.set_verbosity(tf.compat.v1.logging.WARN)

api = Api(app, version='1.0', title='Predict Risk', prefix='/api')


api.add_namespace(ns_predict)


if __name__ == '__main__':
    app.run(port=5000, debug=True)