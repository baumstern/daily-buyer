const request = require('request')
const { v4: uuidv4 } = require('uuid')
const crypto = require('crypto')
const jwt = require('jsonwebtoken')
const queryEncode = require("querystring").encode

const access_key = process.env.UPBIT_OPEN_API_ACCESS_KEY
const secret_key = process.env.UPBIT_OPEN_API_SECRET_KEY
const server_url = process.env.UPBIT_OPEN_API_SERVER_URL || 'https://api.upbit.com'


function buy(market, price) {
  const body = {
    market: market,
    side: 'bid',
    price: price,
    ord_type: 'price',
  }

  const query = queryEncode(body)

  const hash = crypto.createHash('sha512')
  const queryHash = hash.update(query, 'utf-8').digest('hex')

  const payload = {
    access_key: access_key,
    nonce: uuidv4(),
  }

  const jwtToken = jwt.sign(payload, secret_key)

  const options = {
      method: "POST",
      url: server_url + "/v1/orders",
      headers: {Authorization: `Bearer ${jwtToken}`},
      json: body
  }

  request(options, (error, response, body) => {
      if (error) throw new Error(error)
      console.log(body)
  })
}

buy('KRW-BTC', 5000)
buy('KRW-ETH', 5000)