from flask_restplus import Resource, Namespace, reqparse, fields
from werkzeug.exceptions import NotFound

from model.average_model import predict

api = Namespace('PredictRisk', description='预测爆仓风险')

_predict_risk_parser = reqparse.RequestParser()
_predict_risk_parser.add_argument('code', type=str, help='股票代码', required=True)

risk_model = api.model(
    'risk_model',
    {
        'risk': fields.Float,
        'rank': fields.Integer,
        'total': fields.Integer
    }
)


# @cross_origin()
@api.route('/predict')
class PredictRisk(Resource):

    @api.expect(_predict_risk_parser)
    @api.marshal_with(risk_model)
    def post(self):
        """
        通过股票代码和预期平仓线获得爆仓概率
        :return:
        """
        args = _predict_risk_parser.parse_args()
        risk_prob, rank, total = predict(**args)
        return {
            'risk': risk_prob,
            'rank': rank,
            'total': total
        }
