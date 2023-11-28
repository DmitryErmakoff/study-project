/**
 * plugins/index.js
 *
 * Automatically included in `./src/main.js`
 */

// Plugins
import vuetify from './vuetify'
import vuex from '../store'
import router from '../router'

export function registerPlugins (app) {
  app
    .use(vuetify)
    .use(router)
    .use(vuex)
}
