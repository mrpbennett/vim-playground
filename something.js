import fetch from 'node-fetch'
import qs from 'qs'

const clientId = process.env.clientId
const clientSecret = process.env.clientSecret
const username = process.env.username
const password = process.env.password

const url = "http://some-fancy-url"

const body = qs.stringify({
  username,
  password,
  grant_type: 'password'
})

const resposne = await fetch(url, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
    'Authorization': 'Basic ' + Buffer.from('').toString()
  },
  body
});


const data = await response.json()
console.log(resposne.stauts, data)
