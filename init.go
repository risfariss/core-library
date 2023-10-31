package core_library

import (
	"bitbucket.org/kawancicil/core-library/common"
	"bitbucket.org/kawancicil/core-library/external/oss"
	"bitbucket.org/kawancicil/core-library/external/privy"
	"bitbucket.org/kawancicil/core-library/external/rabbitMQ"
	"bitbucket.org/kawancicil/core-library/external/slack"
	"bitbucket.org/kawancicil/core-library/loan"
	"bitbucket.org/kawancicil/core-library/logger"
	"bitbucket.org/kawancicil/core-library/request"
	"bitbucket.org/kawancicil/core-library/response"
	"bitbucket.org/kawancicil/core-library/rest_api"
	"bitbucket.org/kawancicil/core-library/transaction"
)

type CoreModule struct {
	/*todo fill this block if need connection with another interface*/

	Request     request.Request
	Response    response.Response
	Common      common.Common
	Loan        loan.Loan
	Transaction transaction.Transaction
	RestApi     rest_api.RestApi
	Oss         oss.Oss
	Privy       privy.Privy
	Logger      logger.Logger
	RabbitMQ    rabbitMQ.RabbitMQ
	Slack       slack.Slack
}

func InitCoreModule() *CoreModule {
	loggerUtils := logger.InitLoggerUtils()
	rabbitMQUtils := rabbitMQ.InitRabbitMQUtils()
	commonUtils := common.InitCommonUtils()
	utilsHttp := rest_api.NewUtilHttpRequest()
	restApiUtils := rest_api.InitRestApiUtils(utilsHttp)
	requestUtils := request.InitRequestUtils()
	responseUtils := response.InitResponseUtils(loggerUtils, rabbitMQUtils)
	loanUtils := loan.InitLoanUtils(commonUtils)
	transactionUtils := transaction.InitTransactionUtils()
	ossUtils := oss.InitOssUtils()
	privyUtils := privy.InitPrivyUtils(rest_api.InitRestApiUtils(utilsHttp))
	slackUtils := slack.InitSlackUtils(restApiUtils)

	return &CoreModule{
		Request:     requestUtils,
		Response:    responseUtils,
		Common:      commonUtils,
		Loan:        loanUtils,
		Transaction: transactionUtils,
		RestApi:     restApiUtils,
		Oss:         ossUtils,
		Privy:       privyUtils,
		Logger:      loggerUtils,
		RabbitMQ:    rabbitMQUtils,
		Slack:       slackUtils,
	}
}
