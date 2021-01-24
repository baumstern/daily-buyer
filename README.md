# daily-buyer

Buy cryptocurrency using Open API from [Upbit](https://upbit.com/).

Currently, when the program is executed each time, it creates an order for BTC and ETH at a market price equivalent to KRW 5,000.

*Writing go program is in progress ...*

## Quick Start

Install dependencies:
```
npm install
```

Run to buy BTC and ETH once:
```
npm start
```

NOTE: You should provide personal access token to environment variable (`UPBIT_OPEN_API_ACCESS_KEY` and `UPBIT_OPEN_API_SECRET_KEY`).

Token can be issued [here](https://upbit.com/mypage/open_api_management).


# Schedule an order by daily

Here is a cron job example that executes at daily midnight (UTC):
```
0 0 * * * node app.js
```