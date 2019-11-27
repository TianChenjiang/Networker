from flask_restplus import Resource, Namespace, fields
from resources.predict_risk import _code_parser
from model.files import ASSOCIATE_PATH, FAMILY_PATH, CROSS_PATH
from werkzeug.exceptions import NotFound
import json

api = Namespace('Graph', description='网络图')


def get_json_file(filepath):
    if not filepath.is_file():
        raise NotFound('no such company!')

    with open(filepath, 'rb') as f:
        result = json.load(f)

    return result


@api.route('/family')
class Family(Resource):

    @api.expect(_code_parser)
    def post(self):
        args = _code_parser.parse_args()
        code = args['code']
        file = FAMILY_PATH.joinpath(code + '.json')

        return get_json_file(file)


@api.route('/cross')
class Cross(Resource):

    @api.expect(_code_parser)
    def post(self):
        args = _code_parser.parse_args()
        code = args['code']
        file = CROSS_PATH.joinpath(code + '.json')

        return get_json_file(file)


@api.route('/associate')
class Associate(Resource):

    @api.expect(_code_parser)
    def post(self):
        args = _code_parser.parse_args()
        code = args['code']
        file = ASSOCIATE_PATH.joinpath(code + '.json')

        return get_json_file(file)


