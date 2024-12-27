import '../scss/styles.scss'
import Keycloak from 'keycloak-js'
import * as bootstrap from 'bootstrap'

window.addEventListener('DOMContentLoaded', async function () {
    var apiUrl = document.location.protocol + '//' + document.location.hostname + ':8080';

    const keycloak = new Keycloak({
        url: 'http://localhost:9100',
        realm: 'krakend',
        clientId: 'playground',
        flow: 'implicit'
    });

    try {
        await keycloak.init({ onLoad: 'check-sso', flow: 'implicit', checkLoginIframe: false });
    } catch (error) {
        // Skip initialization error
    }

    const loggedInAlert = this.document.getElementById('__alert-logged-in')
    const notLoggedInAlert = this.document.getElementById('__alert-not-logged-in')
    const loggedInMenu = this.document.getElementById('__menu-logged-in')
    const notLoggedInMenu = this.document.getElementById('__menu-not-logged-in')
    const loginBtn = this.document.getElementById('__menu-login')
    const logoutBtn = this.document.getElementById('__menu-logout')
    const homeBtn = this.document.getElementById('__menu-home')
    const profileBtn = this.document.getElementById('__menu-profile')
    const publicCallBtn = this.document.getElementById('__btn-ping-public')
    const privateCallBtn = this.document.getElementById('__btn-ping-private')
    const callResponseSection = this.document.getElementById('__call-response')
    const pingSection = this.document.getElementById('__section-ping')
    const profileSection = this.document.getElementById('__section-profile')
    loginBtn.addEventListener('click', async () => {
        await keycloak.login()
    })
    logoutBtn.addEventListener('click', async () => {
        await keycloak.logout()
    })
    homeBtn.addEventListener('click', () => {
        show(pingSection)
        hide(profileSection)
    })
    profileBtn.addEventListener('click', () => {
        show(profileSection)
        hide(pingSection)
    })
    publicCallBtn.addEventListener('click', async () => {
        const responseBlock = callResponseSection.querySelector('pre > code')
        try {
            hide(callResponseSection)
            responseBlock.innerHTML = ''

            const req = await fetch(apiUrl + '/public')
            if (req.ok) {
                const res = await req.json()
                responseBlock.innerHTML = JSON.stringify(res, null, 4)
            } else {
                responseBlock.innerHTML = `Error (${req.status}): ${await req.text()}`
            }
        } catch (err) {
            responseBlock.innerHTML = 'Error: ' + err
        } finally {
            show(callResponseSection)
        }

        show(callResponseSection)
    })
    privateCallBtn.addEventListener('click', async () => {
        const responseBlock = callResponseSection.querySelector('pre > code')
        try {
            hide(callResponseSection)
            responseBlock.innerHTML = ''

            const req = await fetch(apiUrl + '/private/moderate', {
                headers: {
                    'Authorization': `Bearer ${keycloak.token}`
                }
            })
            if (req.ok) {
                const res = await req.json()
                responseBlock.innerHTML = JSON.stringify(res, null, 4)
            } else {
                responseBlock.innerHTML = `Error (${req.status}): ${await req.text()}`
            }

        } catch (err) {
            responseBlock.innerHTML = 'Error: ' + err
        } finally {
            show(callResponseSection)
        }

        show(callResponseSection)
    })

    if (keycloak.authenticated) {
        show(loggedInMenu)
        show(loggedInAlert)
        show(pingSection)
        hide(notLoggedInMenu)
        hide(notLoggedInAlert)

        privateCallBtn.classList.remove('disabled')
        const info = await keycloak.loadUserInfo()
        this.document.getElementById('__menu-profile-username').innerHTML = info.name
        const table = this.document.querySelector('#__section-profile > * > table > tbody')

        const tableMapping = { sub: "ID", email: "Email", name: "Name", preferred_username: "Username" }
        for (const [k, v] of Object.entries(tableMapping)) {
            const row = this.document.createElement('tr')
            const cellA = this.document.createElement('td')
            const cellB = this.document.createElement('td')
            cellA.innerHTML = v
            cellB.innerHTML = info[k]

            row.appendChild(cellA)
            row.appendChild(cellB)
            table.appendChild(row)
        }
        this.document.querySelector('#__section-profile > pre > code').innerHTML = JSON.stringify(info, null, 4)
    } else {
        show(notLoggedInMenu)
        show(notLoggedInAlert)
        hide(loggedInMenu)
        hide(loggedInAlert)
    }
})

function hide(el) {
    if (!el.classList.contains('visually-hidden')) {
        el.classList.add('visually-hidden')
    }
}

function show(el) {
    if (el.classList.contains('visually-hidden')) {
        el.classList.remove('visually-hidden')
    }
}