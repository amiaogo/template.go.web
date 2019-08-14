'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  // 代理映射表
  PROXY_TABLE: {
    // 后端API地址映射
    '/api/v1/*': 'http://127.0.0.1:8080',
  },
  // 主机名
  HOST: '127.0.0.1',
  // 端口号
  PORT: '8081',
  // 是否使用代码校验
  USE_ESLINT: false,
})
