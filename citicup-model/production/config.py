from os import getenv


class BaseConfig:
    RESTPLUS_MASK_SWAGGER = False
    ERROR_404_HELP = False


class DevelopmentConfig(BaseConfig):
    DEBUG = True


class TestConfig(BaseConfig):
    TESTING = True


class ProductConfig(BaseConfig):
    TF_CPP_MIN_LOG_LEVEL = '3'
    pass


config = {
    'dev': DevelopmentConfig,
    'test': TestConfig,
    'prod': ProductConfig
}