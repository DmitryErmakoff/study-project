<template>
  <v-app class="app">
    <v-navigation-drawer v-model="drawer" temporary location="right" class="pa-2">
      <v-text-field hide-details type="login"    variant="outlined" class="my-2 loginField" label="Login"></v-text-field>
      <v-text-field hide-details type="password" variant="outlined" class="mb-5 loginField" label="Password"></v-text-field>
      <v-row class="mx-0 justify-space-around">
        <v-btn class="">login</v-btn>
        <v-btn class="">register</v-btn>
      </v-row>
    </v-navigation-drawer>
    <v-app-bar density="comfortable" flat color="#354f52">
      <v-app-bar-title>
        <v-btn>
            nash magazinchik
        </v-btn>
        
      </v-app-bar-title>
      <v-spacer></v-spacer>
      <v-text-field hide-details density="compact" flat rounded placeholder="Search" variant="solo-filled" class="mr-2"></v-text-field>
      <v-btn icon class="mr-2"><v-icon>mdi-magnify</v-icon></v-btn>
      <v-btn icon class="mr-2">
        <v-badge :content="3" color="red">
          <v-icon>mdi-cart</v-icon>
        </v-badge>
      </v-btn>

      <v-btn icon @click="toggleDrawer"><v-icon>mdi-account</v-icon></v-btn>
    </v-app-bar>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>

export default {
  data: () => ({
    drawer: false
  }),
  methods: {
    toggleDrawer() {
      this.drawer = !this.drawer
    }
  },
}
</script>

<style>
.drawerWrapper {
  padding: 5px;
}

.loginField {
  width: 100%
}

</style>