import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/user/login',
    method: 'post',
    data
  })
}

export function Reg(data) {
  return request({
    url: '/user/reg',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/user/info',
    method: 'get',
    params: {token}
  })
}

export function userList() {
  return request({
    url: '/user/list',
    method: 'get'
  })
}
