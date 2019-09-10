from flask_restplus import Resource, Namespace, reqparse, fields
from werkzeug.exceptions import NotFound
from flask_cors import cross_origin

from model.average_model import predict

api = Namespace('PredictRisk', description='预测爆仓风险')

_predict_risk_parser = reqparse.RequestParser()
_predict_risk_parser.add_argument('code', type=str, help='股票代码', required=True)
_predict_risk_parser.add_argument('forecast_close_line', type=float, help='预期平仓线', required=True)

risk_model = api.model(
    'risk_model',
    {
        'risk': fields.Float
    }
)


@cross_origin()
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
        risk_prob = predict(**args)
        if risk_prob is None:
            raise NotFound('code {} not found'.format(args['code']))
        return {
            'risk': risk_prob
        }
