const functions = require('firebase-functions')
const { DialogflowApp } = require('actions-on-google')
const axios = require('axios')

const SEND_TEXT = 'text.send'
const NAME_ARGUMENT = 'name'
const MESSAGE_ARGUMENT = 'message'

var contacts = {
    "Michael": "18888888888",
    "Kierstyn": "18888888888",
    "Josh": "18888888888",
    "Wilber": "18888888888"
}
var From = "8888888888"
var actualNum = "88888888888"

function sendText(app) {
    const homePrefix = "Hey this is Michael texting from my Google Home: "
    const homeSuffix = `. Text me back at: ${actualNum}`
    const Body = homePrefix + app.getArgument(MESSAGE_ARGUMENT) + homeSuffix
    const toWho = app.getArgument(NAME_ARGUMENT)
    if (!(toWho in contacts)) {
        app.ask(`Sorry I couldn't find ${To} in your contacts, who would you like to send a text to?`, 
        ['Who do you want to send a text to?','Who should I send this text to?', 
        'Goodbye!'])
    }
    console.log(`Body: ${Body}, toWho: ${toWho}`)
    const To = contacts[toWho]
    let instance = axios.create({
        baseURL: 'http://EXTERNAL_IP:8081/api/',
        headers: {'Content-Type': 'application/json'}
    })
    instance.post('google-home-sms', { To, Body, From, })
        .then(function(response) {
            app.tell(`Text sent succesfully to ${toWho}`)
        })
        .catch(function(err) {
            console.log(err)
            app.tell(`Unable to send text to ${toWho}`)

        })
}

const actionMap = new Map()
actionMap.set(SEND_TEXT, sendText)


exports.googleHomeSMS = functions.https.onRequest((request, response) => {
    const app = new DialogflowApp({ request, response })
    console.log(`Request Headers: ${JSON.stringify(request.headers)}`)
    console.log(`Request Body: ${JSON.stringify(request.body)}`)
    app.handleRequest(actionMap)
})
