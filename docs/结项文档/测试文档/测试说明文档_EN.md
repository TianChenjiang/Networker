

# Test document

## 1. Introduction to Testing

Black box and white box testing for backend server code

## 2. Test purpose

- Evaluation of software quality by analysis of test results
- Evaluate whether the test execution and test plan are consistent
- Analyze system defects and provide suggestions for fixing and preventing bugs

## 3. Test Tools

- Postman for functional testing of services
- Jmeter for interface stress testing

## 4. Test process

### 4.1 Functional Test

#### User Management

| Features                                     | Basic Requirements                                           | Test Situation | By (Y/N) |
| -------------------------------------------- | ------------------------------------------------------------ | -------------- | -------- |
| User login                                   | After the user login authentication succeeds, the correct token is returned. And the subsequent authorization can be provided. | Implemented    | Y        |
| User Information Acquisition                 | User access to personal information, also need to pass stress test | Implemented    | Y        |
| Modify User Information                      | Ability to modify user information normally, and validate via form | Implemented    | Y        |
| User tagged company                          | User's company watchlist added normally                      | Implemented    | Y        |
| Users get attention to the list of companies | Can return the list of companies that the user is concerned with, and the paging is normal | Implemented    | Y        |



#### Company Management

| Features                | Basic Requirements                                           | Test Situation | By (Y/N) |
| ----------------------- | ------------------------------------------------------------ | -------------- | -------- |
| Get company information | Get information on Tushare and display it normally           | Implemented    | Y        |
| Company Fuzzy Search    | Ability to perform fuzzy search of company name based on keywords entered by users, and paging is normal | Implemented    | Y        |

#### Stock Information Management

| Features              | Basic Requirements                                     | Test Situation | By (Y/N) |
| --------------------- | ------------------------------------------------------ | -------------- | -------- |
| Get Stock Information | Get the information on Tushare and display it normally | Implemented    | Y        |

#### Company Financial Information

| Features | Basic Requirements | Test Situation | By (Y/N) |
| ---------------- | -------------------------------- ------------------------ | -------- | --------- |
Get company financial information | Get the information on Tushare according to the company's stock code, and display it normally | Realized | Y |

#### Stock price closing line warning probability

| Features                       | Basic Requirements                                           | Test Situation | By (Y/N) |
| ------------------------------ | ------------------------------------------------------------ | -------------- | -------- |
| Get the probability of a burst | According to the company's stock code and the level of the warehouse line, you should return the probability of a burst (30 days) | Realized       | Y        |

### 4.2 Performance Test

#### Test parameter description

| Parameter Name             | Value            |
| -------------------------- | ---------------- |
| Script Cycles              | 6000             |
| Number of concurrent users | 5/10/50/100/200  |
| Number of real clients     | 1                |
| Script Recording Method    | Automatic        |
| Analog Route Type          | 10/100M Ethernet |

#### Performance test specific results

##### effectiveness

Average response time (s):

| Number of concurrent users | Login | User information acquisition | Get stock information | Get company information |
| -------------------------- | ----- | ---------------------------- | --------------------- | ----------------------- |
| 5                          | 0.195 | 0.258                        | 0.421                 | 0.033                   |
| 10                         | 0.294 | 0.213                        | 0.293                 | 0.057                   |
| 50                         | 2.502 | 1.400                        | 1.501                 | 2.714                   |
| 100                        | 5.228 | 3.775                        | 6.841                 | 1.812                   |

### Performance Analysis

- Backend performance test analysis

According to the results of the JMeter aggregation report, each backend api performs normally. When running 100 threads (users), the average time is 9.862s, and the error rate is 0% in the case of low concurrent (5-100).

- Web performance test analysis

According to the results of the JMeter aggregation report, when running 200 threads (virtual users), when the number of https requests is about 460, the average response time is 1.452s and the error rate is only 0.07%, which is in line with the expected performance results.

![](https://i.loli.net/2019/09/14/sx8vbwlid3O1LSE.png)