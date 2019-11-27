from flask_restplus import Resource, Namespace, fields, reqparse
from resources.predict_risk import _code_parser
from model.files import ASSOCIATE_PATH, FAMILY_PATH, CROSS_PATH
from werkzeug.exceptions import NotFound
import tushare as ts
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


_line_parser = reqparse.RequestParser() \
    .add_argument('code', type=str, help='股票代码', required=True) \
    .add_argument('start_date', type=str, help='开始日期', required=True) \
    .add_argument('end_date', type=str, help='结束日期', required=True)


@api.route('/klines')
class Klines(Resource):

    @api.expect(_line_parser)
    def get(self):
        args = _line_parser.parse_args()
        code, s_date, e_date = args['code'], args['start_date'], args['end_date']
        pro = ts.pro_api('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')
        df = pro.query('daily', ts_code=code, start_date=s_date, end_date=e_date)
        res = []
        for index, row in df.iterrows():
            tmp = {
                'date': row['trade_date'],
                'open': row['open'],
                'close': row['close'],
                'highest': row['high'],
                'lowest': row['low'],
                'vol': row['vol']
            }
            res.append(tmp)
        return res



