const functions = require('firebase-functions')
const { DialogflowApp } = require('actions-on-google')

const WELCOME_INTENT = 'input.welcome'

function welcome(app) {

}

const actionMap = new Map()

exports.googleHomeSMS = functions.https.onRequest((request, response) => {
    const app = new DialogflowApp({ request, response })
    console.log(`Request Headers: ${JSON.stringify(request.headers)}`)
    console.log(`Request Body: ${JSON.stringify(request.body)}`)
})
