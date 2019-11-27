from flask_restplus import Resource, Namespace, reqparse, fields

from model.average_model import predict, predict_given_forecast_close_line

api = Namespace('PredictRisk', description='预测爆仓风险')

_code_parser = reqparse.RequestParser()
_code_parser.add_argument('code', type=str, help='股票代码', required=True)

_code_forecast_parser = reqparse.RequestParser()
_code_forecast_parser.add_argument('code', type=str, help='股票代码', required=True)
_code_forecast_parser.add_argument('forecast_close_line', type=float, help='预期平仓线', required=True)

risk_model = api.model(
    'risk_model',
    {
        'risk': fields.Float,
        'rank': fields.Integer,
        'total': fields.Integer
    }
)

only_risk_model = api.model(
    'only_risk_model',
    {
        'risk': fields.Float
    }
)


# @cross_origin()
@api.route('/predict')
class PredictRisk(Resource):

    @api.expect(_code_parser)
    @api.marshal_with(risk_model)
    def post(self):
        """
        通过股票代码和预期平仓线获得爆仓概率
        :return:
        """
        args = _code_parser.parse_args()
        risk_prob, rank, total = predict(**args)
        return {
            'risk': risk_prob,
            'rank': rank,
            'total': total
        }


@api.route('/predict_forecast')
class PredictRisk(Resource):

    @api.expect(_code_forecast_parser)
    @api.marshal_with(only_risk_model)
    def post(self):
        """
        通过股票代码和预期平仓线获得爆仓概率
        :return:
        """
        args = _code_forecast_parser.parse_args()
        risk_prob = predict_given_forecast_close_line(**args)
        return {
            'risk': risk_prob,
        }
